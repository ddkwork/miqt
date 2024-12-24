// +build ignore

#include <qfloat16.h>
#include "gen_qfloat16.h"

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

qfloat16* qfloat16_new() {
	return new qfloat16();
}

qfloat16* qfloat16_new2(int param1) {
	return new qfloat16(static_cast<Qt::Initialization>(param1));
}

qfloat16* qfloat16_new3(float f) {
	return new qfloat16(static_cast<float>(f));
}

bool qfloat16_IsInf(const qfloat16* self) {
	return self->isInf();
}

bool qfloat16_IsNaN(const qfloat16* self) {
	return self->isNaN();
}

bool qfloat16_IsFinite(const qfloat16* self) {
	return self->isFinite();
}

int qfloat16_FpClassify(const qfloat16* self) {
	return self->fpClassify();
}

bool qfloat16_IsNormal(const qfloat16* self) {
	return self->isNormal();
}

void qfloat16_Delete(qfloat16* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<qfloat16*>( self );
	} else {
		delete self;
	}
}

