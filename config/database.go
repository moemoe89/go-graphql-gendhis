//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// InitDB will create a variable that represent the sqlx.DB
func InitDB() (*sqlx.DB, *sqlx.DB, error) {
	DBRead, err := sqlx.Connect(Configuration.DialectSlave, Configuration.DsnSlave)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open connection to slave database: %s", err.Error())
	}

	err = DBRead.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to ping connection to slave database: %s", err.Error())
	}

	DBRead.SetMaxIdleConns(Configuration.IdleConnSlave)
	DBRead.SetMaxOpenConns(Configuration.MaxConnSlave)

	DBWrite, err := sqlx.Connect(Configuration.DialectMaster, Configuration.DsnMaster)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open connection to master database: %s", err.Error())
	}

	err = DBWrite.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to ping connection to master database: %s", err.Error())
	}

	return DBRead, DBWrite, nil
}
