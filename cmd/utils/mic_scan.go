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

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	devices, err := portaudio.Devices()
	if err != nil {
		panic(err)
	}

	fmt.Println("🕵️‍♂️ ОХОТНИК ЗА ЗВУКОМ: Проверяю все устройства...")
	fmt.Println("(Шуми, хлопай или говори пока идет тест!)")
	fmt.Println("------------------------------------------------")

	for i, dev := range devices {
		// Проверяем только входные устройства
		if dev.MaxInputChannels > 0 {
			fmt.Printf("Testing ID [%d]: %s... ", i, dev.Name)

			// Пытаемся послушать полсекунды
			vol := testDevice(dev)

			if vol > 0.001 {
				fmt.Printf("✅ СИГНАЛ ЕСТЬ! (Громкость: %.5f)\n", vol)
			} else {
				fmt.Printf("❌ Тишина (%.5f)\n", vol)
			}
		}
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("Запомни ID, где была зеленая галочка ✅, и впиши его в receiver.go")
}

func testDevice(info *portaudio.DeviceInfo) float64 {
	in := make([]float32, 1024)
	streamParams := portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{
			Device:   info,
			Channels: 1,
			Latency:  info.DefaultLowInputLatency,
		},
		SampleRate:      44100,
		FramesPerBuffer: len(in),
	}

	stream, err := portaudio.OpenStream(streamParams, in)
	if err != nil {
		return -1 // Ошибка открытия
	}

	stream.Start()

	maxVal := 0.0
	// Слушаем 20 пакетов (около 0.5 сек)
	for j := 0; j < 20; j++ {
		stream.Read()
		for _, v := range in {
			val := math.Abs(float64(v))
			if val > maxVal {
				maxVal = val
			}
		}
	}

	stream.Close()
	return maxVal
}
