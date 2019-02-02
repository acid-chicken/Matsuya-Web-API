db.createUser({
    user: 'matsuya',
    pwd: 'matsuya',
    roles:
        [
            {
                role: 'userAdmin',
                db: 'matsuya'
            }
        ]
})