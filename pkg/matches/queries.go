package matches

// With Artist, Tracks, Tabs, and all lookups
const createTabsView = `
	CREATE VIEW v_tabs AS
	SELECT 
		t.ID AS tab_id,
		t.InstrumentID,
		i.name AS instrument_name,
		t.DifficultyID,
		d.name AS difficulty_name,
		t.Description,
		at.ArtistID,
		a.Name AS artist_name,
		tt.TrackID,
		tr.Title AS track_title,
		tr.Duration AS track_duration
	FROM tabs AS t
	INNER JOIN instruments AS i ON t.InstrumentID = i.ID
	INNER JOIN difficulties AS d ON t.DifficultyID = d.ID
	INNER JOIN track_tab AS tt ON t.ID = tt.TabID
	INNER JOIN artist_track AS at ON tt.TrackID = at.TrackID
	INNER JOIN artists AS a ON at.ArtistID = a.ID
	INNER JOIN tracks AS tr ON tt.TrackID = tr.ID;
`
