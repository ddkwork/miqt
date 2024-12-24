package qt

import (
	"unsafe"
)

type QLatin1StringMatcher struct {
	h          uintptr
	isSubclass bool
}

// NewQLatin1StringMatcher constructs a new QLatin1StringMatcher object.
func NewQLatin1StringMatcher() *QLatin1StringMatcher {
	ret := newQLatin1StringMatcher(QLatin1StringMatcher_new())
	ret.isSubclass = true
	return ret
}

func (this *QLatin1StringMatcher) SetCaseSensitivity(cs CaseSensitivity) {
	QLatin1StringMatcher_SetCaseSensitivity(this.h, (int)(cs))
}

func (this *QLatin1StringMatcher) CaseSensitivity() CaseSensitivity {
	return (CaseSensitivity)(QLatin1StringMatcher_CaseSensitivity(this.h))
}
