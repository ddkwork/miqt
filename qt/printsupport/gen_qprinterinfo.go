package printsupport

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QPrinterInfo struct {
	h          uintptr
	isSubclass bool
}

// NewQPrinterInfo constructs a new QPrinterInfo object.
func NewQPrinterInfo() *QPrinterInfo {
	g := newQPrinterInfo(QPrinterInfo_new())
	g.isSubclass = true
	return g
}

// NewQPrinterInfo2 constructs a new QPrinterInfo object.
func NewQPrinterInfo2(other *QPrinterInfo) *QPrinterInfo {
	g := newQPrinterInfo(QPrinterInfo_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPrinterInfo3 constructs a new QPrinterInfo object.
func NewQPrinterInfo3(printer *QPrinter) *QPrinterInfo {
	g := newQPrinterInfo(QPrinterInfo_new3(printer.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPrinterInfo) OperatorAssign(other *QPrinterInfo) {
	QPrinterInfo_OperatorAssign(this.h, other.cPointer())
}

func (this *QPrinterInfo) PrinterName() string {
	var _ms struct_miqt_string = QPrinterInfo_PrinterName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinterInfo) Description() string {
	var _ms struct_miqt_string = QPrinterInfo_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinterInfo) Location() string {
	var _ms struct_miqt_string = QPrinterInfo_Location(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinterInfo) MakeAndModel() string {
	var _ms struct_miqt_string = QPrinterInfo_MakeAndModel(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinterInfo) IsNull() bool {
	return (bool)(QPrinterInfo_IsNull(this.h))
}

func (this *QPrinterInfo) IsDefault() bool {
	return (bool)(QPrinterInfo_IsDefault(this.h))
}

func (this *QPrinterInfo) IsRemote() bool {
	return (bool)(QPrinterInfo_IsRemote(this.h))
}

func (this *QPrinterInfo) State() QPrinter__PrinterState {
	return (QPrinter__PrinterState)(QPrinterInfo_State(this.h))
}

func (this *QPrinterInfo) SupportedPageSizes() []qt.QPageSize {
	var _ma struct_miqt_array = QPrinterInfo_SupportedPageSizes(this.h)
	_ret := make([]qt.QPageSize, int(_ma.len))
	_outCast := (*[0xffff]*QPageSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := qt.UnsafeNewQPageSize(unsafe.Pointer(_outCast[i]))
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QPrinterInfo) DefaultPageSize() *qt.QPageSize {
	_goptr := qt.UnsafeNewQPageSize(unsafe.Pointer(QPrinterInfo_DefaultPageSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPrinterInfo) SupportsCustomPageSizes() bool {
	return (bool)(QPrinterInfo_SupportsCustomPageSizes(this.h))
}

func (this *QPrinterInfo) MinimumPhysicalPageSize() *qt.QPageSize {
	_goptr := qt.UnsafeNewQPageSize(unsafe.Pointer(QPrinterInfo_MinimumPhysicalPageSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPrinterInfo) MaximumPhysicalPageSize() *qt.QPageSize {
	_goptr := qt.UnsafeNewQPageSize(unsafe.Pointer(QPrinterInfo_MaximumPhysicalPageSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPrinterInfo) SupportedResolutions() []int {
	var _ma struct_miqt_array = QPrinterInfo_SupportedResolutions(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QPrinterInfo) DefaultDuplexMode() QPrinter__DuplexMode {
	return (QPrinter__DuplexMode)(QPrinterInfo_DefaultDuplexMode(this.h))
}

func (this *QPrinterInfo) SupportedDuplexModes() []QPrinter__DuplexMode {
	var _ma struct_miqt_array = QPrinterInfo_SupportedDuplexModes(this.h)
	_ret := make([]QPrinter__DuplexMode, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QPrinter__DuplexMode)(_outCast[i])
	}
	return _ret
}

func (this *QPrinterInfo) DefaultColorMode() QPrinter__ColorMode {
	return (QPrinter__ColorMode)(QPrinterInfo_DefaultColorMode(this.h))
}

func (this *QPrinterInfo) SupportedColorModes() []QPrinter__ColorMode {
	var _ma struct_miqt_array = QPrinterInfo_SupportedColorModes(this.h)
	_ret := make([]QPrinter__ColorMode, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QPrinter__ColorMode)(_outCast[i])
	}
	return _ret
}

func QPrinterInfo_AvailablePrinterNames() []string {
	var _ma struct_miqt_array = QPrinterInfo_AvailablePrinterNames()
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

func QPrinterInfo_AvailablePrinters() []QPrinterInfo {
	var _ma struct_miqt_array = QPrinterInfo_AvailablePrinters()
	_ret := make([]QPrinterInfo, int(_ma.len))
	_outCast := (*[0xffff]*QPrinterInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQPrinterInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QPrinterInfo_DefaultPrinterName() string {
	var _ms struct_miqt_string = QPrinterInfo_DefaultPrinterName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPrinterInfo_DefaultPrinter() *QPrinterInfo {
	_goptr := newQPrinterInfo(QPrinterInfo_DefaultPrinter())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPrinterInfo_PrinterInfo(printerName string) *QPrinterInfo {
	printerName_ms := struct_miqt_string{}
	printerName_ms.data = CString(printerName)
	printerName_ms.len = size_t(len(printerName))
	defer free(unsafe.Pointer(printerName_ms.data))
	_goptr := newQPrinterInfo(QPrinterInfo_PrinterInfo(printerName_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
