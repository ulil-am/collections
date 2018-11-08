# FROM golang:1.9.2-alpine
FROM golang:1.11.1-alpine

### START: Setting Environment ###

ENV GOPATH /go
ENV GOAPP collections
ENV GOENV kube
ENV PATH $GOPATH/bin:$PATH
ADD conf/bin $GOPATH/bin

RUN apk add --no-cache \
	libc6-compat
RUN apk add --no-cache tzdata
### END: Setting Environment ###

### START: add source ###
RUN mkdir -p /go/src/collections
RUN mkdir -p /go/src/collections/logs

# RUN mkdir -p /go/src/collections/helper
# ADD helper /go/src/collections/helper

# RUN mkdir -p /go/src/collections/vendor
# ADD vendor /go/src/collections/vendor

RUN mkdir -p /go/src/collections/conf	
ADD conf /go/src/collections/conf

RUN mkdir -p /go/src/collections/storages
ADD storages /go/src/collections/storages

# RUN mkdir -p /go/src/collections/database
# ADD database /go/src/collections/database

ADD collections /go/src/collections


### END: add source ###


### START: Initialize dependency ###

# RUN apk add --no-cache git mercurial \
# && go get github.com/beego/bee \
# && apk del git mercurial
WORKDIR /go/src/collections
RUN ls
### END: Initialize dependency ###

CMD ["/go/src/collections/collections"]