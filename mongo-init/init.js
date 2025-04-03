db.createUser({
    user: "${MONGO_DB_USERNAME}",
    pwd: "${MONGO_DB_PASSWORD}",
    roles: [{
        role: "readWrite",
        db: "${MONGO_DB_DATABASE}"
    }]
});