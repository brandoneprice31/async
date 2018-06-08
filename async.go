package async

import (
	"sync"
)

// Run functions in parallel and collect the errors.
func Parallel(ff ...func() error) Errors {
	var wg sync.WaitGroup
	var ee errs

	wg.Add(len(ff))

	for i := range ff {
		f := ff[i]

		go func() {
			defer wg.Done()

			if err := f(); err != nil {
				ee.append(err)
			}
		}()
	}

	wg.Wait()

	return &ee
}
