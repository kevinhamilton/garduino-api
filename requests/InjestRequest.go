package requests

// InjestRequest is a request to ingest data from sensors.
type InjestRequest struct {
	SensorID int
	Value    int
}
