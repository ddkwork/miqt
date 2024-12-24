// +build ignore

#include <QFlag>
#include <QIncompatibleFlag>
#include <qflags.h>
#include "gen_qflags.h"

#ifndef _Bool
#define _Bool bool
#endif

QFlag* QFlag_new(int value) {
	return new QFlag(static_cast<int>(value));
}

QFlag* QFlag_new2(QFlag* param1) {
	return new QFlag(*param1);
}

void QFlag_Delete(QFlag* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFlag*>( self );
	} else {
		delete self;
	}
}

QIncompatibleFlag* QIncompatibleFlag_new(int i) {
	return new QIncompatibleFlag(static_cast<int>(i));
}

QIncompatibleFlag* QIncompatibleFlag_new2(QIncompatibleFlag* param1) {
	return new QIncompatibleFlag(*param1);
}

void QIncompatibleFlag_Delete(QIncompatibleFlag* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QIncompatibleFlag*>( self );
	} else {
		delete self;
	}
}

