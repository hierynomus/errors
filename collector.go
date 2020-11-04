package gotils

import (
	"fmt"
)

// ErrCollector is an error that contains an array of errors
type ErrCollector []error

// Collect appends an error to this ErrCollector if it is non-nil.
// if the passed error is a ErrCollector itself, the errors of that Collector will be appended.
func (c *ErrCollector) Collect(e error) {
	if e == nil {
		return
	}

	switch e := e.(type) {
	case *ErrCollector:
		if e.HasErrors() {
			for _, err := range *e { //nolint:gosimple (the simplification results in the Collector itself being appended, instead of appending its errors)
				*c = append(*c, err)
			}
		}
	default:
		*c = append(*c, e)
	}
}

// Error returns the error string containing all the collected errors
func (c *ErrCollector) Error() (err string) {
	err = "Collected errors:\n"
	for i, e := range *c {
		err += fmt.Sprintf("\tError %d: %s\n", i, e.Error())
	}

	return err
}

// HasErrors returns true when this ErrCollector contains 1 or more errors
func (c *ErrCollector) HasErrors() bool {
	return len(*c) != 0
}
