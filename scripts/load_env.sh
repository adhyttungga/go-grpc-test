#!/bin/bash

ENV_FILE=".env"

if [ -f "$ENV_FILE" ]; then
    while IFS= read -r line; do
        if [[ "$line" =~ ^\s*#.*$ || -z "$line" ]]; then
            continue
        fi

        key=$(echo "$line" | cut -d '=' -f 1)
        value=$(echo "$line" | cut -d '=' -f 2-)
        value=$(echo "$value" | sed -e "s/^\'//g" -e "s/\'$//g" -e 's/^"//g' -e 's/"$//g' -e 's/^[ \t]*//;s/[ \t]*$//')

        if [[ "$key" == "DB_HOST" || "$key" == "DB_PORT" || "$key" == "REDIS_ADDR" ]]; then
            export "$key=$value"
        fi

    done < "$ENV_FILE"
fi
