#!/bin/bash

docker stack rm smartcart
echo "Stack removed!"
sleep 15

echo "Deploying the new stack.."
docker stack deploy -c server-docker-swarm-vm.yml smartcart

sleep 20
docker service ls

echo "Removing the temp vite modules"
sudo rm -rf /home/johnt/Documents/STEVEcart/STEVECart/server/steve_display/node_modules/.vite

echo "Deploying the frontend"
cd /home/johnt/Documents/STEVEcart/STEVECart/server/steve_display
npm run dev