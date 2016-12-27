#!/usr/bin/env bash

go vet $(go list ./... | grep -v /vendor/)