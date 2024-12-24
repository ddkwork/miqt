// +build ignore

#include <QByteArray>
#include <QIODevice>
#include <QMetaObject>
#include <QObject>
#include <QPagedPaintDevice>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPdfOutputIntent>
#include <QPdfWriter>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUuid>
#include <qpdfwriter.h>
#include "gen_qpdfwriter.h"
class MiqtVirtualQPdfWriter : public virtual QPdfWriter {
public:
MiqtVirtualQPdfWriter(const QString& filename): QPdfWriter(filename) {};
MiqtVirtualQPdfWriter(QIODevice* device): QPdfWriter(device) {};

virtual ~MiqtVirtualQPdfWriter() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QPdfWriter::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QPdfWriter_MetaObject(const_cast<MiqtVirtualQPdfWriter*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPdfWriter::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QPdfWriter::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPdfWriter_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPdfWriter::qt_metacast(param1);

	}
};
QPdfWriter* QPdfWriter_new(struct miqt_string filename) {
QString filename_QString = QString::fromUtf8(filename.data, filename.len);
	return new MiqtVirtualQPdfWriter(filename_QString);
}
QPdfWriter* QPdfWriter_new2(QIODevice* device) {
return new MiqtVirtualQPdfWriter(device);
}
void QPdfWriter_virtbase(QPdfWriter* src
, QObject** outptr_QObject
, QPagedPaintDevice** outptr_QPagedPaintDevice
) {
*outptr_QObject = static_cast<QObject*>(src);
*outptr_QPagedPaintDevice = static_cast<QPagedPaintDevice*>(src);
}
QMetaObject* QPdfWriter_MetaObject(const QPdfWriter* self) {
	return (QMetaObject*) self->metaObject();
}
void* QPdfWriter_Metacast(QPdfWriter* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QPdfWriter_Tr(const char* s) {
	QString _ret = QPdfWriter::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPdfWriter_SetPdfVersion(QPdfWriter* self, PdfVersion version) {
	self->setPdfVersion(version);
}
PdfVersion QPdfWriter_PdfVersion(const QPdfWriter* self) {
	return self->pdfVersion();
}
struct miqt_string QPdfWriter_Title(const QPdfWriter* self) {
	QString _ret = self->title();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPdfWriter_SetTitle(QPdfWriter* self, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	self->setTitle(title_QString);
}
struct miqt_string QPdfWriter_Creator(const QPdfWriter* self) {
	QString _ret = self->creator();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPdfWriter_SetCreator(QPdfWriter* self, struct miqt_string creator) {
	QString creator_QString = QString::fromUtf8(creator.data, creator.len);
	self->setCreator(creator_QString);
}
QUuid* QPdfWriter_DocumentId(const QPdfWriter* self) {
	return new QUuid(self->documentId());
}
void QPdfWriter_SetDocumentId(QPdfWriter* self, QUuid* documentId) {
	self->setDocumentId(*documentId);
}
struct miqt_string QPdfWriter_Author(const QPdfWriter* self) {
	QString _ret = self->author();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPdfWriter_SetAuthor(QPdfWriter* self, struct miqt_string author) {
	QString author_QString = QString::fromUtf8(author.data, author.len);
	self->setAuthor(author_QString);
}
bool QPdfWriter_NewPage(QPdfWriter* self) {
	return self->newPage();
}
void QPdfWriter_SetResolution(QPdfWriter* self, int resolution) {
	self->setResolution(static_cast<int>(resolution));
}
int QPdfWriter_Resolution(const QPdfWriter* self) {
	return self->resolution();
}
void QPdfWriter_SetDocumentXmpMetadata(QPdfWriter* self, struct miqt_string xmpMetadata) {
	QByteArray xmpMetadata_QByteArray(xmpMetadata.data, xmpMetadata.len);
	self->setDocumentXmpMetadata(xmpMetadata_QByteArray);
}
struct miqt_string QPdfWriter_DocumentXmpMetadata(const QPdfWriter* self) {
	QByteArray _qb = self->documentXmpMetadata();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
void QPdfWriter_AddFileAttachment(QPdfWriter* self, struct miqt_string fileName, struct miqt_string data) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	QByteArray data_QByteArray(data.data, data.len);
	self->addFileAttachment(fileName_QString, data_QByteArray);
}
ColorModel QPdfWriter_ColorModel(const QPdfWriter* self) {
	return self->colorModel();
}
void QPdfWriter_SetColorModel(QPdfWriter* self, ColorModel model) {
	self->setColorModel(model);
}
QPdfOutputIntent* QPdfWriter_OutputIntent(const QPdfWriter* self) {
	return new QPdfOutputIntent(self->outputIntent());
}
void QPdfWriter_SetOutputIntent(QPdfWriter* self, QPdfOutputIntent* intent) {
	self->setOutputIntent(*intent);
}
struct miqt_string QPdfWriter_Tr2(const char* s, const char* c) {
	QString _ret = QPdfWriter::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QPdfWriter_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPdfWriter::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPdfWriter_AddFileAttachment3(QPdfWriter* self, struct miqt_string fileName, struct miqt_string data, struct miqt_string mimeType) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	QByteArray data_QByteArray(data.data, data.len);
	QString mimeType_QString = QString::fromUtf8(mimeType.data, mimeType.len);
	self->addFileAttachment(fileName_QString, data_QByteArray, mimeType_QString);
}
void QPdfWriter_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPdfWriter*>( (QPdfWriter*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QPdfWriter_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPdfWriter*)(self) )->virtualbase_MetaObject();
}
void QPdfWriter_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPdfWriter*>( (QPdfWriter*)(self) )->handle__Metacast = slot;
}
void* QPdfWriter_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPdfWriter*)(self) )->virtualbase_Metacast(param1);
}
void QPdfWriter_Delete(QPdfWriter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPdfWriter*>( self );
	} else {
		delete self;
	}
}
