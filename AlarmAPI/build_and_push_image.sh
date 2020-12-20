#!/usr/bin/env bash

docker build -t meik99/coffee-alarm-api:$1 -t meik99/coffee-alarm-api:latest $(pwd)
docker push meik99/coffee-alarm-api:$1
docker push meik99/coffee-alarm-api:latest
