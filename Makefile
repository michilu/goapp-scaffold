all: app swagger

GAE_DIR=backend
DESIGN_DIR=design
APP_DIR=app
DESIGN=$(wildcard ${DESIGN_DIR}/*.go)
APP=$(sort ${APP_DIR}/controllers.go $(wildcard ${APP_DIR}/* ${APP_DIR}/test/*))
SWAGGER=$(sort ${GAE_DIR}/swagger/swagger.json $(wildcard ${GAE_DIR}/swagger/*))
BIN_FLATC=flatc
FBS_DIR=${GAE_DIR}

REPO=$(shell echo $${PWD\#`go env GOPATH`/src/})

app: $(APP)
$(APP): $(DESIGN)
	goagen app -d ${REPO}/${DESIGN_DIR}

swagger: $(SWAGGER)
$(SWAGGER): $(DESIGN)
	goagen swagger --design ${REPO}/${DESIGN_DIR} --out ${GAE_DIR}

swagger-ui:
	mkdir -p $@ && curl -L `curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest|jq -r .tarball_url`| tar xzfp - -C $@ --strip=1 --no-same-owner --no-same-permissions */dist
	sed -i '' 's;"http://localhost:3200/oauth2-redirect.html";"http://localhost:8080/swagger-ui/oauth2-redirect.html";' swagger-ui/dist/swagger-ui-standalone-preset.js

FBS = $(shell find . -name "*.fbs.txt")
fbs: $(FBS)
	$(BIN_FLATC) --go $(FBS)
	mv ${APP_DIR}/*.go $(FBS_DIR) && rm -r ${APP_DIR}

build:
	go build ${APP_DIR} ${DESIGN_DIR} && goapp build ${GAE_DIR}

test:
	go test ${APP_DIR} ${DESIGN_DIR} && goapp test ${GAE_DIR}

lint:
	golint ${APP_DIR} ${GAE_DIR}

.PHONY: app swagger
