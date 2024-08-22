package models

import (
	"github.com/google/uuid"
)

// TabEntry represents a tab.
type TabEntry struct {
	ID           uuid.UUID
	InstrumentID uint
	DifficultyID uint
	Description  string
}
