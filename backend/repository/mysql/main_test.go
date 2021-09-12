package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/Thewalkers2012/DOJ/settings"
)

func TestMain(m *testing.M) {
	if err := settings.Init(); err != nil {
		log.Fatal("cannot load config file")
	}

	if err := Init(settings.Config.MySQLConfig); err != nil {
		log.Fatal("cannot init mysql")
	}

	os.Exit(m.Run())
}
