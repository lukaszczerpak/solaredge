package solaredgeapi

import (
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/url"
	"solaredge-sync/solaredgeapi/model"
	"solaredge-sync/util"
	"testing"
	"time"
)

func TestFetchData(t *testing.T) {
	client := createRestyClient("")
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~^https://monitoringapi\.solaredge\.com/equipment/\d+/[^/]+/data`,
		httpmock.NewJsonResponderOrPanic(200, util.LoadJsonToMap("../data/solaredge-techdata-2021-09-01-2021-09-02.json")))

	s := &Client{client}
	resp, err := s.FetchData("1234", "dummy", time.Now(), time.Now())

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, 1, httpmock.GetTotalCallCount(), "number of api calls")
	assert.Len(t, resp.Data.Telemetries, 166, "data.telemetries")
	assert.Equal(t, 166, resp.Data.Count, "data.count")
	assert.Equal(t, model.IStatus(5), resp.Data.Telemetries[0].InverterMode, "operation mode")
}

func TestFetchDataError(t *testing.T) {
	client := createRestyClient("")
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~^https://monitoringapi\.solaredge\.com/equipment/\d+/[^/]+/data`,
		httpmock.NewErrorResponder(errors.New("Timeout")))

	s := &Client{client}
	resp, err := s.FetchData("1234", "dummy", time.Now(), time.Now())

	assert.Error(t, err, "expected error")
	assert.Nil(t, resp, "expected nil response")
}

func TestFetchDataUrl(t *testing.T) {
	startTimeStr := "2021-08-13 12:01:02"
	stopTimeStr := "2021-10-11 10:45:59"
	siteId := "1234"
	inverterSn := "aBCd"
	apiKey := "sup3rs3cr3t"
	encodedUrl := fmt.Sprintf("https://monitoringapi.solaredge.com/equipment/%v/%v/data?api_key=%v&endTime=%v&startTime=%v",
		siteId, inverterSn, apiKey, url.QueryEscape(stopTimeStr), url.QueryEscape(startTimeStr))

	client := createRestyClient(apiKey)
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~^https://`,
		httpmock.NewJsonResponderOrPanic(200, make(map[string]string)))

	startTime, _ := time.ParseInLocation(DATETIME_FORMAT, startTimeStr, time.Local)
	stopTime, _ := time.ParseInLocation(DATETIME_FORMAT, stopTimeStr, time.Local)

	s := &Client{client}
	s.FetchData(siteId, inverterSn, startTime, stopTime)

	info := httpmock.GetCallCountInfo()
	assert.Equal(t, 1, info["GET "+encodedUrl])
}
