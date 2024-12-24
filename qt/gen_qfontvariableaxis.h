#pragma once
#ifndef MIQT_QT_GEN_QFONTVARIABLEAXIS_H
#define MIQT_QT_GEN_QFONTVARIABLEAXIS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QFont__Tag)
typedef QFont::Tag QFont__Tag;
typedef struct QFont__Tag QFont__Tag;
typedef struct QFontVariableAxis QFontVariableAxis;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFontVariableAxis* QFontVariableAxis_new();
extern __declspec(dllexport) 
QFontVariableAxis* QFontVariableAxis_new2(QFontVariableAxis* axis);
extern __declspec(dllexport) 
void QFontVariableAxis_Swap(QFontVariableAxis* self, QFontVariableAxis* other);
extern __declspec(dllexport) 
void QFontVariableAxis_OperatorAssign(QFontVariableAxis* self, QFontVariableAxis* axis);
extern __declspec(dllexport) 
QFont__Tag* QFontVariableAxis_Tag(const QFontVariableAxis* self);
extern __declspec(dllexport) 
void QFontVariableAxis_SetTag(QFontVariableAxis* self, QFont__Tag* tag);
extern __declspec(dllexport) 
struct miqt_string QFontVariableAxis_Name(const QFontVariableAxis* self);
extern __declspec(dllexport) 
void QFontVariableAxis_SetName(QFontVariableAxis* self, struct miqt_string name);
extern __declspec(dllexport) 
double QFontVariableAxis_MinimumValue(const QFontVariableAxis* self);
extern __declspec(dllexport) 
void QFontVariableAxis_SetMinimumValue(QFontVariableAxis* self, double minimumValue);
extern __declspec(dllexport) 
double QFontVariableAxis_MaximumValue(const QFontVariableAxis* self);
extern __declspec(dllexport) 
void QFontVariableAxis_SetMaximumValue(QFontVariableAxis* self, double maximumValue);
extern __declspec(dllexport) 
double QFontVariableAxis_DefaultValue(const QFontVariableAxis* self);
extern __declspec(dllexport) 
void QFontVariableAxis_SetDefaultValue(QFontVariableAxis* self, double defaultValue);
extern __declspec(dllexport) 
void QFontVariableAxis_Delete(QFontVariableAxis* self, bool isSubclass);

}
