package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

type IStatus int64

var statuses = map[string]IStatus{
	"OFF":           1,
	"SLEEPING":      2,
	"STARTING":      3,
	"MPPT":          4,
	"THROTTLED":     5,
	"SHUTTING_DOWN": 6,
	"FAULT":         7,
	"STANDBY":       8,
}

func (status *IStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if v, found := statuses[s]; found {
		*status = v
	} else {
		return errors.New(fmt.Sprintf("Unknown status: %v", s))
	}
	return nil
}
