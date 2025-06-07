#!/bin/sh

set -e  # Stop on error

echo "Running migrations..."
./main db:migrate

echo "Seeding database..."
./main db:seed

echo "Starting application..."
exec ./main
