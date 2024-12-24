package qt

import (
	"unsafe"
)

type QIODeviceBase__OpenModeFlag int

const (
	QIODeviceBase__NotOpen      QIODeviceBase__OpenModeFlag = 0
	QIODeviceBase__ReadOnly     QIODeviceBase__OpenModeFlag = 1
	QIODeviceBase__WriteOnly    QIODeviceBase__OpenModeFlag = 2
	QIODeviceBase__ReadWrite    QIODeviceBase__OpenModeFlag = 3
	QIODeviceBase__Append       QIODeviceBase__OpenModeFlag = 4
	QIODeviceBase__Truncate     QIODeviceBase__OpenModeFlag = 8
	QIODeviceBase__Text         QIODeviceBase__OpenModeFlag = 16
	QIODeviceBase__Unbuffered   QIODeviceBase__OpenModeFlag = 32
	QIODeviceBase__NewOnly      QIODeviceBase__OpenModeFlag = 64
	QIODeviceBase__ExistingOnly QIODeviceBase__OpenModeFlag = 128
)

type QIODeviceBase struct {
	h          uintptr
	isSubclass bool
}
