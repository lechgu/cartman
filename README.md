# Cartman

Cartman is a simple local Certificate Authority written in Go. It helps generate certificates for securing **network traffic**, for example:

- gRPC services
- mutual TLS (mTLS)
- internal microservice communication
- test environments
- development and staging deployments
- public-facing services, if the Cartman root certificate is trusted

Cartman is a port and evolution of [TinyCA](https://github.com/lechgu/tinyca), with a focus on cross-platform usability, minimal dependencies, and flexible support for different cryptographic algorithms.

Cartman can be used in **any context where the root certificate is trusted** — it is not limited to internal use only.

## Features

- Initialize a self-signed root CA
- Issue leaf certificates (with DNS names and/or IP addresses)
- Support for multiple signature algorithms
- Outputs standard PEM-encoded certificates and private keys
- Cross-platform: runs on Linux, macOS, Windows

## Installing Cartman

### Pre-built Binaries

Pre-built binaries are available from the [Releases](https://github.com/lechgu/cartman/releases) section.

Download the binary for your platform and place it in your `PATH`.

### Install via `go install`

If you have Go **1.24.0** or newer, you can install Cartman directly:

```bash
go install github.com/lechgu/cartman@latest
```

The `cartman` binary will be placed in your `$GOPATH/bin` (or `$HOME/go/bin`), which should be in your `PATH`.at

### Build from Source

You can also build Cartman manually from source:

```bash
git clone https://github.com/lechgu/cartman.git
cd cartman
go build -o cartman
```

## Running Cartman

### Initialize the Certificate Authority

```bash
cartman init
```

This creates a `.cartman` directory in the current folder, containing:

- `cert.pem` — self-signed root CA certificate
- `key.pem` — root CA private key

By default, the CA certificate is valid for **10 years**.

### Issue Leaf Certificates

```bash
cartman issue --name localhost --dns localhost --ip 127.0.0.1
```

This creates a new folder named `localhost` (matching the `name`) with:

- `cert.pem` — leaf certificate valid for the provided DNS name(s) and/or IP(s)
- `key.pem` — corresponding private key

By default, leaf certificates are valid for **1 year**. You can customize the validity and subject fields via CLI options.

### Example with Multiple DNS and IPs

```bash
cartman issue --name myservice --dns myservice.local --dns myservice.internal --ip 10.0.0.10 --ip 10.0.0.11
```

## Trusting Cartman CA in Your System

If you want your system to **trust certificates issued by Cartman**, you must add the Cartman root CA to the system trust store.

👉 You typically need to do this **once**, after running `cartman init`, if you want browsers or system tools to trust Cartman-issued certificates.

### macOS

Run in Terminal (with `sudo`):

```bash
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain .cartman/cert.pem
```

### Windows

Run in an elevated PowerShell or Command Prompt:
at
```bash
certutil -addstore Root .cartman\cert.pem
```

### Linux

On most Linux distributions:

```bash
sudo cp .cartman/cert.pem /usr/local/share/ca-certificates/cartman.crt
sudo update-ca-certificates
```

## License

BSD 3-Clause License.
