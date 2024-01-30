package flow

import (
	"context"
	"time"
)

type Procedure func() error

func Retry(ctx context.Context, fn Procedure, maxAttempts uint64, sleepOnFail bool) (err error) {
	var retries uint64
	for true {
		if retries >= maxAttempts {
			break
		}
		err = fn()
		if err != nil {
			if sleepOnFail {
				time.Sleep(time.Duration(1<<retries) * time.Second)
			}
			retries++
		} else {
			break
		}
	}
	return
}
