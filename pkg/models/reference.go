package models

// Reference represents a record in the 'references' table.
type Reference struct {
	*ReferenceEntry
	Source *Source
}
