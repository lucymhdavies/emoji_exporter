#!/bin/bash
#
# Launch script for EC2 instance

# Update packages
yum update -y

# Stuff we actually need
yum installl -y docker git

# Configure Docker
service docker start
usermod -a -G docker ec2-user

# Docker Compose
curl -L "https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m)" -o /bin/docker-compose
chmod +x /bin/docker-compose

# Clone Repo
git clone https://github.com/lucymhdavies/emoji_exporter.git

# Run!
cd emoji_exporter
docker-compose up -d

