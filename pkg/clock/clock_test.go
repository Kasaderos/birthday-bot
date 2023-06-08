package clock

import (
	"testing"
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
