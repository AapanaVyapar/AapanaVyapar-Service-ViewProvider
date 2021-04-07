docker-compose -f /home/shitij/go/src/aapanavyapar-service-viewprovider/sharding/config-server/docker-compose.yaml up -d

sleep 30

docker exec -it configServer1 bash -c "mongo docker-entrypoint-initdb.d/config_replica.js"
docker exec -it configServer1 bash -c "echo 'rs.status()'| mongo"
