package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/favorites"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	keys "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/keys/keys.1"
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

	var user response.User
	user, err = recipes.SignIn(lib.ApiUser1, lib.Password)
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
	key, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: FavKeyName,
			Type: lib.TeacherKeyType,
		})
	if err != nil {
		return favorite, err
	}
	var course response.Course
	course, err = courses.AddCourse(
		user.JwtToken,
		request.AddCourseReqBody{Name: FavCourseName},
		key.ID)
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
