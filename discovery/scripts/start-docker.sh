#!bin/bash
docker run --env AWS_ACCESS_KEY=$AWS_ACCESS_KEY --env AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY --env LOGSENE_TOKEN=$LOGSENE_TOKEN -d -p 3000:3000 docker.io/askfermi/cvpm-discovery