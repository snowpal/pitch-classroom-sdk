package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/attributes"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"

	log "github.com/sirupsen/logrus"
)

// GetResourceAttributes sign in, get resource attributes
func GetResourceAttributes() {
	log.Info("Objective: Get resource attributes")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	log.Info("Sign in user, email: ", lib.ApiUser1)
	user, err := recipes.SignIn(lib.ApiUser1, lib.Password)

	log.Info(".get resource attributes")
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
