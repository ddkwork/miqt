// +build ignore

#include <QInputDevice>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QRect>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qinputdevice.h>
#include "gen_qinputdevice.h"

class MiqtVirtualQInputDevice : public virtual QInputDevice {
public:

	MiqtVirtualQInputDevice(): QInputDevice() {};
	MiqtVirtualQInputDevice(const QString& name, qint64 systemId, DeviceType typeVal): QInputDevice(name, systemId, typeVal) {};
	MiqtVirtualQInputDevice(QObject* parent): QInputDevice(parent) {};
	MiqtVirtualQInputDevice(const QString& name, qint64 systemId, DeviceType typeVal, const QString& seatName): QInputDevice(name, systemId, typeVal, seatName) {};
	MiqtVirtualQInputDevice(const QString& name, qint64 systemId, DeviceType typeVal, const QString& seatName, QObject* parent): QInputDevice(name, systemId, typeVal, seatName, parent) {};

	virtual ~MiqtVirtualQInputDevice() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QInputDevice::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QInputDevice_MetaObject(const_cast<MiqtVirtualQInputDevice*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QInputDevice::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QInputDevice::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QInputDevice_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QInputDevice::qt_metacast(param1);

	}

};

QInputDevice* QInputDevice_new() {
	return new MiqtVirtualQInputDevice();
}

QInputDevice* QInputDevice_new2(struct miqt_string name, long long systemId, DeviceType typeVal) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQInputDevice(name_QString, static_cast<qint64>(systemId), typeVal);
}

QInputDevice* QInputDevice_new3(QObject* parent) {
	return new MiqtVirtualQInputDevice(parent);
}

QInputDevice* QInputDevice_new4(struct miqt_string name, long long systemId, DeviceType typeVal, struct miqt_string seatName) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString seatName_QString = QString::fromUtf8(seatName.data, seatName.len);
	return new MiqtVirtualQInputDevice(name_QString, static_cast<qint64>(systemId), typeVal, seatName_QString);
}

QInputDevice* QInputDevice_new5(struct miqt_string name, long long systemId, DeviceType typeVal, struct miqt_string seatName, QObject* parent) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString seatName_QString = QString::fromUtf8(seatName.data, seatName.len);
	return new MiqtVirtualQInputDevice(name_QString, static_cast<qint64>(systemId), typeVal, seatName_QString, parent);
}

void QInputDevice_virtbase(QInputDevice* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QInputDevice_MetaObject(const QInputDevice* self) {
	return (QMetaObject*) self->metaObject();
}

void* QInputDevice_Metacast(QInputDevice* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QInputDevice_Tr(const char* s) {
	QString _ret = QInputDevice::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QInputDevice_Name(const QInputDevice* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

DeviceType QInputDevice_Type(const QInputDevice* self) {
	return self->type();
}

Capabilities QInputDevice_Capabilities(const QInputDevice* self) {
	return self->capabilities();
}

bool QInputDevice_HasCapability(const QInputDevice* self, Capability cap) {
	return self->hasCapability(cap);
}

long long QInputDevice_SystemId(const QInputDevice* self) {
	qint64 _ret = self->systemId();
	return static_cast<long long>(_ret);
}

struct miqt_string QInputDevice_SeatName(const QInputDevice* self) {
	QString _ret = self->seatName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QRect* QInputDevice_AvailableVirtualGeometry(const QInputDevice* self) {
	return new QRect(self->availableVirtualGeometry());
}

struct miqt_array /* of struct miqt_string */  QInputDevice_SeatNames() {
	QStringList _ret = QInputDevice::seatNames();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QInputDevice* */  QInputDevice_Devices() {
	QList<const QInputDevice *> _ret = QInputDevice::devices();
	// Convert QList<> from C++ memory to manually-managed C memory
	QInputDevice** _arr = static_cast<QInputDevice**>(malloc(sizeof(QInputDevice*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = (QInputDevice*) _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QInputDevice* QInputDevice_PrimaryKeyboard() {
	return (QInputDevice*) QInputDevice::primaryKeyboard();
}

bool QInputDevice_OperatorEqual(const QInputDevice* self, QInputDevice* other) {
	return (*self == *other);
}

void QInputDevice_AvailableVirtualGeometryChanged(QInputDevice* self, QRect* area) {
	self->availableVirtualGeometryChanged(*area);
}

void QInputDevice_connect_AvailableVirtualGeometryChanged(QInputDevice* self, intptr_t slot) {
	MiqtVirtualQInputDevice::connect(self, static_cast<void (QInputDevice::*)(QRect)>(&QInputDevice::availableVirtualGeometryChanged), self, [=](QRect area) {
		QRect* sigval1 = new QRect(area);
		miqt_exec_callback_QInputDevice_AvailableVirtualGeometryChanged(slot, sigval1);
	});
}

struct miqt_string QInputDevice_Tr2(const char* s, const char* c) {
	QString _ret = QInputDevice::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QInputDevice_Tr3(const char* s, const char* c, int n) {
	QString _ret = QInputDevice::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QInputDevice* QInputDevice_PrimaryKeyboard1(struct miqt_string seatName) {
	QString seatName_QString = QString::fromUtf8(seatName.data, seatName.len);
	return (QInputDevice*) QInputDevice::primaryKeyboard(seatName_QString);
}

void QInputDevice_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQInputDevice*>( (QInputDevice*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QInputDevice_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQInputDevice*)(self) )->virtualbase_MetaObject();
}

void QInputDevice_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQInputDevice*>( (QInputDevice*)(self) )->handle__Metacast = slot;
}

void* QInputDevice_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQInputDevice*)(self) )->virtualbase_Metacast(param1);
}

void QInputDevice_Delete(QInputDevice* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQInputDevice*>( self );
	} else {
		delete self;
	}
}

