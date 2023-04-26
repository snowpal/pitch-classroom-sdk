package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	blockPods "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/block_pods/block_pods.1"
	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	Key1Name         = "Taxes"
	AnotherKeyName   = "State Taxes"
	Block1Name       = "Form 1040"
	AnotherBlockName = "Form 1120S"
	Pod1Name         = "Income"
	BlockPod1Name    = "Expenses"
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
	log.Info("Add a new custom key")
	recipes.SleepBefore()
	newKey, err = recipes.AddCustomKey(user, Key1Name)
	if err != nil {
		return
	}
	log.Printf(".Key, %s is added successfully.", newKey.Name)
	recipes.SleepAfter()

	var (
		newBlock    response.Block
		newBlockPod response.Pod
	)
	newBlock, newBlockPod, err = addBlocksAndPods(user, newKey)
	if err != nil {
		return
	}

	log.Info("Add another key")
	recipes.SleepBefore()
	var anotherKey response.Key
	anotherKey, err = recipes.AddCustomKey(user, AnotherKeyName)
	if err != nil {
		return
	}

	log.Info("Add block")
	recipes.SleepBefore()
	var anotherBlock response.Block
	anotherBlock, err = recipes.AddBlock(user, AnotherBlockName, newKey)
	if err != nil {
		return
	}
	log.Printf(".Block, %s is created successfully.", newBlock.Name)
	recipes.SleepAfter()

	err = linkResources(user, anotherKey, anotherBlock, newBlock, newBlockPod)
	if err != nil {
		return
	}
}

func linkResources(
	user response.User,
	anotherKey response.Key,
	anotherBlock response.Block,
	newBlock response.Block,
	newBlockPod response.Pod,
) error {
	log.Info("Link block into the other key")
	recipes.SleepBefore()
	err := courses.LinkCourseToKey(user.JwtToken,
		common.ResourceIdParam{
			BlockId: newBlock.ID,
			KeyId:   anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Printf(".Block, %s is linked successfully to %s Key.", newBlock.Name, anotherKey.Name)
	recipes.SleepAfter()
	return nil
}

func addBlocksAndPods(user response.User, newKey response.Key) (response.Block, response.Pod, error) {
	var (
		pod   response.Pod
		block response.Block
	)
	log.Info("Add a new block")
	recipes.SleepBefore()
	newBlock, err := recipes.AddBlock(user, Block1Name, newKey)
	if err != nil {
		return block, pod, err
	}
	log.Printf(".Block, %s is created successfully.", newBlock.Name)
	recipes.SleepAfter()

	log.Info("Add a new block pod in this block")
	recipes.SleepBefore()
	var newBlockPod response.Pod
	newBlockPod, err = blockPods.AddBlockPod(user.JwtToken,
		request.AddPodReqBody{
			Name: BlockPod1Name,
		},
		common.ResourceIdParam{
			BlockId: newBlock.ID,
			KeyId:   newKey.ID,
		})
	if err != nil {
		return block, pod, err
	}
	log.Printf(".Block Pod, %s is created successfully in %s Block.", newBlockPod.Name, newBlock.Name)
	recipes.SleepAfter()
	return newBlock, newBlockPod, nil
}
