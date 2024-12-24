package qt

import (
	"unsafe"
)

type QIcon__Mode int

const (
	QIcon__Normal   QIcon__Mode = 0
	QIcon__Disabled QIcon__Mode = 1
	QIcon__Active   QIcon__Mode = 2
	QIcon__Selected QIcon__Mode = 3
)

type QIcon__State int

const (
	QIcon__On  QIcon__State = 0
	QIcon__Off QIcon__State = 1
)

type QIcon__ThemeIcon int

const (
	QIcon__AddressBookNew          QIcon__ThemeIcon = 0
	QIcon__ApplicationExit         QIcon__ThemeIcon = 1
	QIcon__AppointmentNew          QIcon__ThemeIcon = 2
	QIcon__CallStart               QIcon__ThemeIcon = 3
	QIcon__CallStop                QIcon__ThemeIcon = 4
	QIcon__ContactNew              QIcon__ThemeIcon = 5
	QIcon__DocumentNew             QIcon__ThemeIcon = 6
	QIcon__DocumentOpen            QIcon__ThemeIcon = 7
	QIcon__DocumentOpenRecent      QIcon__ThemeIcon = 8
	QIcon__DocumentPageSetup       QIcon__ThemeIcon = 9
	QIcon__DocumentPrint           QIcon__ThemeIcon = 10
	QIcon__DocumentPrintPreview    QIcon__ThemeIcon = 11
	QIcon__DocumentProperties      QIcon__ThemeIcon = 12
	QIcon__DocumentRevert          QIcon__ThemeIcon = 13
	QIcon__DocumentSave            QIcon__ThemeIcon = 14
	QIcon__DocumentSaveAs          QIcon__ThemeIcon = 15
	QIcon__DocumentSend            QIcon__ThemeIcon = 16
	QIcon__EditClear               QIcon__ThemeIcon = 17
	QIcon__EditCopy                QIcon__ThemeIcon = 18
	QIcon__EditCut                 QIcon__ThemeIcon = 19
	QIcon__EditDelete              QIcon__ThemeIcon = 20
	QIcon__EditFind                QIcon__ThemeIcon = 21
	QIcon__EditPaste               QIcon__ThemeIcon = 22
	QIcon__EditRedo                QIcon__ThemeIcon = 23
	QIcon__EditSelectAll           QIcon__ThemeIcon = 24
	QIcon__EditUndo                QIcon__ThemeIcon = 25
	QIcon__FolderNew               QIcon__ThemeIcon = 26
	QIcon__FormatIndentLess        QIcon__ThemeIcon = 27
	QIcon__FormatIndentMore        QIcon__ThemeIcon = 28
	QIcon__FormatJustifyCenter     QIcon__ThemeIcon = 29
	QIcon__FormatJustifyFill       QIcon__ThemeIcon = 30
	QIcon__FormatJustifyLeft       QIcon__ThemeIcon = 31
	QIcon__FormatJustifyRight      QIcon__ThemeIcon = 32
	QIcon__FormatTextDirectionLtr  QIcon__ThemeIcon = 33
	QIcon__FormatTextDirectionRtl  QIcon__ThemeIcon = 34
	QIcon__FormatTextBold          QIcon__ThemeIcon = 35
	QIcon__FormatTextItalic        QIcon__ThemeIcon = 36
	QIcon__FormatTextUnderline     QIcon__ThemeIcon = 37
	QIcon__FormatTextStrikethrough QIcon__ThemeIcon = 38
	QIcon__GoDown                  QIcon__ThemeIcon = 39
	QIcon__GoHome                  QIcon__ThemeIcon = 40
	QIcon__GoNext                  QIcon__ThemeIcon = 41
	QIcon__GoPrevious              QIcon__ThemeIcon = 42
	QIcon__GoUp                    QIcon__ThemeIcon = 43
	QIcon__HelpAbout               QIcon__ThemeIcon = 44
	QIcon__HelpFaq                 QIcon__ThemeIcon = 45
	QIcon__InsertImage             QIcon__ThemeIcon = 46
	QIcon__InsertLink              QIcon__ThemeIcon = 47
	QIcon__InsertText              QIcon__ThemeIcon = 48
	QIcon__ListAdd                 QIcon__ThemeIcon = 49
	QIcon__ListRemove              QIcon__ThemeIcon = 50
	QIcon__MailForward             QIcon__ThemeIcon = 51
	QIcon__MailMarkImportant       QIcon__ThemeIcon = 52
	QIcon__MailMarkRead            QIcon__ThemeIcon = 53
	QIcon__MailMarkUnread          QIcon__ThemeIcon = 54
	QIcon__MailMessageNew          QIcon__ThemeIcon = 55
	QIcon__MailReplyAll            QIcon__ThemeIcon = 56
	QIcon__MailReplySender         QIcon__ThemeIcon = 57
	QIcon__MailSend                QIcon__ThemeIcon = 58
	QIcon__MediaEject              QIcon__ThemeIcon = 59
	QIcon__MediaPlaybackPause      QIcon__ThemeIcon = 60
	QIcon__MediaPlaybackStart      QIcon__ThemeIcon = 61
	QIcon__MediaPlaybackStop       QIcon__ThemeIcon = 62
	QIcon__MediaRecord             QIcon__ThemeIcon = 63
	QIcon__MediaSeekBackward       QIcon__ThemeIcon = 64
	QIcon__MediaSeekForward        QIcon__ThemeIcon = 65
	QIcon__MediaSkipBackward       QIcon__ThemeIcon = 66
	QIcon__MediaSkipForward        QIcon__ThemeIcon = 67
	QIcon__ObjectRotateLeft        QIcon__ThemeIcon = 68
	QIcon__ObjectRotateRight       QIcon__ThemeIcon = 69
	QIcon__ProcessStop             QIcon__ThemeIcon = 70
	QIcon__SystemLockScreen        QIcon__ThemeIcon = 71
	QIcon__SystemLogOut            QIcon__ThemeIcon = 72
	QIcon__SystemSearch            QIcon__ThemeIcon = 73
	QIcon__SystemReboot            QIcon__ThemeIcon = 74
	QIcon__SystemShutdown          QIcon__ThemeIcon = 75
	QIcon__ToolsCheckSpelling      QIcon__ThemeIcon = 76
	QIcon__ViewFullscreen          QIcon__ThemeIcon = 77
	QIcon__ViewRefresh             QIcon__ThemeIcon = 78
	QIcon__ViewRestore             QIcon__ThemeIcon = 79
	QIcon__WindowClose             QIcon__ThemeIcon = 80
	QIcon__WindowNew               QIcon__ThemeIcon = 81
	QIcon__ZoomFitBest             QIcon__ThemeIcon = 82
	QIcon__ZoomIn                  QIcon__ThemeIcon = 83
	QIcon__ZoomOut                 QIcon__ThemeIcon = 84
	QIcon__AudioCard               QIcon__ThemeIcon = 85
	QIcon__AudioInputMicrophone    QIcon__ThemeIcon = 86
	QIcon__Battery                 QIcon__ThemeIcon = 87
	QIcon__CameraPhoto             QIcon__ThemeIcon = 88
	QIcon__CameraVideo             QIcon__ThemeIcon = 89
	QIcon__CameraWeb               QIcon__ThemeIcon = 90
	QIcon__Computer                QIcon__ThemeIcon = 91
	QIcon__DriveHarddisk           QIcon__ThemeIcon = 92
	QIcon__DriveOptical            QIcon__ThemeIcon = 93
	QIcon__InputGaming             QIcon__ThemeIcon = 94
	QIcon__InputKeyboard           QIcon__ThemeIcon = 95
	QIcon__InputMouse              QIcon__ThemeIcon = 96
	QIcon__InputTablet             QIcon__ThemeIcon = 97
	QIcon__MediaFlash              QIcon__ThemeIcon = 98
	QIcon__MediaOptical            QIcon__ThemeIcon = 99
	QIcon__MediaTape               QIcon__ThemeIcon = 100
	QIcon__MultimediaPlayer        QIcon__ThemeIcon = 101
	QIcon__NetworkWired            QIcon__ThemeIcon = 102
	QIcon__NetworkWireless         QIcon__ThemeIcon = 103
	QIcon__Phone                   QIcon__ThemeIcon = 104
	QIcon__Printer                 QIcon__ThemeIcon = 105
	QIcon__Scanner                 QIcon__ThemeIcon = 106
	QIcon__VideoDisplay            QIcon__ThemeIcon = 107
	QIcon__AppointmentMissed       QIcon__ThemeIcon = 108
	QIcon__AppointmentSoon         QIcon__ThemeIcon = 109
	QIcon__AudioVolumeHigh         QIcon__ThemeIcon = 110
	QIcon__AudioVolumeLow          QIcon__ThemeIcon = 111
	QIcon__AudioVolumeMedium       QIcon__ThemeIcon = 112
	QIcon__AudioVolumeMuted        QIcon__ThemeIcon = 113
	QIcon__BatteryCaution          QIcon__ThemeIcon = 114
	QIcon__BatteryLow              QIcon__ThemeIcon = 115
	QIcon__DialogError             QIcon__ThemeIcon = 116
	QIcon__DialogInformation       QIcon__ThemeIcon = 117
	QIcon__DialogPassword          QIcon__ThemeIcon = 118
	QIcon__DialogQuestion          QIcon__ThemeIcon = 119
	QIcon__DialogWarning           QIcon__ThemeIcon = 120
	QIcon__FolderDragAccept        QIcon__ThemeIcon = 121
	QIcon__FolderOpen              QIcon__ThemeIcon = 122
	QIcon__FolderVisiting          QIcon__ThemeIcon = 123
	QIcon__ImageLoading            QIcon__ThemeIcon = 124
	QIcon__ImageMissing            QIcon__ThemeIcon = 125
	QIcon__MailAttachment          QIcon__ThemeIcon = 126
	QIcon__MailUnread              QIcon__ThemeIcon = 127
	QIcon__MailRead                QIcon__ThemeIcon = 128
	QIcon__MailReplied             QIcon__ThemeIcon = 129
	QIcon__MediaPlaylistRepeat     QIcon__ThemeIcon = 130
	QIcon__MediaPlaylistShuffle    QIcon__ThemeIcon = 131
	QIcon__NetworkOffline          QIcon__ThemeIcon = 132
	QIcon__PrinterPrinting         QIcon__ThemeIcon = 133
	QIcon__SecurityHigh            QIcon__ThemeIcon = 134
	QIcon__SecurityLow             QIcon__ThemeIcon = 135
	QIcon__SoftwareUpdateAvailable QIcon__ThemeIcon = 136
	QIcon__SoftwareUpdateUrgent    QIcon__ThemeIcon = 137
	QIcon__SyncError               QIcon__ThemeIcon = 138
	QIcon__SyncSynchronizing       QIcon__ThemeIcon = 139
	QIcon__UserAvailable           QIcon__ThemeIcon = 140
	QIcon__UserOffline             QIcon__ThemeIcon = 141
	QIcon__WeatherClear            QIcon__ThemeIcon = 142
	QIcon__WeatherClearNight       QIcon__ThemeIcon = 143
	QIcon__WeatherFewClouds        QIcon__ThemeIcon = 144
	QIcon__WeatherFewCloudsNight   QIcon__ThemeIcon = 145
	QIcon__WeatherFog              QIcon__ThemeIcon = 146
	QIcon__WeatherShowers          QIcon__ThemeIcon = 147
	QIcon__WeatherSnow             QIcon__ThemeIcon = 148
	QIcon__WeatherStorm            QIcon__ThemeIcon = 149
	QIcon__NThemeIcons             QIcon__ThemeIcon = 150
)

type QIcon struct {
	h          uintptr
	isSubclass bool
}

// NewQIcon constructs a new QIcon object.
func NewQIcon() *QIcon {
	ret := newQIcon(QIcon_new())
	ret.isSubclass = true
	return ret
}

// NewQIcon2 constructs a new QIcon object.
func NewQIcon2(pixmap *QPixmap) *QIcon {
	ret := newQIcon(QIcon_new2(pixmap.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQIcon3 constructs a new QIcon object.
func NewQIcon3(other *QIcon) *QIcon {
	ret := newQIcon(QIcon_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQIcon4 constructs a new QIcon object.
func NewQIcon4(fileName string) *QIcon {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQIcon(QIcon_new4(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQIcon5 constructs a new QIcon object.
func NewQIcon5(engine *QIconEngine) *QIcon {
	ret := newQIcon(QIcon_new5(engine.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QIcon) OperatorAssign(other *QIcon) {
	QIcon_OperatorAssign(this.h, other.cPointer())
}

func (this *QIcon) Swap(other *QIcon) {
	QIcon_Swap(this.h, other.cPointer())
}

func (this *QIcon) Pixmap(size *QSize) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap(this.h, size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap2(w int, h int) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap2(this.h, (int)(w), (int)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) PixmapWithExtent(extent int) *QPixmap {
	_goptr := newQPixmap(QIcon_PixmapWithExtent(this.h, (int)(extent)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap3(size *QSize, devicePixelRatio float64) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap3(this.h, size.cPointer(), (double)(devicePixelRatio)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap4(window *QWindow, size *QSize) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap4(this.h, window.cPointer(), size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize(size *QSize) *QSize {
	_goptr := newQSize(QIcon_ActualSize(this.h, size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize2(window *QWindow, size *QSize) *QSize {
	_goptr := newQSize(QIcon_ActualSize2(this.h, window.cPointer(), size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Name() string {
	var _ms struct_miqt_string = QIcon_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIcon) Paint(painter *QPainter, rect *QRect) {
	QIcon_Paint(this.h, painter.cPointer(), rect.cPointer())
}

func (this *QIcon) Paint2(painter *QPainter, x int, y int, w int, h int) {
	QIcon_Paint2(this.h, painter.cPointer(), (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QIcon) IsNull() bool {
	return (bool)(QIcon_IsNull(this.h))
}

func (this *QIcon) IsDetached() bool {
	return (bool)(QIcon_IsDetached(this.h))
}

func (this *QIcon) Detach() {
	QIcon_Detach(this.h)
}

func (this *QIcon) CacheKey() int64 {
	return (int64)(QIcon_CacheKey(this.h))
}

func (this *QIcon) AddPixmap(pixmap *QPixmap) {
	QIcon_AddPixmap(this.h, pixmap.cPointer())
}

func (this *QIcon) AddFile(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QIcon_AddFile(this.h, fileName_ms)
}

func (this *QIcon) AvailableSizes() []QSize {
	var _ma struct_miqt_array = QIcon_AvailableSizes(this.h)
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSize(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QIcon) SetIsMask(isMask bool) {
	QIcon_SetIsMask(this.h, (bool)(isMask))
}

func (this *QIcon) IsMask() bool {
	return (bool)(QIcon_IsMask(this.h))
}

func QIcon_FromTheme(name string) *QIcon {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	_goptr := newQIcon(QIcon_FromTheme(name_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QIcon_FromTheme2(name string, fallback *QIcon) *QIcon {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	_goptr := newQIcon(QIcon_FromTheme2(name_ms, fallback.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QIcon_HasThemeIcon(name string) bool {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	return (bool)(QIcon_HasThemeIcon(name_ms))
}

func QIcon_FromThemeWithIcon(icon ThemeIcon) *QIcon {
	_goptr := newQIcon(QIcon_FromThemeWithIcon(icon))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QIcon_FromTheme3(icon ThemeIcon, fallback *QIcon) *QIcon {
	_goptr := newQIcon(QIcon_FromTheme3(icon, fallback.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QIcon_HasThemeIconWithIcon(icon ThemeIcon) bool {
	return (bool)(QIcon_HasThemeIconWithIcon(icon))
}

func QIcon_ThemeSearchPaths() []string {
	var _ma struct_miqt_array = QIcon_ThemeSearchPaths()
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QIcon_SetThemeSearchPaths(searchpath []string) {
	searchpath_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(searchpath))))
	defer free(unsafe.Pointer(searchpath_CArray))
	for i := range searchpath {
		searchpath_i_ms := struct_miqt_string{}
		searchpath_i_ms.data = CString(searchpath[i])
		searchpath_i_ms.len = size_t(len(searchpath[i]))
		defer free(unsafe.Pointer(searchpath_i_ms.data))
		searchpath_CArray[i] = searchpath_i_ms
	}
	searchpath_ma := struct_miqt_array{len: size_t(len(searchpath)), data: unsafe.Pointer(searchpath_CArray)}
	QIcon_SetThemeSearchPaths(searchpath_ma)
}

func QIcon_FallbackSearchPaths() []string {
	var _ma struct_miqt_array = QIcon_FallbackSearchPaths()
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QIcon_SetFallbackSearchPaths(paths []string) {
	paths_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(paths))))
	defer free(unsafe.Pointer(paths_CArray))
	for i := range paths {
		paths_i_ms := struct_miqt_string{}
		paths_i_ms.data = CString(paths[i])
		paths_i_ms.len = size_t(len(paths[i]))
		defer free(unsafe.Pointer(paths_i_ms.data))
		paths_CArray[i] = paths_i_ms
	}
	paths_ma := struct_miqt_array{len: size_t(len(paths)), data: unsafe.Pointer(paths_CArray)}
	QIcon_SetFallbackSearchPaths(paths_ma)
}

func QIcon_ThemeName() string {
	var _ms struct_miqt_string = QIcon_ThemeName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIcon_SetThemeName(path string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QIcon_SetThemeName(path_ms)
}

func QIcon_FallbackThemeName() string {
	var _ms struct_miqt_string = QIcon_FallbackThemeName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIcon_SetFallbackThemeName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QIcon_SetFallbackThemeName(name_ms)
}

func (this *QIcon) DataPtr() *DataPtr {
	xxxxxxxxx
}

func (this *QIcon) Pixmap22(size *QSize, mode Mode) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap22(this.h, size.cPointer(), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap32(size *QSize, mode Mode, state State) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap32(this.h, size.cPointer(), mode, state))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap33(w int, h int, mode Mode) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap33(this.h, (int)(w), (int)(h), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap42(w int, h int, mode Mode, state State) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap42(this.h, (int)(w), (int)(h), mode, state))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap23(extent int, mode Mode) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap23(this.h, (int)(extent), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap34(extent int, mode Mode, state State) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap34(this.h, (int)(extent), mode, state))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap35(size *QSize, devicePixelRatio float64, mode Mode) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap35(this.h, size.cPointer(), (double)(devicePixelRatio), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap43(size *QSize, devicePixelRatio float64, mode Mode, state State) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap43(this.h, size.cPointer(), (double)(devicePixelRatio), mode, state))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap36(window *QWindow, size *QSize, mode Mode) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap36(this.h, window.cPointer(), size.cPointer(), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap44(window *QWindow, size *QSize, mode Mode, state State) *QPixmap {
	_goptr := newQPixmap(QIcon_Pixmap44(this.h, window.cPointer(), size.cPointer(), mode, state))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize22(size *QSize, mode Mode) *QSize {
	_goptr := newQSize(QIcon_ActualSize22(this.h, size.cPointer(), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize3(size *QSize, mode Mode, state State) *QSize {
	_goptr := newQSize(QIcon_ActualSize3(this.h, size.cPointer(), mode, state))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize32(window *QWindow, size *QSize, mode Mode) *QSize {
	_goptr := newQSize(QIcon_ActualSize32(this.h, window.cPointer(), size.cPointer(), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize4(window *QWindow, size *QSize, mode Mode, state State) *QSize {
	_goptr := newQSize(QIcon_ActualSize4(this.h, window.cPointer(), size.cPointer(), mode, state))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Paint3(painter *QPainter, rect *QRect, alignment AlignmentFlag) {
	QIcon_Paint3(this.h, painter.cPointer(), rect.cPointer(), (int)(alignment))
}

func (this *QIcon) Paint4(painter *QPainter, rect *QRect, alignment AlignmentFlag, mode Mode) {
	QIcon_Paint4(this.h, painter.cPointer(), rect.cPointer(), (int)(alignment), mode)
}

func (this *QIcon) Paint5(painter *QPainter, rect *QRect, alignment AlignmentFlag, mode Mode, state State) {
	QIcon_Paint5(this.h, painter.cPointer(), rect.cPointer(), (int)(alignment), mode, state)
}

func (this *QIcon) Paint6(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag) {
	QIcon_Paint6(this.h, painter.cPointer(), (int)(x), (int)(y), (int)(w), (int)(h), (int)(alignment))
}

func (this *QIcon) Paint7(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag, mode Mode) {
	QIcon_Paint7(this.h, painter.cPointer(), (int)(x), (int)(y), (int)(w), (int)(h), (int)(alignment), mode)
}

func (this *QIcon) Paint8(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag, mode Mode, state State) {
	QIcon_Paint8(this.h, painter.cPointer(), (int)(x), (int)(y), (int)(w), (int)(h), (int)(alignment), mode, state)
}

func (this *QIcon) AddPixmap2(pixmap *QPixmap, mode Mode) {
	QIcon_AddPixmap2(this.h, pixmap.cPointer(), mode)
}

func (this *QIcon) AddPixmap3(pixmap *QPixmap, mode Mode, state State) {
	QIcon_AddPixmap3(this.h, pixmap.cPointer(), mode, state)
}

func (this *QIcon) AddFile2(fileName string, size *QSize) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QIcon_AddFile2(this.h, fileName_ms, size.cPointer())
}

func (this *QIcon) AddFile3(fileName string, size *QSize, mode Mode) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QIcon_AddFile3(this.h, fileName_ms, size.cPointer(), mode)
}

func (this *QIcon) AddFile4(fileName string, size *QSize, mode Mode, state State) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QIcon_AddFile4(this.h, fileName_ms, size.cPointer(), mode, state)
}

func (this *QIcon) AvailableSizes1(mode Mode) []QSize {
	var _ma struct_miqt_array = QIcon_AvailableSizes1(this.h, mode)
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSize(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QIcon) AvailableSizes2(mode Mode, state State) []QSize {
	var _ma struct_miqt_array = QIcon_AvailableSizes2(this.h, mode, state)
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSize(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}
