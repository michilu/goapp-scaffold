all: fbs

BIN_FLATC=flatc

FBS_DIR=backend

FBS = $(shell find . -name "*.fbs.txt")
fbs: $(FBS)
	$(BIN_FLATC) --go $(FBS)
	mv app/*.go $(FBS_DIR) && rm -r app

restapi: swagger.yaml
	swagger generate server --flag-strategy=pflag
