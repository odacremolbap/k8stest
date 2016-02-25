# build vars
BIN_NAME=k8stest
DOCKER_PATH=./docker
GO_PACKAGE := github.com/odacremolbap/k8stest

# injected vars
VERSION := $(shell cat VERSION)
DATE= $(shell date +%FT%T)
GIT_COMMIT= $(shell git rev-parse --short HEAD)
LDFLAG_VER := -X $(GO_PACKAGE)/appinfo.Version=$(VERSION)
LDFLAG_DATE := -X $(GO_PACKAGE)/appinfo.Date=$(DATE)
LDFLAG_GIT := -X $(GO_PACKAGE)/appinfo.GitCommit=$(GIT_COMMIT)
LDFLAG_STATIC :=-extldflags "-static"
GOOSES := darwin freebsd linux windows
GOARCHS := amd64 386

define build
	mkdir -p ./releases/$(1)/$(2);
	GOOS=$(1) GOARCH=$(2) go build --ldflags '$(LDFLAG_VER) $(LDFLAG_DATE) $(LDFLAG_GIT) $(LDFLAG_STATIC)' -o ./releases/$(1)/$(2)/$(BIN_NAME) main.go;
endef

default: pushcontainer

compile_linux: *.go VERSION
	$(call build,linux,amd64)

cross_compile: *.go VERSION
	$(foreach GOARCH,$(GOARCHS),$(foreach GOOS,$(GOOSES),$(call build,$(GOOS),$(GOARCH))))

buildcontainer: compile_linux
	cp ./releases/linux/amd64/$(BIN_NAME) ./docker/k8stest
	pushd ./docker && \
	docker build -t pmercado/k8stest . && \
	popd

pushcontainer: buildcontainer
	docker push pmercado/k8stest
