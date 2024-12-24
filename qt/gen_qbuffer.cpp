// +build ignore

#include <QBuffer>
#include <QByteArray>
#include <QIODevice>
#include <QIODeviceBase>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qbuffer.h>
#include "gen_qbuffer.h"
class MiqtVirtualQBuffer : public virtual QBuffer {
public:
MiqtVirtualQBuffer(): QBuffer() {};
MiqtVirtualQBuffer(QObject* parent): QBuffer(parent) {};

virtual ~MiqtVirtualQBuffer() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QBuffer::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QBuffer_MetaObject(const_cast<MiqtVirtualQBuffer*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QBuffer::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QBuffer::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QBuffer_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QBuffer::qt_metacast(param1);

	}
};
QBuffer* QBuffer_new() {
return new MiqtVirtualQBuffer();
}
QBuffer* QBuffer_new2(QObject* parent) {
return new MiqtVirtualQBuffer(parent);
}
void QBuffer_virtbase(QBuffer* src
, QIODevice** outptr_QIODevice
) {
*outptr_QIODevice = static_cast<QIODevice*>(src);
}
QMetaObject* QBuffer_MetaObject(const QBuffer* self) {
	return (QMetaObject*) self->metaObject();
}
void* QBuffer_Metacast(QBuffer* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QBuffer_Tr(const char* s) {
	QString _ret = QBuffer::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QBuffer_Buffer(QBuffer* self) {
	QByteArray _qb = self->buffer();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
struct miqt_string QBuffer_Buffer2(const QBuffer* self) {
	const QByteArray _qb = self->buffer();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
void QBuffer_SetData(QBuffer* self, struct miqt_string data) {
	QByteArray data_QByteArray(data.data, data.len);
	self->setData(data_QByteArray);
}
void QBuffer_SetData2(QBuffer* self, const char* data, ptrdiff_t lenVal) {
	self->setData(data, (qsizetype)(lenVal));
}
struct miqt_string QBuffer_Data(const QBuffer* self) {
	const QByteArray _qb = self->data();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
bool QBuffer_Open(QBuffer* self, OpenMode openMode) {
	return self->open(openMode);
}
void QBuffer_Close(QBuffer* self) {
	self->close();
}
long long QBuffer_Size(const QBuffer* self) {
	qint64 _ret = self->size();
	return static_cast<long long>(_ret);
}
long long QBuffer_Pos(const QBuffer* self) {
	qint64 _ret = self->pos();
	return static_cast<long long>(_ret);
}
bool QBuffer_Seek(QBuffer* self, long long off) {
	return self->seek(static_cast<qint64>(off));
}
bool QBuffer_AtEnd(const QBuffer* self) {
	return self->atEnd();
}
bool QBuffer_CanReadLine(const QBuffer* self) {
	return self->canReadLine();
}
struct miqt_string QBuffer_Tr2(const char* s, const char* c) {
	QString _ret = QBuffer::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QBuffer_Tr3(const char* s, const char* c, int n) {
	QString _ret = QBuffer::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QBuffer_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQBuffer*>( (QBuffer*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QBuffer_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQBuffer*)(self) )->virtualbase_MetaObject();
}
void QBuffer_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQBuffer*>( (QBuffer*)(self) )->handle__Metacast = slot;
}
void* QBuffer_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQBuffer*)(self) )->virtualbase_Metacast(param1);
}
void QBuffer_Delete(QBuffer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQBuffer*>( self );
	} else {
		delete self;
	}
}
