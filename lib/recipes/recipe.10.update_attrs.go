package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/attributes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
)

const (
	AttrsKeyName    = "Birds1"
	AttrsCourseName = "Parrot1"
)

// UpdateAttributes sign in, update key attributes, update course attributes, update assessment attributes, update course assessment attributes,
// get resource attributes
func UpdateAttributes() {
	log.Info("Objective: Update show/hide of key, course, assessment & course assessment attributes")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ApiUser1, lib.Password)
	if err != nil {
		return
	}

	log.Info("Get displayable attributes")
	recipes.SleepBefore()
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}
	log.Info(resourceAttrs)

	log.Info("Update key attributes")
	recipes.SleepBefore()
	key, err := recipes.AddTeacherKey(user, AttrsKeyName)
	if err != nil {
		return
	}
	err = attributes.UpdateKeyAttrs(
		user.JwtToken,
		key.ID,
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Printf(".Attributes for Key %s updated successfully", key.Name)
	recipes.SleepAfter()

	log.Info("Update course attributes")
	recipes.SleepBefore()
	course, err := recipes.AddCourse(user, AttrsCourseName, key)
	if err != nil {
		return
	}
	err = attributes.UpdateCourseDisplayAttributes(
		user.JwtToken,
		common.ResourceIdParam{
			CourseId: course.ID,
			KeyId:    course.Key.ID,
		},
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Printf(".Attributes for course %s updated successfully", key.Name)
	recipes.SleepAfter()
}
