// +build ignore

#include <QByteArray>
#include <QList>
#include <QMetaObject>
#include <QMetaType>
#include <QMimeData>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <QVariant>
#include <qmimedata.h>
#include "gen_qmimedata.h"
class MiqtVirtualQMimeData : public virtual QMimeData {
public:
MiqtVirtualQMimeData(): QMimeData() {};

virtual ~MiqtVirtualQMimeData() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QMimeData::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QMimeData_MetaObject(const_cast<MiqtVirtualQMimeData*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QMimeData::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QMimeData::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QMimeData_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QMimeData::qt_metacast(param1);

	}
};
QMimeData* QMimeData_new() {
return new MiqtVirtualQMimeData();
}
void QMimeData_virtbase(QMimeData* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QMimeData_MetaObject(const QMimeData* self) {
	return (QMetaObject*) self->metaObject();
}
void* QMimeData_Metacast(QMimeData* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QMimeData_Tr(const char* s) {
	QString _ret = QMimeData::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_array /* of QUrl* */  QMimeData_Urls(const QMimeData* self) {
	QList<QUrl> _ret = self->urls();
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QMimeData_SetUrls(QMimeData* self, struct miqt_array /* of QUrl* */  urls) {
	QList<QUrl> urls_QList;
	urls_QList.reserve(urls.len);
	QUrl** urls_arr = static_cast<QUrl**>(urls.data);
	for(size_t i = 0; i < urls.len; ++i) {
		urls_QList.push_back(*(urls_arr[i]));
	}
	self->setUrls(urls_QList);
}
bool QMimeData_HasUrls(const QMimeData* self) {
	return self->hasUrls();
}
struct miqt_string QMimeData_Text(const QMimeData* self) {
	QString _ret = self->text();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMimeData_SetText(QMimeData* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setText(text_QString);
}
bool QMimeData_HasText(const QMimeData* self) {
	return self->hasText();
}
struct miqt_string QMimeData_Html(const QMimeData* self) {
	QString _ret = self->html();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMimeData_SetHtml(QMimeData* self, struct miqt_string html) {
	QString html_QString = QString::fromUtf8(html.data, html.len);
	self->setHtml(html_QString);
}
bool QMimeData_HasHtml(const QMimeData* self) {
	return self->hasHtml();
}
QVariant* QMimeData_ImageData(const QMimeData* self) {
	return new QVariant(self->imageData());
}
void QMimeData_SetImageData(QMimeData* self, QVariant* image) {
	self->setImageData(*image);
}
bool QMimeData_HasImage(const QMimeData* self) {
	return self->hasImage();
}
QVariant* QMimeData_ColorData(const QMimeData* self) {
	return new QVariant(self->colorData());
}
void QMimeData_SetColorData(QMimeData* self, QVariant* color) {
	self->setColorData(*color);
}
bool QMimeData_HasColor(const QMimeData* self) {
	return self->hasColor();
}
struct miqt_string QMimeData_Data(const QMimeData* self, struct miqt_string mimetype) {
	QString mimetype_QString = QString::fromUtf8(mimetype.data, mimetype.len);
	QByteArray _qb = self->data(mimetype_QString);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
void QMimeData_SetData(QMimeData* self, struct miqt_string mimetype, struct miqt_string data) {
	QString mimetype_QString = QString::fromUtf8(mimetype.data, mimetype.len);
	QByteArray data_QByteArray(data.data, data.len);
	self->setData(mimetype_QString, data_QByteArray);
}
void QMimeData_RemoveFormat(QMimeData* self, struct miqt_string mimetype) {
	QString mimetype_QString = QString::fromUtf8(mimetype.data, mimetype.len);
	self->removeFormat(mimetype_QString);
}
bool QMimeData_HasFormat(const QMimeData* self, struct miqt_string mimetype) {
	QString mimetype_QString = QString::fromUtf8(mimetype.data, mimetype.len);
	return self->hasFormat(mimetype_QString);
}
struct miqt_array /* of struct miqt_string */  QMimeData_Formats(const QMimeData* self) {
	QStringList _ret = self->formats();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QMimeData_Clear(QMimeData* self) {
	self->clear();
}
struct miqt_string QMimeData_Tr2(const char* s, const char* c) {
	QString _ret = QMimeData::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMimeData_Tr3(const char* s, const char* c, int n) {
	QString _ret = QMimeData::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMimeData_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMimeData*>( (QMimeData*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QMimeData_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQMimeData*)(self) )->virtualbase_MetaObject();
}
void QMimeData_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMimeData*>( (QMimeData*)(self) )->handle__Metacast = slot;
}
void* QMimeData_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQMimeData*)(self) )->virtualbase_Metacast(param1);
}
void QMimeData_Delete(QMimeData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQMimeData*>( self );
	} else {
		delete self;
	}
}
