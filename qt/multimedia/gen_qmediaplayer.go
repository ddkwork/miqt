package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QMediaPlayer__PlaybackState int

const (
	QMediaPlayer__StoppedState QMediaPlayer__PlaybackState = 0
	QMediaPlayer__PlayingState QMediaPlayer__PlaybackState = 1
	QMediaPlayer__PausedState  QMediaPlayer__PlaybackState = 2
)

type QMediaPlayer__MediaStatus int

const (
	QMediaPlayer__NoMedia        QMediaPlayer__MediaStatus = 0
	QMediaPlayer__LoadingMedia   QMediaPlayer__MediaStatus = 1
	QMediaPlayer__LoadedMedia    QMediaPlayer__MediaStatus = 2
	QMediaPlayer__StalledMedia   QMediaPlayer__MediaStatus = 3
	QMediaPlayer__BufferingMedia QMediaPlayer__MediaStatus = 4
	QMediaPlayer__BufferedMedia  QMediaPlayer__MediaStatus = 5
	QMediaPlayer__EndOfMedia     QMediaPlayer__MediaStatus = 6
	QMediaPlayer__InvalidMedia   QMediaPlayer__MediaStatus = 7
)

type QMediaPlayer__Error int

const (
	QMediaPlayer__NoError           QMediaPlayer__Error = 0
	QMediaPlayer__ResourceError     QMediaPlayer__Error = 1
	QMediaPlayer__FormatError       QMediaPlayer__Error = 2
	QMediaPlayer__NetworkError      QMediaPlayer__Error = 3
	QMediaPlayer__AccessDeniedError QMediaPlayer__Error = 4
)

type QMediaPlayer__Loops int

const (
	QMediaPlayer__Infinite QMediaPlayer__Loops = -1
	QMediaPlayer__Once     QMediaPlayer__Loops = 1
)

type QMediaPlayer struct {
	h          uintptr
	isSubclass bool
}

// NewQMediaPlayer constructs a new QMediaPlayer object.
func NewQMediaPlayer() *QMediaPlayer {
	g := newQMediaPlayer(QMediaPlayer_new())
	g.isSubclass = true
	return g
}

// NewQMediaPlayer2 constructs a new QMediaPlayer object.
func NewQMediaPlayer2(parent *qt.QObject) *QMediaPlayer {
	g := newQMediaPlayer(QMediaPlayer_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QMediaPlayer) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QMediaPlayer_MetaObject(this.h)))
}

func (this *QMediaPlayer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMediaPlayer_Metacast(this.h, param1_Cstring))
}

func QMediaPlayer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMediaPlayer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaPlayer) AudioTracks() []QMediaMetaData {
	var _ma struct_miqt_array = QMediaPlayer_AudioTracks(this.h)
	_ret := make([]QMediaMetaData, int(_ma.len))
	_outCast := (*[0xffff]*QMediaMetaData)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQMediaMetaData(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QMediaPlayer) VideoTracks() []QMediaMetaData {
	var _ma struct_miqt_array = QMediaPlayer_VideoTracks(this.h)
	_ret := make([]QMediaMetaData, int(_ma.len))
	_outCast := (*[0xffff]*QMediaMetaData)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQMediaMetaData(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QMediaPlayer) SubtitleTracks() []QMediaMetaData {
	var _ma struct_miqt_array = QMediaPlayer_SubtitleTracks(this.h)
	_ret := make([]QMediaMetaData, int(_ma.len))
	_outCast := (*[0xffff]*QMediaMetaData)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQMediaMetaData(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QMediaPlayer) ActiveAudioTrack() int {
	return (int)(QMediaPlayer_ActiveAudioTrack(this.h))
}

func (this *QMediaPlayer) ActiveVideoTrack() int {
	return (int)(QMediaPlayer_ActiveVideoTrack(this.h))
}

func (this *QMediaPlayer) ActiveSubtitleTrack() int {
	return (int)(QMediaPlayer_ActiveSubtitleTrack(this.h))
}

func (this *QMediaPlayer) SetActiveAudioTrack(index int) {
	QMediaPlayer_SetActiveAudioTrack(this.h, (int)(index))
}

func (this *QMediaPlayer) SetActiveVideoTrack(index int) {
	QMediaPlayer_SetActiveVideoTrack(this.h, (int)(index))
}

func (this *QMediaPlayer) SetActiveSubtitleTrack(index int) {
	QMediaPlayer_SetActiveSubtitleTrack(this.h, (int)(index))
}

func (this *QMediaPlayer) SetAudioBufferOutput(output *QAudioBufferOutput) {
	QMediaPlayer_SetAudioBufferOutput(this.h, output.cPointer())
}

func (this *QMediaPlayer) AudioBufferOutput() *QAudioBufferOutput {
	return newQAudioBufferOutput(QMediaPlayer_AudioBufferOutput(this.h))
}

func (this *QMediaPlayer) SetAudioOutput(output *QAudioOutput) {
	QMediaPlayer_SetAudioOutput(this.h, output.cPointer())
}

func (this *QMediaPlayer) AudioOutput() *QAudioOutput {
	return newQAudioOutput(QMediaPlayer_AudioOutput(this.h))
}

func (this *QMediaPlayer) SetVideoOutput(videoOutput *qt.QObject) {
	QMediaPlayer_SetVideoOutput(this.h, (*QObject)(videoOutput.UnsafePointer()))
}

func (this *QMediaPlayer) VideoOutput() *qt.QObject {
	return qt.UnsafeNewQObject(unsafe.Pointer(QMediaPlayer_VideoOutput(this.h)))
}

func (this *QMediaPlayer) SetVideoSink(sink *QVideoSink) {
	QMediaPlayer_SetVideoSink(this.h, sink.cPointer())
}

func (this *QMediaPlayer) VideoSink() *QVideoSink {
	return newQVideoSink(QMediaPlayer_VideoSink(this.h))
}

func (this *QMediaPlayer) Source() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QMediaPlayer_Source(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaPlayer) SourceDevice() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QMediaPlayer_SourceDevice(this.h)))
}

func (this *QMediaPlayer) PlaybackState() PlaybackState {
	xxxxxxxxx
}

func (this *QMediaPlayer) MediaStatus() MediaStatus {
	xxxxxxxxx
}

func (this *QMediaPlayer) Duration() int64 {
	return (int64)(QMediaPlayer_Duration(this.h))
}

func (this *QMediaPlayer) Position() int64 {
	return (int64)(QMediaPlayer_Position(this.h))
}

func (this *QMediaPlayer) HasAudio() bool {
	return (bool)(QMediaPlayer_HasAudio(this.h))
}

func (this *QMediaPlayer) HasVideo() bool {
	return (bool)(QMediaPlayer_HasVideo(this.h))
}

func (this *QMediaPlayer) BufferProgress() float32 {
	return (float32)(QMediaPlayer_BufferProgress(this.h))
}

func (this *QMediaPlayer) BufferedTimeRange() *QMediaTimeRange {
	_goptr := newQMediaTimeRange(QMediaPlayer_BufferedTimeRange(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaPlayer) IsSeekable() bool {
	return (bool)(QMediaPlayer_IsSeekable(this.h))
}

func (this *QMediaPlayer) PlaybackRate() float64 {
	return (float64)(QMediaPlayer_PlaybackRate(this.h))
}

func (this *QMediaPlayer) IsPlaying() bool {
	return (bool)(QMediaPlayer_IsPlaying(this.h))
}

func (this *QMediaPlayer) Loops() int {
	return (int)(QMediaPlayer_Loops(this.h))
}

func (this *QMediaPlayer) SetLoops(loops int) {
	QMediaPlayer_SetLoops(this.h, (int)(loops))
}

func (this *QMediaPlayer) Error() Error {
	xxxxxxxxx
}

func (this *QMediaPlayer) ErrorString() string {
	var _ms struct_miqt_string = QMediaPlayer_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaPlayer) IsAvailable() bool {
	return (bool)(QMediaPlayer_IsAvailable(this.h))
}

func (this *QMediaPlayer) MetaData() *QMediaMetaData {
	_goptr := newQMediaMetaData(QMediaPlayer_MetaData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaPlayer) Play() {
	QMediaPlayer_Play(this.h)
}

func (this *QMediaPlayer) Pause() {
	QMediaPlayer_Pause(this.h)
}

func (this *QMediaPlayer) Stop() {
	QMediaPlayer_Stop(this.h)
}

func (this *QMediaPlayer) SetPosition(position int64) {
	QMediaPlayer_SetPosition(this.h, (longlong)(position))
}

func (this *QMediaPlayer) SetPlaybackRate(rate float64) {
	QMediaPlayer_SetPlaybackRate(this.h, (double)(rate))
}

func (this *QMediaPlayer) SetSource(source *qt.QUrl) {
	QMediaPlayer_SetSource(this.h, (*QUrl)(source.UnsafePointer()))
}

func (this *QMediaPlayer) SetSourceDevice(device *qt.QIODevice) {
	QMediaPlayer_SetSourceDevice(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QMediaPlayer) SourceChanged(media *qt.QUrl) {
	QMediaPlayer_SourceChanged(this.h, (*QUrl)(media.UnsafePointer()))
}

func (this *QMediaPlayer) OnSourceChanged(slot func(media *qt.QUrl)) {
	QMediaPlayer_connect_SourceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_SourceChanged
func miqt_exec_callback_QMediaPlayer_SourceChanged(cb intptr_t, media *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(media *qt.QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(media))

	gofunc(slotval1)
}

func (this *QMediaPlayer) PlaybackStateChanged(newState QMediaPlayer__PlaybackState) {
	QMediaPlayer_PlaybackStateChanged(this.h, (int)(newState))
}

func (this *QMediaPlayer) OnPlaybackStateChanged(slot func(newState QMediaPlayer__PlaybackState)) {
	QMediaPlayer_connect_PlaybackStateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_PlaybackStateChanged
func miqt_exec_callback_QMediaPlayer_PlaybackStateChanged(cb intptr_t, newState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newState QMediaPlayer__PlaybackState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QMediaPlayer__PlaybackState)(newState)

	gofunc(slotval1)
}

func (this *QMediaPlayer) MediaStatusChanged(status QMediaPlayer__MediaStatus) {
	QMediaPlayer_MediaStatusChanged(this.h, (int)(status))
}

func (this *QMediaPlayer) OnMediaStatusChanged(slot func(status QMediaPlayer__MediaStatus)) {
	QMediaPlayer_connect_MediaStatusChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_MediaStatusChanged
func miqt_exec_callback_QMediaPlayer_MediaStatusChanged(cb intptr_t, status int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(status QMediaPlayer__MediaStatus))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QMediaPlayer__MediaStatus)(status)

	gofunc(slotval1)
}

func (this *QMediaPlayer) DurationChanged(duration int64) {
	QMediaPlayer_DurationChanged(this.h, (longlong)(duration))
}

func (this *QMediaPlayer) OnDurationChanged(slot func(duration int64)) {
	QMediaPlayer_connect_DurationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_DurationChanged
func miqt_exec_callback_QMediaPlayer_DurationChanged(cb intptr_t, duration longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(duration int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(duration)

	gofunc(slotval1)
}

func (this *QMediaPlayer) PositionChanged(position int64) {
	QMediaPlayer_PositionChanged(this.h, (longlong)(position))
}

func (this *QMediaPlayer) OnPositionChanged(slot func(position int64)) {
	QMediaPlayer_connect_PositionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_PositionChanged
func miqt_exec_callback_QMediaPlayer_PositionChanged(cb intptr_t, position longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(position int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(position)

	gofunc(slotval1)
}

func (this *QMediaPlayer) HasAudioChanged(available bool) {
	QMediaPlayer_HasAudioChanged(this.h, (bool)(available))
}

func (this *QMediaPlayer) OnHasAudioChanged(slot func(available bool)) {
	QMediaPlayer_connect_HasAudioChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_HasAudioChanged
func miqt_exec_callback_QMediaPlayer_HasAudioChanged(cb intptr_t, available bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(available bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(available)

	gofunc(slotval1)
}

func (this *QMediaPlayer) HasVideoChanged(videoAvailable bool) {
	QMediaPlayer_HasVideoChanged(this.h, (bool)(videoAvailable))
}

func (this *QMediaPlayer) OnHasVideoChanged(slot func(videoAvailable bool)) {
	QMediaPlayer_connect_HasVideoChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_HasVideoChanged
func miqt_exec_callback_QMediaPlayer_HasVideoChanged(cb intptr_t, videoAvailable bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(videoAvailable bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(videoAvailable)

	gofunc(slotval1)
}

func (this *QMediaPlayer) BufferProgressChanged(progress float32) {
	QMediaPlayer_BufferProgressChanged(this.h, (float)(progress))
}

func (this *QMediaPlayer) OnBufferProgressChanged(slot func(progress float32)) {
	QMediaPlayer_connect_BufferProgressChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_BufferProgressChanged
func miqt_exec_callback_QMediaPlayer_BufferProgressChanged(cb intptr_t, progress float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(progress float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(progress)

	gofunc(slotval1)
}

func (this *QMediaPlayer) SeekableChanged(seekable bool) {
	QMediaPlayer_SeekableChanged(this.h, (bool)(seekable))
}

func (this *QMediaPlayer) OnSeekableChanged(slot func(seekable bool)) {
	QMediaPlayer_connect_SeekableChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_SeekableChanged
func miqt_exec_callback_QMediaPlayer_SeekableChanged(cb intptr_t, seekable bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(seekable bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(seekable)

	gofunc(slotval1)
}

func (this *QMediaPlayer) PlayingChanged(playing bool) {
	QMediaPlayer_PlayingChanged(this.h, (bool)(playing))
}

func (this *QMediaPlayer) OnPlayingChanged(slot func(playing bool)) {
	QMediaPlayer_connect_PlayingChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_PlayingChanged
func miqt_exec_callback_QMediaPlayer_PlayingChanged(cb intptr_t, playing bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(playing bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(playing)

	gofunc(slotval1)
}

func (this *QMediaPlayer) PlaybackRateChanged(rate float64) {
	QMediaPlayer_PlaybackRateChanged(this.h, (double)(rate))
}

func (this *QMediaPlayer) OnPlaybackRateChanged(slot func(rate float64)) {
	QMediaPlayer_connect_PlaybackRateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_PlaybackRateChanged
func miqt_exec_callback_QMediaPlayer_PlaybackRateChanged(cb intptr_t, rate double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(rate float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(rate)

	gofunc(slotval1)
}

func (this *QMediaPlayer) LoopsChanged() {
	QMediaPlayer_LoopsChanged(this.h)
}

func (this *QMediaPlayer) OnLoopsChanged(slot func()) {
	QMediaPlayer_connect_LoopsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_LoopsChanged
func miqt_exec_callback_QMediaPlayer_LoopsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaPlayer) MetaDataChanged() {
	QMediaPlayer_MetaDataChanged(this.h)
}

func (this *QMediaPlayer) OnMetaDataChanged(slot func()) {
	QMediaPlayer_connect_MetaDataChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_MetaDataChanged
func miqt_exec_callback_QMediaPlayer_MetaDataChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaPlayer) VideoOutputChanged() {
	QMediaPlayer_VideoOutputChanged(this.h)
}

func (this *QMediaPlayer) OnVideoOutputChanged(slot func()) {
	QMediaPlayer_connect_VideoOutputChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_VideoOutputChanged
func miqt_exec_callback_QMediaPlayer_VideoOutputChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaPlayer) AudioOutputChanged() {
	QMediaPlayer_AudioOutputChanged(this.h)
}

func (this *QMediaPlayer) OnAudioOutputChanged(slot func()) {
	QMediaPlayer_connect_AudioOutputChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_AudioOutputChanged
func miqt_exec_callback_QMediaPlayer_AudioOutputChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaPlayer) AudioBufferOutputChanged() {
	QMediaPlayer_AudioBufferOutputChanged(this.h)
}

func (this *QMediaPlayer) OnAudioBufferOutputChanged(slot func()) {
	QMediaPlayer_connect_AudioBufferOutputChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_AudioBufferOutputChanged
func miqt_exec_callback_QMediaPlayer_AudioBufferOutputChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaPlayer) TracksChanged() {
	QMediaPlayer_TracksChanged(this.h)
}

func (this *QMediaPlayer) OnTracksChanged(slot func()) {
	QMediaPlayer_connect_TracksChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_TracksChanged
func miqt_exec_callback_QMediaPlayer_TracksChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaPlayer) ActiveTracksChanged() {
	QMediaPlayer_ActiveTracksChanged(this.h)
}

func (this *QMediaPlayer) OnActiveTracksChanged(slot func()) {
	QMediaPlayer_connect_ActiveTracksChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_ActiveTracksChanged
func miqt_exec_callback_QMediaPlayer_ActiveTracksChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaPlayer) ErrorChanged() {
	QMediaPlayer_ErrorChanged(this.h)
}

func (this *QMediaPlayer) OnErrorChanged(slot func()) {
	QMediaPlayer_connect_ErrorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_ErrorChanged
func miqt_exec_callback_QMediaPlayer_ErrorChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaPlayer) ErrorOccurred(error QMediaPlayer__Error, errorString string) {
	errorString_ms := struct_miqt_string{}
	errorString_ms.data = CString(errorString)
	errorString_ms.len = size_t(len(errorString))
	defer free(unsafe.Pointer(errorString_ms.data))
	QMediaPlayer_ErrorOccurred(this.h, (int)(error), errorString_ms)
}

func (this *QMediaPlayer) OnErrorOccurred(slot func(error QMediaPlayer__Error, errorString string)) {
	QMediaPlayer_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_ErrorOccurred
func miqt_exec_callback_QMediaPlayer_ErrorOccurred(cb intptr_t, error int, errorString struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QMediaPlayer__Error, errorString string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QMediaPlayer__Error)(error)

	var errorString_ms struct_miqt_string = errorString
	errorString_ret := GoStringN(errorString_ms.data, int(int64(errorString_ms.len)))
	free(unsafe.Pointer(errorString_ms.data))
	slotval2 := errorString_ret

	gofunc(slotval1, slotval2)
}

func QMediaPlayer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMediaPlayer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaPlayer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMediaPlayer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaPlayer) SetSourceDevice2(device *qt.QIODevice, sourceUrl *qt.QUrl) {
	QMediaPlayer_SetSourceDevice2(this.h, (*QIODevice)(device.UnsafePointer()), (*QUrl)(sourceUrl.UnsafePointer()))
}

func (this *QMediaPlayer) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QMediaPlayer_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QMediaPlayer) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMediaPlayer_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_MetaObject
func miqt_exec_callback_QMediaPlayer_MetaObject(self QMediaPlayer, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMediaPlayer{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QMediaPlayer) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMediaPlayer_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMediaPlayer) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMediaPlayer_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaPlayer_Metacast
func miqt_exec_callback_QMediaPlayer_Metacast(self QMediaPlayer, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QMediaPlayer{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
