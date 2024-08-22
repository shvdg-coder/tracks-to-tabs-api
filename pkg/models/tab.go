package models

// Tab represents a tab.
type Tab struct {
	*TabEntry
	Instrument *InstrumentEntry
	Difficulty *DifficultyEntry
	References []*Reference
}
