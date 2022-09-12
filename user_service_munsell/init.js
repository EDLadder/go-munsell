// admin
db.auth("mongoadm", "mongoadm")

// user
userdb = db.getSiblingDB("munsell_system")
userdb.createUser({
  "user": "nsuser",
  "pwd" : "nsuser",
  "roles": [
    { "role" : "readWrite", "db" : "munsell_system"}
  ],
  "mechanisms": [ "SCRAM-SHA-1" ],
  "passwordDigestor": "client"
})
userdb.auth("nsuser", "nsuser")
