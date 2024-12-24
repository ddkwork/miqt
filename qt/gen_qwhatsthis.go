package qt

import (
	"unsafe"
)

type QWhatsThis struct {
	h          uintptr
	isSubclass bool
}

func QWhatsThis_EnterWhatsThisMode() {
	QWhatsThis_EnterWhatsThisMode()
}

func QWhatsThis_InWhatsThisMode() bool {
	return (bool)(QWhatsThis_InWhatsThisMode())
}

func QWhatsThis_LeaveWhatsThisMode() {
	QWhatsThis_LeaveWhatsThisMode()
}

func QWhatsThis_ShowText(pos *QPoint, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QWhatsThis_ShowText(pos.cPointer(), text_ms)
}

func QWhatsThis_HideText() {
	QWhatsThis_HideText()
}

func QWhatsThis_CreateAction() *QAction {
	return newQAction(QWhatsThis_CreateAction())
}

func QWhatsThis_ShowText3(pos *QPoint, text string, w *QWidget) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QWhatsThis_ShowText3(pos.cPointer(), text_ms, w.cPointer())
}

func QWhatsThis_CreateAction1(parent *QObject) *QAction {
	return newQAction(QWhatsThis_CreateAction1(parent.cPointer()))
}
