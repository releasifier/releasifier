#!/bin/bash

database="releasifier_db"
user="releasifier_agent"
password="98uhi4q3brjfnsdlzisw2"

# create postgres USER as SUPERUSER
psql -h127.0.0.1 <<< "CREATE USER postgres SUPERUSER;"

# removing tracing
psql -h127.0.0.1 -U postgres <<< "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '$database' AND pid <> pg_backend_pid();"

# dropping database
psql -h127.0.0.1 -U postgres <<< "DROP DATABASE $database;"

# create database react_native_updater_db
psql -h127.0.0.1 -U postgres <<< "CREATE DATABASE $database ENCODING 'UTF-8' LC_COLLATE='en_US.UTF-8' LC_CTYPE='en_US.UTF-8' TEMPLATE template0;"

# create username and password
psql -h127.0.0.1 -U postgres <<< "CREATE USER $user WITH PASSWORD '$password';"

# updated schema
cat data/schemas/users.sql | psql -h127.0.0.1 -U $user -d $database
cat data/schemas/apps.sql | psql -h127.0.0.1 -U $user -d $database
cat data/schemas/apps_users_permissions.sql | psql -h127.0.0.1 -U $user -d $database
cat data/schemas/releases.sql | psql -h127.0.0.1 -U $user -d $database
cat data/schemas/bundles.sql | psql -h127.0.0.1 -U $user -d $database
