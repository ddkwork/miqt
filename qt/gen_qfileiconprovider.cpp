// +build ignore

#include <QAbstractFileIconProvider>
#include <QFileIconProvider>
#include <QFileInfo>
#include <QIcon>
#include <qfileiconprovider.h>
#include "gen_qfileiconprovider.h"

QFileIconProvider* QFileIconProvider_new() {
	return new QFileIconProvider();
}

void QFileIconProvider_virtbase(QFileIconProvider* src, QAbstractFileIconProvider** outptr_QAbstractFileIconProvider) {
	*outptr_QAbstractFileIconProvider = static_cast<QAbstractFileIconProvider*>(src);
}

QIcon* QFileIconProvider_Icon(const QFileIconProvider* self, IconType typeVal) {
	return new QIcon(self->icon(typeVal));
}

QIcon* QFileIconProvider_IconWithInfo(const QFileIconProvider* self, QFileInfo* info) {
	return new QIcon(self->icon(*info));
}

void QFileIconProvider_Delete(QFileIconProvider* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFileIconProvider*>( self );
	} else {
		delete self;
	}
}

