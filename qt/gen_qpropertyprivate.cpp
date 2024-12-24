// +build ignore

#include <QPropertyProxyBindingData>
#include <QUntypedPropertyData>
#include <qpropertyprivate.h>
#include "gen_qpropertyprivate.h"

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

void QUntypedPropertyData_Delete(QUntypedPropertyData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QUntypedPropertyData*>( self );
	} else {
		delete self;
	}
}

void QPropertyProxyBindingData_Delete(QPropertyProxyBindingData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPropertyProxyBindingData*>( self );
	} else {
		delete self;
	}
}

