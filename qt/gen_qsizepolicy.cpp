// +build ignore

#include <QSizePolicy>
#include <qsizepolicy.h>
#include "gen_qsizepolicy.h"

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

QSizePolicy* QSizePolicy_new() {
	return new QSizePolicy();
}

QSizePolicy* QSizePolicy_new2(Policy horizontal, Policy vertical) {
	return new QSizePolicy(horizontal, vertical);
}

QSizePolicy* QSizePolicy_new3(QSizePolicy* param1) {
	return new QSizePolicy(*param1);
}

QSizePolicy* QSizePolicy_new4(Policy horizontal, Policy vertical, ControlType typeVal) {
	return new QSizePolicy(horizontal, vertical, typeVal);
}

Policy QSizePolicy_HorizontalPolicy(const QSizePolicy* self) {
	return self->horizontalPolicy();
}

Policy QSizePolicy_VerticalPolicy(const QSizePolicy* self) {
	return self->verticalPolicy();
}

ControlType QSizePolicy_ControlType(const QSizePolicy* self) {
	return self->controlType();
}

void QSizePolicy_SetHorizontalPolicy(QSizePolicy* self, Policy d) {
	self->setHorizontalPolicy(d);
}

void QSizePolicy_SetVerticalPolicy(QSizePolicy* self, Policy d) {
	self->setVerticalPolicy(d);
}

void QSizePolicy_SetControlType(QSizePolicy* self, ControlType typeVal) {
	self->setControlType(typeVal);
}

int QSizePolicy_ExpandingDirections(const QSizePolicy* self) {
	Qt::Orientations _ret = self->expandingDirections();
	return static_cast<int>(_ret);
}

void QSizePolicy_SetHeightForWidth(QSizePolicy* self, bool b) {
	self->setHeightForWidth(b);
}

bool QSizePolicy_HasHeightForWidth(const QSizePolicy* self) {
	return self->hasHeightForWidth();
}

void QSizePolicy_SetWidthForHeight(QSizePolicy* self, bool b) {
	self->setWidthForHeight(b);
}

bool QSizePolicy_HasWidthForHeight(const QSizePolicy* self) {
	return self->hasWidthForHeight();
}

bool QSizePolicy_OperatorEqual(const QSizePolicy* self, QSizePolicy* s) {
	return (*self == *s);
}

bool QSizePolicy_OperatorNotEqual(const QSizePolicy* self, QSizePolicy* s) {
	return (*self != *s);
}

int QSizePolicy_HorizontalStretch(const QSizePolicy* self) {
	return self->horizontalStretch();
}

int QSizePolicy_VerticalStretch(const QSizePolicy* self) {
	return self->verticalStretch();
}

void QSizePolicy_SetHorizontalStretch(QSizePolicy* self, int stretchFactor) {
	self->setHorizontalStretch(static_cast<int>(stretchFactor));
}

void QSizePolicy_SetVerticalStretch(QSizePolicy* self, int stretchFactor) {
	self->setVerticalStretch(static_cast<int>(stretchFactor));
}

bool QSizePolicy_RetainSizeWhenHidden(const QSizePolicy* self) {
	return self->retainSizeWhenHidden();
}

void QSizePolicy_SetRetainSizeWhenHidden(QSizePolicy* self, bool retainSize) {
	self->setRetainSizeWhenHidden(retainSize);
}

void QSizePolicy_Transpose(QSizePolicy* self) {
	self->transpose();
}

QSizePolicy* QSizePolicy_Transposed(const QSizePolicy* self) {
	return new QSizePolicy(self->transposed());
}

void QSizePolicy_Delete(QSizePolicy* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QSizePolicy*>( self );
	} else {
		delete self;
	}
}

