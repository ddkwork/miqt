package qt
import (
	"unsafe"
)

		type QNativeInterface__QWindowsScreen struct {
			h uintptr
			isSubclass bool
}
		
			// NewQNativeInterface__QWindowsScreen constructs a new QNativeInterface::QWindowsScreen object.
			func NewQNativeInterface__QWindowsScreen() *QNativeInterface__QWindowsScreen {
								
				ret := newQNativeInterface__QWindowsScreen(QNativeInterface__QWindowsScreen_new())
				ret.isSubclass = true
				return ret
			}
			
			
			func (this *QNativeInterface__QWindowsScreen) Handle() *struct HMONITOR__ {
				return (*struct HMONITOR__)(unsafe.Pointer(QNativeInterface__QWindowsScreen_Handle(this.h)))
}
			