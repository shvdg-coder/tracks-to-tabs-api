package queries

/*
CreateTabsTable is a query to create a tracks table.
+--------------------------------------+---------------+---------------+------------------+
|                  id                  | instrument_id | difficulty_id |   description    |
+--------------------------------------+---------------+---------------+------------------+
| 123e4567-e89b-12d3-a456-426614174000 |      580      |      423      | James Hetfield   |
| 123e4567-e89b-12d3-a456-426614174001 |      590      |      420      | Mick Mars        |
+--------------------------------------+---------------+---------------+------------------+

It contains the following columns:
  - 'id': This is the UUID that uniquely identifies a record.
  - 'instrument_id': This column represents the ID of the instrument from a lookup table.
  - 'difficulty_id': This column represents the ID of the difficulty level from a lookup table.
  - 'description': This column records the description of the tab.
*/
const CreateTabsTable = `
	CREATE TABLE IF NOT EXISTS tabs (
	   id UUID PRIMARY KEY,
	   instrument_id INT NOT NULL,
	   difficulty_id INT NOT NULL,
	   description TEXT,
	   CONSTRAINT fk_instrument FOREIGN KEY(instrument_id) REFERENCES instruments(id),
	   CONSTRAINT fk_difficulty	FOREIGN KEY(difficulty_id) REFERENCES difficulties(id)
	);

	CREATE INDEX idx_tab_id ON tabs(id);
`

// DropTabsTable is a SQL query to drop the 'tabs' table if it exists
const DropTabsTable = `
	DROP TABLE IF EXISTS tabs;
`

// GetTabs is a SQL query sting used to retrieve tabs for the provided IDs.
const GetTabs = `SELECT id, instrument_id, difficulty_id, description FROM tabs WHERE id = ANY($1::uuid[])`
