package graphql

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/colachg/pallas/mysql"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	ProjectRepo mysql.ProjectRepo
}
