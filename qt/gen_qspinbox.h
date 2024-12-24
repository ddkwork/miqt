#pragma once
#ifndef MIQT_QT_GEN_QSPINBOX_H
#define MIQT_QT_GEN_QSPINBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractSpinBox QAbstractSpinBox;
typedef struct QDoubleSpinBox QDoubleSpinBox;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QSpinBox QSpinBox;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSpinBox* QSpinBox_new(QWidget* parent);
extern __declspec(dllexport) 
QSpinBox* QSpinBox_new2();
extern __declspec(dllexport) 
void QSpinBox_virtbase(QSpinBox* src
, QAbstractSpinBox** outptr_QAbstractSpinBox
);
extern __declspec(dllexport) 
QMetaObject* QSpinBox_MetaObject(const QSpinBox* self);
extern __declspec(dllexport) 
void* QSpinBox_Metacast(QSpinBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSpinBox_Tr(const char* s);
extern __declspec(dllexport) 
int QSpinBox_Value(const QSpinBox* self);
extern __declspec(dllexport) 
struct miqt_string QSpinBox_Prefix(const QSpinBox* self);
extern __declspec(dllexport) 
void QSpinBox_SetPrefix(QSpinBox* self, struct miqt_string prefix);
extern __declspec(dllexport) 
struct miqt_string QSpinBox_Suffix(const QSpinBox* self);
extern __declspec(dllexport) 
void QSpinBox_SetSuffix(QSpinBox* self, struct miqt_string suffix);
extern __declspec(dllexport) 
struct miqt_string QSpinBox_CleanText(const QSpinBox* self);
extern __declspec(dllexport) 
int QSpinBox_SingleStep(const QSpinBox* self);
extern __declspec(dllexport) 
void QSpinBox_SetSingleStep(QSpinBox* self, int val);
extern __declspec(dllexport) 
int QSpinBox_Minimum(const QSpinBox* self);
extern __declspec(dllexport) 
void QSpinBox_SetMinimum(QSpinBox* self, int min);
extern __declspec(dllexport) 
int QSpinBox_Maximum(const QSpinBox* self);
extern __declspec(dllexport) 
void QSpinBox_SetMaximum(QSpinBox* self, int max);
extern __declspec(dllexport) 
void QSpinBox_SetRange(QSpinBox* self, int min, int max);
extern __declspec(dllexport) 
StepType QSpinBox_StepType(const QSpinBox* self);
extern __declspec(dllexport) 
void QSpinBox_SetStepType(QSpinBox* self, StepType stepType);
extern __declspec(dllexport) 
int QSpinBox_DisplayIntegerBase(const QSpinBox* self);
extern __declspec(dllexport) 
void QSpinBox_SetDisplayIntegerBase(QSpinBox* self, int base);
extern __declspec(dllexport) 
bool QSpinBox_Event(QSpinBox* self, QEvent* event);
extern __declspec(dllexport) 
int QSpinBox_Validate(const QSpinBox* self, struct miqt_string input, int* pos);
extern __declspec(dllexport) 
int QSpinBox_ValueFromText(const QSpinBox* self, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QSpinBox_TextFromValue(const QSpinBox* self, int val);
extern __declspec(dllexport) 
void QSpinBox_Fixup(const QSpinBox* self, struct miqt_string str);
extern __declspec(dllexport) 
void QSpinBox_SetValue(QSpinBox* self, int val);
extern __declspec(dllexport) 
void QSpinBox_ValueChanged(QSpinBox* self, int param1);
void QSpinBox_connect_ValueChanged(QSpinBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QSpinBox_TextChanged(QSpinBox* self, struct miqt_string param1);
void QSpinBox_connect_TextChanged(QSpinBox* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QSpinBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSpinBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSpinBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSpinBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSpinBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QSpinBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSpinBox_Delete(QSpinBox* self, bool isSubclass);

extern __declspec(dllexport) 
QDoubleSpinBox* QDoubleSpinBox_new(QWidget* parent);
extern __declspec(dllexport) 
QDoubleSpinBox* QDoubleSpinBox_new2();
extern __declspec(dllexport) 
void QDoubleSpinBox_virtbase(QDoubleSpinBox* src
, QAbstractSpinBox** outptr_QAbstractSpinBox
);
extern __declspec(dllexport) 
QMetaObject* QDoubleSpinBox_MetaObject(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
void* QDoubleSpinBox_Metacast(QDoubleSpinBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QDoubleSpinBox_Tr(const char* s);
extern __declspec(dllexport) 
double QDoubleSpinBox_Value(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
struct miqt_string QDoubleSpinBox_Prefix(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetPrefix(QDoubleSpinBox* self, struct miqt_string prefix);
extern __declspec(dllexport) 
struct miqt_string QDoubleSpinBox_Suffix(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetSuffix(QDoubleSpinBox* self, struct miqt_string suffix);
extern __declspec(dllexport) 
struct miqt_string QDoubleSpinBox_CleanText(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
double QDoubleSpinBox_SingleStep(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetSingleStep(QDoubleSpinBox* self, double val);
extern __declspec(dllexport) 
double QDoubleSpinBox_Minimum(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetMinimum(QDoubleSpinBox* self, double min);
extern __declspec(dllexport) 
double QDoubleSpinBox_Maximum(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetMaximum(QDoubleSpinBox* self, double max);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetRange(QDoubleSpinBox* self, double min, double max);
extern __declspec(dllexport) 
StepType QDoubleSpinBox_StepType(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetStepType(QDoubleSpinBox* self, StepType stepType);
extern __declspec(dllexport) 
int QDoubleSpinBox_Decimals(const QDoubleSpinBox* self);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetDecimals(QDoubleSpinBox* self, int prec);
extern __declspec(dllexport) 
int QDoubleSpinBox_Validate(const QDoubleSpinBox* self, struct miqt_string input, int* pos);
extern __declspec(dllexport) 
double QDoubleSpinBox_ValueFromText(const QDoubleSpinBox* self, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QDoubleSpinBox_TextFromValue(const QDoubleSpinBox* self, double val);
extern __declspec(dllexport) 
void QDoubleSpinBox_Fixup(const QDoubleSpinBox* self, struct miqt_string str);
extern __declspec(dllexport) 
void QDoubleSpinBox_SetValue(QDoubleSpinBox* self, double val);
extern __declspec(dllexport) 
void QDoubleSpinBox_ValueChanged(QDoubleSpinBox* self, double param1);
void QDoubleSpinBox_connect_ValueChanged(QDoubleSpinBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QDoubleSpinBox_TextChanged(QDoubleSpinBox* self, struct miqt_string param1);
void QDoubleSpinBox_connect_TextChanged(QDoubleSpinBox* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QDoubleSpinBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QDoubleSpinBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QDoubleSpinBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QDoubleSpinBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QDoubleSpinBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QDoubleSpinBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QDoubleSpinBox_Delete(QDoubleSpinBox* self, bool isSubclass);

}
