package main

import "testing"

type ICounter interface {
	Add(int) error
}

type MemoryCounter struct {
	counter int
}

func (m *MemoryCounter) Add(i int) error {
	m.counter += i
	return nil
}

type ServiceWithInterface struct {
	Cnt ICounter
}

func (a *ServiceWithInterface) Add(i int) error {
	return a.Cnt.Add(i)
}

type ServiceWithConcreteType struct {
	Cnt *MemoryCounter
}

func (a *ServiceWithConcreteType) Add(i int) error {
	return a.Cnt.Add(i)
}

func BenchmarkInterfaces(b *testing.B) {
	interCounter := ServiceWithInterface{Cnt: &MemoryCounter{}}
	concrCounter := ServiceWithConcreteType{Cnt: &MemoryCounter{}}

	b.Run("interface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for k := 0; k < 1000_000; k++ {
				err := interCounter.Add(k)
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})
	b.Run("concrete type", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for k := 0; k < 1000_000; k++ {
				err := concrCounter.Add(k)
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})
}
