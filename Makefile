
PROJECT = tsundoku

VERSION = 1.22.2

all: build

include include/golang.mk
include include/build.mk
include include/ci.mk

clean:
	rm -rvf tmp

distclean: clean
	rm -rfv toolchain
	rm -rfv -- "$(TRUNK)"
