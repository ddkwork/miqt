// +build ignore

#include <QAbstractSlider>
#include <QEvent>
#include <QKeyEvent>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <QWheelEvent>
#include <QWidget>
#include <qabstractslider.h>
#include "gen_qabstractslider.h"
class MiqtVirtualQAbstractSlider : public virtual QAbstractSlider {
public:
MiqtVirtualQAbstractSlider(QWidget* parent): QAbstractSlider(parent) {};
MiqtVirtualQAbstractSlider(): QAbstractSlider() {};

virtual ~MiqtVirtualQAbstractSlider() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QAbstractSlider::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QAbstractSlider_MetaObject(const_cast<MiqtVirtualQAbstractSlider*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAbstractSlider::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QAbstractSlider::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAbstractSlider_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAbstractSlider::qt_metacast(param1);

	}
};
QAbstractSlider* QAbstractSlider_new(QWidget* parent) {
return new MiqtVirtualQAbstractSlider(parent);
}
QAbstractSlider* QAbstractSlider_new2() {
return new MiqtVirtualQAbstractSlider();
}
void QAbstractSlider_virtbase(QAbstractSlider* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QAbstractSlider_MetaObject(const QAbstractSlider* self) {
	return (QMetaObject*) self->metaObject();
}
void* QAbstractSlider_Metacast(QAbstractSlider* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QAbstractSlider_Tr(const char* s) {
	QString _ret = QAbstractSlider::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QAbstractSlider_Orientation(const QAbstractSlider* self) {
	Qt::Orientation _ret = self->orientation();
	return static_cast<int>(_ret);
}
void QAbstractSlider_SetMinimum(QAbstractSlider* self, int minimum) {
	self->setMinimum(static_cast<int>(minimum));
}
int QAbstractSlider_Minimum(const QAbstractSlider* self) {
	return self->minimum();
}
void QAbstractSlider_SetMaximum(QAbstractSlider* self, int maximum) {
	self->setMaximum(static_cast<int>(maximum));
}
int QAbstractSlider_Maximum(const QAbstractSlider* self) {
	return self->maximum();
}
void QAbstractSlider_SetSingleStep(QAbstractSlider* self, int singleStep) {
	self->setSingleStep(static_cast<int>(singleStep));
}
int QAbstractSlider_SingleStep(const QAbstractSlider* self) {
	return self->singleStep();
}
void QAbstractSlider_SetPageStep(QAbstractSlider* self, int pageStep) {
	self->setPageStep(static_cast<int>(pageStep));
}
int QAbstractSlider_PageStep(const QAbstractSlider* self) {
	return self->pageStep();
}
void QAbstractSlider_SetTracking(QAbstractSlider* self, bool enable) {
	self->setTracking(enable);
}
bool QAbstractSlider_HasTracking(const QAbstractSlider* self) {
	return self->hasTracking();
}
void QAbstractSlider_SetSliderDown(QAbstractSlider* self, bool sliderDown) {
	self->setSliderDown(sliderDown);
}
bool QAbstractSlider_IsSliderDown(const QAbstractSlider* self) {
	return self->isSliderDown();
}
void QAbstractSlider_SetSliderPosition(QAbstractSlider* self, int sliderPosition) {
	self->setSliderPosition(static_cast<int>(sliderPosition));
}
int QAbstractSlider_SliderPosition(const QAbstractSlider* self) {
	return self->sliderPosition();
}
void QAbstractSlider_SetInvertedAppearance(QAbstractSlider* self, bool invertedAppearance) {
	self->setInvertedAppearance(invertedAppearance);
}
bool QAbstractSlider_InvertedAppearance(const QAbstractSlider* self) {
	return self->invertedAppearance();
}
void QAbstractSlider_SetInvertedControls(QAbstractSlider* self, bool invertedControls) {
	self->setInvertedControls(invertedControls);
}
bool QAbstractSlider_InvertedControls(const QAbstractSlider* self) {
	return self->invertedControls();
}
int QAbstractSlider_Value(const QAbstractSlider* self) {
	return self->value();
}
void QAbstractSlider_TriggerAction(QAbstractSlider* self, SliderAction action) {
	self->triggerAction(action);
}
void QAbstractSlider_SetValue(QAbstractSlider* self, int value) {
	self->setValue(static_cast<int>(value));
}
void QAbstractSlider_SetOrientation(QAbstractSlider* self, int orientation) {
	self->setOrientation(static_cast<Qt::Orientation>(orientation));
}
void QAbstractSlider_SetRange(QAbstractSlider* self, int min, int max) {
	self->setRange(static_cast<int>(min), static_cast<int>(max));
}
void QAbstractSlider_ValueChanged(QAbstractSlider* self, int value) {
	self->valueChanged(static_cast<int>(value));
}
void QAbstractSlider_connect_ValueChanged(QAbstractSlider* self, intptr_t slot) {
	MiqtVirtualQAbstractSlider::connect(self, static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::valueChanged), self, [=](int value) {
		int sigval1 = value;
		miqt_exec_callback_QAbstractSlider_ValueChanged(slot, sigval1);
	});
}
void QAbstractSlider_SliderPressed(QAbstractSlider* self) {
	self->sliderPressed();
}
void QAbstractSlider_connect_SliderPressed(QAbstractSlider* self, intptr_t slot) {
	MiqtVirtualQAbstractSlider::connect(self, static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderPressed), self, [=]() {
		miqt_exec_callback_QAbstractSlider_SliderPressed(slot);
	});
}
void QAbstractSlider_SliderMoved(QAbstractSlider* self, int position) {
	self->sliderMoved(static_cast<int>(position));
}
void QAbstractSlider_connect_SliderMoved(QAbstractSlider* self, intptr_t slot) {
	MiqtVirtualQAbstractSlider::connect(self, static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::sliderMoved), self, [=](int position) {
		int sigval1 = position;
		miqt_exec_callback_QAbstractSlider_SliderMoved(slot, sigval1);
	});
}
void QAbstractSlider_SliderReleased(QAbstractSlider* self) {
	self->sliderReleased();
}
void QAbstractSlider_connect_SliderReleased(QAbstractSlider* self, intptr_t slot) {
	MiqtVirtualQAbstractSlider::connect(self, static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderReleased), self, [=]() {
		miqt_exec_callback_QAbstractSlider_SliderReleased(slot);
	});
}
void QAbstractSlider_RangeChanged(QAbstractSlider* self, int min, int max) {
	self->rangeChanged(static_cast<int>(min), static_cast<int>(max));
}
void QAbstractSlider_connect_RangeChanged(QAbstractSlider* self, intptr_t slot) {
	MiqtVirtualQAbstractSlider::connect(self, static_cast<void (QAbstractSlider::*)(int, int)>(&QAbstractSlider::rangeChanged), self, [=](int min, int max) {
		int sigval1 = min;
		int sigval2 = max;
		miqt_exec_callback_QAbstractSlider_RangeChanged(slot, sigval1, sigval2);
	});
}
void QAbstractSlider_ActionTriggered(QAbstractSlider* self, int action) {
	self->actionTriggered(static_cast<int>(action));
}
void QAbstractSlider_connect_ActionTriggered(QAbstractSlider* self, intptr_t slot) {
	MiqtVirtualQAbstractSlider::connect(self, static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::actionTriggered), self, [=](int action) {
		int sigval1 = action;
		miqt_exec_callback_QAbstractSlider_ActionTriggered(slot, sigval1);
	});
}
struct miqt_string QAbstractSlider_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractSlider::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAbstractSlider_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractSlider::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAbstractSlider_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractSlider*>( (QAbstractSlider*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QAbstractSlider_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAbstractSlider*)(self) )->virtualbase_MetaObject();
}
void QAbstractSlider_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractSlider*>( (QAbstractSlider*)(self) )->handle__Metacast = slot;
}
void* QAbstractSlider_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAbstractSlider*)(self) )->virtualbase_Metacast(param1);
}
void QAbstractSlider_Delete(QAbstractSlider* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAbstractSlider*>( self );
	} else {
		delete self;
	}
}
