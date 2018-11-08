#!/bin/bash
export GOENV=devci
export GOAPP=medium
export GOPATH=/var/lib/jenkins/workspace/Medium
# export GOPATH=/home/araya/TN/medium_service

# process download gometalinter
pwd

go version
go env
echo $GOENV
cd $GOPATH
go get github.com/beego/bee
# chmod +x $GOPATH/src/medium/goget.sh
# $GOPATH/src/medium/goget.sh
go get -u gopkg.in/alecthomas/gometalinter.v1 && $GOPATH/bin/gometalinter.v1 --install

# run linter
chmod +x $GOPATH/src/medium/cicd/linter/runLinter
cd $GOPATH/src/medium && ./cicd/linter/runLinter