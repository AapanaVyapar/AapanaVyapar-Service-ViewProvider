rs.initiate(
    {
        _id: "shard2rs",
        members: [
            { _id : 0, host : "172.22.0.14" },
            { _id : 1, host : "172.22.0.15" },
            { _id : 2, host : "172.22.0.16" }
        ]
    }
)

rs.status()
