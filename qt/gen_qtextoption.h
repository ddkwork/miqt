#pragma once
#ifndef MIQT_QT_GEN_QTEXTOPTION_H
#define MIQT_QT_GEN_QTEXTOPTION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextOption__Tab)
typedef QTextOption::Tab QTextOption__Tab;
typedef struct QChar QChar;
typedef struct QTextOption QTextOption;
typedef struct QTextOption__Tab QTextOption__Tab;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTextOption* QTextOption_new();
extern __declspec(dllexport) 
QTextOption* QTextOption_new2(int alignment);
extern __declspec(dllexport) 
QTextOption* QTextOption_new3(QTextOption* o);
extern __declspec(dllexport) 
void QTextOption_OperatorAssign(QTextOption* self, QTextOption* o);
extern __declspec(dllexport) 
void QTextOption_SetAlignment(QTextOption* self, int alignment);
extern __declspec(dllexport) 
int QTextOption_Alignment(const QTextOption* self);
extern __declspec(dllexport) 
void QTextOption_SetTextDirection(QTextOption* self, int aDirection);
extern __declspec(dllexport) 
int QTextOption_TextDirection(const QTextOption* self);
extern __declspec(dllexport) 
void QTextOption_SetWrapMode(QTextOption* self, WrapMode wrap);
extern __declspec(dllexport) 
WrapMode QTextOption_WrapMode(const QTextOption* self);
extern __declspec(dllexport) 
void QTextOption_SetFlags(QTextOption* self, Flags flags);
extern __declspec(dllexport) 
Flags QTextOption_Flags(const QTextOption* self);
extern __declspec(dllexport) 
void QTextOption_SetTabStopDistance(QTextOption* self, double tabStopDistance);
extern __declspec(dllexport) 
double QTextOption_TabStopDistance(const QTextOption* self);
extern __declspec(dllexport) 
void QTextOption_SetTabArray(QTextOption* self, struct miqt_array /* of double */  tabStops);
extern __declspec(dllexport) 
struct miqt_array /* of double */  QTextOption_TabArray(const QTextOption* self);
extern __declspec(dllexport) 
void QTextOption_SetTabs(QTextOption* self, struct miqt_array /* of Tab */  tabStops);
extern __declspec(dllexport) 
struct miqt_array /* of Tab */  QTextOption_Tabs(const QTextOption* self);
extern __declspec(dllexport) 
void QTextOption_SetUseDesignMetrics(QTextOption* self, bool b);
extern __declspec(dllexport) 
bool QTextOption_UseDesignMetrics(const QTextOption* self);
extern __declspec(dllexport) 
void QTextOption_Delete(QTextOption* self, bool isSubclass);

extern __declspec(dllexport) 
QTextOption__Tab* QTextOption__Tab_new();
extern __declspec(dllexport) 
QTextOption__Tab* QTextOption__Tab_new2(double pos, TabType tabType);
extern __declspec(dllexport) 
QTextOption__Tab* QTextOption__Tab_new3(double pos, TabType tabType, QChar* delim);
extern __declspec(dllexport) 
bool QTextOption__Tab_OperatorEqual(const QTextOption__Tab* self, const Tab* other);
extern __declspec(dllexport) 
bool QTextOption__Tab_OperatorNotEqual(const QTextOption__Tab* self, const Tab* other);
extern __declspec(dllexport) 
void QTextOption__Tab_Delete(QTextOption__Tab* self, bool isSubclass);

}
