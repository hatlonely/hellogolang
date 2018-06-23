package grpclb

import (
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
	return &pseudoWatcher{
		addrs:    r.addrs,
		hashNext: true,
	}, nil
}

type pseudoWatcher struct {
	addrs    []string
	hashNext bool
}

func (w *pseudoWatcher) Next() ([]*naming.Update, error) {
	if !w.hashNext {
		// block forever
		<-make(chan bool)
	}

	var updates []*naming.Update
	for _, addr := range w.addrs {
		updates = append(updates, &naming.Update{Op: naming.Add, Addr: addr})
	}

	return updates, nil
}

func (w *pseudoWatcher) Close() {
	// nothing to do
}
