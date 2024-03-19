package market

import (
	_ "database/sql"
	"errors"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

// postgres://username:password@localhost:5432/dbname?sslmode=disable&search_path=public
// oracle://username:password@:0/?connstr=(description=(address=(protocol=tcp)(host=localhost)(port=1521))(connect_data=(server=dedicated)(sid=dbname)))&persist security info=true&ssl=enable&ssl verify=false
type SQLX struct {
	Client *sqlx.DB
}

func NewSQL(dataSourceName string) (market SQLX, err error) {
	if !strings.Contains(dataSourceName, "://") {
		err = errors.New("market: indefined data source name")
		return
	}
	driverName := strings.ToLower(strings.Split(dataSourceName, "://")[0])

	market.Client, err = sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return
	}
	market.Client.SetMaxOpenConns(20)

	return
}
