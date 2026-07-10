#!/bin/bash

echo "Building images.."
docker build -t steve-api:1.0 ./steve-api
docker build -t steve-frontend:1.0 ./steve_display

docker stack rm smartcart
echo "Stack removed!"
sleep 15

echo "Deploying the new stack.."
docker stack deploy -c server-docker-swarm-vm.yml smartcart

sleep 20
docker service ls