package models

import (
	"github.com/jinzhu/gorm"
)

type Observation struct {
	// Could have a composite primary key of <MessageNumber (constellation) +
	// ReferenceStationId + Epoch>
	gorm.Model
	// MessageNumber encodes constellation atm, could put this into SatelliteData
	// and have each constellation nested under the same "Observation" - though
	// not sure if anything else in Observation can be different per constellation
	MessageNumber          uint16
	// This does not seem to be correctly implemented anywhere at the moment -
	// could use the station name instead (otherwise have the Id link to a table 
	// of ID + station name)
	// Still can't assume that the ReferenceStationId from RTCM is correct
	ReferenceStationId     uint16
	// Should normalize constellation epochs with timestamp
	Epoch                  uint32
	// This can be normalized to a type - spec says 0-4 is not applied, applied,
	// unknown, and reseverd
	ClockSteeringIndicator uint8
	// This can be normalized to a type - spec says 0-4 is internal, external
	// (locked), external (not locked), and unknown
	ExternalClockIndicator uint8
	// This could probably be normalized to SmoothingType - spec says true means
	// divergence-free smoothing and false means any other smoothing type
	SmoothingTypeIndicator bool
	// Could be normalized to seconds (or null for no smoothing)
	SmoothingInterval      uint8
	SatelliteData []SatelliteData `gorm:"foreignkey:ObservationID"`
}

type SatelliteData struct {
	// Not sure if SatelliteData or SignalData should extend gorm.Model as well
	ID                     int `gorm:"primary_key"`
	ObservationID          uint
	SatelliteID            int
	// Can probably be int or some time type
	RoughRangeMilliseconds uint8
	// This is specific for each constellation...
	Extended               uint8
	// Same comment as RangeMilliseconds
	RoughRanges            uint16
	PhaseRangeRates        int16
	SignalData []SignalData `gorm:"foreignkey:SatelliteDataID"`
}

type SignalData struct {
	ID              int `gorm:"primary_key"`
	SatelliteDataID uint
	SignalID        int
	Pseudoranges    int32
	PhaseRanges     int32
	PhaseRangeLocks uint16
	HalfCycles      bool
	CNRs            uint16
	PhaseRangeRates int16
}
