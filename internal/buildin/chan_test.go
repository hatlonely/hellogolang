package buildin

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestProducerConsumer(t *testing.T) {
	type Product struct {
		name  int
		value int
	}

	var wg sync.WaitGroup
	stop := false
	products := make(chan Product, 10)
	producer := func(name int) {
		for !stop {
			product := Product{name: name, value: rand.Int()}
			products <- product
			t.Logf("produce %v a product: %#v\n", name, product)
			time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
		}
		wg.Done()
	}
	consumer := func(name int) {
		for !stop && len(products) != 0 {
			product := <-products
			t.Logf("consume %v a product: %#v\n", name, product)
			time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
		}
		wg.Done()
	}

	for i := 0; i < 5; i++ {
		go producer(i)
		go consumer(i)
		wg.Add(2)
	}

	time.Sleep(time.Duration(1) * time.Second)
	stop = true

	wg.Wait()
}

func TestSelect(t *testing.T) {
	type ProductA struct {
		name  int
		value int
	}

	type ProductB struct {
		name  int
		value int
	}

	productAs := make(chan ProductA, 10)
	productBs := make(chan ProductB, 10)

	stopA := false
	stopB := false
	var wg sync.WaitGroup

	producerA := func(name int) {
		for !stopA {
			product := ProductA{name: name, value: rand.Int()}
			productAs <- product
			t.Logf("produceA %v a product: %#v\n", name, product)
			time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
		}
		wg.Done()
	}
	producerB := func(name int) {
		for !stopB {
			product := ProductB{name: name, value: rand.Int()}
			productBs <- product
			t.Logf("produceB %v a product: %#v\n", name, product)
			time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
		}
		wg.Done()
	}
	consumer := func(name int) {
		ticker := time.Tick(time.Duration(500) * time.Millisecond)
		for !stopA || !stopB || len(productAs) != 0 || len(productBs) != 0 {
			select {
			case product := <-productAs:
				t.Logf("consume %v a productA: %#v\n", name, product)
			case product := <-productBs:
				t.Logf("consume %v a productB: %#v\n", name, product)
			case <-ticker:
				// nothing to do just awake from block
			}
		}
		wg.Done()
	}

	for i := 0; i < 5; i++ {
		go producerA(i)
		go producerB(i)
		go consumer(i)
		wg.Add(3)
	}

	time.Sleep(time.Duration(1) * time.Second)
	stopA = true
	time.Sleep(time.Duration(1) * time.Second)
	stopB = true

	wg.Wait()
}
