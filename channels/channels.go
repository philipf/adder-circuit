package channels

// FanOut fanouts one channel into many channels
// https://stackoverflow.com/questions/16930251/go-one-producer-many-consumers
func FanOut(ch <-chan bool, num int) []chan bool {
	cs := make([]chan bool, num)

	for i := range cs {
		cs[i] = make(chan bool)
	}

	go func() {
		for i := range ch {
			for _, c := range cs {
				c <- i
			}
		}

		for _, c := range cs {
			// close all our fanOut channels when the input channel is exhausted.
			close(c)
		}
	}()

	return cs
}

// Split one channel into two
func Split(ch <-chan bool) (chan bool, chan bool) {
	r := FanOut(ch, 2)
	return r[0], r[1]
}
