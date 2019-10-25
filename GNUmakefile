
PKG_PREFIX=github.com/ibraimgm
APP_NAME=goobi
PKG_NAME=$(PKG_PREFIX)/$(APP_NAME)
PKG_EXE_NAME=$(PKG_NAME)/cmd/$(APP_NAME)

GO=go

UNAME = $(shell uname)

ifneq (,$(findstring MINGW,$(UNAME)))
	EXE = .exe
else ifndef $(UNAME)
	EXE = .exe
endif

APP_EXE=$(APP_NAME)$(EXE)

all: $(APP_EXE)

$(APP_EXE): 
	$(GO) build -o $(APP_EXE) $(PKG_EXE_NAME)

install: 
	$(GO) install $(PKG_EXE_NAME)

build: clean $(APP_EXE)

build-docker: 
	GOOS=linux GOARCH=amd64 $(GO) build -o $(APP_NAME) $(PKG_EXE_NAME)

docker: clean build-docker
	docker build -t goobi .

check: $(APP_EXE)
	go test ./...

clean:
	rm -f $(APP_EXE)
	rm -f $(APP_NAME)

.PHONY: all build check clean install build-docker docker
