package qt

import (
	"unsafe"
)

type QCoreApplication__ int

const (
	QCoreApplication__ApplicationFlags QCoreApplication__ = 395520
)

type QCoreApplication struct {
	h          uintptr
	isSubclass bool
}

// NewQCoreApplication constructs a new QCoreApplication object.
func NewQCoreApplication(args []string) *QCoreApplication {
	argcPtr := uintptr(unsafe.Pointer(&argc))
	var argvPtr uintptr
	if argv != nil && len(argv) > 0 {
		argvPtr = uintptr(unsafe.Pointer(&argv[0])) // 获取argv的指针
	} else {
		argvPtr = 0 // 或者使用nil
	}
	ret := newQCoreApplication(QCoreApplication_new())
	ret.isSubclass = true
	return ret
}

// NewQCoreApplication2 constructs a new QCoreApplication object.
func NewQCoreApplication2(args []string, param3 int) *QCoreApplication {
	argcPtr := uintptr(unsafe.Pointer(&argc))
	var argvPtr uintptr
	if argv != nil && len(argv) > 0 {
		argvPtr = uintptr(unsafe.Pointer(&argv[0])) // 获取argv的指针
	} else {
		argvPtr = 0 // 或者使用nil
	}
	ret := newQCoreApplication(QCoreApplication_new2((int)(param3)))
	ret.isSubclass = true
	return ret
}

func (this *QCoreApplication) MetaObject() *QMetaObject {
	return newQMetaObject(QCoreApplication_MetaObject(this.h))
}

func (this *QCoreApplication) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QCoreApplication_Metacast(this.h, param1_Cstring))
}

func QCoreApplication_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QCoreApplication_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_Arguments() []string {
	var _ma struct_miqt_array = QCoreApplication_Arguments()
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

func QCoreApplication_SetAttribute(attribute ApplicationAttribute) {
	QCoreApplication_SetAttribute((int)(attribute))
}

func QCoreApplication_TestAttribute(attribute ApplicationAttribute) bool {
	return (bool)(QCoreApplication_TestAttribute((int)(attribute)))
}

func QCoreApplication_SetOrganizationDomain(orgDomain string) {
	orgDomain_ms := struct_miqt_string{}
	orgDomain_ms.data = CString(orgDomain)
	orgDomain_ms.len = size_t(len(orgDomain))
	defer free(unsafe.Pointer(orgDomain_ms.data))
	QCoreApplication_SetOrganizationDomain(orgDomain_ms)
}

func QCoreApplication_OrganizationDomain() string {
	var _ms struct_miqt_string = QCoreApplication_OrganizationDomain()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_SetOrganizationName(orgName string) {
	orgName_ms := struct_miqt_string{}
	orgName_ms.data = CString(orgName)
	orgName_ms.len = size_t(len(orgName))
	defer free(unsafe.Pointer(orgName_ms.data))
	QCoreApplication_SetOrganizationName(orgName_ms)
}

func QCoreApplication_OrganizationName() string {
	var _ms struct_miqt_string = QCoreApplication_OrganizationName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_SetApplicationName(application string) {
	application_ms := struct_miqt_string{}
	application_ms.data = CString(application)
	application_ms.len = size_t(len(application))
	defer free(unsafe.Pointer(application_ms.data))
	QCoreApplication_SetApplicationName(application_ms)
}

func QCoreApplication_ApplicationName() string {
	var _ms struct_miqt_string = QCoreApplication_ApplicationName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_SetApplicationVersion(version string) {
	version_ms := struct_miqt_string{}
	version_ms.data = CString(version)
	version_ms.len = size_t(len(version))
	defer free(unsafe.Pointer(version_ms.data))
	QCoreApplication_SetApplicationVersion(version_ms)
}

func QCoreApplication_ApplicationVersion() string {
	var _ms struct_miqt_string = QCoreApplication_ApplicationVersion()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_SetSetuidAllowed(allow bool) {
	QCoreApplication_SetSetuidAllowed((bool)(allow))
}

func QCoreApplication_IsSetuidAllowed() bool {
	return (bool)(QCoreApplication_IsSetuidAllowed())
}

func QCoreApplication_Instance() *QCoreApplication {
	return newQCoreApplication(QCoreApplication_Instance())
}

func QCoreApplication_Exec() int {
	return (int)(QCoreApplication_Exec())
}

func QCoreApplication_ProcessEvents() {
	QCoreApplication_ProcessEvents()
}

func QCoreApplication_ProcessEvents2(flags ProcessEventsFlag, maxtime int) {
	QCoreApplication_ProcessEvents2((int)(flags), (int)(maxtime))
}

func QCoreApplication_ProcessEvents3(flags ProcessEventsFlag, deadline QDeadlineTimer) {
	QCoreApplication_ProcessEvents3((int)(flags), deadline.cPointer())
}

func QCoreApplication_SendEvent(receiver *QObject, event *QEvent) bool {
	return (bool)(QCoreApplication_SendEvent(receiver.cPointer(), event.cPointer()))
}

func QCoreApplication_PostEvent(receiver *QObject, event *QEvent) {
	QCoreApplication_PostEvent(receiver.cPointer(), event.cPointer())
}

func QCoreApplication_SendPostedEvents() {
	QCoreApplication_SendPostedEvents()
}

func QCoreApplication_RemovePostedEvents(receiver *QObject) {
	QCoreApplication_RemovePostedEvents(receiver.cPointer())
}

func QCoreApplication_EventDispatcher() *QAbstractEventDispatcher {
	return newQAbstractEventDispatcher(QCoreApplication_EventDispatcher())
}

func QCoreApplication_SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher) {
	QCoreApplication_SetEventDispatcher(eventDispatcher.cPointer())
}

func (this *QCoreApplication) Notify(param1 *QObject, param2 *QEvent) bool {
	return (bool)(QCoreApplication_Notify(this.h, param1.cPointer(), param2.cPointer()))
}

func QCoreApplication_StartingUp() bool {
	return (bool)(QCoreApplication_StartingUp())
}

func QCoreApplication_ClosingDown() bool {
	return (bool)(QCoreApplication_ClosingDown())
}

func QCoreApplication_ApplicationDirPath() string {
	var _ms struct_miqt_string = QCoreApplication_ApplicationDirPath()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_ApplicationFilePath() string {
	var _ms struct_miqt_string = QCoreApplication_ApplicationFilePath()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_ApplicationPid() int64 {
	return (int64)(QCoreApplication_ApplicationPid())
}

func (this *QCoreApplication) CheckPermission(permission *QPermission) PermissionStatus {
	return (PermissionStatus)(QCoreApplication_CheckPermission(this.h, permission.cPointer()))
}

func QCoreApplication_SetLibraryPaths(libraryPaths []string) {
	libraryPaths_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(libraryPaths))))
	defer free(unsafe.Pointer(libraryPaths_CArray))
	for i := range libraryPaths {
		libraryPaths_i_ms := struct_miqt_string{}
		libraryPaths_i_ms.data = CString(libraryPaths[i])
		libraryPaths_i_ms.len = size_t(len(libraryPaths[i]))
		defer free(unsafe.Pointer(libraryPaths_i_ms.data))
		libraryPaths_CArray[i] = libraryPaths_i_ms
	}
	libraryPaths_ma := struct_miqt_array{len: size_t(len(libraryPaths)), data: unsafe.Pointer(libraryPaths_CArray)}
	QCoreApplication_SetLibraryPaths(libraryPaths_ma)
}

func QCoreApplication_LibraryPaths() []string {
	var _ma struct_miqt_array = QCoreApplication_LibraryPaths()
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

func QCoreApplication_AddLibraryPath(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QCoreApplication_AddLibraryPath(param1_ms)
}

func QCoreApplication_RemoveLibraryPath(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QCoreApplication_RemoveLibraryPath(param1_ms)
}

func QCoreApplication_InstallTranslator(messageFile *QTranslator) bool {
	return (bool)(QCoreApplication_InstallTranslator(messageFile.cPointer()))
}

func QCoreApplication_RemoveTranslator(messageFile *QTranslator) bool {
	return (bool)(QCoreApplication_RemoveTranslator(messageFile.cPointer()))
}

func QCoreApplication_Translate(context string, key string) string {
	context_Cstring := CString(context)
	defer free(unsafe.Pointer(context_Cstring))
	key_Cstring := CString(key)
	defer free(unsafe.Pointer(key_Cstring))
	var _ms struct_miqt_string = QCoreApplication_Translate(context_Cstring, key_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCoreApplication) InstallNativeEventFilter(filterObj *QAbstractNativeEventFilter) {
	QCoreApplication_InstallNativeEventFilter(this.h, filterObj.cPointer())
}
func (this *QCoreApplication) OnInstallNativeEventFilter(slot func(filterObj *QAbstractNativeEventFilter)) {
	QCoreApplication_connect_InstallNativeEventFilter(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_InstallNativeEventFilter
func miqt_exec_callback_QCoreApplication_InstallNativeEventFilter(cb intptr_t, filterObj *QAbstractNativeEventFilter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(filterObj *QAbstractNativeEventFilter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractNativeEventFilter(filterObj)

	gofunc(slotval1)
}

func (this *QCoreApplication) RemoveNativeEventFilter(filterObj *QAbstractNativeEventFilter) {
	QCoreApplication_RemoveNativeEventFilter(this.h, filterObj.cPointer())
}
func (this *QCoreApplication) OnRemoveNativeEventFilter(slot func(filterObj *QAbstractNativeEventFilter)) {
	QCoreApplication_connect_RemoveNativeEventFilter(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_RemoveNativeEventFilter
func miqt_exec_callback_QCoreApplication_RemoveNativeEventFilter(cb intptr_t, filterObj *QAbstractNativeEventFilter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(filterObj *QAbstractNativeEventFilter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractNativeEventFilter(filterObj)

	gofunc(slotval1)
}

func QCoreApplication_IsQuitLockEnabled() bool {
	return (bool)(QCoreApplication_IsQuitLockEnabled())
}

func QCoreApplication_SetQuitLockEnabled(enabled bool) {
	QCoreApplication_SetQuitLockEnabled((bool)(enabled))
}

func QCoreApplication_Quit() {
	QCoreApplication_Quit()
}

func QCoreApplication_Exit() {
	QCoreApplication_Exit()
}

func (this *QCoreApplication) OrganizationNameChanged() {
	QCoreApplication_OrganizationNameChanged(this.h)
}
func (this *QCoreApplication) OnOrganizationNameChanged(slot func()) {
	QCoreApplication_connect_OrganizationNameChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_OrganizationNameChanged
func miqt_exec_callback_QCoreApplication_OrganizationNameChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCoreApplication) OrganizationDomainChanged() {
	QCoreApplication_OrganizationDomainChanged(this.h)
}
func (this *QCoreApplication) OnOrganizationDomainChanged(slot func()) {
	QCoreApplication_connect_OrganizationDomainChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_OrganizationDomainChanged
func miqt_exec_callback_QCoreApplication_OrganizationDomainChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCoreApplication) ApplicationNameChanged() {
	QCoreApplication_ApplicationNameChanged(this.h)
}
func (this *QCoreApplication) OnApplicationNameChanged(slot func()) {
	QCoreApplication_connect_ApplicationNameChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_ApplicationNameChanged
func miqt_exec_callback_QCoreApplication_ApplicationNameChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCoreApplication) ApplicationVersionChanged() {
	QCoreApplication_ApplicationVersionChanged(this.h)
}
func (this *QCoreApplication) OnApplicationVersionChanged(slot func()) {
	QCoreApplication_connect_ApplicationVersionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_ApplicationVersionChanged
func miqt_exec_callback_QCoreApplication_ApplicationVersionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QCoreApplication_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCoreApplication_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCoreApplication_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_SetAttribute2(attribute ApplicationAttribute, on bool) {
	QCoreApplication_SetAttribute2((int)(attribute), (bool)(on))
}

func QCoreApplication_ProcessEvents1(flags ProcessEventsFlag) {
	QCoreApplication_ProcessEvents1((int)(flags))
}

func QCoreApplication_PostEvent3(receiver *QObject, event *QEvent, priority int) {
	QCoreApplication_PostEvent3(receiver.cPointer(), event.cPointer(), (int)(priority))
}

func QCoreApplication_SendPostedEvents1(receiver *QObject) {
	QCoreApplication_SendPostedEvents1(receiver.cPointer())
}

func QCoreApplication_SendPostedEvents2(receiver *QObject, event_type int) {
	QCoreApplication_SendPostedEvents2(receiver.cPointer(), (int)(event_type))
}

func QCoreApplication_RemovePostedEvents2(receiver *QObject, eventType int) {
	QCoreApplication_RemovePostedEvents2(receiver.cPointer(), (int)(eventType))
}

func QCoreApplication_Translate3(context string, key string, disambiguation string) string {
	context_Cstring := CString(context)
	defer free(unsafe.Pointer(context_Cstring))
	key_Cstring := CString(key)
	defer free(unsafe.Pointer(key_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QCoreApplication_Translate3(context_Cstring, key_Cstring, disambiguation_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_Translate4(context string, key string, disambiguation string, n int) string {
	context_Cstring := CString(context)
	defer free(unsafe.Pointer(context_Cstring))
	key_Cstring := CString(key)
	defer free(unsafe.Pointer(key_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QCoreApplication_Translate4(context_Cstring, key_Cstring, disambiguation_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCoreApplication_Exit1(retcode int) {
	QCoreApplication_Exit1((int)(retcode))
}

func (this *QCoreApplication) callVirtualBase_Notify(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QCoreApplication_virtualbase_Notify(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QCoreApplication) OnNotify(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCoreApplication_override_virtual_Notify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_Notify
func miqt_exec_callback_QCoreApplication_Notify(self QCoreApplication, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QCoreApplication{h: self}).callVirtualBase_Notify, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QCoreApplication) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QCoreApplication_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QCoreApplication) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCoreApplication_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_Event
func miqt_exec_callback_QCoreApplication_Event(self QCoreApplication, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QCoreApplication{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QCoreApplication) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QCoreApplication_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QCoreApplication) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCoreApplication_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_EventFilter
func miqt_exec_callback_QCoreApplication_EventFilter(self QCoreApplication, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QCoreApplication{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QCoreApplication) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QCoreApplication_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCoreApplication) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCoreApplication_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_TimerEvent
func miqt_exec_callback_QCoreApplication_TimerEvent(self QCoreApplication, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QCoreApplication{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QCoreApplication) callVirtualBase_ChildEvent(event *QChildEvent) {

	QCoreApplication_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCoreApplication) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCoreApplication_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_ChildEvent
func miqt_exec_callback_QCoreApplication_ChildEvent(self QCoreApplication, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QCoreApplication{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QCoreApplication) callVirtualBase_CustomEvent(event *QEvent) {

	QCoreApplication_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCoreApplication) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCoreApplication_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_CustomEvent
func miqt_exec_callback_QCoreApplication_CustomEvent(self QCoreApplication, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QCoreApplication{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QCoreApplication) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QCoreApplication_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QCoreApplication) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCoreApplication_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_ConnectNotify
func miqt_exec_callback_QCoreApplication_ConnectNotify(self QCoreApplication, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QCoreApplication{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QCoreApplication) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QCoreApplication_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QCoreApplication) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCoreApplication_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCoreApplication_DisconnectNotify
func miqt_exec_callback_QCoreApplication_DisconnectNotify(self QCoreApplication, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QCoreApplication{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
