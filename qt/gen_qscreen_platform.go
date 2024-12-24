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
			func (this *QNativeInterface__QWindowsScreen) OnHandle(slot func() *struct HMONITOR__) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QNativeInterface__QWindowsScreen_override_virtual_Handle(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QNativeInterface__QWindowsScreen_Handle
				func miqt_exec_callback_QNativeInterface__QWindowsScreen_Handle(self QNativeInterface__QWindowsScreen, cb intptr_t) *struct_HMONITOR__{
					gofunc, ok := cgo.Handle(cb).Value().(func() *struct HMONITOR__)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc( )

return virtualReturn

				}
			