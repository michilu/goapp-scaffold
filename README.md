# goapp-scaffold
A scaffold for Google App Engine Golang.

# initial setup

## setup build tools

    $ npm install

## setup GOPATH and GOROOT

    $ goof make goapp-scaffold
    (go:goapp-scaffold)$ vi activate

    export GOPATH=/<path-to>/goapp-scaffold/gopath
    export GOROOT=$HOME/google-cloud-sdk/platform/google_appengine/goroot

# go get

    (go:goapp-scaffold)$ goapp get github.com/GoogleCloudPlatform/go-endpoints/endpoints
    ...

# go serve

    (go:goapp-scaffold)$ goapp serve src
