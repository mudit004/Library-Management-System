#!/bin/bash

read -p "Enter your MySQL Username" DB_USERNAME
read -p -s "Enter your MySQL Password" DB_PASSWORD
read -p "Enter your MySQL Host:port" DB_HOST
read -p "Enter your Database Name" DB_NAME
read -p "Enter your JWT Secret Key" JWT_SECRET_KEY

echo "---------------------------------------------------"

yaml_content=$(cat <<EOF
DB_USERNAME: "$DB_USERNAME"
DB_PASSWORD: "$DB_PASSWORD"
DB_HOST: "$DB_HOST"
DB_NAME: "$DB_NAME"
JWT_SECRET_KEY: "$JWT_SECRET_KEY"
EOF
)

echo "$yaml_content" > config.yaml

migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST)/$DB_NAME" force 1
migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST)/$DB_NAME" -verbose up

echo "Congratulations the yaml file and database setup is complete"
