package database

import "github.com/google/uuid"

type SchemaHistory struct {
	Id       uuid.UUID
	Type     string
	Script   string
	Checksum int64
}

type Migration struct {
	Type        string
	Version     string
	Description string
	Script      string
	Checksum    int64
	Success     bool
}
