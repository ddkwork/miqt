// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVideoFrame>
#include <QVideoSink>
#include <qvideosink.h>
#include "gen_qvideosink.h"
class MiqtVirtualQVideoSink : public virtual QVideoSink {
public:
MiqtVirtualQVideoSink(): QVideoSink() {};
MiqtVirtualQVideoSink(QObject* parent): QVideoSink(parent) {};

virtual ~MiqtVirtualQVideoSink() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QVideoSink::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QVideoSink_MetaObject(const_cast<MiqtVirtualQVideoSink*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QVideoSink::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QVideoSink::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QVideoSink_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QVideoSink::qt_metacast(param1);

	}
};
QVideoSink* QVideoSink_new() {
return new MiqtVirtualQVideoSink();
}
QVideoSink* QVideoSink_new2(QObject* parent) {
return new MiqtVirtualQVideoSink(parent);
}
void QVideoSink_virtbase(QVideoSink* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QVideoSink_MetaObject(const QVideoSink* self) {
	return (QMetaObject*) self->metaObject();
}
void* QVideoSink_Metacast(QVideoSink* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QVideoSink_Tr(const char* s) {
	QString _ret = QVideoSink::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QVideoSink_VideoSize(const QVideoSink* self) {
	return new QSize(self->videoSize());
}
struct miqt_string QVideoSink_SubtitleText(const QVideoSink* self) {
	QString _ret = self->subtitleText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QVideoSink_SetSubtitleText(QVideoSink* self, struct miqt_string subtitle) {
	QString subtitle_QString = QString::fromUtf8(subtitle.data, subtitle.len);
	self->setSubtitleText(subtitle_QString);
}
void QVideoSink_SetVideoFrame(QVideoSink* self, QVideoFrame* frame) {
	self->setVideoFrame(*frame);
}
QVideoFrame* QVideoSink_VideoFrame(const QVideoSink* self) {
	return new QVideoFrame(self->videoFrame());
}
void QVideoSink_VideoFrameChanged(const QVideoSink* self, QVideoFrame* frame) {
	self->videoFrameChanged(*frame);
}
void QVideoSink_connect_VideoFrameChanged(QVideoSink* self, intptr_t slot) {
	MiqtVirtualQVideoSink::connect(self, static_cast<void (QVideoSink::*)(const QVideoFrame&) const>(&QVideoSink::videoFrameChanged), self, [=](const QVideoFrame& frame) {
		const QVideoFrame& frame_ret = frame;
		// Cast returned reference into pointer
		QVideoFrame* sigval1 = const_cast<QVideoFrame*>(&frame_ret);
		miqt_exec_callback_QVideoSink_VideoFrameChanged(slot, sigval1);
	});
}
void QVideoSink_SubtitleTextChanged(const QVideoSink* self, struct miqt_string subtitleText) {
	QString subtitleText_QString = QString::fromUtf8(subtitleText.data, subtitleText.len);
	self->subtitleTextChanged(subtitleText_QString);
}
void QVideoSink_connect_SubtitleTextChanged(QVideoSink* self, intptr_t slot) {
	MiqtVirtualQVideoSink::connect(self, static_cast<void (QVideoSink::*)(const QString&) const>(&QVideoSink::subtitleTextChanged), self, [=](const QString& subtitleText) {
		const QString subtitleText_ret = subtitleText;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray subtitleText_b = subtitleText_ret.toUtf8();
		struct miqt_string subtitleText_ms;
		subtitleText_ms.len = subtitleText_b.length();
		subtitleText_ms.data = static_cast<char*>(malloc(subtitleText_ms.len));
		memcpy(subtitleText_ms.data, subtitleText_b.data(), subtitleText_ms.len);
		struct miqt_string sigval1 = subtitleText_ms;
		miqt_exec_callback_QVideoSink_SubtitleTextChanged(slot, sigval1);
	});
}
void QVideoSink_VideoSizeChanged(QVideoSink* self) {
	self->videoSizeChanged();
}
void QVideoSink_connect_VideoSizeChanged(QVideoSink* self, intptr_t slot) {
	MiqtVirtualQVideoSink::connect(self, static_cast<void (QVideoSink::*)()>(&QVideoSink::videoSizeChanged), self, [=]() {
		miqt_exec_callback_QVideoSink_VideoSizeChanged(slot);
	});
}
struct miqt_string QVideoSink_Tr2(const char* s, const char* c) {
	QString _ret = QVideoSink::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QVideoSink_Tr3(const char* s, const char* c, int n) {
	QString _ret = QVideoSink::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QVideoSink_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQVideoSink*>( (QVideoSink*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QVideoSink_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQVideoSink*)(self) )->virtualbase_MetaObject();
}
void QVideoSink_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQVideoSink*>( (QVideoSink*)(self) )->handle__Metacast = slot;
}
void* QVideoSink_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQVideoSink*)(self) )->virtualbase_Metacast(param1);
}
void QVideoSink_Delete(QVideoSink* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQVideoSink*>( self );
	} else {
		delete self;
	}
}