# Golang TMF Service application
- Application is in developing stage

## Application contains:

### TMF 641 v4.2 - Service Ordering Management
- Version 4.2.0
The Service Order API provides a standardized mechanism for managing Service Order, a type of order which can be used to place an order between internal Customer Order management system to service order management system or between a service provider and a partner and vice versa.
#### Service Order resource
- A service order will describe a list of service order items.  A service order item references an action on an existing  or  future  service.  By  service  we  designed  Customer Facing  Service  (CFS)  as  well  as  Resource Facing Service (RFS).
From a component perspective, a service order should be available
- from a Service Orchestration Component (and it could mix CFS and RFS)
- from an Infrastructure Control & Management component (and it would have only RFS)
#### TMF641 performs the following operations on service order resource :
- Retrieval of a service order or a collection of service orders depending on filter criteria
- Partial update of a service order (including updating rules)
- Creation of a service order (including default values and creation rules)
- Deletion of service order (for administration purposes)
- Notification of events on Service order
The REST API for Service Order Management provides a standardized mechanism for placing a service order with all the necessary order parameters.
It allows users to create, update & retrieve Service Orders and manages related notifications.

## Setup notes
### Install Task
- Install task in your local machine, in order to do that you can follow [installation instructions](https://taskfile.dev/#/installation)\
``` 
$ brew install go-task/tap/go-task 
```

```
  $ task --list
  task: Available tasks for this project:
    * build:        Build the app
    * run:          Run the app
    * swagger.doc:      Doc for swagger
    * swagger.gen:      generate Go code
    * swagger.validate:     Validate swagger
```
### Install Go Swagger

- Install ```go-swagger``` tool. [Installation page](https://goswagger.io/go-swagger/install/)
```aiignore
sudo apt update
sudo apt install -y apt-transport-https gnupg curl debian-keyring debian-archive-keyring
```

- Register GPG signing key
```aiignore
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | sudo gpg --dearmor -o /usr/share/keyrings/go-swagger-go-swagger-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' | sudo tee /etc/apt/sources.list.d/go-swagger-go-swagger.list
```
- Install
```aiignore
sudo apt update 
sudo apt install swagger
```

### Validate Swagger specification
```aiignore
$ task swagger.validate.v4
```

### Generate Swagger documentation
```aiignore
$ task swagger.doc.v4
```

### Generate a self-signed TLS for local development
- Install OpenSSL
```aiignore
sudo apt install openssl
```
- Generate a Private Key

- - A private key is a cryptographic key used to encrypt and decrypt data. It's essential for establishing secure communication in TLS/SSL.
```aiignore
openssl genrsa -out tls.key 2048
```
- Create a Certificate Signing Request (CSR)
- - A CSR is a file that contains your public key and organization details. It’s used when requesting a certificate from a Certificate Authority (CA). For a self-signed certificate, it’s still a necessary step as it includes key metadata and helps create the certificate.

```aiignore
openssl req -new -key tls.key -out tls.csr
```
- Generate a Self-Signed Certificate
- - A self-signed certificate is a certificate you sign yourself (instead of relying on a CA). It’s suitable for local development or testing but not recommended for production.
```aiignore
openssl x509 -req -in tls.csr -signkey tls.key -out tls.crt -days 365
```
### Setup ENV file
```aiignore
cp .env.example .env
```
Set configuration parameters for mongo DB.
### Run server
```aiignore
go run main.go --tls-certificate=tls.crt --tls-key=tls.key  --tls-port=34191
```