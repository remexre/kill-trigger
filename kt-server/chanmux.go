package main

import "sync"

// A ChanMux multiplexes several channels.
type ChanMux struct {
	chans map[chan map[string]string]struct{}
	lock  sync.Mutex
}

// NewChanMux creates a new ChanMux.
func NewChanMux() *ChanMux {
	return &ChanMux{
		chans: make(map[chan map[string]string]struct{}),
	}
}

// NewChan creates a new channel and adds it to the ChanMux.
func (chm *ChanMux) NewChan() chan map[string]string {
	ch := make(chan map[string]string, 1)

	chm.lock.Lock()
	chm.chans[ch] = struct{}{}
	chm.lock.Unlock()

	return ch
}

// Delete removes a channel from the ChanMux.
func (chm *ChanMux) Delete(ch chan map[string]string) {
	chm.lock.Lock()
	delete(chm.chans, ch)
	chm.lock.Unlock()

	close(ch)
}

// Len returns the number of open channels.
func (chm *ChanMux) Len() int {
	chm.lock.Lock()
	defer chm.lock.Unlock()

	return len(chm.chans)
}

// Send sends through all channels.
func (chm *ChanMux) Send(b map[string]string) {
	chm.lock.Lock()
	for ch := range chm.chans {
		// Non-blocking send
		select {
		case ch <- b:
		default:
		}
	}
	chm.lock.Unlock()
}
