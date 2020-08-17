#!/bin/bash

echo "Compiling..."
GOOS=linux GOARCH=amd64 go build -o kashtoshi main.go
echo "Done"
echo "Packaging..."
mkdir dist

mv kashtoshi dist/kashtoshi
echo "Done. Executable can be found in the dist/ directory."