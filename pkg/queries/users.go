package queries

// InsertUser is an SQL query to insert a user in to the 'users' table.
const InsertUser = `
	INSERT INTO users (id, email, password)
    VALUES (gen_random_uuid(), $1, $2) 
`

// SelectUserPassword is an SQL query to get the password of a specific user.
const SelectUserPassword = `SELECT password FROM users WHERE email = $1`
