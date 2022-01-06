package solaredgeapi

import (
	"solaredge/solaredgeapi/model"
	"time"
)

func (s *Client) FetchData(siteId, inverterSn string, startTime, stopTime time.Time) (*model.InverterTechnicalData, error) {

	startTimeStr := startTime.Format(DATETIME_FORMAT)
	stopTimeStr := stopTime.Format(DATETIME_FORMAT)

	resp, err := s.ApiClient.R().
		SetQueryParam("startTime", startTimeStr).
		SetQueryParam("endTime", stopTimeStr).
		SetPathParams(map[string]string{
			"siteId":     siteId,
			"inverterSn": inverterSn,
		}).
		SetResult(&model.InverterTechnicalData{}).
		Get("/equipment/{siteId}/{inverterSn}/data")

	if err == nil {
		return resp.Result().(*model.InverterTechnicalData), nil
	}

	return nil, err
}
