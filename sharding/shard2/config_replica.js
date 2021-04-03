rs.initiate(
    {
        _id: "shard2rs",
        members: [
            { _id : 0, host : "shard2Server1" },
            { _id : 1, host : "shard2Server2" },
            { _id : 2, host : "shard2Server3" }
        ]
    }
)

rs.status()