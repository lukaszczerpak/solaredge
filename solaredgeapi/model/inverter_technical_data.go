package model

type LineData struct {
	AcCurrent     float64 `json:"acCurrent"`
	AcVoltage     float64 `json:"acVoltage"`
	AcFrequency   float64 `json:"acFrequency"`
	ActivePower   float64 `json:"activePower"`
	ApparentPower float64 `json:"apparentPower"`
	ReactivePower float64 `json:"reactivePower"`
}

type InverterTelemetry struct {
	Date             string   `json:"date"`
	TotalActivePower float64  `json:"totalActivePower"`
	DcVoltage        float64  `json:"dcVoltage"`
	TotalEnergy      float64  `json:"totalEnergy"`
	Temperature      float64  `json:"temperature"`
	InverterMode     IStatus  `json:"inverterMode"`
	OperationMode    int      `json:"operationMode"`
	VL1To2           float64  `json:"vL1To2"`
	VL2To3           float64  `json:"vL2To3"`
	VL3To1           float64  `json:"vL3To1"`
	L1Data           LineData `json:"L1Data"`
	L2Data           LineData `json:"L2Data"`
	L3Data           LineData `json:"L3Data"`
}

type InverterTechnicalData struct {
	Data struct {
		Count       int                 `json:"count"`
		Telemetries []InverterTelemetry `json:"telemetries"`
	} `json:"data"`
}
