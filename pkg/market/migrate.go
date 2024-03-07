package market

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"strings"
)

func Migrate(dataSourceName string) (err error) {
	if !strings.Contains(dataSourceName, "://") {
		err = errors.New("market: undefined data source name " + dataSourceName)
		return
	}
	driverName := strings.ToLower(strings.Split(dataSourceName, "://")[0])

	migrations, err := migrate.New(fmt.Sprintf("file://migrations/%s", driverName), dataSourceName)
	if err != nil {
		return
	}

	if err = migrations.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
	}

	return
}
