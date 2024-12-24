#pragma once
#ifndef MIQT_QT_GEN_QSYNTAXHIGHLIGHTER_H
#define MIQT_QT_GEN_QSYNTAXHIGHLIGHTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSyntaxHighlighter QSyntaxHighlighter;
typedef struct QTextBlock QTextBlock;
typedef struct QTextDocument QTextDocument;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSyntaxHighlighter* QSyntaxHighlighter_new(QObject* parent);
extern __declspec(dllexport) QSyntaxHighlighter* QSyntaxHighlighter_new2(QTextDocument* parent);
extern __declspec(dllexport) void QSyntaxHighlighter_virtbase(QSyntaxHighlighter* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QSyntaxHighlighter_MetaObject(const QSyntaxHighlighter* self);
extern __declspec(dllexport) void* QSyntaxHighlighter_Metacast(QSyntaxHighlighter* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSyntaxHighlighter_Tr(const char* s);
extern __declspec(dllexport) void QSyntaxHighlighter_SetDocument(QSyntaxHighlighter* self, QTextDocument* doc);
extern __declspec(dllexport) QTextDocument* QSyntaxHighlighter_Document(const QSyntaxHighlighter* self);
extern __declspec(dllexport) void QSyntaxHighlighter_Rehighlight(QSyntaxHighlighter* self);
extern __declspec(dllexport) void QSyntaxHighlighter_RehighlightBlock(QSyntaxHighlighter* self, QTextBlock* block);
extern __declspec(dllexport) void QSyntaxHighlighter_HighlightBlock(QSyntaxHighlighter* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QSyntaxHighlighter_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSyntaxHighlighter_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSyntaxHighlighter_override_virtual_HighlightBlock(void* self, intptr_t slot);
void QSyntaxHighlighter_virtualbase_HighlightBlock(void* self, struct miqt_string text);
extern __declspec(dllexport) void QSyntaxHighlighter_override_virtual_Event(void* self, intptr_t slot);
bool QSyntaxHighlighter_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSyntaxHighlighter_override_virtual_EventFilter(void* self, intptr_t slot);
bool QSyntaxHighlighter_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QSyntaxHighlighter_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSyntaxHighlighter_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QSyntaxHighlighter_override_virtual_ChildEvent(void* self, intptr_t slot);
void QSyntaxHighlighter_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QSyntaxHighlighter_override_virtual_CustomEvent(void* self, intptr_t slot);
void QSyntaxHighlighter_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QSyntaxHighlighter_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QSyntaxHighlighter_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSyntaxHighlighter_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QSyntaxHighlighter_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSyntaxHighlighter_Delete(QSyntaxHighlighter* self, bool isSubclass);

} 
