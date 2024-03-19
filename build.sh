# bin/bash
echo "Building compiled versions..."

echo "Building for Windows (64-bit)"
env GOOS=windows GOARCH=amd64 go build -o bin/windows_64.exe -ldflags="-s -w" .

echo "Building for Windows (32-bit)"
env GOOS=windows GOARCH=386 go build -o bin/windows_32.exe -ldflags="-s -w" .

echo "Building for Linux (64-bit)"
env GOOS=linux GOARCH=amd64 go build -o bin/linux -ldflags="-s -w" .

echo "Building for Linux (32-bit)"
env GOOS=linux GOARCH=386 go build -o bin/linux_32 -ldflags="-s -w" .

echo "Building for Mac (Intel)"
env GOOS=darwin GOARCH=amd64 go build -o bin/mac_intel -ldflags="-s -w" .

echo "Building for Mac (Apple Silicon)"
env GOOS=darwin GOARCH=arm64 go build -o bin/mac_arm -ldflags="-s -w" .
