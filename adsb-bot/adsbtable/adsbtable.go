package adsbtable

import (
	"fmt"
	"sync"
	"time"

	adsb "github.com/skypies/adsb"
)

// AdsbTable contains active ADS-B records indexed by 24-bit ICAO aircraft address
type AdsbTable struct {
	rec map[adsb.IcaoId]adsb.Msg
	syn sync.RWMutex
}

// NewTable creates new table of ADS-B entries
func NewTable() *AdsbTable {
	tb := &AdsbTable{}
	tb.rec = make(map[adsb.IcaoId]adsb.Msg)
	return tb
}

// Length returns current length of ADS-B table
func (tb *AdsbTable) Length() int {
	tb.syn.Lock()
	defer tb.syn.Unlock()

	return len(tb.rec)
}

// Update parses SBS-1 string and updates ADS-B table
func (tb *AdsbTable) Update(sbs string) (adsb.IcaoId, bool) {
	var new adsb.Msg
	var cur adsb.Msg

	var tsUpdated bool
	var updated bool
	var ok bool

	tsUpdated = false
	updated = false

	tb.syn.Lock()
	defer tb.syn.Unlock()

	if err := new.FromSBS1(sbs); err != nil {
		return "", false
	}

	id := new.Icao24

	// FIXME: incomplete dump1090 output ?
	if id == "000000" {
		return "", false
	}

	if cur, ok = tb.rec[id]; !ok {
		cur.FromSBS1(fmt.Sprintf("MSG,1,,,%s,,1970/01/01,00:00:00.000,1970/01/01,00:00:00.000,,,,,,,,,,,,0", id))
		tb.rec[id] = cur
		updated = true
	}

	if new.HasCallsign() {
		if cur.Callsign != new.Callsign {
			cur.Callsign = new.Callsign
			updated = true
		}
	}

	if new.HasGroundSpeed() {
		if cur.GroundSpeed != new.GroundSpeed {
			cur.GroundSpeed = new.GroundSpeed
			updated = true
		}
	}

	if new.HasPosition() {
		if cur.Position != new.Position {
			cur.Position = new.Position
			updated = true
		}
	}

	if new.HasVerticalRate() {
		if cur.VerticalRate != new.VerticalRate {
			cur.VerticalRate = new.VerticalRate
			updated = true
		}
	}

	if new.Altitude != 0 {
		if cur.Altitude != new.Altitude {
			cur.Altitude = new.Altitude
			updated = true
		}
	}

	if new.GeneratedTimestampUTC.After(cur.GeneratedTimestampUTC) {
		cur.GeneratedTimestampUTC = new.GeneratedTimestampUTC
		// no need to report update: only last generated timestamp updated
		tsUpdated = true
	}

	if updated || tsUpdated {
		tb.rec[id] = cur
	}

	return id, updated
}

// Get returns ADS-B entry for specified 24-bit ICAO aircraft address
func (tb *AdsbTable) Get(id adsb.IcaoId) (adsb.Msg, bool) {
	tb.syn.Lock()
	defer tb.syn.Unlock()

	ac, ok := tb.rec[id]

	return ac, ok
}

// GetString returns ADS-B string for specified 24-bit ICAO aircraft address
func (tb *AdsbTable) GetString(id adsb.IcaoId) string {
	tb.syn.Lock()
	defer tb.syn.Unlock()

	return tb.getString(id)
}

// Age removes outdated ADS-B entries
func (tb *AdsbTable) Age(time time.Time) {
	tb.syn.Lock()
	defer tb.syn.Unlock()

	for id, rec := range tb.rec {
		if rec.GeneratedTimestampUTC.Before(time) {
			delete(tb.rec, id)
		}
	}
}

// Summary provides basic info for active records
func (tb *AdsbTable) Summary() map[adsb.IcaoId]string {
	tb.syn.Lock()
	defer tb.syn.Unlock()

	m := make(map[adsb.IcaoId]string)

	for id := range tb.rec {
		m[id] = tb.getString(id)
	}

	return m
}

// Internal use only: assumed to be used under syn lock
func (tb *AdsbTable) getString(id adsb.IcaoId) string {
	var m adsb.Msg
	var s string
	var ok bool

	if m, ok = tb.rec[id]; !ok {
		return ""
	}

	if m.HasCallsign() {
		s = fmt.Sprintf("%s (%s)", m.Callsign, m.Icao24)
	} else {
		s = fmt.Sprintf("UNKNOWN (%s)", m.Icao24)
	}

	if m.Altitude != 0 {
		s += fmt.Sprintf(" ALT [%d]", m.Altitude)
	}

	if m.HasGroundSpeed() {
		s += fmt.Sprintf(" GND SPEED [%d]", m.GroundSpeed)
	}

	if m.HasVerticalRate() {
		s += fmt.Sprintf(" VERT SPEED [%d]", m.VerticalRate)
	}

	if m.HasPosition() {
		s += fmt.Sprintf(" POS [%s]", m.Position)
	}

	return s
}
