#!/bin/bash
# This is for testing and local development purpose only! DONT use in production.

echo "Generating RootCA"
openssl genrsa -out ./entry/keys/root.key 2048
openssl req -new -x509 -nodes -key ./entry/keys/root.key -sha256 -days 3650 -out ./entry/keys/root.pem 

echo "Generating HTTPS Certificates"
openssl req -new -sha256 -nodes -out ./entry/keys/server.csr -newkey rsa:2048 -keyout ./entry/keys/server.key -config ./scripts/https/server.csr.cnf
openssl x509 -req -in ./entry/keys/server.csr -CA ./entry/keys/root.pem  -CAkey ./entry/keys/root.key -CAcreateserial -out ./entry/keys/server.crt -days 3650 -sha256 -extfile ./scripts/https/v3.ext
