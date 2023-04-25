package lib

const (
	XApiKey      = "REPLACE_WITH_YOUR_KEY"
	XProductCode = "REPLACE_WITH_PRODUCT_CODE"
)

const GatewayHost = "https://gateway.snowpal.com/"

const (
	RouteAttributesGetDisplayableAttributes        = "app/resource/attributes"
	RouteAttributesUpdateKeyDisplayAttributes      = "keys/%s/attributes"
	RouteAttributesUpdateBlockDisplayAttributes    = "blocks/%s/attributes?keyId=%s"
	RouteAttributesUpdateBlockPodDisplayAttributes = "block-pods/%s/attributes?keyId=%s&blockId=%s"
)

const (
	RouteBlocksGetBlocks                             = "keys/%s/blocks?filter=%s&batchIndex=%s&aclWriteOrHigher=%s"
	RouteBlocksAddBlock                              = "keys/%s/blocks"
	RouteBlocksGetBlocksLinkedToPod                  = "pods/:id/linked-to/blocks?keyId=%s"
	RouteBlocksAddBlockBasedOnTemplate               = "keys/%s/blocks/by-template?templateId=%s&excludePods=%s&excludeTasks=%s"
	RouteBlocksLinkBlockToKey                        = "keys/%s/blocks/%s/link"
	RouteBlocksUnlinkBlockFromKey                    = "keys/%s/blocks/%s/unlink"
	RouteBlocksGetBlocksAvailableToBeLinkedToThisKey = "keys/%s/blocks/available-to-link"
	RouteBlocksGetBlock                              = "blocks/%s?keyId=%s"
	RouteBlocksUpdateBlock                           = "blocks/%s?keyId=%s"
	RouteBlocksAddBlockTypeToBlock                   = "blocks/%s/block-types/%s?keyId=%s"
	RouteBlocksDeleteBlockTypeFromBlock              = "blocks/%s/block-types?keyId=%s"
	RouteBlocksAddScaleToBlock                       = "blocks/%s/scales/%s?keyId=%s"
	RouteBlocksDeleteScaleFromBlock                  = "blocks/%s/scales?keyId=%s"
	RouteBlocksUpdateBlockScaleValue                 = "blocks/%s/scale-value?keyId=%s"
	RouteBlocksUpdateBlockDescription                = "blocks/%s/description?keyId=%s"
	RouteBlocksArchiveBlock                          = "blocks/%s/archive?keyId=%s"
	RouteBlocksUnarchiveBlock                        = "blocks/%s/unarchive?keyId=%s"
	RouteBlocksGetArchivedBlocks                     = "blocks/archived?keyId=%s&batchIndex=%s"
	RouteBlocksBulkArchiveBlocks                     = "blocks/archive?keyId=%s"
	RouteBlocksAllowArchivalOfBlock                  = "blocks/%s/allow-archival?keyId=%s"
	RouteBlocksCopyBlock                             = "blocks/%s/copy?keyId=%s&allTasks=%s&podIds=%s&allPods=%s&allChecklists=%s&targetKeyId=%s"
	RouteBlocksMoveBlock                             = "blocks/%s/move?keyId=%s&targetKeyId=%s"
)

const (
	RouteBlocksGetBlockAttachments   = "blocks/%s/attachments?keyId=%s"
	RouteBlocksAddBlockAttachment    = "blocks/%s/attachments?keyId=%s"
	RouteBlocksRenameBlockAttachment = "block-attachments/%s/rename?keyId=%s&blockId=%s"
	RouteBlocksDeleteBlockAttachment = "block-attachments/%s?keyId=%s&blockId=%s"
)

const (
	RouteBlocksGetLinkedBlockPods        = "charts/keys/%s/blocks/%s/linked-resources"
	RouteBlocksGetScaleValuesForScale    = "charts/keys/%s/blocks/%s/scales/%s/grades"
	RouteBlocksGetTaskStatusForBlock     = "charts/keys/%s/blocks/%s/task-status"
	RouteBlocksGetBlockGradesForStudents = "blocks/%s/students/all/grades?keyId=%s"
)

const (
	RouteBlocksGetBlockChecklists         = "blocks/%s/checklists?keyId=%s"
	RouteBlocksAddBlockChecklist          = "blocks/%s/checklists?keyId=%s"
	RouteBlocksReorderBlockChecklists     = "blocks/%s/checklists/reorder?keyId=%s"
	RouteBlocksDeleteBlockChecklist       = "blocks/%s/checklists/%s?keyId=%s"
	RouteBlocksRenameBlockChecklist       = "blocks/%s/checklists/%s?keyId=%s"
	RouteBlocksAddBlockChecklistItem      = "blocks/%s/checklists/%s/checklist-items?keyId=%s"
	RouteBlocksUpdateBlockChecklistItem   = "blocks/%s/checklists/%s/checklist-items/%s?keyId=%s"
	RouteBlocksDeleteBlockChecklistItem   = "blocks/%s/checklists/%s/checklist-items/%s?keyId=%s"
	RouteBlocksReorderBlockChecklistItems = "blocks/%s/checklists/%s/checklist-items/reorder?keyId=%s"
)

const (
	RouteBlocksGetBlockComments   = "blocks/%s/comments?keyId=%s"
	RouteBlocksAddBlockComment    = "blocks/%s/comments?keyId=%s"
	RouteBlocksUpdateBlockComment = "block-comments/%s?keyId=%s&blockId=%s"
	RouteBlocksDeleteBlockComment = "block-comments/%s?keyId=%s&blockId=%s"
)

const (
	RouteBlocksGetBlockNotes   = "blocks/%s/notes?keyId=%s"
	RouteBlocksAddBlockNote    = "blocks/%s/notes?keyId=%s"
	RouteBlocksUpdateBlockNote = "block-notes/%s?keyId=%s&blockId=%s"
	RouteBlocksDeleteBlockNote = "block-notes/%s?keyId=%s&blockId=%s"
)

const (
	RouteBlocksGetBlockTasks     = "blocks/%s/tasks?keyId=%s"
	RouteBlocksAddBlockTask      = "blocks/%s/tasks?keyId=%s"
	RouteBlocksUpdateBlockTask   = "block-tasks/%s?keyId=%s&blockId=%s"
	RouteBlocksDeleteBlockTask   = "block-tasks/%s?keyId=%s&blockId=%s"
	RouteBlocksAssignBlockTask   = "block-tasks/%s/assign?keyId=%s&blockId=%s"
	RouteBlocksUnassignBlockTask = "block-tasks/%s/unassign?keyId=%s&blockId=%s"
	RouteBlocksReorderBlockTasks = "blocks/%s/tasks/reorder?keyId=%s"
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
	RouteBlockTypesGetBlockTypes           = "block-types?includeCounts=%s"
	RouteBlockTypesAddBlockType            = "block-types"
	RouteBlockTypesUpdateBlockType         = "block-types/%s"
	RouteBlockTypesDeleteBlockType         = "block-types/%s"
	RouteBlockTypesGetBlocksUsingBlockType = "block-types/%s/blocks"
)

const (
	RouteCollaborationGetBlockCollaborators                   = "blocks/%s/acl?keyId=%s"
	RouteCollaborationUpdateBlockAcl                          = "blocks/%s/users/%s/acl?keyId=%s"
	RouteCollaborationUnshareBlockFromCollaborator            = "blocks/%s/users/%s/unshare?keyId=%s"
	RouteCollaborationShareBlockWithCollaborator              = "blocks/%s/users/%s/share?keyId=%s"
	RouteCollaborationShareBlockWithCollaboratorAlongWithPods = "blocks/%s/users/%s/share/with/pods?keyId=%s"
	RouteCollaborationGetUsersThisBlockCanBeSharedWith        = "search/blocks/%s/shareable/users?keyId=%s&token=%s"
	RouteCollaborationBulkShareBlocksWithCollaborators        = "blocks/users/%s/share?keyId=%s"
	RouteCollaborationLeaveBlock                              = "blocks/%s/leave?keyId=%s"
)

const (
	RouteCommentsGetRecentComments = "comments"
)

const (
const (
	RouteDashboardGetDashboardDetails              = "dashboard/combined-responses"
	RouteDashboardGetRecentlyModifiedBlocksAndPods = "dashboard/recently-modified"
	RouteDashboardGetUnreadCount                   = "dashboard/unread-count"
	RouteDashboardGetRecentlyModifiedKeys          = "dashboard/recently-modified/keys"
	RouteDashboardGetPodsAndTasksDueShortly        = "dashboard/due-shortly/pods-and-tasks"
	RouteDashboardGetBlocksDueShortly              = "dashboard/due-shortly/blocks"
	RouteDashboardGetUnreadNotifications           = "dashboard/notifications/unread-status"
	RouteDashboardGetUnreadConversations           = "dashboard/conversations/unread-status"
)

const (
	RouteDashboardGetUserKeysBlocksAndPods           = "charts/dashboard/keys-blocks-pods"
	RouteDashboardGetSystemKeysBlocksAndPods         = "charts/dashboard/system-keys"
	RouteDashboardGetFilteredUserKeysBlocksAndPods   = "charts/dashboard/keys/filters"
	RouteDashboardGetFilteredSystemKeysBlocksAndPods = "charts/dashboard/system-keys/filters"
	RouteDashboardGetBlocksBasedOnBlockTypes         = "charts/dashboard/block-types"
	RouteDashboardGetPodsBasedOnPodTypes             = "charts/dashboard/pod-types"
	RouteDashboardGetBlocksAndPodsBasedOnScales      = "charts/dashboard/scales"
	RouteDashboardGetTasksByStatus                   = "charts/dashboard/task-status"
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
	RouteKeysGetKeys               = "keys?batchIndex=%s"
	RouteKeysAddKey                = "keys"
	RouteKeysAddKeyBasedOnTemplate = "keys/by-template?templateId=%s&excludeBlocks=%s&excludePods=%s&excludeTasks=%s"
	RouteKeysGetKey                = "keys/%s"
	RouteKeysUpdateKey             = "keys/%s"
	RouteKeysGetArchivedKeys       = "keys/archived"
	RouteKeysGetKeysLinkedToPod    = "pods/%s/linked-to/keys?keyId=%s"
	RouteKeysGetKeysLinkedToBlock  = "blocks/%s/linked-to/keys?keyId=%s"
	RouteKeysGetKeysFilteredByType = "keys/filtered/by-type?keyType=%s"
	RouteKeysBulkArchiveKeys       = "keys/archive?keyIds=%s"
	RouteKeysArchiveKey            = "keys/%s/archive"
	RouteKeysUnarchiveKey          = "keys/%s/unarchive"
	RouteKeysUpdateKeyDescription  = "keys/%s/description"
)

const (
	RouteKeysGetBlocksAndPodsAssociatedWithKey           = "charts/keys/%s/blocks-pods"
	RouteKeysGetFilteredUserKeysBlocksAndPodsForGivenKey = "charts/keys/%s/filters"
	RouteKeysGetBlockTypesAndBlocksBasedOnThemInKey      = "charts/keys/%s/block-types"
	RouteKeysGetPodsBasedOnPodTypesInKey                 = "charts/keys/%s/pod-types"
	RouteKeysGetScalesAlongWithBlocksAndPodsBasedOnThem  = "charts/keys/%s/scales"
	RouteKeysGetLinkedResources                          = "charts/keys/%s/linked-resources"
	RouteKeysGetKeyPodAndBlockScaleValues                = "charts/keys/%s/scales/%s/scale-values"
	RouteKeysGetTaskStatus                               = "charts/keys/%s/task-status"
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
	RoutePodTypesGetPodTypes         = "pod-types?includeCounts=%s"
	RoutePodTypesAddPodType          = "pod-types"
	RoutePodTypesUpdatePodType       = "pod-types/%s"
	RoutePodTypesDeletePodType       = "pod-types/%s"
	RoutePodTypesGetPodsUsingPodType = "pod-types/%s/pods"
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
