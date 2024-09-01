#ifndef GEN_QVARIANTANIMATION_H
#define GEN_QVARIANTANIMATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QEasingCurve;
class QMetaObject;
class QObject;
class QVariant;
class QVariantAnimation;
#else
typedef struct QEasingCurve QEasingCurve;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QVariant QVariant;
typedef struct QVariantAnimation QVariantAnimation;
#endif

QVariantAnimation* QVariantAnimation_new();
QVariantAnimation* QVariantAnimation_new2(QObject* parent);
QMetaObject* QVariantAnimation_MetaObject(QVariantAnimation* self);
void QVariantAnimation_Tr(const char* s, char** _out, int* _out_Strlen);
void QVariantAnimation_TrUtf8(const char* s, char** _out, int* _out_Strlen);
QVariant* QVariantAnimation_StartValue(QVariantAnimation* self);
void QVariantAnimation_SetStartValue(QVariantAnimation* self, QVariant* value);
QVariant* QVariantAnimation_EndValue(QVariantAnimation* self);
void QVariantAnimation_SetEndValue(QVariantAnimation* self, QVariant* value);
QVariant* QVariantAnimation_KeyValueAt(QVariantAnimation* self, double step);
void QVariantAnimation_SetKeyValueAt(QVariantAnimation* self, double step, QVariant* value);
QVariant* QVariantAnimation_CurrentValue(QVariantAnimation* self);
int QVariantAnimation_Duration(QVariantAnimation* self);
void QVariantAnimation_SetDuration(QVariantAnimation* self, int msecs);
QEasingCurve* QVariantAnimation_EasingCurve(QVariantAnimation* self);
void QVariantAnimation_SetEasingCurve(QVariantAnimation* self, QEasingCurve* easing);
void QVariantAnimation_ValueChanged(QVariantAnimation* self, QVariant* value);
void QVariantAnimation_connect_ValueChanged(QVariantAnimation* self, void* slot);
void QVariantAnimation_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QVariantAnimation_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QVariantAnimation_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QVariantAnimation_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QVariantAnimation_Delete(QVariantAnimation* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif