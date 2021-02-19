# golang-gin



## Go Module Init

```bash
go mod init gitlab.com/alansuparlan/golang-gin
```

## Gin-Gonic library: github.com/gin-gonic/gin

## Run

```bash
go run server.go
```

# Swagger Documentation

```bash
go get -u github.com/swaggo/swag/cmd/swag
```


## Generate Swagger Documentation

```bash
swag init
```

setelah nanti di masuk Swagger Ui masukan jwt token diawali dengan Bearer

# Deploy on AWS ElasticBeanstalk from EB CLI

## 1.- Create user (e.g "beanstalk") and allow Programmatic Access
## 2.- Create new group ( e.g. "Beanstalk")
## 3.- Attach AWSElasticBeanstalkFullAccess policy to the group
## 4.- Add the user to the Group
## 5.- Copy user's aws_access_key_id and aws_secret_access_key to .aws/config file
## 6.- From the application directory run:
## 6.a.- eb init
## 6.b.- eb create --single

# Deploy on AWS ElasticBeanstalk with Docker

## Build the docker image

```bash
docker build --tag alansuparlan/golang-gin .
```

## Run the container locally

```bash
docker run -p 5000:5000 alansuparlan/golang-gin
```

## Push the image to DockerHub (you need a DockerHub account)

```bash
docker login
```

```bash
docker push alansuparlan/golang-gin
```




> ### Golang clean-architecture codebase example containing real world examples (CRUD, auth, advanced patterns, etc) that adheres to the [RealWorld](https://github.com/gothinkster/realworld) spec and API.

### [Demo](https://github.com/gothinkster/realworld)&nbsp;&nbsp;&nbsp;&nbsp;[RealWorld](https://github.com/gothinkster/realworld)

# Getting started
### Build the app 
```
make
```
### Run the app
```bash
./go-realworld-clean
```

### Run the integration tests
Start the server with an existing user
```
./go-realworld-clean --populate=true
```

In another terminal, run the tests against the API
```
newman run api/Conduit.postman_collection.json \
  -e api/Conduit.postman_integration_test_environment.json \
  --global-var "EMAIL=joe@what.com" \
  --global-var "PASSWORD=password"
```
# Additional
## Make Targets

The version is either `0.1.0` if no tag has ever been defined or the latest
tag defined. The build number is the SHA1 of the latest commit.

- **make**: Builds and injects version/build in binary
- **make init**: Sets the pre-commit hook in the repository
- **make docker**: Build docker image and tag it with both `latest` and version
- **make latest**: Build docker image and tag it only with `latest`
- **make test**: Executes the test suite
- **make mock**: Generate the necessary mocks
- **make clean**: Removes the built binary if present

