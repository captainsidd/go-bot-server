from scapy.all import IP, TCP, send
import random
import time

# Define the target and spoofed IP
target_ip = "127.0.0.1"
spoofed_ip = ".".join(str(random.randint(0, 255)) for _ in range(4))

ports = [80, 443, 21, 22, 25, 110, 143, 993, 995, 587]  # Add or modify ports as needed


while True:
  # Pick a random port from the list
  random_port = random.choice(ports)
  # Create a packet
  packet = IP(src=spoofed_ip, dst=target_ip)/TCP(dport=random_port)
  print(packet)
  # Send the packet
  send(packet)
  time.sleep(1)