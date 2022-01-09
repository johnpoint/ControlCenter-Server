#!/usr/bin/env bash
set -x
export USERID=${UID}
export GROUPID=${GID}
export RABBITMQ_USERNAME="admin"
export RABBITMQ_PASSWORD="password"
export MONGODB_USERNAME="admin"
export MONGODB_PASSWORD="password"
export INFLUXDB_USERNAME="admin"
export INFLUXDB_PASSWORD="password"
export REDIS_PASSWORD="123456"

if [[ $1 = "up" ]]; then
  if [[ $2 = "" ]]; then
    docker-compose -f base_service.yaml up -d
    docker-compose -f server.yaml up -d
  fi
  if [[ $2 = "base" ]]; then
    docker-compose -f base_service.yaml up -d
  fi
  if [[ $2 = "dev" ]]; then
    docker-compose -f server.yaml up -d
  fi
fi

if [[ $1 = "down" ]]; then
  if [[ $2 = "" ]]; then
    docker-compose -f base_service.yaml down
    docker-compose -f server.yaml down
  fi
  if [[ $2 = "base" ]]; then
    docker-compose -f base_service.yaml down
  fi
  if [[ $2 = "dev" ]]; then
    docker-compose -f server.yaml down
  fi
fi
