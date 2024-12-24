// +build ignore

#include <QOperatingSystemVersion>
#include <QOperatingSystemVersionBase>
#include <QOperatingSystemVersionUnexported>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVersionNumber>
#include <qoperatingsystemversion.h>
#include "gen_qoperatingsystemversion.h"

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

QOperatingSystemVersionBase* QOperatingSystemVersionBase_new(OSType osType, int vmajor) {
	return new QOperatingSystemVersionBase(osType, static_cast<int>(vmajor));
}

QOperatingSystemVersionBase* QOperatingSystemVersionBase_new2(QOperatingSystemVersionBase* param1) {
	return new QOperatingSystemVersionBase(*param1);
}

QOperatingSystemVersionBase* QOperatingSystemVersionBase_new3(OSType osType, int vmajor, int vminor) {
	return new QOperatingSystemVersionBase(osType, static_cast<int>(vmajor), static_cast<int>(vminor));
}

QOperatingSystemVersionBase* QOperatingSystemVersionBase_new4(OSType osType, int vmajor, int vminor, int vmicro) {
	return new QOperatingSystemVersionBase(osType, static_cast<int>(vmajor), static_cast<int>(vminor), static_cast<int>(vmicro));
}

QOperatingSystemVersionBase* QOperatingSystemVersionBase_Current() {
	return new QOperatingSystemVersionBase(QOperatingSystemVersionBase::current());
}

struct miqt_string QOperatingSystemVersionBase_Name(QOperatingSystemVersionBase* osversion) {
	QString _ret = QOperatingSystemVersionBase::name(*osversion);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

OSType QOperatingSystemVersionBase_CurrentType() {
	return QOperatingSystemVersionBase::currentType();
}

QVersionNumber* QOperatingSystemVersionBase_Version(const QOperatingSystemVersionBase* self) {
	return new QVersionNumber(self->version());
}

int QOperatingSystemVersionBase_MajorVersion(const QOperatingSystemVersionBase* self) {
	return self->majorVersion();
}

int QOperatingSystemVersionBase_MinorVersion(const QOperatingSystemVersionBase* self) {
	return self->minorVersion();
}

int QOperatingSystemVersionBase_MicroVersion(const QOperatingSystemVersionBase* self) {
	return self->microVersion();
}

int QOperatingSystemVersionBase_SegmentCount(const QOperatingSystemVersionBase* self) {
	return self->segmentCount();
}

OSType QOperatingSystemVersionBase_Type(const QOperatingSystemVersionBase* self) {
	return self->type();
}

struct miqt_string QOperatingSystemVersionBase_Name2(const QOperatingSystemVersionBase* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QOperatingSystemVersionBase_Delete(QOperatingSystemVersionBase* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QOperatingSystemVersionBase*>( self );
	} else {
		delete self;
	}
}

QOperatingSystemVersionUnexported* QOperatingSystemVersionUnexported_new(QOperatingSystemVersionBase* other) {
	return new QOperatingSystemVersionUnexported(*other);
}

QOperatingSystemVersionUnexported* QOperatingSystemVersionUnexported_new2() {
	return new QOperatingSystemVersionUnexported();
}

QOperatingSystemVersionUnexported* QOperatingSystemVersionUnexported_new3(QOperatingSystemVersionUnexported* param1) {
	return new QOperatingSystemVersionUnexported(*param1);
}

void QOperatingSystemVersionUnexported_virtbase(QOperatingSystemVersionUnexported* src, QOperatingSystemVersionBase** outptr_QOperatingSystemVersionBase) {
	*outptr_QOperatingSystemVersionBase = static_cast<QOperatingSystemVersionBase*>(src);
}

void QOperatingSystemVersionUnexported_Delete(QOperatingSystemVersionUnexported* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QOperatingSystemVersionUnexported*>( self );
	} else {
		delete self;
	}
}

QOperatingSystemVersion* QOperatingSystemVersion_new(QOperatingSystemVersionBase* osversion) {
	return new QOperatingSystemVersion(*osversion);
}

QOperatingSystemVersion* QOperatingSystemVersion_new2(OSType osType, int vmajor) {
	return new QOperatingSystemVersion(osType, static_cast<int>(vmajor));
}

QOperatingSystemVersion* QOperatingSystemVersion_new3(QOperatingSystemVersion* param1) {
	return new QOperatingSystemVersion(*param1);
}

QOperatingSystemVersion* QOperatingSystemVersion_new4(OSType osType, int vmajor, int vminor) {
	return new QOperatingSystemVersion(osType, static_cast<int>(vmajor), static_cast<int>(vminor));
}

QOperatingSystemVersion* QOperatingSystemVersion_new5(OSType osType, int vmajor, int vminor, int vmicro) {
	return new QOperatingSystemVersion(osType, static_cast<int>(vmajor), static_cast<int>(vminor), static_cast<int>(vmicro));
}

void QOperatingSystemVersion_virtbase(QOperatingSystemVersion* src, QOperatingSystemVersionUnexported** outptr_QOperatingSystemVersionUnexported) {
	*outptr_QOperatingSystemVersionUnexported = static_cast<QOperatingSystemVersionUnexported*>(src);
}

OSType QOperatingSystemVersion_CurrentType() {
	return QOperatingSystemVersion::currentType();
}

OSType QOperatingSystemVersion_Type(const QOperatingSystemVersion* self) {
	return self->type();
}

void QOperatingSystemVersion_Delete(QOperatingSystemVersion* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QOperatingSystemVersion*>( self );
	} else {
		delete self;
	}
}

