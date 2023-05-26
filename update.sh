#!/bin/bash

git add .
git commit -m "update"
git push origin main
set GOOS=linux
set GOARCH=amd64

go build main.go
del main.zip
zip main.zip main
