package qt

import (
	"unsafe"
)

type QJsonParseError__ParseError int

const (
	QJsonParseError__NoError               QJsonParseError__ParseError = 0
	QJsonParseError__UnterminatedObject    QJsonParseError__ParseError = 1
	QJsonParseError__MissingNameSeparator  QJsonParseError__ParseError = 2
	QJsonParseError__UnterminatedArray     QJsonParseError__ParseError = 3
	QJsonParseError__MissingValueSeparator QJsonParseError__ParseError = 4
	QJsonParseError__IllegalValue          QJsonParseError__ParseError = 5
	QJsonParseError__TerminationByNumber   QJsonParseError__ParseError = 6
	QJsonParseError__IllegalNumber         QJsonParseError__ParseError = 7
	QJsonParseError__IllegalEscapeSequence QJsonParseError__ParseError = 8
	QJsonParseError__IllegalUTF8String     QJsonParseError__ParseError = 9
	QJsonParseError__UnterminatedString    QJsonParseError__ParseError = 10
	QJsonParseError__MissingObject         QJsonParseError__ParseError = 11
	QJsonParseError__DeepNesting           QJsonParseError__ParseError = 12
	QJsonParseError__DocumentTooLarge      QJsonParseError__ParseError = 13
	QJsonParseError__GarbageAtEnd          QJsonParseError__ParseError = 14
)

type QJsonParseError struct {
	h          uintptr
	isSubclass bool
}

func (this *QJsonParseError) ErrorString() string {
	var _ms struct_miqt_string = QJsonParseError_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
