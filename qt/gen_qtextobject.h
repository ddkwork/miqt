#pragma once
#ifndef MIQT_QT_GEN_QTEXTOBJECT_H
#define MIQT_QT_GEN_QTEXTOBJECT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextBlock__iterator)
typedef QTextBlock::iterator QTextBlock__iterator;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextFrame__iterator)
typedef QTextFrame::iterator QTextFrame__iterator;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextLayout__FormatRange)
typedef QTextLayout::FormatRange QTextLayout__FormatRange;
typedef struct QGlyphRun QGlyphRun;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTextBlock QTextBlock;
typedef struct QTextBlock__iterator QTextBlock__iterator;
typedef struct QTextBlockFormat QTextBlockFormat;
typedef struct QTextBlockGroup QTextBlockGroup;
typedef struct QTextBlockUserData QTextBlockUserData;
typedef struct QTextCharFormat QTextCharFormat;
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTextFormat QTextFormat;
typedef struct QTextFragment QTextFragment;
typedef struct QTextFrame QTextFrame;
typedef struct QTextFrame__iterator QTextFrame__iterator;
typedef struct QTextFrameFormat QTextFrameFormat;
typedef struct QTextFrameLayoutData QTextFrameLayoutData;
typedef struct QTextLayout QTextLayout;
typedef struct QTextLayout__FormatRange QTextLayout__FormatRange;
typedef struct QTextList QTextList;
typedef struct QTextObject QTextObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
void QTextObject_virtbase(QTextObject* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QTextObject_MetaObject(const QTextObject* self);
extern __declspec(dllexport) 
void* QTextObject_Metacast(QTextObject* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTextObject_Tr(const char* s);
extern __declspec(dllexport) 
QTextFormat* QTextObject_Format(const QTextObject* self);
extern __declspec(dllexport) 
int QTextObject_FormatIndex(const QTextObject* self);
extern __declspec(dllexport) 
QTextDocument* QTextObject_Document(const QTextObject* self);
extern __declspec(dllexport) 
int QTextObject_ObjectIndex(const QTextObject* self);
extern __declspec(dllexport) 
struct miqt_string QTextObject_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTextObject_Tr3(const char* s, const char* c, int n);

extern __declspec(dllexport) 
void QTextBlockGroup_virtbase(QTextBlockGroup* src
, QTextObject** outptr_QTextObject
);
extern __declspec(dllexport) 
QMetaObject* QTextBlockGroup_MetaObject(const QTextBlockGroup* self);
extern __declspec(dllexport) 
void* QTextBlockGroup_Metacast(QTextBlockGroup* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTextBlockGroup_Tr(const char* s);
extern __declspec(dllexport) 
void QTextBlockGroup_BlockInserted(QTextBlockGroup* self, QTextBlock* block);
extern __declspec(dllexport) 
void QTextBlockGroup_BlockRemoved(QTextBlockGroup* self, QTextBlock* block);
extern __declspec(dllexport) 
void QTextBlockGroup_BlockFormatChanged(QTextBlockGroup* self, QTextBlock* block);
extern __declspec(dllexport) 
struct miqt_string QTextBlockGroup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTextBlockGroup_Tr3(const char* s, const char* c, int n);

extern __declspec(dllexport) 
void QTextFrameLayoutData_OperatorAssign(QTextFrameLayoutData* self, QTextFrameLayoutData* param1);
extern __declspec(dllexport) 
void QTextFrameLayoutData_Delete(QTextFrameLayoutData* self, bool isSubclass);

extern __declspec(dllexport) 
QTextFrame* QTextFrame_new(QTextDocument* doc);
extern __declspec(dllexport) 
void QTextFrame_virtbase(QTextFrame* src
, QTextObject** outptr_QTextObject
);
extern __declspec(dllexport) 
QMetaObject* QTextFrame_MetaObject(const QTextFrame* self);
extern __declspec(dllexport) 
void* QTextFrame_Metacast(QTextFrame* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTextFrame_Tr(const char* s);
extern __declspec(dllexport) 
void QTextFrame_SetFrameFormat(QTextFrame* self, QTextFrameFormat* format);
extern __declspec(dllexport) 
QTextFrameFormat* QTextFrame_FrameFormat(const QTextFrame* self);
extern __declspec(dllexport) 
QTextCursor* QTextFrame_FirstCursorPosition(const QTextFrame* self);
extern __declspec(dllexport) 
QTextCursor* QTextFrame_LastCursorPosition(const QTextFrame* self);
extern __declspec(dllexport) 
int QTextFrame_FirstPosition(const QTextFrame* self);
extern __declspec(dllexport) 
int QTextFrame_LastPosition(const QTextFrame* self);
extern __declspec(dllexport) 
QTextFrameLayoutData* QTextFrame_LayoutData(const QTextFrame* self);
extern __declspec(dllexport) 
void QTextFrame_SetLayoutData(QTextFrame* self, QTextFrameLayoutData* data);
extern __declspec(dllexport) 
struct miqt_array /* of QTextFrame* */  QTextFrame_ChildFrames(const QTextFrame* self);
extern __declspec(dllexport) 
QTextFrame* QTextFrame_ParentFrame(const QTextFrame* self);
extern __declspec(dllexport) 
iterator QTextFrame_Begin(const QTextFrame* self);
extern __declspec(dllexport) 
iterator QTextFrame_End(const QTextFrame* self);
extern __declspec(dllexport) 
struct miqt_string QTextFrame_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTextFrame_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTextFrame_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTextFrame_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTextFrame_override_virtual_Metacast(void* self, intptr_t slot);
void* QTextFrame_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTextFrame_Delete(QTextFrame* self, bool isSubclass);

extern __declspec(dllexport) 
void QTextBlockUserData_OperatorAssign(QTextBlockUserData* self, QTextBlockUserData* param1);
extern __declspec(dllexport) 
void QTextBlockUserData_Delete(QTextBlockUserData* self, bool isSubclass);

extern __declspec(dllexport) 
QTextBlock* QTextBlock_new();
extern __declspec(dllexport) 
QTextBlock* QTextBlock_new2(QTextBlock* o);
extern __declspec(dllexport) 
void QTextBlock_OperatorAssign(QTextBlock* self, QTextBlock* o);
extern __declspec(dllexport) 
bool QTextBlock_IsValid(const QTextBlock* self);
extern __declspec(dllexport) 
bool QTextBlock_OperatorEqual(const QTextBlock* self, QTextBlock* o);
extern __declspec(dllexport) 
bool QTextBlock_OperatorNotEqual(const QTextBlock* self, QTextBlock* o);
extern __declspec(dllexport) 
bool QTextBlock_OperatorLesser(const QTextBlock* self, QTextBlock* o);
extern __declspec(dllexport) 
int QTextBlock_Position(const QTextBlock* self);
extern __declspec(dllexport) 
int QTextBlock_Length(const QTextBlock* self);
extern __declspec(dllexport) 
bool QTextBlock_Contains(const QTextBlock* self, int position);
extern __declspec(dllexport) 
QTextLayout* QTextBlock_Layout(const QTextBlock* self);
extern __declspec(dllexport) 
void QTextBlock_ClearLayout(QTextBlock* self);
extern __declspec(dllexport) 
QTextBlockFormat* QTextBlock_BlockFormat(const QTextBlock* self);
extern __declspec(dllexport) 
int QTextBlock_BlockFormatIndex(const QTextBlock* self);
extern __declspec(dllexport) 
QTextCharFormat* QTextBlock_CharFormat(const QTextBlock* self);
extern __declspec(dllexport) 
int QTextBlock_CharFormatIndex(const QTextBlock* self);
extern __declspec(dllexport) 
int QTextBlock_TextDirection(const QTextBlock* self);
extern __declspec(dllexport) 
struct miqt_string QTextBlock_Text(const QTextBlock* self);
extern __declspec(dllexport) 
struct miqt_array /* of QTextLayout__FormatRange* */  QTextBlock_TextFormats(const QTextBlock* self);
extern __declspec(dllexport) 
QTextDocument* QTextBlock_Document(const QTextBlock* self);
extern __declspec(dllexport) 
QTextList* QTextBlock_TextList(const QTextBlock* self);
extern __declspec(dllexport) 
QTextBlockUserData* QTextBlock_UserData(const QTextBlock* self);
extern __declspec(dllexport) 
void QTextBlock_SetUserData(QTextBlock* self, QTextBlockUserData* data);
extern __declspec(dllexport) 
int QTextBlock_UserState(const QTextBlock* self);
extern __declspec(dllexport) 
void QTextBlock_SetUserState(QTextBlock* self, int state);
extern __declspec(dllexport) 
int QTextBlock_Revision(const QTextBlock* self);
extern __declspec(dllexport) 
void QTextBlock_SetRevision(QTextBlock* self, int rev);
extern __declspec(dllexport) 
bool QTextBlock_IsVisible(const QTextBlock* self);
extern __declspec(dllexport) 
void QTextBlock_SetVisible(QTextBlock* self, bool visible);
extern __declspec(dllexport) 
int QTextBlock_BlockNumber(const QTextBlock* self);
extern __declspec(dllexport) 
int QTextBlock_FirstLineNumber(const QTextBlock* self);
extern __declspec(dllexport) 
void QTextBlock_SetLineCount(QTextBlock* self, int count);
extern __declspec(dllexport) 
int QTextBlock_LineCount(const QTextBlock* self);
extern __declspec(dllexport) 
iterator QTextBlock_Begin(const QTextBlock* self);
extern __declspec(dllexport) 
iterator QTextBlock_End(const QTextBlock* self);
extern __declspec(dllexport) 
QTextBlock* QTextBlock_Next(const QTextBlock* self);
extern __declspec(dllexport) 
QTextBlock* QTextBlock_Previous(const QTextBlock* self);
extern __declspec(dllexport) 
int QTextBlock_FragmentIndex(const QTextBlock* self);
extern __declspec(dllexport) 
void QTextBlock_Delete(QTextBlock* self, bool isSubclass);

extern __declspec(dllexport) 
QTextFragment* QTextFragment_new();
extern __declspec(dllexport) 
QTextFragment* QTextFragment_new2(QTextFragment* o);
extern __declspec(dllexport) 
void QTextFragment_OperatorAssign(QTextFragment* self, QTextFragment* o);
extern __declspec(dllexport) 
bool QTextFragment_IsValid(const QTextFragment* self);
extern __declspec(dllexport) 
bool QTextFragment_OperatorEqual(const QTextFragment* self, QTextFragment* o);
extern __declspec(dllexport) 
bool QTextFragment_OperatorNotEqual(const QTextFragment* self, QTextFragment* o);
extern __declspec(dllexport) 
bool QTextFragment_OperatorLesser(const QTextFragment* self, QTextFragment* o);
extern __declspec(dllexport) 
int QTextFragment_Position(const QTextFragment* self);
extern __declspec(dllexport) 
int QTextFragment_Length(const QTextFragment* self);
extern __declspec(dllexport) 
bool QTextFragment_Contains(const QTextFragment* self, int position);
extern __declspec(dllexport) 
QTextCharFormat* QTextFragment_CharFormat(const QTextFragment* self);
extern __declspec(dllexport) 
int QTextFragment_CharFormatIndex(const QTextFragment* self);
extern __declspec(dllexport) 
struct miqt_string QTextFragment_Text(const QTextFragment* self);
extern __declspec(dllexport) 
struct miqt_array /* of QGlyphRun* */  QTextFragment_GlyphRuns(const QTextFragment* self);
extern __declspec(dllexport) 
struct miqt_array /* of QGlyphRun* */  QTextFragment_GlyphRuns1(const QTextFragment* self, int from);
extern __declspec(dllexport) 
struct miqt_array /* of QGlyphRun* */  QTextFragment_GlyphRuns2(const QTextFragment* self, int from, int length);
extern __declspec(dllexport) 
void QTextFragment_Delete(QTextFragment* self, bool isSubclass);

extern __declspec(dllexport) 
QTextFrame__iterator* QTextFrame__iterator_new();
extern __declspec(dllexport) 
QTextFrame__iterator* QTextFrame__iterator_new2(const iterator* param1);
extern __declspec(dllexport) 
QTextFrame* QTextFrame__iterator_ParentFrame(const QTextFrame__iterator* self);
extern __declspec(dllexport) 
QTextFrame* QTextFrame__iterator_CurrentFrame(const QTextFrame__iterator* self);
extern __declspec(dllexport) 
QTextBlock* QTextFrame__iterator_CurrentBlock(const QTextFrame__iterator* self);
extern __declspec(dllexport) 
bool QTextFrame__iterator_AtEnd(const QTextFrame__iterator* self);
extern __declspec(dllexport) 
bool QTextFrame__iterator_OperatorEqual(const QTextFrame__iterator* self, const iterator* o);
extern __declspec(dllexport) 
bool QTextFrame__iterator_OperatorNotEqual(const QTextFrame__iterator* self, const iterator* o);
extern __declspec(dllexport) 
iterator* QTextFrame__iterator_OperatorPlusPlus(QTextFrame__iterator* self);
extern __declspec(dllexport) 
iterator QTextFrame__iterator_OperatorPlusPlusWithInt(QTextFrame__iterator* self, int param1);
extern __declspec(dllexport) 
iterator* QTextFrame__iterator_OperatorMinusMinus(QTextFrame__iterator* self);
extern __declspec(dllexport) 
iterator QTextFrame__iterator_OperatorMinusMinusWithInt(QTextFrame__iterator* self, int param1);
extern __declspec(dllexport) 
void QTextFrame__iterator_Delete(QTextFrame__iterator* self, bool isSubclass);

extern __declspec(dllexport) 
QTextBlock__iterator* QTextBlock__iterator_new();
extern __declspec(dllexport) 
QTextBlock__iterator* QTextBlock__iterator_new2(const iterator* param1);
extern __declspec(dllexport) 
QTextFragment* QTextBlock__iterator_Fragment(const QTextBlock__iterator* self);
extern __declspec(dllexport) 
bool QTextBlock__iterator_AtEnd(const QTextBlock__iterator* self);
extern __declspec(dllexport) 
bool QTextBlock__iterator_OperatorEqual(const QTextBlock__iterator* self, const iterator* o);
extern __declspec(dllexport) 
bool QTextBlock__iterator_OperatorNotEqual(const QTextBlock__iterator* self, const iterator* o);
extern __declspec(dllexport) 
iterator* QTextBlock__iterator_OperatorPlusPlus(QTextBlock__iterator* self);
extern __declspec(dllexport) 
iterator QTextBlock__iterator_OperatorPlusPlusWithInt(QTextBlock__iterator* self, int param1);
extern __declspec(dllexport) 
iterator* QTextBlock__iterator_OperatorMinusMinus(QTextBlock__iterator* self);
extern __declspec(dllexport) 
iterator QTextBlock__iterator_OperatorMinusMinusWithInt(QTextBlock__iterator* self, int param1);
extern __declspec(dllexport) 
void QTextBlock__iterator_Delete(QTextBlock__iterator* self, bool isSubclass);

}
