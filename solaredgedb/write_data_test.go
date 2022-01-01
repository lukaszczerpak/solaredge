package solaredgedb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/stretchr/testify/assert"
	"solaredge-sync/common"
	"solaredge-sync/solaredgeapi/model"
	"solaredge-sync/util"
	"testing"
	"time"
)

type MockDB struct {
	influxdb2.Client
	writeAPI *MockWriteAPI
}

func (mock *MockDB) WriteAPI(org, bucket string) api.WriteAPI {
	mock.writeAPI = new(MockWriteAPI)
	mock.writeAPI.counter = 0
	return mock.writeAPI
}

type MockWriteAPI struct {
	api.WriteAPI
	counter int
}

func (mock *MockWriteAPI) Flush() {
}

func (mock *MockWriteAPI) Errors() <-chan error {
	return make(chan error)
}

func (mock *MockWriteAPI) WritePoint(point *write.Point) {
	mock.counter++
}

func (mock *MockWriteAPI) GetNumberOfCalls() int {
	return mock.counter
}

func TestSolarEdgeDB_WriteMeasurements(t *testing.T) {
	mockdb := new(MockDB)
	db := &SolarEdgeDB{client: mockdb}
	cfg := &common.AppConfig{}
	cfg.General.Location = time.Local
	cfg.SolarEdge.InverterSn = "aBCd"
	cfg.SolarEdge.SiteId = "1234"
	var data = model.InverterTechnicalData{}
	util.LoadJsonToStruct("../data/solaredge-techdata-2021-09-01-2021-09-02.json", &data)
	db.WriteMeasurements(cfg, data.Data.Telemetries)

	assert.Equal(t, data.Data.Count, mockdb.writeAPI.counter)
}
