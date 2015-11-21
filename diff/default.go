package diff

import "log"

type defaultCollector struct{}

func init() {
	var collector defaultCollector
	Register("default", collector)
}

func (c defaultCollector) Diff(device *Device) ([]*Result, error) {
	log.Printf("Diff device Type[%s] Name[%s]\n", device.Type, device.Name)
	return nil, nil
}
