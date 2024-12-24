package qt

import (
	"unsafe"
)

type QCommandLineParser__SingleDashWordOptionMode int

const (
	QCommandLineParser__ParseAsCompactedShortOptions QCommandLineParser__SingleDashWordOptionMode = 0
	QCommandLineParser__ParseAsLongOptions           QCommandLineParser__SingleDashWordOptionMode = 1
)

type QCommandLineParser__OptionsAfterPositionalArgumentsMode int

const (
	QCommandLineParser__ParseAsOptions             QCommandLineParser__OptionsAfterPositionalArgumentsMode = 0
	QCommandLineParser__ParseAsPositionalArguments QCommandLineParser__OptionsAfterPositionalArgumentsMode = 1
)

type QCommandLineParser__MessageType int

const (
	QCommandLineParser__InformationMessage QCommandLineParser__MessageType = 0
	QCommandLineParser__ErrorMessage       QCommandLineParser__MessageType = 1
)

type QCommandLineParser struct {
	h          uintptr
	isSubclass bool
}

// NewQCommandLineParser constructs a new QCommandLineParser object.
func NewQCommandLineParser() *QCommandLineParser {
	g := newQCommandLineParser(QCommandLineParser_new())
	g.isSubclass = true
	return g
}

func QCommandLineParser_Tr(sourceText string) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	var _ms struct_miqt_string = QCommandLineParser_Tr(sourceText_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode SingleDashWordOptionMode) {
	QCommandLineParser_SetSingleDashWordOptionMode(this.h, parsingMode)
}

func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode OptionsAfterPositionalArgumentsMode) {
	QCommandLineParser_SetOptionsAfterPositionalArgumentsMode(this.h, mode)
}

func (this *QCommandLineParser) AddOption(commandLineOption *QCommandLineOption) bool {
	return (bool)(QCommandLineParser_AddOption(this.h, commandLineOption.cPointer()))
}

func (this *QCommandLineParser) AddOptions(options []QCommandLineOption) bool {
	options_CArray := (*[0xffff]*QCommandLineOption)(malloc(size_t(8 * len(options))))
	defer free(unsafe.Pointer(options_CArray))
	for i := range options {
		options_CArray[i] = options[i].cPointer()
	}
	options_ma := struct_miqt_array{len: size_t(len(options)), data: unsafe.Pointer(options_CArray)}
	return (bool)(QCommandLineParser_AddOptions(this.h, options_ma))
}

func (this *QCommandLineParser) AddVersionOption() *QCommandLineOption {
	_goptr := newQCommandLineOption(QCommandLineParser_AddVersionOption(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommandLineParser) AddHelpOption() *QCommandLineOption {
	_goptr := newQCommandLineOption(QCommandLineParser_AddHelpOption(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommandLineParser) SetApplicationDescription(description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QCommandLineParser_SetApplicationDescription(this.h, description_ms)
}

func (this *QCommandLineParser) ApplicationDescription() string {
	var _ms struct_miqt_string = QCommandLineParser_ApplicationDescription(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) AddPositionalArgument(name string, description string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QCommandLineParser_AddPositionalArgument(this.h, name_ms, description_ms)
}

func (this *QCommandLineParser) ClearPositionalArguments() {
	QCommandLineParser_ClearPositionalArguments(this.h)
}

func (this *QCommandLineParser) Process(arguments []string) {
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
	QCommandLineParser_Process(this.h, arguments_ma)
}

func (this *QCommandLineParser) ProcessWithApp(app *QCoreApplication) {
	QCommandLineParser_ProcessWithApp(this.h, app.cPointer())
}

func (this *QCommandLineParser) Parse(arguments []string) bool {
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
	return (bool)(QCommandLineParser_Parse(this.h, arguments_ma))
}

func (this *QCommandLineParser) ErrorText() string {
	var _ms struct_miqt_string = QCommandLineParser_ErrorText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) IsSet(name string) bool {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	return (bool)(QCommandLineParser_IsSet(this.h, name_ms))
}

func (this *QCommandLineParser) Value(name string) string {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	var _ms struct_miqt_string = QCommandLineParser_Value(this.h, name_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) Values(name string) []string {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	var _ma struct_miqt_array = QCommandLineParser_Values(this.h, name_ms)
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

func (this *QCommandLineParser) IsSetWithOption(option *QCommandLineOption) bool {
	return (bool)(QCommandLineParser_IsSetWithOption(this.h, option.cPointer()))
}

func (this *QCommandLineParser) ValueWithOption(option *QCommandLineOption) string {
	var _ms struct_miqt_string = QCommandLineParser_ValueWithOption(this.h, option.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) ValuesWithOption(option *QCommandLineOption) []string {
	var _ma struct_miqt_array = QCommandLineParser_ValuesWithOption(this.h, option.cPointer())
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

func (this *QCommandLineParser) PositionalArguments() []string {
	var _ma struct_miqt_array = QCommandLineParser_PositionalArguments(this.h)
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

func (this *QCommandLineParser) OptionNames() []string {
	var _ma struct_miqt_array = QCommandLineParser_OptionNames(this.h)
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

func (this *QCommandLineParser) UnknownOptionNames() []string {
	var _ma struct_miqt_array = QCommandLineParser_UnknownOptionNames(this.h)
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

func (this *QCommandLineParser) ShowVersion() {
	QCommandLineParser_ShowVersion(this.h)
}

func (this *QCommandLineParser) ShowHelp() {
	QCommandLineParser_ShowHelp(this.h)
}

func (this *QCommandLineParser) HelpText() string {
	var _ms struct_miqt_string = QCommandLineParser_HelpText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCommandLineParser_ShowMessageAndExit(message string, typeVal MessageType) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QCommandLineParser_ShowMessageAndExit(message_ms, typeVal)
}

func QCommandLineParser_Tr2(sourceText string, disambiguation string) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QCommandLineParser_Tr2(sourceText_Cstring, disambiguation_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCommandLineParser_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QCommandLineParser_Tr3(sourceText_Cstring, disambiguation_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) AddPositionalArgument3(name string, description string, syntax string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	syntax_ms := struct_miqt_string{}
	syntax_ms.data = CString(syntax)
	syntax_ms.len = size_t(len(syntax))
	defer free(unsafe.Pointer(syntax_ms.data))
	QCommandLineParser_AddPositionalArgument3(this.h, name_ms, description_ms, syntax_ms)
}

func (this *QCommandLineParser) ShowHelp1(exitCode int) {
	QCommandLineParser_ShowHelp1(this.h, (int)(exitCode))
}

func QCommandLineParser_ShowMessageAndExit3(message string, typeVal MessageType, exitCode int) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QCommandLineParser_ShowMessageAndExit3(message_ms, typeVal, (int)(exitCode))
}
