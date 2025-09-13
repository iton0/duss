#!/bin/bash

cd shared
go mod tidy
cd ../key-gen-service
go mod tidy
cd ../url-shortener-service
go mod tidy
cd ../url-redirect-service
go mod tidy
cd ../api-gateway-service
go mod tidy
cd ..
