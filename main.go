package main

import (
	"encoding/binary"
	"math"
	"os"
)

const (
	hz        = 8000
	sinWaveHz = 1000
)

func main() {
	f, err := os.OpenFile("test.pcm", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	for t := float64(0); t < 5; t = t + (float64(1) / hz) {
		var sample float64 = 1500 * math.Sin(2*math.Pi*sinWaveHz*t)
		s16 := int16(sample)
		c := byte(s16 % 256)
		binary.Write(f, binary.LittleEndian, c)
		c = byte(s16 / 256 % 256)
		binary.Write(f, binary.LittleEndian, c)
	}

	f.Close()
}
