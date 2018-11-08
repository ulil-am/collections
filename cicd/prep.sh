#!/bin/bash
mkdir -p /var/lib/jenkins/workspace/collections/src/collections

rm -R /var/lib/jenkins/workspace/collections/src/collections/*

cp -r /var/lib/jenkins/workspace/collections_Prep/* /var/lib/jenkins/workspace/collections/src/collections

docker start mongo
docker start rabbitmq
docker start spostgres