all: build

prepare:
	easyjson -all ./services/rest/models.go
	protoc -I services/grpc/ services/grpc/service.proto --go_out=plugins=grpc:services/grpc
	protoc -I clients/profiles-service/grpc/ clients/profiles-service/grpc/service.proto \
		--go_out=plugins=grpc:clients/profiles-service/grpc

build:
	go build -v -o ./_build/service .

test_units:
	mkdir -p _test
	go test -run 'Unit' -v -race ./...
	go test -run 'Unit' -v -covermode=count -coverprofile=./_test/coverage.out ./...

clean:
	rm -rf ./_build
	rm -rf ./_test
