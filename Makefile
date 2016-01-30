.PHONY: test

test:
	go test -v -race ./...

deps:
	godep save -r ./...
