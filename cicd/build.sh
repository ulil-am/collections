#!/bin/bash

export DOCKER_USERNAME=tnindo
export DOCKER_PASSWORD=donokasinoindro123
export KEY_SSH=/var/lib/jenkins/workspace/tn-indo-key.pem
export USER_SSH=ubuntu
export HOST_DEV=203.154.91.182
export VERSION_IMG=v1
export VERSION_MINOR_IMG=0
export VERSION_IMG_STAG=v1stag
export VERSION_MINOR_IMG_STAG=0
export GOPATH=/var/lib/jenkins/workspace/collections

#build image
pwd
sed -i -e "s/runmode = dev/runmode = prod/g" $GOPATH/src/collections/conf/app.conf
cat $GOPATH/src/collections/conf/app.conf
cd $GOPATH/src/collections
pwd
go build
ls -lah
docker build -t tnindo/collections:$VERSION_IMG_STAG.$VERSION_MINOR_IMG_STAG.$BUILD_ID -f Dockerfile .
