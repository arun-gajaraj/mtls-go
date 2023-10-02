# mTLS setup
# using openssl

# CA
# ------

# create a private key for Certificate Authority (CA)
openssl genpkey -algorithm RSA -out ca.key.pem

# create a self signed CA certificate
openssl req -x509 -new -nodes -key ca.key.pem -days 365 -out ca.pem 



# Server 
# ------

# create a server key
openssl genpkey -algorithm RSA -out server-key.pem

# create server cert signing request 
openssl req -new -key server-key.pem -out server.csr

# create server certificate signed by CA
openssl x509 -req -in server.csr -CA ca.pem -CAkey ca.key.pem -out server-cert.pem -days 365



# Client
# ------

# create client key
openssl genpkey -algorithm RSA -out client-key.pem

# create client csr 
openssl req -new -key client-key.pem -out client.csr

# create client certificate signed by CA
openssl x509 -req -in client.csr -CA ca.pem -CAkey ca.key.pem -days 365 -out client-cert.pem
