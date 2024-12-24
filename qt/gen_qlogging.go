package qt
import (
	"unsafe"
)

		type QtMsgType int
		const (
QtMsgType__QtDebugMsg QtMsgType = 0
QtMsgType__QtWarningMsg QtMsgType = 1
QtMsgType__QtCriticalMsg QtMsgType = 2
QtMsgType__QtFatalMsg QtMsgType = 3
QtMsgType__QtInfoMsg QtMsgType = 4
QtMsgType__QtSystemMsg QtMsgType = 2

)


		type QMessageLogContext struct {
			h uintptr
			isSubclass bool
}
		
			// NewQMessageLogContext constructs a new QMessageLogContext object.
			func NewQMessageLogContext() *QMessageLogContext {
								
				ret := newQMessageLogContext(QMessageLogContext_new())
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQMessageLogContext2 constructs a new QMessageLogContext object.
			func NewQMessageLogContext2(fileName string, lineNumber int, functionName string, categoryName string) *QMessageLogContext {
				fileName_Cstring := CString(fileName)
defer free(unsafe.Pointer(fileName_Cstring))
functionName_Cstring := CString(functionName)
defer free(unsafe.Pointer(functionName_Cstring))
categoryName_Cstring := CString(categoryName)
defer free(unsafe.Pointer(categoryName_Cstring))
				
				ret := newQMessageLogContext(QMessageLogContext_new2(fileName_Cstring, (int)(lineNumber), functionName_Cstring, categoryName_Cstring))
				ret.isSubclass = true
				return ret
			}
			
			
		type QMessageLogger struct {
			h uintptr
			isSubclass bool
}
		
			// NewQMessageLogger constructs a new QMessageLogger object.
			func NewQMessageLogger() *QMessageLogger {
								
				ret := newQMessageLogger(QMessageLogger_new())
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQMessageLogger2 constructs a new QMessageLogger object.
			func NewQMessageLogger2(file string, line int, function string) *QMessageLogger {
				file_Cstring := CString(file)
defer free(unsafe.Pointer(file_Cstring))
function_Cstring := CString(function)
defer free(unsafe.Pointer(function_Cstring))
				
				ret := newQMessageLogger(QMessageLogger_new2(file_Cstring, (int)(line), function_Cstring))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQMessageLogger3 constructs a new QMessageLogger object.
			func NewQMessageLogger3(file string, line int, function string, category string) *QMessageLogger {
				file_Cstring := CString(file)
defer free(unsafe.Pointer(file_Cstring))
function_Cstring := CString(function)
defer free(unsafe.Pointer(function_Cstring))
category_Cstring := CString(category)
defer free(unsafe.Pointer(category_Cstring))
				
				ret := newQMessageLogger(QMessageLogger_new3(file_Cstring, (int)(line), function_Cstring, category_Cstring))
				ret.isSubclass = true
				return ret
			}
			
			
			func (this *QMessageLogger) Debug(msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Debug(this.h, msg_Cstring, )
}
			
			func (this *QMessageLogger) NoDebug(param1 string,  ...)  {
				param1_Cstring := CString(param1)
defer free(unsafe.Pointer(param1_Cstring))
 QMessageLogger_NoDebug(this.h, param1_Cstring, )
}
			
			func (this *QMessageLogger) Info(msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Info(this.h, msg_Cstring, )
}
			
			func (this *QMessageLogger) Warning(msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Warning(this.h, msg_Cstring, )
}
			
			func (this *QMessageLogger) Critical(msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Critical(this.h, msg_Cstring, )
}
			
			func (this *QMessageLogger) Fatal(msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Fatal(this.h, msg_Cstring, )
}
			
			func (this *QMessageLogger) Debug2(cat *QLoggingCategory, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Debug2(this.h, cat.cPointer(), msg_Cstring, )
}
			
			func (this *QMessageLogger) Debug3(catFunc CategoryFunction, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Debug3(this.h, catFunc, msg_Cstring, )
}
			
			func (this *QMessageLogger) Info2(cat *QLoggingCategory, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Info2(this.h, cat.cPointer(), msg_Cstring, )
}
			
			func (this *QMessageLogger) Info3(catFunc CategoryFunction, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Info3(this.h, catFunc, msg_Cstring, )
}
			
			func (this *QMessageLogger) Warning2(cat *QLoggingCategory, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Warning2(this.h, cat.cPointer(), msg_Cstring, )
}
			
			func (this *QMessageLogger) Warning3(catFunc CategoryFunction, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Warning3(this.h, catFunc, msg_Cstring, )
}
			
			func (this *QMessageLogger) Critical2(cat *QLoggingCategory, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Critical2(this.h, cat.cPointer(), msg_Cstring, )
}
			
			func (this *QMessageLogger) Critical3(catFunc CategoryFunction, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Critical3(this.h, catFunc, msg_Cstring, )
}
			
			func (this *QMessageLogger) Fatal2(cat *QLoggingCategory, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Fatal2(this.h, cat.cPointer(), msg_Cstring, )
}
			
			func (this *QMessageLogger) Fatal3(catFunc CategoryFunction, msg string,  ...)  {
				msg_Cstring := CString(msg)
defer free(unsafe.Pointer(msg_Cstring))
 QMessageLogger_Fatal3(this.h, catFunc, msg_Cstring, )
}
			
			func (this *QMessageLogger) Debug4() *QDebug {
				_goptr := newQDebug(QMessageLogger_Debug4(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) DebugWithCat(cat *QLoggingCategory) *QDebug {
				_goptr := newQDebug(QMessageLogger_DebugWithCat(this.h, cat.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) DebugWithCatFunc(catFunc CategoryFunction) *QDebug {
				_goptr := newQDebug(QMessageLogger_DebugWithCatFunc(this.h, catFunc))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) Info4() *QDebug {
				_goptr := newQDebug(QMessageLogger_Info4(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) InfoWithCat(cat *QLoggingCategory) *QDebug {
				_goptr := newQDebug(QMessageLogger_InfoWithCat(this.h, cat.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) InfoWithCatFunc(catFunc CategoryFunction) *QDebug {
				_goptr := newQDebug(QMessageLogger_InfoWithCatFunc(this.h, catFunc))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) Warning4() *QDebug {
				_goptr := newQDebug(QMessageLogger_Warning4(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) WarningWithCat(cat *QLoggingCategory) *QDebug {
				_goptr := newQDebug(QMessageLogger_WarningWithCat(this.h, cat.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) WarningWithCatFunc(catFunc CategoryFunction) *QDebug {
				_goptr := newQDebug(QMessageLogger_WarningWithCatFunc(this.h, catFunc))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) Critical4() *QDebug {
				_goptr := newQDebug(QMessageLogger_Critical4(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) CriticalWithCat(cat *QLoggingCategory) *QDebug {
				_goptr := newQDebug(QMessageLogger_CriticalWithCat(this.h, cat.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) CriticalWithCatFunc(catFunc CategoryFunction) *QDebug {
				_goptr := newQDebug(QMessageLogger_CriticalWithCatFunc(this.h, catFunc))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) Fatal4() *QDebug {
				_goptr := newQDebug(QMessageLogger_Fatal4(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) FatalWithCat(cat *QLoggingCategory) *QDebug {
				_goptr := newQDebug(QMessageLogger_FatalWithCat(this.h, cat.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) FatalWithCatFunc(catFunc CategoryFunction) *QDebug {
				_goptr := newQDebug(QMessageLogger_FatalWithCatFunc(this.h, catFunc))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageLogger) NoDebug2() *QNoDebug {
				_goptr := newQNoDebug(QMessageLogger_NoDebug2(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			