package tabs

/*
+--------------------------------------+---------------+---------------+-------------+------------------+
|                  ID                  | InstrumentID  | DifficultyID  |  TuningID   |   Description    |
+--------------------------------------+---------------+---------------+-------------+------------------+
| 123e4567-e89b-12d3-a456-426614174000 |      580      |      423      |     300     | James Hetfield   |
| 123e4567-e89b-12d3-a456-426614174001 |      590      |      420      |     324     | Mick Mars        |
+--------------------------------------+---------------+---------------+-------------+------------------+

This table is used to store Tracks in our system.

It contains the following columns:
  - 'ID': This is the UUID that uniquely identifies a record.
  - 'InstrumentID': This column represents the ID of the instrument from a lookup table.
  - 'DifficultyID': This column represents the ID of the difficulty level from a lookup table.
  - 'TuningID': This column represents the ID of the tuning from a lookup table.
  - 'Description': This column records the description of the tab.
*/
const createTabsTableQuery = `
	CREATE TABLE IF NOT EXISTS tabs (
	   ID UUID PRIMARY KEY,
	   InstrumentID INT NOT NULL,
	   DifficultyID INT NOT NULL,
	   TuningID INT NOT NULL,
	   Description TEXT
	);
`

// dropTabsTableQuery is a SQL query to drop the 'tabs' table if it exists
const dropTabsTableQuery = `
	DROP TABLE IF EXISTS tabs;
`

// insertTabQuery is a SQL query string used to insert a tab into the 'tabs' table.
const insertTabQuery = `
	INSERT INTO tabs (id, instrumentId, difficultyId, tuningId, description)
    VALUES (gen_random_uuid(), $1, $2, $3, $4) 
`
