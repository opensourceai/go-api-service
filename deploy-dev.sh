#!/usr/bin/env bash
set -e

docker-compose down --rmi all && docker-compose up -d --build