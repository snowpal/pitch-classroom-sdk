package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ShareContent(user response.User, anotherUserEmail string, keyWithCourses KeyWithCourses) error {
	log.Info("Share key content with ", anotherUserEmail)

	log.Info("Share course ", keyWithCourses.Courses[0].Course.Name, " for teacher key, ", keyWithCourses.Key.Name)
	err := recipes.SearchUserAndShareCourse(user, keyWithCourses.Courses[0].Course, anotherUserEmail, lib.StudentAcl)
	if err != nil {
		return err
	}

	return nil
}
