#!/usr/bin/env bash

golint $(go list ./... | grep -v /vendor/)