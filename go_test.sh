#!/usr/bin/env bash

SECONDS=0
go test -parallel=1 -p=1 $(go list ./... | grep -v /vendor/)
duration=$SECONDS
echo "$(($duration / 60)) minutes and $(($duration % 60)) seconds elapsed."