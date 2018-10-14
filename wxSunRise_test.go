package wxSunRise

import (
	"fmt"
	"testing"
	"time"
)

var dataSunriseSunset = []struct {
	inLatitude  float64
	inLongitude float64
	inYear      int
	inMonth     time.Month
	inDay       int
	outSunrise  time.Time
	outSunset   time.Time
}{
	// 1970-01-01 - prime meridian
	{
		31.045601, 121.399696,
		2018, time.October, 15,
		time.Date(1970, time.January, 1, 5, 59, 54, 0, time.UTC),
		time.Date(1970, time.January, 1, 18, 07, 07, 0, time.UTC),
	},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{
		31.045601, 121.399696,
		1911, time.January, 1,
		time.Date(2000, time.January, 1, 12, 51, 00, 0, time.UTC),
		time.Date(2000, time.January, 1, 21, 50, 36, 0, time.UTC),
	},
	// 2004-04-01 - (52° N, 5° E)
	{
		31.045601, 121.399696,
		2099, time.April, 1,
		time.Date(2004, time.April, 1, 5, 13, 40, 0, time.UTC),
		time.Date(2004, time.April, 1, 18, 13, 27, 0, time.UTC),
	},
}

func TestSunriseSunset(t *testing.T) {
	for _, tt := range dataSunriseSunset {
		vSunrise, vSunset := SunriseSunset(tt.inLatitude, tt.inLongitude, tt.inYear, tt.inMonth, tt.inDay)
		if vSunrise != tt.outSunrise {
			fmt.Println("%s != %s", vSunrise.String(), tt.outSunrise.String())
		}
		if vSunset != tt.outSunset {
			fmt.Println("%s != %s", vSunset.String(), tt.outSunset.String())
		}
		r := vSunrise.Local().String()
		s := vSunset.Local().String()
		fmt.Println(r, s)
	}
}
