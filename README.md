# MaxChat Project

Welcome to the MaxChat project! This guide will help you set up the application for local development.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup Guide](#setup-guide)
  - [Clone the Repository](#clone-the-repository)
  - [Install Dependencies](#install-dependencies)
  - [Install OpenSSL](#install-openssl)
  <!-- - [Set OPENSSL_CFG_PATH](#set-openssl_cfg_path) -->
  - [Generate Certificates](#generating-local-development-certificates)
  - [Run the Server](#run-the-server)
  - [Run the Client](#run-the-client)
<!-- - [Additional Configuration](#additional-configuration) -->
<!-- - [Access the Application](#access-the-application) -->
- [Contributing](#contributing)

## Prerequisites
Ensure that you have the following installed on your system:
- Git
- Golang
- OpenSSL

## Setup Guide

### Clone the Repository

```bash
git clone https://github.com/maxsg5/maxchat.git maxchat
cd maxchat
```

## Install Dependencies

```
go mod download
```

## Install OpenSSL
- Windows: Download from [OpenSSL's official site.](https://www.openssl.org/source/)
- macOS: Use [Homebrew](https://brew.sh/). 
```
brew install openssl
```
- Linux (e.g., Ubuntu/Debian):
```
sudo apt-get update
sudo apt-get install openssl
```



## Generating Local Development Certificates

### Generate the private key and public certificate:
Navigate to the scripts directory and run the appropriate script for your platform:
- Windows (PowerShell):
```
.\generate-certs.ps1`
```
- macOS/Linux (Bash):
```
./generate-certs.sh
```
This will produce two files in your directory: localhost.key (private key) and localhost.crt (public certificate). Use these for local development.

## Run the Server
navigate to the server directory and run the following command:
```
go run server.go
```

## Run the Client
navigate to the client directory and run the following command:
```
go run client.go
```

## Contributing
TODO: Add contributing guidelines.