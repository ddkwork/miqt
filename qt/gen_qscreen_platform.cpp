// +build ignore

#define WORKAROUND_INNER_CLASS_DEFINITION_QNativeInterface__QWindowsScreen
#include <qscreen_platform.h>
#include "gen_qscreen_platform.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQNativeInterfaceQWindowsScreen : public virtual QNativeInterface::QWindowsScreen {
public:

	MiqtVirtualQNativeInterfaceQWindowsScreen(): QNativeInterface::QWindowsScreen() {};

private:
	virtual ~MiqtVirtualQNativeInterfaceQWindowsScreen();

public:

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Handle = 0;

	// Subclass to allow providing a Go implementation
	virtual HMONITOR handle() const override {
		if (handle__Handle == 0) {
			return nullptr; // Pure virtual, there is no base we can call
		}
		

		struct HMONITOR__* callback_return_value = miqt_exec_callback_QNativeInterface__QWindowsScreen_Handle(const_cast<MiqtVirtualQNativeInterfaceQWindowsScreen*>(this), handle__Handle);

		return callback_return_value;
	}

};

QNativeInterface__QWindowsScreen* QNativeInterface__QWindowsScreen_new() {
	return new MiqtVirtualQNativeInterfaceQWindowsScreen();
}

struct HMONITOR__* QNativeInterface__QWindowsScreen_Handle(const QNativeInterface__QWindowsScreen* self) {
	HMONITOR _ret = self->handle();
	return static_cast<struct HMONITOR__*>(_ret);
}

void QNativeInterface__QWindowsScreen_override_virtual_Handle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQNativeInterfaceQWindowsScreen*>( (QNativeInterface__QWindowsScreen*)(self) )->handle__Handle = slot;
}

