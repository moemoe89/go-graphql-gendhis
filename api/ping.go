//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	conf "github.com/moemoe89/go-graphql-gendhis/config"
	cons "github.com/moemoe89/go-graphql-gendhis/constant"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var starTime = time.Now()

// Ping will handle the ping endpoint
func Ping(c *gin.Context) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := starTime.In(loc)
	runMode := conf.Configuration.RunMode

	res := map[string]string{
		"start_time": startTime.Format("[02 January 2006] 15:04:05 MST"),
		"message":    "Version " + cons.APP_VERSION + " run on " + runMode + " mode",
	}
	c.JSON(http.StatusOK, res)
}
