package main

import "errors"

type Measurement struct {
	ID          uint64 `json:"id"`
	Temperature string `json:"temperature"`
}

type Measurements []Measurement

func parseMeasurements(res [][]interface{}) (Measurements, error) {
	measurements := Measurements{}
	for _, item := range res {
		measurement := Measurement{}

		var ok bool
		measurement.ID, ok = item[0].(uint64)
		if !ok {
			return nil, errors.New("Cannot parse ID")
		}
		measurement.Temperature, ok = item[1].(string)
		if !ok {
			return nil, errors.New("Cannot parse temperature")
		}

		measurements = append(measurements, measurement)
	}

	return measurements, nil
}
