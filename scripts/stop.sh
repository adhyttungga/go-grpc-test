#!/bin/bash

# Navigate to the project directory
cd "$(dirname "$0")/.."

# Stop and delete Docker container
docker-compose down --rmi local
