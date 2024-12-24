// +build ignore

#include <QAtomicInt>
#include <qatomic.h>
#include "gen_qatomic.h"

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

QAtomicInt* QAtomicInt_new() {
	return new QAtomicInt();
}

QAtomicInt* QAtomicInt_new2(QAtomicInt* param1) {
	return new QAtomicInt(*param1);
}

QAtomicInt* QAtomicInt_new3(int value) {
	return new QAtomicInt(static_cast<int>(value));
}

void QAtomicInt_OperatorAssign(QAtomicInt* self, QAtomicInt* param1) {
	self->operator=(*param1);
}

void QAtomicInt_Delete(QAtomicInt* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAtomicInt*>( self );
	} else {
		delete self;
	}
}

