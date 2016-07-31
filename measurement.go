package main

import (
	"errors"
)

type Device struct {
	MAC         string  `json:"mac"`
	Temperature string  `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

type Devices []*Device

func parseDevices(res [][]interface{}) (Devices, error) {
	devices := Devices{}
	mac2device := make(map[string]*Device)

	for _, item := range res {
		mac, ok := item[0].(string)

		var device *Device
		var exists bool
		if device, exists = mac2device[mac]; !exists {
			device = &Device{MAC: mac}
			devices = append(devices, device)
			mac2device[mac] = device
		}

		measurement, ok := item[1].(string)
		if !ok {
			return nil, errors.New("Cannot parse measurement type")
		}
		switch measurement {
		case "temperature":
			device.Temperature, ok = item[2].(string)
			if !ok {
				return nil, errors.New("Cannot parse temperature")
			}
		}
	}

	return devices, nil
}
