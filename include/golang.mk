
GO_TOOLCHAIN_VERSION = $(VERSION)

TOOLCHAIN = toolchain/go/bin/go

PACKAGE = go$(GO_TOOLCHAIN_VERSION).linux-amd64.tar.gz

DOWNLOAD = wget -SO- "https://dl.google.com/go/$(PACKAGE)" | tar xzvpf -

$(TOOLCHAIN):
	mkdir -pv toolchain
	cd toolchain && $(DOWNLOAD)
	touch "$(@)"
	$(@) version | grep -- "$(GO_TOOLCHAIN_VERSION)"
