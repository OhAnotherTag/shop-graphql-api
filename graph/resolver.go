package graph

import (
	"github.com/OhAnotherTag/shop-gql-api/config/database"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{

}

func (r *Resolver) Database() *gorm.DB {
	return database.DB
}
