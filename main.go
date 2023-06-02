package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/snowpal/pitch-classroom-sdk/lib/recipes"
)

func main() {
	recipeID := 1
	switch recipeID {
	case 1:
		log.Info("Run Recipe1")
		recipes.RegisterFewUsers()
		break
	case 2:
		log.Info("Run Recipe2")
		recipes.GetResourceAttributes()
		break
	case 3:
		log.Info("Run Recipe3")
		recipes.AddAndLinkResources()
		break
	case 4:
		log.Info("Run Recipe4")
		recipes.ShareCourse()
		break
	case 5:
		log.Info("Run Recipe5")
		recipes.GetAllKeys()
		break
	case 6:
		log.Info("Run Recipe6")
		recipes.AddFavorite()
		break
	case 7:
		log.Info("Run Recipe7")
		recipes.FetchScheduler()
		break
	case 8:
		log.Info("Run Recipe8")
		recipes.AddRelation()
		break
	case 9:
		log.Info("Run Recipe9")
		recipes.PublishStudentGrade()
		break
	case 10:
		log.Info("Run Recipe10")
		recipes.GrantAclOnCustomCourse()
		break
	case 11:
		log.Info("Run Recipe11")
		recipes.UpdateAttributes()
		break
	default:
		log.Info("pick a specific recipe to run")
	}
}
