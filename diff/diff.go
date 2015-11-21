package diff

import (
	"log"
	"sync"
)

var collectors = make(map[string]Collector)

func Run() {
	devices, err := RetrieveDevices()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(devices))

	for _, device := range devices {
		collector, exists := collectors[device.Type]
		if !exists {
			collector = collectors["default"]
		}

		go func(collector Collector, device *Device) {
			Collect(collector, device, results)
			waitGroup.Done()
		}(collector, device)
	}

	go func() {
		waitGroup.Wait()

		close(results)
	}()

	Display(results)
}

func Register(deviceType string, collector Collector) {
	if _, exists := collectors[deviceType]; exists {
		log.Fatalln(deviceType, "Collector already registered")
	}

	log.Println("Register", deviceType, "collector")
	collectors[deviceType] = collector
}
