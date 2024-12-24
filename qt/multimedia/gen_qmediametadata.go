package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QMediaMetaData__Key int

const (
	QMediaMetaData__Title              QMediaMetaData__Key = 0
	QMediaMetaData__Author             QMediaMetaData__Key = 1
	QMediaMetaData__Comment            QMediaMetaData__Key = 2
	QMediaMetaData__Description        QMediaMetaData__Key = 3
	QMediaMetaData__Genre              QMediaMetaData__Key = 4
	QMediaMetaData__Date               QMediaMetaData__Key = 5
	QMediaMetaData__Language           QMediaMetaData__Key = 6
	QMediaMetaData__Publisher          QMediaMetaData__Key = 7
	QMediaMetaData__Copyright          QMediaMetaData__Key = 8
	QMediaMetaData__Url                QMediaMetaData__Key = 9
	QMediaMetaData__Duration           QMediaMetaData__Key = 10
	QMediaMetaData__MediaType          QMediaMetaData__Key = 11
	QMediaMetaData__FileFormat         QMediaMetaData__Key = 12
	QMediaMetaData__AudioBitRate       QMediaMetaData__Key = 13
	QMediaMetaData__AudioCodec         QMediaMetaData__Key = 14
	QMediaMetaData__VideoBitRate       QMediaMetaData__Key = 15
	QMediaMetaData__VideoCodec         QMediaMetaData__Key = 16
	QMediaMetaData__VideoFrameRate     QMediaMetaData__Key = 17
	QMediaMetaData__AlbumTitle         QMediaMetaData__Key = 18
	QMediaMetaData__AlbumArtist        QMediaMetaData__Key = 19
	QMediaMetaData__ContributingArtist QMediaMetaData__Key = 20
	QMediaMetaData__TrackNumber        QMediaMetaData__Key = 21
	QMediaMetaData__Composer           QMediaMetaData__Key = 22
	QMediaMetaData__LeadPerformer      QMediaMetaData__Key = 23
	QMediaMetaData__ThumbnailImage     QMediaMetaData__Key = 24
	QMediaMetaData__CoverArtImage      QMediaMetaData__Key = 25
	QMediaMetaData__Orientation        QMediaMetaData__Key = 26
	QMediaMetaData__Resolution         QMediaMetaData__Key = 27
	QMediaMetaData__HasHdrContent      QMediaMetaData__Key = 28
)

type QMediaMetaData struct {
	h          uintptr
	isSubclass bool
}

// NewQMediaMetaData constructs a new QMediaMetaData object.
func NewQMediaMetaData(param1 *QMediaMetaData) *QMediaMetaData {
	g := newQMediaMetaData(QMediaMetaData_new(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMediaMetaData2 constructs a new QMediaMetaData object.
func NewQMediaMetaData2() *QMediaMetaData {
	g := newQMediaMetaData(QMediaMetaData_new2())
	g.isSubclass = true
	return g
}

func (this *QMediaMetaData) Value(k Key) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QMediaMetaData_Value(this.h, k)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaMetaData) Insert(k Key, value *qt.QVariant) {
	QMediaMetaData_Insert(this.h, k, (*QVariant)(value.UnsafePointer()))
}

func (this *QMediaMetaData) Remove(k Key) {
	QMediaMetaData_Remove(this.h, k)
}

func (this *QMediaMetaData) Keys() []Key {
	var _ma struct_miqt_array = QMediaMetaData_Keys(this.h)
	_ret := make([]Key, int(_ma.len))
	_outCast := (*[0xffff]Key)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QMediaMetaData) OperatorSubscript(k Key) *qt.QVariant {
	return qt.UnsafeNewQVariant(unsafe.Pointer(QMediaMetaData_OperatorSubscript(this.h, k)))
}

func (this *QMediaMetaData) Clear() {
	QMediaMetaData_Clear(this.h)
}

func (this *QMediaMetaData) IsEmpty() bool {
	return (bool)(QMediaMetaData_IsEmpty(this.h))
}

func (this *QMediaMetaData) StringValue(k Key) string {
	var _ms struct_miqt_string = QMediaMetaData_StringValue(this.h, k)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaMetaData_MetaDataKeyToString(k Key) string {
	var _ms struct_miqt_string = QMediaMetaData_MetaDataKeyToString(k)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
