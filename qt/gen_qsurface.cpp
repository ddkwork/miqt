// +build ignore

#include <QSize>
#include <QSurface>
#include <QSurfaceFormat>
#include <qsurface.h>
#include "gen_qsurface.h"
SurfaceClass QSurface_SurfaceClass(const QSurface* self) {
	return self->surfaceClass();
}
QSurfaceFormat* QSurface_Format(const QSurface* self) {
	return new QSurfaceFormat(self->format());
}
SurfaceType QSurface_SurfaceType(const QSurface* self) {
	return self->surfaceType();
}
bool QSurface_SupportsOpenGL(const QSurface* self) {
	return self->supportsOpenGL();
}
QSize* QSurface_Size(const QSurface* self) {
	return new QSize(self->size());
}
void QSurface_Delete(QSurface* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QSurface*>( self );
	} else {
		delete self;
	}
}
