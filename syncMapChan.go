package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type SaveConfigStopChans struct {
	syncMap sync.Map
}

func (m *SaveConfigStopChans) Load(key string) chan bool {
	val, ok := m.syncMap.Load(key)
	if ok {
		return val.(chan bool)
	} else {
		return nil
	}
}

func (m *SaveConfigStopChans) Store(key string, value chan bool) {
	m.syncMap.Store(key, value)
}

func (m *SaveConfigStopChans) Delete(key string) {
	m.syncMap.Delete(key)
}

func (m *SaveConfigStopChans) Exists(key string) bool {
	_, ok := m.syncMap.Load(key)
	return ok
}

var mapChannel SaveConfigStopChans

func main() {

	startInterval := float64(1000)

	go func() {
		ticker := time.NewTicker(time.Duration(startInterval) * time.Millisecond)
		counter := 1.0

		for {
			select {
			case <-ticker.C:
				log.Println("ticker accelerating to " + fmt.Sprint(startInterval/counter) + " ms")
				ticker = time.NewTicker(time.Duration(startInterval/counter) * time.Millisecond)
				counter++
			case v := <-mapChannel.Load("1"):
				fmt.Println("Received on channel 1: %d", v)
				if v {
					return
				}
				//fmt.Println(" not match value,waiting to be stopped")
			default:
				fmt.Println("waiting to be stopped")
				time.Sleep(1 * time.Second)
			}
		}

	}()
	log.Println("stopping ticker...")
	mapChannel.Store("1", make(chan bool))
	if stopCh := mapChannel.Load("1"); stopCh != nil {
		fmt.Printf("current value on channel 1: %d\n", stopCh)
		stopCh <- true
		fmt.Printf("latest value on channel 1: %d\n", stopCh)
	}
	//reported <- true
	fmt.Println("Should stopped") // just to see quit messages
}
