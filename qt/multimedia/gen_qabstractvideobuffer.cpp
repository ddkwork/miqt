// +build ignore

#include <QAbstractVideoBuffer>
#define WORKAROUND_INNER_CLASS_DEFINITION_QAbstractVideoBuffer__MapData
#include <QVideoFrameFormat>
#include <qabstractvideobuffer.h>
#include "gen_qabstractvideobuffer.h"

MapData QAbstractVideoBuffer_Map(QAbstractVideoBuffer* self, int mode) {
	return self->map(static_cast<QVideoFrame::MapMode>(mode));
}

void QAbstractVideoBuffer_Unmap(QAbstractVideoBuffer* self) {
	self->unmap();
}

QVideoFrameFormat* QAbstractVideoBuffer_Format(const QAbstractVideoBuffer* self) {
	return new QVideoFrameFormat(self->format());
}

void QAbstractVideoBuffer_Delete(QAbstractVideoBuffer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractVideoBuffer*>( self );
	} else {
		delete self;
	}
}

void QAbstractVideoBuffer__MapData_Delete(QAbstractVideoBuffer__MapData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractVideoBuffer::MapData*>( self );
	} else {
		delete self;
	}
}

