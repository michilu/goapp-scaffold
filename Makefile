all: fbs

BIN_FLATC=flatc

FBS_DIR=backend

FBS = $(shell find . -name "*.fbs.txt")
fbs: $(FBS)
	$(BIN_FLATC) --go $(FBS)
	mv app/*.go $(FBS_DIR) && rm -r app

swagger-ui:
	mkdir -p $@ && curl -L `curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest|jq -r .tarball_url`| tar xzfp - -C $@ --strip=1
