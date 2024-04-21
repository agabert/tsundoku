
tmp:
	mkdir -pv tmp

EXECUTABLE = tmp/$(PROJECT)

build: $(TOOLCHAIN) src $(EXECUTABLE)

GOOS = linux
# or freebsd

GO = ../$(TOOLCHAIN)

BUILD = mkdir -pv tmp && \
		cd src && \
		$(GO) get && \
		$(GO) fmt && \
		CGO_ENABLED=0 \
		GOOS=$(GOOS) \
		$(GO) build -a -installsuffix cgo -o "../$(EXECUTABLE)" . && \
		file "../$(EXECUTABLE)" && \
		ls -ali "../$(EXECUTABLE)"

# when you run make clean but did not change the code
$(EXECUTABLE):
	$(BUILD)

# anytime the code gets updated this should trigger
src:
	$(BUILD)
