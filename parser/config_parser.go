package parser

import (
	"github.com/Henrod/dummy-server/models"
	"github.com/spf13/viper"
)

// ConfigParser parses config and returns server routes from it
type ConfigParser struct {
}

// NewConfigParser returns a config parser
func NewConfigParser() *ConfigParser {
	return &ConfigParser{}
}

// Parse parses config and returns the routes
func (c *ConfigParser) Parse() ([]*models.Route, error) {
	var routes []*models.Route

	err := viper.UnmarshalKey("paths", &routes)
	if err != nil {
		return nil, err
	}

	return routes, nil
}
