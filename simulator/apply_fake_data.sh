#!/bin/bash

# Determine directory of this script
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SQL_FILE="$DIR/insert_fake_data.sql"

# Check if SQL file exists
if [ ! -f "$SQL_FILE" ]; then
    echo "Error: $SQL_FILE not found."
    exit 1
fi

echo "Looking for MariaDB/MySQL container..."

# Try to find a running container matching mariadb_1 (compose)
CONTAINER_ID=$(docker ps -q -f name=mariadb_1 -f status=running | head -n 1)

# Try swarm name format (smartcart_mariadb_1) if not found
if [ -z "$CONTAINER_ID" ]; then
    CONTAINER_ID=$(docker ps -q -f name=smartcart_mariadb_1 -f status=running | head -n 1)
fi

# Try generic name if not found
if [ -z "$CONTAINER_ID" ]; then
    CONTAINER_ID=$(docker ps -q -f name=mariadb -f status=running | head -n 1)
fi

# Try mysql name if not found
if [ -z "$CONTAINER_ID" ]; then
    CONTAINER_ID=$(docker ps -q -f name=mysql -f status=running | head -n 1)
fi

if [ -n "$CONTAINER_ID" ]; then
    echo "Found running container: $CONTAINER_ID"
    echo "Applying insert_fake_data.sql to container..."
    
    # Try running mariadb client first, fallback to mysql client
    docker exec -i "$CONTAINER_ID" mariadb -uuniwa_admin -padminUNIWA supermarket_db < "$SQL_FILE" 2>/dev/null
    STATUS=$?
    
    if [ $STATUS -ne 0 ]; then
        docker exec -i "$CONTAINER_ID" mysql -uuniwa_admin -padminUNIWA supermarket_db < "$SQL_FILE"
        STATUS=$?
    fi
    
    if [ $STATUS -eq 0 ]; then
        echo "Successfully imported fake data into the container database!"
    else
        echo "Error: Failed to import SQL file into the container database."
    fi
else
    echo "Could not find a running mariadb/mysql container automatically via docker ps."
    echo "Attempting to run local mysql command against 127.0.0.1:3306..."
    mysql -h 127.0.0.1 -P 3306 -u uniwa_admin -padminUNIWA supermarket_db < "$SQL_FILE"
    if [ $? -eq 0 ]; then
        echo "Successfully imported fake data via local mysql command!"
    else
        echo "Could not import data. Please ensure the Docker stack is running and the database is accessible."
    fi
fi
