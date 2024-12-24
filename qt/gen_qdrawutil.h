#pragma once
#ifndef MIQT_QT_GEN_QDRAWUTIL_H
#define MIQT_QT_GEN_QDRAWUTIL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QTileRules QTileRules;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTileRules* QTileRules_new(int horizontalRule, int verticalRule);
extern __declspec(dllexport) 
QTileRules* QTileRules_new2();
extern __declspec(dllexport) 
QTileRules* QTileRules_new3(QTileRules* param1);
extern __declspec(dllexport) 
QTileRules* QTileRules_new4(int rule);
extern __declspec(dllexport) 
void QTileRules_Delete(QTileRules* self, bool isSubclass);

}
