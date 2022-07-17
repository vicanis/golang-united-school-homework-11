package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) []user {
	wg := sync.WaitGroup{}

	wg.Add(int(n))

	result := make([]user, n)

	workers := make(chan struct{}, pool)

	for i := 0; int64(i) < n; i++ {
		go func(i int, res []user) {
			workers <- struct{}{}

			res[i] = getOne(int64(i))

			<-workers

			wg.Done()
		}(i, result)
	}

	wg.Wait()

	return result
}
