package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QMediaFormat__FileFormat int

const (
	QMediaFormat__FileFormat__UnspecifiedFormat QMediaFormat__FileFormat = -1
	QMediaFormat__FileFormat__WMV               QMediaFormat__FileFormat = 0
	QMediaFormat__FileFormat__AVI               QMediaFormat__FileFormat = 1
	QMediaFormat__FileFormat__Matroska          QMediaFormat__FileFormat = 2
	QMediaFormat__FileFormat__MPEG4             QMediaFormat__FileFormat = 3
	QMediaFormat__FileFormat__Ogg               QMediaFormat__FileFormat = 4
	QMediaFormat__FileFormat__QuickTime         QMediaFormat__FileFormat = 5
	QMediaFormat__FileFormat__WebM              QMediaFormat__FileFormat = 6
	QMediaFormat__FileFormat__Mpeg4Audio        QMediaFormat__FileFormat = 7
	QMediaFormat__FileFormat__AAC               QMediaFormat__FileFormat = 8
	QMediaFormat__FileFormat__WMA               QMediaFormat__FileFormat = 9
	QMediaFormat__FileFormat__MP3               QMediaFormat__FileFormat = 10
	QMediaFormat__FileFormat__FLAC              QMediaFormat__FileFormat = 11
	QMediaFormat__FileFormat__Wave              QMediaFormat__FileFormat = 12
	QMediaFormat__FileFormat__LastFileFormat    QMediaFormat__FileFormat = 12
)

type QMediaFormat__AudioCodec int

const (
	QMediaFormat__AudioCodec__Unspecified    QMediaFormat__AudioCodec = -1
	QMediaFormat__AudioCodec__MP3            QMediaFormat__AudioCodec = 0
	QMediaFormat__AudioCodec__AAC            QMediaFormat__AudioCodec = 1
	QMediaFormat__AudioCodec__AC3            QMediaFormat__AudioCodec = 2
	QMediaFormat__AudioCodec__EAC3           QMediaFormat__AudioCodec = 3
	QMediaFormat__AudioCodec__FLAC           QMediaFormat__AudioCodec = 4
	QMediaFormat__AudioCodec__DolbyTrueHD    QMediaFormat__AudioCodec = 5
	QMediaFormat__AudioCodec__Opus           QMediaFormat__AudioCodec = 6
	QMediaFormat__AudioCodec__Vorbis         QMediaFormat__AudioCodec = 7
	QMediaFormat__AudioCodec__Wave           QMediaFormat__AudioCodec = 8
	QMediaFormat__AudioCodec__WMA            QMediaFormat__AudioCodec = 9
	QMediaFormat__AudioCodec__ALAC           QMediaFormat__AudioCodec = 10
	QMediaFormat__AudioCodec__LastAudioCodec QMediaFormat__AudioCodec = 10
)

type QMediaFormat__VideoCodec int

const (
	QMediaFormat__VideoCodec__Unspecified    QMediaFormat__VideoCodec = -1
	QMediaFormat__VideoCodec__MPEG1          QMediaFormat__VideoCodec = 0
	QMediaFormat__VideoCodec__MPEG2          QMediaFormat__VideoCodec = 1
	QMediaFormat__VideoCodec__MPEG4          QMediaFormat__VideoCodec = 2
	QMediaFormat__VideoCodec__H264           QMediaFormat__VideoCodec = 3
	QMediaFormat__VideoCodec__H265           QMediaFormat__VideoCodec = 4
	QMediaFormat__VideoCodec__VP8            QMediaFormat__VideoCodec = 5
	QMediaFormat__VideoCodec__VP9            QMediaFormat__VideoCodec = 6
	QMediaFormat__VideoCodec__AV1            QMediaFormat__VideoCodec = 7
	QMediaFormat__VideoCodec__Theora         QMediaFormat__VideoCodec = 8
	QMediaFormat__VideoCodec__WMV            QMediaFormat__VideoCodec = 9
	QMediaFormat__VideoCodec__MotionJPEG     QMediaFormat__VideoCodec = 10
	QMediaFormat__VideoCodec__LastVideoCodec QMediaFormat__VideoCodec = 10
)

type QMediaFormat__ConversionMode int

const (
	QMediaFormat__Encode QMediaFormat__ConversionMode = 0
	QMediaFormat__Decode QMediaFormat__ConversionMode = 1
)

type QMediaFormat__ResolveFlags int

const (
	QMediaFormat__NoFlags       QMediaFormat__ResolveFlags = 0
	QMediaFormat__RequiresVideo QMediaFormat__ResolveFlags = 1
)

type QMediaFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQMediaFormat constructs a new QMediaFormat object.
func NewQMediaFormat() *QMediaFormat {
	g := newQMediaFormat(QMediaFormat_new())
	g.isSubclass = true
	return g
}

// NewQMediaFormat2 constructs a new QMediaFormat object.
func NewQMediaFormat2(other *QMediaFormat) *QMediaFormat {
	g := newQMediaFormat(QMediaFormat_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMediaFormat3 constructs a new QMediaFormat object.
func NewQMediaFormat3(format FileFormat) *QMediaFormat {
	g := newQMediaFormat(QMediaFormat_new3(format))
	g.isSubclass = true
	return g
}

func (this *QMediaFormat) OperatorAssign(other *QMediaFormat) {
	QMediaFormat_OperatorAssign(this.h, other.cPointer())
}

func (this *QMediaFormat) Swap(other *QMediaFormat) {
	QMediaFormat_Swap(this.h, other.cPointer())
}

func (this *QMediaFormat) FileFormat() FileFormat {
	xxxxxxxxx
}

func (this *QMediaFormat) SetFileFormat(f FileFormat) {
	QMediaFormat_SetFileFormat(this.h, f)
}

func (this *QMediaFormat) SetVideoCodec(codec VideoCodec) {
	QMediaFormat_SetVideoCodec(this.h, codec)
}

func (this *QMediaFormat) VideoCodec() VideoCodec {
	xxxxxxxxx
}

func (this *QMediaFormat) SetAudioCodec(codec AudioCodec) {
	QMediaFormat_SetAudioCodec(this.h, codec)
}

func (this *QMediaFormat) AudioCodec() AudioCodec {
	xxxxxxxxx
}

func (this *QMediaFormat) IsSupported(mode ConversionMode) bool {
	return (bool)(QMediaFormat_IsSupported(this.h, mode))
}

func (this *QMediaFormat) MimeType() *qt.QMimeType {
	_goptr := qt.UnsafeNewQMimeType(unsafe.Pointer(QMediaFormat_MimeType(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaFormat) SupportedFileFormats(m ConversionMode) []FileFormat {
	var _ma struct_miqt_array = QMediaFormat_SupportedFileFormats(this.h, m)
	_ret := make([]FileFormat, int(_ma.len))
	_outCast := (*[0xffff]FileFormat)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QMediaFormat) SupportedVideoCodecs(m ConversionMode) []VideoCodec {
	var _ma struct_miqt_array = QMediaFormat_SupportedVideoCodecs(this.h, m)
	_ret := make([]VideoCodec, int(_ma.len))
	_outCast := (*[0xffff]VideoCodec)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QMediaFormat) SupportedAudioCodecs(m ConversionMode) []AudioCodec {
	var _ma struct_miqt_array = QMediaFormat_SupportedAudioCodecs(this.h, m)
	_ret := make([]AudioCodec, int(_ma.len))
	_outCast := (*[0xffff]AudioCodec)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func QMediaFormat_FileFormatName(fileFormat FileFormat) string {
	var _ms struct_miqt_string = QMediaFormat_FileFormatName(fileFormat)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaFormat_AudioCodecName(codec AudioCodec) string {
	var _ms struct_miqt_string = QMediaFormat_AudioCodecName(codec)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaFormat_VideoCodecName(codec VideoCodec) string {
	var _ms struct_miqt_string = QMediaFormat_VideoCodecName(codec)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaFormat_FileFormatDescription(fileFormat QMediaFormat__FileFormat) string {
	var _ms struct_miqt_string = QMediaFormat_FileFormatDescription((int)(fileFormat))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaFormat_AudioCodecDescription(codec QMediaFormat__AudioCodec) string {
	var _ms struct_miqt_string = QMediaFormat_AudioCodecDescription((int)(codec))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaFormat_VideoCodecDescription(codec QMediaFormat__VideoCodec) string {
	var _ms struct_miqt_string = QMediaFormat_VideoCodecDescription((int)(codec))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaFormat) OperatorEqual(other *QMediaFormat) bool {
	return (bool)(QMediaFormat_OperatorEqual(this.h, other.cPointer()))
}

func (this *QMediaFormat) OperatorNotEqual(other *QMediaFormat) bool {
	return (bool)(QMediaFormat_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QMediaFormat) ResolveForEncoding(flags ResolveFlags) {
	QMediaFormat_ResolveForEncoding(this.h, flags)
}
