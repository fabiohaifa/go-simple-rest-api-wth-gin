// Goal is 95% code coveraged by tests

package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getToDos(t *testing.T) {
	arrToDos := T.GetAllToDos(0, "")
	assert.NotNil(t, arrToDos, "ToDos should not be nil")
	assert.Equal(t, len(arrToDos), 3, "ToDos should have 3 items")
}

func Test_getOneToDos(t *testing.T) {
	arrToDos := T.GetAllToDos(0, "Learn Gin")
	assert.NotNil(t, arrToDos, "ToDos should not be nil")
	assert.Equal(t, len(arrToDos), 1, "ToDos should have 1 items")
}
