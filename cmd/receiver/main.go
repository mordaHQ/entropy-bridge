package main

/*
#cgo LDFLAGS: -L. -lportaudio_x64
#cgo CFLAGS: -I.
*/
import "C"

import (
	"fmt"
	"math"

	"github.com/gordonklaus/portaudio"
)

const (
	SampleRate = 44100.0
	FreqZero   = 2100.0
	FreqOne    = 3100.0
	BaudRate   = 5
	DeviceID   = 1
	Threshold  = 0.005
)

const BufferSize = 8820

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	devices, err := portaudio.Devices()
	if err != nil {
		panic(err)
	}
	mic := devices[DeviceID]

	in := make([]float32, BufferSize)
	streamParams := portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{
			Device: mic, Channels: 1, Latency: mic.DefaultLowInputLatency,
		},
		SampleRate: SampleRate, FramesPerBuffer: len(in),
	}

	stream, err := portaudio.OpenStream(streamParams, in)
	if err != nil {
		panic(err)
	}
	stream.Start()
	defer stream.Close()

	fmt.Printf("🎤 Слушаю: %s\n", mic.Name)
	fmt.Println("📡 ФИЛЬТР: Игнорирую ЗАГЛАВНЫЕ, ловлю строчные.")
	fmt.Println("------------------------------------------------")

	state := "IDLE"
	var currentByte byte = 0
	var bitCount int = 0

	for {
		stream.Read()

		mag0 := goertzelMag(in, FreqZero, SampleRate)
		mag1 := goertzelMag(in, FreqOne, SampleRate)

		if mag0 < Threshold && mag1 < Threshold {
			continue
		}

		var bit int
		if mag0 > mag1 {
			bit = 0
		} else {
			bit = 1
		}

		if state == "IDLE" {
			if bit == 0 {
				state = "READING"
				currentByte = 0
				bitCount = 0
			}
		} else if state == "READING" {
			currentByte = (currentByte << 1) | byte(bit)
			bitCount++

			if bitCount == 8 {
				// --- УМНЫЙ ФИЛЬТР ---
				// 1. WARMUP (большие буквы) -> Игнорируем
				// 2. github (маленькие буквы) -> Печатаем

				isValid := false

				// Разрешаем только маленькие буквы a-z
				if currentByte >= 'a' && currentByte <= 'z' {
					isValid = true
				}
				// Разрешаем символы URL
				if currentByte == '.' || currentByte == '/' || currentByte == ':' {
					isValid = true
				}
				// Разрешаем цифры
				if currentByte >= '0' && currentByte <= '9' {
					isValid = true
				}

				if isValid {
					fmt.Print(string(currentByte))
				}

				state = "IDLE"
			}
		}
	}
}

func goertzelMag(samples []float32, targetFreq, rate float64) float64 {
	k := int(0.5 + (float64(len(samples)) * targetFreq / rate))
	w := (2.0 * math.Pi / float64(len(samples))) * float64(k)
	cosine := math.Cos(w)
	sine := math.Sin(w)
	coeff := 2.0 * cosine
	q1 := 0.0
	q2 := 0.0
	for _, s := range samples {
		q0 := coeff*q1 - q2 + float64(s)
		q2 = q1
		q1 = q0
	}
	real := (q1 - q2*cosine)
	imag := (q2 * sine)
	return math.Sqrt(real*real+imag*imag) / float64(len(samples))
}
