#ifndef GEN_QSYNTAXHIGHLIGHTER_H
#define GEN_QSYNTAXHIGHLIGHTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMetaObject;
class QSyntaxHighlighter;
class QTextBlock;
class QTextDocument;
#else
typedef struct QMetaObject QMetaObject;
typedef struct QSyntaxHighlighter QSyntaxHighlighter;
typedef struct QTextBlock QTextBlock;
typedef struct QTextDocument QTextDocument;
#endif

QMetaObject* QSyntaxHighlighter_MetaObject(const QSyntaxHighlighter* self);
void QSyntaxHighlighter_Tr(const char* s, char** _out, int* _out_Strlen);
void QSyntaxHighlighter_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QSyntaxHighlighter_SetDocument(QSyntaxHighlighter* self, QTextDocument* doc);
QTextDocument* QSyntaxHighlighter_Document(const QSyntaxHighlighter* self);
void QSyntaxHighlighter_Rehighlight(QSyntaxHighlighter* self);
void QSyntaxHighlighter_RehighlightBlock(QSyntaxHighlighter* self, QTextBlock* block);
void QSyntaxHighlighter_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QSyntaxHighlighter_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QSyntaxHighlighter_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QSyntaxHighlighter_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QSyntaxHighlighter_Delete(QSyntaxHighlighter* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
