// +build ignore

#include <QPrintEngine>
#include <QVariant>
#include <qprintengine.h>
#include "gen_qprintengine.h"

#ifndef _Bool
#define _Bool bool
#endif

void QPrintEngine_SetProperty(QPrintEngine* self, PrintEnginePropertyKey key, QVariant* value) {
	self->setProperty(key, *value);
}

QVariant* QPrintEngine_Property(const QPrintEngine* self, PrintEnginePropertyKey key) {
	return new QVariant(self->property(key));
}

bool QPrintEngine_NewPage(QPrintEngine* self) {
	return self->newPage();
}

bool QPrintEngine_Abort(QPrintEngine* self) {
	return self->abort();
}

int QPrintEngine_Metric(const QPrintEngine* self, int param1) {
	return self->metric(static_cast<QPaintDevice::PaintDeviceMetric>(param1));
}

int QPrintEngine_PrinterState(const QPrintEngine* self) {
	QPrinter::PrinterState _ret = self->printerState();
	return static_cast<int>(_ret);
}

void QPrintEngine_OperatorAssign(QPrintEngine* self, QPrintEngine* param1) {
	self->operator=(*param1);
}

void QPrintEngine_Delete(QPrintEngine* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPrintEngine*>( self );
	} else {
		delete self;
	}
}

