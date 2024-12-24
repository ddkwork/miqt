// +build ignore

#include <QAbstractFileIconProvider>
#include <QFileInfo>
#include <QIcon>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qabstractfileiconprovider.h>
#include "gen_qabstractfileiconprovider.h"
QAbstractFileIconProvider* QAbstractFileIconProvider_new() {
return new QAbstractFileIconProvider();
}
QIcon* QAbstractFileIconProvider_Icon(const QAbstractFileIconProvider* self, IconType param1) {
	return new QIcon(self->icon(param1));
}
QIcon* QAbstractFileIconProvider_IconWithQFileInfo(const QAbstractFileIconProvider* self, QFileInfo* param1) {
	return new QIcon(self->icon(*param1));
}
struct miqt_string QAbstractFileIconProvider_Type(const QAbstractFileIconProvider* self, QFileInfo* param1) {
	QString _ret = self->type(*param1);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAbstractFileIconProvider_SetOptions(QAbstractFileIconProvider* self, Options options) {
	self->setOptions(options);
}
Options QAbstractFileIconProvider_Options(const QAbstractFileIconProvider* self) {
	return self->options();
}
void QAbstractFileIconProvider_Delete(QAbstractFileIconProvider* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractFileIconProvider*>( self );
	} else {
		delete self;
	}
}
