echo '#!/bin/bash

# Remove the executable from the system PATH
echo "Uninstalling Docker Container Manager..."
sudo rm /usr/local/bin/docker-container-manager

# Check if the removal was successful
if [ $? -ne 0 ]; then
    echo "Uninstallation failed. Exiting."
    exit 1
fi

echo "Docker Container Manager uninstalled successfully!"' > uninstall.sh