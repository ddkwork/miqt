package printsupport

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QPrinter__PrinterMode int

const (
	QPrinter__ScreenResolution  QPrinter__PrinterMode = 0
	QPrinter__PrinterResolution QPrinter__PrinterMode = 1
	QPrinter__HighResolution    QPrinter__PrinterMode = 2
)

type QPrinter__PageOrder int

const (
	QPrinter__FirstPageFirst QPrinter__PageOrder = 0
	QPrinter__LastPageFirst  QPrinter__PageOrder = 1
)

type QPrinter__ColorMode int

const (
	QPrinter__GrayScale QPrinter__ColorMode = 0
	QPrinter__Color     QPrinter__ColorMode = 1
)

type QPrinter__PaperSource int

const (
	QPrinter__OnlyOne         QPrinter__PaperSource = 0
	QPrinter__Lower           QPrinter__PaperSource = 1
	QPrinter__Middle          QPrinter__PaperSource = 2
	QPrinter__Manual          QPrinter__PaperSource = 3
	QPrinter__Envelope        QPrinter__PaperSource = 4
	QPrinter__EnvelopeManual  QPrinter__PaperSource = 5
	QPrinter__Auto            QPrinter__PaperSource = 6
	QPrinter__Tractor         QPrinter__PaperSource = 7
	QPrinter__SmallFormat     QPrinter__PaperSource = 8
	QPrinter__LargeFormat     QPrinter__PaperSource = 9
	QPrinter__LargeCapacity   QPrinter__PaperSource = 10
	QPrinter__Cassette        QPrinter__PaperSource = 11
	QPrinter__FormSource      QPrinter__PaperSource = 12
	QPrinter__MaxPageSource   QPrinter__PaperSource = 13
	QPrinter__CustomSource    QPrinter__PaperSource = 14
	QPrinter__LastPaperSource QPrinter__PaperSource = 14
	QPrinter__Upper           QPrinter__PaperSource = 0
)

type QPrinter__PrinterState int

const (
	QPrinter__Idle    QPrinter__PrinterState = 0
	QPrinter__Active  QPrinter__PrinterState = 1
	QPrinter__Aborted QPrinter__PrinterState = 2
	QPrinter__Error   QPrinter__PrinterState = 3
)

type QPrinter__OutputFormat int

const (
	QPrinter__NativeFormat QPrinter__OutputFormat = 0
	QPrinter__PdfFormat    QPrinter__OutputFormat = 1
)

type QPrinter__PrintRange int

const (
	QPrinter__AllPages    QPrinter__PrintRange = 0
	QPrinter__Selection   QPrinter__PrintRange = 1
	QPrinter__PageRange   QPrinter__PrintRange = 2
	QPrinter__CurrentPage QPrinter__PrintRange = 3
)

type QPrinter__Unit int

const (
	QPrinter__Millimeter  QPrinter__Unit = 0
	QPrinter__Point       QPrinter__Unit = 1
	QPrinter__Inch        QPrinter__Unit = 2
	QPrinter__Pica        QPrinter__Unit = 3
	QPrinter__Didot       QPrinter__Unit = 4
	QPrinter__Cicero      QPrinter__Unit = 5
	QPrinter__DevicePixel QPrinter__Unit = 6
)

type QPrinter__DuplexMode int

const (
	QPrinter__DuplexNone      QPrinter__DuplexMode = 0
	QPrinter__DuplexAuto      QPrinter__DuplexMode = 1
	QPrinter__DuplexLongSide  QPrinter__DuplexMode = 2
	QPrinter__DuplexShortSide QPrinter__DuplexMode = 3
)

type QPrinter struct {
	h          uintptr
	isSubclass bool
}

// NewQPrinter constructs a new QPrinter object.
func NewQPrinter() *QPrinter {
	ret := newQPrinter(QPrinter_new())
	ret.isSubclass = true
	return ret
}

// NewQPrinter2 constructs a new QPrinter object.
func NewQPrinter2(printer *QPrinterInfo) *QPrinter {
	ret := newQPrinter(QPrinter_new2(printer.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPrinter3 constructs a new QPrinter object.
func NewQPrinter3(mode PrinterMode) *QPrinter {
	ret := newQPrinter(QPrinter_new3(mode))
	ret.isSubclass = true
	return ret
}

// NewQPrinter4 constructs a new QPrinter object.
func NewQPrinter4(printer *QPrinterInfo, mode PrinterMode) *QPrinter {
	ret := newQPrinter(QPrinter_new4(printer.cPointer(), mode))
	ret.isSubclass = true
	return ret
}

func (this *QPrinter) DevType() int {
	return (int)(QPrinter_DevType(this.h))
}

func (this *QPrinter) SetOutputFormat(format OutputFormat) {
	QPrinter_SetOutputFormat(this.h, format)
}

func (this *QPrinter) OutputFormat() OutputFormat {
	xxxxxxxxx
}

func (this *QPrinter) SetPdfVersion(version PdfVersion) {
	QPrinter_SetPdfVersion(this.h, version)
}

func (this *QPrinter) PdfVersion() PdfVersion {
	xxxxxxxxx
}

func (this *QPrinter) SetPrinterName(printerName string) {
	printerName_ms := struct_miqt_string{}
	printerName_ms.data = CString(printerName)
	printerName_ms.len = size_t(len(printerName))
	defer free(unsafe.Pointer(printerName_ms.data))
	QPrinter_SetPrinterName(this.h, printerName_ms)
}

func (this *QPrinter) PrinterName() string {
	var _ms struct_miqt_string = QPrinter_PrinterName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinter) IsValid() bool {
	return (bool)(QPrinter_IsValid(this.h))
}

func (this *QPrinter) SetOutputFileName(outputFileName string) {
	outputFileName_ms := struct_miqt_string{}
	outputFileName_ms.data = CString(outputFileName)
	outputFileName_ms.len = size_t(len(outputFileName))
	defer free(unsafe.Pointer(outputFileName_ms.data))
	QPrinter_SetOutputFileName(this.h, outputFileName_ms)
}

func (this *QPrinter) OutputFileName() string {
	var _ms struct_miqt_string = QPrinter_OutputFileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinter) SetPrintProgram(printProgram string) {
	printProgram_ms := struct_miqt_string{}
	printProgram_ms.data = CString(printProgram)
	printProgram_ms.len = size_t(len(printProgram))
	defer free(unsafe.Pointer(printProgram_ms.data))
	QPrinter_SetPrintProgram(this.h, printProgram_ms)
}

func (this *QPrinter) PrintProgram() string {
	var _ms struct_miqt_string = QPrinter_PrintProgram(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinter) SetDocName(docName string) {
	docName_ms := struct_miqt_string{}
	docName_ms.data = CString(docName)
	docName_ms.len = size_t(len(docName))
	defer free(unsafe.Pointer(docName_ms.data))
	QPrinter_SetDocName(this.h, docName_ms)
}

func (this *QPrinter) DocName() string {
	var _ms struct_miqt_string = QPrinter_DocName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinter) SetCreator(creator string) {
	creator_ms := struct_miqt_string{}
	creator_ms.data = CString(creator)
	creator_ms.len = size_t(len(creator))
	defer free(unsafe.Pointer(creator_ms.data))
	QPrinter_SetCreator(this.h, creator_ms)
}

func (this *QPrinter) Creator() string {
	var _ms struct_miqt_string = QPrinter_Creator(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinter) SetPageOrder(pageOrder PageOrder) {
	QPrinter_SetPageOrder(this.h, pageOrder)
}

func (this *QPrinter) PageOrder() PageOrder {
	xxxxxxxxx
}

func (this *QPrinter) SetResolution(resolution int) {
	QPrinter_SetResolution(this.h, (int)(resolution))
}

func (this *QPrinter) Resolution() int {
	return (int)(QPrinter_Resolution(this.h))
}

func (this *QPrinter) SetColorMode(colorMode ColorMode) {
	QPrinter_SetColorMode(this.h, colorMode)
}

func (this *QPrinter) ColorMode() ColorMode {
	xxxxxxxxx
}

func (this *QPrinter) SetCollateCopies(collate bool) {
	QPrinter_SetCollateCopies(this.h, (bool)(collate))
}

func (this *QPrinter) CollateCopies() bool {
	return (bool)(QPrinter_CollateCopies(this.h))
}

func (this *QPrinter) SetFullPage(fullPage bool) {
	QPrinter_SetFullPage(this.h, (bool)(fullPage))
}

func (this *QPrinter) FullPage() bool {
	return (bool)(QPrinter_FullPage(this.h))
}

func (this *QPrinter) SetCopyCount(copyCount int) {
	QPrinter_SetCopyCount(this.h, (int)(copyCount))
}

func (this *QPrinter) CopyCount() int {
	return (int)(QPrinter_CopyCount(this.h))
}

func (this *QPrinter) SupportsMultipleCopies() bool {
	return (bool)(QPrinter_SupportsMultipleCopies(this.h))
}

func (this *QPrinter) SetPaperSource(paperSource PaperSource) {
	QPrinter_SetPaperSource(this.h, paperSource)
}

func (this *QPrinter) PaperSource() PaperSource {
	xxxxxxxxx
}

func (this *QPrinter) SetDuplex(duplex DuplexMode) {
	QPrinter_SetDuplex(this.h, duplex)
}

func (this *QPrinter) Duplex() DuplexMode {
	xxxxxxxxx
}

func (this *QPrinter) SupportedResolutions() []int {
	var _ma struct_miqt_array = QPrinter_SupportedResolutions(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QPrinter) SupportedPaperSources() []PaperSource {
	var _ma struct_miqt_array = QPrinter_SupportedPaperSources(this.h)
	_ret := make([]PaperSource, int(_ma.len))
	_outCast := (*[0xffff]PaperSource)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QPrinter) SetFontEmbeddingEnabled(enable bool) {
	QPrinter_SetFontEmbeddingEnabled(this.h, (bool)(enable))
}

func (this *QPrinter) FontEmbeddingEnabled() bool {
	return (bool)(QPrinter_FontEmbeddingEnabled(this.h))
}

func (this *QPrinter) PaperRect(param1 Unit) *qt.QRectF {
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QPrinter_PaperRect(this.h, param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPrinter) PageRect(param1 Unit) *qt.QRectF {
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QPrinter_PageRect(this.h, param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPrinter) PrinterSelectionOption() string {
	var _ms struct_miqt_string = QPrinter_PrinterSelectionOption(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinter) SetPrinterSelectionOption(printerSelectionOption string) {
	printerSelectionOption_ms := struct_miqt_string{}
	printerSelectionOption_ms.data = CString(printerSelectionOption)
	printerSelectionOption_ms.len = size_t(len(printerSelectionOption))
	defer free(unsafe.Pointer(printerSelectionOption_ms.data))
	QPrinter_SetPrinterSelectionOption(this.h, printerSelectionOption_ms)
}

func (this *QPrinter) NewPage() bool {
	return (bool)(QPrinter_NewPage(this.h))
}

func (this *QPrinter) Abort() bool {
	return (bool)(QPrinter_Abort(this.h))
}

func (this *QPrinter) PrinterState() PrinterState {
	xxxxxxxxx
}

func (this *QPrinter) PaintEngine() *qt.QPaintEngine {
	return qt.UnsafeNewQPaintEngine(unsafe.Pointer(QPrinter_PaintEngine(this.h)))
}

func (this *QPrinter) PrintEngine() *QPrintEngine {
	return newQPrintEngine(QPrinter_PrintEngine(this.h))
}

func (this *QPrinter) SetFromTo(fromPage int, toPage int) {
	QPrinter_SetFromTo(this.h, (int)(fromPage), (int)(toPage))
}

func (this *QPrinter) FromPage() int {
	return (int)(QPrinter_FromPage(this.h))
}

func (this *QPrinter) ToPage() int {
	return (int)(QPrinter_ToPage(this.h))
}

func (this *QPrinter) SetPrintRange(rangeVal PrintRange) {
	QPrinter_SetPrintRange(this.h, rangeVal)
}

func (this *QPrinter) PrintRange() PrintRange {
	xxxxxxxxx
}
