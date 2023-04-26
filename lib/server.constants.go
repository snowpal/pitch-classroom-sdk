package lib

const (
	XApiKey      = "REPLACE_WITH_YOUR_KEY"
	XProductCode = "REPLACE_WITH_PRODUCT_CODE"
)

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
	RouteCoursesGetLinkedCourseAssessments = "charts/keys/%s/courses/%s/linked-resources"
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
	RouteBlockPodsGetBlockPods                          = "blocks/%s/pods?batchIndex=%s&keyId=%s"
	RouteBlockPodsAddBlockPod                           = "blocks/%s/pods?keyId=%s"
	RouteBlockPodsAddBlockPodBasedOnTemplate            = "blocks/%s/pods/by-template?keyId=%s&templateId=%s&excludeTasks=%s"
	RouteBlockPodsLinkPodToBlock                        = "blocks/%s/pods/%s/link?keyId=%s"
	RouteBlockPodsUnlinkPodFromBlock                    = "blocks/%s/pods/%s/unlink?keyId=%s"
	RouteBlockPodsGetBlockPod                           = "block-pods/%s?keyId=%s&blockId=%s"
	RouteBlockPodsUpdateBlockPod                        = "block-pods/%s?keyId=%s&blockId=%s"
	RouteBlockPodsUpdateBlockPodCompletionStatus        = "block-pods/%s/by-completion-status?keyId=%s&blockId=%s"
	RouteBlockPodsAddPodTypeToBlockPod                  = "block-pods/%s/pod-types/%s?keyId=%s&blockId=%s"
	RouteBlockPodsDeletePodTypeFromBlockPod             = "block-pods/%s/pod-types?keyId=%s&blockId=%s"
	RouteBlockPodsAddScaleToBlockPod                    = "block-pods/%s/scales/%s?keyId=%s&blockId=%s"
	RouteBlockPodsDeleteScaleFromBlockPod               = "block-pods/%s/scales?keyId=%s&blockId=%s"
	RouteBlockPodsUpdateBlockPodScaleValue              = "block-pods/%s/scale-value?keyId=%s&blockId=%s"
	RouteBlockPodsArchiveBlockPod                       = "block-pods/%s/archive?keyId=%s&block=%s"
	RouteBlockPodsGetArchivedBlockPods                  = "block-pods/archived?batchIndex=%s&keyId=%s&blockId=%s"
	RouteBlockPodsGetPodsAvailableToBeLinkedToThisBlock = "blocks/%s/pods/available-to-link?keyId=%s"
	RouteBlockPodsUnarchiveBlockPod                     = "block-pods/%s/unarchive?keyId=%s&blockId=%s"
	RouteBlockPodsBulkArchiveBlockPods                  = "block-pods/archive?keyId=%s&blockId=%s"
	RouteBlockPodsUpdateBlockPodDescription             = "block-pods/%s/description?keyId=%s&blockId=%s"
	RouteBlockPodsAllowArchivalOfBlockPod               = "block-pods/%s/allow-archival?keyId=%s&blockId=%s"
	RouteBlockPodsCopyBlockPod                          = "block-pods/%s/copy?keyId=%s&blockId=%s&allTasks=%s&allChecklists=%s&targetKeyId=%s&targetBlockId=%s"
	RouteBlockPodsMoveBlockPod                          = "block-pods/%s/move?keyId=%s&blockId=%s&targetKeyId=%s&targetBlockId=%s"
)

const (
	RouteBlockPodsGetBlockPodAttachments   = "block-pods/%s/attachments?keyId=%s&blockId=%s"
	RouteBlockPodsAddBlockPodAttachment    = "block-pods/%s/attachments?keyId=%s&blockId=%s"
	RouteBlockPodsRenameBlockPodAttachment = "block-pod-attachments/%s/rename?keyId=%s&blockId=%s&podId=%s"
	RouteBlockPodsDeleteBlockPodAttachment = "block-pod-attachments/%s?keyId=%s&blockId=%s&podId=%s"
)

const (
	RouteBlockPodsGetBlockPodTasksForCharts    = "charts/block-pods/%s/tasks?keyId=%s&blockId=%s"
	RouteBlockPodsGetBlockPodGradesForStudents = "charts/classroom-pods/%s/students/grades?keyId=%s&blockId=%s"
)

const (
	RouteBlockPodsGetBlockPodChecklists         = "block-pods/%s/checklists?keyId=%s&blockId=%s"
	RouteBlockPodsAddBlockPodChecklist          = "block-pods/%s/checklists?keyId=%s&blockId=%s"
	RouteBlockPodsReorderBlockPodChecklists     = "block-pods/%s/checklists/reorder?keyId=%s&blockId=%s"
	RouteBlockPodsDeleteBlockPodChecklist       = "block-pods/%s/checklists/%s?keyId=%s&blockId=%s"
	RouteBlockPodsRenameBlockPodChecklist       = "block-pods/%s/checklists/%s?keyId=%s&blockId=%s"
	RouteBlockPodsAddBlockPodChecklistItem      = "block-pods/%s/checklists/%s/checklist-items?keyId=%s&blockId=%s"
	RouteBlockPodsUpdateBlockPodChecklistItem   = "block-pods/%s/checklists/%s/checklist-items/%s?keyId=%s&blockId=%s"
	RouteBlockPodsDeleteBlockPodChecklistItem   = "block-pods/%s/checklists/%s/checklist-items/%s?keyId=%s&blockId=%s"
	RouteBlockPodsReorderBlockPodChecklistItems = "block-pods/%s/checklists/%s/checklist-items/reorder?keyId=%s&blockId=%s"
)

const (
	RouteBlockPodsGetBlockPodComments   = "block-pods/%s/comments?keyId=%s&blockId=%s"
	RouteBlockPodsAddBlockPodComment    = "block-pods/%s/comments?keyId=%s&blockId=%s"
	RouteBlockPodsUpdateBlockPodComment = "block-pod-comments/%s?keyId=%s&blockId=%s&podId=%s"
	RouteBlockPodsDeleteBlockPodComment = "block-pod-comments/%s?keyId=%s&blockId=%s&podId=%s"
)

const (
	RouteBlockPodsGetBlockPodNotes   = "block-pods/%s/notes?keyId=%s&blockId=%s"
	RouteBlockPodsAddBlockPodNote    = "block-pods/%s/notes?keyId=%s&blockId=%s"
	RouteBlockPodsUpdateBlockPodNote = "block-pod-notes/%s?keyId=%s&blockId=%s&podId=%s"
	RouteBlockPodsDeleteBlockPodNote = "block-pod-notes/%s?keyId=%s&blockId=%s&podId=%s"
)

const (
	RouteBlockPodsGetBlockPodTasks     = "block-pods/%s/tasks?keyId=%s&blockId=%s"
	RouteBlockPodsAddBlockPodTask      = "block-pods/%s/tasks?keyId=%s&blockId=%s"
	RouteBlockPodsUpdateBlockPodTask   = "block-pod-tasks/%s?keyId=%s&blockId=%s&podId=%s"
	RouteBlockPodsDeleteBlockPodTask   = "block-pod-tasks/%s?keyId=%s&blockId=%s&podId=%s"
	RouteBlockPodsAssignBlockPodTask   = "block-pod-tasks/%s/assign?keyId=%s&blockId=%s&podId=%s"
	RouteBlockPodsUnassignBlockPodTask = "block-pod-tasks/%s/unassign?keyId=%s&blockId=%s&podId=%s"
	RouteBlockPodsReorderBlockPodTasks = "block-pods/%s/tasks/reorder?keyId=%s&blockId=%s"
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
	RouteFavoritesGetFavorites          = "favorites"
	RouteFavoritesAddKeyAsFavorite      = "favorites/keys/%s"
	RouteFavoritesAddBlockAsFavorite    = "favorites/blocks/%s?keyId=%s"
	RouteFavoritesAddBlockPodAsFavorite = "favorites/block-pods/%s?keyId=%s&blockId=%s"
	RouteFavoritesDeleteFavorite        = "favorites/%s"
)

const (
	RouteFollowersAddUserToFollowUsList = "app/users/follow-us"
	RouteFollowersGetFollowers          = "app/followers"
)

const (
	RouteKeysGetKeys                   = "keys?batchIndex=%s"
	RouteKeysAddKey                    = "keys"
	RouteKeysAddKeyBasedOnTemplate     = "keys/by-template?templateId=%s&excludeBlocks=%s&excludePods=%s&excludeTasks=%s"
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
	RouteProfileGetUsersProfile               = "profiles"
	RouteProfileUpdateUsersProfile            = "profiles"
	RouteProfileUpdateUsername                = "profiles/username/%s"
	RouteProfileBlocksUserFromSendingMessages = "users/%s/block"
	RouteProfileUnblocksUser                  = "users/%s/unblock"
)

const (
	RouteRegistrationRegisterNewUserByEmail = "app/users/sign-up"
	RouteRegistrationSignInByEmail          = "app/users/sign-in"
	RouteRegistrationResetPassword          = "app/users/reset-password"
	RouteRegistrationActivateUser           = "app/user-verified/%s"
)

// TODO(Anish,3,03/23/23): As the endpoint is same for key pod & block pod. We need to create these constants to
// support both with the same endpoint with different query string params.
const (
	RouteRelationsGetRelationsForKeyMatchingSearchToken      = "search/relations?token=%s&currentKeyId=%s"
	RouteRelationsGetRelationsForBlockMatchingSearchToken    = "search/relations?token=%s&currentBlockId=%s"
	RouteRelationsGetRelationsForBlockPodMatchingSearchToken = "search/relations?token=%s&currentPodId=%s&keyId=%s&blockId=%s"

	RouteRelationsRelateBlockPodToKey          = "keys/%s/pods/%s/relate?targetKeyId=%s&targetBlockId=%s"
	RouteRelationsUnrelateBlockPodFromKey      = "keys/%s/pods/%s/unrelate?targetKeyId=%s&targetBlockId=%s"
	RouteRelationsRelateBlockPodToBlock        = "blocks/%s/pods/%s/relate?targetKeyId=%s&targetBlockId=%s"
	RouteRelationsUnrelateBlockPodFromBlock    = "blocks/%s/pods/%s/unrelate?targetKeyId=%s&targetBlockId=%s"
	RouteRelationsRelateBlockPodToBlockPod     = "pods/%s/pods/%s/relate?sourceKeyId=%s&sourceBlockId=%s&targetKeyId=%s&targetBlockId=%s"
	RouteRelationsUnrelateBlockPodFromBlockPod = "pods/%s/pods/%s/unrelate?sourceKeyId=%s&sourceBlockId=%s&targetKeyId=%s&targetBlockId=%s"
)

const (
	RouteRelationsGetRelationsForKey      = "keys/%s/relations"
	RouteRelationsGetRelationsForBlock    = "blocks/%s/relations?keyId=%s"
	RouteRelationsGetRelationsForBlockPod = "block-pods/%s/relations?keyId=%s&blockId=%s"
	RouteRelationsRelateKeyToKey          = "keys/%s/keys/%s/relate"
	RouteRelationsUnrelateKeyFromKey      = "keys/%s/keys/%s/unrelate"
	RouteRelationsRelateBlockToKey        = "keys/%s/blocks/%s/relate"
	RouteRelationsUnrelateBlockFromKey    = "keys/%s/blocks/%s/unrelate"
	RouteRelationsRelateBlockToBlock      = "blocks/%s/blocks/%s/relate"
	RouteRelationsUnrelateBlockFromBlock  = "blocks/%s/blocks/%s/unrelate"
	RouteRelationsRelatePodToPod          = "pods/%s/pods/%s/relate?sourceKeyId=%s&targetKeyId=%s"
	RouteRelationsUnrelatePodFromPod      = "pods/%s/pods/%s/unrelate?sourceKeyId=%s&targetKeyId=%s"
)

const (
	RouteScalesGetScales           = "scales?includeCounts=%s&excludeEmpty=%s"
	RouteScalesAddScale            = "scales"
	RouteScalesGetScale            = "scales/%s"
	RouteScalesUpdateScale         = "scales/%s"
	RouteScalesDeleteScale         = "scales/%s"
	RouteScalesGetBlocksUsingScale = "scales/%s/blocks"
	RouteScalesGetPodsUsingScale   = "scales/%s/pods"
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
	RouteSearchSearchKeyBlockOrPodByToken = "search?token=%s"
	RouteSearchSearchUserByToken          = "search/users?token=%s"
)

const (
	RouteTeacherKeysGetAttachmentSubmissionsAsStudent = "classroom-pods/%s/submissions/attachments/as-student?keyId=%s&blockId=%s"
	RouteTeacherKeysGetCommentSubmissionsAsStudent    = "classroom-pods/%s/submissions/comments/as-student?keyId=%s&blockId=%s"
	RouteTeacherKeysGetStudentsInABlock               = "classroom-blocks/%s/students?keyId=%s"
)

const (
	RouteTeacherKeysGetStudentAttachmentSubmissionsAsTeacher  = "classroom-pods/%s/submissions/attachments/as-teacher?studentId=%s&keyId=%s&blockId=%s"
	RouteTeacherKeysGetStudentCommentSubmissionsAsTeacher     = "classroom-pods/%s/submissions/comments/as-teacher?studentId=%s&keyId=%s&blockId=%s"
	RouteTeacherKeysAddAttachmentToTeacherPodAsTeacher        = "classroom-pods/%s/attachments/as-teacher?keyId=%s&blockId=%s"
	RouteTeacherKeysAddCommentToTeacherPodAsTeacher           = "classroom-pods/%s/comments/as-teacher?keyId=%s&blockId=%s"
	RouteTeacherKeysGetBlockAndPodsGradesForAStudentAsTeacher = "classroom-blocks/%s/student-grades/as-teacher?studentUserId=%s&keyId=%s"
	RouteTeacherKeysPublishStudentGradesForABlock             = "classroom-blocks/%s/student-grades/publish?keyId=%s"
	RouteTeacherKeysBulkPublishPodGradesForAStudent           = "classroom-pods/students/%s/grades/publish?keyId=%s&blockId=%s"
	RouteTeacherKeysBulkPublishPodGradesForStudents           = "classroom-pods/%s/students/grades/publish?keyId=%s&blockId=%s"
	RouteTeacherKeysGetBlockGradesForStudents                 = "classroom-blocks/%s/students/grades?keyId=%s"
	RouteTeacherKeysGetPodGradesForStudents                   = "classroom-pods/%s/students/grades?keyId=%s&blockId=%s"
	RouteTeacherKeysAssignGradeToStudent                      = "classroom-blocks/%s/student/grade?studentUserId=%s&keyId=%s"
	RouteTeacherKeysAssignPodGradeForAStudentAsTeacher        = "classroom-pods/%s/student/grade?studentUserId=%s&keyId=%s&blockId=%s"
	RouteTeacherKeysGetStudentProfile                         = "classroom/students/%s/profile?keyId=%s&blockId=%s"
)

const (
	RouteTemplatesGetKeyTemplates   = "templates/keys"
	RouteTemplatesGetBlockTemplates = "templates/blocks"
	RouteTemplatesGetPodTemplates   = "templates/pods"
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
