package views

const createTabsView = `
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

const createSourcesToEndpointsView = `
	CREATE VIEW v_source_endpoints AS
		SELECT sources.id as source_id, sources.name as source_name, sources.category AS source_category, 
			   endpoints.category AS endpoint_category, endpoints.type as endpoint_type, endpoints.url as endpoint_url
		FROM sources
		INNER JOIN endpoints
		ON sources.id = endpoints.source_id;
`
