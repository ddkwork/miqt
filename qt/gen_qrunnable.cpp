// +build ignore

#include <QGenericRunnable>
#include <QRunnable>
#include <qrunnable.h>
#include "gen_qrunnable.h"

QRunnable* QRunnable_new() {
	return new QRunnable();
}

void QRunnable_Run(QRunnable* self) {
	self->run();
}

bool QRunnable_AutoDelete(const QRunnable* self) {
	return self->autoDelete();
}

void QRunnable_SetAutoDelete(QRunnable* self, bool autoDelete) {
	self->setAutoDelete(autoDelete);
}

void QRunnable_Delete(QRunnable* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QRunnable*>( self );
	} else {
		delete self;
	}
}

void QGenericRunnable_virtbase(QGenericRunnable* src, QRunnable** outptr_QRunnable) {
	*outptr_QRunnable = static_cast<QRunnable*>(src);
}

void QGenericRunnable_Run(QGenericRunnable* self) {
	self->run();
}

void QGenericRunnable_Delete(QGenericRunnable* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QGenericRunnable*>( self );
	} else {
		delete self;
	}
}

