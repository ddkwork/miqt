#pragma once
#ifndef MIQT_QT_GEN_QPDFOUTPUTINTENT_H
#define MIQT_QT_GEN_QPDFOUTPUTINTENT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QColorSpace QColorSpace;
typedef struct QPdfOutputIntent QPdfOutputIntent;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPdfOutputIntent* QPdfOutputIntent_new();
extern __declspec(dllexport) 
QPdfOutputIntent* QPdfOutputIntent_new2(QPdfOutputIntent* other);
extern __declspec(dllexport) 
void QPdfOutputIntent_OperatorAssign(QPdfOutputIntent* self, QPdfOutputIntent* other);
extern __declspec(dllexport) 
void QPdfOutputIntent_Swap(QPdfOutputIntent* self, QPdfOutputIntent* other);
extern __declspec(dllexport) 
struct miqt_string QPdfOutputIntent_OutputConditionIdentifier(const QPdfOutputIntent* self);
extern __declspec(dllexport) 
void QPdfOutputIntent_SetOutputConditionIdentifier(QPdfOutputIntent* self, struct miqt_string identifier);
extern __declspec(dllexport) 
struct miqt_string QPdfOutputIntent_OutputCondition(const QPdfOutputIntent* self);
extern __declspec(dllexport) 
void QPdfOutputIntent_SetOutputCondition(QPdfOutputIntent* self, struct miqt_string condition);
extern __declspec(dllexport) 
QUrl* QPdfOutputIntent_RegistryName(const QPdfOutputIntent* self);
extern __declspec(dllexport) 
void QPdfOutputIntent_SetRegistryName(QPdfOutputIntent* self, QUrl* name);
extern __declspec(dllexport) 
QColorSpace* QPdfOutputIntent_OutputProfile(const QPdfOutputIntent* self);
extern __declspec(dllexport) 
void QPdfOutputIntent_SetOutputProfile(QPdfOutputIntent* self, QColorSpace* profile);
extern __declspec(dllexport) 
void QPdfOutputIntent_Delete(QPdfOutputIntent* self, bool isSubclass);

}
