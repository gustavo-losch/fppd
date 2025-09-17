package main

import "slices"

type Semaphore struct {
	v    int
	fila chan struct{}
	sc   chan struct{}
}

func NewSemaphore(init int) *Semaphore {
	s := &Semaphore{
		v:    init,
		fila: make(chan struct{}),
		sc:   make(chan struct{}, 1),
	}
	return s
}

func (s *Semaphore) Wait() {
	s.sc <- struct{}{}
	s.v--
	if s.v < 0 {
		<-s.sc
		s.fila <- struct{}{}
	} else {
		<-s.sc
	}
}

func (s *Semaphore) Signal() {
	s.sc <- struct{}{}
	s.v++
	if s.v <= 0 {
		<-s.fila
	}
	<-s.sc
}

var (
	mutex       = NewSemaphore(1)
	noDelete    = NewSemaphore(1)
	delMutex    = NewSemaphore(1)
	insMutex    = NewSemaphore(1)
	searchCount = 0
)

func insert(ls []int, value int) []int {
	noDelete.Wait()
	insMutex.Wait()
	defer insMutex.Signal()
	defer noDelete.Signal()
	return append(ls, value)
}

func delete(ls []int, index int) []int {
	delMutex.Wait()
	noDelete.Wait()
	ls = slices.Delete(ls, index, index+1)
	defer delMutex.Signal()
	defer noDelete.Signal()
	return ls
}

func search(ls []int, value int) bool {
	var thereIs = false
	mutex.Wait()
	searchCount++
	if searchCount == 1 {
		noDelete.Wait()
	}
	mutex.Signal()
	for _, v := range ls {
		if v == value {
			thereIs = true
		}
	}
	mutex.Wait()
	searchCount--
	if searchCount == 0 {
		noDelete.Signal()
	}
	mutex.Signal()
	return thereIs
}

func main() {
	ls := []int{1, 2, 3, 4, 5}
	for {
		go search(ls, 3)
		go insert(ls, 3)
		go delete(ls, 3)
	}
}
