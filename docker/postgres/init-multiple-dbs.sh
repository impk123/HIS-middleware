#!/bin/bash
set -e

# Split environment variables into arrays
IFS=',' read -ra dbs <<< "$POSTGRES_MULTIPLE_DATABASES"
IFS=',' read -ra users <<< "$POSTGRES_MULTIPLE_USERS"
IFS=',' read -ra passwords <<< "$POSTGRES_MULTIPLE_PASSWORDS"

# Create each database and user
for i in "${!dbs[@]}"; do
  db=${dbs[$i]}
  user=${users[$i]}
  password=${passwords[$i]}

  echo "Creating user '$user' and database '$db'"
    DB_EXISTS=$(psql -U "$POSTGRES_USER" -tAc "SELECT 1 FROM pg_database WHERE datname='$db'")
    if [ "$DB_EXISTS" != "1" ]; then
      psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -c "CREATE DATABASE \"$db\";"
    else
      echo "Database '$db' already exists. Skipping creation."
    fi
done