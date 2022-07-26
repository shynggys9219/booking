package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// not the best idea to have this package

// AppConfig - a global config for an app to hold some app related info
type AppConfig struct {
	UseCache     bool // is used to get updated page if it has changes
	TemplateCach map[string]*template.Template
	InfoLog      *log.Logger // write it somewhere (terminal, file, db etc.)
	InProduction bool
	Session      *scs.SessionManager
}
