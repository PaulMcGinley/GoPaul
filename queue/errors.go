package queue

import (
	"errors"
)

var IsEmptyError = errors.New("queue is empty")
var DoesNotContainError = errors.New("queue does not contain element")
var DoesNotHaveError = errors.New("queue does not have n elements")
