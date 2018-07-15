package parser

import (
	"github.com/Henrod/dummy-server/models"
)

// Parser has a parse method to return the server routes
type Parser interface {
	Parse() ([]*models.Route, error)
}
