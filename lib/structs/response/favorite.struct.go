package response

import (
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type Favorites struct {
	Favorites []Favorite `json:"favorites"`
}

type AddFavorite struct {
	ID string `json:"id"`
}

type Favorite struct {
	ID       string                  `json:"id"`
	Modifier common.ResourceModifier `json:"modifier"`
	Resource FavoriteResource        `json:"resource"`
}

type FavoriteKey struct {
	ID       string                  `json:"id"`
	Name     string                  `json:"keyName"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type FavoriteCourse struct {
	ID       string                  `json:"id"`
	Name     string                  `json:"courseName"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type FavoriteResource struct {
	ID             string                  `json:"id"`
	ResourceType   string                  `json:"resourceType"`
	KeyName        *string                 `json:"keyName"`
	Coursename     *string                 `json:"courseName"`
	AssessmentName *string                 `json:"assessmentName"`
	Key            *FavoriteKey            `json:"key"`
	Course         *FavoriteCourse         `json:"course"`
	Modifier       common.ResourceModifier `json:"modifier"`
}
