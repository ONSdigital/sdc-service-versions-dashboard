# Cross-compilation values.
ARCH=amd64
OS_LINUX=linux
OS_MAC=darwin

# Output directory structures.
BUILD=build
LINUX_BUILD_ARCH=$(BUILD)/$(OS_LINUX)-$(ARCH)
MAC_BUILD_ARCH=$(BUILD)/$(OS_MAC)-$(ARCH)

# Cross-compile the binary for Linux and macOS.
build: clean
	GOOS=$(OS_LINUX) GOARCH=$(ARCH) go build -o $(LINUX_BUILD_ARCH)/bin/versions-dashboard
	GOOS=$(OS_MAC) GOARCH=$(ARCH) go build -o $(MAC_BUILD_ARCH)/bin/versions-dashboard
	cp ./index.html $(LINUX_BUILD_ARCH)/bin
	cp ./index.html $(MAC_BUILD_ARCH)/bin

# Remove the build directory tree.
clean:
	if [ -d $(BUILD) ]; then rm -r $(BUILD); fi;