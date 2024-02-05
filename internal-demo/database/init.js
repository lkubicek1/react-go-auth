// Create the users collection with an index on username
db.createCollection('users');
db.users.createIndex({ username: 1 }, { unique: true });

// Create the sessions collection with necessary indexes
db.createCollection('sessions');
db.sessions.createIndex({ session_id: 1 }, { unique: true });
db.sessions.createIndex({ expires_at: 1 }, { expireAfterSeconds: 0 }); // TTL index for auto-deletion of expired sessions
