# Service Ordering Management API
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