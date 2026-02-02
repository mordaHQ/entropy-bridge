package main

/*
#cgo LDFLAGS: -L. -lportaudio_x64
#cgo CFLAGS: -I.
*/
import "C"

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	devices, err := portaudio.Devices()
	if err != nil {
		panic(err)
	}

	fmt.Println("🔎 ПОИСК МИКРОФОНОВ:")
	fmt.Println("------------------------------------------------")

	for i, dev := range devices {
		// Показываем только устройства, у которых есть ВХОД (Input > 0)
		if dev.MaxInputChannels > 0 {
			fmt.Printf("ID: [%d]  Имя: %s  (Каналов: %d)\n", i, dev.Name, dev.MaxInputChannels)
		}
	}
	fmt.Println("------------------------------------------------")
}
