export DOCKER_USERNAME=tnindo
export DOCKER_PASSWORD=donokasinoindro123
export VERSION_IMG_STAG=v1stag
export VERSION_MINOR_IMG_STAG=0
export GOPATH=/var/lib/jenkins/workspace/collections


cd $GOPATH/src/collections

docker login -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD";
docker push tnindo/collections:$VERSION_IMG_STAG.$VERSION_MINOR_IMG_STAG.$BUILD_ID;     