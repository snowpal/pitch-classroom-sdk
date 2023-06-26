package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ShareData(user response.User, anotherUserEmail string, keyWithCourses KeyWithCourses) error {
	log.Info(fmt.Sprintf("Share course with %s", anotherUserEmail))

	course := keyWithCourses.Courses[0].Course
	log.Info(fmt.Sprintf("Share course %s for teacher key, %s", course.Name, keyWithCourses.Key.Name))
	err := recipes.SearchUserAndShareCourse(user, course, anotherUserEmail, lib.StudentAcl)
	if err != nil {
		return err
	}

	return nil
}
