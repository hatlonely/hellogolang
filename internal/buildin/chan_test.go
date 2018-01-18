package buildin

import (
	"testing"
	"time"
	"math/rand"
)

func TestProducerConsumer(t *testing.T) {
	type Product struct {
		name  int
		value int
	}
	products := make(chan Product, 10)
	producer := func(name int) {
		for {
			product := Product{name: name, value: rand.Int()}
			products <- product
			t.Logf("produce %v a product: %#v\n", name, product)
			time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
		}
	}
	consumer := func(name int) {
		for {
			product := <-products
			t.Logf("consume %v a product: %#v\n", name, product)
			time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
		}
	}

	for i := 0; i < 5; i++ {
		go producer(i)
		go consumer(i)
	}

	time.Sleep(time.Duration(1) * time.Second)
}
