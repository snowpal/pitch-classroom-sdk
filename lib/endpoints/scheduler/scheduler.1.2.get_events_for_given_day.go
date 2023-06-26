package scheduler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetEventsForGivenDay(jwtToken string, dateParam request.EventDateParam) (response.AllEvents, error) {
	resAllEvents := response.AllEvents{}
	route, err := helpers2.GetRoute(lib.RouteSchedulerGetEventsForGivenDay, dateParam.StartDate)
	if err != nil {
		return resAllEvents, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAllEvents, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resAllEvents, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		return resAllEvents, err
	}

	err = json.Unmarshal(body, &resAllEvents)
	if err != nil {
		return resAllEvents, err
	}
	return resAllEvents, nil
}
