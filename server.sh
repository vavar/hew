#!/usr/bin/env bash

cd server && gin -a 8080 -i -g /main.go
