package users

/*
CreateUsersTableQuery is a query to create a users table.
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
const CreateUsersTableQuery = `
	CREATE TABLE IF NOT EXISTS users  (
	   id UUID PRIMARY KEY,
	   email VARCHAR(255) UNIQUE NOT NULL,
	   password VARCHAR(60) NOT NULL
	);
`

// DropUsersTableQuery is an SQL query to drop the 'users' table from the database.
const DropUsersTableQuery = `
	DROP TABLE IF EXISTS users;
`
