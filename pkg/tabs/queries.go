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

/*
+------------------+--------------+--------------+--------------------+------------------+---------------+-------------------+------------------+----------------+------------------+----------------------+
|  artist_id       | artist_name  |  track_id    |   track_title      |  track_duration  |     tab_id    |   instrument_id   | instrument_name  | difficulty_id  | difficulty_name  | description          |
+---------------------------------+--------------+--------------------+------------------+---------------+-------------------+------------------+----------------+------------------+----------------------+
| 234e5678-e90c... | Sting        |     600      | Englishman in NY   |   140305         | 123e4567-e... |  700              |  Bass guitar     | 500            | Intermediate     | Sting playing bass   |
| 234e5678-e90c... | B.B. King    |     650      | The Thrill is Gone |   124022         | 123e4569-e... |  785              |  Electric guitar | 530            | Advanced         | B.B. King's solo     |
+------------------+--------------+--------------+--------------------+------------------+---------------+-------------------+------------------+----------------+------------------+----------------------+

This view is used to display a comprehensive listing of Artist, Track, and Tab information in our system.

The view combines several tables and consists of the following columns:
  - 'artist_id': The ID of the artist from the `artists` table.
  - 'artist_name': The name of the artist from the `artists` table.
  - 'track_id': The ID of the track from the `tracks` table.
  - 'track_title': The title of the track from the `tracks` table.
  - 'track_duration': The duration in milliseconds of the track from the `tracks` table.
  - 'tab_id': The ID of the tab from the `tabs` table.
  - 'instrument_id': The ID of the instrument from the `instruments` table.
  - 'instrument_name': The name of the instrument from the `instruments` table.
  - 'difficulty_id': The ID of the difficulty from the `difficulties` table.
  - 'difficulty_name': The name of the difficulty from the `difficulties` table.
  - 'description': The description of the tab from the `tabs` table.
*/
const createTabsViewQuery = `
	CREATE VIEW v_tabs AS
	SELECT 
		a.id AS artist_id,
		a.name AS artist_name,
		tr.id AS track_id,
		tr.title AS track_title,
		tr.duration AS track_duration,
		t.id AS tab_id,
		t.instrument_id,
		i.name AS instrument_name,
		t.difficulty_id,
		d.name AS difficulty_name,
		t.description
	FROM artist_track AS at
	INNER JOIN artists AS a ON at.artist_id = a.id
	INNER JOIN track_tab AS tt ON at.track_id = tt.track_id
	INNER JOIN tracks AS tr ON tt.track_id = tr.id
	INNER JOIN tabs AS t ON tt.tab_id = t.id
	INNER JOIN instruments AS i ON t.instrument_id = i.id
	INNER JOIN difficulties AS d ON t.difficulty_id = d.id;
`

// dropTabsViewQuery is a SQL query to drop the 'tabs' view if it exists
const dropTabsViewQuery = `
	DROP VIEW IF EXISTS "v_tabs";
`
