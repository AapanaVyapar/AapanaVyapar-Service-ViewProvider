db = db.getSiblingDB('admin');
db.auth("api_user", "api_serrate_pass_unknown");

sh.addShard('shard1rs/192.168.43.189:50001,192.168.43.189:50002,192.168.43.189:50003');
sh.status()

sh.addShard('shard2rs/192.168.43.189:50004,192.168.43.189:50005,192.168.43.189:50006');
sh.status()
