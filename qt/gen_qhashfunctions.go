package qt

import (
	"unsafe"
)

type QHashSeed struct {
	h          uintptr
	isSubclass bool
}

// NewQHashSeed constructs a new QHashSeed object.
func NewQHashSeed() *QHashSeed {
	ret := newQHashSeed(QHashSeed_new())
	ret.isSubclass = true
	return ret
}

// NewQHashSeed2 constructs a new QHashSeed object.
func NewQHashSeed2(d uint64) *QHashSeed {
	ret := newQHashSeed(QHashSeed_new2((ulonglong)(d)))
	ret.isSubclass = true
	return ret
}

func QHashSeed_GlobalSeed() *QHashSeed {
	_goptr := newQHashSeed(QHashSeed_GlobalSeed())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QHashSeed_SetDeterministicGlobalSeed() {
	QHashSeed_SetDeterministicGlobalSeed()
}

func QHashSeed_ResetRandomGlobalSeed() {
	QHashSeed_ResetRandomGlobalSeed()
}
