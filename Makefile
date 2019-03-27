GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOPATH?=$$(go env GOPATH)

fmt:
	gofmt -w $(GOFMT_FILES)

imports: fmt
	$(GOPATH)/bin/goimports -w $(GOFMT_FILES)
