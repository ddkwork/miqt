// +build ignore

#include <QAudioFormat>
#include <QIODevice>
#include <QIODeviceBase>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWaveDecoder>
#include <qwavedecoder.h>
#include "gen_qwavedecoder.h"
class MiqtVirtualQWaveDecoder : public virtual QWaveDecoder {
public:
MiqtVirtualQWaveDecoder(QIODevice* device): QWaveDecoder(device) {};
MiqtVirtualQWaveDecoder(QIODevice* device, const QAudioFormat& format): QWaveDecoder(device, format) {};
MiqtVirtualQWaveDecoder(QIODevice* device, QObject* parent): QWaveDecoder(device, parent) {};
MiqtVirtualQWaveDecoder(QIODevice* device, const QAudioFormat& format, QObject* parent): QWaveDecoder(device, format, parent) {};

virtual ~MiqtVirtualQWaveDecoder() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QWaveDecoder::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QWaveDecoder_MetaObject(const_cast<MiqtVirtualQWaveDecoder*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QWaveDecoder::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QWaveDecoder::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QWaveDecoder_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QWaveDecoder::qt_metacast(param1);

	}
};
QWaveDecoder* QWaveDecoder_new(QIODevice* device) {
return new MiqtVirtualQWaveDecoder(device);
}
QWaveDecoder* QWaveDecoder_new2(QIODevice* device, QAudioFormat* format) {
return new MiqtVirtualQWaveDecoder(device, *format);
}
QWaveDecoder* QWaveDecoder_new3(QIODevice* device, QObject* parent) {
return new MiqtVirtualQWaveDecoder(device, parent);
}
QWaveDecoder* QWaveDecoder_new4(QIODevice* device, QAudioFormat* format, QObject* parent) {
return new MiqtVirtualQWaveDecoder(device, *format, parent);
}
void QWaveDecoder_virtbase(QWaveDecoder* src
, QIODevice** outptr_QIODevice
) {
*outptr_QIODevice = static_cast<QIODevice*>(src);
}
QMetaObject* QWaveDecoder_MetaObject(const QWaveDecoder* self) {
	return (QMetaObject*) self->metaObject();
}
void* QWaveDecoder_Metacast(QWaveDecoder* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QWaveDecoder_Tr(const char* s) {
	QString _ret = QWaveDecoder::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QAudioFormat* QWaveDecoder_AudioFormat(const QWaveDecoder* self) {
	return new QAudioFormat(self->audioFormat());
}
QIODevice* QWaveDecoder_GetDevice(QWaveDecoder* self) {
	return self->getDevice();
}
int QWaveDecoder_Duration(const QWaveDecoder* self) {
	return self->duration();
}
long long QWaveDecoder_HeaderLength() {
	qint64 _ret = QWaveDecoder::headerLength();
	return static_cast<long long>(_ret);
}
bool QWaveDecoder_Open(QWaveDecoder* self, int mode) {
	return self->open(static_cast<QIODevice::OpenMode>(mode));
}
void QWaveDecoder_Close(QWaveDecoder* self) {
	self->close();
}
bool QWaveDecoder_Seek(QWaveDecoder* self, long long pos) {
	return self->seek(static_cast<qint64>(pos));
}
long long QWaveDecoder_Pos(const QWaveDecoder* self) {
	qint64 _ret = self->pos();
	return static_cast<long long>(_ret);
}
long long QWaveDecoder_Size(const QWaveDecoder* self) {
	qint64 _ret = self->size();
	return static_cast<long long>(_ret);
}
bool QWaveDecoder_IsSequential(const QWaveDecoder* self) {
	return self->isSequential();
}
long long QWaveDecoder_BytesAvailable(const QWaveDecoder* self) {
	qint64 _ret = self->bytesAvailable();
	return static_cast<long long>(_ret);
}
void QWaveDecoder_FormatKnown(QWaveDecoder* self) {
	self->formatKnown();
}
void QWaveDecoder_connect_FormatKnown(QWaveDecoder* self, intptr_t slot) {
	MiqtVirtualQWaveDecoder::connect(self, static_cast<void (QWaveDecoder::*)()>(&QWaveDecoder::formatKnown), self, [=]() {
		miqt_exec_callback_QWaveDecoder_FormatKnown(slot);
	});
}
void QWaveDecoder_ParsingError(QWaveDecoder* self) {
	self->parsingError();
}
void QWaveDecoder_connect_ParsingError(QWaveDecoder* self, intptr_t slot) {
	MiqtVirtualQWaveDecoder::connect(self, static_cast<void (QWaveDecoder::*)()>(&QWaveDecoder::parsingError), self, [=]() {
		miqt_exec_callback_QWaveDecoder_ParsingError(slot);
	});
}
struct miqt_string QWaveDecoder_Tr2(const char* s, const char* c) {
	QString _ret = QWaveDecoder::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QWaveDecoder_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWaveDecoder::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWaveDecoder_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWaveDecoder*>( (QWaveDecoder*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QWaveDecoder_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQWaveDecoder*)(self) )->virtualbase_MetaObject();
}
void QWaveDecoder_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWaveDecoder*>( (QWaveDecoder*)(self) )->handle__Metacast = slot;
}
void* QWaveDecoder_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQWaveDecoder*)(self) )->virtualbase_Metacast(param1);
}
void QWaveDecoder_Delete(QWaveDecoder* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWaveDecoder*>( self );
	} else {
		delete self;
	}
}
