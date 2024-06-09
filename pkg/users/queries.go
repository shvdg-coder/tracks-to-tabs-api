package users

const createUsersTableQuery = `
	CREATE TABLE IF NOT EXISTS users  (
	   ID UUID PRIMARY KEY,
	   email VARCHAR(255) UNIQUE NOT NULL,
	   password VARCHAR(60) NOT NULL
	);
`

const dropUsersTableQuery = `
	DROP TABLE IF EXISTS users;
`

const insertUserQuery = `
	INSERT INTO users (id, email, password)
    VALUES (gen_random_uuid(), $1, $2) 
`

const selectUserPasswordQuery = `SELECT password FROM users WHERE email = $1`
