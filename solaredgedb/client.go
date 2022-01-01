package solaredgedb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type SolarEdgeDB struct {
	client influxdb2.Client
	org    string
	bucket string
}

func New(url, token, org, bucket string) *SolarEdgeDB {
	db := SolarEdgeDB{}
	db.org = org
	db.bucket = bucket
	db.client = influxdb2.NewClient(url, token)
	return &db
}

func (db *SolarEdgeDB) Close() {
	db.client.Close()
}
