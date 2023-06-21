package relations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

type SearchKeyRelationParam struct {
	Token        string
	CurrentKeyId string
}

type SearchCourseRelationParam struct {
	Token           string
	CurrentCourseId string
	KeyId           string
}

type SearchAssessmentRelationParam struct {
	Token               string
	CurrentAssessmentId string
	KeyId               string
	CourseId            string
}

func searchRelationsMatchingSearchToken(jwtToken string, route string) ([]response.SearchResource, error) {
	resSearchResources := response.SearchResources{}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resSearchResources.Results, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resSearchResources.Results, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resSearchResources.Results, err
	}

	err = json.Unmarshal(body, &resSearchResources)
	if err != nil {
		return resSearchResources.Results, err
	}
	return resSearchResources.Results, nil
}

func SearchRelationsForKeyMatchingSearchToken(
	jwtToken string,
	relationParam SearchKeyRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers2.GetRoute(
		lib.RouteRelationsGetRelationsForKeyMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentKeyId,
	)
	if err != nil {
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		return searchResults, err
	}
	return searchResults, nil
}

func SearchRelationsForCourseMatchingSearchToken(
	jwtToken string,
	relationParam SearchCourseRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers2.GetRoute(
		lib.RouteRelationsGetRelationsForCourseMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentCourseId,
		relationParam.KeyId,
	)
	if err != nil {
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		return searchResults, err
	}
	return searchResults, nil
}

func SearchRelationsForAssessmentMatchingSearchToken(
	jwtToken string,
	relationParam SearchAssessmentRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers2.GetRoute(
		lib.RouteRelationsGetRelationsForAssessmentMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentAssessmentId,
		relationParam.KeyId,
		relationParam.CourseId,
	)
	if err != nil {
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		return searchResults, err
	}
	return searchResults, nil
}
