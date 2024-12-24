// +build ignore

#include <QScrollerProperties>
#include <QVariant>
#include <qscrollerproperties.h>
#include "gen_qscrollerproperties.h"

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
	return (*self == *sp);
}

bool QScrollerProperties_OperatorNotEqual(const QScrollerProperties* self, QScrollerProperties* sp) {
	return (*self != *sp);
}

void QScrollerProperties_SetDefaultScrollerProperties(QScrollerProperties* sp) {
	QScrollerProperties::setDefaultScrollerProperties(*sp);
}

void QScrollerProperties_UnsetDefaultScrollerProperties() {
	QScrollerProperties::unsetDefaultScrollerProperties();
}

QVariant* QScrollerProperties_ScrollMetric(const QScrollerProperties* self, ScrollMetric metric) {
	return new QVariant(self->scrollMetric(metric));
}

void QScrollerProperties_SetScrollMetric(QScrollerProperties* self, ScrollMetric metric, QVariant* value) {
	self->setScrollMetric(metric, *value);
}

void QScrollerProperties_Delete(QScrollerProperties* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QScrollerProperties*>( self );
	} else {
		delete self;
	}
}

