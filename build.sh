# bin/bash
echo "Building compiled versions..."

echo "Building for Windows (64-bit)"
env GOOS=windows GOARCH=amd64 go build -o bin/windows-64.exe .

echo "Building for Windows (32-bit)"
env GOOS=windows GOARCH=386 go build -o bin/windows-32

echo "Building for Linux (64-bit)"
env GOOS=linux GOARCH=amd64 go build -o bin/linux .

echo "Building for Linux (32-bit)"
env GOOS=linux GOARCH=386 go build -o bin/linux-32 .

echo "Building for Mac (Intel)"
env GOOS=darwin GOARCH=amd64 go build -o bin/mac_intel .

echo "Building for Mac (Apple Silicon)"
env GOOS=darwin GOARCH=arm64 go build -o bin/mac_arm .
