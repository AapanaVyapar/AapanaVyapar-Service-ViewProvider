docker-compose -f /home/shitij/go/src/aapanavyapar-service-viewprovider/sharding/config-server/docker-compose.yaml up -d

sleep 35

docker exec -it configServer1 bash -c "mongo docker-entrypoint-initdb.d/config_replica.js"
docker exec -it configServer1 bash -c "echo 'rs.status()'| mongo"

docker exec -it configServer1 bash -c "mongo docker-entrypoint-initdb.d/create-users.js"
docker exec -it configServer2 bash -c "mongo docker-entrypoint-initdb.d/create-users.js"
docker exec -it configServer3 bash -c "mongo docker-entrypoint-initdb.d/create-users.js"
