
#!/bin/bash

# Build the Wails application
echo "Building StockScribe FIFO desktop application..."

# Install Wails CLI if not already installed
if ! command -v wails &> /dev/null; then
    echo "Installing Wails CLI..."
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
fi

# Clean previous builds
echo "Cleaning previous builds..."
rm -rf build

# Build the application for production
echo "Building for production..."
wails build -platform windows/amd64 -clean -webview2 embed

echo "Build complete! You can find your executable in the build/bin directory."
