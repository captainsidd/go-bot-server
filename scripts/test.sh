#!/bin/bash

# List of ports to choose from
PORTS=(80 443 21 22 25 110 143 993 995 587)

# Function to generate a random IP address
generate_random_ip() {
  echo "$((RANDOM % 256)).$((RANDOM % 256)).$((RANDOM % 256)).$((RANDOM % 256))"
}

# Function to choose a random port from the list
choose_random_port() {
  echo "${PORTS[RANDOM % ${#PORTS[@]}]}"
}

# Target server IP
TARGET_IP="127.0.0.1"

# Main loop that runs every 30 seconds
while true; do
  RANDOM_IP=$(generate_random_ip)
  RANDOM_PORT=$(choose_random_port)

  echo "Sending packet from $RANDOM_IP to $TARGET_IP:$DEST_PORT on local port $RANDOM_PORT..."

  # Use nping to send a spoofed packet
  hping3 -S -a $RANDOM_IP -s $RANDOM_PORT -p $RANDOM_PORT $TARGET_IP

  # Sleep for 30 seconds before the next iteration
  sleep 30
done
