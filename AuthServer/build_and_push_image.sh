#!/usr/bin/env bash

docker build -t meik99/coffee-auth-server:$1 -t meik99/coffee-auth-server:latest $(pwd)
docker push meik99/coffee-auth-server:$1
docker push meik99/coffee-auth-server:latest
