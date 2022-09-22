// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"net/url"
	"sync"
	"time"
)

const DefaultExpirationTime time.Duration = time.Minute * 5

const (
	none opType = iota
	read
	write
	all
)

type opType int

type locationUnavailabilityInfo struct {
	lastCheckTime  time.Time
	unavailableOps opType
}

type dbAcctLocationsInfo struct {
	prefLocations              []string
	availWriteLocations        []string
	availReadLocations         []string
	availWriteEndptsByLocation map[string]url.URL
	availReadEndptsByLocation  map[string]url.URL
	writeEndpts                []url.URL
	readEndpts                 []url.URL
}

type AcctRegion struct {
	name     string
	endpoint string // make url?
	props    map[string]interface{}
}
type AcctProperties struct {
	readRegions                  []AcctRegion
	writeRegions                 []AcctRegion
	enableMultipleWriteLocations bool
}

type LocationCache struct {
	locationInfo                      dbAcctLocationsInfo
	defaultEndpt                      url.URL
	enableEndptDiscovery              bool
	useMultipleWriteLocations         bool
	connLimit                         int
	updateMutex                       sync.Mutex
	mapMutex                          sync.Mutex // possibly try RWMutex
	locationUnavailabilityInfoMap     map[url.URL]locationUnavailabilityInfo
	lastUpdateTime                    time.Time
	enableMultipleWriteLocations      bool
	unavailableLocationExpirationTime time.Duration
}

func NewLocationCache(prefLocations []string, defaultEndpt url.URL) *LocationCache {
	return &LocationCache{
		defaultEndpt:                  defaultEndpt,
		locationInfo:                  *NewdbdbAcctLocationsInfo(prefLocations, defaultEndpt),
		locationUnavailabilityInfoMap: make(map[url.URL]locationUnavailabilityInfo),
	}
}

func (lc *LocationCache) update(writeLocations []AcctRegion, readLocations []AcctRegion, prefList []string, enableMultipleWriteLocations bool) {
	lc.updateMutex.Lock()
	nextLoc := Copy(lc.locationInfo)
	if prefList != nil {
		nextLoc.prefLocations = prefList
	}
	lc.enableMultipleWriteLocations = enableMultipleWriteLocations
	lc.RefreshStaleEndpts()
	if readLocations != nil {

	}

	if writeLocations != nil {

	}

	//write/read endpts
	lc.lastUpdateTime = time.Now()
	lc.updateMutex.Unlock()
}

func (lc *LocationCache) ReadEndpoints() []url.URL {
	if time.Since(lc.lastUpdateTime) > lc.unavailableLocationExpirationTime && len(lc.locationUnavailabilityInfoMap) > 0 {
		lc.update(nil, nil, nil, lc.enableMultipleWriteLocations)
	}
	return lc.locationInfo.readEndpts
}

func (lc *LocationCache) WriteEndpoints() []url.URL {
	if time.Since(lc.lastUpdateTime) > lc.unavailableLocationExpirationTime && len(lc.locationUnavailabilityInfoMap) > 0 {
		lc.update(nil, nil, nil, lc.enableMultipleWriteLocations)
	}
	return lc.locationInfo.writeEndpts
}

func (lc *LocationCache) GetLocation(endpoint url.URL) string {
	firstLoc := ""
	for location, uri := range lc.locationInfo.availWriteEndptsByLocation {
		if uri == endpoint {
			return location
		}
		if firstLoc == "" {
			firstLoc = location
		}
	}

	for location, uri := range lc.locationInfo.availReadEndptsByLocation {
		if uri == endpoint {
			return location
		}
	}

	if endpoint == lc.defaultEndpt && !lc.CanUseMultipleWriteLocs() {
		if len(lc.locationInfo.availWriteEndptsByLocation) > 0 {
			return firstLoc
		}
	}
	return ""
}

func (lc *LocationCache) CanUseMultipleWriteLocs() bool {
	return lc.useMultipleWriteLocations && lc.enableMultipleWriteLocations
}

func (lc *LocationCache) MarkEndptUnavailableForRead(endpoint url.URL) {
	lc.MarkEndptUnavailable(endpoint, read)
}

func (lc *LocationCache) MarkEndptUnavailableForWrite(endpoint url.URL) {
	lc.MarkEndptUnavailable(endpoint, write)
}

func (lc *LocationCache) MarkEndptUnavailable(endpoint url.URL, op opType) {
	currTime := time.Now()
	lc.mapMutex.Lock()
	if info, ok := lc.locationUnavailabilityInfoMap[endpoint]; ok {
		info.lastCheckTime = currTime
		info.unavailableOps |= op
		lc.locationUnavailabilityInfoMap[endpoint] = info
	} else {
		info = locationUnavailabilityInfo{
			lastCheckTime:  currTime,
			unavailableOps: op,
		}
		lc.locationUnavailabilityInfoMap[endpoint] = info
	}
	lc.mapMutex.Unlock()
	lc.update(nil, nil, nil, lc.enableMultipleWriteLocations)
}

func (lc *LocationCache) DbAcctRead(dbAcct AcctProperties) {
	lc.update(dbAcct.writeRegions, dbAcct.readRegions, nil, dbAcct.enableMultipleWriteLocations)
}

func (lc *LocationCache) RefreshStaleEndpts() {
	lc.mapMutex.Lock()
	for endpoint, info := range lc.locationUnavailabilityInfoMap {
		if time.Since(info.lastCheckTime) > lc.unavailableLocationExpirationTime {
			delete(lc.locationUnavailabilityInfoMap, endpoint)
		}
	}
	lc.mapMutex.Unlock()
}

func (lc *LocationCache) IsEndptUnavailable(endpoint url.URL, ops opType) bool {
	lc.mapMutex.Lock()
	info, ok := lc.locationUnavailabilityInfoMap[endpoint]
	if ops == none || !ok || ops&info.unavailableOps == 0 {
		return false
	} else {
		if time.Since(info.lastCheckTime) > lc.unavailableLocationExpirationTime {
			return false
		}
		return true
	}
}

func NewdbdbAcctLocationsInfo(prefLocations []string, defaultEndpt url.URL) *dbAcctLocationsInfo {
	availWriteLocs := make([]string, 0)
	availReadLocs := make([]string, 0)
	availWriteEndptsByLocation := make(map[string]url.URL)
	availReadEndptsByLocation := make(map[string]url.URL)
	writeEndpts := make([]url.URL, 0)
	readEndpts := make([]url.URL, 0)
	writeEndpts = append(writeEndpts, defaultEndpt)
	readEndpts = append(readEndpts, defaultEndpt)
	return &dbAcctLocationsInfo{
		prefLocations:              prefLocations,
		availWriteLocations:        availWriteLocs,
		availReadLocations:         availReadLocs,
		availWriteEndptsByLocation: availWriteEndptsByLocation,
		availReadEndptsByLocation:  availReadEndptsByLocation,
		writeEndpts:                writeEndpts,
		readEndpts:                 readEndpts,
	}
}

func Copy(other dbAcctLocationsInfo) dbAcctLocationsInfo {
	return dbAcctLocationsInfo{
		prefLocations:              other.prefLocations,
		availWriteLocations:        other.availWriteLocations,
		availReadLocations:         other.availReadLocations,
		availWriteEndptsByLocation: other.availWriteEndptsByLocation,
		availReadEndptsByLocation:  other.availReadEndptsByLocation,
		writeEndpts:                other.writeEndpts,
		readEndpts:                 other.readEndpts,
	}
}
