// +build ignore

#include <qvarlengtharray.h>
#include "gen_qvarlengtharray.h"

#ifndef _Bool
#define _Bool bool
#endif

size_type QVLABaseBase_Capacity(const QVLABaseBase* self) {
	return self->capacity();
}

size_type QVLABaseBase_Size(const QVLABaseBase* self) {
	return self->size();
}

bool QVLABaseBase_Empty(const QVLABaseBase* self) {
	return self->empty();
}

