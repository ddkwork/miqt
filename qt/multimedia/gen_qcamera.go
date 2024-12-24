package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QCamera__Error int

const (
	QCamera__NoError     QCamera__Error = 0
	QCamera__CameraError QCamera__Error = 1
)

type QCamera__FocusMode int

const (
	QCamera__FocusModeAuto       QCamera__FocusMode = 0
	QCamera__FocusModeAutoNear   QCamera__FocusMode = 1
	QCamera__FocusModeAutoFar    QCamera__FocusMode = 2
	QCamera__FocusModeHyperfocal QCamera__FocusMode = 3
	QCamera__FocusModeInfinity   QCamera__FocusMode = 4
	QCamera__FocusModeManual     QCamera__FocusMode = 5
)

type QCamera__FlashMode int

const (
	QCamera__FlashOff  QCamera__FlashMode = 0
	QCamera__FlashOn   QCamera__FlashMode = 1
	QCamera__FlashAuto QCamera__FlashMode = 2
)

type QCamera__TorchMode int

const (
	QCamera__TorchOff  QCamera__TorchMode = 0
	QCamera__TorchOn   QCamera__TorchMode = 1
	QCamera__TorchAuto QCamera__TorchMode = 2
)

type QCamera__ExposureMode int

const (
	QCamera__ExposureAuto          QCamera__ExposureMode = 0
	QCamera__ExposureManual        QCamera__ExposureMode = 1
	QCamera__ExposurePortrait      QCamera__ExposureMode = 2
	QCamera__ExposureNight         QCamera__ExposureMode = 3
	QCamera__ExposureSports        QCamera__ExposureMode = 4
	QCamera__ExposureSnow          QCamera__ExposureMode = 5
	QCamera__ExposureBeach         QCamera__ExposureMode = 6
	QCamera__ExposureAction        QCamera__ExposureMode = 7
	QCamera__ExposureLandscape     QCamera__ExposureMode = 8
	QCamera__ExposureNightPortrait QCamera__ExposureMode = 9
	QCamera__ExposureTheatre       QCamera__ExposureMode = 10
	QCamera__ExposureSunset        QCamera__ExposureMode = 11
	QCamera__ExposureSteadyPhoto   QCamera__ExposureMode = 12
	QCamera__ExposureFireworks     QCamera__ExposureMode = 13
	QCamera__ExposureParty         QCamera__ExposureMode = 14
	QCamera__ExposureCandlelight   QCamera__ExposureMode = 15
	QCamera__ExposureBarcode       QCamera__ExposureMode = 16
)

type QCamera__WhiteBalanceMode int

const (
	QCamera__WhiteBalanceAuto        QCamera__WhiteBalanceMode = 0
	QCamera__WhiteBalanceManual      QCamera__WhiteBalanceMode = 1
	QCamera__WhiteBalanceSunlight    QCamera__WhiteBalanceMode = 2
	QCamera__WhiteBalanceCloudy      QCamera__WhiteBalanceMode = 3
	QCamera__WhiteBalanceShade       QCamera__WhiteBalanceMode = 4
	QCamera__WhiteBalanceTungsten    QCamera__WhiteBalanceMode = 5
	QCamera__WhiteBalanceFluorescent QCamera__WhiteBalanceMode = 6
	QCamera__WhiteBalanceFlash       QCamera__WhiteBalanceMode = 7
	QCamera__WhiteBalanceSunset      QCamera__WhiteBalanceMode = 8
)

type QCamera__Feature int

const (
	QCamera__ColorTemperature     QCamera__Feature = 1
	QCamera__ExposureCompensation QCamera__Feature = 2
	QCamera__IsoSensitivity       QCamera__Feature = 4
	QCamera__ManualExposureTime   QCamera__Feature = 8
	QCamera__CustomFocusPoint     QCamera__Feature = 16
	QCamera__FocusDistance        QCamera__Feature = 32
)

type QCamera struct {
	h          uintptr
	isSubclass bool
}

// NewQCamera constructs a new QCamera object.
func NewQCamera() *QCamera {
	g := newQCamera(QCamera_new())
	g.isSubclass = true
	return g
}

// NewQCamera2 constructs a new QCamera object.
func NewQCamera2(cameraDevice *QCameraDevice) *QCamera {
	g := newQCamera(QCamera_new2(cameraDevice.cPointer()))
	g.isSubclass = true
	return g
}

// NewQCamera3 constructs a new QCamera object.
func NewQCamera3(position QCameraDevice__Position) *QCamera {
	g := newQCamera(QCamera_new3((int)(position)))
	g.isSubclass = true
	return g
}

// NewQCamera4 constructs a new QCamera object.
func NewQCamera4(parent *qt.QObject) *QCamera {
	g := newQCamera(QCamera_new4((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQCamera5 constructs a new QCamera object.
func NewQCamera5(cameraDevice *QCameraDevice, parent *qt.QObject) *QCamera {
	g := newQCamera(QCamera_new5(cameraDevice.cPointer(), (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQCamera6 constructs a new QCamera object.
func NewQCamera6(position QCameraDevice__Position, parent *qt.QObject) *QCamera {
	g := newQCamera(QCamera_new6((int)(position), (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QCamera) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QCamera_MetaObject(this.h)))
}

func (this *QCamera) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QCamera_Metacast(this.h, param1_Cstring))
}

func QCamera_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QCamera_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCamera) IsAvailable() bool {
	return (bool)(QCamera_IsAvailable(this.h))
}

func (this *QCamera) IsActive() bool {
	return (bool)(QCamera_IsActive(this.h))
}

func (this *QCamera) CaptureSession() *QMediaCaptureSession {
	return newQMediaCaptureSession(QCamera_CaptureSession(this.h))
}

func (this *QCamera) CameraDevice() *QCameraDevice {
	_goptr := newQCameraDevice(QCamera_CameraDevice(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCamera) SetCameraDevice(cameraDevice *QCameraDevice) {
	QCamera_SetCameraDevice(this.h, cameraDevice.cPointer())
}

func (this *QCamera) CameraFormat() *QCameraFormat {
	_goptr := newQCameraFormat(QCamera_CameraFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCamera) SetCameraFormat(format *QCameraFormat) {
	QCamera_SetCameraFormat(this.h, format.cPointer())
}

func (this *QCamera) Error() Error {
	xxxxxxxxx
}

func (this *QCamera) ErrorString() string {
	var _ms struct_miqt_string = QCamera_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCamera) SupportedFeatures() Features {
	xxxxxxxxx
}

func (this *QCamera) FocusMode() FocusMode {
	xxxxxxxxx
}

func (this *QCamera) SetFocusMode(mode FocusMode) {
	QCamera_SetFocusMode(this.h, mode)
}

func (this *QCamera) IsFocusModeSupported(mode FocusMode) bool {
	return (bool)(QCamera_IsFocusModeSupported(this.h, mode))
}

func (this *QCamera) FocusPoint() *qt.QPointF {
	_goptr := qt.UnsafeNewQPointF(unsafe.Pointer(QCamera_FocusPoint(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCamera) CustomFocusPoint() *qt.QPointF {
	_goptr := qt.UnsafeNewQPointF(unsafe.Pointer(QCamera_CustomFocusPoint(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCamera) SetCustomFocusPoint(point *qt.QPointF) {
	QCamera_SetCustomFocusPoint(this.h, (*QPointF)(point.UnsafePointer()))
}

func (this *QCamera) SetFocusDistance(d float32) {
	QCamera_SetFocusDistance(this.h, (float)(d))
}

func (this *QCamera) FocusDistance() float32 {
	return (float32)(QCamera_FocusDistance(this.h))
}

func (this *QCamera) MinimumZoomFactor() float32 {
	return (float32)(QCamera_MinimumZoomFactor(this.h))
}

func (this *QCamera) MaximumZoomFactor() float32 {
	return (float32)(QCamera_MaximumZoomFactor(this.h))
}

func (this *QCamera) ZoomFactor() float32 {
	return (float32)(QCamera_ZoomFactor(this.h))
}

func (this *QCamera) SetZoomFactor(factor float32) {
	QCamera_SetZoomFactor(this.h, (float)(factor))
}

func (this *QCamera) FlashMode() FlashMode {
	xxxxxxxxx
}

func (this *QCamera) IsFlashModeSupported(mode FlashMode) bool {
	return (bool)(QCamera_IsFlashModeSupported(this.h, mode))
}

func (this *QCamera) IsFlashReady() bool {
	return (bool)(QCamera_IsFlashReady(this.h))
}

func (this *QCamera) TorchMode() TorchMode {
	xxxxxxxxx
}

func (this *QCamera) IsTorchModeSupported(mode TorchMode) bool {
	return (bool)(QCamera_IsTorchModeSupported(this.h, mode))
}

func (this *QCamera) ExposureMode() ExposureMode {
	xxxxxxxxx
}

func (this *QCamera) IsExposureModeSupported(mode ExposureMode) bool {
	return (bool)(QCamera_IsExposureModeSupported(this.h, mode))
}

func (this *QCamera) ExposureCompensation() float32 {
	return (float32)(QCamera_ExposureCompensation(this.h))
}

func (this *QCamera) IsoSensitivity() int {
	return (int)(QCamera_IsoSensitivity(this.h))
}

func (this *QCamera) ManualIsoSensitivity() int {
	return (int)(QCamera_ManualIsoSensitivity(this.h))
}

func (this *QCamera) ExposureTime() float32 {
	return (float32)(QCamera_ExposureTime(this.h))
}

func (this *QCamera) ManualExposureTime() float32 {
	return (float32)(QCamera_ManualExposureTime(this.h))
}

func (this *QCamera) MinimumIsoSensitivity() int {
	return (int)(QCamera_MinimumIsoSensitivity(this.h))
}

func (this *QCamera) MaximumIsoSensitivity() int {
	return (int)(QCamera_MaximumIsoSensitivity(this.h))
}

func (this *QCamera) MinimumExposureTime() float32 {
	return (float32)(QCamera_MinimumExposureTime(this.h))
}

func (this *QCamera) MaximumExposureTime() float32 {
	return (float32)(QCamera_MaximumExposureTime(this.h))
}

func (this *QCamera) WhiteBalanceMode() WhiteBalanceMode {
	xxxxxxxxx
}

func (this *QCamera) IsWhiteBalanceModeSupported(mode WhiteBalanceMode) bool {
	return (bool)(QCamera_IsWhiteBalanceModeSupported(this.h, mode))
}

func (this *QCamera) ColorTemperature() int {
	return (int)(QCamera_ColorTemperature(this.h))
}

func (this *QCamera) SetActive(active bool) {
	QCamera_SetActive(this.h, (bool)(active))
}

func (this *QCamera) Start() {
	QCamera_Start(this.h)
}

func (this *QCamera) Stop() {
	QCamera_Stop(this.h)
}

func (this *QCamera) ZoomTo(zoom float32, rate float32) {
	QCamera_ZoomTo(this.h, (float)(zoom), (float)(rate))
}

func (this *QCamera) SetFlashMode(mode FlashMode) {
	QCamera_SetFlashMode(this.h, mode)
}

func (this *QCamera) SetTorchMode(mode TorchMode) {
	QCamera_SetTorchMode(this.h, mode)
}

func (this *QCamera) SetExposureMode(mode ExposureMode) {
	QCamera_SetExposureMode(this.h, mode)
}

func (this *QCamera) SetExposureCompensation(ev float32) {
	QCamera_SetExposureCompensation(this.h, (float)(ev))
}

func (this *QCamera) SetManualIsoSensitivity(iso int) {
	QCamera_SetManualIsoSensitivity(this.h, (int)(iso))
}

func (this *QCamera) SetAutoIsoSensitivity() {
	QCamera_SetAutoIsoSensitivity(this.h)
}

func (this *QCamera) SetManualExposureTime(seconds float32) {
	QCamera_SetManualExposureTime(this.h, (float)(seconds))
}

func (this *QCamera) SetAutoExposureTime() {
	QCamera_SetAutoExposureTime(this.h)
}

func (this *QCamera) SetWhiteBalanceMode(mode WhiteBalanceMode) {
	QCamera_SetWhiteBalanceMode(this.h, mode)
}

func (this *QCamera) SetColorTemperature(colorTemperature int) {
	QCamera_SetColorTemperature(this.h, (int)(colorTemperature))
}

func (this *QCamera) ActiveChanged(param1 bool) {
	QCamera_ActiveChanged(this.h, (bool)(param1))
}

func (this *QCamera) OnActiveChanged(slot func(param1 bool)) {
	QCamera_connect_ActiveChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ActiveChanged
func miqt_exec_callback_QCamera_ActiveChanged(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QCamera) ErrorChanged() {
	QCamera_ErrorChanged(this.h)
}

func (this *QCamera) OnErrorChanged(slot func()) {
	QCamera_connect_ErrorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ErrorChanged
func miqt_exec_callback_QCamera_ErrorChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) ErrorOccurred(error QCamera__Error, errorString string) {
	errorString_ms := struct_miqt_string{}
	errorString_ms.data = CString(errorString)
	errorString_ms.len = size_t(len(errorString))
	defer free(unsafe.Pointer(errorString_ms.data))
	QCamera_ErrorOccurred(this.h, (int)(error), errorString_ms)
}

func (this *QCamera) OnErrorOccurred(slot func(error QCamera__Error, errorString string)) {
	QCamera_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ErrorOccurred
func miqt_exec_callback_QCamera_ErrorOccurred(cb intptr_t, error int, errorString struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QCamera__Error, errorString string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QCamera__Error)(error)

	var errorString_ms struct_miqt_string = errorString
	errorString_ret := GoStringN(errorString_ms.data, int(int64(errorString_ms.len)))
	free(unsafe.Pointer(errorString_ms.data))
	slotval2 := errorString_ret

	gofunc(slotval1, slotval2)
}

func (this *QCamera) CameraDeviceChanged() {
	QCamera_CameraDeviceChanged(this.h)
}

func (this *QCamera) OnCameraDeviceChanged(slot func()) {
	QCamera_connect_CameraDeviceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_CameraDeviceChanged
func miqt_exec_callback_QCamera_CameraDeviceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) CameraFormatChanged() {
	QCamera_CameraFormatChanged(this.h)
}

func (this *QCamera) OnCameraFormatChanged(slot func()) {
	QCamera_connect_CameraFormatChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_CameraFormatChanged
func miqt_exec_callback_QCamera_CameraFormatChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) SupportedFeaturesChanged() {
	QCamera_SupportedFeaturesChanged(this.h)
}

func (this *QCamera) OnSupportedFeaturesChanged(slot func()) {
	QCamera_connect_SupportedFeaturesChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_SupportedFeaturesChanged
func miqt_exec_callback_QCamera_SupportedFeaturesChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) FocusModeChanged() {
	QCamera_FocusModeChanged(this.h)
}

func (this *QCamera) OnFocusModeChanged(slot func()) {
	QCamera_connect_FocusModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_FocusModeChanged
func miqt_exec_callback_QCamera_FocusModeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) ZoomFactorChanged(param1 float32) {
	QCamera_ZoomFactorChanged(this.h, (float)(param1))
}

func (this *QCamera) OnZoomFactorChanged(slot func(param1 float32)) {
	QCamera_connect_ZoomFactorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ZoomFactorChanged
func miqt_exec_callback_QCamera_ZoomFactorChanged(cb intptr_t, param1 float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(param1)

	gofunc(slotval1)
}

func (this *QCamera) MinimumZoomFactorChanged(param1 float32) {
	QCamera_MinimumZoomFactorChanged(this.h, (float)(param1))
}

func (this *QCamera) OnMinimumZoomFactorChanged(slot func(param1 float32)) {
	QCamera_connect_MinimumZoomFactorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_MinimumZoomFactorChanged
func miqt_exec_callback_QCamera_MinimumZoomFactorChanged(cb intptr_t, param1 float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(param1)

	gofunc(slotval1)
}

func (this *QCamera) MaximumZoomFactorChanged(param1 float32) {
	QCamera_MaximumZoomFactorChanged(this.h, (float)(param1))
}

func (this *QCamera) OnMaximumZoomFactorChanged(slot func(param1 float32)) {
	QCamera_connect_MaximumZoomFactorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_MaximumZoomFactorChanged
func miqt_exec_callback_QCamera_MaximumZoomFactorChanged(cb intptr_t, param1 float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(param1)

	gofunc(slotval1)
}

func (this *QCamera) FocusDistanceChanged(param1 float32) {
	QCamera_FocusDistanceChanged(this.h, (float)(param1))
}

func (this *QCamera) OnFocusDistanceChanged(slot func(param1 float32)) {
	QCamera_connect_FocusDistanceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_FocusDistanceChanged
func miqt_exec_callback_QCamera_FocusDistanceChanged(cb intptr_t, param1 float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(param1)

	gofunc(slotval1)
}

func (this *QCamera) FocusPointChanged() {
	QCamera_FocusPointChanged(this.h)
}

func (this *QCamera) OnFocusPointChanged(slot func()) {
	QCamera_connect_FocusPointChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_FocusPointChanged
func miqt_exec_callback_QCamera_FocusPointChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) CustomFocusPointChanged() {
	QCamera_CustomFocusPointChanged(this.h)
}

func (this *QCamera) OnCustomFocusPointChanged(slot func()) {
	QCamera_connect_CustomFocusPointChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_CustomFocusPointChanged
func miqt_exec_callback_QCamera_CustomFocusPointChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) FlashReady(param1 bool) {
	QCamera_FlashReady(this.h, (bool)(param1))
}

func (this *QCamera) OnFlashReady(slot func(param1 bool)) {
	QCamera_connect_FlashReady(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_FlashReady
func miqt_exec_callback_QCamera_FlashReady(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QCamera) FlashModeChanged() {
	QCamera_FlashModeChanged(this.h)
}

func (this *QCamera) OnFlashModeChanged(slot func()) {
	QCamera_connect_FlashModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_FlashModeChanged
func miqt_exec_callback_QCamera_FlashModeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) TorchModeChanged() {
	QCamera_TorchModeChanged(this.h)
}

func (this *QCamera) OnTorchModeChanged(slot func()) {
	QCamera_connect_TorchModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_TorchModeChanged
func miqt_exec_callback_QCamera_TorchModeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) ExposureTimeChanged(speed float32) {
	QCamera_ExposureTimeChanged(this.h, (float)(speed))
}

func (this *QCamera) OnExposureTimeChanged(slot func(speed float32)) {
	QCamera_connect_ExposureTimeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ExposureTimeChanged
func miqt_exec_callback_QCamera_ExposureTimeChanged(cb intptr_t, speed float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(speed float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(speed)

	gofunc(slotval1)
}

func (this *QCamera) ManualExposureTimeChanged(speed float32) {
	QCamera_ManualExposureTimeChanged(this.h, (float)(speed))
}

func (this *QCamera) OnManualExposureTimeChanged(slot func(speed float32)) {
	QCamera_connect_ManualExposureTimeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ManualExposureTimeChanged
func miqt_exec_callback_QCamera_ManualExposureTimeChanged(cb intptr_t, speed float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(speed float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(speed)

	gofunc(slotval1)
}

func (this *QCamera) IsoSensitivityChanged(param1 int) {
	QCamera_IsoSensitivityChanged(this.h, (int)(param1))
}

func (this *QCamera) OnIsoSensitivityChanged(slot func(param1 int)) {
	QCamera_connect_IsoSensitivityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_IsoSensitivityChanged
func miqt_exec_callback_QCamera_IsoSensitivityChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QCamera) ManualIsoSensitivityChanged(param1 int) {
	QCamera_ManualIsoSensitivityChanged(this.h, (int)(param1))
}

func (this *QCamera) OnManualIsoSensitivityChanged(slot func(param1 int)) {
	QCamera_connect_ManualIsoSensitivityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ManualIsoSensitivityChanged
func miqt_exec_callback_QCamera_ManualIsoSensitivityChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QCamera) ExposureCompensationChanged(param1 float32) {
	QCamera_ExposureCompensationChanged(this.h, (float)(param1))
}

func (this *QCamera) OnExposureCompensationChanged(slot func(param1 float32)) {
	QCamera_connect_ExposureCompensationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ExposureCompensationChanged
func miqt_exec_callback_QCamera_ExposureCompensationChanged(cb intptr_t, param1 float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(param1)

	gofunc(slotval1)
}

func (this *QCamera) ExposureModeChanged() {
	QCamera_ExposureModeChanged(this.h)
}

func (this *QCamera) OnExposureModeChanged(slot func()) {
	QCamera_connect_ExposureModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ExposureModeChanged
func miqt_exec_callback_QCamera_ExposureModeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) WhiteBalanceModeChanged() {
	QCamera_WhiteBalanceModeChanged(this.h)
}

func (this *QCamera) OnWhiteBalanceModeChanged(slot func()) {
	QCamera_connect_WhiteBalanceModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_WhiteBalanceModeChanged
func miqt_exec_callback_QCamera_WhiteBalanceModeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) ColorTemperatureChanged() {
	QCamera_ColorTemperatureChanged(this.h)
}

func (this *QCamera) OnColorTemperatureChanged(slot func()) {
	QCamera_connect_ColorTemperatureChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ColorTemperatureChanged
func miqt_exec_callback_QCamera_ColorTemperatureChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) BrightnessChanged() {
	QCamera_BrightnessChanged(this.h)
}

func (this *QCamera) OnBrightnessChanged(slot func()) {
	QCamera_connect_BrightnessChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_BrightnessChanged
func miqt_exec_callback_QCamera_BrightnessChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) ContrastChanged() {
	QCamera_ContrastChanged(this.h)
}

func (this *QCamera) OnContrastChanged(slot func()) {
	QCamera_connect_ContrastChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_ContrastChanged
func miqt_exec_callback_QCamera_ContrastChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) SaturationChanged() {
	QCamera_SaturationChanged(this.h)
}

func (this *QCamera) OnSaturationChanged(slot func()) {
	QCamera_connect_SaturationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_SaturationChanged
func miqt_exec_callback_QCamera_SaturationChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCamera) HueChanged() {
	QCamera_HueChanged(this.h)
}

func (this *QCamera) OnHueChanged(slot func()) {
	QCamera_connect_HueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_HueChanged
func miqt_exec_callback_QCamera_HueChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QCamera_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCamera_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCamera_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCamera_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCamera) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QCamera_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QCamera) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCamera_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_MetaObject
func miqt_exec_callback_QCamera_MetaObject(self QCamera, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCamera{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QCamera) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QCamera_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QCamera) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCamera_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCamera_Metacast
func miqt_exec_callback_QCamera_Metacast(self QCamera, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QCamera{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
