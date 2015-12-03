#!/bin/bash

database="releasifier_db"
user="releasifier_agent"
password="98uhi4q3brjfnsdlzisw2"


# removing tracing
psql -U postgres <<< "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '$database' AND pid <> pg_backend_pid();"

# dropping database
psql -U postgres <<< "DROP DATABASE $database;"

# create database react_native_updater_db
psql -U postgres <<< "CREATE DATABASE $database ENCODING 'UTF-8' LC_COLLATE='en_US.UTF-8' LC_CTYPE='en_US.UTF-8' TEMPLATE template0;"

# create username and password
psql -U postgres <<< "CREATE USER $user WITH PASSWORD '$password';"

# updated schema
cat /migration/users.sql | psql -U $user -d $database
cat /migration/apps.sql | psql -U $user -d $database
cat /migration/apps_users_permissions.sql | psql -U $user -d $database
cat /migration/releases.sql | psql -U $user -d $database
cat /migration/bundles.sql | psql -U $user -d $database
