// +build ignore

#include <QDateTime>
#include <QDirListing>
#define WORKAROUND_INNER_CLASS_DEFINITION_QDirListing__DirEntry
#define WORKAROUND_INNER_CLASS_DEFINITION_QDirListing__const_iterator
#define WORKAROUND_INNER_CLASS_DEFINITION_QDirListing__sentinel
#include <QFileInfo>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimeZone>
#include <qdirlisting.h>
#include "gen_qdirlisting.h"

#ifndef _Bool
#define _Bool bool
#endif

QDirListing* QDirListing_new(struct miqt_string path) {
	QString path_QString = QString::fromUtf8(path.data, path.len);
	return new QDirListing(path_QString);
}

QDirListing* QDirListing_new2(struct miqt_string path, struct miqt_array /* of struct miqt_string */  nameFilters) {
	QString path_QString = QString::fromUtf8(path.data, path.len);
	QStringList nameFilters_QList;
	nameFilters_QList.reserve(nameFilters.len);
	struct miqt_string* nameFilters_arr = static_cast<struct miqt_string*>(nameFilters.data);
	for(size_t i = 0; i < nameFilters.len; ++i) {
		QString nameFilters_arr_i_QString = QString::fromUtf8(nameFilters_arr[i].data, nameFilters_arr[i].len);
		nameFilters_QList.push_back(nameFilters_arr_i_QString);
	}
	return new QDirListing(path_QString, nameFilters_QList);
}

QDirListing* QDirListing_new3(struct miqt_string path, IteratorFlags flags) {
	QString path_QString = QString::fromUtf8(path.data, path.len);
	return new QDirListing(path_QString, flags);
}

QDirListing* QDirListing_new4(struct miqt_string path, struct miqt_array /* of struct miqt_string */  nameFilters, IteratorFlags flags) {
	QString path_QString = QString::fromUtf8(path.data, path.len);
	QStringList nameFilters_QList;
	nameFilters_QList.reserve(nameFilters.len);
	struct miqt_string* nameFilters_arr = static_cast<struct miqt_string*>(nameFilters.data);
	for(size_t i = 0; i < nameFilters.len; ++i) {
		QString nameFilters_arr_i_QString = QString::fromUtf8(nameFilters_arr[i].data, nameFilters_arr[i].len);
		nameFilters_QList.push_back(nameFilters_arr_i_QString);
	}
	return new QDirListing(path_QString, nameFilters_QList, flags);
}

void QDirListing_Swap(QDirListing* self, QDirListing* other) {
	self->swap(*other);
}

struct miqt_string QDirListing_IteratorPath(const QDirListing* self) {
	QString _ret = self->iteratorPath();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

IteratorFlags QDirListing_IteratorFlags(const QDirListing* self) {
	return self->iteratorFlags();
}

struct miqt_array /* of struct miqt_string */  QDirListing_NameFilters(const QDirListing* self) {
	QStringList _ret = self->nameFilters();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

const_iterator QDirListing_Begin(const QDirListing* self) {
	return self->begin();
}

const_iterator QDirListing_Cbegin(const QDirListing* self) {
	return self->cbegin();
}

sentinel QDirListing_End(const QDirListing* self) {
	return self->end();
}

sentinel QDirListing_Cend(const QDirListing* self) {
	return self->cend();
}

const_iterator QDirListing_ConstBegin(const QDirListing* self) {
	return self->constBegin();
}

sentinel QDirListing_ConstEnd(const QDirListing* self) {
	return self->constEnd();
}

void QDirListing_Delete(QDirListing* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDirListing*>( self );
	} else {
		delete self;
	}
}

QDirListing__DirEntry* QDirListing__DirEntry_new(const DirEntry* param1) {
	return new QDirListing::DirEntry(*param1);
}

QDirListing__DirEntry* QDirListing__DirEntry_new2() {
	return new QDirListing::DirEntry();
}

struct miqt_string QDirListing__DirEntry_FileName(const QDirListing__DirEntry* self) {
	QString _ret = self->fileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDirListing__DirEntry_BaseName(const QDirListing__DirEntry* self) {
	QString _ret = self->baseName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDirListing__DirEntry_CompleteBaseName(const QDirListing__DirEntry* self) {
	QString _ret = self->completeBaseName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDirListing__DirEntry_Suffix(const QDirListing__DirEntry* self) {
	QString _ret = self->suffix();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDirListing__DirEntry_BundleName(const QDirListing__DirEntry* self) {
	QString _ret = self->bundleName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDirListing__DirEntry_CompleteSuffix(const QDirListing__DirEntry* self) {
	QString _ret = self->completeSuffix();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDirListing__DirEntry_FilePath(const QDirListing__DirEntry* self) {
	QString _ret = self->filePath();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QDirListing__DirEntry_IsDir(const QDirListing__DirEntry* self) {
	return self->isDir();
}

bool QDirListing__DirEntry_IsFile(const QDirListing__DirEntry* self) {
	return self->isFile();
}

bool QDirListing__DirEntry_IsSymLink(const QDirListing__DirEntry* self) {
	return self->isSymLink();
}

bool QDirListing__DirEntry_Exists(const QDirListing__DirEntry* self) {
	return self->exists();
}

bool QDirListing__DirEntry_IsHidden(const QDirListing__DirEntry* self) {
	return self->isHidden();
}

bool QDirListing__DirEntry_IsReadable(const QDirListing__DirEntry* self) {
	return self->isReadable();
}

bool QDirListing__DirEntry_IsWritable(const QDirListing__DirEntry* self) {
	return self->isWritable();
}

bool QDirListing__DirEntry_IsExecutable(const QDirListing__DirEntry* self) {
	return self->isExecutable();
}

QFileInfo* QDirListing__DirEntry_FileInfo(const QDirListing__DirEntry* self) {
	return new QFileInfo(self->fileInfo());
}

struct miqt_string QDirListing__DirEntry_CanonicalFilePath(const QDirListing__DirEntry* self) {
	QString _ret = self->canonicalFilePath();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDirListing__DirEntry_AbsoluteFilePath(const QDirListing__DirEntry* self) {
	QString _ret = self->absoluteFilePath();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDirListing__DirEntry_AbsolutePath(const QDirListing__DirEntry* self) {
	QString _ret = self->absolutePath();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

long long QDirListing__DirEntry_Size(const QDirListing__DirEntry* self) {
	qint64 _ret = self->size();
	return static_cast<long long>(_ret);
}

QDateTime* QDirListing__DirEntry_BirthTime(const QDirListing__DirEntry* self, QTimeZone* tz) {
	return new QDateTime(self->birthTime(*tz));
}

QDateTime* QDirListing__DirEntry_MetadataChangeTime(const QDirListing__DirEntry* self, QTimeZone* tz) {
	return new QDateTime(self->metadataChangeTime(*tz));
}

QDateTime* QDirListing__DirEntry_LastModified(const QDirListing__DirEntry* self, QTimeZone* tz) {
	return new QDateTime(self->lastModified(*tz));
}

QDateTime* QDirListing__DirEntry_LastRead(const QDirListing__DirEntry* self, QTimeZone* tz) {
	return new QDateTime(self->lastRead(*tz));
}

QDateTime* QDirListing__DirEntry_FileTime(const QDirListing__DirEntry* self, int typeVal, QTimeZone* tz) {
	return new QDateTime(self->fileTime(static_cast<QFileDevice::FileTime>(typeVal), *tz));
}

void QDirListing__DirEntry_OperatorAssign(QDirListing__DirEntry* self, const DirEntry* param1) {
	self->operator=(*param1);
}

void QDirListing__DirEntry_Delete(QDirListing__DirEntry* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDirListing::DirEntry*>( self );
	} else {
		delete self;
	}
}

QDirListing__sentinel* QDirListing__sentinel_new() {
	return new QDirListing::sentinel();
}

QDirListing__sentinel* QDirListing__sentinel_new2(const sentinel* param1) {
	return new QDirListing::sentinel(*param1);
}

void QDirListing__sentinel_Delete(QDirListing__sentinel* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDirListing::sentinel*>( self );
	} else {
		delete self;
	}
}

QDirListing__const_iterator* QDirListing__const_iterator_new() {
	return new QDirListing::const_iterator();
}

reference QDirListing__const_iterator_OperatorMultiply(const QDirListing__const_iterator* self) {
	return self->operator*();
}

pointer QDirListing__const_iterator_OperatorMinusGreater(const QDirListing__const_iterator* self) {
	return self->operator->();
}

const_iterator* QDirListing__const_iterator_OperatorPlusPlus(QDirListing__const_iterator* self) {
	return &self->operator++();
}

void QDirListing__const_iterator_OperatorPlusPlusWithInt(QDirListing__const_iterator* self, int param1) {
	self->operator++(static_cast<int>(param1));
}

void QDirListing__const_iterator_Delete(QDirListing__const_iterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDirListing::const_iterator*>( self );
	} else {
		delete self;
	}
}

