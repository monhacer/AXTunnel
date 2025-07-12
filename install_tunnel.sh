#!/bin/bash

# GitHub repository details
REPO="https://raw.githubusercontent.com/monhacer/AXTunnel/main"
INSTALL_DIR="/opt/axtunnel"
mkdir -p $INSTALL_DIR && cd $INSTALL_DIR

echo "[INFO] Downloading source files..."
curl -O $REPO/client.go
curl -O $REPO/server.go
curl -O $REPO/crypto.go

# Create default config.json
cat <<EOF > config.json
{
  "RemoteIP": "1.2.3.4",
  "RemotePorts": ["443", "8443", "80"],
  "LocalPort": "1080",
  "Password": "SuperSecretPassword"
}
EOF

# Install Go if not available
if ! command -v go &> /dev/null; then
  echo "[INFO] Installing Go..."
  apt update && apt install -y golang
fi

# Compile client and server
echo "[INFO] Compiling sources..."
go build -o tunnel-client client.go crypto.go
go build -o tunnel-server server.go crypto.go

# Generate TLS certificates
echo "[INFO] Generating TLS certificate..."
openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out cert.pem -subj "/CN=google.com"

# Ask server role
read -p "Is this the client machine? (y/n): " IS_CLIENT

if [[ "$IS_CLIENT" == "y" ]]; then
  echo "[INFO] Starting client in background..."
  nohup ./tunnel-client > client.log 2>&1 &
else
  echo "[INFO] Starting server on multiple ports..."
  for port in 443 8443 80; do
    nohup ./tunnel-server "$port" > "server-$port.log" 2>&1 &
    echo "[OK] Server running on port $port"
  done
fi

echo "[DONE] Installation complete! Files are located in $INSTALL_DIR"
