// +build ignore

#include <QLockFile>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qlockfile.h>
#include "gen_qlockfile.h"

QLockFile* QLockFile_new(struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new QLockFile(fileName_QString);
}

struct miqt_string QLockFile_FileName(const QLockFile* self) {
	QString _ret = self->fileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QLockFile_Lock(QLockFile* self) {
	return self->lock();
}

bool QLockFile_TryLock(QLockFile* self, int timeout) {
	return self->tryLock(static_cast<int>(timeout));
}

void QLockFile_Unlock(QLockFile* self) {
	self->unlock();
}

void QLockFile_SetStaleLockTime(QLockFile* self, int staleLockTime) {
	self->setStaleLockTime(static_cast<int>(staleLockTime));
}

int QLockFile_StaleLockTime(const QLockFile* self) {
	return self->staleLockTime();
}

bool QLockFile_TryLock2(QLockFile* self) {
	return self->tryLock();
}

bool QLockFile_IsLocked(const QLockFile* self) {
	return self->isLocked();
}

bool QLockFile_RemoveStaleLockFile(QLockFile* self) {
	return self->removeStaleLockFile();
}

LockError QLockFile_Error(const QLockFile* self) {
	return self->error();
}

void QLockFile_Delete(QLockFile* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QLockFile*>( self );
	} else {
		delete self;
	}
}

