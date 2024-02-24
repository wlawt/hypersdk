// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package fetcher

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"

	"github.com/ava-labs/hypersdk/keys"
	"github.com/ava-labs/hypersdk/state"
)

// Fetcher retrieves values on-the-fly and
// ensures that a value is only fetched from
// disk once. Subsequent requests are fetched
// from cache.
type Fetcher struct {
	im        state.Immutable
	cacheLock sync.RWMutex
	Cache     map[string]*FetchData

	TxnsToFetch map[ids.ID]*sync.WaitGroup // Number of keys a txn is waiting on

	stopOnce  sync.Once
	stop      chan struct{}
	err       error
	fetchable chan *task
}

// Data to insert into the cache
type FetchData struct {
	Val    []byte
	Exists bool

	Chunks uint16
}

// task holds the information that a worker needs to fetch values
type task struct {
	ctx      context.Context
	id       ids.ID
	toLookup []string
}

// New creates a new [Fetcher]
func New(numTxs int, concurrency int, im state.Immutable) *Fetcher {
	f := &Fetcher{
		TxnsToFetch: make(map[ids.ID]*sync.WaitGroup, numTxs),
		fetchable:   make(chan *task),
		im:          im,
		Cache:       make(map[string]*FetchData, numTxs),
		stop:        make(chan struct{}),
	}
	for i := 0; i < concurrency; i++ {
		f.createWorker()
	}
	return f
}

func (f *Fetcher) createWorker() {
	go func() {
		for {
			select {
			case t, ok := <-f.fetchable:
				if !ok {
					return
				}
				// Fetch from disk that aren't already in cache
				fmt.Printf("lookup %v\n", t.toLookup)
				for _, k := range t.toLookup {
					fmt.Printf("fetching key %v for %v\n", k, t.id)
					f.cacheLock.Lock()
					if _, ok := f.Cache[k]; ok {
						fmt.Printf("cache hit for key %v %v\n", k, t.id)
						f.TxnsToFetch[t.id].Done()
						f.cacheLock.Unlock()
						continue
					}
					f.cacheLock.Unlock()

					fmt.Printf("cache miss %v %v\n", k, t.id)
					v, err := f.im.GetValue(t.ctx, []byte(k))
					if errors.Is(err, database.ErrNotFound) {
						// Update the cache
						f.cacheLock.Lock()
						f.Cache[k] = &FetchData{nil, false, 0}
						f.cacheLock.Unlock()
						// Decrement the count as we fetched one of the keys
						f.TxnsToFetch[t.id].Done()
						continue
					} else if err != nil {
						f.stopOnce.Do(func() {
							f.err = err
							close(f.stop)
						})
						return
					}

					// We verify that the [NumChunks] is already less than the number
					// added on the write path, so we don't need to do so again here.
					numChunks, ok := keys.NumChunks(v)
					if !ok {
						f.stopOnce.Do(func() {
							f.err = ErrInvalidKeyValue
							close(f.stop)
						})
						return
					}

					f.cacheLock.Lock()
					f.Cache[k] = &FetchData{v, true, numChunks}
					f.cacheLock.Unlock()

					f.TxnsToFetch[t.id].Done()
				}
			case <-f.stop:
				return
			}
		}
	}()
}

// Lookup enqueues a set of stateKey values that we need to fetch from disk, and
// returns a WaitGroup for a given transaction.
func (f *Fetcher) Lookup(ctx context.Context, txID ids.ID, stateKeys state.Keys) *sync.WaitGroup {
	// Get key names
	toLookup := make([]string, 0, len(stateKeys))
	for k := range stateKeys {
		toLookup = append(toLookup, k)
	}
	// Increment counter to number of keys to fetch
	f.TxnsToFetch[txID] = &sync.WaitGroup{}
	f.TxnsToFetch[txID].Add(len(toLookup))
	t := &task{
		ctx:      ctx,
		id:       txID,
		toLookup: toLookup,
	}
	// Go fetch from disk or cache
	f.fetchable <- t
	return f.TxnsToFetch[txID]
}