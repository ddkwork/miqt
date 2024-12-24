// +build ignore

#include <QPixmap>
#include <QPixmapCache>
#define WORKAROUND_INNER_CLASS_DEFINITION_QPixmapCache__Key
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qpixmapcache.h>
#include "gen_qpixmapcache.h"
int QPixmapCache_CacheLimit() {
	return QPixmapCache::cacheLimit();
}
void QPixmapCache_SetCacheLimit(int cacheLimit) {
	QPixmapCache::setCacheLimit(static_cast<int>(cacheLimit));
}
bool QPixmapCache_Find(struct miqt_string key, QPixmap* pixmap) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return QPixmapCache::find(key_QString, pixmap);
}
bool QPixmapCache_Find2(const Key* key, QPixmap* pixmap) {
	return QPixmapCache::find(*key, pixmap);
}
bool QPixmapCache_Insert(struct miqt_string key, QPixmap* pixmap) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return QPixmapCache::insert(key_QString, *pixmap);
}
Key QPixmapCache_InsertWithPixmap(QPixmap* pixmap) {
	return QPixmapCache::insert(*pixmap);
}
bool QPixmapCache_Replace(const Key* key, QPixmap* pixmap) {
	return QPixmapCache::replace(*key, *pixmap);
}
void QPixmapCache_Remove(struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	QPixmapCache::remove(key_QString);
}
void QPixmapCache_RemoveWithKey(const Key* key) {
	QPixmapCache::remove(*key);
}
void QPixmapCache_Clear() {
	QPixmapCache::clear();
}
void QPixmapCache_Delete(QPixmapCache* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPixmapCache*>( self );
	} else {
		delete self;
	}
}
QPixmapCache__Key* QPixmapCache__Key_new() {
return new QPixmapCache::Key();
}
QPixmapCache__Key* QPixmapCache__Key_new2(const Key* other) {
return new QPixmapCache::Key(*other);
}
bool QPixmapCache__Key_OperatorEqual(const QPixmapCache__Key* self, const Key* key) {
	return (*self == *key);
}
bool QPixmapCache__Key_OperatorNotEqual(const QPixmapCache__Key* self, const Key* key) {
	return (*self != *key);
}
void QPixmapCache__Key_OperatorAssign(QPixmapCache__Key* self, const Key* other) {
	self->operator=(*other);
}
void QPixmapCache__Key_Swap(QPixmapCache__Key* self, Key* other) {
	self->swap(*other);
}
bool QPixmapCache__Key_IsValid(const QPixmapCache__Key* self) {
	return self->isValid();
}
void QPixmapCache__Key_Delete(QPixmapCache__Key* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPixmapCache::Key*>( self );
	} else {
		delete self;
	}
}
