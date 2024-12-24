#pragma once
#ifndef MIQT_QT_GEN_QLCDNUMBER_H
#define MIQT_QT_GEN_QLCDNUMBER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QLCDNumber QLCDNumber;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QLCDNumber* QLCDNumber_new(QWidget* parent);
extern __declspec(dllexport) 
QLCDNumber* QLCDNumber_new2();
extern __declspec(dllexport) 
QLCDNumber* QLCDNumber_new3(unsigned int numDigits);
extern __declspec(dllexport) 
QLCDNumber* QLCDNumber_new4(unsigned int numDigits, QWidget* parent);
extern __declspec(dllexport) 
void QLCDNumber_virtbase(QLCDNumber* src
, QFrame** outptr_QFrame
);
extern __declspec(dllexport) 
QMetaObject* QLCDNumber_MetaObject(const QLCDNumber* self);
extern __declspec(dllexport) 
void* QLCDNumber_Metacast(QLCDNumber* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QLCDNumber_Tr(const char* s);
extern __declspec(dllexport) 
bool QLCDNumber_SmallDecimalPoint(const QLCDNumber* self);
extern __declspec(dllexport) 
int QLCDNumber_DigitCount(const QLCDNumber* self);
extern __declspec(dllexport) 
void QLCDNumber_SetDigitCount(QLCDNumber* self, int nDigits);
extern __declspec(dllexport) 
bool QLCDNumber_CheckOverflow(const QLCDNumber* self, double num);
extern __declspec(dllexport) 
bool QLCDNumber_CheckOverflowWithNum(const QLCDNumber* self, int num);
extern __declspec(dllexport) 
Mode QLCDNumber_Mode(const QLCDNumber* self);
extern __declspec(dllexport) 
void QLCDNumber_SetMode(QLCDNumber* self, Mode mode);
extern __declspec(dllexport) 
SegmentStyle QLCDNumber_SegmentStyle(const QLCDNumber* self);
extern __declspec(dllexport) 
void QLCDNumber_SetSegmentStyle(QLCDNumber* self, SegmentStyle segmentStyle);
extern __declspec(dllexport) 
double QLCDNumber_Value(const QLCDNumber* self);
extern __declspec(dllexport) 
int QLCDNumber_IntValue(const QLCDNumber* self);
extern __declspec(dllexport) 
QSize* QLCDNumber_SizeHint(const QLCDNumber* self);
extern __declspec(dllexport) 
void QLCDNumber_Display(QLCDNumber* self, struct miqt_string str);
extern __declspec(dllexport) 
void QLCDNumber_DisplayWithNum(QLCDNumber* self, int num);
extern __declspec(dllexport) 
void QLCDNumber_Display2(QLCDNumber* self, double num);
extern __declspec(dllexport) 
void QLCDNumber_SetHexMode(QLCDNumber* self);
extern __declspec(dllexport) 
void QLCDNumber_SetDecMode(QLCDNumber* self);
extern __declspec(dllexport) 
void QLCDNumber_SetOctMode(QLCDNumber* self);
extern __declspec(dllexport) 
void QLCDNumber_SetBinMode(QLCDNumber* self);
extern __declspec(dllexport) 
void QLCDNumber_SetSmallDecimalPoint(QLCDNumber* self, bool smallDecimalPoint);
extern __declspec(dllexport) 
void QLCDNumber_Overflow(QLCDNumber* self);
void QLCDNumber_connect_Overflow(QLCDNumber* self, intptr_t slot);
extern __declspec(dllexport) 
bool QLCDNumber_Event(QLCDNumber* self, QEvent* e);
extern __declspec(dllexport) 
void QLCDNumber_PaintEvent(QLCDNumber* self, QPaintEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QLCDNumber_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QLCDNumber_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QLCDNumber_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QLCDNumber_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QLCDNumber_override_virtual_Metacast(void* self, intptr_t slot);
void* QLCDNumber_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QLCDNumber_Delete(QLCDNumber* self, bool isSubclass);

}
