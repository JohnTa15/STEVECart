#!/bin/bash

docker stack rm smartcart
echo "Stack removed!"
sleep 20

docker stack deploy -c server-docker-swarm-vm.yml smartcart
echo "Deploying the new stack.."

docker service ls
