package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QNetworkInformation__Reachability int

const (
	QNetworkInformation__Reachability__Unknown      QNetworkInformation__Reachability = 0
	QNetworkInformation__Reachability__Disconnected QNetworkInformation__Reachability = 1
	QNetworkInformation__Reachability__Local        QNetworkInformation__Reachability = 2
	QNetworkInformation__Reachability__Site         QNetworkInformation__Reachability = 3
	QNetworkInformation__Reachability__Online       QNetworkInformation__Reachability = 4
)

type QNetworkInformation__TransportMedium int

const (
	QNetworkInformation__TransportMedium__Unknown   QNetworkInformation__TransportMedium = 0
	QNetworkInformation__TransportMedium__Ethernet  QNetworkInformation__TransportMedium = 1
	QNetworkInformation__TransportMedium__Cellular  QNetworkInformation__TransportMedium = 2
	QNetworkInformation__TransportMedium__WiFi      QNetworkInformation__TransportMedium = 3
	QNetworkInformation__TransportMedium__Bluetooth QNetworkInformation__TransportMedium = 4
)

type QNetworkInformation__Feature int

const (
	QNetworkInformation__Feature__Reachability    QNetworkInformation__Feature = 1
	QNetworkInformation__Feature__CaptivePortal   QNetworkInformation__Feature = 2
	QNetworkInformation__Feature__TransportMedium QNetworkInformation__Feature = 4
	QNetworkInformation__Feature__Metered         QNetworkInformation__Feature = 8
)

type QNetworkInformation struct {
	h          uintptr
	isSubclass bool
}

func (this *QNetworkInformation) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QNetworkInformation_MetaObject(this.h)))
}

func (this *QNetworkInformation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QNetworkInformation_Metacast(this.h, param1_Cstring))
}

func QNetworkInformation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QNetworkInformation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkInformation) Reachability() Reachability {
	xxxxxxxxx
}

func (this *QNetworkInformation) IsBehindCaptivePortal() bool {
	return (bool)(QNetworkInformation_IsBehindCaptivePortal(this.h))
}

func (this *QNetworkInformation) TransportMedium() TransportMedium {
	xxxxxxxxx
}

func (this *QNetworkInformation) IsMetered() bool {
	return (bool)(QNetworkInformation_IsMetered(this.h))
}

func (this *QNetworkInformation) BackendName() string {
	var _ms struct_miqt_string = QNetworkInformation_BackendName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkInformation) Supports(features Features) bool {
	return (bool)(QNetworkInformation_Supports(this.h, features))
}

func (this *QNetworkInformation) SupportedFeatures() Features {
	xxxxxxxxx
}

func QNetworkInformation_LoadDefaultBackend() bool {
	return (bool)(QNetworkInformation_LoadDefaultBackend())
}

func QNetworkInformation_LoadBackendByFeatures(features Features) bool {
	return (bool)(QNetworkInformation_LoadBackendByFeatures(features))
}

func QNetworkInformation_LoadWithFeatures(features Features) bool {
	return (bool)(QNetworkInformation_LoadWithFeatures(features))
}

func QNetworkInformation_AvailableBackends() []string {
	var _ma struct_miqt_array = QNetworkInformation_AvailableBackends()
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

func QNetworkInformation_Instance() *QNetworkInformation {
	return newQNetworkInformation(QNetworkInformation_Instance())
}

func (this *QNetworkInformation) ReachabilityChanged(newReachability QNetworkInformation__Reachability) {
	QNetworkInformation_ReachabilityChanged(this.h, (int)(newReachability))
}

func (this *QNetworkInformation) OnReachabilityChanged(slot func(newReachability QNetworkInformation__Reachability)) {
	QNetworkInformation_connect_ReachabilityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkInformation_ReachabilityChanged
func miqt_exec_callback_QNetworkInformation_ReachabilityChanged(cb intptr_t, newReachability int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newReachability QNetworkInformation__Reachability))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkInformation__Reachability)(newReachability)

	gofunc(slotval1)
}

func (this *QNetworkInformation) IsBehindCaptivePortalChanged(state bool) {
	QNetworkInformation_IsBehindCaptivePortalChanged(this.h, (bool)(state))
}

func (this *QNetworkInformation) OnIsBehindCaptivePortalChanged(slot func(state bool)) {
	QNetworkInformation_connect_IsBehindCaptivePortalChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkInformation_IsBehindCaptivePortalChanged
func miqt_exec_callback_QNetworkInformation_IsBehindCaptivePortalChanged(cb intptr_t, state bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(state)

	gofunc(slotval1)
}

func (this *QNetworkInformation) TransportMediumChanged(current QNetworkInformation__TransportMedium) {
	QNetworkInformation_TransportMediumChanged(this.h, (int)(current))
}

func (this *QNetworkInformation) OnTransportMediumChanged(slot func(current QNetworkInformation__TransportMedium)) {
	QNetworkInformation_connect_TransportMediumChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkInformation_TransportMediumChanged
func miqt_exec_callback_QNetworkInformation_TransportMediumChanged(cb intptr_t, current int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current QNetworkInformation__TransportMedium))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkInformation__TransportMedium)(current)

	gofunc(slotval1)
}

func (this *QNetworkInformation) IsMeteredChanged(isMetered bool) {
	QNetworkInformation_IsMeteredChanged(this.h, (bool)(isMetered))
}

func (this *QNetworkInformation) OnIsMeteredChanged(slot func(isMetered bool)) {
	QNetworkInformation_connect_IsMeteredChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkInformation_IsMeteredChanged
func miqt_exec_callback_QNetworkInformation_IsMeteredChanged(cb intptr_t, isMetered bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(isMetered bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(isMetered)

	gofunc(slotval1)
}

func QNetworkInformation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkInformation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkInformation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkInformation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
