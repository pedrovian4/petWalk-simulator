# Pet Walk Simulator
### Overview

The Pet Walk Simulator is a tool designed to simulate the movement of a pet for testing GPS tracking systems. Developed in Go, it allows the simulation of linear movements, sending geographical location data (latitude and longitude) to a specified server via a TCP connection. This tool is particularly useful for developers and testers working on pet tracking solutions.
Features

> Simulate linear movement over a configurable distance.
Send location data via TCP.
Support for JWT authentication for secure connections.

### Prerequisites
- Go programming environment
- Access to a TCP server for receiving data

### Installation
``
    git clone https://github.com/pedrovian4/petWalk-simulato
    cd petWalkSimulator
    sudo ./install.sh
``
 
### Usage

After installation, you can run the Pet Walk Simulator using the following command format:

Command Line Arguments

    --server: The TCP server address (e.g., "127.0.0.1:12345") where the simulator will send data.
    --distance: The total distance (in kilometers) for the movement simulation.
    --jwt (optional): JWT token for authentication if the server requires it.

Example
`petWalkSimulator --server="127.0.0.1:12345" --distance=3`

Contributing

Contributions to the Pet Walk Simulator are welcome. Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.
License

This project is licensed under the MIT License - see the LICENSE file for details.
Contact

Author: Pedro Viana
