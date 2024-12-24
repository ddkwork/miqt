// +build ignore

#include <QSharedData>
#include <qshareddata.h>
#include "gen_qshareddata.h"

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

QSharedData* QSharedData_new() {
	return new QSharedData();
}

QSharedData* QSharedData_new2(QSharedData* param1) {
	return new QSharedData(*param1);
}

void QSharedData_Delete(QSharedData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QSharedData*>( self );
	} else {
		delete self;
	}
}

QAdoptSharedDataTag* QAdoptSharedDataTag_new() {
	return new QAdoptSharedDataTag();
}

void QAdoptSharedDataTag_Delete(QAdoptSharedDataTag* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAdoptSharedDataTag*>( self );
	} else {
		delete self;
	}
}

