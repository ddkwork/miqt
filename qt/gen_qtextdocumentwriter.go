package qt

import (
	"unsafe"
)

type QTextDocumentWriter struct {
	h          uintptr
	isSubclass bool
}

// NewQTextDocumentWriter constructs a new QTextDocumentWriter object.
func NewQTextDocumentWriter() *QTextDocumentWriter {
	g := newQTextDocumentWriter(QTextDocumentWriter_new())
	g.isSubclass = true
	return g
}

// NewQTextDocumentWriter2 constructs a new QTextDocumentWriter object.
func NewQTextDocumentWriter2(device *QIODevice, format []byte) *QTextDocumentWriter {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	g := newQTextDocumentWriter(QTextDocumentWriter_new2(device.cPointer(), format_alias))
	g.isSubclass = true
	return g
}

// NewQTextDocumentWriter3 constructs a new QTextDocumentWriter object.
func NewQTextDocumentWriter3(fileName string) *QTextDocumentWriter {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQTextDocumentWriter(QTextDocumentWriter_new3(fileName_ms))
	g.isSubclass = true
	return g
}

// NewQTextDocumentWriter4 constructs a new QTextDocumentWriter object.
func NewQTextDocumentWriter4(fileName string, format []byte) *QTextDocumentWriter {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	g := newQTextDocumentWriter(QTextDocumentWriter_new4(fileName_ms, format_alias))
	g.isSubclass = true
	return g
}

func (this *QTextDocumentWriter) SetFormat(format []byte) {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	QTextDocumentWriter_SetFormat(this.h, format_alias)
}

func (this *QTextDocumentWriter) Format() []byte {
	var _bytearray struct_miqt_string = QTextDocumentWriter_Format(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QTextDocumentWriter) SetDevice(device *QIODevice) {
	QTextDocumentWriter_SetDevice(this.h, device.cPointer())
}

func (this *QTextDocumentWriter) Device() *QIODevice {
	return newQIODevice(QTextDocumentWriter_Device(this.h))
}

func (this *QTextDocumentWriter) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QTextDocumentWriter_SetFileName(this.h, fileName_ms)
}

func (this *QTextDocumentWriter) FileName() string {
	var _ms struct_miqt_string = QTextDocumentWriter_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocumentWriter) Write(document *QTextDocument) bool {
	return (bool)(QTextDocumentWriter_Write(this.h, document.cPointer()))
}

func (this *QTextDocumentWriter) WriteWithFragment(fragment *QTextDocumentFragment) bool {
	return (bool)(QTextDocumentWriter_WriteWithFragment(this.h, fragment.cPointer()))
}

func QTextDocumentWriter_SupportedDocumentFormats() [][]byte {
	var _ma struct_miqt_array = QTextDocumentWriter_SupportedDocumentFormats()
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}
