#!/bin/bash 

set -e
pwd && ls -lart && go mod download && go install github.com/air-verse/air@latest && air