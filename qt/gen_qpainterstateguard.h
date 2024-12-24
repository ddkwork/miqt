#pragma once
#ifndef MIQT_QT_GEN_QPAINTERSTATEGUARD_H
#define MIQT_QT_GEN_QPAINTERSTATEGUARD_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QPainter QPainter;
typedef struct QPainterStateGuard QPainterStateGuard;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPainterStateGuard* QPainterStateGuard_new(QPainter* painter);
extern __declspec(dllexport) 
QPainterStateGuard* QPainterStateGuard_new2(QPainter* painter, InitialState state);
extern __declspec(dllexport) 
void QPainterStateGuard_Save(QPainterStateGuard* self);
extern __declspec(dllexport) 
void QPainterStateGuard_Restore(QPainterStateGuard* self);
extern __declspec(dllexport) 
void QPainterStateGuard_Delete(QPainterStateGuard* self, bool isSubclass);

}
