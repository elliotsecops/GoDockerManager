echo '#!/bin/bash

# Build the project
echo "Building Docker Container Manager..."
go build -o docker-container-manager ./cmd/docker-container-manager

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Exiting."
    exit 1
fi

# Move the executable to a directory in the system PATH
echo "Installing Docker Container Manager..."
sudo mv docker-container-manager /usr/local/bin/

# Check if the move was successful
if [ $? -ne 0 ]; then
    echo "Installation failed. Exiting."
    exit 1
fi

echo "Docker Container Manager installed successfully!"' > install.sh