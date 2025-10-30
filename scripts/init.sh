#!/bin/bash

cd shared || exit 1
go mod tidy
cd ../key-gen-service || exit 1
go mod tidy
cd ../url-shortener-service || exit 1
go mod tidy
cd ../url-redirect-service || exit 1
go mod tidy
cd ../api-gateway-service || exit 1
go mod tidy
cd .. || exit 1
