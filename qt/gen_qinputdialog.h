#pragma once
#ifndef MIQT_QT_GEN_QINPUTDIALOG_H
#define MIQT_QT_GEN_QINPUTDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCloseEvent;
class QContextMenuEvent;
class QDialog;
class QEvent;
class QInputDialog;
class QKeyEvent;
class QMetaObject;
class QObject;
class QPaintDevice;
class QResizeEvent;
class QShowEvent;
class QSize;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QInputDialog QInputDialog;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QInputDialog* QInputDialog_new(QWidget* parent);
extern __declspec(dllexport) QInputDialog* QInputDialog_new2();
extern __declspec(dllexport) QInputDialog* QInputDialog_new3(QWidget* parent, int flags);
extern __declspec(dllexport) void QInputDialog_virtbase(QInputDialog* src, QDialog** outptr_QDialog);
extern __declspec(dllexport) QMetaObject* QInputDialog_MetaObject(const QInputDialog* self);
extern __declspec(dllexport) void* QInputDialog_Metacast(QInputDialog* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QInputDialog_Tr(const char* s);
extern __declspec(dllexport) void QInputDialog_SetInputMode(QInputDialog* self, InputMode mode);
extern __declspec(dllexport) InputMode QInputDialog_InputMode(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetLabelText(QInputDialog* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QInputDialog_LabelText(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetOption(QInputDialog* self, InputDialogOption option);
extern __declspec(dllexport) bool QInputDialog_TestOption(const QInputDialog* self, InputDialogOption option);
extern __declspec(dllexport) void QInputDialog_SetOptions(QInputDialog* self, InputDialogOptions options);
extern __declspec(dllexport) InputDialogOptions QInputDialog_Options(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetTextValue(QInputDialog* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QInputDialog_TextValue(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetTextEchoMode(QInputDialog* self, int mode);
extern __declspec(dllexport) int QInputDialog_TextEchoMode(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetComboBoxEditable(QInputDialog* self, bool editable);
extern __declspec(dllexport) bool QInputDialog_IsComboBoxEditable(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetComboBoxItems(QInputDialog* self, struct miqt_array /* of struct miqt_string */  items);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QInputDialog_ComboBoxItems(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetIntValue(QInputDialog* self, int value);
extern __declspec(dllexport) int QInputDialog_IntValue(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetIntMinimum(QInputDialog* self, int min);
extern __declspec(dllexport) int QInputDialog_IntMinimum(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetIntMaximum(QInputDialog* self, int max);
extern __declspec(dllexport) int QInputDialog_IntMaximum(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetIntRange(QInputDialog* self, int min, int max);
extern __declspec(dllexport) void QInputDialog_SetIntStep(QInputDialog* self, int step);
extern __declspec(dllexport) int QInputDialog_IntStep(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetDoubleValue(QInputDialog* self, double value);
extern __declspec(dllexport) double QInputDialog_DoubleValue(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetDoubleMinimum(QInputDialog* self, double min);
extern __declspec(dllexport) double QInputDialog_DoubleMinimum(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetDoubleMaximum(QInputDialog* self, double max);
extern __declspec(dllexport) double QInputDialog_DoubleMaximum(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetDoubleRange(QInputDialog* self, double min, double max);
extern __declspec(dllexport) void QInputDialog_SetDoubleDecimals(QInputDialog* self, int decimals);
extern __declspec(dllexport) int QInputDialog_DoubleDecimals(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetOkButtonText(QInputDialog* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QInputDialog_OkButtonText(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetCancelButtonText(QInputDialog* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QInputDialog_CancelButtonText(const QInputDialog* self);
extern __declspec(dllexport) QSize* QInputDialog_MinimumSizeHint(const QInputDialog* self);
extern __declspec(dllexport) QSize* QInputDialog_SizeHint(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_SetVisible(QInputDialog* self, bool visible);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetText(QWidget* parent, struct miqt_string title, struct miqt_string label);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetMultiLineText(QWidget* parent, struct miqt_string title, struct miqt_string label);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetItem(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_array /* of struct miqt_string */  items);
extern __declspec(dllexport) int QInputDialog_GetInt(QWidget* parent, struct miqt_string title, struct miqt_string label);
extern __declspec(dllexport) double QInputDialog_GetDouble(QWidget* parent, struct miqt_string title, struct miqt_string label);
extern __declspec(dllexport) void QInputDialog_SetDoubleStep(QInputDialog* self, double step);
extern __declspec(dllexport) double QInputDialog_DoubleStep(const QInputDialog* self);
extern __declspec(dllexport) void QInputDialog_TextValueChanged(QInputDialog* self, struct miqt_string text);
void QInputDialog_connect_TextValueChanged(QInputDialog* self, intptr_t slot);
extern __declspec(dllexport) void QInputDialog_TextValueSelected(QInputDialog* self, struct miqt_string text);
void QInputDialog_connect_TextValueSelected(QInputDialog* self, intptr_t slot);
extern __declspec(dllexport) void QInputDialog_IntValueChanged(QInputDialog* self, int value);
void QInputDialog_connect_IntValueChanged(QInputDialog* self, intptr_t slot);
extern __declspec(dllexport) void QInputDialog_IntValueSelected(QInputDialog* self, int value);
void QInputDialog_connect_IntValueSelected(QInputDialog* self, intptr_t slot);
extern __declspec(dllexport) void QInputDialog_DoubleValueChanged(QInputDialog* self, double value);
void QInputDialog_connect_DoubleValueChanged(QInputDialog* self, intptr_t slot);
extern __declspec(dllexport) void QInputDialog_DoubleValueSelected(QInputDialog* self, double value);
void QInputDialog_connect_DoubleValueSelected(QInputDialog* self, intptr_t slot);
extern __declspec(dllexport) void QInputDialog_Done(QInputDialog* self, int result);
extern __declspec(dllexport) struct miqt_string QInputDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QInputDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QInputDialog_SetOption2(QInputDialog* self, InputDialogOption option, bool on);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetText4(QWidget* parent, struct miqt_string title, struct miqt_string label, int echo);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetText5(QWidget* parent, struct miqt_string title, struct miqt_string label, int echo, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetText6(QWidget* parent, struct miqt_string title, struct miqt_string label, int echo, struct miqt_string text, bool* ok);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetText7(QWidget* parent, struct miqt_string title, struct miqt_string label, int echo, struct miqt_string text, bool* ok, int flags);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetText8(QWidget* parent, struct miqt_string title, struct miqt_string label, int echo, struct miqt_string text, bool* ok, int flags, int inputMethodHints);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetMultiLineText4(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetMultiLineText5(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_string text, bool* ok);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetMultiLineText6(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_string text, bool* ok, int flags);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetMultiLineText7(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_string text, bool* ok, int flags, int inputMethodHints);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetItem5(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_array /* of struct miqt_string */  items, int current);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetItem6(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_array /* of struct miqt_string */  items, int current, bool editable);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetItem7(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_array /* of struct miqt_string */  items, int current, bool editable, bool* ok);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetItem8(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_array /* of struct miqt_string */  items, int current, bool editable, bool* ok, int flags);
extern __declspec(dllexport) struct miqt_string QInputDialog_GetItem9(QWidget* parent, struct miqt_string title, struct miqt_string label, struct miqt_array /* of struct miqt_string */  items, int current, bool editable, bool* ok, int flags, int inputMethodHints);
extern __declspec(dllexport) int QInputDialog_GetInt4(QWidget* parent, struct miqt_string title, struct miqt_string label, int value);
extern __declspec(dllexport) int QInputDialog_GetInt5(QWidget* parent, struct miqt_string title, struct miqt_string label, int value, int minValue);
extern __declspec(dllexport) int QInputDialog_GetInt6(QWidget* parent, struct miqt_string title, struct miqt_string label, int value, int minValue, int maxValue);
extern __declspec(dllexport) int QInputDialog_GetInt7(QWidget* parent, struct miqt_string title, struct miqt_string label, int value, int minValue, int maxValue, int step);
extern __declspec(dllexport) int QInputDialog_GetInt8(QWidget* parent, struct miqt_string title, struct miqt_string label, int value, int minValue, int maxValue, int step, bool* ok);
extern __declspec(dllexport) int QInputDialog_GetInt9(QWidget* parent, struct miqt_string title, struct miqt_string label, int value, int minValue, int maxValue, int step, bool* ok, int flags);
extern __declspec(dllexport) double QInputDialog_GetDouble4(QWidget* parent, struct miqt_string title, struct miqt_string label, double value);
extern __declspec(dllexport) double QInputDialog_GetDouble5(QWidget* parent, struct miqt_string title, struct miqt_string label, double value, double minValue);
extern __declspec(dllexport) double QInputDialog_GetDouble6(QWidget* parent, struct miqt_string title, struct miqt_string label, double value, double minValue, double maxValue);
extern __declspec(dllexport) double QInputDialog_GetDouble7(QWidget* parent, struct miqt_string title, struct miqt_string label, double value, double minValue, double maxValue, int decimals);
extern __declspec(dllexport) double QInputDialog_GetDouble8(QWidget* parent, struct miqt_string title, struct miqt_string label, double value, double minValue, double maxValue, int decimals, bool* ok);
extern __declspec(dllexport) double QInputDialog_GetDouble9(QWidget* parent, struct miqt_string title, struct miqt_string label, double value, double minValue, double maxValue, int decimals, bool* ok, int flags);
extern __declspec(dllexport) double QInputDialog_GetDouble10(QWidget* parent, struct miqt_string title, struct miqt_string label, double value, double minValue, double maxValue, int decimals, bool* ok, int flags, double step);
extern __declspec(dllexport) void QInputDialog_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QInputDialog_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QInputDialog_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QInputDialog_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QInputDialog_override_virtual_SetVisible(void* self, intptr_t slot);
void QInputDialog_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QInputDialog_override_virtual_Done(void* self, intptr_t slot);
void QInputDialog_virtualbase_Done(void* self, int result);
extern __declspec(dllexport) void QInputDialog_override_virtual_Open(void* self, intptr_t slot);
void QInputDialog_virtualbase_Open(void* self);
extern __declspec(dllexport) void QInputDialog_override_virtual_Exec(void* self, intptr_t slot);
int QInputDialog_virtualbase_Exec(void* self);
extern __declspec(dllexport) void QInputDialog_override_virtual_Accept(void* self, intptr_t slot);
void QInputDialog_virtualbase_Accept(void* self);
extern __declspec(dllexport) void QInputDialog_override_virtual_Reject(void* self, intptr_t slot);
void QInputDialog_virtualbase_Reject(void* self);
extern __declspec(dllexport) void QInputDialog_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QInputDialog_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QInputDialog_override_virtual_CloseEvent(void* self, intptr_t slot);
void QInputDialog_virtualbase_CloseEvent(void* self, QCloseEvent* param1);
extern __declspec(dllexport) void QInputDialog_override_virtual_ShowEvent(void* self, intptr_t slot);
void QInputDialog_virtualbase_ShowEvent(void* self, QShowEvent* param1);
extern __declspec(dllexport) void QInputDialog_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QInputDialog_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
extern __declspec(dllexport) void QInputDialog_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QInputDialog_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QInputDialog_override_virtual_EventFilter(void* self, intptr_t slot);
bool QInputDialog_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QInputDialog_Delete(QInputDialog* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
