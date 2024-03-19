# bin/bash
env GOOS=windows GOARCH=amd64 go build -o bin/build.exe .
env GOOS=linux GOARCH=amd64 go build -o bin/build .
env GOOS=darwin GOARCH=amd64 go build -o bin/build_mac .

env GOOS=darwin GOARCH=arm64 go build -o bin/build_mac_arm .
