package count

import (
	"fmt"
	"sync"
	"testing"
)

func TestCount_Add(t *testing.T) {

	var wg sync.WaitGroup
	x := New()

	for ix := 1; ix <= 1000000; ix++ {
		wg.Add(1)
		go func(i int) {
			x.Add(i)
			wg.Done()
		}(ix * 2)
	}

	fmt.Printf("%+v\n", x.Value())
	fmt.Printf("%+v\n", x.Called())
	wg.Wait()
	x.Close()
	fmt.Printf("%+v\n", x.Value())
	fmt.Printf("%+v\n", x.Called())

	x.Add(1000)
}
