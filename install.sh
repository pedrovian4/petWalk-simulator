#!/bin/bash

if [ "$(id -u)" != "0" ]; then
   echo "This script must be run as root" 1>&2
   exit 1
fi

echo "Updating the system and installing Git..."
apt-get update
apt-get install -y git

if ! command -v go &> /dev/null
then
    echo "Go not found. Please install Go before continuing."
    exit 1
fi

echo "Cloning the repository..."
git clone [URL of the Repository] petWalkSimulator
cd petWalkSimulator

echo "Building the application..."
go build

echo "Moving the binary to /usr/local/bin"
mv petWalkSimulator /usr/local/bin/petWalkSimulator

echo "Installation completed. You can now run the application with 'petWalkSimulator'"
