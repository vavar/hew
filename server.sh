#!/usr/bin/env bash

cd server && ENVIRONMENT=development gin -a 8080 -i -g server/main.go
