#!/usr/bin/env bash
pwd
set -e
echo "" > coverage.txt
for d in $(go list ./... | grep -v vendor); do
    echo $d
    go test -coverprofile=profile.out $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        # go tool cover -html=profile.out
        rm profile.out
    fi
done
