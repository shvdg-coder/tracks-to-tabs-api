package tabs

/*
+--------------------------------------+---------------+---------------+------------------+
|                  id                  | instrument_id | difficulty_id |   description    |
+--------------------------------------+---------------+---------------+------------------+
| 123e4567-e89b-12d3-a456-426614174000 |      580      |      423      | James Hetfield   |
| 123e4567-e89b-12d3-a456-426614174001 |      590      |      420      | Mick Mars        |
+--------------------------------------+---------------+---------------+------------------+

This table is used to store Tracks in our system.

It contains the following columns:
  - 'id': This is the UUID that uniquely identifies a record.
  - 'instrument_id': This column represents the ID of the instrument from a lookup table.
  - 'difficulty_id': This column represents the ID of the difficulty level from a lookup table.
  - 'description': This column records the description of the tab.
*/
const createTabsTableQuery = `
	CREATE TABLE IF NOT EXISTS tabs (
	   id UUID PRIMARY KEY,
	   instrument_id INT NOT NULL,
	   difficulty_id INT NOT NULL,
	   description TEXT,
	   CONSTRAINT fk_instrument FOREIGN KEY(instrument_id) REFERENCES instruments(id),
	   CONSTRAINT fk_difficulty	FOREIGN KEY(difficulty_id) REFERENCES difficulties(id)
	);
`

// dropTabsTableQuery is a SQL query to drop the 'tabs' table if it exists
const dropTabsTableQuery = `
	DROP TABLE IF EXISTS tabs;
`

// insertTabQuery is a SQL query string used to insert a tab into the 'tabs' table.
const insertTabQuery = `
	INSERT INTO tabs (id, instrument_id, difficulty_id, description)
    VALUES ($1, $2, $3, $4) 
`

// getTabsQuery is a SQL query sting used to retrieve tabs for the provided IDs.
const getTabsQuery = `SELECT id, instrument_id, difficulty_id, description FROM tabs WHERE id IN ($1)`
