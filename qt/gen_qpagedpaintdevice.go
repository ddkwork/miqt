package qt

import (
	"unsafe"
)

type QPagedPaintDevice__PdfVersion int

const (
	QPagedPaintDevice__PdfVersion_1_4 QPagedPaintDevice__PdfVersion = 0
	QPagedPaintDevice__PdfVersion_A1b QPagedPaintDevice__PdfVersion = 1
	QPagedPaintDevice__PdfVersion_1_6 QPagedPaintDevice__PdfVersion = 2
	QPagedPaintDevice__PdfVersion_X4  QPagedPaintDevice__PdfVersion = 3
)

type QPagedPaintDevice struct {
	h          uintptr
	isSubclass bool
}

func (this *QPagedPaintDevice) NewPage() bool {
	return (bool)(QPagedPaintDevice_NewPage(this.h))
}

func (this *QPagedPaintDevice) SetPageLayout(pageLayout *QPageLayout) bool {
	return (bool)(QPagedPaintDevice_SetPageLayout(this.h, pageLayout.cPointer()))
}

func (this *QPagedPaintDevice) SetPageSize(pageSize *QPageSize) bool {
	return (bool)(QPagedPaintDevice_SetPageSize(this.h, pageSize.cPointer()))
}

func (this *QPagedPaintDevice) SetPageOrientation(orientation QPageLayout__Orientation) bool {
	return (bool)(QPagedPaintDevice_SetPageOrientation(this.h, (int)(orientation)))
}

func (this *QPagedPaintDevice) SetPageMargins(margins *QMarginsF, units QPageLayout__Unit) bool {
	return (bool)(QPagedPaintDevice_SetPageMargins(this.h, margins.cPointer(), (int)(units)))
}

func (this *QPagedPaintDevice) PageLayout() *QPageLayout {
	_goptr := newQPageLayout(QPagedPaintDevice_PageLayout(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPagedPaintDevice) SetPageRanges(ranges *QPageRanges) {
	QPagedPaintDevice_SetPageRanges(this.h, ranges.cPointer())
}

func (this *QPagedPaintDevice) PageRanges() *QPageRanges {
	_goptr := newQPageRanges(QPagedPaintDevice_PageRanges(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
