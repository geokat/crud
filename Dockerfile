FROM golang:1.8

ADD . /go/src/github.com/geokat/crud
CMD cd /go/src/github.com/geokat/crud && go build -o /go/bin/crud && /go/bin/crud
