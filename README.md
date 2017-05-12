# goapp-scaffold
A scaffold for Google App Engine Golang.

[![wercker status](https://app.wercker.com/status/01a88b7c9b69392341715dc7b9a234ba/m/master "wercker status")](https://app.wercker.com/project/byKey/01a88b7c9b69392341715dc7b9a234ba)

# clone repository

    $ cd `goapp env GOPATH`
    $ mkdir -p src/github.com/MiCHiLU
    $ cd src/github.com/MiCHiLU
    $ git clone --depth 1 https://github.com/MiCHiLU/goapp-scaffold.git
    $ cd goapp-scaffold

# dep

    $ go get -u github.com/golang/dep/...
    $ dep init
    $ dep ensure

# start server

    $ dev_appserver.py backend

# with [goa](https://github.com/goadesign/goa)

Install goa:

    $ go get -u github.com/goadesign/goa/...

Generate code:

    $ goagen app -d github.com/MiCHiLU/goapp-scaffold/design
    $ goagen swagger -d github.com/MiCHiLU/goapp-scaffold/design

# with swagger

Get swagger-ui:

    $ make swagger-ui

Open [swagger-ui](http://localhost:8080/swagger-ui/?url=http://localhost:8080/swagger.json).

    $ open "http://localhost:8080/swagger-ui/?url=http://localhost:8080/swagger.json"

# Deploy to the Google App Engine

Create a Google App Engine application:

    $ CLOUDSDK_CORE_PROJECT=<APP-ID> gcloud app create

Get an OAuth 2.0 refresh token:

    $ jq .refresh_token ~/.appcfg_oauth2_tokens

Then set to an `APP_ENGINE_TOKEN` of the application environment variables on the Wercker CI.
