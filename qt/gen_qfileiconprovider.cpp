#include <QFileIconProvider>
#include <QFileInfo>
#include <QIcon>
#include <QString>
#include <QByteArray>
#include <cstring>
#include "qfileiconprovider.h"

#include "gen_qfileiconprovider.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QFileIconProvider* QFileIconProvider_new() {
	return new QFileIconProvider();
}

QIcon* QFileIconProvider_Icon(const QFileIconProvider* self, uintptr_t typeVal) {
	QIcon ret = self->icon(static_cast<QFileIconProvider::IconType>(typeVal));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QIcon*>(new QIcon(ret));
}

QIcon* QFileIconProvider_IconWithInfo(const QFileIconProvider* self, QFileInfo* info) {
	QIcon ret = self->icon(*info);
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QIcon*>(new QIcon(ret));
}

void QFileIconProvider_Type(const QFileIconProvider* self, QFileInfo* info, char** _out, int* _out_Strlen) {
	QString ret = self->type(*info);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QFileIconProvider_SetOptions(QFileIconProvider* self, int options) {
	self->setOptions(static_cast<QFileIconProvider::Options>(options));
}

int QFileIconProvider_Options(const QFileIconProvider* self) {
	QFileIconProvider::Options ret = self->options();
	return static_cast<int>(ret);
}

void QFileIconProvider_Delete(QFileIconProvider* self) {
	delete self;
}

