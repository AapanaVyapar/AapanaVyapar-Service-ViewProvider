db = db.getSiblingDB('admin');
db.auth("api_user", "api_serrate_pass_unknown");

rs.initiate(
    {
        _id: "cfgrs",
        configsvr: true,
        members: [
            { _id : 0, host : "172.22.0.17" },
            { _id : 1, host : "172.22.0.18" },
            { _id : 2, host : "172.22.0.19" }
        ]
    }
)

rs.status()
