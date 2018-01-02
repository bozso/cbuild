PREFIX ?= /usr/local
EXEC   ?= cbuild

all : cbuild test

test :
	cd test && $(MAKE)

install: cbuild
	cp cbuild $(PREFIX)/bin/$(EXEC)

dist: cbuild
	$(MAKE) -f cbuild.mk CLEAN_cbuild
	cd test && $(MAKE) CLEAN_test

include cbuild.mk
.PHONY: test
