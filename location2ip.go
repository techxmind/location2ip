package location2ip

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrInvalidGeo = errors.New("Invalid geo")

func Generate(geo string) (ip string, err error) {
	var (
		ipRanges []uint32
		ok       bool
	)

	if ipRanges, ok = _geoIpData[geo]; !ok {
		if geoIDs, ok := _locationGeoIDMapper[geo]; ok {
			n := uint32(len(geoIDs))
			geoID := geoIDs[rand.Uint32()%n]
			if ipRanges, ok = _geoIpData[geoID]; !ok {
				return "", ErrInvalidGeo
			}
		} else {
			return "", ErrInvalidGeo
		}
	}

	randIndex := (rand.Uint32() % uint32(len(ipRanges)/2)) * 2
	ipRangeStart := ipRanges[randIndex]
	ipRangeEnd := ipRanges[randIndex+1]

	ipNum := ipRangeStart

	if ipRangeEnd > ipRangeStart {
		ipNum += rand.Uint32() % uint32(ipRangeEnd-ipRangeStart)
	}

	ip = fmt.Sprintf(
		"%d.%d.%d.%d",
		(ipNum>>24)&0xFF,
		(ipNum>>16)&0xFF,
		(ipNum>>8)&0xFF,
		ipNum&0xFF,
	)

	return
}
