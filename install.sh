#!/bin/bash

# Check if the script is run as root
if [ "$(id -u)" != "0" ]; then
   echo "This script must be run as root" 1>&2
   exit 1
fi

# Update the system and install Git if necessary
echo "Updating the system and installing Git..."
apt-get update
apt-get install -y git

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go not found. Please install Go before continuing."
    exit 1
fi

# Clone the project repository
echo "Cloning the repository..."
git clone [URL of the Repository] petWalkSimulator
cd petWalkSimulator

# Build the application
echo "Building the application..."
go build

# Move the binary to a global path
echo "Moving the binary to /usr/local/bin"
mv petWalkSimulator /usr/local/bin/petWalkSimulator

# Completion message
echo "Installation completed. You can now run the application with 'petWalkSimulator'"
