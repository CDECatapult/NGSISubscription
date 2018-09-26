#!/bin/bash

echo "Building executables..."

GOOS=darwin GOARCH=386 go build main.go utility.go
mv main ./bin/subscribe_osx
echo "MAC OSX version created"
GOOS=windows GOARCH=386 go build main.go utility.go
mv main.exe ./bin/subscribe_win.exe
echo "Windows version created"
GOOS=linux GOARCH=386 go build main.go utility.go
mv main ./bin/subscribe_linux
echo "Linux version created"

echo "All done!"