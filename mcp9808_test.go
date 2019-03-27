package mcp9808

import (
	"fmt"
	"testing"
)

func TestTempRegisterToCelsius(t *testing.T) {
	cases := []struct {
		reg  uint16
		want float32
	}{
		// The temp register uses two's complement. See pages 24-25 of the datasheet.
		{0x1800, -128.0},
		{0x1f60, -10.0},
		{0x1ff9, -0.4375},
		{0x0000, 0},
		{0x0001, 0.0625},
		{0x0003, 0.1875},
		{0x0007, 0.4375},
		{0x000f, 0.9375},
		{0x001f, 1.9375},
		{0x003f, 3.9375},
		{0x00a0, 10.0},
		{0x0800, 128.0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.want), func(t *testing.T) {
			got := tempRegisterToCelsius(c.reg)
			if got != c.want {
				t.Errorf("Got %v, want %v", got, c.want)
			}
		})
	}
}
