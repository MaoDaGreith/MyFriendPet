package repositories

import (
	"github.com/upper/db/v4"
	// "github.com/upper/db/v4/lib/sqlbuilder"
	"github.com/upper/db/v4/adapter/mysql"

	"github.com/MaoDaGreith/MyFriendPet/config"
)

func GetDBConnection(dbConnSettings config.DBConnSettings) (db.Session, error) {
	options := map[string]string{}
	if dbConnSettings.EnableMultiStatements {
		options["multiStatements"] = "true"
	}
	connection := mysql.ConnectionURL{
		User:     dbConnSettings.User,
		Password: dbConnSettings.Password,
		Database: dbConnSettings.Database,
		Host:     dbConnSettings.Host,
		Options:  options,
	}

	session, err := mysql.Open(connection)
	if err != nil {
		return nil, err
	}

	return session, nil
}
