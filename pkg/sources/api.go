package sources

import (
	"database/sql"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing sources.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateSourcesTable creates a sources table if it doesn't already exist.
func (a *API) CreateSourcesTable() {
	_, err := a.Database.DB.Exec(createSourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sources' table.")
	}
}

// DropSourcesTable drops the sources table if it exists.
func (a *API) DropSourcesTable() {
	_, err := a.Database.DB.Exec(dropSourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sources' table.")
	}
}

// InsertSources inserts multiple sources in the sources table.
func (a *API) InsertSources(sources ...*Source) {
	for _, source := range sources {
		a.InsertSource(source)
	}
}

// InsertSource inserts a new source in the sources table.
func (a *API) InsertSource(source *Source) {
	_, err := a.Database.DB.Exec(insertSourceQuery, source.ID, source.Name, source.Category)
	if err != nil {
		log.Printf("Failed inserting source with name: '%s': %s", source.Name, err.Error())
	} else {
		log.Printf("Successfully inserted source with name: '%s'", source.Name)
	}
}

// GetSources retrieves the sources.
func (a *API) GetSources() []*Source {
	rows, err := a.Database.DB.Query(getSourcesQuery)
	if err != nil {
		log.Printf("Failed to get sources: %s", err)
		return nil
	}

	sources := rowsToSources(rows)
	defer closeRows(rows)

	return sources
}

// rowsToSources converts the given *sql.Rows into a slice of *Source objects.
func rowsToSources(rows *sql.Rows) []*Source {
	var sources []*Source
	for rows.Next() {
		source := rowsToSource(rows)
		if source != nil {
			sources = append(sources, source)
		}
	}
	return sources
}

// rowsToSource scans the SQL row into a Source struct.
func rowsToSource(rows *sql.Rows) *Source {
	var source Source
	err := rows.Scan(&source.ID, &source.Name)
	if err != nil {
		log.Printf("Unable to scan source: %s", err.Error())
		return nil
	}
	return &source
}

// closeRows closes the SQL rows and logs error if any.
func closeRows(rows *sql.Rows) {
	err := rows.Err()
	if err != nil {
		log.Printf("Error while processing rows: %s", err.Error())
	}
	err = rows.Close()
	if err != nil {
		log.Printf("Failed to close rows: %s", err.Error())
	}
}

// CreateSourcesToEndpointsView creates the sources to endpoints view.
func (a *API) CreateSourcesToEndpointsView() {
	_, err := a.Database.DB.Exec(createSourcesToEndpointsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sources' to 'endpoints' view")
	}
}

// DropSourcesToEndpointsView drops the sources to endpoints view if it exists.
func (a *API) DropSourcesToEndpointsView() {
	_, err := a.Database.DB.Exec(dropSourcesToEndpointsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sources' to 'endpoints' view")
	}
}
