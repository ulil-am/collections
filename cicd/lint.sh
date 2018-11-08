#!/bin/bash
export GOENV=devci
export GOAPP=collections
export GOPATH=/var/lib/jenkins/workspace/collections
# export GOPATH=/home/araya/TN/collections_service

# process download gometalinter
pwd

go version
go env
echo $GOENV
cd $GOPATH
go get github.com/beego/bee
# chmod +x $GOPATH/src/collections/goget.sh
# $GOPATH/src/collections/goget.sh
go get -u gopkg.in/alecthomas/gometalinter.v1 && $GOPATH/bin/gometalinter.v1 --install

# run linter
chmod +x $GOPATH/src/collections/cicd/linter/runLinter
cd $GOPATH/src/collections && ./cicd/linter/runLinter