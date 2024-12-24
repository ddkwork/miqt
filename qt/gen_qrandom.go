package qt

import (
	"unsafe"
)

type QRandomGenerator struct {
	h          uintptr
	isSubclass bool
}

// NewQRandomGenerator constructs a new QRandomGenerator object.
func NewQRandomGenerator() *QRandomGenerator {
	g := newQRandomGenerator(QRandomGenerator_new())
	g.isSubclass = true
	return g
}

// NewQRandomGenerator2 constructs a new QRandomGenerator object.
func NewQRandomGenerator2(seedBuffer *uint, lenVal int64) *QRandomGenerator {
	g := newQRandomGenerator(QRandomGenerator_new2((*uint)(unsafe.Pointer(seedBuffer)), (ptrdiff_t)(lenVal)))
	g.isSubclass = true
	return g
}

// NewQRandomGenerator3 constructs a new QRandomGenerator object.
func NewQRandomGenerator3(begin *uint, end *uint) *QRandomGenerator {
	g := newQRandomGenerator(QRandomGenerator_new3((*uint)(unsafe.Pointer(begin)), (*uint)(unsafe.Pointer(end))))
	g.isSubclass = true
	return g
}

// NewQRandomGenerator4 constructs a new QRandomGenerator object.
func NewQRandomGenerator4(other *QRandomGenerator) *QRandomGenerator {
	g := newQRandomGenerator(QRandomGenerator_new4(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQRandomGenerator5 constructs a new QRandomGenerator object.
func NewQRandomGenerator5(seedValue uint) *QRandomGenerator {
	g := newQRandomGenerator(QRandomGenerator_new5((uint)(seedValue)))
	g.isSubclass = true
	return g
}

func (this *QRandomGenerator) OperatorAssign(other *QRandomGenerator) {
	QRandomGenerator_OperatorAssign(this.h, other.cPointer())
}

func (this *QRandomGenerator) Generate() uint {
	return (uint)(QRandomGenerator_Generate(this.h))
}

func (this *QRandomGenerator) Generate64() uint64 {
	return (uint64)(QRandomGenerator_Generate64(this.h))
}

func (this *QRandomGenerator) GenerateDouble() float64 {
	return (float64)(QRandomGenerator_GenerateDouble(this.h))
}

func (this *QRandomGenerator) Bounded(highest float64) float64 {
	return (float64)(QRandomGenerator_Bounded(this.h, (double)(highest)))
}

func (this *QRandomGenerator) BoundedWithHighest(highest uint) uint {
	return (uint)(QRandomGenerator_BoundedWithHighest(this.h, (uint)(highest)))
}

func (this *QRandomGenerator) Bounded2(lowest uint, highest uint) uint {
	return (uint)(QRandomGenerator_Bounded2(this.h, (uint)(lowest), (uint)(highest)))
}

func (this *QRandomGenerator) Bounded3(highest int) int {
	return (int)(QRandomGenerator_Bounded3(this.h, (int)(highest)))
}

func (this *QRandomGenerator) Bounded4(lowest int, highest int) int {
	return (int)(QRandomGenerator_Bounded4(this.h, (int)(lowest), (int)(highest)))
}

func (this *QRandomGenerator) Bounded5(highest uint64) uint64 {
	return (uint64)(QRandomGenerator_Bounded5(this.h, (ulonglong)(highest)))
}

func (this *QRandomGenerator) Bounded6(lowest uint64, highest uint64) uint64 {
	return (uint64)(QRandomGenerator_Bounded6(this.h, (ulonglong)(lowest), (ulonglong)(highest)))
}

func (this *QRandomGenerator) Bounded7(highest int64) int64 {
	return (int64)(QRandomGenerator_Bounded7(this.h, (longlong)(highest)))
}

func (this *QRandomGenerator) Bounded8(lowest int64, highest int64) int64 {
	return (int64)(QRandomGenerator_Bounded8(this.h, (longlong)(lowest), (longlong)(highest)))
}

func (this *QRandomGenerator) Bounded9(lowest int, highest int64) int64 {
	return (int64)(QRandomGenerator_Bounded9(this.h, (int)(lowest), (longlong)(highest)))
}

func (this *QRandomGenerator) Bounded10(lowest int64, highest int) int64 {
	return (int64)(QRandomGenerator_Bounded10(this.h, (longlong)(lowest), (int)(highest)))
}

func (this *QRandomGenerator) Bounded11(lowest uint, highest uint64) uint64 {
	return (uint64)(QRandomGenerator_Bounded11(this.h, (uint)(lowest), (ulonglong)(highest)))
}

func (this *QRandomGenerator) Bounded12(lowest uint64, highest uint) uint64 {
	return (uint64)(QRandomGenerator_Bounded12(this.h, (ulonglong)(lowest), (uint)(highest)))
}

func (this *QRandomGenerator) Generate2(begin *uint, end *uint) {
	QRandomGenerator_Generate2(this.h, (*uint)(unsafe.Pointer(begin)), (*uint)(unsafe.Pointer(end)))
}

func (this *QRandomGenerator) OperatorCall() result_type {
	xxxxxxxxx
}

func (this *QRandomGenerator) Seed() {
	QRandomGenerator_Seed(this.h)
}

func (this *QRandomGenerator) Discard(z uint64) {
	QRandomGenerator_Discard(this.h, (ulonglong)(z))
}

func QRandomGenerator_Min() result_type {
	xxxxxxxxx
}

func QRandomGenerator_Max() result_type {
	xxxxxxxxx
}

func QRandomGenerator_System() *QRandomGenerator {
	return newQRandomGenerator(QRandomGenerator_System())
}

func QRandomGenerator_Global() *QRandomGenerator {
	return newQRandomGenerator(QRandomGenerator_Global())
}

func QRandomGenerator_SecurelySeeded() *QRandomGenerator {
	_goptr := newQRandomGenerator(QRandomGenerator_SecurelySeeded())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRandomGenerator) Seed1(s uint) {
	QRandomGenerator_Seed1(this.h, (uint)(s))
}

type QRandomGenerator64 struct {
	h          uintptr
	isSubclass bool
}

// NewQRandomGenerator64 constructs a new QRandomGenerator64 object.
func NewQRandomGenerator64() *QRandomGenerator64 {
	g := newQRandomGenerator64(QRandomGenerator64_new())
	g.isSubclass = true
	return g
}

// NewQRandomGenerator642 constructs a new QRandomGenerator64 object.
func NewQRandomGenerator642(seedBuffer *uint, lenVal int64) *QRandomGenerator64 {
	g := newQRandomGenerator64(QRandomGenerator64_new2((*uint)(unsafe.Pointer(seedBuffer)), (ptrdiff_t)(lenVal)))
	g.isSubclass = true
	return g
}

// NewQRandomGenerator643 constructs a new QRandomGenerator64 object.
func NewQRandomGenerator643(begin *uint, end *uint) *QRandomGenerator64 {
	g := newQRandomGenerator64(QRandomGenerator64_new3((*uint)(unsafe.Pointer(begin)), (*uint)(unsafe.Pointer(end))))
	g.isSubclass = true
	return g
}

// NewQRandomGenerator644 constructs a new QRandomGenerator64 object.
func NewQRandomGenerator644(other *QRandomGenerator) *QRandomGenerator64 {
	g := newQRandomGenerator64(QRandomGenerator64_new4(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQRandomGenerator645 constructs a new QRandomGenerator64 object.
func NewQRandomGenerator645(param1 *QRandomGenerator64) *QRandomGenerator64 {
	g := newQRandomGenerator64(QRandomGenerator64_new5(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQRandomGenerator646 constructs a new QRandomGenerator64 object.
func NewQRandomGenerator646(seedValue uint) *QRandomGenerator64 {
	g := newQRandomGenerator64(QRandomGenerator64_new6((uint)(seedValue)))
	g.isSubclass = true
	return g
}

func (this *QRandomGenerator64) Generate() uint64 {
	return (uint64)(QRandomGenerator64_Generate(this.h))
}

func (this *QRandomGenerator64) OperatorCall() result_type {
	xxxxxxxxx
}

func (this *QRandomGenerator64) Discard(z uint64) {
	QRandomGenerator64_Discard(this.h, (ulonglong)(z))
}

func QRandomGenerator64_Min() result_type {
	xxxxxxxxx
}

func QRandomGenerator64_Max() result_type {
	xxxxxxxxx
}

func QRandomGenerator64_System() *QRandomGenerator64 {
	return newQRandomGenerator64(QRandomGenerator64_System())
}

func QRandomGenerator64_Global() *QRandomGenerator64 {
	return newQRandomGenerator64(QRandomGenerator64_Global())
}

func QRandomGenerator64_SecurelySeeded() *QRandomGenerator64 {
	_goptr := newQRandomGenerator64(QRandomGenerator64_SecurelySeeded())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRandomGenerator64) OperatorAssign(param1 *QRandomGenerator64) {
	QRandomGenerator64_OperatorAssign(this.h, param1.cPointer())
}
