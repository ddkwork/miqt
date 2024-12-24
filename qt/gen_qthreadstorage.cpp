// +build ignore

#include <QThreadStorageData>
#include <qthreadstorage.h>
#include "gen_qthreadstorage.h"

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

void QThreadStorageData_Delete(QThreadStorageData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QThreadStorageData*>( self );
	} else {
		delete self;
	}
}

