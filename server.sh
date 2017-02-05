#!/usr/bin/env bash

cd server && go get -v -d && go install -v && ENVIRONMENT=development gin -a 8080 -i main.go
