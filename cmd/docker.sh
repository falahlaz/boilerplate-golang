docker network create -d bridge db-network
docker network create -d bridge redis-network
docker volume create dbdata
docker volume create db2data
docker volume create redisdata

docker run --name db-boiler-plate -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=supertestuser -e POSTGRES_DB=boilerplate_golang_db -p 5432:5432 -v dbdata:/postgresql/lib/data --network=db-network -d postgres
docker run --name db2-boiler-plate -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=supertestuser -e POSTGRES_DB=boilerplate_golang_db -p 5433:5432 -v db2data:/postgresql/lib/data --network=db-network -d postgres
docker run --name redis-boiler-plate -e REDIS_PASSWORD=testingredis -v redisdata:/redis/lib/data --network=redis-network -p 6379:6379 -d redis
docker build -t boiler-plate .
docker run --name boiler-plate-go --network=db-network -dp 8080:3000 boiler-plate .

docker container stop boiler-plate-go
docker container rm boiler-plate-go
docker image rm boiler-plate