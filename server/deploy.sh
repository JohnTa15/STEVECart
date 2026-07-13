#!/bin/bash

echo "Building images.."
docker build -t steve-api:1.0 ./steve-api
docker build -t steve-frontend:1.0 ./steve_display

# Hash (bcrypt) the default superadmin password and inject it into init.sql.
echo "Hashing superadmin password.."
# Runs a temporary (throwaway) container just to borrow the htpasswd tool.
SUPERADMIN_HASH=$(docker run --rm httpd:alpine htpasswd -bnBC 10 "" superadmin123 | tr -d ':\n')
echo "Superadmin bcrypt hash: $SUPERADMIN_HASH"

#adding to the hash to the code block in db/init.sql
sed -i "s|VALUES ('superadmin', 'superadmin@uniwa.gr', '[^']*'|VALUES ('superadmin', 'superadmin@uniwa.gr', '$SUPERADMIN_HASH'|" db/init.sql

docker stack rm smartcart
echo "Stack removed!"
sleep 15

echo "Deploying the new stack.."
docker stack deploy -c server-docker-swarm-vm.yml smartcart

sleep 20
docker service ls