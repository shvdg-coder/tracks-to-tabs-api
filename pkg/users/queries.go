package users

// insertUserQuery is an SQL query to insert a user in to the 'users' table.
const insertUserQuery = `
	INSERT INTO users (id, email, password)
    VALUES (gen_random_uuid(), $1, $2) 
`

// selectUserPasswordQuery is an SQL query to get the password of a specific user.
const selectUserPasswordQuery = `SELECT password FROM users WHERE email = $1`
