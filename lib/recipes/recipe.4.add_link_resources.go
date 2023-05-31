package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	Key1Name          = "Taxes"
	AnotherKeyName    = "State Taxes"
	Block1Name        = "Form 1040"
	AnotherCourseName = "Form 1120S"
	Pod1Name          = "Income"
	BlockPod1Name     = "Expenses"
)

// AddAndLinkResources Add block, pod & block pod to a key and link them into another key
func AddAndLinkResources() {
	log.Info("Objective: Add keys and blocks, and link blocks")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var newKey response.Key
	log.Info("Add a new teacher key")
	recipes.SleepBefore()
	newKey, err = recipes.AddTeacherKey(user, Key1Name)
	if err != nil {
		return
	}
	log.Printf(".Key, %s is added successfully.", newKey.Name)
	recipes.SleepAfter()

	var (
		newBlock    response.Course
		newBlockPod response.Assessment
	)
	newBlock, newBlockPod, err = addBlocksAndPods(user, newKey)
	if err != nil {
		return
	}

	log.Info("Add another key")
	recipes.SleepBefore()
	var anotherKey response.Key
	anotherKey, err = recipes.AddStudentKey(user, AnotherKeyName)
	if err != nil {
		return
	}

	log.Info("Add course")
	recipes.SleepBefore()
	var anotherBlock response.Course
	anotherBlock, err = recipes.AddCourse(user, AnotherCourseName, newKey)
	if err != nil {
		return
	}
	log.Printf(".Course, %s is created successfully.", newBlock.Name)
	recipes.SleepAfter()

	err = linkResources(user, anotherKey, anotherBlock, newBlock, newBlockPod)
	if err != nil {
		return
	}
}

func linkResources(
	user response.User,
	anotherKey response.Key,
	anotherBlock response.Course,
	newBlock response.Course,
	newBlockPod response.Assessment,
) error {
	log.Info("Link course into the other key")
	recipes.SleepBefore()
	err := courses.LinkCourseToKey(user.JwtToken,
		common.ResourceIdParam{
			CourseId: newBlock.ID,
			KeyId:    anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Printf(".Course, %s is linked successfully to %s Key.", newBlock.Name, anotherKey.Name)
	recipes.SleepAfter()
	return nil
}

func addBlocksAndPods(user response.User, newKey response.Key) (response.Course, response.Assessment, error) {
	var (
		pod   response.Assessment
		block response.Course
	)
	log.Info("Add a new block")
	recipes.SleepBefore()
	newBlock, err := recipes.AddCourse(user, Block1Name, newKey)
	if err != nil {
		return block, pod, err
	}
	log.Printf(".Course, %s is created successfully.", newBlock.Name)
	recipes.SleepAfter()

	log.Info("Add a new block pod in this block")
	recipes.SleepBefore()
	var newBlockPod response.Assessment
	newBlockPod, err = assessments.AddAssessment(user.JwtToken,
		request.AddAssessmentReqBody{
			Name: BlockPod1Name,
		},
		common.ResourceIdParam{
			CourseId: newBlock.ID,
			KeyId:    newKey.ID,
		})
	if err != nil {
		return block, pod, err
	}
	log.Printf(".Assessment, %s is created successfully in %s Course.", newBlockPod.Name, newBlock.Name)
	recipes.SleepAfter()
	return newBlock, newBlockPod, nil
}
