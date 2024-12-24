// +build ignore

#include <QDoubleValidator>
#include <QIntValidator>
#include <QLocale>
#include <QMetaObject>
#include <QObject>
#include <QRegularExpression>
#include <QRegularExpressionValidator>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QValidator>
#include <qvalidator.h>
#include "gen_qvalidator.h"
class MiqtVirtualQValidator : public virtual QValidator {
public:
MiqtVirtualQValidator(): QValidator() {};
MiqtVirtualQValidator(QObject* parent): QValidator(parent) {};

virtual ~MiqtVirtualQValidator() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QValidator::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QValidator_MetaObject(const_cast<MiqtVirtualQValidator*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QValidator::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QValidator::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QValidator_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QValidator::qt_metacast(param1);

	}
};
QValidator* QValidator_new() {
return new MiqtVirtualQValidator();
}
QValidator* QValidator_new2(QObject* parent) {
return new MiqtVirtualQValidator(parent);
}
void QValidator_virtbase(QValidator* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QValidator_MetaObject(const QValidator* self) {
	return (QMetaObject*) self->metaObject();
}
void* QValidator_Metacast(QValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QValidator_Tr(const char* s) {
	QString _ret = QValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QValidator_SetLocale(QValidator* self, QLocale* locale) {
	self->setLocale(*locale);
}
QLocale* QValidator_Locale(const QValidator* self) {
	return new QLocale(self->locale());
}
State QValidator_Validate(const QValidator* self, struct miqt_string param1, int* param2) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	return self->validate(param1_QString, static_cast<int&>(*param2));
}
void QValidator_Fixup(const QValidator* self, struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	self->fixup(param1_QString);
}
void QValidator_Changed(QValidator* self) {
	self->changed();
}
void QValidator_connect_Changed(QValidator* self, intptr_t slot) {
	MiqtVirtualQValidator::connect(self, static_cast<void (QValidator::*)()>(&QValidator::changed), self, [=]() {
		miqt_exec_callback_QValidator_Changed(slot);
	});
}
struct miqt_string QValidator_Tr2(const char* s, const char* c) {
	QString _ret = QValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QValidator_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQValidator*>( (QValidator*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QValidator_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQValidator*)(self) )->virtualbase_MetaObject();
}
void QValidator_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQValidator*>( (QValidator*)(self) )->handle__Metacast = slot;
}
void* QValidator_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQValidator*)(self) )->virtualbase_Metacast(param1);
}
void QValidator_Delete(QValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQValidator*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQIntValidator : public virtual QIntValidator {
public:
MiqtVirtualQIntValidator(): QIntValidator() {};
MiqtVirtualQIntValidator(int bottom, int top): QIntValidator(bottom, top) {};
MiqtVirtualQIntValidator(QObject* parent): QIntValidator(parent) {};
MiqtVirtualQIntValidator(int bottom, int top, QObject* parent): QIntValidator(bottom, top, parent) {};

virtual ~MiqtVirtualQIntValidator() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QIntValidator::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QIntValidator_MetaObject(const_cast<MiqtVirtualQIntValidator*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QIntValidator::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QIntValidator::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QIntValidator_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QIntValidator::qt_metacast(param1);

	}
};
QIntValidator* QIntValidator_new() {
return new MiqtVirtualQIntValidator();
}
QIntValidator* QIntValidator_new2(int bottom, int top) {
return new MiqtVirtualQIntValidator(static_cast<int>(bottom), static_cast<int>(top));
}
QIntValidator* QIntValidator_new3(QObject* parent) {
return new MiqtVirtualQIntValidator(parent);
}
QIntValidator* QIntValidator_new4(int bottom, int top, QObject* parent) {
return new MiqtVirtualQIntValidator(static_cast<int>(bottom), static_cast<int>(top), parent);
}
void QIntValidator_virtbase(QIntValidator* src
, QValidator** outptr_QValidator
) {
*outptr_QValidator = static_cast<QValidator*>(src);
}
QMetaObject* QIntValidator_MetaObject(const QIntValidator* self) {
	return (QMetaObject*) self->metaObject();
}
void* QIntValidator_Metacast(QIntValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QIntValidator_Tr(const char* s) {
	QString _ret = QIntValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QIntValidator_Validate(const QIntValidator* self, struct miqt_string param1, int* param2) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QValidator::State _ret = self->validate(param1_QString, static_cast<int&>(*param2));
	return static_cast<int>(_ret);
}
void QIntValidator_Fixup(const QIntValidator* self, struct miqt_string input) {
	QString input_QString = QString::fromUtf8(input.data, input.len);
	self->fixup(input_QString);
}
void QIntValidator_SetBottom(QIntValidator* self, int bottom) {
	self->setBottom(static_cast<int>(bottom));
}
void QIntValidator_SetTop(QIntValidator* self, int top) {
	self->setTop(static_cast<int>(top));
}
void QIntValidator_SetRange(QIntValidator* self, int bottom, int top) {
	self->setRange(static_cast<int>(bottom), static_cast<int>(top));
}
int QIntValidator_Bottom(const QIntValidator* self) {
	return self->bottom();
}
int QIntValidator_Top(const QIntValidator* self) {
	return self->top();
}
void QIntValidator_BottomChanged(QIntValidator* self, int bottom) {
	self->bottomChanged(static_cast<int>(bottom));
}
void QIntValidator_connect_BottomChanged(QIntValidator* self, intptr_t slot) {
	MiqtVirtualQIntValidator::connect(self, static_cast<void (QIntValidator::*)(int)>(&QIntValidator::bottomChanged), self, [=](int bottom) {
		int sigval1 = bottom;
		miqt_exec_callback_QIntValidator_BottomChanged(slot, sigval1);
	});
}
void QIntValidator_TopChanged(QIntValidator* self, int top) {
	self->topChanged(static_cast<int>(top));
}
void QIntValidator_connect_TopChanged(QIntValidator* self, intptr_t slot) {
	MiqtVirtualQIntValidator::connect(self, static_cast<void (QIntValidator::*)(int)>(&QIntValidator::topChanged), self, [=](int top) {
		int sigval1 = top;
		miqt_exec_callback_QIntValidator_TopChanged(slot, sigval1);
	});
}
struct miqt_string QIntValidator_Tr2(const char* s, const char* c) {
	QString _ret = QIntValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QIntValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QIntValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QIntValidator_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIntValidator*>( (QIntValidator*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QIntValidator_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQIntValidator*)(self) )->virtualbase_MetaObject();
}
void QIntValidator_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIntValidator*>( (QIntValidator*)(self) )->handle__Metacast = slot;
}
void* QIntValidator_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQIntValidator*)(self) )->virtualbase_Metacast(param1);
}
void QIntValidator_Delete(QIntValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQIntValidator*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQDoubleValidator : public virtual QDoubleValidator {
public:
MiqtVirtualQDoubleValidator(): QDoubleValidator() {};
MiqtVirtualQDoubleValidator(double bottom, double top, int decimals): QDoubleValidator(bottom, top, decimals) {};
MiqtVirtualQDoubleValidator(QObject* parent): QDoubleValidator(parent) {};
MiqtVirtualQDoubleValidator(double bottom, double top, int decimals, QObject* parent): QDoubleValidator(bottom, top, decimals, parent) {};

virtual ~MiqtVirtualQDoubleValidator() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QDoubleValidator::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QDoubleValidator_MetaObject(const_cast<MiqtVirtualQDoubleValidator*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QDoubleValidator::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QDoubleValidator::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QDoubleValidator_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QDoubleValidator::qt_metacast(param1);

	}
};
QDoubleValidator* QDoubleValidator_new() {
return new MiqtVirtualQDoubleValidator();
}
QDoubleValidator* QDoubleValidator_new2(double bottom, double top, int decimals) {
return new MiqtVirtualQDoubleValidator(static_cast<double>(bottom), static_cast<double>(top), static_cast<int>(decimals));
}
QDoubleValidator* QDoubleValidator_new3(QObject* parent) {
return new MiqtVirtualQDoubleValidator(parent);
}
QDoubleValidator* QDoubleValidator_new4(double bottom, double top, int decimals, QObject* parent) {
return new MiqtVirtualQDoubleValidator(static_cast<double>(bottom), static_cast<double>(top), static_cast<int>(decimals), parent);
}
void QDoubleValidator_virtbase(QDoubleValidator* src
, QValidator** outptr_QValidator
) {
*outptr_QValidator = static_cast<QValidator*>(src);
}
QMetaObject* QDoubleValidator_MetaObject(const QDoubleValidator* self) {
	return (QMetaObject*) self->metaObject();
}
void* QDoubleValidator_Metacast(QDoubleValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QDoubleValidator_Tr(const char* s) {
	QString _ret = QDoubleValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QDoubleValidator_Validate(const QDoubleValidator* self, struct miqt_string param1, int* param2) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QValidator::State _ret = self->validate(param1_QString, static_cast<int&>(*param2));
	return static_cast<int>(_ret);
}
void QDoubleValidator_Fixup(const QDoubleValidator* self, struct miqt_string input) {
	QString input_QString = QString::fromUtf8(input.data, input.len);
	self->fixup(input_QString);
}
void QDoubleValidator_SetRange(QDoubleValidator* self, double bottom, double top, int decimals) {
	self->setRange(static_cast<double>(bottom), static_cast<double>(top), static_cast<int>(decimals));
}
void QDoubleValidator_SetRange2(QDoubleValidator* self, double bottom, double top) {
	self->setRange(static_cast<double>(bottom), static_cast<double>(top));
}
void QDoubleValidator_SetBottom(QDoubleValidator* self, double bottom) {
	self->setBottom(static_cast<double>(bottom));
}
void QDoubleValidator_SetTop(QDoubleValidator* self, double top) {
	self->setTop(static_cast<double>(top));
}
void QDoubleValidator_SetDecimals(QDoubleValidator* self, int decimals) {
	self->setDecimals(static_cast<int>(decimals));
}
void QDoubleValidator_SetNotation(QDoubleValidator* self, Notation notation) {
	self->setNotation(notation);
}
double QDoubleValidator_Bottom(const QDoubleValidator* self) {
	return self->bottom();
}
double QDoubleValidator_Top(const QDoubleValidator* self) {
	return self->top();
}
int QDoubleValidator_Decimals(const QDoubleValidator* self) {
	return self->decimals();
}
Notation QDoubleValidator_Notation(const QDoubleValidator* self) {
	return self->notation();
}
void QDoubleValidator_BottomChanged(QDoubleValidator* self, double bottom) {
	self->bottomChanged(static_cast<double>(bottom));
}
void QDoubleValidator_connect_BottomChanged(QDoubleValidator* self, intptr_t slot) {
	MiqtVirtualQDoubleValidator::connect(self, static_cast<void (QDoubleValidator::*)(double)>(&QDoubleValidator::bottomChanged), self, [=](double bottom) {
		double sigval1 = bottom;
		miqt_exec_callback_QDoubleValidator_BottomChanged(slot, sigval1);
	});
}
void QDoubleValidator_TopChanged(QDoubleValidator* self, double top) {
	self->topChanged(static_cast<double>(top));
}
void QDoubleValidator_connect_TopChanged(QDoubleValidator* self, intptr_t slot) {
	MiqtVirtualQDoubleValidator::connect(self, static_cast<void (QDoubleValidator::*)(double)>(&QDoubleValidator::topChanged), self, [=](double top) {
		double sigval1 = top;
		miqt_exec_callback_QDoubleValidator_TopChanged(slot, sigval1);
	});
}
void QDoubleValidator_DecimalsChanged(QDoubleValidator* self, int decimals) {
	self->decimalsChanged(static_cast<int>(decimals));
}
void QDoubleValidator_connect_DecimalsChanged(QDoubleValidator* self, intptr_t slot) {
	MiqtVirtualQDoubleValidator::connect(self, static_cast<void (QDoubleValidator::*)(int)>(&QDoubleValidator::decimalsChanged), self, [=](int decimals) {
		int sigval1 = decimals;
		miqt_exec_callback_QDoubleValidator_DecimalsChanged(slot, sigval1);
	});
}
void QDoubleValidator_NotationChanged(QDoubleValidator* self, int notation) {
	self->notationChanged(static_cast<QDoubleValidator::Notation>(notation));
}
void QDoubleValidator_connect_NotationChanged(QDoubleValidator* self, intptr_t slot) {
	MiqtVirtualQDoubleValidator::connect(self, static_cast<void (QDoubleValidator::*)(QDoubleValidator::Notation)>(&QDoubleValidator::notationChanged), self, [=](QDoubleValidator::Notation notation) {
		QDoubleValidator::Notation notation_ret = notation;
		int sigval1 = static_cast<int>(notation_ret);
		miqt_exec_callback_QDoubleValidator_NotationChanged(slot, sigval1);
	});
}
struct miqt_string QDoubleValidator_Tr2(const char* s, const char* c) {
	QString _ret = QDoubleValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QDoubleValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDoubleValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QDoubleValidator_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDoubleValidator*>( (QDoubleValidator*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QDoubleValidator_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQDoubleValidator*)(self) )->virtualbase_MetaObject();
}
void QDoubleValidator_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDoubleValidator*>( (QDoubleValidator*)(self) )->handle__Metacast = slot;
}
void* QDoubleValidator_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQDoubleValidator*)(self) )->virtualbase_Metacast(param1);
}
void QDoubleValidator_Delete(QDoubleValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQDoubleValidator*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQRegularExpressionValidator : public virtual QRegularExpressionValidator {
public:
MiqtVirtualQRegularExpressionValidator(): QRegularExpressionValidator() {};
MiqtVirtualQRegularExpressionValidator(const QRegularExpression& re): QRegularExpressionValidator(re) {};
MiqtVirtualQRegularExpressionValidator(QObject* parent): QRegularExpressionValidator(parent) {};
MiqtVirtualQRegularExpressionValidator(const QRegularExpression& re, QObject* parent): QRegularExpressionValidator(re, parent) {};

virtual ~MiqtVirtualQRegularExpressionValidator() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QRegularExpressionValidator::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QRegularExpressionValidator_MetaObject(const_cast<MiqtVirtualQRegularExpressionValidator*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QRegularExpressionValidator::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QRegularExpressionValidator::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QRegularExpressionValidator_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QRegularExpressionValidator::qt_metacast(param1);

	}
};
QRegularExpressionValidator* QRegularExpressionValidator_new() {
return new MiqtVirtualQRegularExpressionValidator();
}
QRegularExpressionValidator* QRegularExpressionValidator_new2(QRegularExpression* re) {
return new MiqtVirtualQRegularExpressionValidator(*re);
}
QRegularExpressionValidator* QRegularExpressionValidator_new3(QObject* parent) {
return new MiqtVirtualQRegularExpressionValidator(parent);
}
QRegularExpressionValidator* QRegularExpressionValidator_new4(QRegularExpression* re, QObject* parent) {
return new MiqtVirtualQRegularExpressionValidator(*re, parent);
}
void QRegularExpressionValidator_virtbase(QRegularExpressionValidator* src
, QValidator** outptr_QValidator
) {
*outptr_QValidator = static_cast<QValidator*>(src);
}
QMetaObject* QRegularExpressionValidator_MetaObject(const QRegularExpressionValidator* self) {
	return (QMetaObject*) self->metaObject();
}
void* QRegularExpressionValidator_Metacast(QRegularExpressionValidator* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QRegularExpressionValidator_Tr(const char* s) {
	QString _ret = QRegularExpressionValidator::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QRegularExpressionValidator_Validate(const QRegularExpressionValidator* self, struct miqt_string input, int* pos) {
	QString input_QString = QString::fromUtf8(input.data, input.len);
	QValidator::State _ret = self->validate(input_QString, static_cast<int&>(*pos));
	return static_cast<int>(_ret);
}
QRegularExpression* QRegularExpressionValidator_RegularExpression(const QRegularExpressionValidator* self) {
	return new QRegularExpression(self->regularExpression());
}
void QRegularExpressionValidator_SetRegularExpression(QRegularExpressionValidator* self, QRegularExpression* re) {
	self->setRegularExpression(*re);
}
void QRegularExpressionValidator_RegularExpressionChanged(QRegularExpressionValidator* self, QRegularExpression* re) {
	self->regularExpressionChanged(*re);
}
void QRegularExpressionValidator_connect_RegularExpressionChanged(QRegularExpressionValidator* self, intptr_t slot) {
	MiqtVirtualQRegularExpressionValidator::connect(self, static_cast<void (QRegularExpressionValidator::*)(const QRegularExpression&)>(&QRegularExpressionValidator::regularExpressionChanged), self, [=](const QRegularExpression& re) {
		const QRegularExpression& re_ret = re;
		// Cast returned reference into pointer
		QRegularExpression* sigval1 = const_cast<QRegularExpression*>(&re_ret);
		miqt_exec_callback_QRegularExpressionValidator_RegularExpressionChanged(slot, sigval1);
	});
}
struct miqt_string QRegularExpressionValidator_Tr2(const char* s, const char* c) {
	QString _ret = QRegularExpressionValidator::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QRegularExpressionValidator_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRegularExpressionValidator::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QRegularExpressionValidator_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRegularExpressionValidator*>( (QRegularExpressionValidator*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QRegularExpressionValidator_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQRegularExpressionValidator*)(self) )->virtualbase_MetaObject();
}
void QRegularExpressionValidator_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRegularExpressionValidator*>( (QRegularExpressionValidator*)(self) )->handle__Metacast = slot;
}
void* QRegularExpressionValidator_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQRegularExpressionValidator*)(self) )->virtualbase_Metacast(param1);
}
void QRegularExpressionValidator_Delete(QRegularExpressionValidator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRegularExpressionValidator*>( self );
	} else {
		delete self;
	}
}
