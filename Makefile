all: protoc-go
  
clean:
	rm -f pbgen/pbgen pkg/*.go

pbgen: clean
	cd pbgen && go build .
	pbgen/pbgen

fmt: pbgen
	for f in $$(find pb -type f); do \
		prototool format -w $$f; \
	done

protoc-go: fmt
	for f in $$(find pb -type f -name *.proto); do \
		protoc -I . \
			--go_out=plugins=grpc:$$GOPATH/src \
			$$f ;\
	done
	sed -i "" -e "s/,omitempty//g" pkg/*.go

.PHONY: clean pbgen fmt protoc-go
