#include <QScrollerProperties>
#include <QVariant>
#include "qscrollerproperties.h"

#include "gen_qscrollerproperties.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QScrollerProperties* QScrollerProperties_new() {
	return new QScrollerProperties();
}

QScrollerProperties* QScrollerProperties_new2(QScrollerProperties* sp) {
	return new QScrollerProperties(*sp);
}

void QScrollerProperties_OperatorAssign(QScrollerProperties* self, QScrollerProperties* sp) {
	self->operator=(*sp);
}

bool QScrollerProperties_OperatorEqual(const QScrollerProperties* self, QScrollerProperties* sp) {
	return self->operator==(*sp);
}

bool QScrollerProperties_OperatorNotEqual(const QScrollerProperties* self, QScrollerProperties* sp) {
	return self->operator!=(*sp);
}

void QScrollerProperties_SetDefaultScrollerProperties(QScrollerProperties* sp) {
	QScrollerProperties::setDefaultScrollerProperties(*sp);
}

void QScrollerProperties_UnsetDefaultScrollerProperties() {
	QScrollerProperties::unsetDefaultScrollerProperties();
}

QVariant* QScrollerProperties_ScrollMetric(const QScrollerProperties* self, uintptr_t metric) {
	QVariant ret = self->scrollMetric(static_cast<QScrollerProperties::ScrollMetric>(metric));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QVariant*>(new QVariant(ret));
}

void QScrollerProperties_SetScrollMetric(QScrollerProperties* self, uintptr_t metric, QVariant* value) {
	self->setScrollMetric(static_cast<QScrollerProperties::ScrollMetric>(metric), *value);
}

void QScrollerProperties_Delete(QScrollerProperties* self) {
	delete self;
}

