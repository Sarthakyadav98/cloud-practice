```bash
#!/bin/bash

set -e

GO_VERSION="1.22.5"

echo "📦 Installing Go $GO_VERSION..."

# Download Go
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz

# Remove old Go (if exists)
sudo rm -rf /usr/local/go

# Extract
sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

# Cleanup
rm go${GO_VERSION}.linux-amd64.tar.gz

# Add to PATH if not already added
if ! grep -q "/usr/local/go/bin" ~/.bashrc; then
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
fi

# Apply changes
source ~/.bashrc

# Verify
echo "✅ Go installed:"
go version
```

```bash
#!/bin/bash

set -e

echo "📦 Installing Node.js via NVM..."

# Install dependencies
sudo apt update
sudo apt install -y curl build-essential

# Install NVM
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash

# Load NVM immediately
export NVM_DIR="$HOME/.nvm"
source "$NVM_DIR/nvm.sh"

# Install latest LTS Node
nvm install --lts
nvm use --lts

# Verify
echo "✅ Node installed:"
node -v
npm -v
```