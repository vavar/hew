#!/usr/bin/env bash

go install $(go list ./... | grep -v /vendor/)