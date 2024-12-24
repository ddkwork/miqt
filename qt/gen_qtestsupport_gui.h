#pragma once
#ifndef MIQT_QT_GEN_QTESTSUPPORT_GUI_H
#define MIQT_QT_GEN_QTESTSUPPORT_GUI_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPoint;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTest__QTouchEventSequence)
typedef QTest::QTouchEventSequence QTest__QTouchEventSequence;
#else
class QTest__QTouchEventSequence;
#endif
class QWindow;
class _GUID;
class type_info;
#else
typedef struct QPoint QPoint;
typedef struct QTest__QTouchEventSequence QTest__QTouchEventSequence;
typedef struct QWindow QWindow;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTouchEventSequence* QTest__QTouchEventSequence_Press(QTest__QTouchEventSequence* self, int touchId, QPoint* pt);
extern __declspec(dllexport) QTouchEventSequence* QTest__QTouchEventSequence_Move(QTest__QTouchEventSequence* self, int touchId, QPoint* pt);
extern __declspec(dllexport) QTouchEventSequence* QTest__QTouchEventSequence_Release(QTest__QTouchEventSequence* self, int touchId, QPoint* pt);
extern __declspec(dllexport) QTouchEventSequence* QTest__QTouchEventSequence_Stationary(QTest__QTouchEventSequence* self, int touchId);
extern __declspec(dllexport) bool QTest__QTouchEventSequence_Commit(QTest__QTouchEventSequence* self, bool processEvents);
extern __declspec(dllexport) QTouchEventSequence* QTest__QTouchEventSequence_Press3(QTest__QTouchEventSequence* self, int touchId, QPoint* pt, QWindow* window);
extern __declspec(dllexport) QTouchEventSequence* QTest__QTouchEventSequence_Move3(QTest__QTouchEventSequence* self, int touchId, QPoint* pt, QWindow* window);
extern __declspec(dllexport) QTouchEventSequence* QTest__QTouchEventSequence_Release3(QTest__QTouchEventSequence* self, int touchId, QPoint* pt, QWindow* window);
extern __declspec(dllexport) void QTest__QTouchEventSequence_Delete(QTest__QTouchEventSequence* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
