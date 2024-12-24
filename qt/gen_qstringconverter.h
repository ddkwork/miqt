#pragma once
#ifndef MIQT_QT_GEN_QSTRINGCONVERTER_H
#define MIQT_QT_GEN_QSTRINGCONVERTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QByteArrayView QByteArrayView;
typedef struct QChar QChar;
typedef struct QStringConverter QStringConverter;
typedef struct QStringConverterBase QStringConverterBase;
typedef struct QStringDecoder QStringDecoder;
typedef struct QStringEncoder QStringEncoder;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStringEncoder* QStringEncoder_new();
extern __declspec(dllexport) 
QStringEncoder* QStringEncoder_new2(Encoding encoding);
extern __declspec(dllexport) 
QStringEncoder* QStringEncoder_new3(QAnyStringView* name);
extern __declspec(dllexport) 
QStringEncoder* QStringEncoder_new4(Encoding encoding, Flags flags);
extern __declspec(dllexport) 
QStringEncoder* QStringEncoder_new5(QAnyStringView* name, Flags flags);
extern __declspec(dllexport) 
void QStringEncoder_virtbase(QStringEncoder* src
, QStringConverter** outptr_QStringConverter
);
extern __declspec(dllexport) 
ptrdiff_t QStringEncoder_RequiredSpace(const QStringEncoder* self, ptrdiff_t inputLength);
extern __declspec(dllexport) 
void QStringEncoder_Delete(QStringEncoder* self, bool isSubclass);

extern __declspec(dllexport) 
QStringDecoder* QStringDecoder_new(Encoding encoding);
extern __declspec(dllexport) 
QStringDecoder* QStringDecoder_new2();
extern __declspec(dllexport) 
QStringDecoder* QStringDecoder_new3(QAnyStringView* name);
extern __declspec(dllexport) 
QStringDecoder* QStringDecoder_new4(Encoding encoding, Flags flags);
extern __declspec(dllexport) 
QStringDecoder* QStringDecoder_new5(QAnyStringView* name, Flags f);
extern __declspec(dllexport) 
void QStringDecoder_virtbase(QStringDecoder* src
, QStringConverter** outptr_QStringConverter
);
extern __declspec(dllexport) 
ptrdiff_t QStringDecoder_RequiredSpace(const QStringDecoder* self, ptrdiff_t inputLength);
extern __declspec(dllexport) 
QChar* QStringDecoder_AppendToBuffer(QStringDecoder* self, QChar* out, QByteArrayView* ba);
extern __declspec(dllexport) 
QStringDecoder* QStringDecoder_DecoderForHtml(QByteArrayView* data);
extern __declspec(dllexport) 
void QStringDecoder_Delete(QStringDecoder* self, bool isSubclass);

}
