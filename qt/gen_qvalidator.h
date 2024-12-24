#pragma once
#ifndef MIQT_QT_GEN_QVALIDATOR_H
#define MIQT_QT_GEN_QVALIDATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QDoubleValidator;
class QEvent;
class QIntValidator;
class QLocale;
class QMetaMethod;
class QMetaObject;
class QObject;
class QRegularExpression;
class QRegularExpressionValidator;
class QTimerEvent;
class QValidator;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QDoubleValidator QDoubleValidator;
typedef struct QEvent QEvent;
typedef struct QIntValidator QIntValidator;
typedef struct QLocale QLocale;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRegularExpression QRegularExpression;
typedef struct QRegularExpressionValidator QRegularExpressionValidator;
typedef struct QTimerEvent QTimerEvent;
typedef struct QValidator QValidator;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QValidator* QValidator_new();
extern __declspec(dllexport) QValidator* QValidator_new2(QObject* parent);
extern __declspec(dllexport) void QValidator_virtbase(QValidator* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QValidator_MetaObject(const QValidator* self);
extern __declspec(dllexport) void* QValidator_Metacast(QValidator* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QValidator_Tr(const char* s);
extern __declspec(dllexport) void QValidator_SetLocale(QValidator* self, QLocale* locale);
extern __declspec(dllexport) QLocale* QValidator_Locale(const QValidator* self);
extern __declspec(dllexport) State QValidator_Validate(const QValidator* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) void QValidator_Fixup(const QValidator* self, struct miqt_string param1);
extern __declspec(dllexport) void QValidator_Changed(QValidator* self);
void QValidator_connect_Changed(QValidator* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QValidator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QValidator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QValidator_override_virtual_Validate(void* self, intptr_t slot);
State QValidator_virtualbase_Validate(const void* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) void QValidator_override_virtual_Fixup(void* self, intptr_t slot);
void QValidator_virtualbase_Fixup(const void* self, struct miqt_string param1);
extern __declspec(dllexport) void QValidator_override_virtual_Event(void* self, intptr_t slot);
bool QValidator_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QValidator_override_virtual_EventFilter(void* self, intptr_t slot);
bool QValidator_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QValidator_override_virtual_TimerEvent(void* self, intptr_t slot);
void QValidator_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QValidator_override_virtual_ChildEvent(void* self, intptr_t slot);
void QValidator_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QValidator_override_virtual_CustomEvent(void* self, intptr_t slot);
void QValidator_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QValidator_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QValidator_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QValidator_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QValidator_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QValidator_Delete(QValidator* self, bool isSubclass);

extern __declspec(dllexport) QIntValidator* QIntValidator_new();
extern __declspec(dllexport) QIntValidator* QIntValidator_new2(int bottom, int top);
extern __declspec(dllexport) QIntValidator* QIntValidator_new3(QObject* parent);
extern __declspec(dllexport) QIntValidator* QIntValidator_new4(int bottom, int top, QObject* parent);
extern __declspec(dllexport) void QIntValidator_virtbase(QIntValidator* src, QValidator** outptr_QValidator);
extern __declspec(dllexport) QMetaObject* QIntValidator_MetaObject(const QIntValidator* self);
extern __declspec(dllexport) void* QIntValidator_Metacast(QIntValidator* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QIntValidator_Tr(const char* s);
extern __declspec(dllexport) int QIntValidator_Validate(const QIntValidator* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) void QIntValidator_Fixup(const QIntValidator* self, struct miqt_string input);
extern __declspec(dllexport) void QIntValidator_SetBottom(QIntValidator* self, int bottom);
extern __declspec(dllexport) void QIntValidator_SetTop(QIntValidator* self, int top);
extern __declspec(dllexport) void QIntValidator_SetRange(QIntValidator* self, int bottom, int top);
extern __declspec(dllexport) int QIntValidator_Bottom(const QIntValidator* self);
extern __declspec(dllexport) int QIntValidator_Top(const QIntValidator* self);
extern __declspec(dllexport) void QIntValidator_BottomChanged(QIntValidator* self, int bottom);
void QIntValidator_connect_BottomChanged(QIntValidator* self, intptr_t slot);
extern __declspec(dllexport) void QIntValidator_TopChanged(QIntValidator* self, int top);
void QIntValidator_connect_TopChanged(QIntValidator* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QIntValidator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QIntValidator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QIntValidator_override_virtual_Validate(void* self, intptr_t slot);
int QIntValidator_virtualbase_Validate(const void* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) void QIntValidator_override_virtual_Fixup(void* self, intptr_t slot);
void QIntValidator_virtualbase_Fixup(const void* self, struct miqt_string input);
extern __declspec(dllexport) void QIntValidator_Delete(QIntValidator* self, bool isSubclass);

extern __declspec(dllexport) QDoubleValidator* QDoubleValidator_new();
extern __declspec(dllexport) QDoubleValidator* QDoubleValidator_new2(double bottom, double top, int decimals);
extern __declspec(dllexport) QDoubleValidator* QDoubleValidator_new3(QObject* parent);
extern __declspec(dllexport) QDoubleValidator* QDoubleValidator_new4(double bottom, double top, int decimals, QObject* parent);
extern __declspec(dllexport) void QDoubleValidator_virtbase(QDoubleValidator* src, QValidator** outptr_QValidator);
extern __declspec(dllexport) QMetaObject* QDoubleValidator_MetaObject(const QDoubleValidator* self);
extern __declspec(dllexport) void* QDoubleValidator_Metacast(QDoubleValidator* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QDoubleValidator_Tr(const char* s);
extern __declspec(dllexport) int QDoubleValidator_Validate(const QDoubleValidator* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) void QDoubleValidator_Fixup(const QDoubleValidator* self, struct miqt_string input);
extern __declspec(dllexport) void QDoubleValidator_SetRange(QDoubleValidator* self, double bottom, double top, int decimals);
extern __declspec(dllexport) void QDoubleValidator_SetRange2(QDoubleValidator* self, double bottom, double top);
extern __declspec(dllexport) void QDoubleValidator_SetBottom(QDoubleValidator* self, double bottom);
extern __declspec(dllexport) void QDoubleValidator_SetTop(QDoubleValidator* self, double top);
extern __declspec(dllexport) void QDoubleValidator_SetDecimals(QDoubleValidator* self, int decimals);
extern __declspec(dllexport) void QDoubleValidator_SetNotation(QDoubleValidator* self, Notation notation);
extern __declspec(dllexport) double QDoubleValidator_Bottom(const QDoubleValidator* self);
extern __declspec(dllexport) double QDoubleValidator_Top(const QDoubleValidator* self);
extern __declspec(dllexport) int QDoubleValidator_Decimals(const QDoubleValidator* self);
extern __declspec(dllexport) Notation QDoubleValidator_Notation(const QDoubleValidator* self);
extern __declspec(dllexport) void QDoubleValidator_BottomChanged(QDoubleValidator* self, double bottom);
void QDoubleValidator_connect_BottomChanged(QDoubleValidator* self, intptr_t slot);
extern __declspec(dllexport) void QDoubleValidator_TopChanged(QDoubleValidator* self, double top);
void QDoubleValidator_connect_TopChanged(QDoubleValidator* self, intptr_t slot);
extern __declspec(dllexport) void QDoubleValidator_DecimalsChanged(QDoubleValidator* self, int decimals);
void QDoubleValidator_connect_DecimalsChanged(QDoubleValidator* self, intptr_t slot);
extern __declspec(dllexport) void QDoubleValidator_NotationChanged(QDoubleValidator* self, int notation);
void QDoubleValidator_connect_NotationChanged(QDoubleValidator* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QDoubleValidator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QDoubleValidator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QDoubleValidator_override_virtual_Validate(void* self, intptr_t slot);
int QDoubleValidator_virtualbase_Validate(const void* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) void QDoubleValidator_override_virtual_Fixup(void* self, intptr_t slot);
void QDoubleValidator_virtualbase_Fixup(const void* self, struct miqt_string input);
extern __declspec(dllexport) void QDoubleValidator_Delete(QDoubleValidator* self, bool isSubclass);

extern __declspec(dllexport) QRegularExpressionValidator* QRegularExpressionValidator_new();
extern __declspec(dllexport) QRegularExpressionValidator* QRegularExpressionValidator_new2(QRegularExpression* re);
extern __declspec(dllexport) QRegularExpressionValidator* QRegularExpressionValidator_new3(QObject* parent);
extern __declspec(dllexport) QRegularExpressionValidator* QRegularExpressionValidator_new4(QRegularExpression* re, QObject* parent);
extern __declspec(dllexport) void QRegularExpressionValidator_virtbase(QRegularExpressionValidator* src, QValidator** outptr_QValidator);
extern __declspec(dllexport) QMetaObject* QRegularExpressionValidator_MetaObject(const QRegularExpressionValidator* self);
extern __declspec(dllexport) void* QRegularExpressionValidator_Metacast(QRegularExpressionValidator* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QRegularExpressionValidator_Tr(const char* s);
extern __declspec(dllexport) int QRegularExpressionValidator_Validate(const QRegularExpressionValidator* self, struct miqt_string input, int* pos);
extern __declspec(dllexport) QRegularExpression* QRegularExpressionValidator_RegularExpression(const QRegularExpressionValidator* self);
extern __declspec(dllexport) void QRegularExpressionValidator_SetRegularExpression(QRegularExpressionValidator* self, QRegularExpression* re);
extern __declspec(dllexport) void QRegularExpressionValidator_RegularExpressionChanged(QRegularExpressionValidator* self, QRegularExpression* re);
void QRegularExpressionValidator_connect_RegularExpressionChanged(QRegularExpressionValidator* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QRegularExpressionValidator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QRegularExpressionValidator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QRegularExpressionValidator_override_virtual_Validate(void* self, intptr_t slot);
int QRegularExpressionValidator_virtualbase_Validate(const void* self, struct miqt_string input, int* pos);
extern __declspec(dllexport) void QRegularExpressionValidator_override_virtual_Fixup(void* self, intptr_t slot);
void QRegularExpressionValidator_virtualbase_Fixup(const void* self, struct miqt_string param1);
extern __declspec(dllexport) void QRegularExpressionValidator_Delete(QRegularExpressionValidator* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
