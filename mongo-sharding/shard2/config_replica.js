rs.initiate(
    {
        _id: "shard2rs",
        members: [
            { _id : 0, host : "192.168.43.189:50004" },
            { _id : 1, host : "192.168.43.189:50005" },
            { _id : 2, host : "192.168.43.189:50006" }
        ]
    }
)

rs.status()