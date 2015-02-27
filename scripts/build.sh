#!/bin/bash

OUTPUT_PATH=./out
SCRIPT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
SOURCE_PATH=$SCRIPT_PATH/..

cd $SCRIPT_PATH

rm -rf $OUTPUT_PATH
mkdir $OUTPUT_PATH

rm -f ../client/client.go

cd $SOURCE_PATH
go-bindata -pkg="client" -o=client/client.go -ignore=\\.gitignore client
cd $SCRIPT_PATH

if [[ $FLOOD_DEBUG = 1 ]]; then
    go build -gcflags "-N -l"  ../hookinator.go
else
    go build  -ldflags "-w" ../hookinator.go
fi

chmod +x hookinator
mv hookinator $OUTPUT_PATH/
