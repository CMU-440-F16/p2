#!/bin/bash

if [ -z $GOPATH ]; then
    echo "WARNING! GOPATH environment variable is not set!"
    exit 1
fi

if [ -n "$(go version | grep 'darwin/amd64')" ]; then    
    GOOS="darwin_amd64"
elif [ -n "$(go version | grep 'linux/amd64')" ]; then
    GOOS="linux_amd64"
else
    echo "FAIL: only 64-bit Mac OS X and Linux operating systems are supported"
    exit 1
fi

$GOPATH/tests_cp/tribtest.sh
$GOPATH/tests_cp/libtest.sh
$GOPATH/tests_cp/storagetest.sh
