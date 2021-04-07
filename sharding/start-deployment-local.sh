bash ./config-server/start-containers.sh
bash ./shard1/start-containers.sh 
bash ./shard2/start-containers.sh 

wait

bash ./mongos/start-containers.sh

