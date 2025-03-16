package core

import (
	"fmt"
	"sync"
	"time"
)

type worker struct {
	wg      *sync.WaitGroup
	consume <-chan string
	label   string
	r       bool
}

func newWorker(tenant string, ch <-chan string, wg *sync.WaitGroup) *worker {
	w := &worker{}
	wg.Add(1)
	w.wg = wg
	w.consume = ch
	w.label = tenant
	w.r = true
	return w
}

func (w *worker) run() {
	c := w.consume
	for i := 0; ; i++ {
		select {
		case msg := <-c:
			if w.r {
				fmt.Printf("%d tenant %s received message: %s\n", i, w.label, msg)
			}
		case <-time.After(time.Second * 1):
		default:
		}
		if !w.r {
			break
		}
	}
	w.wg.Done()
}

func (w *worker) closeWorker() {
	w.r = false
}

type businessWorkers struct {
	workers map[string]*worker
	lock    sync.Mutex
	wg      sync.WaitGroup
}

func newBusinessWorkers() *businessWorkers {
	bw := &businessWorkers{}
	bw.workers = make(map[string]*worker)
	bw.lock = sync.Mutex{}
	bw.wg = sync.WaitGroup{}
	return bw
}

func (bw *businessWorkers) listenerRegister(tenantChan <-chan string) {
	for {
		select {
		case msg := <-tenantChan:
			bw.lock.Lock()
			c := getChannel(msg)
			worker := newWorker(msg, c, &bw.wg)
			bw.workers[msg] = worker
			go worker.run()
			bw.lock.Unlock()
		default:
		}
	}
}

func (bw *businessWorkers) listenerUnregister(tenantChan <-chan string) {
	for {
		select {
		case msg := <-tenantChan:
			bw.lock.Lock()
			if c, ok := bw.workers[msg]; ok {
				delete(bw.workers, msg)
				c.closeWorker()
			}
			bw.lock.Unlock()
		default:
		}
	}
}
