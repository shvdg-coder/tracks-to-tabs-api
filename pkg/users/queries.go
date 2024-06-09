package users

/*
+--------------------------------------+---------------+-------------+
|                   ID                 |     Email     | Password 	 |
+--------------------------------------+---------------+-------------+
| 123e4567-e89b-12d3-a456-426614174000 | john@doe.com  | hashedPw123 |
| 123e4567-e89b-12d3-a456-426614174001 | jane@doe.com  | hashedPw456 |
+--------------------------------------+---------------+-------------+

This table is used to store a user with their credentials in our system.

It consists of the following columns:
  - 'ID': This is the UUID that uniquely identifies a user in our system.
  - 'Email': This is the user's email address.
  - 'Password': This stores the hashed password of the user.
*/
const createUsersTableQuery = `
	CREATE TABLE IF NOT EXISTS users  (
	   ID UUID PRIMARY KEY,
	   email VARCHAR(255) UNIQUE NOT NULL,
	   password VARCHAR(60) NOT NULL
	);
`

// dropUsersTableQuery is an SQL query to drop the 'users' table from the database.
const dropUsersTableQuery = `
	DROP TABLE IF EXISTS users;
`

// insertUserQuery is an SQL query to insert a user in to the 'users' table.
const insertUserQuery = `
	INSERT INTO users (id, email, password)
    VALUES (gen_random_uuid(), $1, $2) 
`

// selectUserPasswordQuery is an SQL query to get the password of a specific user.
const selectUserPasswordQuery = `SELECT password FROM users WHERE email = $1`
