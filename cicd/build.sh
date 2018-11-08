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
export GOPATH=/var/lib/jenkins/workspace/Medium

#build image
pwd
sed -i -e "s/runmode = dev/runmode = prod/g" $GOPATH/src/medium/conf/app.conf
cat $GOPATH/src/medium/conf/app.conf
cd $GOPATH/src/medium
pwd
go build
ls -lah
docker build -t tnindo/medium:$VERSION_IMG_STAG.$VERSION_MINOR_IMG_STAG.$BUILD_ID -f Dockerfile .
