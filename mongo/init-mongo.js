db = db.getSiblingDB('creators')
db.createUser({
  user: 'soul',
  pwd: 'u5pssqjPV9q6Liu5pssqjPV9q6Li',
  roles: [
    {
      role: 'root',
      db: 'admin',
    },
  ],
});
