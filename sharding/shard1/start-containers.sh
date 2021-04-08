docker-compose -f /home/shitij/go/src/aapanavyapar-service-viewprovider/sharding/shard1/docker-compose.yaml up -d

sleep 35

docker exec -it shard1Server1 bash -c "mongo docker-entrypoint-initdb.d/config_replica.js"
docker exec -it shard1Server1 bash -c "echo 'rs.status()'| mongo"
