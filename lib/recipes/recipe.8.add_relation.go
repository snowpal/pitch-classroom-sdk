package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/relations"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	keys "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/keys/keys.1"
	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

const (
	RelationKeyName        = "Animals"
	RelationCourseName     = "Tiger"
	RelationAssessmentName = "Cat"
)

func AddRelation() {
	log.Info("Objective: Add and Remove Relations")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	log.Info("Create a key and course & assessment into that key")
	user, err := recipes.SignIn(lib.ApiUser1, lib.Password)
	if err != nil {
		return
	}

	log.Info("Relate the course with key")
	key, course, err := addRelation(user)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Course %s is related with key %s successfully", course.Name, key.Name))

	log.Info("Unrelate the course from key")
	err = removeRelation(user, key, course)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Course %s is unrelated from key %s successfully", course.Name, key.Name))
}

func removeRelation(user response.User, key response.Key, course response.Course) error {
	err := relations.UnrelateCourseFromKey(
		user.JwtToken,
		request.KeyToCourseRelationParam{
			KeyId:          key.ID,
			TargetCourseId: course.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func addRelation(user response.User) (response.Key, response.Course, error) {
	var (
		key    response.Key
		course response.Course
		err    error
	)
	key, err = keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: RelationKeyName,
			Type: lib.TeacherKeyType,
		})
	if err != nil {
		return key, course, err
	}
	course, err = courses.AddCourse(
		user.JwtToken,
		request.AddCourseReqBody{Name: RelationCourseName},
		key.ID)
	if err != nil {
		return key, course, err
	}
	err = relations.RelateCourseToKey(
		user.JwtToken,
		request.KeyToCourseRelationParam{
			KeyId:          key.ID,
			TargetCourseId: course.ID,
		},
	)
	if err != nil {
		return key, course, err
	}
	return key, course, nil

}
