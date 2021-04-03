db.getSiblingDB("admin").auth("shitij_cluster", "password_shitij");
db.getSiblingDB("admin").auth("shitij", "password_shitij");

sh.addShard('shard1rs/shard1Server1,shard1Server2,shard1Server3');
sh.status()

sh.addShard('shard2rs/shard2Server1,shard2Server2,shard2Server3');
sh.status()
