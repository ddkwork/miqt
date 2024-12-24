// +build ignore

#include <QPainter>
#include <QPainterStateGuard>
#include <qpainterstateguard.h>
#include "gen_qpainterstateguard.h"
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
