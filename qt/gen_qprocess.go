package qt

import (
	"unsafe"
)

type QProcessEnvironment__Initialization int

const (
	QProcessEnvironment__InheritFromParent QProcessEnvironment__Initialization = 0
)

type QProcess__ProcessError int

const (
	QProcess__FailedToStart QProcess__ProcessError = 0
	QProcess__Crashed       QProcess__ProcessError = 1
	QProcess__Timedout      QProcess__ProcessError = 2
	QProcess__ReadError     QProcess__ProcessError = 3
	QProcess__WriteError    QProcess__ProcessError = 4
	QProcess__UnknownError  QProcess__ProcessError = 5
)

type QProcess__ProcessState int

const (
	QProcess__NotRunning QProcess__ProcessState = 0
	QProcess__Starting   QProcess__ProcessState = 1
	QProcess__Running    QProcess__ProcessState = 2
)

type QProcess__ProcessChannel int

const (
	QProcess__StandardOutput QProcess__ProcessChannel = 0
	QProcess__StandardError  QProcess__ProcessChannel = 1
)

type QProcess__ProcessChannelMode int

const (
	QProcess__SeparateChannels       QProcess__ProcessChannelMode = 0
	QProcess__MergedChannels         QProcess__ProcessChannelMode = 1
	QProcess__ForwardedChannels      QProcess__ProcessChannelMode = 2
	QProcess__ForwardedOutputChannel QProcess__ProcessChannelMode = 3
	QProcess__ForwardedErrorChannel  QProcess__ProcessChannelMode = 4
)

type QProcess__InputChannelMode int

const (
	QProcess__ManagedInputChannel   QProcess__InputChannelMode = 0
	QProcess__ForwardedInputChannel QProcess__InputChannelMode = 1
)

type QProcess__ExitStatus int

const (
	QProcess__NormalExit QProcess__ExitStatus = 0
	QProcess__CrashExit  QProcess__ExitStatus = 1
)

type QProcessEnvironment struct {
	h          uintptr
	isSubclass bool
}

// NewQProcessEnvironment constructs a new QProcessEnvironment object.
func NewQProcessEnvironment() *QProcessEnvironment {
	g := newQProcessEnvironment(QProcessEnvironment_new())
	g.isSubclass = true
	return g
}

// NewQProcessEnvironment2 constructs a new QProcessEnvironment object.
func NewQProcessEnvironment2(param1 Initialization) *QProcessEnvironment {
	g := newQProcessEnvironment(QProcessEnvironment_new2(param1))
	g.isSubclass = true
	return g
}

// NewQProcessEnvironment3 constructs a new QProcessEnvironment object.
func NewQProcessEnvironment3(other *QProcessEnvironment) *QProcessEnvironment {
	g := newQProcessEnvironment(QProcessEnvironment_new3(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QProcessEnvironment) OperatorAssign(other *QProcessEnvironment) {
	QProcessEnvironment_OperatorAssign(this.h, other.cPointer())
}

func (this *QProcessEnvironment) Swap(other *QProcessEnvironment) {
	QProcessEnvironment_Swap(this.h, other.cPointer())
}

func (this *QProcessEnvironment) IsEmpty() bool {
	return (bool)(QProcessEnvironment_IsEmpty(this.h))
}

func (this *QProcessEnvironment) InheritsFromParent() bool {
	return (bool)(QProcessEnvironment_InheritsFromParent(this.h))
}

func (this *QProcessEnvironment) Clear() {
	QProcessEnvironment_Clear(this.h)
}

func (this *QProcessEnvironment) Contains(name string) bool {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	return (bool)(QProcessEnvironment_Contains(this.h, name_ms))
}

func (this *QProcessEnvironment) Insert(name string, value string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	value_ms := struct_miqt_string{}
	value_ms.data = CString(value)
	value_ms.len = size_t(len(value))
	defer free(unsafe.Pointer(value_ms.data))
	QProcessEnvironment_Insert(this.h, name_ms, value_ms)
}

func (this *QProcessEnvironment) Remove(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QProcessEnvironment_Remove(this.h, name_ms)
}

func (this *QProcessEnvironment) Value(name string) string {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	var _ms struct_miqt_string = QProcessEnvironment_Value(this.h, name_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProcessEnvironment) ToStringList() []string {
	var _ma struct_miqt_array = QProcessEnvironment_ToStringList(this.h)
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

func (this *QProcessEnvironment) Keys() []string {
	var _ma struct_miqt_array = QProcessEnvironment_Keys(this.h)
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

func (this *QProcessEnvironment) InsertWithQProcessEnvironment(e *QProcessEnvironment) {
	QProcessEnvironment_InsertWithQProcessEnvironment(this.h, e.cPointer())
}

func QProcessEnvironment_SystemEnvironment() *QProcessEnvironment {
	_goptr := newQProcessEnvironment(QProcessEnvironment_SystemEnvironment())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProcessEnvironment) Value2(name string, defaultValue string) string {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))
	var _ms struct_miqt_string = QProcessEnvironment_Value2(this.h, name_ms, defaultValue_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QProcess struct {
	h          uintptr
	isSubclass bool
}

// NewQProcess constructs a new QProcess object.
func NewQProcess() *QProcess {
	g := newQProcess(QProcess_new())
	g.isSubclass = true
	return g
}

// NewQProcess2 constructs a new QProcess object.
func NewQProcess2(parent *QObject) *QProcess {
	g := newQProcess(QProcess_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QProcess) MetaObject() *QMetaObject {
	return newQMetaObject(QProcess_MetaObject(this.h))
}

func (this *QProcess) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QProcess_Metacast(this.h, param1_Cstring))
}

func QProcess_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QProcess_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProcess) Start(program string) {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	QProcess_Start(this.h, program_ms)
}

func (this *QProcess) Start2() {
	QProcess_Start2(this.h)
}

func (this *QProcess) StartCommand(command string) {
	command_ms := struct_miqt_string{}
	command_ms.data = CString(command)
	command_ms.len = size_t(len(command))
	defer free(unsafe.Pointer(command_ms.data))
	QProcess_StartCommand(this.h, command_ms)
}

func (this *QProcess) StartDetached() bool {
	return (bool)(QProcess_StartDetached(this.h))
}

func (this *QProcess) Open(mode OpenMode) bool {
	return (bool)(QProcess_Open(this.h, mode))
}

func (this *QProcess) Program() string {
	var _ms struct_miqt_string = QProcess_Program(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProcess) SetProgram(program string) {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	QProcess_SetProgram(this.h, program_ms)
}

func (this *QProcess) Arguments() []string {
	var _ma struct_miqt_array = QProcess_Arguments(this.h)
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

func (this *QProcess) SetArguments(arguments []string) {
	arguments_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(arguments))))
	defer free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := struct_miqt_string{}
		arguments_i_ms.data = CString(arguments[i])
		arguments_i_ms.len = size_t(len(arguments[i]))
		defer free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := struct_miqt_array{len: size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	QProcess_SetArguments(this.h, arguments_ma)
}

func (this *QProcess) ProcessChannelMode() ProcessChannelMode {
	xxxxxxxxx
}

func (this *QProcess) SetProcessChannelMode(mode ProcessChannelMode) {
	QProcess_SetProcessChannelMode(this.h, mode)
}

func (this *QProcess) InputChannelMode() InputChannelMode {
	xxxxxxxxx
}

func (this *QProcess) SetInputChannelMode(mode InputChannelMode) {
	QProcess_SetInputChannelMode(this.h, mode)
}

func (this *QProcess) ReadChannel() ProcessChannel {
	xxxxxxxxx
}

func (this *QProcess) SetReadChannel(channel ProcessChannel) {
	QProcess_SetReadChannel(this.h, channel)
}

func (this *QProcess) CloseReadChannel(channel ProcessChannel) {
	QProcess_CloseReadChannel(this.h, channel)
}

func (this *QProcess) CloseWriteChannel() {
	QProcess_CloseWriteChannel(this.h)
}

func (this *QProcess) SetStandardInputFile(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QProcess_SetStandardInputFile(this.h, fileName_ms)
}

func (this *QProcess) SetStandardOutputFile(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QProcess_SetStandardOutputFile(this.h, fileName_ms)
}

func (this *QProcess) SetStandardErrorFile(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QProcess_SetStandardErrorFile(this.h, fileName_ms)
}

func (this *QProcess) SetStandardOutputProcess(destination *QProcess) {
	QProcess_SetStandardOutputProcess(this.h, destination.cPointer())
}

func (this *QProcess) NativeArguments() string {
	var _ms struct_miqt_string = QProcess_NativeArguments(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProcess) SetNativeArguments(arguments string) {
	arguments_ms := struct_miqt_string{}
	arguments_ms.data = CString(arguments)
	arguments_ms.len = size_t(len(arguments))
	defer free(unsafe.Pointer(arguments_ms.data))
	QProcess_SetNativeArguments(this.h, arguments_ms)
}

func (this *QProcess) CreateProcessArgumentsModifier() CreateProcessArgumentModifier {
	xxxxxxxxx
}

func (this *QProcess) SetCreateProcessArgumentsModifier(modifier CreateProcessArgumentModifier) {
	QProcess_SetCreateProcessArgumentsModifier(this.h, modifier)
}

func (this *QProcess) WorkingDirectory() string {
	var _ms struct_miqt_string = QProcess_WorkingDirectory(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProcess) SetWorkingDirectory(dir string) {
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	QProcess_SetWorkingDirectory(this.h, dir_ms)
}

func (this *QProcess) SetEnvironment(environment []string) {
	environment_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(environment))))
	defer free(unsafe.Pointer(environment_CArray))
	for i := range environment {
		environment_i_ms := struct_miqt_string{}
		environment_i_ms.data = CString(environment[i])
		environment_i_ms.len = size_t(len(environment[i]))
		defer free(unsafe.Pointer(environment_i_ms.data))
		environment_CArray[i] = environment_i_ms
	}
	environment_ma := struct_miqt_array{len: size_t(len(environment)), data: unsafe.Pointer(environment_CArray)}
	QProcess_SetEnvironment(this.h, environment_ma)
}

func (this *QProcess) Environment() []string {
	var _ma struct_miqt_array = QProcess_Environment(this.h)
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

func (this *QProcess) SetProcessEnvironment(environment *QProcessEnvironment) {
	QProcess_SetProcessEnvironment(this.h, environment.cPointer())
}

func (this *QProcess) ProcessEnvironment() *QProcessEnvironment {
	_goptr := newQProcessEnvironment(QProcess_ProcessEnvironment(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProcess) Error() QProcess__ProcessError {
	return (QProcess__ProcessError)(QProcess_Error(this.h))
}

func (this *QProcess) State() QProcess__ProcessState {
	return (QProcess__ProcessState)(QProcess_State(this.h))
}

func (this *QProcess) ProcessId() int64 {
	return (int64)(QProcess_ProcessId(this.h))
}

func (this *QProcess) WaitForStarted() bool {
	return (bool)(QProcess_WaitForStarted(this.h))
}

func (this *QProcess) WaitForReadyRead(msecs int) bool {
	return (bool)(QProcess_WaitForReadyRead(this.h, (int)(msecs)))
}

func (this *QProcess) WaitForBytesWritten(msecs int) bool {
	return (bool)(QProcess_WaitForBytesWritten(this.h, (int)(msecs)))
}

func (this *QProcess) WaitForFinished() bool {
	return (bool)(QProcess_WaitForFinished(this.h))
}

func (this *QProcess) ReadAllStandardOutput() []byte {
	var _bytearray struct_miqt_string = QProcess_ReadAllStandardOutput(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QProcess) ReadAllStandardError() []byte {
	var _bytearray struct_miqt_string = QProcess_ReadAllStandardError(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QProcess) ExitCode() int {
	return (int)(QProcess_ExitCode(this.h))
}

func (this *QProcess) ExitStatus() QProcess__ExitStatus {
	return (QProcess__ExitStatus)(QProcess_ExitStatus(this.h))
}

func (this *QProcess) BytesToWrite() int64 {
	return (int64)(QProcess_BytesToWrite(this.h))
}

func (this *QProcess) IsSequential() bool {
	return (bool)(QProcess_IsSequential(this.h))
}

func (this *QProcess) Close() {
	QProcess_Close(this.h)
}

func QProcess_Execute(program string) int {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	return (int)(QProcess_Execute(program_ms))
}

func QProcess_StartDetachedWithProgram(program string) bool {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	return (bool)(QProcess_StartDetachedWithProgram(program_ms))
}

func QProcess_SystemEnvironment() []string {
	var _ma struct_miqt_array = QProcess_SystemEnvironment()
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

func QProcess_NullDevice() string {
	var _ms struct_miqt_string = QProcess_NullDevice()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProcess) Terminate() {
	QProcess_Terminate(this.h)
}

func (this *QProcess) Kill() {
	QProcess_Kill(this.h)
}

func (this *QProcess) Finished(exitCode int) {
	QProcess_Finished(this.h, (int)(exitCode))
}

func (this *QProcess) OnFinished(slot func(exitCode int)) {
	QProcess_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProcess_Finished
func miqt_exec_callback_QProcess_Finished(cb intptr_t, exitCode int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(exitCode int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(exitCode)

	gofunc(slotval1)
}

func (this *QProcess) ErrorOccurred(error QProcess__ProcessError) {
	QProcess_ErrorOccurred(this.h, (int)(error))
}

func (this *QProcess) OnErrorOccurred(slot func(error QProcess__ProcessError)) {
	QProcess_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProcess_ErrorOccurred
func miqt_exec_callback_QProcess_ErrorOccurred(cb intptr_t, error int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QProcess__ProcessError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QProcess__ProcessError)(error)

	gofunc(slotval1)
}

func QProcess_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProcess_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProcess_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProcess_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProcess) Start22(program string, arguments []string) {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	arguments_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(arguments))))
	defer free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := struct_miqt_string{}
		arguments_i_ms.data = CString(arguments[i])
		arguments_i_ms.len = size_t(len(arguments[i]))
		defer free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := struct_miqt_array{len: size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	QProcess_Start22(this.h, program_ms, arguments_ma)
}

func (this *QProcess) Start3(program string, arguments []string, mode OpenMode) {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	arguments_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(arguments))))
	defer free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := struct_miqt_string{}
		arguments_i_ms.data = CString(arguments[i])
		arguments_i_ms.len = size_t(len(arguments[i]))
		defer free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := struct_miqt_array{len: size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	QProcess_Start3(this.h, program_ms, arguments_ma, mode)
}

func (this *QProcess) Start1(mode OpenMode) {
	QProcess_Start1(this.h, mode)
}

func (this *QProcess) StartCommand2(command string, mode OpenMode) {
	command_ms := struct_miqt_string{}
	command_ms.data = CString(command)
	command_ms.len = size_t(len(command))
	defer free(unsafe.Pointer(command_ms.data))
	QProcess_StartCommand2(this.h, command_ms, mode)
}

func (this *QProcess) StartDetached1(pid *int64) bool {
	return (bool)(QProcess_StartDetached1(this.h, (*longlong)(unsafe.Pointer(pid))))
}

func (this *QProcess) SetStandardOutputFile2(fileName string, mode OpenMode) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QProcess_SetStandardOutputFile2(this.h, fileName_ms, mode)
}

func (this *QProcess) SetStandardErrorFile2(fileName string, mode OpenMode) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QProcess_SetStandardErrorFile2(this.h, fileName_ms, mode)
}

func (this *QProcess) WaitForStarted1(msecs int) bool {
	return (bool)(QProcess_WaitForStarted1(this.h, (int)(msecs)))
}

func (this *QProcess) WaitForFinished1(msecs int) bool {
	return (bool)(QProcess_WaitForFinished1(this.h, (int)(msecs)))
}

func QProcess_Execute2(program string, arguments []string) int {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	arguments_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(arguments))))
	defer free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := struct_miqt_string{}
		arguments_i_ms.data = CString(arguments[i])
		arguments_i_ms.len = size_t(len(arguments[i]))
		defer free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := struct_miqt_array{len: size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	return (int)(QProcess_Execute2(program_ms, arguments_ma))
}

func QProcess_StartDetached2(program string, arguments []string) bool {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	arguments_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(arguments))))
	defer free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := struct_miqt_string{}
		arguments_i_ms.data = CString(arguments[i])
		arguments_i_ms.len = size_t(len(arguments[i]))
		defer free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := struct_miqt_array{len: size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	return (bool)(QProcess_StartDetached2(program_ms, arguments_ma))
}

func QProcess_StartDetached3(program string, arguments []string, workingDirectory string) bool {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	arguments_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(arguments))))
	defer free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := struct_miqt_string{}
		arguments_i_ms.data = CString(arguments[i])
		arguments_i_ms.len = size_t(len(arguments[i]))
		defer free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := struct_miqt_array{len: size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	workingDirectory_ms := struct_miqt_string{}
	workingDirectory_ms.data = CString(workingDirectory)
	workingDirectory_ms.len = size_t(len(workingDirectory))
	defer free(unsafe.Pointer(workingDirectory_ms.data))
	return (bool)(QProcess_StartDetached3(program_ms, arguments_ma, workingDirectory_ms))
}

func QProcess_StartDetached4(program string, arguments []string, workingDirectory string, pid *int64) bool {
	program_ms := struct_miqt_string{}
	program_ms.data = CString(program)
	program_ms.len = size_t(len(program))
	defer free(unsafe.Pointer(program_ms.data))
	arguments_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(arguments))))
	defer free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := struct_miqt_string{}
		arguments_i_ms.data = CString(arguments[i])
		arguments_i_ms.len = size_t(len(arguments[i]))
		defer free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := struct_miqt_array{len: size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	workingDirectory_ms := struct_miqt_string{}
	workingDirectory_ms.data = CString(workingDirectory)
	workingDirectory_ms.len = size_t(len(workingDirectory))
	defer free(unsafe.Pointer(workingDirectory_ms.data))
	return (bool)(QProcess_StartDetached4(program_ms, arguments_ma, workingDirectory_ms, (*longlong)(unsafe.Pointer(pid))))
}

func (this *QProcess) Finished2(exitCode int, exitStatus QProcess__ExitStatus) {
	QProcess_Finished2(this.h, (int)(exitCode), (int)(exitStatus))
}

func (this *QProcess) OnFinished2(slot func(exitCode int, exitStatus QProcess__ExitStatus)) {
	QProcess_connect_Finished2(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProcess_Finished2
func miqt_exec_callback_QProcess_Finished2(cb intptr_t, exitCode int, exitStatus int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(exitCode int, exitStatus QProcess__ExitStatus))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(exitCode)

	slotval2 := (QProcess__ExitStatus)(exitStatus)

	gofunc(slotval1, slotval2)
}

func (this *QProcess) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QProcess_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QProcess) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProcess_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProcess_MetaObject
func miqt_exec_callback_QProcess_MetaObject(self QProcess, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProcess{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QProcess) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QProcess_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QProcess) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProcess_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProcess_Metacast
func miqt_exec_callback_QProcess_Metacast(self QProcess, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QProcess{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QProcess__CreateProcessArguments struct {
	h          uintptr
	isSubclass bool
}
