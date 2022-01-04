BINARY_OUT := imgur.out

# executables
GO         := go
GORELEASER := goreleaser

# build flags
BUILD_FLAGS := -v

build:
	$(GO) mod tidy
	$(GO) build ${BUILD_FLAGS} -o ${BINARY_OUT} ./cmd/imgur

snapshot:
	$(GORELEASER) --snapshot --rm-dist

clean:
	$(GO) clean
	rm -f ${BINARY_OUT}

install:
	cp ${BINARY_OUT} /usr/bin/imgur

uninstall:
	rm -f /usr/bin/imgur
