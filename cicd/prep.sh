#!/bin/bash
mkdir -p /var/lib/jenkins/workspace/Medium/src/medium

rm -R /var/lib/jenkins/workspace/Medium/src/medium/*

cp -r /var/lib/jenkins/workspace/Medium_Prep/* /var/lib/jenkins/workspace/Medium/src/medium

docker start mongo
docker start rabbitmq
docker start spostgres