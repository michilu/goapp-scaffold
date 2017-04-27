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
