package qt

import (
	"unsafe"
)

type QSessionManager__RestartHint int

const (
	QSessionManager__RestartIfRunning   QSessionManager__RestartHint = 0
	QSessionManager__RestartAnyway      QSessionManager__RestartHint = 1
	QSessionManager__RestartImmediately QSessionManager__RestartHint = 2
	QSessionManager__RestartNever       QSessionManager__RestartHint = 3
)

type QSessionManager struct {
	h          uintptr
	isSubclass bool
}

func (this *QSessionManager) MetaObject() *QMetaObject {
	return newQMetaObject(QSessionManager_MetaObject(this.h))
}

func (this *QSessionManager) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSessionManager_Metacast(this.h, param1_Cstring))
}

func QSessionManager_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSessionManager_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSessionManager) SessionId() string {
	var _ms struct_miqt_string = QSessionManager_SessionId(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSessionManager) SessionKey() string {
	var _ms struct_miqt_string = QSessionManager_SessionKey(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSessionManager) AllowsInteraction() bool {
	return (bool)(QSessionManager_AllowsInteraction(this.h))
}

func (this *QSessionManager) AllowsErrorInteraction() bool {
	return (bool)(QSessionManager_AllowsErrorInteraction(this.h))
}

func (this *QSessionManager) Release() {
	QSessionManager_Release(this.h)
}

func (this *QSessionManager) Cancel() {
	QSessionManager_Cancel(this.h)
}

func (this *QSessionManager) SetRestartHint(restartHint RestartHint) {
	QSessionManager_SetRestartHint(this.h, restartHint)
}

func (this *QSessionManager) RestartHint() RestartHint {
	xxxxxxxxx
}

func (this *QSessionManager) SetRestartCommand(restartCommand []string) {
	restartCommand_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(restartCommand))))
	defer free(unsafe.Pointer(restartCommand_CArray))
	for i := range restartCommand {
		restartCommand_i_ms := struct_miqt_string{}
		restartCommand_i_ms.data = CString(restartCommand[i])
		restartCommand_i_ms.len = size_t(len(restartCommand[i]))
		defer free(unsafe.Pointer(restartCommand_i_ms.data))
		restartCommand_CArray[i] = restartCommand_i_ms
	}
	restartCommand_ma := struct_miqt_array{len: size_t(len(restartCommand)), data: unsafe.Pointer(restartCommand_CArray)}
	QSessionManager_SetRestartCommand(this.h, restartCommand_ma)
}

func (this *QSessionManager) RestartCommand() []string {
	var _ma struct_miqt_array = QSessionManager_RestartCommand(this.h)
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

func (this *QSessionManager) SetDiscardCommand(discardCommand []string) {
	discardCommand_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(discardCommand))))
	defer free(unsafe.Pointer(discardCommand_CArray))
	for i := range discardCommand {
		discardCommand_i_ms := struct_miqt_string{}
		discardCommand_i_ms.data = CString(discardCommand[i])
		discardCommand_i_ms.len = size_t(len(discardCommand[i]))
		defer free(unsafe.Pointer(discardCommand_i_ms.data))
		discardCommand_CArray[i] = discardCommand_i_ms
	}
	discardCommand_ma := struct_miqt_array{len: size_t(len(discardCommand)), data: unsafe.Pointer(discardCommand_CArray)}
	QSessionManager_SetDiscardCommand(this.h, discardCommand_ma)
}

func (this *QSessionManager) DiscardCommand() []string {
	var _ma struct_miqt_array = QSessionManager_DiscardCommand(this.h)
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

func (this *QSessionManager) SetManagerProperty(name string, value string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	value_ms := struct_miqt_string{}
	value_ms.data = CString(value)
	value_ms.len = size_t(len(value))
	defer free(unsafe.Pointer(value_ms.data))
	QSessionManager_SetManagerProperty(this.h, name_ms, value_ms)
}

func (this *QSessionManager) SetManagerProperty2(name string, value []string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	value_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(value))))
	defer free(unsafe.Pointer(value_CArray))
	for i := range value {
		value_i_ms := struct_miqt_string{}
		value_i_ms.data = CString(value[i])
		value_i_ms.len = size_t(len(value[i]))
		defer free(unsafe.Pointer(value_i_ms.data))
		value_CArray[i] = value_i_ms
	}
	value_ma := struct_miqt_array{len: size_t(len(value)), data: unsafe.Pointer(value_CArray)}
	QSessionManager_SetManagerProperty2(this.h, name_ms, value_ma)
}

func (this *QSessionManager) IsPhase2() bool {
	return (bool)(QSessionManager_IsPhase2(this.h))
}

func (this *QSessionManager) RequestPhase2() {
	QSessionManager_RequestPhase2(this.h)
}

func QSessionManager_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSessionManager_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSessionManager_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSessionManager_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
