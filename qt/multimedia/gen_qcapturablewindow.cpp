// +build ignore

#include <QCapturableWindow>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qcapturablewindow.h>
#include "gen_qcapturablewindow.h"

QCapturableWindow* QCapturableWindow_new() {
	return new QCapturableWindow();
}

QCapturableWindow* QCapturableWindow_new2(QCapturableWindow* other) {
	return new QCapturableWindow(*other);
}

void QCapturableWindow_OperatorAssign(QCapturableWindow* self, QCapturableWindow* other) {
	self->operator=(*other);
}

void QCapturableWindow_Swap(QCapturableWindow* self, QCapturableWindow* other) {
	self->swap(*other);
}

bool QCapturableWindow_IsValid(const QCapturableWindow* self) {
	return self->isValid();
}

struct miqt_string QCapturableWindow_Description(const QCapturableWindow* self) {
	QString _ret = self->description();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QCapturableWindow_Delete(QCapturableWindow* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCapturableWindow*>( self );
	} else {
		delete self;
	}
}

