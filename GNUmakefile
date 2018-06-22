
PKG_PREFIX=github.com/ibraimgm
APP_NAME=goobi
PKG_NAME=$(PKG_PREFIX)/$(APP_NAME)/cmd/$(APP_NAME)

DEP=dep
GO=go

UNAME = $(shell uname)

ifneq (,$(findstring MINGW,$(UNAME)))
	EXE = .exe
else ifndef $(UNAME)
	EXE = .exe
endif

APP_EXE=$(APP_NAME)$(EXE)

all: $(APP_EXE)

$(APP_EXE): deps
	$(GO) build -o $(APP_EXE) $(PKG_NAME)

install: deps
	$(GO) install $(PKG_NAME)

build: clean $(APP_EXE)

deps: Gopkg.lock Gopkg.toml
	$(DEP) ensure

check:
	echo Nao implementado

clean:
	rm -f $(APP_NAME)

.PHONY: all build deps check clean install
