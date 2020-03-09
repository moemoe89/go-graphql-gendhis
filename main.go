//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	"github.com/moemoe89/practicing-graphql-golang/api/v1/user"
	conf "github.com/moemoe89/practicing-graphql-golang/config"
	_ "github.com/moemoe89/practicing-graphql-golang/docs"
	"github.com/moemoe89/practicing-graphql-golang/routers"

	"fmt"

	"github.com/DeanThompson/ginpprof"
)

// @title Simple REST API
// @version 1.0.0
// @description This is a documentation of Simple REST API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url -
// @contact.email bismobaruno@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /api/v1
func main() {
	dbR, dbW, err := conf.InitDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbR.Close()
		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		err := dbW.Close()
		if err != nil {
			panic(err)
		}
	}()

	lang, err := conf.InitLang()
	if err != nil {
		panic(err)
	}

	log := conf.InitLog()

	userRepo := user.NewPostgresRepository(dbR, dbW)
	userSvc := user.NewService(log, userRepo)

	app := routers.GetRouter(lang, log, userSvc)
	ginpprof.Wrap(app)
	err = app.Run(":" + conf.Configuration.Port)
	if err != nil {
		panic(fmt.Sprintf("Can't start the app: %s", err.Error()))
	}
}
