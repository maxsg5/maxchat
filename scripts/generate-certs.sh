#!/bin/bash

# Check if openssl is installed
if ! command -v openssl &> /dev/null; then
    echo "openssl could not be found. Please install it and try again."
    exit 1
fi

# Generate a private key for our server
openssl genpkey -algorithm RSA -out server.key

# Create a certificate signing request (CSR) for our private key
openssl req -new -key server.key -out server.csr -config ./local_cert.cnf

# Generate the self-signed certificate using the CSR and the private key, referencing the v3_req extensions from local_cert.cnf
openssl x509 -req -in server.csr -signkey server.key -out server.crt -days 365 -extensions v3_req -extfile ./local_cert.cnf

# Clean up the CSR, we don't need it anymore
rm server.csr
