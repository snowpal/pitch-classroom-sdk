package setupnewuser

import (
	"fmt"

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
	log.Info(fmt.Sprintf("- %s | %s", email, user.JwtToken))
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

		log.Info(fmt.Sprintf("%d. %s | %s", kIndex+1, key.Name, key.Type))
		courses, err := courses.GetCourses(user.JwtToken, request.GetCoursesParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		log.Info(fmt.Sprintf("List of Courses inside %s", key.Name))
		for bIndex, course := range courses {
			log.Info(fmt.Sprintf("%d. %s", bIndex+1, course.Name))

			assessments, err := assessments.GetAssessments(user.JwtToken, request.GetAssessmentsParam{
				KeyId:    key.ID,
				CourseId: &course.ID,
			})
			if err != nil {
				return
			}

			log.Info(fmt.Sprintf("List of Assessments inside %s and %s", course.Name, key.Name))
			for bpIndex, assessment := range assessments {
				log.Info(fmt.Sprintf("%d. %s", bpIndex+1, assessment.Name))
			}
		}
	}
}

func displayAllNotifications(user response.User) {
	allNotifications, err := notifications.GetNotifications(user.JwtToken)
	if err != nil {
		return
	}

	for index, notification := range allNotifications {
		log.Info(fmt.Sprintf("%d. %s", index+1, notification.Text))
	}
}

func DisplayData(user response.User, anotherUserEmail string) {
	log.Info("## Registered Users")
	displayUser(user.Email)
	displayUser(anotherUserEmail)

	log.Info(fmt.Sprintf("## Resources Created for user: %s", user.Email))
	displayAllKeys(user)

	anotherUser, err := recipes.SignIn(anotherUserEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info(fmt.Sprintf("## Notifications for shared data as user: %s", anotherUserEmail))
	displayAllNotifications(anotherUser)
}
