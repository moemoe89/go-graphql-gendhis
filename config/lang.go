//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/moemoe89/go-localization"
)

// InitLang will create a variable that represent the language.Config
func InitLang() (*language.Config, error) {

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	cfg := language.New()
	basepath = strings.Replace(basepath, "config", "", -1)
	cfg.BindPath(basepath + "languages/lang.json")
	cfg.BindMainLocale("en")

	lang, err := cfg.Init()
	if err != nil {
		return nil, fmt.Errorf("Failed init localization: %s", err.Error())
	}

	return lang, nil
}
