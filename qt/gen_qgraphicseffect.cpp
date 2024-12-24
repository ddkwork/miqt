// +build ignore

#include <QBrush>
#include <QColor>
#include <QGraphicsBlurEffect>
#include <QGraphicsColorizeEffect>
#include <QGraphicsDropShadowEffect>
#include <QGraphicsEffect>
#include <QGraphicsOpacityEffect>
#include <QMetaObject>
#include <QObject>
#include <QPainter>
#include <QPointF>
#include <QRectF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qgraphicseffect.h>
#include "gen_qgraphicseffect.h"
class MiqtVirtualQGraphicsEffect : public virtual QGraphicsEffect {
public:
MiqtVirtualQGraphicsEffect(): QGraphicsEffect() {};
MiqtVirtualQGraphicsEffect(QObject* parent): QGraphicsEffect(parent) {};

virtual ~MiqtVirtualQGraphicsEffect() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QGraphicsEffect::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QGraphicsEffect_MetaObject(const_cast<MiqtVirtualQGraphicsEffect*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGraphicsEffect::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QGraphicsEffect::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGraphicsEffect_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGraphicsEffect::qt_metacast(param1);

	}
};
QGraphicsEffect* QGraphicsEffect_new() {
return new MiqtVirtualQGraphicsEffect();
}
QGraphicsEffect* QGraphicsEffect_new2(QObject* parent) {
return new MiqtVirtualQGraphicsEffect(parent);
}
void QGraphicsEffect_virtbase(QGraphicsEffect* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QGraphicsEffect_MetaObject(const QGraphicsEffect* self) {
	return (QMetaObject*) self->metaObject();
}
void* QGraphicsEffect_Metacast(QGraphicsEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QGraphicsEffect_Tr(const char* s) {
	QString _ret = QGraphicsEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QRectF* QGraphicsEffect_BoundingRectFor(const QGraphicsEffect* self, QRectF* sourceRect) {
	return new QRectF(self->boundingRectFor(*sourceRect));
}
QRectF* QGraphicsEffect_BoundingRect(const QGraphicsEffect* self) {
	return new QRectF(self->boundingRect());
}
bool QGraphicsEffect_IsEnabled(const QGraphicsEffect* self) {
	return self->isEnabled();
}
void QGraphicsEffect_SetEnabled(QGraphicsEffect* self, bool enable) {
	self->setEnabled(enable);
}
void QGraphicsEffect_Update(QGraphicsEffect* self) {
	self->update();
}
void QGraphicsEffect_EnabledChanged(QGraphicsEffect* self, bool enabled) {
	self->enabledChanged(enabled);
}
void QGraphicsEffect_connect_EnabledChanged(QGraphicsEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsEffect::connect(self, static_cast<void (QGraphicsEffect::*)(bool)>(&QGraphicsEffect::enabledChanged), self, [=](bool enabled) {
		bool sigval1 = enabled;
		miqt_exec_callback_QGraphicsEffect_EnabledChanged(slot, sigval1);
	});
}
struct miqt_string QGraphicsEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QGraphicsEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsEffect_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsEffect*>( (QGraphicsEffect*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QGraphicsEffect_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGraphicsEffect*)(self) )->virtualbase_MetaObject();
}
void QGraphicsEffect_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsEffect*>( (QGraphicsEffect*)(self) )->handle__Metacast = slot;
}
void* QGraphicsEffect_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGraphicsEffect*)(self) )->virtualbase_Metacast(param1);
}
void QGraphicsEffect_Delete(QGraphicsEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsEffect*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQGraphicsColorizeEffect : public virtual QGraphicsColorizeEffect {
public:
MiqtVirtualQGraphicsColorizeEffect(): QGraphicsColorizeEffect() {};
MiqtVirtualQGraphicsColorizeEffect(QObject* parent): QGraphicsColorizeEffect(parent) {};

virtual ~MiqtVirtualQGraphicsColorizeEffect() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QGraphicsColorizeEffect::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QGraphicsColorizeEffect_MetaObject(const_cast<MiqtVirtualQGraphicsColorizeEffect*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGraphicsColorizeEffect::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QGraphicsColorizeEffect::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGraphicsColorizeEffect_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGraphicsColorizeEffect::qt_metacast(param1);

	}
};
QGraphicsColorizeEffect* QGraphicsColorizeEffect_new() {
return new MiqtVirtualQGraphicsColorizeEffect();
}
QGraphicsColorizeEffect* QGraphicsColorizeEffect_new2(QObject* parent) {
return new MiqtVirtualQGraphicsColorizeEffect(parent);
}
void QGraphicsColorizeEffect_virtbase(QGraphicsColorizeEffect* src
, QGraphicsEffect** outptr_QGraphicsEffect
) {
*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(src);
}
QMetaObject* QGraphicsColorizeEffect_MetaObject(const QGraphicsColorizeEffect* self) {
	return (QMetaObject*) self->metaObject();
}
void* QGraphicsColorizeEffect_Metacast(QGraphicsColorizeEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QGraphicsColorizeEffect_Tr(const char* s) {
	QString _ret = QGraphicsColorizeEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QColor* QGraphicsColorizeEffect_Color(const QGraphicsColorizeEffect* self) {
	return new QColor(self->color());
}
double QGraphicsColorizeEffect_Strength(const QGraphicsColorizeEffect* self) {
	qreal _ret = self->strength();
	return static_cast<double>(_ret);
}
void QGraphicsColorizeEffect_SetColor(QGraphicsColorizeEffect* self, QColor* c) {
	self->setColor(*c);
}
void QGraphicsColorizeEffect_SetStrength(QGraphicsColorizeEffect* self, double strength) {
	self->setStrength(static_cast<qreal>(strength));
}
void QGraphicsColorizeEffect_ColorChanged(QGraphicsColorizeEffect* self, QColor* color) {
	self->colorChanged(*color);
}
void QGraphicsColorizeEffect_connect_ColorChanged(QGraphicsColorizeEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsColorizeEffect::connect(self, static_cast<void (QGraphicsColorizeEffect::*)(const QColor&)>(&QGraphicsColorizeEffect::colorChanged), self, [=](const QColor& color) {
		const QColor& color_ret = color;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&color_ret);
		miqt_exec_callback_QGraphicsColorizeEffect_ColorChanged(slot, sigval1);
	});
}
void QGraphicsColorizeEffect_StrengthChanged(QGraphicsColorizeEffect* self, double strength) {
	self->strengthChanged(static_cast<qreal>(strength));
}
void QGraphicsColorizeEffect_connect_StrengthChanged(QGraphicsColorizeEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsColorizeEffect::connect(self, static_cast<void (QGraphicsColorizeEffect::*)(qreal)>(&QGraphicsColorizeEffect::strengthChanged), self, [=](qreal strength) {
		qreal strength_ret = strength;
		double sigval1 = static_cast<double>(strength_ret);
		miqt_exec_callback_QGraphicsColorizeEffect_StrengthChanged(slot, sigval1);
	});
}
struct miqt_string QGraphicsColorizeEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsColorizeEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QGraphicsColorizeEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsColorizeEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsColorizeEffect_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsColorizeEffect*>( (QGraphicsColorizeEffect*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QGraphicsColorizeEffect_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGraphicsColorizeEffect*)(self) )->virtualbase_MetaObject();
}
void QGraphicsColorizeEffect_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsColorizeEffect*>( (QGraphicsColorizeEffect*)(self) )->handle__Metacast = slot;
}
void* QGraphicsColorizeEffect_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGraphicsColorizeEffect*)(self) )->virtualbase_Metacast(param1);
}
void QGraphicsColorizeEffect_Delete(QGraphicsColorizeEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsColorizeEffect*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQGraphicsBlurEffect : public virtual QGraphicsBlurEffect {
public:
MiqtVirtualQGraphicsBlurEffect(): QGraphicsBlurEffect() {};
MiqtVirtualQGraphicsBlurEffect(QObject* parent): QGraphicsBlurEffect(parent) {};

virtual ~MiqtVirtualQGraphicsBlurEffect() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QGraphicsBlurEffect::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QGraphicsBlurEffect_MetaObject(const_cast<MiqtVirtualQGraphicsBlurEffect*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGraphicsBlurEffect::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QGraphicsBlurEffect::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGraphicsBlurEffect_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGraphicsBlurEffect::qt_metacast(param1);

	}
};
QGraphicsBlurEffect* QGraphicsBlurEffect_new() {
return new MiqtVirtualQGraphicsBlurEffect();
}
QGraphicsBlurEffect* QGraphicsBlurEffect_new2(QObject* parent) {
return new MiqtVirtualQGraphicsBlurEffect(parent);
}
void QGraphicsBlurEffect_virtbase(QGraphicsBlurEffect* src
, QGraphicsEffect** outptr_QGraphicsEffect
) {
*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(src);
}
QMetaObject* QGraphicsBlurEffect_MetaObject(const QGraphicsBlurEffect* self) {
	return (QMetaObject*) self->metaObject();
}
void* QGraphicsBlurEffect_Metacast(QGraphicsBlurEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QGraphicsBlurEffect_Tr(const char* s) {
	QString _ret = QGraphicsBlurEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QRectF* QGraphicsBlurEffect_BoundingRectFor(const QGraphicsBlurEffect* self, QRectF* rect) {
	return new QRectF(self->boundingRectFor(*rect));
}
double QGraphicsBlurEffect_BlurRadius(const QGraphicsBlurEffect* self) {
	qreal _ret = self->blurRadius();
	return static_cast<double>(_ret);
}
BlurHints QGraphicsBlurEffect_BlurHints(const QGraphicsBlurEffect* self) {
	return self->blurHints();
}
void QGraphicsBlurEffect_SetBlurRadius(QGraphicsBlurEffect* self, double blurRadius) {
	self->setBlurRadius(static_cast<qreal>(blurRadius));
}
void QGraphicsBlurEffect_SetBlurHints(QGraphicsBlurEffect* self, BlurHints hints) {
	self->setBlurHints(hints);
}
void QGraphicsBlurEffect_BlurRadiusChanged(QGraphicsBlurEffect* self, double blurRadius) {
	self->blurRadiusChanged(static_cast<qreal>(blurRadius));
}
void QGraphicsBlurEffect_connect_BlurRadiusChanged(QGraphicsBlurEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsBlurEffect::connect(self, static_cast<void (QGraphicsBlurEffect::*)(qreal)>(&QGraphicsBlurEffect::blurRadiusChanged), self, [=](qreal blurRadius) {
		qreal blurRadius_ret = blurRadius;
		double sigval1 = static_cast<double>(blurRadius_ret);
		miqt_exec_callback_QGraphicsBlurEffect_BlurRadiusChanged(slot, sigval1);
	});
}
void QGraphicsBlurEffect_BlurHintsChanged(QGraphicsBlurEffect* self, BlurHints hints) {
	self->blurHintsChanged(hints);
}
void QGraphicsBlurEffect_connect_BlurHintsChanged(QGraphicsBlurEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsBlurEffect::connect(self, static_cast<void (QGraphicsBlurEffect::*)(BlurHints)>(&QGraphicsBlurEffect::blurHintsChanged), self, [=](BlurHints hints) {
		BlurHints sigval1 = hints;
		miqt_exec_callback_QGraphicsBlurEffect_BlurHintsChanged(slot, sigval1);
	});
}
struct miqt_string QGraphicsBlurEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsBlurEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QGraphicsBlurEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsBlurEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsBlurEffect_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsBlurEffect*>( (QGraphicsBlurEffect*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QGraphicsBlurEffect_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGraphicsBlurEffect*)(self) )->virtualbase_MetaObject();
}
void QGraphicsBlurEffect_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsBlurEffect*>( (QGraphicsBlurEffect*)(self) )->handle__Metacast = slot;
}
void* QGraphicsBlurEffect_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGraphicsBlurEffect*)(self) )->virtualbase_Metacast(param1);
}
void QGraphicsBlurEffect_Delete(QGraphicsBlurEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsBlurEffect*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQGraphicsDropShadowEffect : public virtual QGraphicsDropShadowEffect {
public:
MiqtVirtualQGraphicsDropShadowEffect(): QGraphicsDropShadowEffect() {};
MiqtVirtualQGraphicsDropShadowEffect(QObject* parent): QGraphicsDropShadowEffect(parent) {};

virtual ~MiqtVirtualQGraphicsDropShadowEffect() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QGraphicsDropShadowEffect::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QGraphicsDropShadowEffect_MetaObject(const_cast<MiqtVirtualQGraphicsDropShadowEffect*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGraphicsDropShadowEffect::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QGraphicsDropShadowEffect::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGraphicsDropShadowEffect_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGraphicsDropShadowEffect::qt_metacast(param1);

	}
};
QGraphicsDropShadowEffect* QGraphicsDropShadowEffect_new() {
return new MiqtVirtualQGraphicsDropShadowEffect();
}
QGraphicsDropShadowEffect* QGraphicsDropShadowEffect_new2(QObject* parent) {
return new MiqtVirtualQGraphicsDropShadowEffect(parent);
}
void QGraphicsDropShadowEffect_virtbase(QGraphicsDropShadowEffect* src
, QGraphicsEffect** outptr_QGraphicsEffect
) {
*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(src);
}
QMetaObject* QGraphicsDropShadowEffect_MetaObject(const QGraphicsDropShadowEffect* self) {
	return (QMetaObject*) self->metaObject();
}
void* QGraphicsDropShadowEffect_Metacast(QGraphicsDropShadowEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QGraphicsDropShadowEffect_Tr(const char* s) {
	QString _ret = QGraphicsDropShadowEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QRectF* QGraphicsDropShadowEffect_BoundingRectFor(const QGraphicsDropShadowEffect* self, QRectF* rect) {
	return new QRectF(self->boundingRectFor(*rect));
}
QPointF* QGraphicsDropShadowEffect_Offset(const QGraphicsDropShadowEffect* self) {
	return new QPointF(self->offset());
}
double QGraphicsDropShadowEffect_XOffset(const QGraphicsDropShadowEffect* self) {
	qreal _ret = self->xOffset();
	return static_cast<double>(_ret);
}
double QGraphicsDropShadowEffect_YOffset(const QGraphicsDropShadowEffect* self) {
	qreal _ret = self->yOffset();
	return static_cast<double>(_ret);
}
double QGraphicsDropShadowEffect_BlurRadius(const QGraphicsDropShadowEffect* self) {
	qreal _ret = self->blurRadius();
	return static_cast<double>(_ret);
}
QColor* QGraphicsDropShadowEffect_Color(const QGraphicsDropShadowEffect* self) {
	return new QColor(self->color());
}
void QGraphicsDropShadowEffect_SetOffset(QGraphicsDropShadowEffect* self, QPointF* ofs) {
	self->setOffset(*ofs);
}
void QGraphicsDropShadowEffect_SetOffset2(QGraphicsDropShadowEffect* self, double dx, double dy) {
	self->setOffset(static_cast<qreal>(dx), static_cast<qreal>(dy));
}
void QGraphicsDropShadowEffect_SetOffsetWithQreal(QGraphicsDropShadowEffect* self, double d) {
	self->setOffset(static_cast<qreal>(d));
}
void QGraphicsDropShadowEffect_SetXOffset(QGraphicsDropShadowEffect* self, double dx) {
	self->setXOffset(static_cast<qreal>(dx));
}
void QGraphicsDropShadowEffect_SetYOffset(QGraphicsDropShadowEffect* self, double dy) {
	self->setYOffset(static_cast<qreal>(dy));
}
void QGraphicsDropShadowEffect_SetBlurRadius(QGraphicsDropShadowEffect* self, double blurRadius) {
	self->setBlurRadius(static_cast<qreal>(blurRadius));
}
void QGraphicsDropShadowEffect_SetColor(QGraphicsDropShadowEffect* self, QColor* color) {
	self->setColor(*color);
}
void QGraphicsDropShadowEffect_OffsetChanged(QGraphicsDropShadowEffect* self, QPointF* offset) {
	self->offsetChanged(*offset);
}
void QGraphicsDropShadowEffect_connect_OffsetChanged(QGraphicsDropShadowEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsDropShadowEffect::connect(self, static_cast<void (QGraphicsDropShadowEffect::*)(const QPointF&)>(&QGraphicsDropShadowEffect::offsetChanged), self, [=](const QPointF& offset) {
		const QPointF& offset_ret = offset;
		// Cast returned reference into pointer
		QPointF* sigval1 = const_cast<QPointF*>(&offset_ret);
		miqt_exec_callback_QGraphicsDropShadowEffect_OffsetChanged(slot, sigval1);
	});
}
void QGraphicsDropShadowEffect_BlurRadiusChanged(QGraphicsDropShadowEffect* self, double blurRadius) {
	self->blurRadiusChanged(static_cast<qreal>(blurRadius));
}
void QGraphicsDropShadowEffect_connect_BlurRadiusChanged(QGraphicsDropShadowEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsDropShadowEffect::connect(self, static_cast<void (QGraphicsDropShadowEffect::*)(qreal)>(&QGraphicsDropShadowEffect::blurRadiusChanged), self, [=](qreal blurRadius) {
		qreal blurRadius_ret = blurRadius;
		double sigval1 = static_cast<double>(blurRadius_ret);
		miqt_exec_callback_QGraphicsDropShadowEffect_BlurRadiusChanged(slot, sigval1);
	});
}
void QGraphicsDropShadowEffect_ColorChanged(QGraphicsDropShadowEffect* self, QColor* color) {
	self->colorChanged(*color);
}
void QGraphicsDropShadowEffect_connect_ColorChanged(QGraphicsDropShadowEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsDropShadowEffect::connect(self, static_cast<void (QGraphicsDropShadowEffect::*)(const QColor&)>(&QGraphicsDropShadowEffect::colorChanged), self, [=](const QColor& color) {
		const QColor& color_ret = color;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&color_ret);
		miqt_exec_callback_QGraphicsDropShadowEffect_ColorChanged(slot, sigval1);
	});
}
struct miqt_string QGraphicsDropShadowEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsDropShadowEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QGraphicsDropShadowEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsDropShadowEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsDropShadowEffect_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsDropShadowEffect*>( (QGraphicsDropShadowEffect*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QGraphicsDropShadowEffect_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGraphicsDropShadowEffect*)(self) )->virtualbase_MetaObject();
}
void QGraphicsDropShadowEffect_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsDropShadowEffect*>( (QGraphicsDropShadowEffect*)(self) )->handle__Metacast = slot;
}
void* QGraphicsDropShadowEffect_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGraphicsDropShadowEffect*)(self) )->virtualbase_Metacast(param1);
}
void QGraphicsDropShadowEffect_Delete(QGraphicsDropShadowEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsDropShadowEffect*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQGraphicsOpacityEffect : public virtual QGraphicsOpacityEffect {
public:
MiqtVirtualQGraphicsOpacityEffect(): QGraphicsOpacityEffect() {};
MiqtVirtualQGraphicsOpacityEffect(QObject* parent): QGraphicsOpacityEffect(parent) {};

virtual ~MiqtVirtualQGraphicsOpacityEffect() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QGraphicsOpacityEffect::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QGraphicsOpacityEffect_MetaObject(const_cast<MiqtVirtualQGraphicsOpacityEffect*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGraphicsOpacityEffect::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QGraphicsOpacityEffect::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGraphicsOpacityEffect_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGraphicsOpacityEffect::qt_metacast(param1);

	}
};
QGraphicsOpacityEffect* QGraphicsOpacityEffect_new() {
return new MiqtVirtualQGraphicsOpacityEffect();
}
QGraphicsOpacityEffect* QGraphicsOpacityEffect_new2(QObject* parent) {
return new MiqtVirtualQGraphicsOpacityEffect(parent);
}
void QGraphicsOpacityEffect_virtbase(QGraphicsOpacityEffect* src
, QGraphicsEffect** outptr_QGraphicsEffect
) {
*outptr_QGraphicsEffect = static_cast<QGraphicsEffect*>(src);
}
QMetaObject* QGraphicsOpacityEffect_MetaObject(const QGraphicsOpacityEffect* self) {
	return (QMetaObject*) self->metaObject();
}
void* QGraphicsOpacityEffect_Metacast(QGraphicsOpacityEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QGraphicsOpacityEffect_Tr(const char* s) {
	QString _ret = QGraphicsOpacityEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
double QGraphicsOpacityEffect_Opacity(const QGraphicsOpacityEffect* self) {
	qreal _ret = self->opacity();
	return static_cast<double>(_ret);
}
QBrush* QGraphicsOpacityEffect_OpacityMask(const QGraphicsOpacityEffect* self) {
	return new QBrush(self->opacityMask());
}
void QGraphicsOpacityEffect_SetOpacity(QGraphicsOpacityEffect* self, double opacity) {
	self->setOpacity(static_cast<qreal>(opacity));
}
void QGraphicsOpacityEffect_SetOpacityMask(QGraphicsOpacityEffect* self, QBrush* mask) {
	self->setOpacityMask(*mask);
}
void QGraphicsOpacityEffect_OpacityChanged(QGraphicsOpacityEffect* self, double opacity) {
	self->opacityChanged(static_cast<qreal>(opacity));
}
void QGraphicsOpacityEffect_connect_OpacityChanged(QGraphicsOpacityEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsOpacityEffect::connect(self, static_cast<void (QGraphicsOpacityEffect::*)(qreal)>(&QGraphicsOpacityEffect::opacityChanged), self, [=](qreal opacity) {
		qreal opacity_ret = opacity;
		double sigval1 = static_cast<double>(opacity_ret);
		miqt_exec_callback_QGraphicsOpacityEffect_OpacityChanged(slot, sigval1);
	});
}
void QGraphicsOpacityEffect_OpacityMaskChanged(QGraphicsOpacityEffect* self, QBrush* mask) {
	self->opacityMaskChanged(*mask);
}
void QGraphicsOpacityEffect_connect_OpacityMaskChanged(QGraphicsOpacityEffect* self, intptr_t slot) {
	MiqtVirtualQGraphicsOpacityEffect::connect(self, static_cast<void (QGraphicsOpacityEffect::*)(const QBrush&)>(&QGraphicsOpacityEffect::opacityMaskChanged), self, [=](const QBrush& mask) {
		const QBrush& mask_ret = mask;
		// Cast returned reference into pointer
		QBrush* sigval1 = const_cast<QBrush*>(&mask_ret);
		miqt_exec_callback_QGraphicsOpacityEffect_OpacityMaskChanged(slot, sigval1);
	});
}
struct miqt_string QGraphicsOpacityEffect_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsOpacityEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QGraphicsOpacityEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsOpacityEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsOpacityEffect_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsOpacityEffect*>( (QGraphicsOpacityEffect*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QGraphicsOpacityEffect_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGraphicsOpacityEffect*)(self) )->virtualbase_MetaObject();
}
void QGraphicsOpacityEffect_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsOpacityEffect*>( (QGraphicsOpacityEffect*)(self) )->handle__Metacast = slot;
}
void* QGraphicsOpacityEffect_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGraphicsOpacityEffect*)(self) )->virtualbase_Metacast(param1);
}
void QGraphicsOpacityEffect_Delete(QGraphicsOpacityEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsOpacityEffect*>( self );
	} else {
		delete self;
	}
}
