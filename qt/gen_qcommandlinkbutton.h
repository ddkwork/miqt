#ifndef GEN_QCOMMANDLINKBUTTON_H
#define GEN_QCOMMANDLINKBUTTON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCommandLinkButton;
class QMetaObject;
class QWidget;
#else
typedef struct QCommandLinkButton QCommandLinkButton;
typedef struct QMetaObject QMetaObject;
typedef struct QWidget QWidget;
#endif

QCommandLinkButton* QCommandLinkButton_new();
QCommandLinkButton* QCommandLinkButton_new2(const char* text, size_t text_Strlen);
QCommandLinkButton* QCommandLinkButton_new3(const char* text, size_t text_Strlen, const char* description, size_t description_Strlen);
QCommandLinkButton* QCommandLinkButton_new4(QWidget* parent);
QCommandLinkButton* QCommandLinkButton_new5(const char* text, size_t text_Strlen, QWidget* parent);
QCommandLinkButton* QCommandLinkButton_new6(const char* text, size_t text_Strlen, const char* description, size_t description_Strlen, QWidget* parent);
QMetaObject* QCommandLinkButton_MetaObject(const QCommandLinkButton* self);
void QCommandLinkButton_Tr(const char* s, char** _out, int* _out_Strlen);
void QCommandLinkButton_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QCommandLinkButton_Description(const QCommandLinkButton* self, char** _out, int* _out_Strlen);
void QCommandLinkButton_SetDescription(QCommandLinkButton* self, const char* description, size_t description_Strlen);
void QCommandLinkButton_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QCommandLinkButton_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QCommandLinkButton_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QCommandLinkButton_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QCommandLinkButton_Delete(QCommandLinkButton* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
