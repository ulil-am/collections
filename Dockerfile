# FROM golang:1.9.2-alpine
FROM golang:1.11.1-alpine

### START: Setting Environment ###

ENV GOPATH /go
ENV GOAPP medium
ENV GOENV kube
ENV PATH $GOPATH/bin:$PATH
ADD conf/bin $GOPATH/bin

RUN apk add --no-cache \
	libc6-compat
RUN apk add --no-cache tzdata
### END: Setting Environment ###

### START: add source ###
RUN mkdir -p /go/src/medium
RUN mkdir -p /go/src/medium/logs

# RUN mkdir -p /go/src/medium/helper
# ADD helper /go/src/medium/helper

# RUN mkdir -p /go/src/medium/vendor
# ADD vendor /go/src/medium/vendor

RUN mkdir -p /go/src/medium/conf	
ADD conf /go/src/medium/conf

RUN mkdir -p /go/src/medium/storages
ADD storages /go/src/medium/storages

# RUN mkdir -p /go/src/medium/database
# ADD database /go/src/medium/database

ADD medium /go/src/medium


### END: add source ###


### START: Initialize dependency ###

# RUN apk add --no-cache git mercurial \
# && go get github.com/beego/bee \
# && apk del git mercurial
WORKDIR /go/src/medium
RUN ls
### END: Initialize dependency ###

CMD ["/go/src/medium/medium"]