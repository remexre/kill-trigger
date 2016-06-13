package main

import "sync"

// A ChanMux multiplexes several channels.
type ChanMux struct {
	chans map[chan byte]struct{}
	lock  sync.Mutex
}

// NewChanMux creates a new ChanMux.
func NewChanMux() *ChanMux {
	return &ChanMux{
		chans: make(map[chan byte]struct{}),
	}
}

// NewChan creates a new channel and adds it to the ChanMux.
func (chm *ChanMux) NewChan() chan byte {
	ch := make(chan byte, 1)

	chm.lock.Lock()
	chm.chans[ch] = struct{}{}
	chm.lock.Unlock()

	return ch
}

// Delete removes a channel from the ChanMux.
func (chm *ChanMux) Delete(ch chan byte) {
	chm.lock.Lock()
	delete(chm.chans, ch)
	chm.lock.Unlock()

	close(ch)
}

// Send sends a byte through all channels.
func (chm *ChanMux) Send(b byte) {
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
