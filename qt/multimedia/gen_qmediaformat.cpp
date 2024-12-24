// +build ignore

#include <QList>
#include <QMediaFormat>
#include <QMimeType>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qmediaformat.h>
#include "gen_qmediaformat.h"
QMediaFormat* QMediaFormat_new() {
return new QMediaFormat();
}
QMediaFormat* QMediaFormat_new2(QMediaFormat* other) {
return new QMediaFormat(*other);
}
QMediaFormat* QMediaFormat_new3(FileFormat format) {
return new QMediaFormat(format);
}
void QMediaFormat_OperatorAssign(QMediaFormat* self, QMediaFormat* other) {
	self->operator=(*other);
}
void QMediaFormat_Swap(QMediaFormat* self, QMediaFormat* other) {
	self->swap(*other);
}
FileFormat QMediaFormat_FileFormat(const QMediaFormat* self) {
	return self->fileFormat();
}
void QMediaFormat_SetFileFormat(QMediaFormat* self, FileFormat f) {
	self->setFileFormat(f);
}
void QMediaFormat_SetVideoCodec(QMediaFormat* self, VideoCodec codec) {
	self->setVideoCodec(codec);
}
VideoCodec QMediaFormat_VideoCodec(const QMediaFormat* self) {
	return self->videoCodec();
}
void QMediaFormat_SetAudioCodec(QMediaFormat* self, AudioCodec codec) {
	self->setAudioCodec(codec);
}
AudioCodec QMediaFormat_AudioCodec(const QMediaFormat* self) {
	return self->audioCodec();
}
bool QMediaFormat_IsSupported(const QMediaFormat* self, ConversionMode mode) {
	return self->isSupported(mode);
}
QMimeType* QMediaFormat_MimeType(const QMediaFormat* self) {
	return new QMimeType(self->mimeType());
}
struct miqt_array /* of FileFormat */  QMediaFormat_SupportedFileFormats(QMediaFormat* self, ConversionMode m) {
	QList<FileFormat> _ret = self->supportedFileFormats(m);
	// Convert QList<> from C++ memory to manually-managed C memory
	FileFormat* _arr = static_cast<FileFormat*>(malloc(sizeof(FileFormat) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_array /* of VideoCodec */  QMediaFormat_SupportedVideoCodecs(QMediaFormat* self, ConversionMode m) {
	QList<VideoCodec> _ret = self->supportedVideoCodecs(m);
	// Convert QList<> from C++ memory to manually-managed C memory
	VideoCodec* _arr = static_cast<VideoCodec*>(malloc(sizeof(VideoCodec) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_array /* of AudioCodec */  QMediaFormat_SupportedAudioCodecs(QMediaFormat* self, ConversionMode m) {
	QList<AudioCodec> _ret = self->supportedAudioCodecs(m);
	// Convert QList<> from C++ memory to manually-managed C memory
	AudioCodec* _arr = static_cast<AudioCodec*>(malloc(sizeof(AudioCodec) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_string QMediaFormat_FileFormatName(FileFormat fileFormat) {
	QString _ret = QMediaFormat::fileFormatName(fileFormat);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMediaFormat_AudioCodecName(AudioCodec codec) {
	QString _ret = QMediaFormat::audioCodecName(codec);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMediaFormat_VideoCodecName(VideoCodec codec) {
	QString _ret = QMediaFormat::videoCodecName(codec);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMediaFormat_FileFormatDescription(int fileFormat) {
	QString _ret = QMediaFormat::fileFormatDescription(static_cast<QMediaFormat::FileFormat>(fileFormat));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMediaFormat_AudioCodecDescription(int codec) {
	QString _ret = QMediaFormat::audioCodecDescription(static_cast<QMediaFormat::AudioCodec>(codec));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMediaFormat_VideoCodecDescription(int codec) {
	QString _ret = QMediaFormat::videoCodecDescription(static_cast<QMediaFormat::VideoCodec>(codec));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QMediaFormat_OperatorEqual(const QMediaFormat* self, QMediaFormat* other) {
	return (*self == *other);
}
bool QMediaFormat_OperatorNotEqual(const QMediaFormat* self, QMediaFormat* other) {
	return (*self != *other);
}
void QMediaFormat_ResolveForEncoding(QMediaFormat* self, ResolveFlags flags) {
	self->resolveForEncoding(flags);
}
void QMediaFormat_Delete(QMediaFormat* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QMediaFormat*>( self );
	} else {
		delete self;
	}
}
