// +build ignore

#define WORKAROUND_INNER_CLASS_DEFINITION_QNativeInterface__QWindowsScreen
#include <qscreen_platform.h>
#include "gen_qscreen_platform.h"

QNativeInterface__QWindowsScreen* QNativeInterface__QWindowsScreen_new() {
	return new QNativeInterface::QWindowsScreen();
}

struct HMONITOR__* QNativeInterface__QWindowsScreen_Handle(const QNativeInterface__QWindowsScreen* self) {
	HMONITOR _ret = self->handle();
	return static_cast<struct HMONITOR__*>(_ret);
}

