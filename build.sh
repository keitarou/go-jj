#!/bin/sh

set -e

if [ $# != 1 ]; then
	echo "Usage: $0 [binary name]"
	exit 0
fi

GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ./bin/linux64/$1
GO111MODULE=on GOOS=linux GOARCH=386 go build -o ./bin/linux386/$1

GO111MODULE=on GOOS=windows GOARCH=386 go build -o ./bin/windows386/$1.exe
GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o ./bin/windows64/$1.exe

GO111MODULE=on GOOS=darwin GOARCH=386 go build -o ./bin/darwin386/$1
GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin64/$1
