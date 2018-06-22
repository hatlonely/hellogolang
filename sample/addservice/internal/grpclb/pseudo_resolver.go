package grpclb

// https://github.com/hakobe/grpc-go-client-side-load-balancing-example/blob/master/client/resolver.go

import (
	"fmt"

	"google.golang.org/grpc/naming"
)

// NewPseudoResolver creates a new pseudo resolver which returns fixed addrs.
func NewPseudoResolver(addrs []string) naming.Resolver {
	return &pseudoResolver{addrs}
}

type pseudoResolver struct {
	addrs []string
}

func (r *pseudoResolver) Resolve(target string) (naming.Watcher, error) {
	w := &pseudoWatcher{
		updatesChan: make(chan []*naming.Update, 1),
	}
	updates := []*naming.Update{}
	for _, addr := range r.addrs {
		updates = append(updates, &naming.Update{Op: naming.Add, Addr: addr})
	}
	w.updatesChan <- updates
	return w, nil
}

// This watcher is implemented based on ipwatcher below
// https://github.com/grpc/grpc-go/blob/30fb59a4304034ce78ff68e21bd25776b1d79488/naming/dns_resolver.go#L151-L171
type pseudoWatcher struct {
	updatesChan chan []*naming.Update
}

func (w *pseudoWatcher) Next() ([]*naming.Update, error) {
	us, ok := <-w.updatesChan
	if !ok {
		return nil, fmt.Errorf("watcher has been closed")
	}
	return us, nil
}

func (w *pseudoWatcher) Close() {
	close(w.updatesChan)
}
