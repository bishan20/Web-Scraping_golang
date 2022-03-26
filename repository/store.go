package repository

import (
	"gorm.io/gorm"
)

type Store interface {
	//Querier
}

type SQLStore struct {
	db *gorm.DB
	// *Queries
}

//NewStore creates a new store
func NewStore(db *gorm.DB) Store {
	return &SQLStore{
		db: db,
		//Queries: New(db),
	}
}
