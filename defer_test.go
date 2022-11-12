package main

import (
	"sync"
	"testing"
)

const countMutex = 10

func deferLoop() (mutex [countMutex]sync.Mutex) {
	for i := 0; i < countMutex; i++ {
		mutex[i].Lock()
		defer mutex[i].Unlock()
	}
	return
}

func deferLoopFunc() (mutex [countMutex]sync.Mutex) {
	for i := 0; i < countMutex; i++ {
		func() {
			mutex[i].Lock()
			defer mutex[i].Unlock()
		}()
	}
	return
}

func deferNoLoop() (mutex [countMutex]sync.Mutex) {
	for i := 0; i < countMutex; i++ {
		mutex[i].Lock()
	}
	for i := 0; i < countMutex; i++ {
		mutex[i].Unlock()
	}
	return
}

func BenchmarkDefer(b *testing.B) {
	b.Run("defer loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mutex := deferLoop()
			// be sure all mutex unlock
			for k := range mutex {
				mutex[k].Lock()
				mutex[k].Unlock()
			}
		}
	})
	b.Run("defer loop func", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mutex := deferLoopFunc()
			// be sure all mutex unlock
			for k := range mutex {
				mutex[k].Lock()
				mutex[k].Unlock()
			}
		}
	})
	b.Run("no defer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mutex := deferNoLoop()
			// be sure all mutex unlock
			for k := range mutex {
				mutex[k].Lock()
				mutex[k].Unlock()
			}
		}
	})
}
