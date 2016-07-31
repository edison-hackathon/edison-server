package main

import (
	"errors"
	"strings"
	"strconv"
)

type Device struct {
	MAC         string  `json:"mac"`
	Temperature string  `json:"temperature"`
	Humidity    string `json:"humidity"`
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
		case "humidity":
			device.Humidity, ok = item[2].(string)
			if !ok {
				return nil, errors.New("Cannot parse humidity")
			}
		case "latlon":
			rawLatlon, ok := item[2].(string)
			if !ok {
				return nil, errors.New("Cannot parse latitude and longitude")
			}
			latlon := strings.Split(rawLatlon, ",")

			device.Lat, _ = strconv.ParseFloat(latlon[0], 64)
			device.Lat /= 10000000.0
			device.Lon, _ = strconv.ParseFloat(latlon[1], 64)
			device.Lon /= 10000000.0
		default:
			return nil, errors.New("Unknown measurement")
		}
	}

	return devices, nil
}
