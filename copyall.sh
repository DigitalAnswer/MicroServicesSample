#!/bin/bash

cd accountservice
go get
GOARCH=amd64 GOOS=linux go build -o accountservice-linux-amd64
echo built `pwd`
cd ..

cd healthchecker
go get
GOARCH=amd64 GOOS=linux go build -o healthchecker-linux-amd64
echo built `pwd`
cd ..

cp healthchecker/healthchecker-linux-amd64 accountservice/

docker build -t benprieur/accountservice accountservice/

# docker service rm accountservice
# docker service create --name=accountservice --replicas=1 --network=my_network -p=8080:8080 benprieur/accountservice