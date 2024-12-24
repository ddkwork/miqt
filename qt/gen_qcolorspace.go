package qt

import (
	"unsafe"
)

type QColorSpace__NamedColorSpace int

const (
	QColorSpace__NamedColorSpace__SRgb        QColorSpace__NamedColorSpace = 1
	QColorSpace__NamedColorSpace__SRgbLinear  QColorSpace__NamedColorSpace = 2
	QColorSpace__NamedColorSpace__AdobeRgb    QColorSpace__NamedColorSpace = 3
	QColorSpace__NamedColorSpace__DisplayP3   QColorSpace__NamedColorSpace = 4
	QColorSpace__NamedColorSpace__ProPhotoRgb QColorSpace__NamedColorSpace = 5
	QColorSpace__NamedColorSpace__Bt2020      QColorSpace__NamedColorSpace = 6
	QColorSpace__NamedColorSpace__Bt2100Pq    QColorSpace__NamedColorSpace = 7
	QColorSpace__NamedColorSpace__Bt2100Hlg   QColorSpace__NamedColorSpace = 8
)

type QColorSpace__Primaries int

const (
	QColorSpace__Primaries__Custom      QColorSpace__Primaries = 0
	QColorSpace__Primaries__SRgb        QColorSpace__Primaries = 1
	QColorSpace__Primaries__AdobeRgb    QColorSpace__Primaries = 2
	QColorSpace__Primaries__DciP3D65    QColorSpace__Primaries = 3
	QColorSpace__Primaries__ProPhotoRgb QColorSpace__Primaries = 4
	QColorSpace__Primaries__Bt2020      QColorSpace__Primaries = 5
)

type QColorSpace__TransferFunction int

const (
	QColorSpace__TransferFunction__Custom      QColorSpace__TransferFunction = 0
	QColorSpace__TransferFunction__Linear      QColorSpace__TransferFunction = 1
	QColorSpace__TransferFunction__Gamma       QColorSpace__TransferFunction = 2
	QColorSpace__TransferFunction__SRgb        QColorSpace__TransferFunction = 3
	QColorSpace__TransferFunction__ProPhotoRgb QColorSpace__TransferFunction = 4
	QColorSpace__TransferFunction__Bt2020      QColorSpace__TransferFunction = 5
	QColorSpace__TransferFunction__St2084      QColorSpace__TransferFunction = 6
	QColorSpace__TransferFunction__Hlg         QColorSpace__TransferFunction = 7
)

type QColorSpace__TransformModel uint8_t

const (
	QColorSpace__ThreeComponentMatrix  QColorSpace__TransformModel = 0
	QColorSpace__ElementListProcessing QColorSpace__TransformModel = 1
)

type QColorSpace__ColorModel uint8_t

const (
	QColorSpace__Undefined QColorSpace__ColorModel = 0
	QColorSpace__Rgb       QColorSpace__ColorModel = 1
	QColorSpace__Gray      QColorSpace__ColorModel = 2
	QColorSpace__Cmyk      QColorSpace__ColorModel = 3
)

type QColorSpace struct {
	h          uintptr
	isSubclass bool
}

// NewQColorSpace constructs a new QColorSpace object.
func NewQColorSpace() *QColorSpace {
	g := newQColorSpace(QColorSpace_new())
	g.isSubclass = true
	return g
}

// NewQColorSpace2 constructs a new QColorSpace object.
func NewQColorSpace2(namedColorSpace NamedColorSpace) *QColorSpace {
	g := newQColorSpace(QColorSpace_new2(namedColorSpace))
	g.isSubclass = true
	return g
}

// NewQColorSpace3 constructs a new QColorSpace object.
func NewQColorSpace3(whitePoint QPointF, transferFunction TransferFunction) *QColorSpace {
	g := newQColorSpace(QColorSpace_new3(whitePoint.cPointer(), transferFunction))
	g.isSubclass = true
	return g
}

// NewQColorSpace4 constructs a new QColorSpace object.
func NewQColorSpace4(whitePoint QPointF, transferFunctionTable []uint16) *QColorSpace {
	transferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(transferFunctionTable))))
	defer free(unsafe.Pointer(transferFunctionTable_CArray))
	for i := range transferFunctionTable {
		transferFunctionTable_CArray[i] = (uint16_t)(transferFunctionTable[i])
	}
	transferFunctionTable_ma := struct_miqt_array{len: size_t(len(transferFunctionTable)), data: unsafe.Pointer(transferFunctionTable_CArray)}
	g := newQColorSpace(QColorSpace_new4(whitePoint.cPointer(), transferFunctionTable_ma))
	g.isSubclass = true
	return g
}

// NewQColorSpace5 constructs a new QColorSpace object.
func NewQColorSpace5(primaries Primaries, transferFunction TransferFunction) *QColorSpace {
	g := newQColorSpace(QColorSpace_new5(primaries, transferFunction))
	g.isSubclass = true
	return g
}

// NewQColorSpace6 constructs a new QColorSpace object.
func NewQColorSpace6(primaries Primaries, gamma float32) *QColorSpace {
	g := newQColorSpace(QColorSpace_new6(primaries, (float)(gamma)))
	g.isSubclass = true
	return g
}

// NewQColorSpace7 constructs a new QColorSpace object.
func NewQColorSpace7(primaries Primaries, transferFunctionTable []uint16) *QColorSpace {
	transferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(transferFunctionTable))))
	defer free(unsafe.Pointer(transferFunctionTable_CArray))
	for i := range transferFunctionTable {
		transferFunctionTable_CArray[i] = (uint16_t)(transferFunctionTable[i])
	}
	transferFunctionTable_ma := struct_miqt_array{len: size_t(len(transferFunctionTable)), data: unsafe.Pointer(transferFunctionTable_CArray)}
	g := newQColorSpace(QColorSpace_new7(primaries, transferFunctionTable_ma))
	g.isSubclass = true
	return g
}

// NewQColorSpace8 constructs a new QColorSpace object.
func NewQColorSpace8(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF, transferFunction TransferFunction) *QColorSpace {
	g := newQColorSpace(QColorSpace_new8(whitePoint.cPointer(), redPoint.cPointer(), greenPoint.cPointer(), bluePoint.cPointer(), transferFunction))
	g.isSubclass = true
	return g
}

// NewQColorSpace9 constructs a new QColorSpace object.
func NewQColorSpace9(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF, transferFunctionTable []uint16) *QColorSpace {
	transferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(transferFunctionTable))))
	defer free(unsafe.Pointer(transferFunctionTable_CArray))
	for i := range transferFunctionTable {
		transferFunctionTable_CArray[i] = (uint16_t)(transferFunctionTable[i])
	}
	transferFunctionTable_ma := struct_miqt_array{len: size_t(len(transferFunctionTable)), data: unsafe.Pointer(transferFunctionTable_CArray)}
	g := newQColorSpace(QColorSpace_new9(whitePoint.cPointer(), redPoint.cPointer(), greenPoint.cPointer(), bluePoint.cPointer(), transferFunctionTable_ma))
	g.isSubclass = true
	return g
}

// NewQColorSpace10 constructs a new QColorSpace object.
func NewQColorSpace10(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF, redTransferFunctionTable []uint16, greenTransferFunctionTable []uint16, blueTransferFunctionTable []uint16) *QColorSpace {
	redTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(redTransferFunctionTable))))
	defer free(unsafe.Pointer(redTransferFunctionTable_CArray))
	for i := range redTransferFunctionTable {
		redTransferFunctionTable_CArray[i] = (uint16_t)(redTransferFunctionTable[i])
	}
	redTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(redTransferFunctionTable)), data: unsafe.Pointer(redTransferFunctionTable_CArray)}
	greenTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(greenTransferFunctionTable))))
	defer free(unsafe.Pointer(greenTransferFunctionTable_CArray))
	for i := range greenTransferFunctionTable {
		greenTransferFunctionTable_CArray[i] = (uint16_t)(greenTransferFunctionTable[i])
	}
	greenTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(greenTransferFunctionTable)), data: unsafe.Pointer(greenTransferFunctionTable_CArray)}
	blueTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(blueTransferFunctionTable))))
	defer free(unsafe.Pointer(blueTransferFunctionTable_CArray))
	for i := range blueTransferFunctionTable {
		blueTransferFunctionTable_CArray[i] = (uint16_t)(blueTransferFunctionTable[i])
	}
	blueTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(blueTransferFunctionTable)), data: unsafe.Pointer(blueTransferFunctionTable_CArray)}
	g := newQColorSpace(QColorSpace_new10(whitePoint.cPointer(), redPoint.cPointer(), greenPoint.cPointer(), bluePoint.cPointer(), redTransferFunctionTable_ma, greenTransferFunctionTable_ma, blueTransferFunctionTable_ma))
	g.isSubclass = true
	return g
}

// NewQColorSpace11 constructs a new QColorSpace object.
func NewQColorSpace11(colorSpace *QColorSpace) *QColorSpace {
	g := newQColorSpace(QColorSpace_new11(colorSpace.cPointer()))
	g.isSubclass = true
	return g
}

// NewQColorSpace12 constructs a new QColorSpace object.
func NewQColorSpace12(whitePoint QPointF, transferFunction TransferFunction, gamma float32) *QColorSpace {
	g := newQColorSpace(QColorSpace_new12(whitePoint.cPointer(), transferFunction, (float)(gamma)))
	g.isSubclass = true
	return g
}

// NewQColorSpace13 constructs a new QColorSpace object.
func NewQColorSpace13(primaries Primaries, transferFunction TransferFunction, gamma float32) *QColorSpace {
	g := newQColorSpace(QColorSpace_new13(primaries, transferFunction, (float)(gamma)))
	g.isSubclass = true
	return g
}

// NewQColorSpace14 constructs a new QColorSpace object.
func NewQColorSpace14(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF, transferFunction TransferFunction, gamma float32) *QColorSpace {
	g := newQColorSpace(QColorSpace_new14(whitePoint.cPointer(), redPoint.cPointer(), greenPoint.cPointer(), bluePoint.cPointer(), transferFunction, (float)(gamma)))
	g.isSubclass = true
	return g
}

func (this *QColorSpace) OperatorAssign(colorSpace *QColorSpace) {
	QColorSpace_OperatorAssign(this.h, colorSpace.cPointer())
}

func (this *QColorSpace) Swap(colorSpace *QColorSpace) {
	QColorSpace_Swap(this.h, colorSpace.cPointer())
}

func (this *QColorSpace) Primaries() Primaries {
	xxxxxxxxx
}

func (this *QColorSpace) TransferFunction() TransferFunction {
	xxxxxxxxx
}

func (this *QColorSpace) Gamma() float32 {
	return (float32)(QColorSpace_Gamma(this.h))
}

func (this *QColorSpace) Description() string {
	var _ms struct_miqt_string = QColorSpace_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QColorSpace) SetDescription(description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QColorSpace_SetDescription(this.h, description_ms)
}

func (this *QColorSpace) SetTransferFunction(transferFunction TransferFunction) {
	QColorSpace_SetTransferFunction(this.h, transferFunction)
}

func (this *QColorSpace) SetTransferFunctionWithTransferFunctionTable(transferFunctionTable []uint16) {
	transferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(transferFunctionTable))))
	defer free(unsafe.Pointer(transferFunctionTable_CArray))
	for i := range transferFunctionTable {
		transferFunctionTable_CArray[i] = (uint16_t)(transferFunctionTable[i])
	}
	transferFunctionTable_ma := struct_miqt_array{len: size_t(len(transferFunctionTable)), data: unsafe.Pointer(transferFunctionTable_CArray)}
	QColorSpace_SetTransferFunctionWithTransferFunctionTable(this.h, transferFunctionTable_ma)
}

func (this *QColorSpace) SetTransferFunctions(redTransferFunctionTable []uint16, greenTransferFunctionTable []uint16, blueTransferFunctionTable []uint16) {
	redTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(redTransferFunctionTable))))
	defer free(unsafe.Pointer(redTransferFunctionTable_CArray))
	for i := range redTransferFunctionTable {
		redTransferFunctionTable_CArray[i] = (uint16_t)(redTransferFunctionTable[i])
	}
	redTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(redTransferFunctionTable)), data: unsafe.Pointer(redTransferFunctionTable_CArray)}
	greenTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(greenTransferFunctionTable))))
	defer free(unsafe.Pointer(greenTransferFunctionTable_CArray))
	for i := range greenTransferFunctionTable {
		greenTransferFunctionTable_CArray[i] = (uint16_t)(greenTransferFunctionTable[i])
	}
	greenTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(greenTransferFunctionTable)), data: unsafe.Pointer(greenTransferFunctionTable_CArray)}
	blueTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(blueTransferFunctionTable))))
	defer free(unsafe.Pointer(blueTransferFunctionTable_CArray))
	for i := range blueTransferFunctionTable {
		blueTransferFunctionTable_CArray[i] = (uint16_t)(blueTransferFunctionTable[i])
	}
	blueTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(blueTransferFunctionTable)), data: unsafe.Pointer(blueTransferFunctionTable_CArray)}
	QColorSpace_SetTransferFunctions(this.h, redTransferFunctionTable_ma, greenTransferFunctionTable_ma, blueTransferFunctionTable_ma)
}

func (this *QColorSpace) WithTransferFunction(transferFunction TransferFunction) *QColorSpace {
	_goptr := newQColorSpace(QColorSpace_WithTransferFunction(this.h, transferFunction))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorSpace) WithTransferFunctionWithTransferFunctionTable(transferFunctionTable []uint16) *QColorSpace {
	transferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(transferFunctionTable))))
	defer free(unsafe.Pointer(transferFunctionTable_CArray))
	for i := range transferFunctionTable {
		transferFunctionTable_CArray[i] = (uint16_t)(transferFunctionTable[i])
	}
	transferFunctionTable_ma := struct_miqt_array{len: size_t(len(transferFunctionTable)), data: unsafe.Pointer(transferFunctionTable_CArray)}
	_goptr := newQColorSpace(QColorSpace_WithTransferFunctionWithTransferFunctionTable(this.h, transferFunctionTable_ma))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorSpace) WithTransferFunctions(redTransferFunctionTable []uint16, greenTransferFunctionTable []uint16, blueTransferFunctionTable []uint16) *QColorSpace {
	redTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(redTransferFunctionTable))))
	defer free(unsafe.Pointer(redTransferFunctionTable_CArray))
	for i := range redTransferFunctionTable {
		redTransferFunctionTable_CArray[i] = (uint16_t)(redTransferFunctionTable[i])
	}
	redTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(redTransferFunctionTable)), data: unsafe.Pointer(redTransferFunctionTable_CArray)}
	greenTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(greenTransferFunctionTable))))
	defer free(unsafe.Pointer(greenTransferFunctionTable_CArray))
	for i := range greenTransferFunctionTable {
		greenTransferFunctionTable_CArray[i] = (uint16_t)(greenTransferFunctionTable[i])
	}
	greenTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(greenTransferFunctionTable)), data: unsafe.Pointer(greenTransferFunctionTable_CArray)}
	blueTransferFunctionTable_CArray := (*[0xffff]uint16_t)(malloc(size_t(8 * len(blueTransferFunctionTable))))
	defer free(unsafe.Pointer(blueTransferFunctionTable_CArray))
	for i := range blueTransferFunctionTable {
		blueTransferFunctionTable_CArray[i] = (uint16_t)(blueTransferFunctionTable[i])
	}
	blueTransferFunctionTable_ma := struct_miqt_array{len: size_t(len(blueTransferFunctionTable)), data: unsafe.Pointer(blueTransferFunctionTable_CArray)}
	_goptr := newQColorSpace(QColorSpace_WithTransferFunctions(this.h, redTransferFunctionTable_ma, greenTransferFunctionTable_ma, blueTransferFunctionTable_ma))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorSpace) SetPrimaries(primariesId Primaries) {
	QColorSpace_SetPrimaries(this.h, primariesId)
}

func (this *QColorSpace) SetPrimaries2(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF) {
	QColorSpace_SetPrimaries2(this.h, whitePoint.cPointer(), redPoint.cPointer(), greenPoint.cPointer(), bluePoint.cPointer())
}

func (this *QColorSpace) SetWhitePoint(whitePoint QPointF) {
	QColorSpace_SetWhitePoint(this.h, whitePoint.cPointer())
}

func (this *QColorSpace) WhitePoint() *QPointF {
	_goptr := newQPointF(QColorSpace_WhitePoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorSpace) TransformModel() TransformModel {
	xxxxxxxxx
}

func (this *QColorSpace) ColorModel() ColorModel {
	xxxxxxxxx
}

func (this *QColorSpace) Detach() {
	QColorSpace_Detach(this.h)
}

func (this *QColorSpace) IsValid() bool {
	return (bool)(QColorSpace_IsValid(this.h))
}

func (this *QColorSpace) IsValidTarget() bool {
	return (bool)(QColorSpace_IsValidTarget(this.h))
}

func QColorSpace_FromIccProfile(iccProfile []byte) *QColorSpace {
	iccProfile_alias := struct_miqt_string{}
	iccProfile_alias.data = (char)(unsafe.Pointer(&iccProfile[0]))
	iccProfile_alias.len = size_t(len(iccProfile))
	_goptr := newQColorSpace(QColorSpace_FromIccProfile(iccProfile_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorSpace) IccProfile() []byte {
	var _bytearray struct_miqt_string = QColorSpace_IccProfile(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QColorSpace) TransformationToColorSpace(colorspace *QColorSpace) *QColorTransform {
	_goptr := newQColorTransform(QColorSpace_TransformationToColorSpace(this.h, colorspace.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorSpace) SetTransferFunction2(transferFunction TransferFunction, gamma float32) {
	QColorSpace_SetTransferFunction2(this.h, transferFunction, (float)(gamma))
}

func (this *QColorSpace) WithTransferFunction2(transferFunction TransferFunction, gamma float32) *QColorSpace {
	_goptr := newQColorSpace(QColorSpace_WithTransferFunction2(this.h, transferFunction, (float)(gamma)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
