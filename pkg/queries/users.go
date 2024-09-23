package queries

/*
CreateUsersTable is a query to create a users table.
+--------------------------------------+---------------+-------------+
|                   id                 |     email     | password 	 |
+--------------------------------------+---------------+-------------+
| 123e4567-e89b-12d3-a456-426614174000 | john@doe.com  | hashedPw123 |
| 123e4567-e89b-12d3-a456-426614174001 | jane@doe.com  | hashedPw456 |
+--------------------------------------+---------------+-------------+

It consists of the following columns:
  - 'id': This is the UUID that uniquely identifies a user in our system.
  - 'email': This is the user's email address.
  - 'password': This stores the hashed password of the user.
*/
const CreateUsersTable = `
	CREATE TABLE IF NOT EXISTS users  (
	   id UUID PRIMARY KEY,
	   email VARCHAR(255) UNIQUE NOT NULL,
	   password VARCHAR(60) NOT NULL
	);

	CREATE INDEX idx_user_id ON users(id);
`

// DropUsersTable is an SQL query to drop the 'users' table from the database.
const DropUsersTable = `
	DROP TABLE IF EXISTS users;
`

// InsertUser is an SQL query to insert a user in to the 'users' table.
const InsertUser = `
	INSERT INTO users (id, email, password)
    VALUES (gen_random_uuid(), $1, $2) 
`

// SelectUserPassword is an SQL query to get the password of a specific user.
const SelectUserPassword = `SELECT password FROM users WHERE email = $1`
