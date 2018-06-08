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

// Return []error.
func (ee *errs) All() []error {
	return ee.all
}

// Return all errors as a single error.
func (ee *errs) ToError() error {
	if ee.IsEmpty() {
		return nil
	}

	return ee
}

// Return true if there are no errors.
func (ee *errs) IsEmpty() bool {
	return len(ee.all) == 0
}

// Implement the error interface for errs.
func (ee *errs) Error() string {
	errorStr := ""
	for _, err := range ee.All() {
		errorStr = fmt.Sprintf("%s\n%s", errorStr, err.Error())
	}

	return errorStr
}

// Safely append to the []error in errs struct.
func (ee *errs) append(err error) {
	ee.mutex.Lock()
	defer ee.mutex.Unlock()
	ee.all = append(ee.all, err)
}
