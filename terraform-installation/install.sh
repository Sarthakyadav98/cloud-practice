#!/usr/bin/env bash

set -e  # exit on error

echo "🔄 Updating system..."
sudo apt update -y

echo "📦 Installing dependencies..."
sudo apt install -y gnupg software-properties-common curl

echo "🔑 Adding HashiCorp GPG key..."
curl -fsSL https://apt.releases.hashicorp.com/gpg | \
sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg

echo "📁 Adding HashiCorp repository..."
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \
sudo tee /etc/apt/sources.list.d/hashicorp.list

echo "🔄 Updating package list..."
sudo apt update -y

echo "⬇️ Installing Terraform..."
sudo apt install -y terraform

echo "✅ Verifying installation..."
terraform -version

echo "🎉 Terraform installation complete!"