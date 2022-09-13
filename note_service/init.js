// admin
db.auth("mongoadm", "mongoadm")

// user
userdb = db.getSiblingDB("munsell_system")
userdb.createUser({
  "user": "munselluser",
  "pwd" : "123qweASD!A",
  "roles": [
    { "role" : "readWrite", "db" : "munsell_system"}
  ],
  "mechanisms": [ "SCRAM-SHA-1" ],
  "passwordDigestor": "client"
})
userdb.auth("munselluser", "123qweASD!A")