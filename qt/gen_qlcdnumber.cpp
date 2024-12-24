// +build ignore

#include <QEvent>
#include <QFrame>
#include <QLCDNumber>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qlcdnumber.h>
#include "gen_qlcdnumber.h"
class MiqtVirtualQLCDNumber : public virtual QLCDNumber {
public:
MiqtVirtualQLCDNumber(QWidget* parent): QLCDNumber(parent) {};
MiqtVirtualQLCDNumber(): QLCDNumber() {};
MiqtVirtualQLCDNumber(uint numDigits): QLCDNumber(numDigits) {};
MiqtVirtualQLCDNumber(uint numDigits, QWidget* parent): QLCDNumber(numDigits, parent) {};

virtual ~MiqtVirtualQLCDNumber() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QLCDNumber::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QLCDNumber_MetaObject(const_cast<MiqtVirtualQLCDNumber*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QLCDNumber::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QLCDNumber::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QLCDNumber_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QLCDNumber::qt_metacast(param1);

	}
};
QLCDNumber* QLCDNumber_new(QWidget* parent) {
return new MiqtVirtualQLCDNumber(parent);
}
QLCDNumber* QLCDNumber_new2() {
return new MiqtVirtualQLCDNumber();
}
QLCDNumber* QLCDNumber_new3(unsigned int numDigits) {
return new MiqtVirtualQLCDNumber(static_cast<uint>(numDigits));
}
QLCDNumber* QLCDNumber_new4(unsigned int numDigits, QWidget* parent) {
return new MiqtVirtualQLCDNumber(static_cast<uint>(numDigits), parent);
}
void QLCDNumber_virtbase(QLCDNumber* src
, QFrame** outptr_QFrame
) {
*outptr_QFrame = static_cast<QFrame*>(src);
}
QMetaObject* QLCDNumber_MetaObject(const QLCDNumber* self) {
	return (QMetaObject*) self->metaObject();
}
void* QLCDNumber_Metacast(QLCDNumber* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QLCDNumber_Tr(const char* s) {
	QString _ret = QLCDNumber::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QLCDNumber_SmallDecimalPoint(const QLCDNumber* self) {
	return self->smallDecimalPoint();
}
int QLCDNumber_DigitCount(const QLCDNumber* self) {
	return self->digitCount();
}
void QLCDNumber_SetDigitCount(QLCDNumber* self, int nDigits) {
	self->setDigitCount(static_cast<int>(nDigits));
}
bool QLCDNumber_CheckOverflow(const QLCDNumber* self, double num) {
	return self->checkOverflow(static_cast<double>(num));
}
bool QLCDNumber_CheckOverflowWithNum(const QLCDNumber* self, int num) {
	return self->checkOverflow(static_cast<int>(num));
}
Mode QLCDNumber_Mode(const QLCDNumber* self) {
	return self->mode();
}
void QLCDNumber_SetMode(QLCDNumber* self, Mode mode) {
	self->setMode(mode);
}
SegmentStyle QLCDNumber_SegmentStyle(const QLCDNumber* self) {
	return self->segmentStyle();
}
void QLCDNumber_SetSegmentStyle(QLCDNumber* self, SegmentStyle segmentStyle) {
	self->setSegmentStyle(segmentStyle);
}
double QLCDNumber_Value(const QLCDNumber* self) {
	return self->value();
}
int QLCDNumber_IntValue(const QLCDNumber* self) {
	return self->intValue();
}
QSize* QLCDNumber_SizeHint(const QLCDNumber* self) {
	return new QSize(self->sizeHint());
}
void QLCDNumber_Display(QLCDNumber* self, struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	self->display(str_QString);
}
void QLCDNumber_DisplayWithNum(QLCDNumber* self, int num) {
	self->display(static_cast<int>(num));
}
void QLCDNumber_Display2(QLCDNumber* self, double num) {
	self->display(static_cast<double>(num));
}
void QLCDNumber_SetHexMode(QLCDNumber* self) {
	self->setHexMode();
}
void QLCDNumber_SetDecMode(QLCDNumber* self) {
	self->setDecMode();
}
void QLCDNumber_SetOctMode(QLCDNumber* self) {
	self->setOctMode();
}
void QLCDNumber_SetBinMode(QLCDNumber* self) {
	self->setBinMode();
}
void QLCDNumber_SetSmallDecimalPoint(QLCDNumber* self, bool smallDecimalPoint) {
	self->setSmallDecimalPoint(smallDecimalPoint);
}
void QLCDNumber_Overflow(QLCDNumber* self) {
	self->overflow();
}
void QLCDNumber_connect_Overflow(QLCDNumber* self, intptr_t slot) {
	MiqtVirtualQLCDNumber::connect(self, static_cast<void (QLCDNumber::*)()>(&QLCDNumber::overflow), self, [=]() {
		miqt_exec_callback_QLCDNumber_Overflow(slot);
	});
}
struct miqt_string QLCDNumber_Tr2(const char* s, const char* c) {
	QString _ret = QLCDNumber::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QLCDNumber_Tr3(const char* s, const char* c, int n) {
	QString _ret = QLCDNumber::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QLCDNumber_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLCDNumber*>( (QLCDNumber*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QLCDNumber_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQLCDNumber*)(self) )->virtualbase_MetaObject();
}
void QLCDNumber_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLCDNumber*>( (QLCDNumber*)(self) )->handle__Metacast = slot;
}
void* QLCDNumber_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQLCDNumber*)(self) )->virtualbase_Metacast(param1);
}
void QLCDNumber_Delete(QLCDNumber* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQLCDNumber*>( self );
	} else {
		delete self;
	}
}
