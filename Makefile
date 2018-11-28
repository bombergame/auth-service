all: build

generate:
	go generate ./...

prepare:
	protoc -I services/grpc/ services/grpc/service.proto --go_out=plugins=grpc:services/grpc

build:
	go build -v -o ./_build/service .

test_units:
	mkdir -p _test
	go test -run 'Unit' -v -race ./...
	go test -run 'Unit' -v -covermode=count -coverprofile=./_test/coverage.out ./...

clean:
	rm -rf ./_build
	rm -rf ./_test
