package async

import (
	"fmt"
	"sync"
)

type (
	Errors interface {
		All() []error
		ToError() error
		IsEmpty() bool
	}

	errs struct {
		mutex sync.Mutex
		all   []error
	}
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

// Return []error.
func (ee *errs) All() []error {
	return ee.all
}

// Return all errors as a single error.
func (ee *errs) ToError() error {
	return ee
}

// Return true if there are no errors.
func (ee *errs) IsEmpty() bool {
	return len(ee.all) == 0
}

// Implement the error interface for errors.
func (ee *errs) Error() string {
	errorStr := ""
	for _, err := range ee.All() {
		errorStr = fmt.Sprintf("%s\n%s", errorStr, err.Error())
	}

	return errorStr
}

// Safely append to the []error in errors struct.
func (ee *errs) append(err error) {
	ee.mutex.Lock()
	defer ee.mutex.Unlock()
	ee.all = append(ee.all, err)
}
