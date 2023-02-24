package goroutine

import (
	"fmt"
	"sync"
)

// StorageMain 通过 struct 实现类似本地存储
func StorageMain() {

	var wg sync.WaitGroup

	var ts [5]struct{
		id,result 		int
	}

	for i:=0; i<len(ts); i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			ts[x].id = x
			ts[x].result = (x + 1) * 100
		}(i)
	}

	wg.Wait()

	fmt.Printf("the ts is: %+v:\n",ts)
}



