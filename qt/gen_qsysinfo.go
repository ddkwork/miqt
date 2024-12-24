package qt

import (
	"unsafe"
)

type QSysInfo__Sizes int

const (
	QSysInfo__WordSize QSysInfo__Sizes = 64
)

type QSysInfo__Endian int

const (
	QSysInfo__BigEndian    QSysInfo__Endian = 0
	QSysInfo__LittleEndian QSysInfo__Endian = 1
	QSysInfo__ByteOrder    QSysInfo__Endian = 1
)

type QSysInfo struct {
	h          uintptr
	isSubclass bool
}

func QSysInfo_BuildCpuArchitecture() string {
	var _ms struct_miqt_string = QSysInfo_BuildCpuArchitecture()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_CurrentCpuArchitecture() string {
	var _ms struct_miqt_string = QSysInfo_CurrentCpuArchitecture()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_BuildAbi() string {
	var _ms struct_miqt_string = QSysInfo_BuildAbi()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_KernelType() string {
	var _ms struct_miqt_string = QSysInfo_KernelType()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_KernelVersion() string {
	var _ms struct_miqt_string = QSysInfo_KernelVersion()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_ProductType() string {
	var _ms struct_miqt_string = QSysInfo_ProductType()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_ProductVersion() string {
	var _ms struct_miqt_string = QSysInfo_ProductVersion()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_PrettyProductName() string {
	var _ms struct_miqt_string = QSysInfo_PrettyProductName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_MachineHostName() string {
	var _ms struct_miqt_string = QSysInfo_MachineHostName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSysInfo_MachineUniqueId() []byte {
	var _bytearray struct_miqt_string = QSysInfo_MachineUniqueId()
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QSysInfo_BootUniqueId() []byte {
	var _bytearray struct_miqt_string = QSysInfo_BootUniqueId()
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}
