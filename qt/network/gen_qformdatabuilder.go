package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QFormDataBuilder__Option int

const (
	QFormDataBuilder__Default                          QFormDataBuilder__Option = 0
	QFormDataBuilder__OmitRfc8187EncodedFilename       QFormDataBuilder__Option = 1
	QFormDataBuilder__UseRfc7578PercentEncodedFilename QFormDataBuilder__Option = 2
	QFormDataBuilder__PreferLatin1EncodedFilename      QFormDataBuilder__Option = 4
	QFormDataBuilder__StrictRfc7578                    QFormDataBuilder__Option = 3
)

type QFormDataPartBuilder struct {
	h          uintptr
	isSubclass bool
}

// NewQFormDataPartBuilder constructs a new QFormDataPartBuilder object.
func NewQFormDataPartBuilder() *QFormDataPartBuilder {
	g := newQFormDataPartBuilder(QFormDataPartBuilder_new())
	g.isSubclass = true
	return g
}

// NewQFormDataPartBuilder2 constructs a new QFormDataPartBuilder object.
func NewQFormDataPartBuilder2(param1 *QFormDataPartBuilder) *QFormDataPartBuilder {
	g := newQFormDataPartBuilder(QFormDataPartBuilder_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QFormDataPartBuilder) Swap(other *QFormDataPartBuilder) {
	QFormDataPartBuilder_Swap(this.h, other.cPointer())
}

func (this *QFormDataPartBuilder) SetBody(data qt.QByteArrayView) *QFormDataPartBuilder {
	_goptr := newQFormDataPartBuilder(QFormDataPartBuilder_SetBody(this.h, (*QByteArrayView)(data.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormDataPartBuilder) SetBodyDevice(body *qt.QIODevice) *QFormDataPartBuilder {
	_goptr := newQFormDataPartBuilder(QFormDataPartBuilder_SetBodyDevice(this.h, (*QIODevice)(body.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormDataPartBuilder) SetHeaders(headers *QHttpHeaders) *QFormDataPartBuilder {
	_goptr := newQFormDataPartBuilder(QFormDataPartBuilder_SetHeaders(this.h, headers.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormDataPartBuilder) SetBody2(data qt.QByteArrayView, fileName qt.QAnyStringView) *QFormDataPartBuilder {
	_goptr := newQFormDataPartBuilder(QFormDataPartBuilder_SetBody2(this.h, (*QByteArrayView)(data.UnsafePointer()), (*QAnyStringView)(fileName.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormDataPartBuilder) SetBody3(data qt.QByteArrayView, fileName qt.QAnyStringView, mimeType qt.QAnyStringView) *QFormDataPartBuilder {
	_goptr := newQFormDataPartBuilder(QFormDataPartBuilder_SetBody3(this.h, (*QByteArrayView)(data.UnsafePointer()), (*QAnyStringView)(fileName.UnsafePointer()), (*QAnyStringView)(mimeType.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormDataPartBuilder) SetBodyDevice2(body *qt.QIODevice, fileName qt.QAnyStringView) *QFormDataPartBuilder {
	_goptr := newQFormDataPartBuilder(QFormDataPartBuilder_SetBodyDevice2(this.h, (*QIODevice)(body.UnsafePointer()), (*QAnyStringView)(fileName.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormDataPartBuilder) SetBodyDevice3(body *qt.QIODevice, fileName qt.QAnyStringView, mimeType qt.QAnyStringView) *QFormDataPartBuilder {
	_goptr := newQFormDataPartBuilder(QFormDataPartBuilder_SetBodyDevice3(this.h, (*QIODevice)(body.UnsafePointer()), (*QAnyStringView)(fileName.UnsafePointer()), (*QAnyStringView)(mimeType.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QFormDataBuilder struct {
	h          uintptr
	isSubclass bool
}

// NewQFormDataBuilder constructs a new QFormDataBuilder object.
func NewQFormDataBuilder() *QFormDataBuilder {
	g := newQFormDataBuilder(QFormDataBuilder_new())
	g.isSubclass = true
	return g
}

func (this *QFormDataBuilder) Swap(other *QFormDataBuilder) {
	QFormDataBuilder_Swap(this.h, other.cPointer())
}

func (this *QFormDataBuilder) Part(name qt.QAnyStringView) *QFormDataPartBuilder {
	_goptr := newQFormDataPartBuilder(QFormDataBuilder_Part(this.h, (*QAnyStringView)(name.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
