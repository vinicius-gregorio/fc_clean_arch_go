#!/bin/sh
set -e

# Debug: Echo something to confirm the script is running
echo "Running entrypoint.sh"

# Navigate to the directory containing main.go and wire_gen.go
cd cmd/ordersystem

# Debug: List files to check if wire_gen.go is present
echo "Listing files in cmd/ordersystem:"
ls -la

# Run the Go application with all files in the current directory
go run .

# Execute the CMD from the Dockerfile
exec "$@"
