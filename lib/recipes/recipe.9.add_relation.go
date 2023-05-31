package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/relations"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

const (
	RelationKeyName   = "Animals"
	RelationBlockName = "Tiger"
	RelationPodName   = "Cat"
)

func AddRelation() {
	log.Info("Objective: Add and Remove Relations")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	log.Info("Create a key and block & pod into that key")
	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Relate the block with key pod")
	key, block, err := addRelation(user)
	if err != nil {
		return
	}
	log.Printf(".Course %s is related with pod %s successfully", block.Name, key.Name)

	log.Info("Unrelate the block from key pod")
	err = removeRelation(user, key, block)
	if err != nil {
		return
	}
	log.Printf(".Course %s is unrelated from pod %s successfully", block.Name, key.Name)
}

func removeRelation(user response.User, key response.Key, block response.Course) error {
	err := relations.UnrelateCourseFromKey(
		user.JwtToken,
		request.KeyToCourseRelationParam{
			KeyId:         key.ID,
			TargetBlockId: block.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func addRelation(user response.User) (response.Key, response.Course, error) {
	var (
		key   response.Key
		block response.Course
	)
	key, err := recipes.AddTeacherKey(user, RelationKeyName)
	if err != nil {
		return key, block, err
	}
	block, err = recipes.AddCourse(user, RelationBlockName, key)
	if err != nil {
		return key, block, err
	}
	err = relations.RelateCourseToKey(
		user.JwtToken,
		request.KeyToCourseRelationParam{
			KeyId:         key.ID,
			TargetBlockId: block.ID,
		},
	)
	if err != nil {
		return key, block, err
	}
	return key, block, nil

}
