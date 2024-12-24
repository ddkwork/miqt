#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QMEDIAPLAYER_H
#define MIQT_QT_MULTIMEDIA_GEN_QMEDIAPLAYER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioBufferOutput QAudioBufferOutput;
typedef struct QAudioOutput QAudioOutput;
typedef struct QIODevice QIODevice;
typedef struct QMediaMetaData QMediaMetaData;
typedef struct QMediaPlayer QMediaPlayer;
typedef struct QMediaTimeRange QMediaTimeRange;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QUrl QUrl;
typedef struct QVideoSink QVideoSink;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMediaPlayer* QMediaPlayer_new();
extern __declspec(dllexport) 
QMediaPlayer* QMediaPlayer_new2(QObject* parent);
extern __declspec(dllexport) 
void QMediaPlayer_virtbase(QMediaPlayer* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QMediaPlayer_MetaObject(const QMediaPlayer* self);
extern __declspec(dllexport) 
void* QMediaPlayer_Metacast(QMediaPlayer* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMediaPlayer_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_array /* of QMediaMetaData* */  QMediaPlayer_AudioTracks(const QMediaPlayer* self);
extern __declspec(dllexport) 
struct miqt_array /* of QMediaMetaData* */  QMediaPlayer_VideoTracks(const QMediaPlayer* self);
extern __declspec(dllexport) 
struct miqt_array /* of QMediaMetaData* */  QMediaPlayer_SubtitleTracks(const QMediaPlayer* self);
extern __declspec(dllexport) 
int QMediaPlayer_ActiveAudioTrack(const QMediaPlayer* self);
extern __declspec(dllexport) 
int QMediaPlayer_ActiveVideoTrack(const QMediaPlayer* self);
extern __declspec(dllexport) 
int QMediaPlayer_ActiveSubtitleTrack(const QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_SetActiveAudioTrack(QMediaPlayer* self, int index);
extern __declspec(dllexport) 
void QMediaPlayer_SetActiveVideoTrack(QMediaPlayer* self, int index);
extern __declspec(dllexport) 
void QMediaPlayer_SetActiveSubtitleTrack(QMediaPlayer* self, int index);
extern __declspec(dllexport) 
void QMediaPlayer_SetAudioBufferOutput(QMediaPlayer* self, QAudioBufferOutput* output);
extern __declspec(dllexport) 
QAudioBufferOutput* QMediaPlayer_AudioBufferOutput(const QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_SetAudioOutput(QMediaPlayer* self, QAudioOutput* output);
extern __declspec(dllexport) 
QAudioOutput* QMediaPlayer_AudioOutput(const QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_SetVideoOutput(QMediaPlayer* self, QObject* videoOutput);
extern __declspec(dllexport) 
QObject* QMediaPlayer_VideoOutput(const QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_SetVideoSink(QMediaPlayer* self, QVideoSink* sink);
extern __declspec(dllexport) 
QVideoSink* QMediaPlayer_VideoSink(const QMediaPlayer* self);
extern __declspec(dllexport) 
QUrl* QMediaPlayer_Source(const QMediaPlayer* self);
extern __declspec(dllexport) 
QIODevice* QMediaPlayer_SourceDevice(const QMediaPlayer* self);
extern __declspec(dllexport) 
PlaybackState QMediaPlayer_PlaybackState(const QMediaPlayer* self);
extern __declspec(dllexport) 
MediaStatus QMediaPlayer_MediaStatus(const QMediaPlayer* self);
extern __declspec(dllexport) 
long long QMediaPlayer_Duration(const QMediaPlayer* self);
extern __declspec(dllexport) 
long long QMediaPlayer_Position(const QMediaPlayer* self);
extern __declspec(dllexport) 
bool QMediaPlayer_HasAudio(const QMediaPlayer* self);
extern __declspec(dllexport) 
bool QMediaPlayer_HasVideo(const QMediaPlayer* self);
extern __declspec(dllexport) 
float QMediaPlayer_BufferProgress(const QMediaPlayer* self);
extern __declspec(dllexport) 
QMediaTimeRange* QMediaPlayer_BufferedTimeRange(const QMediaPlayer* self);
extern __declspec(dllexport) 
bool QMediaPlayer_IsSeekable(const QMediaPlayer* self);
extern __declspec(dllexport) 
double QMediaPlayer_PlaybackRate(const QMediaPlayer* self);
extern __declspec(dllexport) 
bool QMediaPlayer_IsPlaying(const QMediaPlayer* self);
extern __declspec(dllexport) 
int QMediaPlayer_Loops(const QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_SetLoops(QMediaPlayer* self, int loops);
extern __declspec(dllexport) 
Error QMediaPlayer_Error(const QMediaPlayer* self);
extern __declspec(dllexport) 
struct miqt_string QMediaPlayer_ErrorString(const QMediaPlayer* self);
extern __declspec(dllexport) 
bool QMediaPlayer_IsAvailable(const QMediaPlayer* self);
extern __declspec(dllexport) 
QMediaMetaData* QMediaPlayer_MetaData(const QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_Play(QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_Pause(QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_Stop(QMediaPlayer* self);
extern __declspec(dllexport) 
void QMediaPlayer_SetPosition(QMediaPlayer* self, long long position);
extern __declspec(dllexport) 
void QMediaPlayer_SetPlaybackRate(QMediaPlayer* self, double rate);
extern __declspec(dllexport) 
void QMediaPlayer_SetSource(QMediaPlayer* self, QUrl* source);
extern __declspec(dllexport) 
void QMediaPlayer_SetSourceDevice(QMediaPlayer* self, QIODevice* device);
extern __declspec(dllexport) 
void QMediaPlayer_SourceChanged(QMediaPlayer* self, QUrl* media);
void QMediaPlayer_connect_SourceChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_PlaybackStateChanged(QMediaPlayer* self, int newState);
void QMediaPlayer_connect_PlaybackStateChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_MediaStatusChanged(QMediaPlayer* self, int status);
void QMediaPlayer_connect_MediaStatusChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_DurationChanged(QMediaPlayer* self, long long duration);
void QMediaPlayer_connect_DurationChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_PositionChanged(QMediaPlayer* self, long long position);
void QMediaPlayer_connect_PositionChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_HasAudioChanged(QMediaPlayer* self, bool available);
void QMediaPlayer_connect_HasAudioChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_HasVideoChanged(QMediaPlayer* self, bool videoAvailable);
void QMediaPlayer_connect_HasVideoChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_BufferProgressChanged(QMediaPlayer* self, float progress);
void QMediaPlayer_connect_BufferProgressChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_SeekableChanged(QMediaPlayer* self, bool seekable);
void QMediaPlayer_connect_SeekableChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_PlayingChanged(QMediaPlayer* self, bool playing);
void QMediaPlayer_connect_PlayingChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_PlaybackRateChanged(QMediaPlayer* self, double rate);
void QMediaPlayer_connect_PlaybackRateChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_LoopsChanged(QMediaPlayer* self);
void QMediaPlayer_connect_LoopsChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_MetaDataChanged(QMediaPlayer* self);
void QMediaPlayer_connect_MetaDataChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_VideoOutputChanged(QMediaPlayer* self);
void QMediaPlayer_connect_VideoOutputChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_AudioOutputChanged(QMediaPlayer* self);
void QMediaPlayer_connect_AudioOutputChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_AudioBufferOutputChanged(QMediaPlayer* self);
void QMediaPlayer_connect_AudioBufferOutputChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_TracksChanged(QMediaPlayer* self);
void QMediaPlayer_connect_TracksChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_ActiveTracksChanged(QMediaPlayer* self);
void QMediaPlayer_connect_ActiveTracksChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_ErrorChanged(QMediaPlayer* self);
void QMediaPlayer_connect_ErrorChanged(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaPlayer_ErrorOccurred(QMediaPlayer* self, int error, struct miqt_string errorString);
void QMediaPlayer_connect_ErrorOccurred(QMediaPlayer* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QMediaPlayer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMediaPlayer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMediaPlayer_SetSourceDevice2(QMediaPlayer* self, QIODevice* device, QUrl* sourceUrl);
extern __declspec(dllexport) 
void QMediaPlayer_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMediaPlayer_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMediaPlayer_override_virtual_Metacast(void* self, intptr_t slot);
void* QMediaPlayer_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMediaPlayer_Delete(QMediaPlayer* self, bool isSubclass);

}
