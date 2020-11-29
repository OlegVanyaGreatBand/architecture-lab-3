package main

import "github.com/OlegVanyaGreatBand/architecture-lab-3/server/models"

func main() {
	tablet := models.TelemetryData{
		TabletId:   0,
		TabletName: nil,
	}
	println(tablet.TabletId)
	println(tablet.TabletName)
}
