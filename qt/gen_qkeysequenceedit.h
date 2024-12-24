#pragma once
#ifndef MIQT_QT_GEN_QKEYSEQUENCEEDIT_H
#define MIQT_QT_GEN_QKEYSEQUENCEEDIT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QKeyCombination QKeyCombination;
typedef struct QKeyEvent QKeyEvent;
typedef struct QKeySequence QKeySequence;
typedef struct QKeySequenceEdit QKeySequenceEdit;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QKeySequenceEdit* QKeySequenceEdit_new(QWidget* parent);
extern __declspec(dllexport) 
QKeySequenceEdit* QKeySequenceEdit_new2();
extern __declspec(dllexport) 
QKeySequenceEdit* QKeySequenceEdit_new3(QKeySequence* keySequence);
extern __declspec(dllexport) 
QKeySequenceEdit* QKeySequenceEdit_new4(QKeySequence* keySequence, QWidget* parent);
extern __declspec(dllexport) 
void QKeySequenceEdit_virtbase(QKeySequenceEdit* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QKeySequenceEdit_MetaObject(const QKeySequenceEdit* self);
extern __declspec(dllexport) 
void* QKeySequenceEdit_Metacast(QKeySequenceEdit* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QKeySequenceEdit_Tr(const char* s);
extern __declspec(dllexport) 
QKeySequence* QKeySequenceEdit_KeySequence(const QKeySequenceEdit* self);
extern __declspec(dllexport) 
ptrdiff_t QKeySequenceEdit_MaximumSequenceLength(const QKeySequenceEdit* self);
extern __declspec(dllexport) 
void QKeySequenceEdit_SetClearButtonEnabled(QKeySequenceEdit* self, bool enable);
extern __declspec(dllexport) 
bool QKeySequenceEdit_IsClearButtonEnabled(const QKeySequenceEdit* self);
extern __declspec(dllexport) 
void QKeySequenceEdit_SetFinishingKeyCombinations(QKeySequenceEdit* self, struct miqt_array /* of QKeyCombination* */  finishingKeyCombinations);
extern __declspec(dllexport) 
struct miqt_array /* of QKeyCombination* */  QKeySequenceEdit_FinishingKeyCombinations(const QKeySequenceEdit* self);
extern __declspec(dllexport) 
void QKeySequenceEdit_SetKeySequence(QKeySequenceEdit* self, QKeySequence* keySequence);
extern __declspec(dllexport) 
void QKeySequenceEdit_Clear(QKeySequenceEdit* self);
extern __declspec(dllexport) 
void QKeySequenceEdit_SetMaximumSequenceLength(QKeySequenceEdit* self, ptrdiff_t count);
extern __declspec(dllexport) 
void QKeySequenceEdit_EditingFinished(QKeySequenceEdit* self);
void QKeySequenceEdit_connect_EditingFinished(QKeySequenceEdit* self, intptr_t slot);
extern __declspec(dllexport) 
void QKeySequenceEdit_KeySequenceChanged(QKeySequenceEdit* self, QKeySequence* keySequence);
void QKeySequenceEdit_connect_KeySequenceChanged(QKeySequenceEdit* self, intptr_t slot);
extern __declspec(dllexport) 
bool QKeySequenceEdit_Event(QKeySequenceEdit* self, QEvent* param1);
extern __declspec(dllexport) 
void QKeySequenceEdit_KeyPressEvent(QKeySequenceEdit* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QKeySequenceEdit_KeyReleaseEvent(QKeySequenceEdit* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QKeySequenceEdit_TimerEvent(QKeySequenceEdit* self, QTimerEvent* param1);
extern __declspec(dllexport) 
void QKeySequenceEdit_FocusOutEvent(QKeySequenceEdit* self, QFocusEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QKeySequenceEdit_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QKeySequenceEdit_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QKeySequenceEdit_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QKeySequenceEdit_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QKeySequenceEdit_override_virtual_Metacast(void* self, intptr_t slot);
void* QKeySequenceEdit_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QKeySequenceEdit_Delete(QKeySequenceEdit* self, bool isSubclass);

}
