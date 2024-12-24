package cbor

import (
	"unsafe"
)

type QCborSimpleType byte

const (
	QCborSimpleType__False     QCborSimpleType = 20
	QCborSimpleType__True      QCborSimpleType = 21
	QCborSimpleType__Null      QCborSimpleType = 22
	QCborSimpleType__Undefined QCborSimpleType = 23
)

type (
	QCborTag       uint64
	QCborKnownTags int
)

const (
	QCborKnownTags__DateTimeString    QCborKnownTags = 0
	QCborKnownTags__UnixTime_t        QCborKnownTags = 1
	QCborKnownTags__PositiveBignum    QCborKnownTags = 2
	QCborKnownTags__NegativeBignum    QCborKnownTags = 3
	QCborKnownTags__Decimal           QCborKnownTags = 4
	QCborKnownTags__Bigfloat          QCborKnownTags = 5
	QCborKnownTags__COSE_Encrypt0     QCborKnownTags = 16
	QCborKnownTags__COSE_Mac0         QCborKnownTags = 17
	QCborKnownTags__COSE_Sign1        QCborKnownTags = 18
	QCborKnownTags__ExpectedBase64url QCborKnownTags = 21
	QCborKnownTags__ExpectedBase64    QCborKnownTags = 22
	QCborKnownTags__ExpectedBase16    QCborKnownTags = 23
	QCborKnownTags__EncodedCbor       QCborKnownTags = 24
	QCborKnownTags__Url               QCborKnownTags = 32
	QCborKnownTags__Base64url         QCborKnownTags = 33
	QCborKnownTags__Base64            QCborKnownTags = 34
	QCborKnownTags__RegularExpression QCborKnownTags = 35
	QCborKnownTags__MimeMessage       QCborKnownTags = 36
	QCborKnownTags__Uuid              QCborKnownTags = 37
	QCborKnownTags__COSE_Encrypt      QCborKnownTags = 96
	QCborKnownTags__COSE_Mac          QCborKnownTags = 97
	QCborKnownTags__COSE_Sign         QCborKnownTags = 98
	QCborKnownTags__Signature         QCborKnownTags = 55799
)

type (
	QCborNegativeInteger uint64
	QCborError__Code     int
)

const (
	QCborError__UnknownError      QCborError__Code = 1
	QCborError__AdvancePastEnd    QCborError__Code = 3
	QCborError__InputOutputError  QCborError__Code = 4
	QCborError__GarbageAtEnd      QCborError__Code = 256
	QCborError__EndOfFile         QCborError__Code = 257
	QCborError__UnexpectedBreak   QCborError__Code = 258
	QCborError__UnknownType       QCborError__Code = 259
	QCborError__IllegalType       QCborError__Code = 260
	QCborError__IllegalNumber     QCborError__Code = 261
	QCborError__IllegalSimpleType QCborError__Code = 262
	QCborError__InvalidUtf8String QCborError__Code = 516
	QCborError__DataTooLarge      QCborError__Code = 1024
	QCborError__NestingTooDeep    QCborError__Code = 1025
	QCborError__UnsupportedType   QCborError__Code = 1026
	QCborError__NoError           QCborError__Code = 0
)

type QCborError struct {
	h          uintptr
	isSubclass bool
}

func (this *QCborError) ToString() string {
	var _ms struct_miqt_string = QCborError_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
