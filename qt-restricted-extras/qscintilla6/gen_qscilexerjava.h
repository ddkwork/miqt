#pragma once
#ifndef MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA6_GEN_QSCILEXERJAVA_H
#define MIQT_QT_RESTRICTED_EXTRAS_QSCINTILLA6_GEN_QSCILEXERJAVA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QColor;
class QEvent;
class QFont;
class QMetaMethod;
class QMetaObject;
class QObject;
class QSettings;
class QTimerEvent;
class QsciLexer;
class QsciLexerCPP;
class QsciLexerJava;
class QsciScintilla;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QColor QColor;
typedef struct QEvent QEvent;
typedef struct QFont QFont;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSettings QSettings;
typedef struct QTimerEvent QTimerEvent;
typedef struct QsciLexer QsciLexer;
typedef struct QsciLexerCPP QsciLexerCPP;
typedef struct QsciLexerJava QsciLexerJava;
typedef struct QsciScintilla QsciScintilla;
#endif

QsciLexerJava* QsciLexerJava_new();
QsciLexerJava* QsciLexerJava_new2(QObject* parent);
void QsciLexerJava_virtbase(QsciLexerJava* src, QsciLexerCPP** outptr_QsciLexerCPP);
QMetaObject* QsciLexerJava_metaObject(const QsciLexerJava* self);
void* QsciLexerJava_metacast(QsciLexerJava* self, const char* param1);
struct miqt_string QsciLexerJava_tr(const char* s);
const char* QsciLexerJava_language(const QsciLexerJava* self);
const char* QsciLexerJava_keywords(const QsciLexerJava* self, int set);
struct miqt_string QsciLexerJava_tr2(const char* s, const char* c);
struct miqt_string QsciLexerJava_tr3(const char* s, const char* c, int n);
bool QsciLexerJava_override_virtual_setFoldAtElse(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setFoldAtElse(void* self, bool fold);
bool QsciLexerJava_override_virtual_setFoldComments(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setFoldComments(void* self, bool fold);
bool QsciLexerJava_override_virtual_setFoldCompact(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setFoldCompact(void* self, bool fold);
bool QsciLexerJava_override_virtual_setFoldPreprocessor(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setFoldPreprocessor(void* self, bool fold);
bool QsciLexerJava_override_virtual_setStylePreprocessor(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setStylePreprocessor(void* self, bool style);
bool QsciLexerJava_override_virtual_language(void* self, intptr_t slot);
const char* QsciLexerJava_virtualbase_language(const void* self);
bool QsciLexerJava_override_virtual_lexer(void* self, intptr_t slot);
const char* QsciLexerJava_virtualbase_lexer(const void* self);
bool QsciLexerJava_override_virtual_lexerId(void* self, intptr_t slot);
int QsciLexerJava_virtualbase_lexerId(const void* self);
bool QsciLexerJava_override_virtual_autoCompletionFillups(void* self, intptr_t slot);
const char* QsciLexerJava_virtualbase_autoCompletionFillups(const void* self);
bool QsciLexerJava_override_virtual_autoCompletionWordSeparators(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QsciLexerJava_virtualbase_autoCompletionWordSeparators(const void* self);
bool QsciLexerJava_override_virtual_blockEnd(void* self, intptr_t slot);
const char* QsciLexerJava_virtualbase_blockEnd(const void* self, int* style);
bool QsciLexerJava_override_virtual_blockLookback(void* self, intptr_t slot);
int QsciLexerJava_virtualbase_blockLookback(const void* self);
bool QsciLexerJava_override_virtual_blockStart(void* self, intptr_t slot);
const char* QsciLexerJava_virtualbase_blockStart(const void* self, int* style);
bool QsciLexerJava_override_virtual_blockStartKeyword(void* self, intptr_t slot);
const char* QsciLexerJava_virtualbase_blockStartKeyword(const void* self, int* style);
bool QsciLexerJava_override_virtual_braceStyle(void* self, intptr_t slot);
int QsciLexerJava_virtualbase_braceStyle(const void* self);
bool QsciLexerJava_override_virtual_caseSensitive(void* self, intptr_t slot);
bool QsciLexerJava_virtualbase_caseSensitive(const void* self);
bool QsciLexerJava_override_virtual_color(void* self, intptr_t slot);
QColor* QsciLexerJava_virtualbase_color(const void* self, int style);
bool QsciLexerJava_override_virtual_eolFill(void* self, intptr_t slot);
bool QsciLexerJava_virtualbase_eolFill(const void* self, int style);
bool QsciLexerJava_override_virtual_font(void* self, intptr_t slot);
QFont* QsciLexerJava_virtualbase_font(const void* self, int style);
bool QsciLexerJava_override_virtual_indentationGuideView(void* self, intptr_t slot);
int QsciLexerJava_virtualbase_indentationGuideView(const void* self);
bool QsciLexerJava_override_virtual_keywords(void* self, intptr_t slot);
const char* QsciLexerJava_virtualbase_keywords(const void* self, int set);
bool QsciLexerJava_override_virtual_defaultStyle(void* self, intptr_t slot);
int QsciLexerJava_virtualbase_defaultStyle(const void* self);
bool QsciLexerJava_override_virtual_description(void* self, intptr_t slot);
struct miqt_string QsciLexerJava_virtualbase_description(const void* self, int style);
bool QsciLexerJava_override_virtual_paper(void* self, intptr_t slot);
QColor* QsciLexerJava_virtualbase_paper(const void* self, int style);
bool QsciLexerJava_override_virtual_defaultColorWithStyle(void* self, intptr_t slot);
QColor* QsciLexerJava_virtualbase_defaultColorWithStyle(const void* self, int style);
bool QsciLexerJava_override_virtual_defaultEolFill(void* self, intptr_t slot);
bool QsciLexerJava_virtualbase_defaultEolFill(const void* self, int style);
bool QsciLexerJava_override_virtual_defaultFontWithStyle(void* self, intptr_t slot);
QFont* QsciLexerJava_virtualbase_defaultFontWithStyle(const void* self, int style);
bool QsciLexerJava_override_virtual_defaultPaperWithStyle(void* self, intptr_t slot);
QColor* QsciLexerJava_virtualbase_defaultPaperWithStyle(const void* self, int style);
bool QsciLexerJava_override_virtual_setEditor(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setEditor(void* self, QsciScintilla* editor);
bool QsciLexerJava_override_virtual_refreshProperties(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_refreshProperties(void* self);
bool QsciLexerJava_override_virtual_styleBitsNeeded(void* self, intptr_t slot);
int QsciLexerJava_virtualbase_styleBitsNeeded(const void* self);
bool QsciLexerJava_override_virtual_wordCharacters(void* self, intptr_t slot);
const char* QsciLexerJava_virtualbase_wordCharacters(const void* self);
bool QsciLexerJava_override_virtual_setAutoIndentStyle(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setAutoIndentStyle(void* self, int autoindentstyle);
bool QsciLexerJava_override_virtual_setColor(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setColor(void* self, QColor* c, int style);
bool QsciLexerJava_override_virtual_setEolFill(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setEolFill(void* self, bool eoffill, int style);
bool QsciLexerJava_override_virtual_setFont(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setFont(void* self, QFont* f, int style);
bool QsciLexerJava_override_virtual_setPaper(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_setPaper(void* self, QColor* c, int style);
bool QsciLexerJava_override_virtual_readProperties(void* self, intptr_t slot);
bool QsciLexerJava_virtualbase_readProperties(void* self, QSettings* qs, struct miqt_string prefix);
bool QsciLexerJava_override_virtual_writeProperties(void* self, intptr_t slot);
bool QsciLexerJava_virtualbase_writeProperties(const void* self, QSettings* qs, struct miqt_string prefix);
bool QsciLexerJava_override_virtual_event(void* self, intptr_t slot);
bool QsciLexerJava_virtualbase_event(void* self, QEvent* event);
bool QsciLexerJava_override_virtual_eventFilter(void* self, intptr_t slot);
bool QsciLexerJava_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event);
bool QsciLexerJava_override_virtual_timerEvent(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_timerEvent(void* self, QTimerEvent* event);
bool QsciLexerJava_override_virtual_childEvent(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_childEvent(void* self, QChildEvent* event);
bool QsciLexerJava_override_virtual_customEvent(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_customEvent(void* self, QEvent* event);
bool QsciLexerJava_override_virtual_connectNotify(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_connectNotify(void* self, QMetaMethod* signal);
bool QsciLexerJava_override_virtual_disconnectNotify(void* self, intptr_t slot);
void QsciLexerJava_virtualbase_disconnectNotify(void* self, QMetaMethod* signal);
bool QsciLexerJava_protectedbase_readProperties(bool* _dynamic_cast_ok, void* self, QSettings* qs, struct miqt_string prefix);
bool QsciLexerJava_protectedbase_writeProperties(bool* _dynamic_cast_ok, const void* self, QSettings* qs, struct miqt_string prefix);
QObject* QsciLexerJava_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
int QsciLexerJava_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
int QsciLexerJava_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
bool QsciLexerJava_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);
void QsciLexerJava_delete(QsciLexerJava* self);

#ifdef __cplusplus
} /* extern C */
#endif

#endif
