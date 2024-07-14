docker-compose up -d
wait

sleep 60

docker exec -it configServer1 bash -c "mongo docker-entrypoint-initdb.d/config_replica.js"
docker exec -it configServer1 bash -c "echo 'rs.status()'| mongo"

docker exec -it shard1Server1 bash -c "mongo docker-entrypoint-initdb.d/config_replica.js"
docker exec -it shard1Server1 bash -c "echo 'rs.status()'| mongo"

docker exec -it shard2Server1 bash -c "mongo docker-entrypoint-initdb.d/config_replica.js"
docker exec -it shard2Server1 bash -c "echo 'rs.status()'| mongo"

wait

docker stop mongos
docker rm mongos

docker-compose up -d mongos

wait

sleep 5

docker exec -it mongos bash -c "mongo /docker-entrypoint-initdb.d/make-authorized.js"

