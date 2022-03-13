package count

import (
	"log"
	"sync"

	"gosub/caller"
)

type Count struct {
	receiveChannel chan int
	receiveClosed  bool
	accumulator    int
	called         int
	wg             sync.WaitGroup
}

// New - prepare accumulation environment
func New() *Count {
	var x Count
	x.receiveChannel = make(chan int, 10)
	x.receiveClosed = false
	go func() {
		for vNew := range x.receiveChannel {
			x.accumulator += vNew
			x.called++
		}
		x.wg.Done()
	}()
	return &x
}

// Value - get result of accumulation
func (x *Count) Value() int {
	if x.receiveClosed {
		return x.accumulator
	} else {
		log.Printf("count 'value' query before close.\n\t%v", caller.Caller())
		return -1
	}
}

// Called - get result of accumulation
func (x *Count) Called() int {
	if x.receiveClosed {
		return x.called
	} else {
		log.Printf("count 'called' query before close.\n\t%v", caller.Caller())
		return -1
	}
}

// Close - end of accumulation
func (x *Count) Close() {
	if !x.receiveClosed {
		// terminate accumulation
		x.wg.Add(1)
		x.receiveClosed = true
		close(x.receiveChannel)
	}
	x.wg.Wait()
}

// Add - accumulate value
func (x *Count) Add(value int) {
	defer func() {
		// negate error when add after close
		v := recover()
		if v != nil {
			log.Printf("count add after close\n\t%v", caller.Caller())
		}
	}()
	x.receiveChannel <- value
}
