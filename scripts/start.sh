. scripts/init.sh

docker network create $APP_PREFIX
docker run -d --network $APP_PREFIX --name $APP_PREFIX-postgres -p $START_PORT:5432 -e "POSTGRES_USER=$APP_PREFIX" -e "POSTGRES_PASSWORD=abc@123" -e "POSTGRES_DB=$APP_PREFIX" postgres