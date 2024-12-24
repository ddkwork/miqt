// +build ignore

#include <QAbstractSpinBox>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QKeyEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionSpinBox>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <qabstractspinbox.h>
#include "gen_qabstractspinbox.h"
class MiqtVirtualQAbstractSpinBox : public virtual QAbstractSpinBox {
public:
MiqtVirtualQAbstractSpinBox(QWidget* parent): QAbstractSpinBox(parent) {};
MiqtVirtualQAbstractSpinBox(): QAbstractSpinBox() {};

virtual ~MiqtVirtualQAbstractSpinBox() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QAbstractSpinBox::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QAbstractSpinBox_MetaObject(const_cast<MiqtVirtualQAbstractSpinBox*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAbstractSpinBox::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QAbstractSpinBox::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAbstractSpinBox_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAbstractSpinBox::qt_metacast(param1);

	}
};
QAbstractSpinBox* QAbstractSpinBox_new(QWidget* parent) {
return new MiqtVirtualQAbstractSpinBox(parent);
}
QAbstractSpinBox* QAbstractSpinBox_new2() {
return new MiqtVirtualQAbstractSpinBox();
}
void QAbstractSpinBox_virtbase(QAbstractSpinBox* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QAbstractSpinBox_MetaObject(const QAbstractSpinBox* self) {
	return (QMetaObject*) self->metaObject();
}
void* QAbstractSpinBox_Metacast(QAbstractSpinBox* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QAbstractSpinBox_Tr(const char* s) {
	QString _ret = QAbstractSpinBox::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
ButtonSymbols QAbstractSpinBox_ButtonSymbols(const QAbstractSpinBox* self) {
	return self->buttonSymbols();
}
void QAbstractSpinBox_SetButtonSymbols(QAbstractSpinBox* self, ButtonSymbols bs) {
	self->setButtonSymbols(bs);
}
void QAbstractSpinBox_SetCorrectionMode(QAbstractSpinBox* self, CorrectionMode cm) {
	self->setCorrectionMode(cm);
}
CorrectionMode QAbstractSpinBox_CorrectionMode(const QAbstractSpinBox* self) {
	return self->correctionMode();
}
bool QAbstractSpinBox_HasAcceptableInput(const QAbstractSpinBox* self) {
	return self->hasAcceptableInput();
}
struct miqt_string QAbstractSpinBox_Text(const QAbstractSpinBox* self) {
	QString _ret = self->text();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAbstractSpinBox_SpecialValueText(const QAbstractSpinBox* self) {
	QString _ret = self->specialValueText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAbstractSpinBox_SetSpecialValueText(QAbstractSpinBox* self, struct miqt_string txt) {
	QString txt_QString = QString::fromUtf8(txt.data, txt.len);
	self->setSpecialValueText(txt_QString);
}
bool QAbstractSpinBox_Wrapping(const QAbstractSpinBox* self) {
	return self->wrapping();
}
void QAbstractSpinBox_SetWrapping(QAbstractSpinBox* self, bool w) {
	self->setWrapping(w);
}
void QAbstractSpinBox_SetReadOnly(QAbstractSpinBox* self, bool r) {
	self->setReadOnly(r);
}
bool QAbstractSpinBox_IsReadOnly(const QAbstractSpinBox* self) {
	return self->isReadOnly();
}
void QAbstractSpinBox_SetKeyboardTracking(QAbstractSpinBox* self, bool kt) {
	self->setKeyboardTracking(kt);
}
bool QAbstractSpinBox_KeyboardTracking(const QAbstractSpinBox* self) {
	return self->keyboardTracking();
}
void QAbstractSpinBox_SetAlignment(QAbstractSpinBox* self, int flag) {
	self->setAlignment(static_cast<Qt::Alignment>(flag));
}
int QAbstractSpinBox_Alignment(const QAbstractSpinBox* self) {
	Qt::Alignment _ret = self->alignment();
	return static_cast<int>(_ret);
}
void QAbstractSpinBox_SetFrame(QAbstractSpinBox* self, bool frame) {
	self->setFrame(frame);
}
bool QAbstractSpinBox_HasFrame(const QAbstractSpinBox* self) {
	return self->hasFrame();
}
void QAbstractSpinBox_SetAccelerated(QAbstractSpinBox* self, bool on) {
	self->setAccelerated(on);
}
bool QAbstractSpinBox_IsAccelerated(const QAbstractSpinBox* self) {
	return self->isAccelerated();
}
void QAbstractSpinBox_SetGroupSeparatorShown(QAbstractSpinBox* self, bool shown) {
	self->setGroupSeparatorShown(shown);
}
bool QAbstractSpinBox_IsGroupSeparatorShown(const QAbstractSpinBox* self) {
	return self->isGroupSeparatorShown();
}
QSize* QAbstractSpinBox_SizeHint(const QAbstractSpinBox* self) {
	return new QSize(self->sizeHint());
}
QSize* QAbstractSpinBox_MinimumSizeHint(const QAbstractSpinBox* self) {
	return new QSize(self->minimumSizeHint());
}
void QAbstractSpinBox_InterpretText(QAbstractSpinBox* self) {
	self->interpretText();
}
bool QAbstractSpinBox_Event(QAbstractSpinBox* self, QEvent* event) {
	return self->event(event);
}
QVariant* QAbstractSpinBox_InputMethodQuery(const QAbstractSpinBox* self, int param1) {
	return new QVariant(self->inputMethodQuery(static_cast<Qt::InputMethodQuery>(param1)));
}
int QAbstractSpinBox_Validate(const QAbstractSpinBox* self, struct miqt_string input, int* pos) {
	QString input_QString = QString::fromUtf8(input.data, input.len);
	QValidator::State _ret = self->validate(input_QString, static_cast<int&>(*pos));
	return static_cast<int>(_ret);
}
void QAbstractSpinBox_Fixup(const QAbstractSpinBox* self, struct miqt_string input) {
	QString input_QString = QString::fromUtf8(input.data, input.len);
	self->fixup(input_QString);
}
void QAbstractSpinBox_StepBy(QAbstractSpinBox* self, int steps) {
	self->stepBy(static_cast<int>(steps));
}
void QAbstractSpinBox_StepUp(QAbstractSpinBox* self) {
	self->stepUp();
}
void QAbstractSpinBox_StepDown(QAbstractSpinBox* self) {
	self->stepDown();
}
void QAbstractSpinBox_SelectAll(QAbstractSpinBox* self) {
	self->selectAll();
}
void QAbstractSpinBox_Clear(QAbstractSpinBox* self) {
	self->clear();
}
void QAbstractSpinBox_EditingFinished(QAbstractSpinBox* self) {
	self->editingFinished();
}
void QAbstractSpinBox_connect_EditingFinished(QAbstractSpinBox* self, intptr_t slot) {
	MiqtVirtualQAbstractSpinBox::connect(self, static_cast<void (QAbstractSpinBox::*)()>(&QAbstractSpinBox::editingFinished), self, [=]() {
		miqt_exec_callback_QAbstractSpinBox_EditingFinished(slot);
	});
}
struct miqt_string QAbstractSpinBox_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractSpinBox::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAbstractSpinBox_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractSpinBox::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAbstractSpinBox_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractSpinBox*>( (QAbstractSpinBox*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QAbstractSpinBox_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAbstractSpinBox*)(self) )->virtualbase_MetaObject();
}
void QAbstractSpinBox_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractSpinBox*>( (QAbstractSpinBox*)(self) )->handle__Metacast = slot;
}
void* QAbstractSpinBox_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAbstractSpinBox*)(self) )->virtualbase_Metacast(param1);
}
void QAbstractSpinBox_Delete(QAbstractSpinBox* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAbstractSpinBox*>( self );
	} else {
		delete self;
	}
}
