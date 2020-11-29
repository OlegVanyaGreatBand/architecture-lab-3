package models

type TelemetryData struct {
	TabletId   int          `json:"id"`
	TabletName *string      `json:"name"`
	Telemetry  []*Telemetry `json:"telemetry"`
}

type Telemetry struct {
	Battery int `json:"battery"`
	DeviceTime string `json:"deviceTime"`
	Timestamp string `json:"timestamp"`
	CurrentVideo *string `json:"current_video"`
}
