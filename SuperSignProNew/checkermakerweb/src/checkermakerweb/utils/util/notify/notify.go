/*Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
 */
package notify

import (
	"errors"
	"sync"
	"time"
)

const E_NOT_FOUND = "E_NOT_FOUND"

// returns the current version
func Version() string {
	return "0.2"
}

// internal mapping of event names to observing channels
var events = make(map[string][]chan interface{})

// mutex for touching the event map
var rwMutex sync.RWMutex

// Start observing the specified event via provided output channel
func Start(event string, outputChan chan interface{}) {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	events[event] = append(events[event], outputChan)
}

// Stop observing the specified event on the provided output channel
func Stop(event string, outputChan chan interface{}) error {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	newArray := make([]chan interface{}, 0)
	outChans, ok := events[event]
	if !ok {
		return errors.New(E_NOT_FOUND)
	}
	for _, ch := range outChans {
		if ch != outputChan {
			newArray = append(newArray, ch)
		} else {
			close(ch)
		}
	}
	events[event] = newArray

	return nil
}

// Stop observing the specified event on all channels
func StopAll(event string) error {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	outChans, ok := events[event]
	if !ok {
		return errors.New(E_NOT_FOUND)
	}
	for _, ch := range outChans {
		close(ch)
	}
	delete(events, event)

	return nil
}

// Post a notification (arbitrary data) to the specified event
func Post(event string, data interface{}) error {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	outChans, ok := events[event]
	if !ok {
		return errors.New(E_NOT_FOUND)
	}
	for _, outputChan := range outChans {
		outputChan <- data
	}

	return nil
}

// Post a notification to the specified event using the provided timeout for
// any output channels that are blocking
func PostTimeout(event string, data interface{}, timeout time.Duration) error {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	outChans, ok := events[event]
	if !ok {
		return errors.New(E_NOT_FOUND)
	}
	for _, outputChan := range outChans {
		select {
		case outputChan <- data:
		case <-time.After(timeout):
		}
	}

	return nil
}
