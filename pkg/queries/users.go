package queries

// InsertUserQuery is an SQL query to insert a user in to the 'users' table.
const InsertUserQuery = `
	INSERT INTO users (id, email, password)
    VALUES (gen_random_uuid(), $1, $2) 
`

// SelectUserPasswordQuery is an SQL query to get the password of a specific user.
const SelectUserPasswordQuery = `SELECT password FROM users WHERE email = $1`
