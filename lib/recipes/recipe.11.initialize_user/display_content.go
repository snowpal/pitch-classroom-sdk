package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/notifications"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func displayUser(email string) {
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return
	}
	log.Info("- ", email, " | ", user.JwtToken)
}

func displayAllKeys(user response.User) {
	keys, err := keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		return
	}
	log.Info("List of Keys")
	for kIndex, key := range keys {
		if key.Type != lib.StudentKeyType && key.Type != lib.TeacherKeyType {
			continue
		}

		log.Info(kIndex+1, ". ", key.Name, " | ", key.Type)
		courses, err := courses.GetCourses(user.JwtToken, request.GetCoursesParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		log.Info("List of Courses inside ", key.Name)
		for bIndex, course := range courses {
			log.Info(bIndex+1, ". ", course.Name)

			assessments, err := assessments.GetAssessments(user.JwtToken, request.GetAssessmentsParam{
				KeyId:    key.ID,
				CourseId: &course.ID,
			})
			if err != nil {
				return
			}

			log.Info("List of Assessments inside ", course.Name, " and ", key.Name)
			for bpIndex, assessment := range assessments {
				log.Info(bpIndex+1, ". ", assessment.Name)
			}
		}
	}
}

func displayAllNotifications(user response.User) {
	notifications, err := notifications.GetNotifications(user.JwtToken)
	if err != nil {
		return
	}

	for index, notification := range notifications {
		log.Info(index+1, ". ", notification.Text)
	}
}

func DisplayContent(user response.User, anotherUserEmail string) {
	log.Info("## Registered Users")
	displayUser(user.Email)
	displayUser(anotherUserEmail)

	log.Info("## Resources Created for user: ", user.Email)
	displayAllKeys(user)

	anotherUser, err := recipes.SignIn(anotherUserEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info("## Notifications for shared content as user: ", anotherUserEmail)
	displayAllNotifications(anotherUser)
}
