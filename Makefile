NAME=olive
IMAGE_NAME=olive-io/$(NAME)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
GIT_VERSION=github.com/olive-io/olive/pkg/version
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
LDFLAGS=-X $(GIT_VERSION).GitCommit=$(GIT_COMMIT) -X $(GIT_VERSION).GitTag=$(GIT_TAG) -X $(GIT_VERSION).BuildDate=$(BUILD_DATE)
IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)
ROOT=github.com/olive-io/olive

all: build

vendor:
	go mod vendor

test-coverage:
	go test ./... -bench=. -coverage

lint:
	golint -set_exit_status ./..

install:

genclients:
	go run ./tools/olive-runtime-gen \
		-g client-gen \
		-g deepcopy-gen \
		-g go-to-protobuf \
		-g informer-gen \
		-g lister-gen \
		-g openapi-gen \
		--module "github.com/olive-io/olive" \
		--versions "github.com/olive-io/olive/apis/discovery"

	go-to-protobuf --apimachinery-packages "+k8s.io/apimachinery/pkg/util/intstr,+k8s.io/apimachinery/pkg/api/resource,+k8s.io/apimachinery/pkg/runtime/schema,+k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1" \
		--output-base $(GOPATH)/src \
 		--go-header-file hack/boilerplate.go.txt \
 		--proto-import=vendor --proto-import=vendor/k8s.io/kubernetes/third_party/protobuf \
 		--packages github.com/olive-io/olive/apis/discovery


generate:
	cd $(GOPATH)/src && \
	protoc --go_out=. github.com/olive-io/olive/apis/pb/discovery/discovery.proto && \
	protoc --go_out=. github.com/olive-io/olive/apis/pb/discovery/activity.proto && \
	protoc --go_out=. github.com/olive-io/olive/apis/pb/auth/auth.proto && \
	protoc -I. -I github.com/googleapis/googleapis --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-olive_out=. --go-olive_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --openapiv2_out=. --openapiv2_opt use_go_templates=true github.com/olive-io/olive/apis/pb/gateway/rpc.proto && \
	protoc -I. -I github.com/googleapis/googleapis --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative github.com/olive-io/olive/apis/pb/olive/internal.proto && \
	protoc -I. -I github.com/googleapis/googleapis --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative github.com/olive-io/olive/apis/pb/olive/raft.proto && \
	protoc -I. -I github.com/googleapis/googleapis --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative github.com/olive-io/olive/apis/pb/olive/rpc.proto

	goimports -w api/*/**.go
	rm -fr api/*/**swagger.json api/*/**.bak api/*/**_olive.pb.go

docker:


vet:
	go vet ./...

test: vet
	go test -v ./...

clean:
	rm -fr ./_output

