// Copyright 2023 The olive Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gateway

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/olive-io/olive/runner/internal/codec"
	"github.com/olive-io/olive/runner/internal/codec/jsonrpc"
	"github.com/olive-io/olive/runner/internal/codec/protorpc"
	"github.com/olive-io/olive/runner/internal/context/metadata"
	"github.com/olive-io/olive/runner/internal/qson"

	jsonpatch "github.com/evanphx/json-patch/v5"
	json "github.com/json-iterator/go"
	"github.com/oxtoacart/bpool"
)

var (
	// supported json codecs
	jsonCodecs = []string{
		"application/grpc+json",
		"application/json",
		"application/json-rpc",
	}

	// supported proto codecs
	protoCodecs = []string{
		"application/grpc",
		"application/grpc+proto",
		"application/proto",
		"application/protobuf",
		"application/proto-rpc",
		"application/octet-stream",
	}

	// supported multipart/form-data codecs
	dataCodecs = []string{
		"multipart/form-data",
	}

	bufferPool = bpool.NewSizedBufferPool(1024, 8)
)

type RawMessage json.RawMessage

// MarshalJSON returns m as the JSON encoding of m.
func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return fmt.Errorf("RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

type buffer struct {
	io.ReadCloser
}

func (b *buffer) Write(_ []byte) (int, error) {
	return 0, nil
}

type Message struct {
	data []byte
}

func (m *Message) ProtoMessage() {}

func (m *Message) Reset() {
	*m = Message{}
}

func (m *Message) String() string {
	return string(m.data)
}

func (m *Message) Marshal() ([]byte, error) {
	return m.data, nil
}

func (m *Message) Unmarshal(data []byte) error {
	m.data = data
	return nil
}

func NewMessage(data []byte) *Message {
	return &Message{data}
}

// requestPayload takes a *http.Request.
// If the request is a GET the query string parameters are extracted and marshaled to JSON and the raw bytes are returned.
// If the request method is a POST the request body is read and returned
func requestPayload(r *http.Request) ([]byte, error) {
	var err error

	// we have to decode json-rpc and proto-rpc because we suck
	// well actually because there's no proxy codec right now
	ct := r.Header.Get("Content-Type")
	switch {
	case strings.Contains(ct, "application/json-rpc"):
		msg := codec.Message{
			Type:   codec.Request,
			Header: make(map[string]string),
		}

		c := jsonrpc.NewCodec(&buffer{r.Body})
		if err = c.ReadHeader(&msg, codec.Request); err != nil {
			return nil, err
		}
		var raw RawMessage
		if err = c.ReadBody(&raw); err != nil {
			return nil, err
		}
		return raw, nil
	case strings.Contains(ct, "application/proto-rpc"), strings.Contains(ct, "application/octet-stream"):
		msg := codec.Message{
			Type:   codec.Request,
			Header: make(map[string]string),
		}
		c := protorpc.NewCodec(&buffer{r.Body})
		if err = c.ReadHeader(&msg, codec.Request); err != nil {
			return nil, err
		}
		var raw Message
		if err = c.ReadBody(&raw); err != nil {
			return nil, err
		}
		return raw.Marshal()
	case strings.Contains(ct, "application/x-www-form-urlencoded"):
		// generate a new set of values from the form
		vals := make(map[string]string)
		for key, values := range r.PostForm {
			vals[key] = strings.Join(values, ",")
		}
		for key, values := range r.URL.Query() {
			vv, ok := vals[key]
			if !ok {
				vals[key] = strings.Join(values, ",")
			} else {
				vals[key] = vv + "," + strings.Join(values, ",")
			}
		}

		// marshal
		return json.Marshal(vals)
		// TODO: application/grpc
	}

	// otherwise as per usual
	rctx := r.Context()
	// dont user metadata.FromContext as it mangles names
	md, ok := metadata.FromContext(rctx)
	if !ok {
		md = make(map[string]string)
	}

	// allocate maximum
	matches := make(map[string]interface{}, len(md))
	bodydst := ""

	// get fields from url path
	for k, v := range md {
		k = strings.ToLower(k)
		// filter own keys
		if strings.HasPrefix(k, "x-api-field-") {
			matches[strings.TrimPrefix(k, "x-api-field-")] = v
			delete(md, k)
		} else if k == "x-api-body" {
			bodydst = v
			delete(md, k)
		}
	}

	// get fields from url values
	if len(r.URL.RawQuery) > 0 {
		umd := make(map[string]interface{})
		err = qson.Unmarshal(&umd, r.URL.RawQuery)
		if err != nil {
			return nil, err
		}
		for k, v := range umd {
			matches[k] = v
		}
	}

	// restore context without fields
	*r = *r.Clone(metadata.NewContext(rctx, md))

	// map of all fields
	fields := nestField(matches)
	pathbuf := []byte("{}")
	if len(fields) > 0 {
		pathbuf, err = json.Marshal(fields)
		if err != nil {
			return nil, err
		}
	}

	urlbuf := []byte("{}")
	out, err := jsonpatch.MergeMergePatches(urlbuf, pathbuf)
	if err != nil {
		return nil, err
	}

	switch r.Method {
	case "GET":
		// empty response
		if strings.Contains(ct, "application/json") && string(out) == "{}" {
			return out, nil
		} else if string(out) == "{}" && !strings.Contains(ct, "application/json") {
			return []byte{}, nil
		}
		return out, nil
	case "PATCH", "POST", "PUT", "DELETE":
		bodybuf := []byte("{}")
		buf := bufferPool.Get()
		defer bufferPool.Put(buf)
		if _, err = buf.ReadFrom(r.Body); err != nil {
			return nil, err
		}
		if b := buf.Bytes(); len(b) > 0 {
			bodybuf = b
		}
		if bodydst == "" || bodydst == "*" {
			if out, err = jsonpatch.MergeMergePatches(out, bodybuf); err == nil {
				return out, nil
			}
		}
		var jsonbody map[string]interface{}
		if json.Valid(bodybuf) {
			if err = json.Unmarshal(bodybuf, &jsonbody); err != nil {
				return nil, err
			}
		}
		dstmap := make(map[string]interface{})
		ps := strings.Split(bodydst, ".")
		if len(ps) == 1 {
			if jsonbody != nil {
				dstmap[ps[0]] = jsonbody
			} else {
				// old unexpected behaviour
				dstmap[ps[0]] = bodybuf
			}
		} else {
			em := make(map[string]interface{})
			if jsonbody != nil {
				em[ps[len(ps)-1]] = jsonbody
			} else {
				// old unexpected behaviour
				em[ps[len(ps)-1]] = bodybuf
			}
			for i := len(ps) - 2; i > 0; i-- {
				nm := make(map[string]interface{})
				nm[ps[i]] = em
				em = nm
			}
			dstmap[ps[0]] = em
		}

		bodyout, err := json.Marshal(dstmap)
		if err != nil {
			return nil, err
		}

		if out, err = jsonpatch.MergeMergePatches(out, bodyout); err == nil {
			return out, nil
		}

		//fallback to previous unknown behaviour
		return bodybuf, nil
	}

	return []byte{}, nil
}

func nestField(matches map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})

	for k, v := range matches {
		ps := strings.Split(k, ".")
		if len(ps) == 1 {
			req[k] = v
			continue
		}

		em := req
		for i := 0; i < len(ps)-1; i++ {
			if vm, ok := em[ps[i]]; !ok {
				vmm := make(map[string]interface{})
				em[ps[i]] = vmm
				em = vmm
			} else {
				vv, ok := vm.(map[string]interface{})
				if !ok {
					em = make(map[string]interface{})
				} else {
					em = vv
				}
			}
		}

		em[ps[len(ps)-1]] = v
	}
	return req
}
