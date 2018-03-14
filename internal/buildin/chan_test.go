package buildin

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Product struct {
	name  int
	value int
}

func producer(wg *sync.WaitGroup, products chan<- Product, name int, stop *bool) {
	for !*stop {
		product := Product{name: name, value: rand.Int()}
		products <- product
		fmt.Printf("producer %v produce a product: %#v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, products <-chan Product, name int) {
	for product := range products {
		fmt.Printf("consumer %v consume a product: %#v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}

func TestProducerConsumer(t *testing.T) {
	var wgp sync.WaitGroup
	var wgc sync.WaitGroup
	stop := false
	products := make(chan Product, 10)
	for i := 0; i < 5; i++ {
		go producer(&wgp, products, i, &stop)
		go consumer(&wgc, products, i)
		wgp.Add(1)
		wgc.Add(1)
	}

	time.Sleep(time.Duration(1) * time.Second)
	stop = true
	wgp.Wait()
	close(products)
	wgc.Wait()
}

type ProductA struct {
	name  int
	value int
}

type ProductB struct {
	name  int
	value int
}

func producerA(wg *sync.WaitGroup, productAs chan<- ProductA, name int, stop *bool) {
	for !*stop {
		product := ProductA{name: name, value: rand.Int()}
		productAs <- product
		fmt.Printf("producerA %v produce a productA: %#v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}

func producerB(wg *sync.WaitGroup, productBs chan<- ProductB, name int, stop *bool) {
	for !*stop {
		product := ProductB{name: name, value: rand.Int()}
		productBs <- product
		fmt.Printf("producerB %v produce a productB: %#v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}

func consumerAB(wg *sync.WaitGroup, productAs <-chan ProductA, productBs <-chan ProductB, name int, stopA *bool, stopB *bool) {
	ticker := time.Tick(time.Duration(500) * time.Millisecond)
	for !*stopA || !*stopB || len(productAs) != 0 || len(productBs) != 0 {
		select {
		case product := <-productAs:
			fmt.Printf("consumerAB %v consume a productA: %#v\n", name, product)
		case product := <-productBs:
			fmt.Printf("consumerAB %v consume a productB: %#v\n", name, product)
		case <-ticker:
			// nothing to do just awake from block
		}
	}
	wg.Done()
}

func TestSelect(t *testing.T) {
	productAs := make(chan ProductA, 10)
	productBs := make(chan ProductB, 10)

	stopA := false
	stopB := false
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go producerA(&wg, productAs, i, &stopA)
		go producerB(&wg, productBs, i, &stopB)
		go consumerAB(&wg, productAs, productBs, i, &stopA, &stopB)
		wg.Add(3)
	}

	time.Sleep(time.Duration(1) * time.Second)
	stopA = true
	time.Sleep(time.Duration(1) * time.Second)
	stopB = true

	wg.Wait()
}
