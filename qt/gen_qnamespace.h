#pragma once
#ifndef MIQT_QT_GEN_QNAMESPACE_H
#define MIQT_QT_GEN_QNAMESPACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_Disambiguated_t)
typedef Qt::Disambiguated_t Disambiguated_t;
typedef struct QInternal QInternal;
typedef struct QKeyCombination QKeyCombination;
typedef struct Disambiguated_t Disambiguated_t;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
Disambiguated_t* Disambiguated_t_new();
extern __declspec(dllexport) 
Disambiguated_t* Disambiguated_t_new2(const Disambiguated_t* param1);
extern __declspec(dllexport) 
void Disambiguated_t_Delete(Disambiguated_t* self, bool isSubclass);

extern __declspec(dllexport) 
void QInternal_Delete(QInternal* self, bool isSubclass);

extern __declspec(dllexport) 
QKeyCombination* QKeyCombination_new();
extern __declspec(dllexport) 
QKeyCombination* QKeyCombination_new2(int modifiers);
extern __declspec(dllexport) 
QKeyCombination* QKeyCombination_new3(int modifiers);
extern __declspec(dllexport) 
QKeyCombination* QKeyCombination_new4(QKeyCombination* param1);
extern __declspec(dllexport) 
QKeyCombination* QKeyCombination_new5(int key);
extern __declspec(dllexport) 
QKeyCombination* QKeyCombination_new6(int modifiers, int key);
extern __declspec(dllexport) 
QKeyCombination* QKeyCombination_new7(int modifiers, int key);
extern __declspec(dllexport) 
int QKeyCombination_KeyboardModifiers(const QKeyCombination* self);
extern __declspec(dllexport) 
int QKeyCombination_Key(const QKeyCombination* self);
extern __declspec(dllexport) 
QKeyCombination* QKeyCombination_FromCombined(int combined);
extern __declspec(dllexport) 
int QKeyCombination_ToCombined(const QKeyCombination* self);
extern __declspec(dllexport) 
void QKeyCombination_Delete(QKeyCombination* self, bool isSubclass);

}
