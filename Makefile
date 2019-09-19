
build:
	go get github.com/golang/protobuf/protoc-gen-go

	protoc --go_out=. OpenAPIv2/OpenAPIv2.proto
	protoc --go_out=. OpenAPIv3/OpenAPIv3.proto
	protoc --go_out=. discovery/discovery.proto
	protoc --go_out=. plugins/plugin.proto
	protoc --go_out=. extensions/extension.proto
	protoc --go_out=. surface/surface.proto
	go get ./...
	go install ./...
	cd extensions/sample; make

