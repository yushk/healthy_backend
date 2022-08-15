.PHONY: build
build:
	echo "nothing to do"

tidy:
	go mod tidy

gen: client prt
	@echo pkg gen ok

client: clean

prt:
	protoc -I/usr/local/include \
		-Ipb \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		pb/*.proto && \
		sed -i -E 's/json:("[^"]+),omitempty"/json:\1" bson:\1"/' pb/*.pb.go

.PHONY: clean
clean:
	rm -rf ./pb/*.pb.go