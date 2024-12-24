// +build ignore

#include <QAnyStringView>
#include <QIODevice>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QXmlStreamAttribute>
#include <QXmlStreamEntityDeclaration>
#include <QXmlStreamEntityResolver>
#include <QXmlStreamNamespaceDeclaration>
#include <QXmlStreamNotationDeclaration>
#include <QXmlStreamReader>
#include <QXmlStreamWriter>
#include <qxmlstream.h>
#include "gen_qxmlstream.h"
QXmlStreamAttribute* QXmlStreamAttribute_new() {
return new QXmlStreamAttribute();
}
QXmlStreamAttribute* QXmlStreamAttribute_new2(struct miqt_string qualifiedName, struct miqt_string value) {
QString qualifiedName_QString = QString::fromUtf8(qualifiedName.data, qualifiedName.len);
	QString value_QString = QString::fromUtf8(value.data, value.len);
	return new QXmlStreamAttribute(qualifiedName_QString, value_QString);
}
QXmlStreamAttribute* QXmlStreamAttribute_new3(struct miqt_string namespaceUri, struct miqt_string name, struct miqt_string value) {
QString namespaceUri_QString = QString::fromUtf8(namespaceUri.data, namespaceUri.len);
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString value_QString = QString::fromUtf8(value.data, value.len);
	return new QXmlStreamAttribute(namespaceUri_QString, name_QString, value_QString);
}
QXmlStreamAttribute* QXmlStreamAttribute_new4(QXmlStreamAttribute* param1) {
return new QXmlStreamAttribute(*param1);
}
bool QXmlStreamAttribute_IsDefault(const QXmlStreamAttribute* self) {
	return self->isDefault();
}
void QXmlStreamAttribute_Delete(QXmlStreamAttribute* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QXmlStreamAttribute*>( self );
	} else {
		delete self;
	}
}
QXmlStreamNamespaceDeclaration* QXmlStreamNamespaceDeclaration_new() {
return new QXmlStreamNamespaceDeclaration();
}
QXmlStreamNamespaceDeclaration* QXmlStreamNamespaceDeclaration_new2(struct miqt_string prefix, struct miqt_string namespaceUri) {
QString prefix_QString = QString::fromUtf8(prefix.data, prefix.len);
	QString namespaceUri_QString = QString::fromUtf8(namespaceUri.data, namespaceUri.len);
	return new QXmlStreamNamespaceDeclaration(prefix_QString, namespaceUri_QString);
}
QXmlStreamNamespaceDeclaration* QXmlStreamNamespaceDeclaration_new3(QXmlStreamNamespaceDeclaration* param1) {
return new QXmlStreamNamespaceDeclaration(*param1);
}
void QXmlStreamNamespaceDeclaration_Delete(QXmlStreamNamespaceDeclaration* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QXmlStreamNamespaceDeclaration*>( self );
	} else {
		delete self;
	}
}
QXmlStreamNotationDeclaration* QXmlStreamNotationDeclaration_new() {
return new QXmlStreamNotationDeclaration();
}
QXmlStreamNotationDeclaration* QXmlStreamNotationDeclaration_new2(QXmlStreamNotationDeclaration* param1) {
return new QXmlStreamNotationDeclaration(*param1);
}
void QXmlStreamNotationDeclaration_Delete(QXmlStreamNotationDeclaration* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QXmlStreamNotationDeclaration*>( self );
	} else {
		delete self;
	}
}
QXmlStreamEntityDeclaration* QXmlStreamEntityDeclaration_new() {
return new QXmlStreamEntityDeclaration();
}
QXmlStreamEntityDeclaration* QXmlStreamEntityDeclaration_new2(QXmlStreamEntityDeclaration* param1) {
return new QXmlStreamEntityDeclaration(*param1);
}
void QXmlStreamEntityDeclaration_Delete(QXmlStreamEntityDeclaration* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QXmlStreamEntityDeclaration*>( self );
	} else {
		delete self;
	}
}
QXmlStreamEntityResolver* QXmlStreamEntityResolver_new() {
return new QXmlStreamEntityResolver();
}
struct miqt_string QXmlStreamEntityResolver_ResolveEntity(QXmlStreamEntityResolver* self, struct miqt_string publicId, struct miqt_string systemId) {
	QString publicId_QString = QString::fromUtf8(publicId.data, publicId.len);
	QString systemId_QString = QString::fromUtf8(systemId.data, systemId.len);
	QString _ret = self->resolveEntity(publicId_QString, systemId_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QXmlStreamEntityResolver_ResolveUndeclaredEntity(QXmlStreamEntityResolver* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString _ret = self->resolveUndeclaredEntity(name_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QXmlStreamEntityResolver_Delete(QXmlStreamEntityResolver* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QXmlStreamEntityResolver*>( self );
	} else {
		delete self;
	}
}
QXmlStreamReader* QXmlStreamReader_new() {
return new QXmlStreamReader();
}
QXmlStreamReader* QXmlStreamReader_new2(QIODevice* device) {
return new QXmlStreamReader(device);
}
QXmlStreamReader* QXmlStreamReader_new3(QAnyStringView* data) {
return new QXmlStreamReader(*data);
}
void QXmlStreamReader_SetDevice(QXmlStreamReader* self, QIODevice* device) {
	self->setDevice(device);
}
QIODevice* QXmlStreamReader_Device(const QXmlStreamReader* self) {
	return self->device();
}
void QXmlStreamReader_AddData(QXmlStreamReader* self, QAnyStringView* data) {
	self->addData(*data);
}
void QXmlStreamReader_Clear(QXmlStreamReader* self) {
	self->clear();
}
bool QXmlStreamReader_AtEnd(const QXmlStreamReader* self) {
	return self->atEnd();
}
TokenType QXmlStreamReader_ReadNext(QXmlStreamReader* self) {
	return self->readNext();
}
bool QXmlStreamReader_ReadNextStartElement(QXmlStreamReader* self) {
	return self->readNextStartElement();
}
void QXmlStreamReader_SkipCurrentElement(QXmlStreamReader* self) {
	self->skipCurrentElement();
}
TokenType QXmlStreamReader_TokenType(const QXmlStreamReader* self) {
	return self->tokenType();
}
struct miqt_string QXmlStreamReader_TokenString(const QXmlStreamReader* self) {
	QString _ret = self->tokenString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QXmlStreamReader_SetNamespaceProcessing(QXmlStreamReader* self, bool namespaceProcessing) {
	self->setNamespaceProcessing(namespaceProcessing);
}
bool QXmlStreamReader_NamespaceProcessing(const QXmlStreamReader* self) {
	return self->namespaceProcessing();
}
bool QXmlStreamReader_IsStartDocument(const QXmlStreamReader* self) {
	return self->isStartDocument();
}
bool QXmlStreamReader_IsEndDocument(const QXmlStreamReader* self) {
	return self->isEndDocument();
}
bool QXmlStreamReader_IsStartElement(const QXmlStreamReader* self) {
	return self->isStartElement();
}
bool QXmlStreamReader_IsEndElement(const QXmlStreamReader* self) {
	return self->isEndElement();
}
bool QXmlStreamReader_IsCharacters(const QXmlStreamReader* self) {
	return self->isCharacters();
}
bool QXmlStreamReader_IsWhitespace(const QXmlStreamReader* self) {
	return self->isWhitespace();
}
bool QXmlStreamReader_IsCDATA(const QXmlStreamReader* self) {
	return self->isCDATA();
}
bool QXmlStreamReader_IsComment(const QXmlStreamReader* self) {
	return self->isComment();
}
bool QXmlStreamReader_IsDTD(const QXmlStreamReader* self) {
	return self->isDTD();
}
bool QXmlStreamReader_IsEntityReference(const QXmlStreamReader* self) {
	return self->isEntityReference();
}
bool QXmlStreamReader_IsProcessingInstruction(const QXmlStreamReader* self) {
	return self->isProcessingInstruction();
}
bool QXmlStreamReader_IsStandaloneDocument(const QXmlStreamReader* self) {
	return self->isStandaloneDocument();
}
bool QXmlStreamReader_HasStandaloneDeclaration(const QXmlStreamReader* self) {
	return self->hasStandaloneDeclaration();
}
long long QXmlStreamReader_LineNumber(const QXmlStreamReader* self) {
	qint64 _ret = self->lineNumber();
	return static_cast<long long>(_ret);
}
long long QXmlStreamReader_ColumnNumber(const QXmlStreamReader* self) {
	qint64 _ret = self->columnNumber();
	return static_cast<long long>(_ret);
}
long long QXmlStreamReader_CharacterOffset(const QXmlStreamReader* self) {
	qint64 _ret = self->characterOffset();
	return static_cast<long long>(_ret);
}
struct miqt_string QXmlStreamReader_ReadElementText(QXmlStreamReader* self) {
	QString _ret = self->readElementText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_array /* of QXmlStreamNamespaceDeclaration* */  QXmlStreamReader_NamespaceDeclarations(const QXmlStreamReader* self) {
	QXmlStreamNamespaceDeclarations _ret = self->namespaceDeclarations();
	// Convert QList<> from C++ memory to manually-managed C memory
	QXmlStreamNamespaceDeclaration** _arr = static_cast<QXmlStreamNamespaceDeclaration**>(malloc(sizeof(QXmlStreamNamespaceDeclaration*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QXmlStreamNamespaceDeclaration(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QXmlStreamReader_AddExtraNamespaceDeclaration(QXmlStreamReader* self, QXmlStreamNamespaceDeclaration* extraNamespaceDeclaraction) {
	self->addExtraNamespaceDeclaration(*extraNamespaceDeclaraction);
}
void QXmlStreamReader_AddExtraNamespaceDeclarations(QXmlStreamReader* self, struct miqt_array /* of QXmlStreamNamespaceDeclaration* */  extraNamespaceDeclaractions) {
	QXmlStreamNamespaceDeclarations extraNamespaceDeclaractions_QList;
	extraNamespaceDeclaractions_QList.reserve(extraNamespaceDeclaractions.len);
	QXmlStreamNamespaceDeclaration** extraNamespaceDeclaractions_arr = static_cast<QXmlStreamNamespaceDeclaration**>(extraNamespaceDeclaractions.data);
	for(size_t i = 0; i < extraNamespaceDeclaractions.len; ++i) {
		extraNamespaceDeclaractions_QList.push_back(*(extraNamespaceDeclaractions_arr[i]));
	}
	self->addExtraNamespaceDeclarations(extraNamespaceDeclaractions_QList);
}
struct miqt_array /* of QXmlStreamNotationDeclaration* */  QXmlStreamReader_NotationDeclarations(const QXmlStreamReader* self) {
	QXmlStreamNotationDeclarations _ret = self->notationDeclarations();
	// Convert QList<> from C++ memory to manually-managed C memory
	QXmlStreamNotationDeclaration** _arr = static_cast<QXmlStreamNotationDeclaration**>(malloc(sizeof(QXmlStreamNotationDeclaration*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QXmlStreamNotationDeclaration(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_array /* of QXmlStreamEntityDeclaration* */  QXmlStreamReader_EntityDeclarations(const QXmlStreamReader* self) {
	QXmlStreamEntityDeclarations _ret = self->entityDeclarations();
	// Convert QList<> from C++ memory to manually-managed C memory
	QXmlStreamEntityDeclaration** _arr = static_cast<QXmlStreamEntityDeclaration**>(malloc(sizeof(QXmlStreamEntityDeclaration*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QXmlStreamEntityDeclaration(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
int QXmlStreamReader_EntityExpansionLimit(const QXmlStreamReader* self) {
	return self->entityExpansionLimit();
}
void QXmlStreamReader_SetEntityExpansionLimit(QXmlStreamReader* self, int limit) {
	self->setEntityExpansionLimit(static_cast<int>(limit));
}
void QXmlStreamReader_RaiseError(QXmlStreamReader* self) {
	self->raiseError();
}
struct miqt_string QXmlStreamReader_ErrorString(const QXmlStreamReader* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
Error QXmlStreamReader_Error(const QXmlStreamReader* self) {
	return self->error();
}
bool QXmlStreamReader_HasError(const QXmlStreamReader* self) {
	return self->hasError();
}
void QXmlStreamReader_SetEntityResolver(QXmlStreamReader* self, QXmlStreamEntityResolver* resolver) {
	self->setEntityResolver(resolver);
}
QXmlStreamEntityResolver* QXmlStreamReader_EntityResolver(const QXmlStreamReader* self) {
	return self->entityResolver();
}
struct miqt_string QXmlStreamReader_ReadElementText1(QXmlStreamReader* self, ReadElementTextBehaviour behaviour) {
	QString _ret = self->readElementText(behaviour);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QXmlStreamReader_RaiseError1(QXmlStreamReader* self, struct miqt_string message) {
	QString message_QString = QString::fromUtf8(message.data, message.len);
	self->raiseError(message_QString);
}
void QXmlStreamReader_Delete(QXmlStreamReader* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QXmlStreamReader*>( self );
	} else {
		delete self;
	}
}
QXmlStreamWriter* QXmlStreamWriter_new() {
return new QXmlStreamWriter();
}
QXmlStreamWriter* QXmlStreamWriter_new2(QIODevice* device) {
return new QXmlStreamWriter(device);
}
void QXmlStreamWriter_SetDevice(QXmlStreamWriter* self, QIODevice* device) {
	self->setDevice(device);
}
QIODevice* QXmlStreamWriter_Device(const QXmlStreamWriter* self) {
	return self->device();
}
void QXmlStreamWriter_SetAutoFormatting(QXmlStreamWriter* self, bool autoFormatting) {
	self->setAutoFormatting(autoFormatting);
}
bool QXmlStreamWriter_AutoFormatting(const QXmlStreamWriter* self) {
	return self->autoFormatting();
}
void QXmlStreamWriter_SetAutoFormattingIndent(QXmlStreamWriter* self, int spacesOrTabs) {
	self->setAutoFormattingIndent(static_cast<int>(spacesOrTabs));
}
int QXmlStreamWriter_AutoFormattingIndent(const QXmlStreamWriter* self) {
	return self->autoFormattingIndent();
}
void QXmlStreamWriter_WriteAttribute(QXmlStreamWriter* self, QAnyStringView* qualifiedName, QAnyStringView* value) {
	self->writeAttribute(*qualifiedName, *value);
}
void QXmlStreamWriter_WriteAttribute2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* name, QAnyStringView* value) {
	self->writeAttribute(*namespaceUri, *name, *value);
}
void QXmlStreamWriter_WriteAttributeWithAttribute(QXmlStreamWriter* self, QXmlStreamAttribute* attribute) {
	self->writeAttribute(*attribute);
}
void QXmlStreamWriter_WriteCDATA(QXmlStreamWriter* self, QAnyStringView* text) {
	self->writeCDATA(*text);
}
void QXmlStreamWriter_WriteCharacters(QXmlStreamWriter* self, QAnyStringView* text) {
	self->writeCharacters(*text);
}
void QXmlStreamWriter_WriteComment(QXmlStreamWriter* self, QAnyStringView* text) {
	self->writeComment(*text);
}
void QXmlStreamWriter_WriteDTD(QXmlStreamWriter* self, QAnyStringView* dtd) {
	self->writeDTD(*dtd);
}
void QXmlStreamWriter_WriteEmptyElement(QXmlStreamWriter* self, QAnyStringView* qualifiedName) {
	self->writeEmptyElement(*qualifiedName);
}
void QXmlStreamWriter_WriteEmptyElement2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* name) {
	self->writeEmptyElement(*namespaceUri, *name);
}
void QXmlStreamWriter_WriteTextElement(QXmlStreamWriter* self, QAnyStringView* qualifiedName, QAnyStringView* text) {
	self->writeTextElement(*qualifiedName, *text);
}
void QXmlStreamWriter_WriteTextElement2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* name, QAnyStringView* text) {
	self->writeTextElement(*namespaceUri, *name, *text);
}
void QXmlStreamWriter_WriteEndDocument(QXmlStreamWriter* self) {
	self->writeEndDocument();
}
void QXmlStreamWriter_WriteEndElement(QXmlStreamWriter* self) {
	self->writeEndElement();
}
void QXmlStreamWriter_WriteEntityReference(QXmlStreamWriter* self, QAnyStringView* name) {
	self->writeEntityReference(*name);
}
void QXmlStreamWriter_WriteNamespace(QXmlStreamWriter* self, QAnyStringView* namespaceUri) {
	self->writeNamespace(*namespaceUri);
}
void QXmlStreamWriter_WriteDefaultNamespace(QXmlStreamWriter* self, QAnyStringView* namespaceUri) {
	self->writeDefaultNamespace(*namespaceUri);
}
void QXmlStreamWriter_WriteProcessingInstruction(QXmlStreamWriter* self, QAnyStringView* target) {
	self->writeProcessingInstruction(*target);
}
void QXmlStreamWriter_WriteStartDocument(QXmlStreamWriter* self) {
	self->writeStartDocument();
}
void QXmlStreamWriter_WriteStartDocumentWithVersion(QXmlStreamWriter* self, QAnyStringView* version) {
	self->writeStartDocument(*version);
}
void QXmlStreamWriter_WriteStartDocument2(QXmlStreamWriter* self, QAnyStringView* version, bool standalone) {
	self->writeStartDocument(*version, standalone);
}
void QXmlStreamWriter_WriteStartElement(QXmlStreamWriter* self, QAnyStringView* qualifiedName) {
	self->writeStartElement(*qualifiedName);
}
void QXmlStreamWriter_WriteStartElement2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* name) {
	self->writeStartElement(*namespaceUri, *name);
}
void QXmlStreamWriter_WriteCurrentToken(QXmlStreamWriter* self, QXmlStreamReader* reader) {
	self->writeCurrentToken(*reader);
}
bool QXmlStreamWriter_HasError(const QXmlStreamWriter* self) {
	return self->hasError();
}
void QXmlStreamWriter_WriteNamespace2(QXmlStreamWriter* self, QAnyStringView* namespaceUri, QAnyStringView* prefix) {
	self->writeNamespace(*namespaceUri, *prefix);
}
void QXmlStreamWriter_WriteProcessingInstruction2(QXmlStreamWriter* self, QAnyStringView* target, QAnyStringView* data) {
	self->writeProcessingInstruction(*target, *data);
}
void QXmlStreamWriter_Delete(QXmlStreamWriter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QXmlStreamWriter*>( self );
	} else {
		delete self;
	}
}
