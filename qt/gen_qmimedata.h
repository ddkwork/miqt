#ifndef GEN_QMIMEDATA_H
#define GEN_QMIMEDATA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QMetaObject;
class QMimeData;
class QUrl;
class QVariant;
#else
typedef struct QByteArray QByteArray;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;
#endif

QMimeData* QMimeData_new();
QMetaObject* QMimeData_MetaObject(const QMimeData* self);
void QMimeData_Tr(const char* s, char** _out, int* _out_Strlen);
void QMimeData_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QMimeData_Urls(const QMimeData* self, QUrl*** _out, size_t* _out_len);
void QMimeData_SetUrls(QMimeData* self, QUrl** urls, size_t urls_len);
bool QMimeData_HasUrls(const QMimeData* self);
void QMimeData_Text(const QMimeData* self, char** _out, int* _out_Strlen);
void QMimeData_SetText(QMimeData* self, const char* text, size_t text_Strlen);
bool QMimeData_HasText(const QMimeData* self);
void QMimeData_Html(const QMimeData* self, char** _out, int* _out_Strlen);
void QMimeData_SetHtml(QMimeData* self, const char* html, size_t html_Strlen);
bool QMimeData_HasHtml(const QMimeData* self);
QVariant* QMimeData_ImageData(const QMimeData* self);
void QMimeData_SetImageData(QMimeData* self, QVariant* image);
bool QMimeData_HasImage(const QMimeData* self);
QVariant* QMimeData_ColorData(const QMimeData* self);
void QMimeData_SetColorData(QMimeData* self, QVariant* color);
bool QMimeData_HasColor(const QMimeData* self);
QByteArray* QMimeData_Data(const QMimeData* self, const char* mimetype, size_t mimetype_Strlen);
void QMimeData_SetData(QMimeData* self, const char* mimetype, size_t mimetype_Strlen, QByteArray* data);
void QMimeData_RemoveFormat(QMimeData* self, const char* mimetype, size_t mimetype_Strlen);
bool QMimeData_HasFormat(const QMimeData* self, const char* mimetype, size_t mimetype_Strlen);
void QMimeData_Formats(const QMimeData* self, char*** _out, int** _out_Lengths, size_t* _out_len);
void QMimeData_Clear(QMimeData* self);
void QMimeData_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QMimeData_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QMimeData_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QMimeData_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QMimeData_Delete(QMimeData* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
