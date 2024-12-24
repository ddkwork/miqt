// +build ignore

#include <QScopedPointerPodDeleter>
#include <qscopedpointer.h>
#include "gen_qscopedpointer.h"

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

void QScopedPointerPodDeleter_Cleanup(void* pointer) {
	QScopedPointerPodDeleter::cleanup(pointer);
}

void QScopedPointerPodDeleter_OperatorCall(const QScopedPointerPodDeleter* self, void* pointer) {
	self->operator()(pointer);
}

void QScopedPointerPodDeleter_Delete(QScopedPointerPodDeleter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QScopedPointerPodDeleter*>( self );
	} else {
		delete self;
	}
}

