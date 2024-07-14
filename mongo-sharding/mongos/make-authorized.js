db = db.getSiblingDB('admin');
db.auth("api_user", "api_serrate_pass_unknown");

sh.addShard('shard1rs/172.22.0.11,172.22.0.12,172.22.0.13');
sh.status()

sh.addShard('shard2rs/172.22.0.14,172.22.0.15,172.22.0.16');
sh.status()
