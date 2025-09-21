#!/bin/bash

# Navigate to the project directory
cd "$(dirname "$0")/.."

# Build and start Docker container
docker-compose up --build -d
