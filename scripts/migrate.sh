#!/bin/bash

# This script is used to run database migrations for the news shared service.

set -e

# Define the database connection parameters
DB_USER="your_db_user"
DB_PASSWORD="your_db_password"
DB_NAME="your_db_name"
DB_HOST="localhost"
DB_PORT="5432"

# Run the migration using the SQL file
echo "Running migrations..."
psql -U $DB_USER -d $DB_NAME -h $DB_HOST -p $DB_PORT -f ../migrations/0001_init.up.sql

echo "Migrations completed successfully."