try {
    admin = db.getSiblingDB("admin")
    admin.createUser(
        {
            user: "shitij",
            pwd: "password_shitij",     // or cleartext password
            roles: [{role: "userAdminAnyDatabase", db: "admin"}]
        }
    )

    db.getSiblingDB("admin").auth("shitij", "password_shitij")

    db.getSiblingDB("admin").createUser(
        {
            "user" : "shitij_cluster",
            "pwd" : "password_shitij",     // or cleartext password
            roles: [ { "role" : "clusterAdmin", "db" : "admin" } ]
        }
    )

    db.getSiblingDB("admin").auth("shitij_cluster", "password_shitij")
}catch (e) {

}
