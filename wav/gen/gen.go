package main

import (
	"os"

	"code.google.com/p/jra-go/wav"
)

func main() {
	w := new(wav.Wav)
	w.BitsPerSample = 8
	w.Channels = 1
	w.SamplesPerSecond = 44100
	w.Samples[wav.Mono] = make([]int16, w.SamplesPerSecond*5)

	// put in a square wave

	halfWave := 600
	high := int16(255)
	low := int16(0)

	value := low
	for i := 0; i < len(w.Samples[wav.Mono]); i++ {
		if i % halfWave == 0 {
			if halfWave > 10 {
				halfWave -= 1
			}
			if value == low {
				value = high
			} else {
				value = low
			}
		}
		w.Samples[wav.Mono][i] = value
	}

	w.Write(os.Stdout)
}

// ex: ts=2
