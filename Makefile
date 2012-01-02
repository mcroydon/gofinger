all:
	$(MAKE) -C src/cmd/finger
	$(MAKE) -C src/cmd/fingerd

install:
	$(MAKE) -C src/cmd/finger install
	$(MAKE) -C src/cmd/fingerd install

gofmt:
	gofmt -w src/cmd/finger/finger.go
	gofmt -w src/cmd/fingerd/fingerd.go

.PHONY: all
