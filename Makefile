all: css js

NODE_MODULES_BASE=node_modules
BIN_COFFEE=$(NODE_MODULES_BASE)/.bin/coffee
BIN_UGLIFYJS=$(NODE_MODULES_BASE)/.bin/uglifyjs
BIN_FLATC=flatc

JS_DIR=src/_js
CSS_DIR=src/_css
FBS_DIR=src

.SUFFIXES: .coffee .js
.coffee.js:
	$(BIN_COFFEE) -c $<
.SUFFIXES: .js .min.js
.js.min.js:
	$(BIN_UGLIFYJS) --define PRODUCTION=0 -nc -m -r "$$" -o $@ $<
COFFEE = $(foreach dir,$(JS_DIR),$(wildcard $(dir)/*.coffee))
JS = $(COFFEE:.coffee=.js)
MINJS = $(JS:.js=.min.js)

.SUFFIXES: .sass .css
.sass.css:
	compass compile $< -c $(CSS_DIR)/config.rb
.SUFFIXES: .sass .min.css
.sass.min.css:
	compass compile --environment production $< -c $(CSS_DIR)/config.rb
	mv $*.css $@
SASS = $(wildcard $(CSS_DIR)/*.sass)
CSS = $(SASS:.sass=.css)
MINCSS = $(SASS:.sass=.min.css)

js: $(MINJS) $(JS)
css: $(MINCSS) $(CSS)

FBS = $(wildcard $(FBS_DIR)/**/*.fbs)
fbs: $(FBS)
	$(BIN_FLATC) -g -o src/fbs $<
