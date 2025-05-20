#!/bin/bash
set -e

# Print Go version
go version

# Clean any existing build artifacts
rm -f out

# Download dependencies
go mod download

# Build the application
go build -o out

# Verify the build
if [ -f "out" ]; then
    echo "Build successful!"
    ls -l out
else
    echo "Build failed!"
    exit 1
fi 