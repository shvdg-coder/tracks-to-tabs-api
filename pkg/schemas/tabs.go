package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/queries"
	"log"
)

// TabSchema represents schema operations related to 'tabs' in the database.
type TabSchema interface {
	CreateTabsTable()
	DropTabsTable()
}

// TabSvc is for managing 'tabs' tables in the database.
type TabSvc struct {
	logic.DbOperations
}

// NewTabSvc creates a new instance of the TabSvc struct.
func NewTabSvc(database logic.DbOperations) TabSchema {
	return &TabSvc{database}
}

// CreateTabsTable creates a tabs table if it doesn't already exist.
func (s *TabSvc) CreateTabsTable() {
	_, err := s.Exec(queries.CreateTabsTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropTabsTable drops the tabs table if it exists.
func (s *TabSvc) DropTabsTable() {
	_, err := s.Exec(queries.DropTabsTable)
	if err != nil {
		log.Fatal(err)
	}
}
