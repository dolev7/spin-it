#!/bin/bash
export $(grep -v '^#' .env | xargs)
migrate -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" -path migrations up
