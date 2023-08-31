PROJECT_NAME := Pulumi Vco Resource Provider

PACK             := vco
PACKDIR          := sdk
PROJECT          := github.com/fabianv-cloud/pulumi-vco-native
NODE_MODULE_NAME := @fabianv-cloud/vco
NUGET_PKG_NAME   := Fabianv-cloud.Vco

PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         ?= $(shell pulumictl get version)
PROVIDER_PATH   := provider
VERSION_PATH     := ${PROVIDER_PATH}/cmd/main.Version

SCHEMA_FILE     := provider/cmd/pulumi-resource-vco/schema.json
GOPATH			:= $(shell go env GOPATH)

WORKING_DIR     := $(shell pwd)
TESTPARALLELISM := 4

ensure::
	cd provider && go mod tidy
	cd sdk && go mod tidy
	cd tests && go mod tidy

codegen::
	(cd provider && VERSION=${VERSION} go generate cmd/${PROVIDER}/main.go)
	(cd provider && go build -o $(WORKING_DIR)/bin/${CODEGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/$(CODEGEN))
	$(WORKING_DIR)/bin/${CODEGEN} $(SCHEMA_FILE) --version ${VERSION}

provider::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

provider_debug::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -gcflags="all=-N -l" -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

test_provider::
	cd provider/pkg && go test -short -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM} ./...

dotnet_sdk:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
dotnet_sdk::
	rm -rf sdk/dotnet
	chmod +x $(WORKING_DIR)/bin/$(PROVIDER)
	pulumi package gen-sdk --language dotnet $(SCHEMA_FILE)
	cd ${PACKDIR}/dotnet/&& \
		echo "${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}
	cp docs/dotnet/README.md ${PACKDIR}/dotnet/

go_sdk::
	rm -rf sdk/go
	pulumi package gen-sdk --language go $(SCHEMA_FILE)

nodejs_sdk:: VERSION := $(shell pulumictl get version --language javascript)
nodejs_sdk::
	rm -rf sdk/nodejs
	chmod +x $(WORKING_DIR)/bin/$(PROVIDER)
	pulumi package gen-sdk --language nodejs $(SCHEMA_FILE)
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc
	cp docs/nodejs/README.md ${PACKDIR}/nodejs/
	cp docs/nodejs/README.md LICENSE ${PACKDIR}/nodejs/package.json ${PACKDIR}/nodejs/yarn.lock ${PACKDIR}/nodejs/bin/
	sed -i.bak 's/$${VERSION}/$(VERSION)/g' ${PACKDIR}/nodejs/bin/package.json
	grep -rlZ '@pulumi/vco' $(WORKING_DIR)/${PACKDIR}/nodejs/ | xargs -0 sed -i 's/@pulumi\/vco/@fabianv-cloud\/vco/g'
	sed -i.bak '/"install"/d' ${PACKDIR}/nodejs/bin/package.json
	sed -i.bak '/"build": "tsc",/s/,//' ${PACKDIR}/nodejs/bin/package.json
	sed -i.bak '/"install"/d' ${PACKDIR}/nodejs/package.json
	sed -i.bak '/"build": "tsc",/s/,//' ${PACKDIR}/nodejs/package.json

python_sdk:: PYPI_VERSION := $(shell pulumictl get version --language python)
python_sdk::
	rm -rf sdk/python
	chmod +x $(WORKING_DIR)/bin/$(PROVIDER)
	pulumi package gen-sdk --language python $(SCHEMA_FILE)
	cp docs/python/README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist
	sed -i '/class InstallPluginCommand/,/^[[:space:]]*$$/d' ${PACKDIR}/python/setup.py
	sed -i '/class InstallPluginCommand/,/^[[:space:]]*$$/d' ${PACKDIR}/python/bin/setup.py
	sed -i '/cmdclass={/,/},/d' ${PACKDIR}/python/setup.py
	sed -i '/cmdclass={/,/},/d' ${PACKDIR}/python/bin/setup.py

.PHONY: build
build:: provider build_sdks

.PHONY: build_sdks
build_sdks: codegen dotnet_sdk go_sdk nodejs_sdk python_sdk

only_build:: build

lint::
	for DIR in "provider" "sdk" "tests" ; do \
		pushd $$DIR && golangci-lint run -c ../.golangci.yml --timeout 10m && popd ; \
	done

install:: install_nodejs_sdk install_dotnet_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

GO_TEST := go test -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM}

test_all::
	cd provider/pkg && $(GO_TEST) ./...
	cd tests/sdk/nodejs && $(GO_TEST) ./...
	cd tests/sdk/python && $(GO_TEST) ./...
	cd tests/sdk/dotnet && $(GO_TEST) ./...
	cd tests/sdk/go && $(GO_TEST) ./...

install_dotnet_sdk::
	rm -rf $(WORKING_DIR)/nuget/$(NUGET_PKG_NAME).*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::
	#target intentionally blank

install_go_sdk::
	#target intentionally blank

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

test::
	cd examples && go test -v -tags=all -timeout 2h
