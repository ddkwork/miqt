#pragma once
#ifndef MIQT_QT_GEN_QTREEWIDGETITEMITERATOR_H
#define MIQT_QT_GEN_QTREEWIDGETITEMITERATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QTreeWidget QTreeWidget;
typedef struct QTreeWidgetItem QTreeWidgetItem;
typedef struct QTreeWidgetItemIterator QTreeWidgetItemIterator;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_new(QTreeWidgetItemIterator* it);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_new2(QTreeWidget* widget);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_new3(QTreeWidgetItem* item);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_new4(QTreeWidget* widget, IteratorFlags flags);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_new5(QTreeWidgetItem* item, IteratorFlags flags);
extern __declspec(dllexport) void QTreeWidgetItemIterator_OperatorAssign(QTreeWidgetItemIterator* self, QTreeWidgetItemIterator* it);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_OperatorPlusPlus(QTreeWidgetItemIterator* self);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_OperatorPlusPlusWithInt(QTreeWidgetItemIterator* self, int param1);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_OperatorPlusAssign(QTreeWidgetItemIterator* self, int n);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_OperatorMinusMinus(QTreeWidgetItemIterator* self);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_OperatorMinusMinusWithInt(QTreeWidgetItemIterator* self, int param1);
extern __declspec(dllexport) QTreeWidgetItemIterator* QTreeWidgetItemIterator_OperatorMinusAssign(QTreeWidgetItemIterator* self, int n);
extern __declspec(dllexport) QTreeWidgetItem* QTreeWidgetItemIterator_OperatorMultiply(const QTreeWidgetItemIterator* self);
extern __declspec(dllexport) void QTreeWidgetItemIterator_Delete(QTreeWidgetItemIterator* self, bool isSubclass);

} 
