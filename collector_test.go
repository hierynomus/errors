package gotils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyCollectorHasNoErrors(t *testing.T) {
	collector := new(ErrCollector)

	assert.False(t, collector.HasErrors())
}

func TestCollectorWithErrorHasErrors(t *testing.T) {
	collector := new(ErrCollector)
	collector.Collect(fmt.Errorf("an error"))
	assert.True(t, collector.HasErrors())
}

func TestShouldFlattenNestedCollectors(t *testing.T) {
	collector := new(ErrCollector)
	nested := new(ErrCollector)

	collector.Collect(fmt.Errorf("an error"))
	nested.Collect(fmt.Errorf("a nested error"))

	collector.Collect(nested)

	assert.Error(t, collector)
	assert.Contains(t, collector.Error(), "a nested error")
}

func TestShouldNotCollectNilError(t *testing.T) {
	collector := new(ErrCollector)
	collector.Collect(nil)

	assert.False(t, collector.HasErrors())
}

func TestShouldNotCollectEmptyNestedCollector(t *testing.T) {
	collector := new(ErrCollector)
	nested := new(ErrCollector)

	collector.Collect(nested)

	assert.False(t, collector.HasErrors())
}
