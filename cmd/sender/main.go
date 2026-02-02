package main

/*
#cgo LDFLAGS: -L. -lportaudio_x64
#cgo CFLAGS: -I.
*/
import "C"

import (
	"fmt"
	"math"
	"time"

	"github.com/gordonklaus/portaudio"
)

// НАСТРОЙКИ: High-Alt (Пробивная способность)
const (
	SampleRate = 44100.0
	FreqZero   = 2100.0
	FreqOne    = 3100.0
	BaudRate   = 5
)

var (
	phase       float64
	currentFreq float64 = 0
)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	// ХИТРОСТЬ:
	// Мы шлем РАЗНЫЕ буквы в верхнем регистре.
	// Это создает сложный, меняющийся звук ("мелодию"),
	// который Windows считает "полезным сигналом" и не глушит.
	// А приемник эти большие буквы просто проигнорирует.
	warmup := "WARMUPWARMUPWARMUPWARMUP"
	payload := "github.com/devmo"

	// Склеиваем
	message := warmup + payload

	bits := encodeUART(message)

	fmt.Println("Entropy Bridge: Режим 'Anti-Noise' (Сложный разогрев).")
	fmt.Printf("Сообщение: %s\n", message)
	fmt.Println(">> СТАРТ ЧЕРЕЗ 1 СЕКУНДУ...")
	time.Sleep(1 * time.Second)

	stream, err := portaudio.OpenDefaultStream(0, 1, SampleRate, 0, processAudio)
	if err != nil {
		panic(err)
	}
	stream.Start()
	defer stream.Close()

	// Пилот-тон (2 секунды)
	fmt.Println(">>> PILOT TONE...")
	currentFreq = FreqOne
	time.Sleep(2 * time.Second)

	// Передача
	fmt.Println(">>> ПЕРЕДАЧА...")
	bitDuration := time.Second / time.Duration(BaudRate)

	for _, bit := range bits {
		if bit == 1 {
			currentFreq = FreqOne
			fmt.Print("1")
		} else {
			currentFreq = FreqZero
			fmt.Print("0")
		}
		time.Sleep(bitDuration)
	}

	currentFreq = 0
	fmt.Println("\n>>> ГОТОВО.")
	time.Sleep(1 * time.Second)
}

func processAudio(out []float32) {
	for i := range out {
		if currentFreq == 0 {
			out[i] = 0
			continue
		}
		step := currentFreq / SampleRate
		out[i] = float32(0.5 * math.Sin(2*math.Pi*phase))
		phase += step
		if phase > 1.0 {
			phase -= 1.0
		}
	}
}

func encodeUART(s string) []int {
	var bits []int
	for _, char := range s {
		bits = append(bits, 0)
		b := int(char)
		for i := 7; i >= 0; i-- {
			bits = append(bits, (b>>i)&1)
		}
		bits = append(bits, 1, 1, 1)
	}
	return bits
}
