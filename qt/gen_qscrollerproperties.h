#pragma once
#ifndef MIQT_QT_GEN_QSCROLLERPROPERTIES_H
#define MIQT_QT_GEN_QSCROLLERPROPERTIES_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QScrollerProperties QScrollerProperties;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QScrollerProperties* QScrollerProperties_new();
extern __declspec(dllexport) 
QScrollerProperties* QScrollerProperties_new2(QScrollerProperties* sp);
extern __declspec(dllexport) 
void QScrollerProperties_OperatorAssign(QScrollerProperties* self, QScrollerProperties* sp);
extern __declspec(dllexport) 
bool QScrollerProperties_OperatorEqual(const QScrollerProperties* self, QScrollerProperties* sp);
extern __declspec(dllexport) 
bool QScrollerProperties_OperatorNotEqual(const QScrollerProperties* self, QScrollerProperties* sp);
extern __declspec(dllexport) 
void QScrollerProperties_SetDefaultScrollerProperties(QScrollerProperties* sp);
extern __declspec(dllexport) 
void QScrollerProperties_UnsetDefaultScrollerProperties();
extern __declspec(dllexport) 
QVariant* QScrollerProperties_ScrollMetric(const QScrollerProperties* self, ScrollMetric metric);
extern __declspec(dllexport) 
void QScrollerProperties_SetScrollMetric(QScrollerProperties* self, ScrollMetric metric, QVariant* value);
extern __declspec(dllexport) 
void QScrollerProperties_Delete(QScrollerProperties* self, bool isSubclass);

}
