openssl genrsa -out ./entry/server.key 2048
openssl req -new -x509 -key ./entry/server.key -out ./entry/server.pem -days 365
