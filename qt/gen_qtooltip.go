package qt

import (
	"unsafe"
)

type QToolTip struct {
	h          uintptr
	isSubclass bool
}

func QToolTip_ShowText(pos *QPoint, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QToolTip_ShowText(pos.cPointer(), text_ms)
}

func QToolTip_HideText() {
	QToolTip_HideText()
}

func QToolTip_IsVisible() bool {
	return (bool)(QToolTip_IsVisible())
}

func QToolTip_Text() string {
	var _ms struct_miqt_string = QToolTip_Text()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QToolTip_Palette() *QPalette {
	_goptr := newQPalette(QToolTip_Palette())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QToolTip_SetPalette(palette *QPalette) {
	QToolTip_SetPalette(palette.cPointer())
}

func QToolTip_Font() *QFont {
	_goptr := newQFont(QToolTip_Font())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QToolTip_SetFont(font *QFont) {
	QToolTip_SetFont(font.cPointer())
}

func QToolTip_ShowText3(pos *QPoint, text string, w *QWidget) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QToolTip_ShowText3(pos.cPointer(), text_ms, w.cPointer())
}

func QToolTip_ShowText4(pos *QPoint, text string, w *QWidget, rect *QRect) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QToolTip_ShowText4(pos.cPointer(), text_ms, w.cPointer(), rect.cPointer())
}

func QToolTip_ShowText5(pos *QPoint, text string, w *QWidget, rect *QRect, msecShowTime int) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QToolTip_ShowText5(pos.cPointer(), text_ms, w.cPointer(), rect.cPointer(), (int)(msecShowTime))
}
