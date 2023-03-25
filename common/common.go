package common

import "golang.org/x/exp/constraints"

func Equal[T constraints.Ordered](self []T, other []T) bool {
	if len(self) != len(other) {
		return false
	}
	for _, element := range self {
		for _, otherElement := range other {
			if element != otherElement {
				return false
			}
		}
	}
	return true
}
