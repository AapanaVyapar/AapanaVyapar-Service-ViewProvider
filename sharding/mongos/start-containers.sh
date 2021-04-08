docker-compose -f /home/shitij/go/src/aapanavyapar-service-viewprovider/sharding/mongos/docker-compose.yaml up -d

sleep 35

docker exec -it mongos bash -c "mongo /docker-entrypoint-initdb.d/make-authorized.js"
