package lib


const GatewayHost = "https://gateway.snowpal.com/"

const (
	RouteAttributesGetDisplayableAttributes          = "app/resource/attributes"
	RouteAttributesUpdateKeyDisplayAttributes        = "keys/%s/attributes"
	RouteAttributesUpdateCourseDisplayAttributes     = "courses/%s/attributes?keyId=%s"
	RouteAttributesUpdateAssessmentDisplayAttributes = "assessments/%s/attributes?keyId=%s&courseId=%s"
)

const (
	RouteCoursesGetCourses                             = "keys/%s/courses?filter=%s&batchIndex=%s&aclWriteOrHigher=%s"
	RouteCoursesAddCourse                              = "keys/%s/courses"
	RouteCoursesGetCoursesLinkedToAssessment           = "assessments/:id/linked-to/courses?keyId=%s"
	RouteCoursesAddCourseBasedOnTemplate               = "keys/%s/courses/by-template?templateId=%s&excludeAssessments=%s&excludeTasks=%s"
	RouteCoursesLinkCourseToKey                        = "keys/%s/courses/%s/link"
	RouteCoursesUnlinkCourseFromKey                    = "keys/%s/courses/%s/unlink"
	RouteCoursesGetCoursesAvailableToBeLinkedToThisKey = "keys/%s/courses/available-to-link"
	RouteCoursesGetCourse                              = "courses/%s?keyId=%s"
	RouteCoursesUpdateCourse                           = "courses/%s?keyId=%s"
	RouteCoursesAddCourseTypeToCourse                  = "courses/%s/course-types/%s?keyId=%s"
	RouteCoursesDeleteCourseTypeFromCourse             = "courses/%s/course-types?keyId=%s"
	RouteCoursesAddScaleToCourse                       = "courses/%s/scales/%s?keyId=%s"
	RouteCoursesDeleteScaleFromCourse                  = "courses/%s/scales?keyId=%s"
	RouteCoursesUpdateCourseScaleValue                 = "courses/%s/scale-value?keyId=%s"
	RouteCoursesUpdateCourseDescription                = "courses/%s/description?keyId=%s"
	RouteCoursesArchiveCourse                          = "courses/%s/archive?keyId=%s"
	RouteCoursesUnarchiveCourse                        = "courses/%s/unarchive?keyId=%s"
	RouteCoursesGetArchivedCourses                     = "courses/archived?keyId=%s&batchIndex=%s"
	RouteCoursesBulkArchiveCourses                     = "courses/archive?keyId=%s"
	RouteCoursesAllowArchivalOfCourse                  = "courses/%s/allow-archival?keyId=%s"
	RouteCoursesCopyCourse                             = "courses/%s/copy?keyId=%s&allTasks=%s&assessmentIds=%s&allAssessments=%s&allChecklists=%s&targetKeyId=%s"
	RouteCoursesMoveCourse                             = "courses/%s/move?keyId=%s&targetKeyId=%s"
)

const (
	RouteCoursesGetCourseAttachments   = "courses/%s/attachments?keyId=%s"
	RouteCoursesAddCourseAttachment    = "courses/%s/attachments?keyId=%s"
	RouteCoursesRenameCourseAttachment = "course-attachments/%s/rename?keyId=%s&courseId=%s"
	RouteCoursesDeleteCourseAttachment = "course-attachments/%s?keyId=%s&courseId=%s"
)

const (
	RouteCoursesGetLinkedAssessments       = "charts/keys/%s/courses/%s/linked-resources"
	RouteCoursesGetScaleValuesForScale     = "charts/keys/%s/courses/%s/scales/%s/grades"
	RouteCoursesGetTaskStatusForCourse     = "charts/keys/%s/courses/%s/task-status"
	RouteCoursesGetCourseGradesForStudents = "courses/%s/students/all/grades?keyId=%s"
)

const (
	RouteCoursesGetCourseChecklists         = "courses/%s/checklists?keyId=%s"
	RouteCoursesAddCourseChecklist          = "courses/%s/checklists?keyId=%s"
	RouteCoursesReorderCourseChecklists     = "courses/%s/checklists/reorder?keyId=%s"
	RouteCoursesDeleteCourseChecklist       = "courses/%s/checklists/%s?keyId=%s"
	RouteCoursesRenameCourseChecklist       = "courses/%s/checklists/%s?keyId=%s"
	RouteCoursesAddCourseChecklistItem      = "courses/%s/checklists/%s/checklist-items?keyId=%s"
	RouteCoursesUpdateCourseChecklistItem   = "courses/%s/checklists/%s/checklist-items/%s?keyId=%s"
	RouteCoursesDeleteCourseChecklistItem   = "courses/%s/checklists/%s/checklist-items/%s?keyId=%s"
	RouteCoursesReorderCourseChecklistItems = "courses/%s/checklists/%s/checklist-items/reorder?keyId=%s"
)

const (
	RouteCoursesGetCourseComments   = "courses/%s/comments?keyId=%s"
	RouteCoursesAddCourseComment    = "courses/%s/comments?keyId=%s"
	RouteCoursesUpdateCourseComment = "course-comments/%s?keyId=%s&courseId=%s"
	RouteCoursesDeleteCourseComment = "course-comments/%s?keyId=%s&courseId=%s"
)

const (
	RouteCoursesGetCourseNotes   = "courses/%s/notes?keyId=%s"
	RouteCoursesAddCourseNote    = "courses/%s/notes?keyId=%s"
	RouteCoursesUpdateCourseNote = "course-notes/%s?keyId=%s&courseId=%s"
	RouteCoursesDeleteCourseNote = "course-notes/%s?keyId=%s&courseId=%s"
)

const (
	RouteCoursesGetCourseTasks     = "courses/%s/tasks?keyId=%s"
	RouteCoursesAddCourseTask      = "courses/%s/tasks?keyId=%s"
	RouteCoursesUpdateCourseTask   = "course-tasks/%s?keyId=%s&courseId=%s"
	RouteCoursesDeleteCourseTask   = "course-tasks/%s?keyId=%s&courseId=%s"
	RouteCoursesAssignCourseTask   = "course-tasks/%s/assign?keyId=%s&courseId=%s"
	RouteCoursesUnassignCourseTask = "course-tasks/%s/unassign?keyId=%s&courseId=%s"
	RouteCoursesReorderCourseTasks = "courses/%s/tasks/reorder?keyId=%s"
)

const (
	RouteAssessmentsGetAssessments                                = "courses/%s/assessments?batchIndex=%s&keyId=%s"
	RouteAssessmentsAddAssessment                                 = "courses/%s/assessments?keyId=%s"
	RouteAssessmentsAddAssessmentBasedOnTemplate                  = "courses/%s/assessments/by-template?keyId=%s&templateId=%s&excludeTasks=%s"
	RouteAssessmentsLinkAssessmentToCourse                        = "courses/%s/assessments/%s/link?keyId=%s"
	RouteAssessmentsUnlinkAssessmentFromCourse                    = "courses/%s/assessments/%s/unlink?keyId=%s"
	RouteAssessmentsGetAssessment                                 = "assessments/%s?keyId=%s&courseId=%s"
	RouteAssessmentsUpdateAssessment                              = "assessments/%s?keyId=%s&courseId=%s"
	RouteAssessmentsUpdateAssessmentCompletionStatus              = "assessments/%s/by-completion-status?keyId=%s&courseId=%s"
	RouteAssessmentsAddAssessmentTypeToAssessment                 = "assessments/%s/assessment-types/%s?keyId=%s&courseId=%s"
	RouteAssessmentsDeleteAssessmentTypeFromAssessment            = "assessments/%s/assessment-types?keyId=%s&courseId=%s"
	RouteAssessmentsAddScaleToAssessment                          = "assessments/%s/scales/%s?keyId=%s&courseId=%s"
	RouteAssessmentsDeleteScaleFromAssessment                     = "assessments/%s/scales?keyId=%s&courseId=%s"
	RouteAssessmentsUpdateAssessmentScaleValue                    = "assessments/%s/scale-value?keyId=%s&courseId=%s"
	RouteAssessmentsArchiveAssessment                             = "assessments/%s/archive?keyId=%s&courseId=%s"
	RouteAssessmentsGetArchivedAssessments                        = "assessments/archived?batchIndex=%s&keyId=%s&courseId=%s"
	RouteAssessmentsGetAssessmentsAvailableToBeLinkedToThisCourse = "courses/%s/assessments/available-to-link?keyId=%s"
	RouteAssessmentsUnarchiveAssessment                           = "assessments/%s/unarchive?keyId=%s&courseId=%s"
	RouteAssessmentsBulkArchiveAssessments                        = "assessments/archive?keyId=%s&courseId=%s"
	RouteAssessmentsUpdateAssessmentDescription                   = "assessments/%s/description?keyId=%s&courseId=%s"
	RouteAssessmentsAllowArchivalOfAssessment                     = "assessments/%s/allow-archival?keyId=%s&courseId=%s"
	RouteAssessmentsCopyAssessment                                = "assessments/%s/copy?keyId=%s&courseId=%s&allTasks=%s&allChecklists=%s&targetKeyId=%s&targetCourseId=%s"
	RouteAssessmentsMoveAssessment                                = "assessments/%s/move?keyId=%s&courseId=%s&targetKeyId=%s&targetCourseId=%s"
)

const (
	RouteAssessmentsGetAssessmentAttachments   = "assessments/%s/attachments?keyId=%s&courseId=%s"
	RouteAssessmentsAddAssessmentAttachment    = "assessments/%s/attachments?keyId=%s&courseId=%s"
	RouteAssessmentsRenameAssessmentAttachment = "assessment-attachments/%s/rename?keyId=%s&courseId=%s&assessmentId=%s"
	RouteAssessmentsDeleteAssessmentAttachment = "assessment-attachments/%s?keyId=%s&courseId=%s&assessmentId=%s"
)

const (
	RouteAssessmentsGetAssessmentTasksForCharts    = "charts/assessments/%s/tasks?keyId=%s&courseId=%s"
	RouteAssessmentsGetAssessmentGradesForStudents = "charts/assessments/%s/students/grades?keyId=%s&courseId=%s"
)

const (
	RouteAssessmentsGetAssessmentChecklists         = "assessments/%s/checklists?keyId=%s&courseId=%s"
	RouteAssessmentsAddAssessmentChecklist          = "assessments/%s/checklists?keyId=%s&courseId=%s"
	RouteAssessmentsReorderAssessmentChecklists     = "assessments/%s/checklists/reorder?keyId=%s&courseId=%s"
	RouteAssessmentsDeleteAssessmentChecklist       = "assessments/%s/checklists/%s?keyId=%s&courseId=%s"
	RouteAssessmentsRenameAssessmentChecklist       = "assessments/%s/checklists/%s?keyId=%s&courseId=%s"
	RouteAssessmentsAddAssessmentChecklistItem      = "assessments/%s/checklists/%s/checklist-items?keyId=%s&courseId=%s"
	RouteAssessmentsUpdateAssessmentChecklistItem   = "assessments/%s/checklists/%s/checklist-items/%s?keyId=%s&courseId=%s"
	RouteAssessmentsDeleteAssessmentChecklistItem   = "assessments/%s/checklists/%s/checklist-items/%s?keyId=%s&courseId=%s"
	RouteAssessmentsReorderAssessmentChecklistItems = "assessments/%s/checklists/%s/checklist-items/reorder?keyId=%s&courseId=%s"
)

const (
	RouteAssessmentsGetAssessmentComments   = "assessments/%s/comments?keyId=%s&courseId=%s"
	RouteAssessmentsAddAssessmentComment    = "assessments/%s/comments?keyId=%s&courseId=%s"
	RouteAssessmentsUpdateAssessmentComment = "assessment-comments/%s?keyId=%s&courseId=%s&assessmentId=%s"
	RouteAssessmentsDeleteAssessmentComment = "assessment-comments/%s?keyId=%s&courseId=%s&assessmentId=%s"
)

const (
	RouteAssessmentsGetAssessmentNotes   = "assessments/%s/notes?keyId=%s&courseId=%s"
	RouteAssessmentsAddAssessmentNote    = "assessments/%s/notes?keyId=%s&courseId=%s"
	RouteAssessmentsUpdateAssessmentNote = "assessment-notes/%s?keyId=%s&courseId=%s&assessmentId=%s"
	RouteAssessmentsDeleteAssessmentNote = "assessment-notes/%s?keyId=%s&courseId=%s&assessmentId=%s"
)

const (
	RouteAssessmentsGetAssessmentTasks     = "assessments/%s/tasks?keyId=%s&courseId=%s"
	RouteAssessmentsAddAssessmentTask      = "assessments/%s/tasks?keyId=%s&courseId=%s"
	RouteAssessmentsUpdateAssessmentTask   = "assessment-tasks/%s?keyId=%s&courseId=%s&assessmentId=%s"
	RouteAssessmentsDeleteAssessmentTask   = "assessment-tasks/%s?keyId=%s&courseId=%s&assessmentId=%s"
	RouteAssessmentsAssignAssessmentTask   = "assessment-tasks/%s/assign?keyId=%s&courseId=%s&assessmentId=%s"
	RouteAssessmentsUnassignAssessmentTask = "assessment-tasks/%s/unassign?keyId=%s&courseId=%s&assessmentId=%s"
	RouteAssessmentsReorderAssessmentTasks = "assessments/%s/tasks/reorder?keyId=%s&courseId=%s"
)

const (
	RouteCourseTypesGetCourseTypes            = "course-types?includeCounts=%s"
	RouteCourseTypesAddCourseType             = "course-types"
	RouteCourseTypesUpdateCourseType          = "course-types/%s"
	RouteCourseTypesDeleteCourseType          = "course-types/%s"
	RouteCourseTypesGetCoursesUsingCourseType = "course-types/%s/courses"
)

const (
	RouteCollaborationGetCourseCollaborators            = "courses/%s/acl?keyId=%s"
	RouteCollaborationUpdateCourseAcl                   = "courses/%s/users/%s/acl?keyId=%s"
	RouteCollaborationUnshareCourseFromCollaborator     = "courses/%s/users/%s/unshare?keyId=%s"
	RouteCollaborationShareCourseWithCollaborator       = "courses/%s/users/%s/share?keyId=%s"
	RouteCollaborationGetUsersThisCourseCanBeSharedWith = "search/courses/%s/shareable/users?keyId=%s&token=%s"
	RouteCollaborationBulkShareCoursesWithCollaborators = "courses/users/%s/share?keyId=%s"
	RouteCollaborationLeaveCourse                       = "courses/%s/leave?keyId=%s"
)

const (
	RouteCommentsGetRecentComments = "comments"
)

const (
	RouteConversationsGetUnreadConversationsCount         = "conversations/unread-status"
	RouteConversationsGetUserConversations                = "conversations"
	RouteConversationsAddPrivateOrGroupConversation       = "conversations"
	RouteConversationsGetConversationForGivenUsernames    = "conversations/by-usernames?userNames=%s"
	RouteConversationsSendMessageToAnExistingConversation = "conversations/%s/messages"
	RouteConversationsGetConversation                     = "conversations/%s"
	RouteConversationsDeleteConversation                  = "conversations/%s"
	RouteConversationsLeaveConversation                   = "conversations/%s/leave"
	RouteConversationsArchiveConversation                 = "conversations/%s/archive"
)

const (
	RouteDashboardGetDashboardDetails                      = "dashboard/combined-responses"
	RouteDashboardGetRecentlyModifiedCoursesAndAssessments = "dashboard/recently-modified"
	RouteDashboardGetUnreadCount                           = "dashboard/unread-count"
	RouteDashboardGetRecentlyModifiedKeys                  = "dashboard/recently-modified/keys"
	RouteDashboardGetAssessmentsAndTasksDueShortly         = "dashboard/due-shortly/assessments-and-tasks"
	RouteDashboardGetCoursesDueShortly                     = "dashboard/due-shortly/courses"
	RouteDashboardGetUnreadNotifications                   = "dashboard/notifications/unread-status"
)

const (
	RouteDashboardGetUserKeysCoursesAndAssessments           = "charts/dashboard/keys-courses-assessments"
	RouteDashboardGetSystemKeysCoursesAndAssessments         = "charts/dashboard/system-keys"
	RouteDashboardGetFilteredUserKeysCoursesAndAssessments   = "charts/dashboard/keys/filters"
	RouteDashboardGetFilteredSystemKeysCoursesAndAssessments = "charts/dashboard/system-keys/filters"
	RouteDashboardGetCoursesBasedOnCourseTypes               = "charts/dashboard/course-types"
	RouteDashboardGetAssessmentsBasedOnAssessmentTypes       = "charts/dashboard/assessment-types"
	RouteDashboardGetCoursesAndAssessmentsBasedOnScales      = "charts/dashboard/scales"
	RouteDashboardGetTasksByStatus                           = "charts/dashboard/task-status"
)

const (
	RouteFavoritesGetFavorites            = "favorites"
	RouteFavoritesAddKeyAsFavorite        = "favorites/keys/%s"
	RouteFavoritesAddCourseAsFavorite     = "favorites/courses/%s?keyId=%s"
	RouteFavoritesAddAssessmentAsFavorite = "favorites/assessments/%s?keyId=%s&courseId=%s"
	RouteFavoritesDeleteFavorite          = "favorites/%s"
)

const (
	RouteFollowersAddUserToFollowUsList = "app/users/follow-us"
	RouteFollowersGetFollowers          = "app/followers"
)

const (
	RouteKeysGetKeys                   = "keys?batchIndex=%s"
	RouteKeysAddKey                    = "keys"
	RouteKeysAddKeyBasedOnTemplate     = "keys/by-template?templateId=%s&excludeCourses=%s&excludeAssessments=%s&excludeTasks=%s"
	RouteKeysGetKey                    = "keys/%s"
	RouteKeysUpdateKey                 = "keys/%s"
	RouteKeysGetArchivedKeys           = "keys/archived"
	RouteKeysGetKeysLinkedToAssessment = "assessments/%s/linked-to/keys?keyId=%s"
	RouteKeysGetKeysLinkedToCourse     = "courses/%s/linked-to/keys?keyId=%s"
	RouteKeysGetKeysFilteredByType     = "keys/filtered/by-type?keyType=%s"
	RouteKeysBulkArchiveKeys           = "keys/archive?keyIds=%s"
	RouteKeysArchiveKey                = "keys/%s/archive"
	RouteKeysUnarchiveKey              = "keys/%s/unarchive"
	RouteKeysUpdateKeyDescription      = "keys/%s/description"
)

const (
	RouteKeysGetCoursesAndAssessmentsAssociatedWithKey           = "charts/keys/%s/courses-assessments"
	RouteKeysGetFilteredUserKeysCoursesAndAssessmentsForGivenKey = "charts/keys/%s/filters"
	RouteKeysGetCourseTypesAndCoursesBasedOnThemInKey            = "charts/keys/%s/course-types"
	RouteKeysGetAssessmentsBasedOnAssessmentTypesInKey           = "charts/keys/%s/assessment-types"
	RouteKeysGetScalesAlongWithCoursesAndAssessmentsBasedOnThem  = "charts/keys/%s/scales"
	RouteKeysGetLinkedResources                                  = "charts/keys/%s/linked-resources"
	RouteKeysGetCourseScaleValues                                = "charts/keys/%s/scales/%s/scale-values"
	RouteKeysGetTaskStatus                                       = "charts/keys/%s/task-status"
)

const (
	RouteKeysGetKeyChecklists         = "keys/%s/checklists"
	RouteKeysAddKeyChecklist          = "keys/%s/checklists"
	RouteKeysReorderKeyChecklists     = "keys/%s/checklists/reorder"
	RouteKeysDeleteKeyChecklist       = "keys/%s/checklists/%s"
	RouteKeysRenameKeyChecklist       = "keys/%s/checklists/%s"
	RouteKeysAddKeyChecklistItem      = "keys/%s/checklists/%s/checklist-items"
	RouteKeysUpdateKeyChecklistItem   = "keys/%s/checklists/%s/checklist-items/%s"
	RouteKeysDeleteKeyChecklistItem   = "keys/%s/checklists/%s/checklist-items/%s"
	RouteKeysReorderKeyChecklistItems = "keys/%s/checklists/%s/checklist-items/reorder"
)

const (
	RouteKeysGetKeyNotes   = "keys/%s/notes"
	RouteKeysAddKeyNote    = "keys/%s/notes"
	RouteKeysUpdateKeyNote = "key-notes/%s?keyId=%s"
	RouteKeysDeleteKeyNote = "key-notes/%s?keyId=%s"
)

const (
	RouteKeysGetKeyTasks     = "keys/%s/tasks"
	RouteKeysAddKeyTask      = "keys/%s/tasks"
	RouteKeysUpdateKeyTask   = "key-tasks/%s?keyId=%s"
	RouteKeysDeleteKeyTask   = "key-tasks/%s?keyId=%s"
	RouteKeysReorderKeyTasks = "keys/%s/tasks/reorder"
)

const (
	RouteNotificationsGetNotifications              = "notifications"
	RouteNotificationsGetUnreadNotifications        = "notifications/unread"
	RouteNotificationsGetUnreadNotificationCount    = "notifications/unread/count"
	RouteNotificationsMarkNotificationAsRead        = "notifications/%s/read"
	RouteNotificationsMarkNotificationsAsReadInBulk = "notifications/read"
)

const (
	RouteAssessmentTypesGetAssessmentTypes                = "assessment-types?includeCounts=%s"
	RouteAssessmentTypesAddAssessmentType                 = "assessment-types"
	RouteAssessmentTypesUpdateAssessmentType              = "assessment-types/%s"
	RouteAssessmentTypesDeleteAssessmentType              = "assessment-types/%s"
	RouteAssessmentTypesGetAssessmentsUsingAssessmentType = "assessment-types/%s/assessments"
)

const (
	RouteProfileGetUsersProfile              = "profiles"
	RouteProfileUpdateUsersProfile           = "profiles"
	RouteProfileUpdateUsername               = "profiles/username/%s"
	RouteProfileBlockUserFromSendingMessages = "users/%s/block"
	RouteProfileUnblockUser                  = "users/%s/unblock"
)

const (
	RouteRegistrationRegisterNewUserByEmail = "app/users/sign-up"
	RouteRegistrationSignInByEmail          = "app/users/sign-in"
	RouteRegistrationResetPassword          = "app/users/reset-password"
	RouteRegistrationActivateUser           = "app/user-verified/%s"
)

const (
	RouteRelationsGetRelationsForKeyMatchingSearchToken        = "search/relations?token=%s&currentKeyId=%s"
	RouteRelationsGetRelationsForCourseMatchingSearchToken     = "search/relations?token=%s&currentCourseId=%s"
	RouteRelationsGetRelationsForAssessmentMatchingSearchToken = "search/relations?token=%s&currentAssessmentId=%s&keyId=%s&courseId=%s"

	RouteRelationsRelateAssessmentToKey            = "keys/%s/assessments/%s/relate?targetKeyId=%s&targetCourseId=%s"
	RouteRelationsUnrelateAssessmentFromKey        = "keys/%s/assessments/%s/unrelate?targetKeyId=%s&targetCourseId=%s"
	RouteRelationsRelateAssessmentToCourse         = "courses/%s/assessments/%s/relate?targetKeyId=%s&targetCourseId=%s"
	RouteRelationsUnrelateAssessmentFromCourse     = "courses/%s/assessments/%s/unrelate?targetKeyId=%s&targetCourseId=%s"
	RouteRelationsRelateAssessmentToAssessment     = "assessments/%s/assessments/%s/relate?sourceKeyId=%s&sourceCourseId=%s&targetKeyId=%s&targetCourseId=%s"
	RouteRelationsUnrelateAssessmentFromAssessment = "assessments/%s/assessments/%s/unrelate?sourceKeyId=%s&sourceCourseId=%s&targetKeyId=%s&targetCourseId=%s"
)

const (
	RouteRelationsGetRelationsForKey        = "keys/%s/relations"
	RouteRelationsGetRelationsForCourse     = "courses/%s/relations?keyId=%s"
	RouteRelationsGetRelationsForAssessment = "assessments/%s/relations?keyId=%s&courseId=%s"
	RouteRelationsRelateKeyToKey            = "keys/%s/keys/%s/relate"
	RouteRelationsUnrelateKeyFromKey        = "keys/%s/keys/%s/unrelate"
	RouteRelationsRelateCourseToKey         = "keys/%s/courses/%s/relate"
	RouteRelationsUnrelateCourseFromKey     = "keys/%s/courses/%s/unrelate"
	RouteRelationsRelateCourseToCourse      = "courses/%s/courses/%s/relate"
	RouteRelationsUnrelateCourseFromCourse  = "courses/%s/courses/%s/unrelate"
)

const (
	RouteScalesGetScales                = "scales?includeCounts=%s&excludeEmpty=%s"
	RouteScalesAddScale                 = "scales"
	RouteScalesGetScale                 = "scales/%s"
	RouteScalesUpdateScale              = "scales/%s"
	RouteScalesDeleteScale              = "scales/%s"
	RouteScalesGetCoursesUsingScale     = "scales/%s/courses"
	RouteScalesGetAssessmentsUsingScale = "scales/%s/assessments"
)

const (
	RouteSchedulerGetEventsInGivenWindow = "scheduler/all-events?startDate=%s&endDate=%s"
	RouteSchedulerGetEventsForGivenDay   = "scheduler/all-events/by-start-date?startDate=%s"
	RouteSchedulerGetStandaloneEvents    = "scheduler/standalone-events"
	RouteSchedulerAddStandaloneEvent     = "scheduler/standalone-events"
	RouteSchedulerUpdateStandaloneEvent  = "scheduler/standalone-events/%s"
	RouteSchedulerDeleteStandaloneEvent  = "scheduler/standalone-events/%s"
)

const (
	RouteSearchSearchKeyCourseOrAssessmentByToken = "search?token=%s"
	RouteSearchSearchUserByToken                  = "search/users?token=%s"
)

const (
	RouteTeacherKeysGetAttachmentSubmissionsAsStudent = "assessments/%s/submissions/attachments/as-student?keyId=%s&courseId=%s"
	RouteTeacherKeysGetCommentSubmissionsAsStudent    = "assessments/%s/submissions/comments/as-student?keyId=%s&courseId=%s"
	RouteTeacherKeysGetStudentsInACourse              = "courses/%s/students?keyId=%s"
)

const (
	RouteTeacherKeysGetStudentAttachmentSubmissionsAsTeacher          = "assessments/%s/submissions/attachments/as-teacher?studentId=%s&keyId=%s&courseId=%s"
	RouteTeacherKeysGetStudentCommentSubmissionsAsTeacher             = "assessments/%s/submissions/comments/as-teacher?studentId=%s&keyId=%s&courseId=%s"
	RouteTeacherKeysAddAttachmentToTeacherAssessmentAsTeacher         = "assessments/%s/attachments/as-teacher?keyId=%s&courseId=%s"
	RouteTeacherKeysAddCommentToTeacherAssessmentAsTeacher            = "assessments/%s/comments/as-teacher?keyId=%s&courseId=%s"
	RouteTeacherKeysGetCourseAndAssessmentsGradesForAStudentAsTeacher = "courses/%s/student-grades/as-teacher?studentUserId=%s&keyId=%s"
	RouteTeacherKeysPublishStudentGradesForACourse                    = "courses/%s/student-grades/publish?keyId=%s"
	RouteTeacherKeysBulkPublishAssessmentGradesForAStudent            = "assessments/students/%s/grades/publish?keyId=%s&courseId=%s"
	RouteTeacherKeysBulkPublishAssessmentGradesForStudents            = "assessments/%s/students/grades/publish?keyId=%s&courseId=%s"
	RouteTeacherKeysGetCourseGradesForStudents                        = "courses/%s/students/grades?keyId=%s"
	RouteTeacherKeysGetAssessmentGradesForStudents                    = "assessments/%s/students/grades?keyId=%s&courseId=%s"
	RouteTeacherKeysAssignGradeToStudent                              = "courses/%s/student/grade?studentUserId=%s&keyId=%s"
	RouteTeacherKeysAssignAssessmentGradeForAStudentAsTeacher         = "assessments/%s/student/grade?studentUserId=%s&keyId=%s&courseId=%s"
	RouteTeacherKeysGetStudentProfile                                 = "students/%s/profile?keyId=%s&courseId=%s"
)

const (
	RouteTemplatesGetKeyTemplates        = "templates/keys"
	RouteTemplatesGetCourseTemplates     = "templates/courses"
	RouteTemplatesGetAssessmentTemplates = "templates/assessments"
)

const (
	RouteUserGetUsers              = "users"
	RouteUserGetUserByUuid         = "users/uuid/%s"
	RouteUserGetUserByEmail        = "users/email/%s"
	RouteUserDeactivateUserAccount = "users/deactivate-account"
)

const (
	RouteVersionGetLatestVersion = "app/latest-version"
	RouteVersionGetAppStatus     = "app/status"
)
