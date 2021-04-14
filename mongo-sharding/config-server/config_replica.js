db = db.getSiblingDB('admin');
db.auth("api_user", "api_serrate_pass_unknown");

rs.initiate(
    {
        _id: "cfgrs",
        configsvr: true,
        members: [
            { _id : 0, host : "192.168.43.189:40001" },
            { _id : 1, host : "192.168.43.189:40002" },
            { _id : 2, host : "192.168.43.189:40003" }
        ]
    }
)

rs.status()
