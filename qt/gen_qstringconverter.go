package qt

import (
	"unsafe"
)

type QStringEncoder struct {
	h          uintptr
	isSubclass bool
}

// NewQStringEncoder constructs a new QStringEncoder object.
func NewQStringEncoder() *QStringEncoder {
	g := newQStringEncoder(QStringEncoder_new())
	g.isSubclass = true
	return g
}

// NewQStringEncoder2 constructs a new QStringEncoder object.
func NewQStringEncoder2(encoding Encoding) *QStringEncoder {
	g := newQStringEncoder(QStringEncoder_new2(encoding))
	g.isSubclass = true
	return g
}

// NewQStringEncoder3 constructs a new QStringEncoder object.
func NewQStringEncoder3(name QAnyStringView) *QStringEncoder {
	g := newQStringEncoder(QStringEncoder_new3(name.cPointer()))
	g.isSubclass = true
	return g
}

// NewQStringEncoder4 constructs a new QStringEncoder object.
func NewQStringEncoder4(encoding Encoding, flags Flags) *QStringEncoder {
	g := newQStringEncoder(QStringEncoder_new4(encoding, flags))
	g.isSubclass = true
	return g
}

// NewQStringEncoder5 constructs a new QStringEncoder object.
func NewQStringEncoder5(name QAnyStringView, flags Flags) *QStringEncoder {
	g := newQStringEncoder(QStringEncoder_new5(name.cPointer(), flags))
	g.isSubclass = true
	return g
}

func (this *QStringEncoder) RequiredSpace(inputLength int64) int64 {
	return (int64)(QStringEncoder_RequiredSpace(this.h, (ptrdiff_t)(inputLength)))
}

type QStringDecoder struct {
	h          uintptr
	isSubclass bool
}

// NewQStringDecoder constructs a new QStringDecoder object.
func NewQStringDecoder(encoding Encoding) *QStringDecoder {
	g := newQStringDecoder(QStringDecoder_new(encoding))
	g.isSubclass = true
	return g
}

// NewQStringDecoder2 constructs a new QStringDecoder object.
func NewQStringDecoder2() *QStringDecoder {
	g := newQStringDecoder(QStringDecoder_new2())
	g.isSubclass = true
	return g
}

// NewQStringDecoder3 constructs a new QStringDecoder object.
func NewQStringDecoder3(name QAnyStringView) *QStringDecoder {
	g := newQStringDecoder(QStringDecoder_new3(name.cPointer()))
	g.isSubclass = true
	return g
}

// NewQStringDecoder4 constructs a new QStringDecoder object.
func NewQStringDecoder4(encoding Encoding, flags Flags) *QStringDecoder {
	g := newQStringDecoder(QStringDecoder_new4(encoding, flags))
	g.isSubclass = true
	return g
}

// NewQStringDecoder5 constructs a new QStringDecoder object.
func NewQStringDecoder5(name QAnyStringView, f Flags) *QStringDecoder {
	g := newQStringDecoder(QStringDecoder_new5(name.cPointer(), f))
	g.isSubclass = true
	return g
}

func (this *QStringDecoder) RequiredSpace(inputLength int64) int64 {
	return (int64)(QStringDecoder_RequiredSpace(this.h, (ptrdiff_t)(inputLength)))
}

func (this *QStringDecoder) AppendToBuffer(out *QChar, ba QByteArrayView) *QChar {
	return newQChar(QStringDecoder_AppendToBuffer(this.h, out.cPointer(), ba.cPointer()))
}

func QStringDecoder_DecoderForHtml(data QByteArrayView) *QStringDecoder {
	_goptr := newQStringDecoder(QStringDecoder_DecoderForHtml(data.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
