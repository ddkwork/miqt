// +build ignore

#include <QAtomicInt>
#include <qatomic.h>
#include "gen_qatomic.h"

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

