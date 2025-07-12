# 🚀 AXTunnel

**AXTunnel** is a custom-built encrypted TCP tunnel with TLS obfuscation and multi-port fallback, designed to bypass censorship and DPI systems.

---

## ✨ Features

- 🔐 AES-256-GCM ready (optional payload encryption)
- 🌐 TLS 1.3 transport with fake SNI (e.g., `google.com`)
- 🔁 Multi-port failover (tries 443, 8443, 80...)
- ⚙️ Standalone executable — no external dependencies
- 📦 Auto-installer script for fast deployment
- 🎯 Supports local forwarding and remote redirection
- 🛠 Easily extensible for advanced routing

---

## 📦 Installation

Run the installer script on either the client (inside region) or server (outside region):

```
bash <(curl -Ls https://raw.githubusercontent.com/monhacer/AXTunnel/main/install_tunnel.sh)
```
You will be asked whether this machine is a client or server, and the script will:

Download latest source files

Generate TLS certificates

Compile the Go binaries

Start the client or launch multiple server ports

🧩 Configuration
Edit the config.json file:

```
{
  "RemoteIP": "x.x.x.x",
  "RemotePorts": ["443", "8443", "80"],
  "LocalPort": "1080",
  "Password": "YourSecurePassword"
}

```

RemoteIP: IP of your outside server

RemotePorts: list of ports to try for connection

LocalPort: local listener for apps

Password: optional passphrase for encryption (future use)

⚡ Usage
Once installed:

Client listens on local port (e.g., 1080)

Server listens on multiple TLS ports

Traffic gets forwarded from client to server through TLS with SNI masking

To run manually:

```
./tunnel-client
./tunnel-server 443
```

## 📁 Files

| File name            | Description                          | Role                              |
|----------------------|--------------------------------------|-----------------------------------|
| `client.go`          | TLS client for local connections     | Runs on inside (Iran) server      |
| `server.go`          | TLS server for incoming traffic      | Runs on outside server            |
| `crypto.go`          | AES-256 encryption library           | Used for payload encryption       |
| `config.json`        | Client configuration                 | Defines IPs, ports, and password  |
| `cert.pem` / `key.pem` | TLS certificate and private key   | Used for secure TLS handshake     |
| `install_tunnel.sh`  | Auto-installer Bash script           | Sets up and launches the tunnel   |
| `client.log` / `server-xxx.log` | Runtime logs             | Optional for debugging            |


📜 License
This project is licensed under MIT. Feel free to fork, improve and build on top.

👤 Author
Created by monhacer Made with ❤️ for freedom, privacy and innovation.
