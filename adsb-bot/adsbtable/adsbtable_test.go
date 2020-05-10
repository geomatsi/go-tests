package adsbtable

import (
	"testing"
	"time"

	"github.com/skypies/adsb"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	table := NewTable()
	assert.Equal(t, 0, table.Length())

	id, ok := table.Update("MSG,5,111,11111,42426B,111111,2020/05/05,21:37:49.938,2020/05/05,21:37:49.924,,8175,,,,,,,0,,0,0")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)

}

func Test2(t *testing.T) {
	table := NewTable()
	assert.Equal(t, 0, table.Length())

	id, ok := table.Update("this,is,invalid,sbs1,record")
	assert.Equal(t, 0, table.Length())
	assert.Equal(t, false, ok)
	assert.Equal(t, adsb.IcaoId(""), id)
}

func Test3(t *testing.T) {
	table := NewTable()
	assert.Equal(t, 0, table.Length())

	v, ok := table.Get("42426B")
	assert.Equal(t, false, ok)
	assert.Equal(t, adsb.IcaoId(""), v.Icao24)

	table.Update("MSG,5,111,11111,42426B,111111,2020/05/05,21:37:49.938,2020/05/05,21:37:49.924,,8175,,,,,,,0,,0,0")
	assert.Equal(t, 1, table.Length())

	v, ok = table.Get("42426B")
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), v.Icao24)
}

func Test4(t *testing.T) {
	table := NewTable()
	assert.Equal(t, 0, table.Length())

	id, ok := table.Update("MSG,5,111,11111,42426B,111111,2020/05/05,21:37:49.938,2020/05/05,21:37:49.924,,8175,,,,,,,0,,0,0")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)

	id, ok = table.Update("MSG,5,111,11111,42426B,111111,2020/05/05,21:37:49.938,2020/05/05,21:37:49.924,,8000,,,,,,,0,,0,0")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)

	id, ok = table.Update("MSG,6,111,11111,42426B,111111,2020/05/05,21:37:50.662,2020/05/05,21:37:50.645,SDM6346 ,,,,,,,3103,0,0,0,0")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)
}

func Test5(t *testing.T) {
	table := NewTable()
	assert.Equal(t, 0, table.Length())

	id, ok := table.Update("MSG,5,111,11111,42426B,111111,2020/05/05,21:37:49.938,2020/05/05,21:37:49.924,,8175,,,,,,,0,,0,0")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)

	// only GeneratedTime updated: no actual update reported
	id, ok = table.Update("MSG,5,111,11111,42426B,111111,2020/05/05,21:37:59.938,2020/05/05,21:37:59.924,,,,,,,,,,,,")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, false, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)

	id, ok = table.Update("MSG,6,111,11111,42426B,111111,2020/05/05,21:37:50.662,2020/05/05,21:37:50.645,SDM6346 ,,,,,,,3103,0,0,0,0")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)
}

func Test6(t *testing.T) {
	table := NewTable()
	assert.Equal(t, 0, table.Length())

	id, ok := table.Update("MSG,5,111,11111,42426B,111111,2020/05/05,21:37:49.938,2020/05/05,21:37:49.924,,,,,,,,,0,,0,0")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)

	id, ok = table.Update("MSG,8,111,11111,1407C4,111111,2020/05/05,21:39:24.538,2020/05/05,21:39:24.494,,,,,,,,,,,,0")
	assert.Equal(t, 2, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("1407C4"), id)

	id, ok = table.Update("MSG,3,111,11111,42426B,111111,2020/05/05,21:39:18.377,2020/05/05,21:39:18.336,,6850,230,,60.00645,30.09212,,,,,,0")
	assert.Equal(t, 2, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)

	id, ok = table.Update("MSG,8,111,11111,1407C4,111111,2020/05/05,21:39:24.538,2020/05/05,21:39:24.494,,9000,,,,,,,,,,0")
	assert.Equal(t, 2, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("1407C4"), id)
}

func Test7(t *testing.T) {
	table := NewTable()
	assert.Equal(t, 0, table.Length())

	id, ok := table.Update("MSG,5,111,11111,42426B,111111,2020/05/05,21:31:00.000,2020/05/05,21:31:01.000,,,,,,,,,0,,0,0")
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), id)

	id, ok = table.Update("MSG,8,111,11111,1407C4,111111,2020/05/05,21:35:00.000,2020/05/05,21:35:01.000,,,,,,,,,,,,0")
	assert.Equal(t, 2, table.Length())
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("1407C4"), id)

	v, ok := table.Get("42426B")
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), v.Icao24)

	cutoff := v.GeneratedTimestampUTC

	v, ok = table.Get("1407C4")
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("1407C4"), v.Icao24)

	// age all entries before 2020/05/05 21:30:00
	table.Age(cutoff.Add(time.Duration(-1 * time.Minute)))

	v, ok = table.Get("42426B")
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("42426B"), v.Icao24)

	v, ok = table.Get("1407C4")
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("1407C4"), v.Icao24)

	// age all entries before 2020/05/05 21:32:00
	table.Age(cutoff.Add(time.Duration(2 * time.Minute)))

	v, ok = table.Get("42426B")
	assert.Equal(t, false, ok)
	assert.Equal(t, adsb.IcaoId(""), v.Icao24)

	v, ok = table.Get("1407C4")
	assert.Equal(t, true, ok)
	assert.Equal(t, adsb.IcaoId("1407C4"), v.Icao24)

	// age all entries before 2020/05/05 21:36:00
	table.Age(cutoff.Add(time.Duration(5 * time.Minute)))

	v, ok = table.Get("42426B")
	assert.Equal(t, false, ok)
	assert.Equal(t, adsb.IcaoId(""), v.Icao24)

	v, ok = table.Get("1407C4")
	assert.Equal(t, false, ok)
	assert.Equal(t, adsb.IcaoId(""), v.Icao24)
}
