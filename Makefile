all: app swagger

GAE_DIR=backend
DESIGN=$(wildcard design/*.go)
APP=$(sort app/controllers.go $(wildcard app/* app/test/*))
SWAGGER=$(sort ${GAE_DIR}/swagger/swagger.json $(wildcard ${GAE_DIR}/swagger/*))
BIN_FLATC=flatc
FBS_DIR=${GAE_DIR}

REPO=$(shell echo $${PWD\#`go env GOPATH`/src/})

app: $(APP)
$(APP): $(DESIGN)
	goagen app -d ${REPO}/design

swagger: $(SWAGGER)
$(SWAGGER): $(DESIGN)
	goagen swagger --design ${REPO}/design --out ${GAE_DIR}

swagger-ui:
	mkdir -p $@ && curl -L `curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest|jq -r .tarball_url`| tar xzfp - -C $@ --strip=1 --no-same-owner --no-same-permissions */dist
	sed -i '' 's;"http://localhost:3200/oauth2-redirect.html";"http://localhost:8080/swagger-ui/oauth2-redirect.html";' swagger-ui/dist/swagger-ui-standalone-preset.js

FBS = $(shell find . -name "*.fbs.txt")
fbs: $(FBS)
	$(BIN_FLATC) --go $(FBS)
	mv app/*.go $(FBS_DIR) && rm -r app

build:
	go build ./app ./design && goapp build ${GAE_DIR}

test:
	go test ./app ./design && goapp test ${GAE_DIR}

lint:
	golint ./app ${GAE_DIR}

.PHONY: app swagger
