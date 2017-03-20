#!/usr/bin/env bash

go get ./...
go install urlshortener
mkdir -p $URLSHORTENER_LOG_DIR
urlshortener api