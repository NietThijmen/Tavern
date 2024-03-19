# bin/bash
echo "Building compiled versions..."

echo "Building for Windows"
env GOOS=windows GOARCH=amd64 go build -o bin/windows.exe .

echo "Building for Linux"
env GOOS=linux GOARCH=amd64 go build -o bin/linux .

echo "Building for Mac (Intel)"
env GOOS=darwin GOARCH=amd64 go build -o bin/mac_intel .

echo "Building for Mac (Apple Silicon)"
env GOOS=darwin GOARCH=arm64 go build -o bin/mac_arm .
