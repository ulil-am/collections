#!/bin/bash
export GOPATH=/var/lib/jenkins/workspace/collections
# export GOPATH=/home/araya/TN/collections_service
# export GOPATH=/home/jannes/work/TN/dlor
export WORKDIR=$GOPATH/src/collections

export GOENV=devci
export GOAPP=collections

# cat $GOPATH/src/collections/conf/ci/mq.json
# cat $GOPATH/src/collections/conf/ci/mongodb.json

# cd $GOPATH/src/collections && $GOPATH/bin/bee migrate -driver=postgres -conn="postgres://postgres:root@172.17.0.1:5432/postgres?sslmode=disable"

# Unit test 
# For Coverage
cd $WORKDIR &&
go test -v \
./models/logic/products/... \
./models/logic/genappid/... \
-coverprofile=$WORKDIR/cicd/sonarqube-report/coverage-report.out

# For Unit test
go test -v \
./models/logic/products/... \
./models/logic/genappid/... \
-json >> $WORKDIR/cicd/sonarqube-report/unit-report.json

# Component test
cd $GOPATH/src/collections/routers/componenttest && 
go test -v \
./http/... \
./grpc/...

# Run SonarQube
# cd $WORKDIR &&
# docker run --rm \
#     -v $(pwd):$WORKDIR \
#     -w=$WORKDIR --network=sonar \
#     nikhuber/sonar-scanner:latest sonar-scanner