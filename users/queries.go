package users

const createUsersTableQuery = `
	CREATE TABLE IF NOT EXISTS users  (
	   ID UUID PRIMARY KEY,
	   email VARCHAR(255) UNIQUE NOT NULL,
	   password VARCHAR(255) NOT NULL
	);
`

const insertUserQuery = `
	INSERT INTO users (id, email, password)
    VALUES (gen_random_uuid(), $1, $2) 
    ON CONFLICT DO NOTHING;
`

const selectUserPasswordQuery = `SELECT password FROM users WHERE email = $1`
