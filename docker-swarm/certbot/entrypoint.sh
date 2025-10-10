#!/bin/bash
set -e

DOMAIN="steve-cart.ydns.eu"
EMAIL="ice18390290@uniwa.gr"

echo "=== Generating SSL certificates for $DOMAIN ==="

certbot certonly --standalone \
  -d $DOMAIN \
  --agree-tos \
  --no-eff-email \
  -m $EMAIL \
  --non-interactive

echo "=== Starting main service ==="

exec "$@"