db.createUser({
  user: 'myuser',
  pwd: 'mypassword',
  roles: [
    {
      role: 'readWrite',
      db: 'mydb',
    },
  ],
});

db.createCollection('users');
db.users.insert({ username: 'admin', password: 'admin' });
