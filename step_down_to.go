package lib

// StepDownTo is used with for...range, e.g.:
//
//	for n := range StepDownTo(10,0,2) {
//			println(n)
//	}
func StepDownTo(start, end, step int64) chan int64 {
	ch := make(chan int64)
	go func() {
		var n, x int64
		for n = start; n > end; n -= step {
			ch <- n
			x = n
		}
		if x-step == end {
			ch <- end
		}
		close(ch)
	}()
	return ch
}
