#pragma once
#ifndef MIQT_QT_GEN_QCOMPARE_H
#define MIQT_QT_GEN_QCOMPARE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_partial_ordering)
typedef Qt::partial_ordering partial_ordering;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_strong_ordering)
typedef Qt::strong_ordering strong_ordering;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_weak_ordering)
typedef Qt::weak_ordering weak_ordering;
typedef struct QPartialOrdering QPartialOrdering;
typedef struct partial_ordering partial_ordering;
typedef struct strong_ordering strong_ordering;
typedef struct weak_ordering weak_ordering;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
partial_ordering* partial_ordering_new(const partial_ordering* param1);
extern __declspec(dllexport) 
void partial_ordering_Delete(partial_ordering* self, bool isSubclass);

extern __declspec(dllexport) 
weak_ordering* weak_ordering_new(const weak_ordering* param1);
extern __declspec(dllexport) 
void weak_ordering_Delete(weak_ordering* self, bool isSubclass);

extern __declspec(dllexport) 
strong_ordering* strong_ordering_new(const strong_ordering* param1);
extern __declspec(dllexport) 
void strong_ordering_Delete(strong_ordering* self, bool isSubclass);

extern __declspec(dllexport) 
QPartialOrdering* QPartialOrdering_new(partial_ordering* order);
extern __declspec(dllexport) 
QPartialOrdering* QPartialOrdering_new2(weak_ordering* stdorder);
extern __declspec(dllexport) 
QPartialOrdering* QPartialOrdering_new3(strong_ordering* stdorder);
extern __declspec(dllexport) 
QPartialOrdering* QPartialOrdering_new4(QPartialOrdering* param1);
extern __declspec(dllexport) 
void QPartialOrdering_Delete(QPartialOrdering* self, bool isSubclass);

}
