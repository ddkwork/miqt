// +build ignore

#include <QTypeRevision>
#include <qtyperevision.h>
#include "gen_qtyperevision.h"

#ifndef _Bool
#define _Bool bool
#endif

QTypeRevision* QTypeRevision_new() {
	return new QTypeRevision();
}

QTypeRevision* QTypeRevision_new2(QTypeRevision* param1) {
	return new QTypeRevision(*param1);
}

QTypeRevision* QTypeRevision_Zero() {
	return new QTypeRevision(QTypeRevision::zero());
}

bool QTypeRevision_HasMajorVersion(const QTypeRevision* self) {
	return self->hasMajorVersion();
}

unsigned char QTypeRevision_MajorVersion(const QTypeRevision* self) {
	quint8 _ret = self->majorVersion();
	return static_cast<unsigned char>(_ret);
}

bool QTypeRevision_HasMinorVersion(const QTypeRevision* self) {
	return self->hasMinorVersion();
}

unsigned char QTypeRevision_MinorVersion(const QTypeRevision* self) {
	quint8 _ret = self->minorVersion();
	return static_cast<unsigned char>(_ret);
}

bool QTypeRevision_IsValid(const QTypeRevision* self) {
	return self->isValid();
}

void QTypeRevision_Delete(QTypeRevision* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTypeRevision*>( self );
	} else {
		delete self;
	}
}

