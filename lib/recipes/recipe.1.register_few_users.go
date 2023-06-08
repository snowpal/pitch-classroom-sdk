package recipes

import (
	log "github.com/sirupsen/logrus"
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
)

func RegisterFewUsers() {
	log.Info("Objective: Sign up a few new users")
	_, err := recipes.RegisterUser(lib.DefaultEmail)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(lib.ActiveUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(lib.Teacher)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(lib.Student)
	if err != nil {
		return
	}
}
