package solaredgedb

import (
	"solaredge-sync/solaredgeapi"
	"solaredge-sync/solaredgeapi/model"
	"time"
)

func InverterTelemetryToInflux(t model.InverterTelemetry, loc *time.Location) (time.Time, map[string]interface{}, error) {
	ts, err := time.ParseInLocation(solaredgeapi.DATETIME_FORMAT, t.Date, loc)
	if err != nil {
		return time.Time{}, nil, err
	}

	fields := map[string]interface{}{
		"I_AC_Current":   t.L1Data.AcCurrent + t.L2Data.AcCurrent + t.L3Data.AcCurrent,
		"I_AC_CurrentA":  t.L1Data.AcCurrent,
		"I_AC_CurrentB":  t.L2Data.AcCurrent,
		"I_AC_CurrentC":  t.L3Data.AcCurrent,
		"I_AC_Energy_WH": t.TotalEnergy,
		"I_AC_Power":     t.TotalActivePower,
		"I_AC_VA":        t.L1Data.ApparentPower + t.L2Data.ApparentPower + t.L3Data.ApparentPower,
		"I_AC_VAR":       t.L1Data.ReactivePower + t.L2Data.ReactivePower + t.L3Data.ReactivePower,
		"I_AC_VoltageAB": t.VL1To2,
		"I_AC_VoltageBC": t.VL2To3,
		"I_AC_VoltageCA": t.VL3To1,
		"I_AC_VoltageAN": t.L1Data.AcVoltage,
		"I_AC_VoltageBN": t.L2Data.AcVoltage,
		"I_AC_VoltageCN": t.L3Data.AcVoltage,
		"I_DC_Voltage":   t.DcVoltage,
		"I_Temp":         t.Temperature,
		"I_Status":       int64(t.InverterMode),
	}

	if sumAP := t.L1Data.ApparentPower + t.L2Data.ApparentPower + t.L3Data.ApparentPower; sumAP > 0 {
		fields["I_AC_PF"] = (t.L1Data.ActivePower + t.L2Data.ActivePower + t.L3Data.ActivePower) / sumAP
	}

	return ts, fields, nil
}
