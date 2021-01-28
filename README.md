# Delinkcious

A delicious-like link management platform implemented using Go microservices


# Directory Structure

## pkg
The core logic is implemented by libraries in this directory

## svc

The microservices are in this directory. They use the excellent [gokit](https://gokit.io) microservice framework.


## cmd

Various utilities and one-time commands live here


docker run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_DB=social_graph_manager postgres

docker run -e ALLOW_NONE_AUTHENTICATION=yes -it --rm --name etcd bitnami/etcd

ALLOW_NONE_AUTHENTICATION 

docker run -p 2379:2379 -e ALLOW_NONE_AUTHENTICATION=yes -it --name etcd bitnami/etcd:3.3.15


docker stop $(docker ps -a -q)
docker container prune



undefined: balancer.PickOptions
https://github.com/etcd-io/etcd/issues/12577

https://github.com/etcd-io/etcd/issues/12124

// google.golang.org/grpc v1.34.0 // indirect
google.golang.org/grpc v1.26.0

v1.26.0 is ok
