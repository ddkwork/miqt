// +build ignore

#include <QBindingStatus>
#include <QBindingStorage>
#include <QUntypedPropertyData>
#include <qbindingstorage.h>
#include "gen_qbindingstorage.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

void QBindingStatus_Delete(QBindingStatus* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QBindingStatus*>( self );
	} else {
		delete self;
	}
}

QBindingStorage* QBindingStorage_new() {
	return new QBindingStorage();
}

bool QBindingStorage_IsEmpty(QBindingStorage* self) {
	return self->isEmpty();
}

bool QBindingStorage_IsValid(const QBindingStorage* self) {
	return self->isValid();
}

void QBindingStorage_RegisterDependency(const QBindingStorage* self, QUntypedPropertyData* data) {
	self->registerDependency(data);
}

void QBindingStorage_Delete(QBindingStorage* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QBindingStorage*>( self );
	} else {
		delete self;
	}
}

