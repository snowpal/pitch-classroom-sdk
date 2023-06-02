package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/favorites"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

const (
	FavKeyName    = "FavoriteKey"
	FavCourseName = "FavoriteCourse"
)

func AddFavorite() {
	log.Info("Objective: Add and Remove Favorites")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Create a key and a course into it. Then add that course as favorite")
	var favorite response.AddFavorite
	favorite, err = addFavorite(user)
	if err != nil {
		return
	}
	log.Info(".course added as favorite")

	err = removeFavorite(user, favorite)
	if err != nil {
		return
	}
	log.Info(".course removed from favorite")
}

func removeFavorite(user response.User, favorite response.AddFavorite) error {
	err := favorites.DeleteFavorite(user.JwtToken, favorite.ID)
	if err != nil {
		return err
	}
	return nil
}

func addFavorite(user response.User) (response.AddFavorite, error) {
	var favorite response.AddFavorite
	key, err := recipes.AddTeacherKey(user, FavKeyName)
	if err != nil {
		return favorite, err
	}
	course, err := recipes.AddCourse(user, FavCourseName, key)
	if err != nil {
		return favorite, err
	}
	favorite, err = favorites.AddCourseAsFavorite(
		user.JwtToken,
		common.ResourceIdParam{CourseId: course.ID, KeyId: key.ID})
	if err != nil {
		return favorite, err
	}
	return favorite, nil
}
