package qt

import (
	"unsafe"
)

type QPixmapCache struct {
	h          uintptr
	isSubclass bool
}

func QPixmapCache_CacheLimit() int {
	return (int)(QPixmapCache_CacheLimit())
}

func QPixmapCache_SetCacheLimit(cacheLimit int) {
	QPixmapCache_SetCacheLimit((int)(cacheLimit))
}

func QPixmapCache_Find(key string, pixmap *QPixmap) bool {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	return (bool)(QPixmapCache_Find(key_ms, pixmap.cPointer()))
}

func QPixmapCache_Find2(key *Key, pixmap *QPixmap) bool {
	return (bool)(QPixmapCache_Find2(key, pixmap.cPointer()))
}

func QPixmapCache_Insert(key string, pixmap *QPixmap) bool {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	return (bool)(QPixmapCache_Insert(key_ms, pixmap.cPointer()))
}

func QPixmapCache_InsertWithPixmap(pixmap *QPixmap) Key {
	xxxxxxxxx
}

func QPixmapCache_Replace(key *Key, pixmap *QPixmap) bool {
	return (bool)(QPixmapCache_Replace(key, pixmap.cPointer()))
}

func QPixmapCache_Remove(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QPixmapCache_Remove(key_ms)
}

func QPixmapCache_RemoveWithKey(key *Key) {
	QPixmapCache_RemoveWithKey(key)
}

func QPixmapCache_Clear() {
	QPixmapCache_Clear()
}

type QPixmapCache__Key struct {
	h          uintptr
	isSubclass bool
}

// NewQPixmapCache__Key constructs a new QPixmapCache::Key object.
func NewQPixmapCache__Key() *QPixmapCache__Key {
	g := newQPixmapCache__Key(QPixmapCache__Key_new())
	g.isSubclass = true
	return g
}

// NewQPixmapCache__Key2 constructs a new QPixmapCache::Key object.
func NewQPixmapCache__Key2(other *Key) *QPixmapCache__Key {
	g := newQPixmapCache__Key(QPixmapCache__Key_new2(other))
	g.isSubclass = true
	return g
}

func (this *QPixmapCache__Key) OperatorEqual(key *Key) bool {
	return (bool)(QPixmapCache__Key_OperatorEqual(this.h, key))
}

func (this *QPixmapCache__Key) OperatorNotEqual(key *Key) bool {
	return (bool)(QPixmapCache__Key_OperatorNotEqual(this.h, key))
}

func (this *QPixmapCache__Key) OperatorAssign(other *Key) {
	QPixmapCache__Key_OperatorAssign(this.h, other)
}

func (this *QPixmapCache__Key) Swap(other *Key) {
	QPixmapCache__Key_Swap(this.h, other)
}

func (this *QPixmapCache__Key) IsValid() bool {
	return (bool)(QPixmapCache__Key_IsValid(this.h))
}
