// +build ignore

#include <QContiguousCacheData>
#include <qcontiguouscache.h>
#include "gen_qcontiguouscache.h"

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

QContiguousCacheData* QContiguousCacheData_AllocateData(ptrdiff_t size, ptrdiff_t alignment) {
	return QContiguousCacheData::allocateData((qsizetype)(size), (qsizetype)(alignment));
}

void QContiguousCacheData_FreeData(QContiguousCacheData* data) {
	QContiguousCacheData::freeData(data);
}

void QContiguousCacheData_Delete(QContiguousCacheData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QContiguousCacheData*>( self );
	} else {
		delete self;
	}
}

