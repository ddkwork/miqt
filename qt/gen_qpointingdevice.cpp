// +build ignore

#include <QEventPoint>
#include <QInputDevice>
#include <QMetaObject>
#include <QObject>
#include <QPointerEvent>
#include <QPointingDevice>
#include <QPointingDeviceUniqueId>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qpointingdevice.h>
#include "gen_qpointingdevice.h"
QPointingDeviceUniqueId* QPointingDeviceUniqueId_new() {
return new QPointingDeviceUniqueId();
}
QPointingDeviceUniqueId* QPointingDeviceUniqueId_new2(QPointingDeviceUniqueId* param1) {
return new QPointingDeviceUniqueId(*param1);
}
QPointingDeviceUniqueId* QPointingDeviceUniqueId_FromNumericId(long long id) {
	return new QPointingDeviceUniqueId(QPointingDeviceUniqueId::fromNumericId(static_cast<qint64>(id)));
}
bool QPointingDeviceUniqueId_IsValid(const QPointingDeviceUniqueId* self) {
	return self->isValid();
}
long long QPointingDeviceUniqueId_NumericId(const QPointingDeviceUniqueId* self) {
	qint64 _ret = self->numericId();
	return static_cast<long long>(_ret);
}
void QPointingDeviceUniqueId_Delete(QPointingDeviceUniqueId* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPointingDeviceUniqueId*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQPointingDevice : public virtual QPointingDevice {
public:
MiqtVirtualQPointingDevice(): QPointingDevice() {};
MiqtVirtualQPointingDevice(const QString& name, qint64 systemId, QInputDevice::DeviceType devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount): QPointingDevice(name, systemId, devType, pType, caps, maxPoints, buttonCount) {};
MiqtVirtualQPointingDevice(QObject* parent): QPointingDevice(parent) {};
MiqtVirtualQPointingDevice(const QString& name, qint64 systemId, QInputDevice::DeviceType devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, const QString& seatName): QPointingDevice(name, systemId, devType, pType, caps, maxPoints, buttonCount, seatName) {};
MiqtVirtualQPointingDevice(const QString& name, qint64 systemId, QInputDevice::DeviceType devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, const QString& seatName, QPointingDeviceUniqueId uniqueId): QPointingDevice(name, systemId, devType, pType, caps, maxPoints, buttonCount, seatName, uniqueId) {};
MiqtVirtualQPointingDevice(const QString& name, qint64 systemId, QInputDevice::DeviceType devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, const QString& seatName, QPointingDeviceUniqueId uniqueId, QObject* parent): QPointingDevice(name, systemId, devType, pType, caps, maxPoints, buttonCount, seatName, uniqueId, parent) {};

virtual ~MiqtVirtualQPointingDevice() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QPointingDevice::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QPointingDevice_MetaObject(const_cast<MiqtVirtualQPointingDevice*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPointingDevice::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QPointingDevice::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPointingDevice_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPointingDevice::qt_metacast(param1);

	}
};
QPointingDevice* QPointingDevice_new() {
return new MiqtVirtualQPointingDevice();
}
QPointingDevice* QPointingDevice_new2(struct miqt_string name, long long systemId, int devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount) {
QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQPointingDevice(name_QString, static_cast<qint64>(systemId), static_cast<QInputDevice::DeviceType>(devType), pType, caps, static_cast<int>(maxPoints), static_cast<int>(buttonCount));
}
QPointingDevice* QPointingDevice_new3(QObject* parent) {
return new MiqtVirtualQPointingDevice(parent);
}
QPointingDevice* QPointingDevice_new4(struct miqt_string name, long long systemId, int devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, struct miqt_string seatName) {
QString name_QString = QString::fromUtf8(name.data, name.len);
	QString seatName_QString = QString::fromUtf8(seatName.data, seatName.len);
	return new MiqtVirtualQPointingDevice(name_QString, static_cast<qint64>(systemId), static_cast<QInputDevice::DeviceType>(devType), pType, caps, static_cast<int>(maxPoints), static_cast<int>(buttonCount), seatName_QString);
}
QPointingDevice* QPointingDevice_new5(struct miqt_string name, long long systemId, int devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, struct miqt_string seatName, QPointingDeviceUniqueId* uniqueId) {
QString name_QString = QString::fromUtf8(name.data, name.len);
	QString seatName_QString = QString::fromUtf8(seatName.data, seatName.len);
	return new MiqtVirtualQPointingDevice(name_QString, static_cast<qint64>(systemId), static_cast<QInputDevice::DeviceType>(devType), pType, caps, static_cast<int>(maxPoints), static_cast<int>(buttonCount), seatName_QString, *uniqueId);
}
QPointingDevice* QPointingDevice_new6(struct miqt_string name, long long systemId, int devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, struct miqt_string seatName, QPointingDeviceUniqueId* uniqueId, QObject* parent) {
QString name_QString = QString::fromUtf8(name.data, name.len);
	QString seatName_QString = QString::fromUtf8(seatName.data, seatName.len);
	return new MiqtVirtualQPointingDevice(name_QString, static_cast<qint64>(systemId), static_cast<QInputDevice::DeviceType>(devType), pType, caps, static_cast<int>(maxPoints), static_cast<int>(buttonCount), seatName_QString, *uniqueId, parent);
}
void QPointingDevice_virtbase(QPointingDevice* src
, QInputDevice** outptr_QInputDevice
) {
*outptr_QInputDevice = static_cast<QInputDevice*>(src);
}
QMetaObject* QPointingDevice_MetaObject(const QPointingDevice* self) {
	return (QMetaObject*) self->metaObject();
}
void* QPointingDevice_Metacast(QPointingDevice* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QPointingDevice_Tr(const char* s) {
	QString _ret = QPointingDevice::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPointingDevice_SetType(QPointingDevice* self, DeviceType devType) {
	self->setType(devType);
}
void QPointingDevice_SetCapabilities(QPointingDevice* self, int caps) {
	self->setCapabilities(static_cast<QInputDevice::Capabilities>(caps));
}
void QPointingDevice_SetMaximumTouchPoints(QPointingDevice* self, int c) {
	self->setMaximumTouchPoints(static_cast<int>(c));
}
PointerType QPointingDevice_PointerType(const QPointingDevice* self) {
	return self->pointerType();
}
int QPointingDevice_MaximumPoints(const QPointingDevice* self) {
	return self->maximumPoints();
}
int QPointingDevice_ButtonCount(const QPointingDevice* self) {
	return self->buttonCount();
}
QPointingDeviceUniqueId* QPointingDevice_UniqueId(const QPointingDevice* self) {
	return new QPointingDeviceUniqueId(self->uniqueId());
}
QPointingDevice* QPointingDevice_PrimaryPointingDevice() {
	return (QPointingDevice*) QPointingDevice::primaryPointingDevice();
}
bool QPointingDevice_OperatorEqual(const QPointingDevice* self, QPointingDevice* other) {
	return (*self == *other);
}
void QPointingDevice_GrabChanged(const QPointingDevice* self, QObject* grabber, GrabTransition transition, QPointerEvent* event, QEventPoint* point) {
	self->grabChanged(grabber, transition, event, *point);
}
void QPointingDevice_connect_GrabChanged(QPointingDevice* self, intptr_t slot) {
	MiqtVirtualQPointingDevice::connect(self, static_cast<void (QPointingDevice::*)(QObject*, GrabTransition, const QPointerEvent*, const QEventPoint&) const>(&QPointingDevice::grabChanged), self, [=](QObject* grabber, GrabTransition transition, const QPointerEvent* event, const QEventPoint& point) {
		QObject* sigval1 = grabber;
		GrabTransition sigval2 = transition;
		QPointerEvent* sigval3 = (QPointerEvent*) event;
		const QEventPoint& point_ret = point;
		// Cast returned reference into pointer
		QEventPoint* sigval4 = const_cast<QEventPoint*>(&point_ret);
		miqt_exec_callback_QPointingDevice_GrabChanged(slot, sigval1, sigval2, sigval3, sigval4);
	});
}
struct miqt_string QPointingDevice_Tr2(const char* s, const char* c) {
	QString _ret = QPointingDevice::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QPointingDevice_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPointingDevice::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QPointingDevice* QPointingDevice_PrimaryPointingDevice1(struct miqt_string seatName) {
	QString seatName_QString = QString::fromUtf8(seatName.data, seatName.len);
	return (QPointingDevice*) QPointingDevice::primaryPointingDevice(seatName_QString);
}
void QPointingDevice_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPointingDevice*>( (QPointingDevice*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QPointingDevice_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPointingDevice*)(self) )->virtualbase_MetaObject();
}
void QPointingDevice_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPointingDevice*>( (QPointingDevice*)(self) )->handle__Metacast = slot;
}
void* QPointingDevice_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPointingDevice*)(self) )->virtualbase_Metacast(param1);
}
void QPointingDevice_Delete(QPointingDevice* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPointingDevice*>( self );
	} else {
		delete self;
	}
}
