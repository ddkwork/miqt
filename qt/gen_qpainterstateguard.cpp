// +build ignore

#include <QPainter>
#include <QPainterStateGuard>
#include <qpainterstateguard.h>
#include "gen_qpainterstateguard.h"

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

QPainterStateGuard* QPainterStateGuard_new(QPainter* painter) {
	return new QPainterStateGuard(painter);
}

QPainterStateGuard* QPainterStateGuard_new2(QPainter* painter, InitialState state) {
	return new QPainterStateGuard(painter, state);
}

void QPainterStateGuard_Save(QPainterStateGuard* self) {
	self->save();
}

void QPainterStateGuard_Restore(QPainterStateGuard* self) {
	self->restore();
}

void QPainterStateGuard_Delete(QPainterStateGuard* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPainterStateGuard*>( self );
	} else {
		delete self;
	}
}

