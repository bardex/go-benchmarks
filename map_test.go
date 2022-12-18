package main

import (
	"sync"
	"testing"
)

func BenchmarkMaps50_50(b *testing.B) {
	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg := sync.WaitGroup{}
			mu := sync.Mutex{}
			items := map[int]int{}
			for k := 0; k < 1000; k++ {
				k := k
				if k%10 == 0 {
					wg.Add(1)
					go func() {
						defer wg.Done()
						mu.Lock()
						defer mu.Unlock()
						items[k] = k
					}()
				} else {
					wg.Add(1)
					go func() {
						defer wg.Done()
						mu.Lock()
						defer mu.Unlock()
						_ = items[k]
					}()
				}
			}
			wg.Wait()
		}
	})
	b.Run("rw_mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg := sync.WaitGroup{}
			mu := sync.RWMutex{}
			items := map[int]int{}
			for k := 0; k < 1000; k++ {
				k := k
				if k%10 == 0 {
					wg.Add(1)
					go func() {
						defer wg.Done()
						mu.Lock()
						defer mu.Unlock()
						items[k] = k
					}()
				} else {
					wg.Add(1)
					go func() {
						defer wg.Done()
						mu.RLock()
						defer mu.RUnlock()
						_ = items[k]
					}()
				}
			}
			wg.Wait()
		}
	})
	b.Run("sync_map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg := sync.WaitGroup{}
			items := sync.Map{}
			for k := 0; k < 1000; k++ {
				k := k
				if k%10 == 0 {
					wg.Add(1)
					go func() {
						defer wg.Done()
						items.Store(k, k)
					}()
				} else {
					wg.Add(1)
					go func() {
						defer wg.Done()
						_, _ = items.Load(k)
					}()
				}
			}
			wg.Wait()
		}
	})
}
