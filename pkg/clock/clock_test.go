package clock

import (
	"testing"
	"time"
)

func TestNewTimeInterval(t *testing.T) {
	type args struct {
		intervalConf string
	}

	tests := []struct {
		name     string
		args     args
		hours1   int
		minutes1 int
		hours2   int
		minutes2 int
		wantErr  bool
	}{
		{
			"good case 1",
			args{"10:00-12:00"},
			10,
			0,
			12,
			0,
			false,
		},
		{
			"good case 2",
			args{"9:30-18:30"},
			9,
			30,
			18,
			30,
			false,
		},
		{
			"good case 3",
			args{"00:00-00:00"},
			0,
			0,
			0,
			0,
			false,
		},

		{
			"equal times",
			args{"10:00-10:00"},
			10,
			0,
			10,
			0,
			false,
		},
		{
			"invalid",
			args{"12:00"},
			0,
			0,
			0,
			0,
			true,
		},
		{
			"start before end",
			args{"9:00-00:00"},
			0,
			0,
			0,
			0,
			true,
		},
		{
			"start before end",
			args{"12:00-10:00"},
			0,
			0,
			0,
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeInterval(tt.args.intervalConf)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTimeInterval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h, m, _ := got.Start.Clock()
			if h != tt.hours1 || m != tt.minutes1 {
				t.Errorf("invalid parse expected %v %v, actual %v %v", tt.hours1, tt.minutes1, h, m)
			}
			h, m, _ = got.End.Clock()
			if h != tt.hours2 || m != tt.minutes2 {
				t.Errorf("invalid parse expected %v %v, actual %v %v", tt.hours1, tt.minutes1, h, m)
			}
		})
	}
}

func TestTimeInterval_getNextStartTime(t *testing.T) {
	type args struct {
		now time.Time
	}
	time1, err := time.Parse(time.DateTime, "2006-01-02 19:04:05")
	if err != nil {
		t.Fatal(err)
	}
	time2, err := time.Parse(time.DateTime, "2006-01-02 01:04:05")
	if err != nil {
		t.Fatal(err)
	}
	time3, err := time.Parse(time.DateTime, "2006-01-02 10:04:05")
	if err != nil {
		t.Fatal(err)
	}
	nextTime, err := time.Parse(time.DateTime, "2006-01-03 10:00:00")
	if err != nil {
		t.Fatal(err)
	}

	tInterval, err := NewTimeInterval("10:00-18:00")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		tInterval TimeInterval
		args      args
		want      time.Time
	}{
		{"case 1", *tInterval, args{time1}, nextTime},
		{"case 2", *tInterval, args{time2}, nextTime},
		{"case 2", *tInterval, args{time3}, nextTime},
	}
	compare := func(t1, t2 time.Time) bool {
		y, m, d := t1.Date()
		y2, m2, d2 := t1.Date()
		hh, mm, ss := t1.Clock()
		hh2, mm2, ss2 := t2.Clock()
		return y == y2 && m == m2 && d == d2-1 && hh == hh2 && mm == mm2 && ss == ss2
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tInterval.getNextStartTime(tt.args.now); compare(got, tt.want) {
				t.Errorf("TimeInterval.getNextStartTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
