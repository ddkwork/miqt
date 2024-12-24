#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTSLIDER_H
#define MIQT_QT_GEN_QABSTRACTSLIDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractSlider QAbstractSlider;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractSlider* QAbstractSlider_new(QWidget* parent);
extern __declspec(dllexport) 
QAbstractSlider* QAbstractSlider_new2();
extern __declspec(dllexport) 
void QAbstractSlider_virtbase(QAbstractSlider* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QAbstractSlider_MetaObject(const QAbstractSlider* self);
extern __declspec(dllexport) 
void* QAbstractSlider_Metacast(QAbstractSlider* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractSlider_Tr(const char* s);
extern __declspec(dllexport) 
int QAbstractSlider_Orientation(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetMinimum(QAbstractSlider* self, int minimum);
extern __declspec(dllexport) 
int QAbstractSlider_Minimum(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetMaximum(QAbstractSlider* self, int maximum);
extern __declspec(dllexport) 
int QAbstractSlider_Maximum(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetSingleStep(QAbstractSlider* self, int singleStep);
extern __declspec(dllexport) 
int QAbstractSlider_SingleStep(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetPageStep(QAbstractSlider* self, int pageStep);
extern __declspec(dllexport) 
int QAbstractSlider_PageStep(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetTracking(QAbstractSlider* self, bool enable);
extern __declspec(dllexport) 
bool QAbstractSlider_HasTracking(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetSliderDown(QAbstractSlider* self, bool sliderDown);
extern __declspec(dllexport) 
bool QAbstractSlider_IsSliderDown(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetSliderPosition(QAbstractSlider* self, int sliderPosition);
extern __declspec(dllexport) 
int QAbstractSlider_SliderPosition(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetInvertedAppearance(QAbstractSlider* self, bool invertedAppearance);
extern __declspec(dllexport) 
bool QAbstractSlider_InvertedAppearance(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_SetInvertedControls(QAbstractSlider* self, bool invertedControls);
extern __declspec(dllexport) 
bool QAbstractSlider_InvertedControls(const QAbstractSlider* self);
extern __declspec(dllexport) 
int QAbstractSlider_Value(const QAbstractSlider* self);
extern __declspec(dllexport) 
void QAbstractSlider_TriggerAction(QAbstractSlider* self, SliderAction action);
extern __declspec(dllexport) 
void QAbstractSlider_SetValue(QAbstractSlider* self, int value);
extern __declspec(dllexport) 
void QAbstractSlider_SetOrientation(QAbstractSlider* self, int orientation);
extern __declspec(dllexport) 
void QAbstractSlider_SetRange(QAbstractSlider* self, int min, int max);
extern __declspec(dllexport) 
void QAbstractSlider_ValueChanged(QAbstractSlider* self, int value);
void QAbstractSlider_connect_ValueChanged(QAbstractSlider* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSlider_SliderPressed(QAbstractSlider* self);
void QAbstractSlider_connect_SliderPressed(QAbstractSlider* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSlider_SliderMoved(QAbstractSlider* self, int position);
void QAbstractSlider_connect_SliderMoved(QAbstractSlider* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSlider_SliderReleased(QAbstractSlider* self);
void QAbstractSlider_connect_SliderReleased(QAbstractSlider* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSlider_RangeChanged(QAbstractSlider* self, int min, int max);
void QAbstractSlider_connect_RangeChanged(QAbstractSlider* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSlider_ActionTriggered(QAbstractSlider* self, int action);
void QAbstractSlider_connect_ActionTriggered(QAbstractSlider* self, intptr_t slot);
extern __declspec(dllexport) 
bool QAbstractSlider_Event(QAbstractSlider* self, QEvent* e);
extern __declspec(dllexport) 
void QAbstractSlider_SliderChange(QAbstractSlider* self, SliderChange change);
extern __declspec(dllexport) 
void QAbstractSlider_KeyPressEvent(QAbstractSlider* self, QKeyEvent* ev);
extern __declspec(dllexport) 
void QAbstractSlider_TimerEvent(QAbstractSlider* self, QTimerEvent* param1);
extern __declspec(dllexport) 
void QAbstractSlider_WheelEvent(QAbstractSlider* self, QWheelEvent* e);
extern __declspec(dllexport) 
void QAbstractSlider_ChangeEvent(QAbstractSlider* self, QEvent* e);
extern __declspec(dllexport) 
struct miqt_string QAbstractSlider_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractSlider_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractSlider_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractSlider_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractSlider_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractSlider_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractSlider_Delete(QAbstractSlider* self, bool isSubclass);

}
