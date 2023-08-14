# Check if openssl is installed
if (-not (Get-Command openssl -ErrorAction SilentlyContinue)) {
    Write-Output "openssl could not be found. Please install it and try again."
    exit
}

# Generate a private key for our server
openssl genpkey -algorithm RSA -out localhost.key

# Create a certificate signing request (CSR) for our private key
openssl req -new -key localhost.key -out localhost.csr -config ./local_cert.cnf

# Generate the self-signed certificate using the CSR and the private key, referencing the v3_req extensions from local_cert.cnf
openssl x509 -req -in localhost.csr -signkey localhost.key -out localhost.crt -days 365 -extensions v3_req -extfile ./local_cert.cnf

# Clean up the CSR, we don't need it anymore
Remove-Item localhost.csr
