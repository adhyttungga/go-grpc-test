#!/bin/bash

# Load the .env file to environment variable
# export $(grep -v '^#' .env | xargs)
# source ./load_env.sh
export \
    $(grep 'DB_HOST' .env | xargs -0) \
    $(grep 'DB_PORT' .env | xargs -0) \
    $(grep 'REDIS_ADDR' .env | xargs -0)

# Wait for postgres connection 
echo "wait for ${DB_HOST}:${DB_PORT}"
wait-for "${DB_HOST}:${DB_PORT}" -- "$@"

# Wait for redis connection
echo "wait for ${REDIS_ADDR}"
wait-for "${REDIS_ADDR}" -- "$@"

# Run the app
./main
