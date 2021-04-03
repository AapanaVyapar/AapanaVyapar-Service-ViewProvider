rs.initiate(
    {
        _id: "cfgrs",
        configsvr: true,
        members: [
            { _id : 0, host : "configServer1" },
            { _id : 1, host : "configServer2" },
            { _id : 2, host : "configServer3" }
        ]
    }
)

rs.status()