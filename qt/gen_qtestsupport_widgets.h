#pragma once
#ifndef MIQT_QT_GEN_QTESTSUPPORT_WIDGETS_H
#define MIQT_QT_GEN_QTESTSUPPORT_WIDGETS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTest__QTouchEventWidgetSequence)
typedef QTest::QTouchEventWidgetSequence QTest__QTouchEventWidgetSequence;
typedef struct QPoint QPoint;
typedef struct QTest__QTouchEventWidgetSequence QTest__QTouchEventWidgetSequence;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_new(const QTouchEventWidgetSequence* param1);
extern __declspec(dllexport) QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Press(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt);
extern __declspec(dllexport) QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Move(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt);
extern __declspec(dllexport) QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Release(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt);
extern __declspec(dllexport) QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Stationary(QTest__QTouchEventWidgetSequence* self, int touchId);
extern __declspec(dllexport) bool QTest__QTouchEventWidgetSequence_Commit(QTest__QTouchEventWidgetSequence* self, bool processEvents);
extern __declspec(dllexport) QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Press3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget);
extern __declspec(dllexport) QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Move3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget);
extern __declspec(dllexport) QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Release3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget);
extern __declspec(dllexport) void QTest__QTouchEventWidgetSequence_override_virtual_Stationary(void* self, intptr_t slot);
QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_virtualbase_Stationary(void* self, int touchId);
extern __declspec(dllexport) void QTest__QTouchEventWidgetSequence_override_virtual_Commit(void* self, intptr_t slot);
bool QTest__QTouchEventWidgetSequence_virtualbase_Commit(void* self, bool processEvents);
extern __declspec(dllexport) void QTest__QTouchEventWidgetSequence_Delete(QTest__QTouchEventWidgetSequence* self, bool isSubclass);

} 
