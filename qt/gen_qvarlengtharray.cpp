// +build ignore

#include <qvarlengtharray.h>
#include "gen_qvarlengtharray.h"

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

size_type QVLABaseBase_Capacity(const QVLABaseBase* self) {
	return self->capacity();
}

size_type QVLABaseBase_Size(const QVLABaseBase* self) {
	return self->size();
}

bool QVLABaseBase_Empty(const QVLABaseBase* self) {
	return self->empty();
}

