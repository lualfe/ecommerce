. scripts/init.sh

docker stop $APP_PREFIX-postgres
docker rm $APP_PREFIX-postgres
docker network rm $APP_PREFIX