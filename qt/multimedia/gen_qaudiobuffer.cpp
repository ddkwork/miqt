#include <QAudioBuffer>
#include <QAudioFormat>
#include <QByteArray>
#include <qaudiobuffer.h>
#include "gen_qaudiobuffer.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
} /* extern C */
#endif

QAudioBuffer* QAudioBuffer_new() {
	return new QAudioBuffer();
}

QAudioBuffer* QAudioBuffer_new2(QAudioBuffer* other) {
	return new QAudioBuffer(*other);
}

QAudioBuffer* QAudioBuffer_new3(struct miqt_string data, QAudioFormat* format) {
	QByteArray data_QByteArray(data.data, data.len);
	return new QAudioBuffer(data_QByteArray, *format);
}

QAudioBuffer* QAudioBuffer_new4(int numFrames, QAudioFormat* format) {
	return new QAudioBuffer(static_cast<int>(numFrames), *format);
}

QAudioBuffer* QAudioBuffer_new5(struct miqt_string data, QAudioFormat* format, long long startTime) {
	QByteArray data_QByteArray(data.data, data.len);
	return new QAudioBuffer(data_QByteArray, *format, static_cast<qint64>(startTime));
}

QAudioBuffer* QAudioBuffer_new6(int numFrames, QAudioFormat* format, long long startTime) {
	return new QAudioBuffer(static_cast<int>(numFrames), *format, static_cast<qint64>(startTime));
}

void QAudioBuffer_OperatorAssign(QAudioBuffer* self, QAudioBuffer* other) {
	self->operator=(*other);
}

bool QAudioBuffer_IsValid(const QAudioBuffer* self) {
	return self->isValid();
}

QAudioFormat* QAudioBuffer_Format(const QAudioBuffer* self) {
	return new QAudioFormat(self->format());
}

int QAudioBuffer_FrameCount(const QAudioBuffer* self) {
	return self->frameCount();
}

int QAudioBuffer_SampleCount(const QAudioBuffer* self) {
	return self->sampleCount();
}

int QAudioBuffer_ByteCount(const QAudioBuffer* self) {
	return self->byteCount();
}

long long QAudioBuffer_Duration(const QAudioBuffer* self) {
	qint64 _ret = self->duration();
	return static_cast<long long>(_ret);
}

long long QAudioBuffer_StartTime(const QAudioBuffer* self) {
	qint64 _ret = self->startTime();
	return static_cast<long long>(_ret);
}

const void* QAudioBuffer_ConstData(const QAudioBuffer* self) {
	return (const void*) self->constData();
}

const void* QAudioBuffer_Data(const QAudioBuffer* self) {
	return (const void*) self->data();
}

void* QAudioBuffer_Data2(QAudioBuffer* self) {
	return self->data();
}

void QAudioBuffer_Delete(QAudioBuffer* self) {
	delete self;
}

