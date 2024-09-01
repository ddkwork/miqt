#ifndef GEN_QSTATICTEXT_H
#define GEN_QSTATICTEXT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QFont;
class QSizeF;
class QStaticText;
class QTextOption;
class QTransform;
#else
typedef struct QFont QFont;
typedef struct QSizeF QSizeF;
typedef struct QStaticText QStaticText;
typedef struct QTextOption QTextOption;
typedef struct QTransform QTransform;
#endif

QStaticText* QStaticText_new();
QStaticText* QStaticText_new2(const char* text, size_t text_Strlen);
QStaticText* QStaticText_new3(QStaticText* other);
void QStaticText_OperatorAssign(QStaticText* self, QStaticText* param1);
void QStaticText_Swap(QStaticText* self, QStaticText* other);
void QStaticText_SetText(QStaticText* self, const char* text, size_t text_Strlen);
void QStaticText_Text(QStaticText* self, char** _out, int* _out_Strlen);
void QStaticText_SetTextFormat(QStaticText* self, uintptr_t textFormat);
uintptr_t QStaticText_TextFormat(QStaticText* self);
void QStaticText_SetTextWidth(QStaticText* self, double textWidth);
double QStaticText_TextWidth(QStaticText* self);
void QStaticText_SetTextOption(QStaticText* self, QTextOption* textOption);
QTextOption* QStaticText_TextOption(QStaticText* self);
QSizeF* QStaticText_Size(QStaticText* self);
void QStaticText_Prepare(QStaticText* self);
void QStaticText_SetPerformanceHint(QStaticText* self, uintptr_t performanceHint);
uintptr_t QStaticText_PerformanceHint(QStaticText* self);
bool QStaticText_OperatorEqual(QStaticText* self, QStaticText* param1);
bool QStaticText_OperatorNotEqual(QStaticText* self, QStaticText* param1);
void QStaticText_Prepare1(QStaticText* self, QTransform* matrix);
void QStaticText_Prepare2(QStaticText* self, QTransform* matrix, QFont* font);
void QStaticText_Delete(QStaticText* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif