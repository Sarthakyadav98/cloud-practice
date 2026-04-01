#!/bin/bash

set -e  # exit on error

echo "🚀 Starting Docker installation..."

# Ensure script is run as root
if [ "$EUID" -ne 0 ]; then
  echo "❌ Please run as root (use sudo)"
  exit 1
fi

echo "🧹 Removing old Docker versions..."
apt remove -y docker.io docker-compose docker-compose-v2 docker-doc podman-docker containerd runc || true

echo "🔄 Updating packages..."
apt update
apt install -y ca-certificates curl

echo "🔐 Setting up Docker GPG key..."
install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
chmod a+r /etc/apt/keyrings/docker.asc

echo "📦 Adding Docker repository..."
tee /etc/apt/sources.list.d/docker.sources > /dev/null <<EOF
Types: deb
URIs: https://download.docker.com/linux/ubuntu
Suites: $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}")
Components: stable
Architectures: $(dpkg --print-architecture)
Signed-By: /etc/apt/keyrings/docker.asc
EOF

echo "🔄 Updating package list again..."
apt update

echo "📥 Installing Docker..."
apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

echo "⚙️ Enabling and starting Docker..."
systemctl enable docker
systemctl start docker

echo "👤 Adding current user to docker group..."
usermod -aG docker $SUDO_USER

echo "✅ Installation complete!"
echo "👉 Please log out and log back in to use Docker without sudo."
echo "👉 Test with: docker run hello-world"