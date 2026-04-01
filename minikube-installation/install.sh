#!/bin/bash

set -e

echo "🚀 Installing Minikube and kubectl..."

# Check root
if [ "$EUID" -ne 0 ]; then
  echo "❌ Run with sudo"
  exit 1
fi

# Install dependencies
echo "📦 Installing dependencies..."
apt update
apt install -y curl apt-transport-https

# Install kubectl
echo "⬇️ Installing kubectl..."
KUBECTL_VERSION=$(curl -L -s https://dl.k8s.io/release/stable.txt)
curl -LO https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl
install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
rm kubectl

# Install Minikube
echo "⬇️ Installing Minikube..."
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
install minikube-linux-amd64 /usr/local/bin/minikube
rm minikube-linux-amd64

# Verify installs
echo "🔍 Verifying installations..."
kubectl version --client
minikube version

# Add user to docker group (if not already)
echo "👤 Ensuring Docker permissions..."
usermod -aG docker $SUDO_USER || true

echo "✅ Installation complete!"
echo ""
echo "👉 Start cluster with:"
echo "   minikube start --driver=docker"
echo ""
echo "👉 Then check:"
echo "   kubectl get nodes"