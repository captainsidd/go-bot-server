#!/bin/bash

# List of ports to choose from
PORTS=(
  66
	80
	81
	443
	445
	457
	1080
	1241
	1352
	1433)

# Function to choose a random port from the list
choose_random_port() {
  echo "${PORTS[RANDOM % ${#PORTS[@]}]}"
}
# Target server IP
TARGET_IP="127.0.0.1"
# Main loop that runs every 10 seconds
while true; do
  RANDOM_PORT=$(choose_random_port)
  echo "Sending request to $TARGET_IP:$RANDOM_PORT..."
  # curl the endpoint
  curl "$TARGET_IP:$RANDOM_PORT"
  # Sleep for 10 seconds before the next iteration
  sleep 10
done
