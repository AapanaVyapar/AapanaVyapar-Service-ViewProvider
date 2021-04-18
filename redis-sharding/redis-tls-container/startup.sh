sudo ./gen-certs.sh certs

sudo chmod 400 certs/
sudo chown 999 certs/

docker-compose up -d --build
#docker-compose up -d

wait
sleep 30

docker exec -d redis_tls_userview bash -c "redis-server /usr/local/etc/redis/redis.conf --loadmodule /usr/lib/redis/modules/redisearch.so"

sudo chmod 555 -R certs
