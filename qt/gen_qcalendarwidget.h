#pragma once
#ifndef MIQT_QT_GEN_QCALENDARWIDGET_H
#define MIQT_QT_GEN_QCALENDARWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCalendar QCalendar;
typedef struct QCalendarWidget QCalendarWidget;
typedef struct QDate QDate;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPainter QPainter;
typedef struct QRect QRect;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QTextCharFormat QTextCharFormat;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QCalendarWidget* QCalendarWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QCalendarWidget* QCalendarWidget_new2();
extern __declspec(dllexport) 
void QCalendarWidget_virtbase(QCalendarWidget* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QCalendarWidget_MetaObject(const QCalendarWidget* self);
extern __declspec(dllexport) 
void* QCalendarWidget_Metacast(QCalendarWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QCalendarWidget_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QCalendarWidget_SizeHint(const QCalendarWidget* self);
extern __declspec(dllexport) 
QSize* QCalendarWidget_MinimumSizeHint(const QCalendarWidget* self);
extern __declspec(dllexport) 
QDate* QCalendarWidget_SelectedDate(const QCalendarWidget* self);
extern __declspec(dllexport) 
int QCalendarWidget_YearShown(const QCalendarWidget* self);
extern __declspec(dllexport) 
int QCalendarWidget_MonthShown(const QCalendarWidget* self);
extern __declspec(dllexport) 
QDate* QCalendarWidget_MinimumDate(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetMinimumDate(QCalendarWidget* self, QDate* date);
extern __declspec(dllexport) 
void QCalendarWidget_ClearMinimumDate(QCalendarWidget* self);
extern __declspec(dllexport) 
QDate* QCalendarWidget_MaximumDate(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetMaximumDate(QCalendarWidget* self, QDate* date);
extern __declspec(dllexport) 
void QCalendarWidget_ClearMaximumDate(QCalendarWidget* self);
extern __declspec(dllexport) 
int QCalendarWidget_FirstDayOfWeek(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetFirstDayOfWeek(QCalendarWidget* self, int dayOfWeek);
extern __declspec(dllexport) 
bool QCalendarWidget_IsNavigationBarVisible(const QCalendarWidget* self);
extern __declspec(dllexport) 
bool QCalendarWidget_IsGridVisible(const QCalendarWidget* self);
extern __declspec(dllexport) 
QCalendar* QCalendarWidget_Calendar(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetCalendar(QCalendarWidget* self, QCalendar* calendar);
extern __declspec(dllexport) 
SelectionMode QCalendarWidget_SelectionMode(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetSelectionMode(QCalendarWidget* self, SelectionMode mode);
extern __declspec(dllexport) 
HorizontalHeaderFormat QCalendarWidget_HorizontalHeaderFormat(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetHorizontalHeaderFormat(QCalendarWidget* self, HorizontalHeaderFormat format);
extern __declspec(dllexport) 
VerticalHeaderFormat QCalendarWidget_VerticalHeaderFormat(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetVerticalHeaderFormat(QCalendarWidget* self, VerticalHeaderFormat format);
extern __declspec(dllexport) 
QTextCharFormat* QCalendarWidget_HeaderTextFormat(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetHeaderTextFormat(QCalendarWidget* self, QTextCharFormat* format);
extern __declspec(dllexport) 
QTextCharFormat* QCalendarWidget_WeekdayTextFormat(const QCalendarWidget* self, int dayOfWeek);
extern __declspec(dllexport) 
void QCalendarWidget_SetWeekdayTextFormat(QCalendarWidget* self, int dayOfWeek, QTextCharFormat* format);
extern __declspec(dllexport) 
struct miqt_map /* of QDate* to QTextCharFormat* */  QCalendarWidget_DateTextFormat(const QCalendarWidget* self);
extern __declspec(dllexport) 
QTextCharFormat* QCalendarWidget_DateTextFormatWithDate(const QCalendarWidget* self, QDate* date);
extern __declspec(dllexport) 
void QCalendarWidget_SetDateTextFormat(QCalendarWidget* self, QDate* date, QTextCharFormat* format);
extern __declspec(dllexport) 
bool QCalendarWidget_IsDateEditEnabled(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetDateEditEnabled(QCalendarWidget* self, bool enable);
extern __declspec(dllexport) 
int QCalendarWidget_DateEditAcceptDelay(const QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SetDateEditAcceptDelay(QCalendarWidget* self, int delay);
extern __declspec(dllexport) 
bool QCalendarWidget_Event(QCalendarWidget* self, QEvent* event);
extern __declspec(dllexport) 
bool QCalendarWidget_EventFilter(QCalendarWidget* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) 
void QCalendarWidget_MousePressEvent(QCalendarWidget* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QCalendarWidget_ResizeEvent(QCalendarWidget* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QCalendarWidget_KeyPressEvent(QCalendarWidget* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QCalendarWidget_PaintCell(const QCalendarWidget* self, QPainter* painter, QRect* rect, QDate* date);
extern __declspec(dllexport) 
void QCalendarWidget_SetSelectedDate(QCalendarWidget* self, QDate* date);
extern __declspec(dllexport) 
void QCalendarWidget_SetDateRange(QCalendarWidget* self, QDate* min, QDate* max);
extern __declspec(dllexport) 
void QCalendarWidget_SetCurrentPage(QCalendarWidget* self, int year, int month);
extern __declspec(dllexport) 
void QCalendarWidget_SetGridVisible(QCalendarWidget* self, bool show);
extern __declspec(dllexport) 
void QCalendarWidget_SetNavigationBarVisible(QCalendarWidget* self, bool visible);
extern __declspec(dllexport) 
void QCalendarWidget_ShowNextMonth(QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_ShowPreviousMonth(QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_ShowNextYear(QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_ShowPreviousYear(QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_ShowSelectedDate(QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_ShowToday(QCalendarWidget* self);
extern __declspec(dllexport) 
void QCalendarWidget_SelectionChanged(QCalendarWidget* self);
void QCalendarWidget_connect_SelectionChanged(QCalendarWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QCalendarWidget_Clicked(QCalendarWidget* self, QDate* date);
void QCalendarWidget_connect_Clicked(QCalendarWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QCalendarWidget_Activated(QCalendarWidget* self, QDate* date);
void QCalendarWidget_connect_Activated(QCalendarWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QCalendarWidget_CurrentPageChanged(QCalendarWidget* self, int year, int month);
void QCalendarWidget_connect_CurrentPageChanged(QCalendarWidget* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QCalendarWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QCalendarWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QCalendarWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QCalendarWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QCalendarWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QCalendarWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QCalendarWidget_Delete(QCalendarWidget* self, bool isSubclass);

}
