package views

import (
	"database/sql"
	logic "github.com/shvdg-dev/base-logic/pkg"
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"log"
)

// API is for managing 'sources, endpoints' views.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateSourcesEndpointsView creates the view with the sources and endpoints.
func (a *API) CreateSourcesEndpointsView() {
	_, err := a.Database.DB.Exec(createSourcesEndpointsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'src' to 'end' view")
	}
}

// DropSourcesEndpointsView drops the view with the sources and endpoints.
func (a *API) DropSourcesEndpointsView() {
	_, err := a.Database.DB.Exec(dropSourcesEndpointsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'src' to 'end' view")
	}
}

// GetAllSourcesWithEndpoints gets all the sources and endpoints from the database.
func (a *API) GetAllSourcesWithEndpoints() ([]*src.Source, error) {
	rows, err := a.Database.DB.Query(selectSourcesEndpoints)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var sources []*src.Source

	for rows.Next() {
		var srcID uint
		var srcName, srcCategory, endCategory, endType, endUrl string
		err := rows.Scan(&srcID, &srcName, &srcCategory, &endCategory, &endType, &endUrl)

		if err != nil {
			return nil, err
		}

		endpoint := end.NewEndpoint(srcID, endCategory, endType, endUrl)
		existingSource := findSourceByID(sources, srcID)
		if existingSource != nil {
			existingSource.Endpoints = append(existingSource.Endpoints, endpoint)
		} else {
			source := src.NewSource(srcID, srcName, src.WithCategory(srcCategory), src.WithEndpoint(endpoint))
			sources = append(sources, source)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sources, nil
}

// findSourceByID attempts to return a source with the provided ID, if not present, return nil
func findSourceByID(sources []*src.Source, srcID uint) *src.Source {
	for _, source := range sources {
		if source.ID == srcID {
			return source
		}
	}
	return nil
}
