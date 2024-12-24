#pragma once
#ifndef MIQT_QT_GEN_QXMLSTREAM_H
#define MIQT_QT_GEN_QXMLSTREAM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAnyStringView;
class QIODevice;
class QXmlStreamAttribute;
class QXmlStreamEntityDeclaration;
class QXmlStreamEntityResolver;
class QXmlStreamNamespaceDeclaration;
class QXmlStreamNotationDeclaration;
class QXmlStreamReader;
class QXmlStreamWriter;
class _GUID;
class type_info;
#else
typedef struct QAnyStringView QAnyStringView;
typedef struct QIODevice QIODevice;
typedef struct QXmlStreamAttribute QXmlStreamAttribute;
typedef struct QXmlStreamEntityDeclaration QXmlStreamEntityDeclaration;
typedef struct QXmlStreamEntityResolver QXmlStreamEntityResolver;
typedef struct QXmlStreamNamespaceDeclaration QXmlStreamNamespaceDeclaration;
typedef struct QXmlStreamNotationDeclaration QXmlStreamNotationDeclaration;
typedef struct QXmlStreamReader QXmlStreamReader;
typedef struct QXmlStreamWriter QXmlStreamWriter;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QXmlStreamAttribute* QXmlStreamAttribute_new();
extern __declspec(dllexport) QXmlStreamAttribute* QXmlStreamAttribute_new2(struct miqt_string qualifiedName, struct miqt_string value);
extern __declspec(dllexport) QXmlStreamAttribute* QXmlStreamAttribute_new3(struct miqt_string namespaceUri, struct miqt_string name, struct miqt_string value);
extern __declspec(dllexport) QXmlStreamAttribute* QXmlStreamAttribute_new4(QXmlStreamAttribute* param1);
extern __declspec(dllexport) bool QXmlStreamAttribute_IsDefault(const QXmlStreamAttribute* self);
extern __declspec(dllexport) void QXmlStreamAttribute_Delete(QXmlStreamAttribute* self, bool isSubclass);

extern __declspec(dllexport) QXmlStreamNamespaceDeclaration* QXmlStreamNamespaceDeclaration_new();
extern __declspec(dllexport) QXmlStreamNamespaceDeclaration* QXmlStreamNamespaceDeclaration_new2(struct miqt_string prefix, struct miqt_string namespaceUri);
extern __declspec(dllexport) QXmlStreamNamespaceDeclaration* QXmlStreamNamespaceDeclaration_new3(QXmlStreamNamespaceDeclaration* param1);
extern __declspec(dllexport) void QXmlStreamNamespaceDeclaration_Delete(QXmlStreamNamespaceDeclaration* self, bool isSubclass);

extern __declspec(dllexport) QXmlStreamNotationDeclaration* QXmlStreamNotationDeclaration_new();
extern __declspec(dllexport) QXmlStreamNotationDeclaration* QXmlStreamNotationDeclaration_new2(QXmlStreamNotationDeclaration* param1);
extern __declspec(dllexport) void QXmlStreamNotationDeclaration_Delete(QXmlStreamNotationDeclaration* self, bool isSubclass);

extern __declspec(dllexport) QXmlStreamEntityDeclaration* QXmlStreamEntityDeclaration_new();
extern __declspec(dllexport) QXmlStreamEntityDeclaration* QXmlStreamEntityDeclaration_new2(QXmlStreamEntityDeclaration* param1);
extern __declspec(dllexport) void QXmlStreamEntityDeclaration_Delete(QXmlStreamEntityDeclaration* self, bool isSubclass);

extern __declspec(dllexport) QXmlStreamEntityResolver* QXmlStreamEntityResolver_new();
extern __declspec(dllexport) struct miqt_string QXmlStreamEntityResolver_ResolveEntity(QXmlStreamEntityResolver* self, struct miqt_string publicId, struct miqt_string systemId);
extern __declspec(dllexport) struct miqt_string QXmlStreamEntityResolver_ResolveUndeclaredEntity(QXmlStreamEntityResolver* self, struct miqt_string name);
extern __declspec(dllexport) void QXmlStreamEntityResolver_override_virtual_ResolveEntity(void* self, intptr_t slot);
struct miqt_string QXmlStreamEntityResolver_virtualbase_ResolveEntity(void* self, struct miqt_string publicId, struct miqt_string systemId);
extern __declspec(dllexport) void QXmlStreamEntityResolver_override_virtual_ResolveUndeclaredEntity(void* self, intptr_t slot);
struct miqt_string QXmlStreamEntityResolver_virtualbase_ResolveUndeclaredEntity(void* self, struct miqt_string name);
extern __declspec(dllexport) void QXmlStreamEntityResolver_Delete(QXmlStreamEntityResolver* self, bool isSubclass);

extern __declspec(dllexport) QXmlStreamReader* QXmlStreamReader_new();
extern __declspec(dllexport) QXmlStreamReader* QXmlStreamReader_new2(QIODevice* device);
extern __declspec(dllexport) QXmlStreamReader* QXmlStreamReader_new3(QAnyStringView* data);
extern __declspec(dllexport) void QXmlStreamReader_SetDevice(QXmlStreamReader* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QXmlStreamReader_Device(const QXmlStreamReader* self);
extern __declspec(dllexport) void QXmlStreamReader_AddData(QXmlStreamReader* self, QAnyStringView* data);
extern __declspec(dllexport) void QXmlStreamReader_Clear(QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_AtEnd(const QXmlStreamReader* self);
extern __declspec(dllexport) TokenType QXmlStreamReader_ReadNext(QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_ReadNextStartElement(QXmlStreamReader* self);
extern __declspec(dllexport) void QXmlStreamReader_SkipCurrentElement(QXmlStreamReader* self);
extern __declspec(dllexport) TokenType QXmlStreamReader_TokenType(const QXmlStreamReader* self);
extern __declspec(dllexport) struct miqt_string QXmlStreamReader_TokenString(const QXmlStreamReader* self);
extern __declspec(dllexport) void QXmlStreamReader_SetNamespaceProcessing(QXmlStreamReader* self, bool namespaceProcessing);
extern __declspec(dllexport) bool QXmlStreamReader_NamespaceProcessing(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsStartDocument(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsEndDocument(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsStartElement(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsEndElement(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsCharacters(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsWhitespace(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsCDATA(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsComment(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsDTD(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsEntityReference(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsProcessingInstruction(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_IsStandaloneDocument(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_HasStandaloneDeclaration(const QXmlStreamReader* self);
extern __declspec(dllexport) long long QXmlStreamReader_LineNumber(const QXmlStreamReader* self);
extern __declspec(dllexport) long long QXmlStreamReader_ColumnNumber(const QXmlStreamReader* self);
extern __declspec(dllexport) long long QXmlStreamReader_CharacterOffset(const QXmlStreamReader* self);
extern __declspec(dllexport) struct miqt_string QXmlStreamReader_ReadElementText(QXmlStreamReader* self);
extern __declspec(dllexport) struct miqt_array /* of QXmlStreamNamespaceDeclaration* */  QXmlStreamReader_NamespaceDeclarations(const QXmlStreamReader* self);
extern __declspec(dllexport) void QXmlStreamReader_AddExtraNamespaceDeclaration(QXmlStreamReader* self, QXmlStreamNamespaceDeclaration* extraNamespaceDeclaraction);
extern __declspec(dllexport) void QXmlStreamReader_AddExtraNamespaceDeclarations(QXmlStreamReader* self, struct miqt_array /* of QXmlStreamNamespaceDeclaration* */  extraNamespaceDeclaractions);
extern __declspec(dllexport) struct miqt_array /* of QXmlStreamNotationDeclaration* */  QXmlStreamReader_NotationDeclarations(const QXmlStreamReader* self);
extern __declspec(dllexport) struct miqt_array /* of QXmlStreamEntityDeclaration* */  QXmlStreamReader_EntityDeclarations(const QXmlStreamReader* self);
extern __declspec(dllexport) int QXmlStreamReader_EntityExpansionLimit(const QXmlStreamReader* self);
extern __declspec(dllexport) void QXmlStreamReader_SetEntityExpansionLimit(QXmlStreamReader* self, int limit);
extern __declspec(dllexport) void QXmlStreamReader_RaiseError(QXmlStreamReader* self);
extern __declspec(dllexport) struct miqt_string QXmlStreamReader_ErrorString(const QXmlStreamReader* self);
extern __declspec(dllexport) Error QXmlStreamReader_Error(const QXmlStreamReader* self);
extern __declspec(dllexport) bool QXmlStreamReader_HasError(const QXmlStreamReader* self);
extern __declspec(dllexport) void QXmlStreamReader_SetEntityResolver(QXmlStreamReader* self, QXmlStreamEntityResolver* resolver);
extern __declspec(dllexport) QXmlStreamEntityResolver* QXmlStreamReader_EntityResolver(const QXmlStreamReader* self);
extern __declspec(dllexport) struct miqt_string QXmlStreamReader_ReadElementText1(QXmlStreamReader* self, ReadElementTextBehaviour behaviour);
extern __declspec(dllexport) void QXmlStreamReader_RaiseError1(QXmlStreamReader* self, struct miqt_string message);
extern __declspec(dllexport) void QXmlStreamReader_Delete(QXmlStreamReader* self, bool isSubclass);

extern __declspec(dllexport) QXmlStreamWriter* QXmlStreamWriter_new();
extern __declspec(dllexport) QXmlStreamWriter* QXmlStreamWriter_new2(QIODevice* device);
extern __declspec(dllexport) void QXmlStreamWriter_SetDevice(QXmlStreamWriter* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QXmlStreamWriter_Device(const QXmlStreamWriter* self);
extern __declspec(dllexport) void QXmlStreamWriter_SetAutoFormatting(QXmlStreamWriter* self, bool autoFormatting);
extern __declspec(dllexport) bool QXmlStreamWriter_AutoFormatting(const QXmlStreamWriter* self);
extern __declspec(dllexport) void QXmlStreamWriter_SetAutoFormattingIndent(QXmlStreamWriter* self, int spacesOrTabs);
extern __declspec(dllexport) int QXmlStreamWriter_AutoFormattingIndent(const QXmlStreamWriter* self);
extern __declspec(dllexport) void QXmlStreamWriter_WriteAttribute(QXmlStreamWriter* self, QAnyStringView* qualifiedName, QAnyStringView* value);
extern __declspec(dllexport) void QXmlStreamWriter_WriteAttribute2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* name, QAnyStringView* value);
extern __declspec(dllexport) void QXmlStreamWriter_WriteAttributeWithAttribute(QXmlStreamWriter* self, QXmlStreamAttribute* attribute);
extern __declspec(dllexport) void QXmlStreamWriter_WriteCDATA(QXmlStreamWriter* self, QAnyStringView* text);
extern __declspec(dllexport) void QXmlStreamWriter_WriteCharacters(QXmlStreamWriter* self, QAnyStringView* text);
extern __declspec(dllexport) void QXmlStreamWriter_WriteComment(QXmlStreamWriter* self, QAnyStringView* text);
extern __declspec(dllexport) void QXmlStreamWriter_WriteDTD(QXmlStreamWriter* self, QAnyStringView* dtd);
extern __declspec(dllexport) void QXmlStreamWriter_WriteEmptyElement(QXmlStreamWriter* self, QAnyStringView* qualifiedName);
extern __declspec(dllexport) void QXmlStreamWriter_WriteEmptyElement2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* name);
extern __declspec(dllexport) void QXmlStreamWriter_WriteTextElement(QXmlStreamWriter* self, QAnyStringView* qualifiedName, QAnyStringView* text);
extern __declspec(dllexport) void QXmlStreamWriter_WriteTextElement2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* name, QAnyStringView* text);
extern __declspec(dllexport) void QXmlStreamWriter_WriteEndDocument(QXmlStreamWriter* self);
extern __declspec(dllexport) void QXmlStreamWriter_WriteEndElement(QXmlStreamWriter* self);
extern __declspec(dllexport) void QXmlStreamWriter_WriteEntityReference(QXmlStreamWriter* self, QAnyStringView* name);
extern __declspec(dllexport) void QXmlStreamWriter_WriteNamespace(QXmlStreamWriter* self, QAnyStringView* namespaceUri);
extern __declspec(dllexport) void QXmlStreamWriter_WriteDefaultNamespace(QXmlStreamWriter* self, QAnyStringView* namespaceUri);
extern __declspec(dllexport) void QXmlStreamWriter_WriteProcessingInstruction(QXmlStreamWriter* self, QAnyStringView* target);
extern __declspec(dllexport) void QXmlStreamWriter_WriteStartDocument(QXmlStreamWriter* self);
extern __declspec(dllexport) void QXmlStreamWriter_WriteStartDocumentWithVersion(QXmlStreamWriter* self, QAnyStringView* version);
extern __declspec(dllexport) void QXmlStreamWriter_WriteStartDocument2(QXmlStreamWriter* self, QAnyStringView* version, bool standalone);
extern __declspec(dllexport) void QXmlStreamWriter_WriteStartElement(QXmlStreamWriter* self, QAnyStringView* qualifiedName);
extern __declspec(dllexport) void QXmlStreamWriter_WriteStartElement2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* name);
extern __declspec(dllexport) void QXmlStreamWriter_WriteCurrentToken(QXmlStreamWriter* self, QXmlStreamReader* reader);
extern __declspec(dllexport) bool QXmlStreamWriter_HasError(const QXmlStreamWriter* self);
extern __declspec(dllexport) void QXmlStreamWriter_WriteNamespace2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* prefix);
extern __declspec(dllexport) void QXmlStreamWriter_WriteProcessingInstruction2(QXmlStreamWriter* self, QAnyStringView* target, QAnyStringView* data);
extern __declspec(dllexport) void QXmlStreamWriter_Delete(QXmlStreamWriter* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
