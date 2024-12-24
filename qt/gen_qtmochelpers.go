package qt

import (
	"unsafe"
)

type QtMocHelpers__detail__TypeCompletenessForMetaType bool

const (
	QtMocHelpers__detail__TypeMayBeIncomplete QtMocHelpers__detail__TypeCompletenessForMetaType = false
	QtMocHelpers__detail__TypeMustBeComplete  QtMocHelpers__detail__TypeCompletenessForMetaType = true
)

type QtMocHelpers__NoType struct {
	h          uintptr
	isSubclass bool
}
