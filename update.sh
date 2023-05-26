#!/bin/bash

git add .
git commit -m "update"
git push origin main

go build
rm main.zip
zip main.zip main
