db = db.getSiblingDB('admin');
db.auth("api_user", "api_serrate_pass_unknown");

db = db.getSiblingDB('db_aapanavypar');

db.createCollection('userData');
db.createCollection('orderData');
db.createCollection('shopData');
db.createCollection('productData');
db.createCollection('analyticalData');

rs.initiate(
    {
        _id: "shard1rs",
        members: [
            { _id : 0, host : "192.168.8.100:50001" },
            { _id : 1, host : "192.168.8.100:50002" },
            { _id : 2, host : "192.168.8.100:50003" }
        ]
    }
)

rs.status()

//sh.shardCollection("db_aapanavypar.productData", { _id: "hashed" }, true)
