// +build ignore

#include <QList>
#include <QMarginsF>
#include <QPageLayout>
#include <QPageRanges>
#include <QPageSize>
#include <QPagedPaintDevice>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPrintEngine>
#include <QPrinter>
#include <QPrinterInfo>
#include <QRectF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qprinter.h>
#include "gen_qprinter.h"

#ifndef _Bool
#define _Bool bool
#endif

class MiqtVirtualQPrinter : public virtual QPrinter {
public:

	MiqtVirtualQPrinter(): QPrinter() {};
	MiqtVirtualQPrinter(const QPrinterInfo& printer): QPrinter(printer) {};
	MiqtVirtualQPrinter(PrinterMode mode): QPrinter(mode) {};
	MiqtVirtualQPrinter(const QPrinterInfo& printer, PrinterMode mode): QPrinter(printer, mode) {};

	virtual ~MiqtVirtualQPrinter() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DevType = 0;

	// Subclass to allow providing a Go implementation
	virtual int devType() const override {
		if (handle__DevType == 0) {
			return QPrinter::devType();
		}
		

		int callback_return_value = miqt_exec_callback_QPrinter_DevType(const_cast<MiqtVirtualQPrinter*>(this), handle__DevType);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_DevType() const {

		return QPrinter::devType();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__NewPage = 0;

	// Subclass to allow providing a Go implementation
	virtual bool newPage() override {
		if (handle__NewPage == 0) {
			return QPrinter::newPage();
		}
		

		bool callback_return_value = miqt_exec_callback_QPrinter_NewPage(this, handle__NewPage);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_NewPage() {

		return QPrinter::newPage();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEngine = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintEngine* paintEngine() const override {
		if (handle__PaintEngine == 0) {
			return QPrinter::paintEngine();
		}
		

		QPaintEngine* callback_return_value = miqt_exec_callback_QPrinter_PaintEngine(const_cast<MiqtVirtualQPrinter*>(this), handle__PaintEngine);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintEngine* virtualbase_PaintEngine() const {

		return QPrinter::paintEngine();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metric = 0;

	// Subclass to allow providing a Go implementation
	virtual int metric(PaintDeviceMetric param1) const override {
		if (handle__Metric == 0) {
			return QPrinter::metric(param1);
		}
		
		PaintDeviceMetric sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QPrinter_Metric(const_cast<MiqtVirtualQPrinter*>(this), handle__Metric, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Metric(PaintDeviceMetric param1) const {

		return QPrinter::metric(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetPageLayout = 0;

	// Subclass to allow providing a Go implementation
	virtual bool setPageLayout(const QPageLayout& pageLayout) override {
		if (handle__SetPageLayout == 0) {
			return QPrinter::setPageLayout(pageLayout);
		}
		
		const QPageLayout& pageLayout_ret = pageLayout;
		// Cast returned reference into pointer
		QPageLayout* sigval1 = const_cast<QPageLayout*>(&pageLayout_ret);

		bool callback_return_value = miqt_exec_callback_QPrinter_SetPageLayout(this, handle__SetPageLayout, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_SetPageLayout(QPageLayout* pageLayout) {

		return QPrinter::setPageLayout(*pageLayout);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetPageSize = 0;

	// Subclass to allow providing a Go implementation
	virtual bool setPageSize(const QPageSize& pageSize) override {
		if (handle__SetPageSize == 0) {
			return QPrinter::setPageSize(pageSize);
		}
		
		const QPageSize& pageSize_ret = pageSize;
		// Cast returned reference into pointer
		QPageSize* sigval1 = const_cast<QPageSize*>(&pageSize_ret);

		bool callback_return_value = miqt_exec_callback_QPrinter_SetPageSize(this, handle__SetPageSize, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_SetPageSize(QPageSize* pageSize) {

		return QPrinter::setPageSize(*pageSize);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetPageOrientation = 0;

	// Subclass to allow providing a Go implementation
	virtual bool setPageOrientation(QPageLayout::Orientation orientation) override {
		if (handle__SetPageOrientation == 0) {
			return QPrinter::setPageOrientation(orientation);
		}
		
		QPageLayout::Orientation orientation_ret = orientation;
		int sigval1 = static_cast<int>(orientation_ret);

		bool callback_return_value = miqt_exec_callback_QPrinter_SetPageOrientation(this, handle__SetPageOrientation, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_SetPageOrientation(int orientation) {

		return QPrinter::setPageOrientation(static_cast<QPageLayout::Orientation>(orientation));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetPageMargins = 0;

	// Subclass to allow providing a Go implementation
	virtual bool setPageMargins(const QMarginsF& margins, QPageLayout::Unit units) override {
		if (handle__SetPageMargins == 0) {
			return QPrinter::setPageMargins(margins, units);
		}
		
		const QMarginsF& margins_ret = margins;
		// Cast returned reference into pointer
		QMarginsF* sigval1 = const_cast<QMarginsF*>(&margins_ret);
		QPageLayout::Unit units_ret = units;
		int sigval2 = static_cast<int>(units_ret);

		bool callback_return_value = miqt_exec_callback_QPrinter_SetPageMargins(this, handle__SetPageMargins, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_SetPageMargins(QMarginsF* margins, int units) {

		return QPrinter::setPageMargins(*margins, static_cast<QPageLayout::Unit>(units));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetPageRanges = 0;

	// Subclass to allow providing a Go implementation
	virtual void setPageRanges(const QPageRanges& ranges) override {
		if (handle__SetPageRanges == 0) {
			QPrinter::setPageRanges(ranges);
			return;
		}
		
		const QPageRanges& ranges_ret = ranges;
		// Cast returned reference into pointer
		QPageRanges* sigval1 = const_cast<QPageRanges*>(&ranges_ret);

		miqt_exec_callback_QPrinter_SetPageRanges(this, handle__SetPageRanges, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetPageRanges(QPageRanges* ranges) {

		QPrinter::setPageRanges(*ranges);

	}

};

QPrinter* QPrinter_new() {
	return new MiqtVirtualQPrinter();
}

QPrinter* QPrinter_new2(QPrinterInfo* printer) {
	return new MiqtVirtualQPrinter(*printer);
}

QPrinter* QPrinter_new3(PrinterMode mode) {
	return new MiqtVirtualQPrinter(mode);
}

QPrinter* QPrinter_new4(QPrinterInfo* printer, PrinterMode mode) {
	return new MiqtVirtualQPrinter(*printer, mode);
}

void QPrinter_virtbase(QPrinter* src, QPagedPaintDevice** outptr_QPagedPaintDevice) {
	*outptr_QPagedPaintDevice = static_cast<QPagedPaintDevice*>(src);
}

int QPrinter_DevType(const QPrinter* self) {
	return self->devType();
}

void QPrinter_SetOutputFormat(QPrinter* self, OutputFormat format) {
	self->setOutputFormat(format);
}

OutputFormat QPrinter_OutputFormat(const QPrinter* self) {
	return self->outputFormat();
}

void QPrinter_SetPdfVersion(QPrinter* self, PdfVersion version) {
	self->setPdfVersion(version);
}

PdfVersion QPrinter_PdfVersion(const QPrinter* self) {
	return self->pdfVersion();
}

void QPrinter_SetPrinterName(QPrinter* self, struct miqt_string printerName) {
	QString printerName_QString = QString::fromUtf8(printerName.data, printerName.len);
	self->setPrinterName(printerName_QString);
}

struct miqt_string QPrinter_PrinterName(const QPrinter* self) {
	QString _ret = self->printerName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QPrinter_IsValid(const QPrinter* self) {
	return self->isValid();
}

void QPrinter_SetOutputFileName(QPrinter* self, struct miqt_string outputFileName) {
	QString outputFileName_QString = QString::fromUtf8(outputFileName.data, outputFileName.len);
	self->setOutputFileName(outputFileName_QString);
}

struct miqt_string QPrinter_OutputFileName(const QPrinter* self) {
	QString _ret = self->outputFileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPrinter_SetPrintProgram(QPrinter* self, struct miqt_string printProgram) {
	QString printProgram_QString = QString::fromUtf8(printProgram.data, printProgram.len);
	self->setPrintProgram(printProgram_QString);
}

struct miqt_string QPrinter_PrintProgram(const QPrinter* self) {
	QString _ret = self->printProgram();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPrinter_SetDocName(QPrinter* self, struct miqt_string docName) {
	QString docName_QString = QString::fromUtf8(docName.data, docName.len);
	self->setDocName(docName_QString);
}

struct miqt_string QPrinter_DocName(const QPrinter* self) {
	QString _ret = self->docName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPrinter_SetCreator(QPrinter* self, struct miqt_string creator) {
	QString creator_QString = QString::fromUtf8(creator.data, creator.len);
	self->setCreator(creator_QString);
}

struct miqt_string QPrinter_Creator(const QPrinter* self) {
	QString _ret = self->creator();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPrinter_SetPageOrder(QPrinter* self, PageOrder pageOrder) {
	self->setPageOrder(pageOrder);
}

PageOrder QPrinter_PageOrder(const QPrinter* self) {
	return self->pageOrder();
}

void QPrinter_SetResolution(QPrinter* self, int resolution) {
	self->setResolution(static_cast<int>(resolution));
}

int QPrinter_Resolution(const QPrinter* self) {
	return self->resolution();
}

void QPrinter_SetColorMode(QPrinter* self, ColorMode colorMode) {
	self->setColorMode(colorMode);
}

ColorMode QPrinter_ColorMode(const QPrinter* self) {
	return self->colorMode();
}

void QPrinter_SetCollateCopies(QPrinter* self, bool collate) {
	self->setCollateCopies(collate);
}

bool QPrinter_CollateCopies(const QPrinter* self) {
	return self->collateCopies();
}

void QPrinter_SetFullPage(QPrinter* self, bool fullPage) {
	self->setFullPage(fullPage);
}

bool QPrinter_FullPage(const QPrinter* self) {
	return self->fullPage();
}

void QPrinter_SetCopyCount(QPrinter* self, int copyCount) {
	self->setCopyCount(static_cast<int>(copyCount));
}

int QPrinter_CopyCount(const QPrinter* self) {
	return self->copyCount();
}

bool QPrinter_SupportsMultipleCopies(const QPrinter* self) {
	return self->supportsMultipleCopies();
}

void QPrinter_SetPaperSource(QPrinter* self, PaperSource paperSource) {
	self->setPaperSource(paperSource);
}

PaperSource QPrinter_PaperSource(const QPrinter* self) {
	return self->paperSource();
}

void QPrinter_SetDuplex(QPrinter* self, DuplexMode duplex) {
	self->setDuplex(duplex);
}

DuplexMode QPrinter_Duplex(const QPrinter* self) {
	return self->duplex();
}

struct miqt_array /* of int */  QPrinter_SupportedResolutions(const QPrinter* self) {
	QList<int> _ret = self->supportedResolutions();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* _arr = static_cast<int*>(malloc(sizeof(int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of PaperSource */  QPrinter_SupportedPaperSources(const QPrinter* self) {
	QList<PaperSource> _ret = self->supportedPaperSources();
	// Convert QList<> from C++ memory to manually-managed C memory
	PaperSource* _arr = static_cast<PaperSource*>(malloc(sizeof(PaperSource) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QPrinter_SetFontEmbeddingEnabled(QPrinter* self, bool enable) {
	self->setFontEmbeddingEnabled(enable);
}

bool QPrinter_FontEmbeddingEnabled(const QPrinter* self) {
	return self->fontEmbeddingEnabled();
}

QRectF* QPrinter_PaperRect(const QPrinter* self, Unit param1) {
	return new QRectF(self->paperRect(param1));
}

QRectF* QPrinter_PageRect(const QPrinter* self, Unit param1) {
	return new QRectF(self->pageRect(param1));
}

struct miqt_string QPrinter_PrinterSelectionOption(const QPrinter* self) {
	QString _ret = self->printerSelectionOption();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPrinter_SetPrinterSelectionOption(QPrinter* self, struct miqt_string printerSelectionOption) {
	QString printerSelectionOption_QString = QString::fromUtf8(printerSelectionOption.data, printerSelectionOption.len);
	self->setPrinterSelectionOption(printerSelectionOption_QString);
}

bool QPrinter_NewPage(QPrinter* self) {
	return self->newPage();
}

bool QPrinter_Abort(QPrinter* self) {
	return self->abort();
}

PrinterState QPrinter_PrinterState(const QPrinter* self) {
	return self->printerState();
}

QPaintEngine* QPrinter_PaintEngine(const QPrinter* self) {
	return self->paintEngine();
}

QPrintEngine* QPrinter_PrintEngine(const QPrinter* self) {
	return self->printEngine();
}

void QPrinter_SetFromTo(QPrinter* self, int fromPage, int toPage) {
	self->setFromTo(static_cast<int>(fromPage), static_cast<int>(toPage));
}

int QPrinter_FromPage(const QPrinter* self) {
	return self->fromPage();
}

int QPrinter_ToPage(const QPrinter* self) {
	return self->toPage();
}

void QPrinter_SetPrintRange(QPrinter* self, PrintRange rangeVal) {
	self->setPrintRange(rangeVal);
}

PrintRange QPrinter_PrintRange(const QPrinter* self) {
	return self->printRange();
}

void QPrinter_override_virtual_DevType(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__DevType = slot;
}

int QPrinter_virtualbase_DevType(const void* self) {
	return ( (const MiqtVirtualQPrinter*)(self) )->virtualbase_DevType();
}

void QPrinter_override_virtual_NewPage(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__NewPage = slot;
}

bool QPrinter_virtualbase_NewPage(void* self) {
	return ( (MiqtVirtualQPrinter*)(self) )->virtualbase_NewPage();
}

void QPrinter_override_virtual_PaintEngine(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__PaintEngine = slot;
}

QPaintEngine* QPrinter_virtualbase_PaintEngine(const void* self) {
	return ( (const MiqtVirtualQPrinter*)(self) )->virtualbase_PaintEngine();
}

void QPrinter_override_virtual_Metric(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__Metric = slot;
}

int QPrinter_virtualbase_Metric(const void* self, PaintDeviceMetric param1) {
	return ( (const MiqtVirtualQPrinter*)(self) )->virtualbase_Metric(param1);
}

void QPrinter_override_virtual_SetPageLayout(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__SetPageLayout = slot;
}

bool QPrinter_virtualbase_SetPageLayout(void* self, QPageLayout* pageLayout) {
	return ( (MiqtVirtualQPrinter*)(self) )->virtualbase_SetPageLayout(pageLayout);
}

void QPrinter_override_virtual_SetPageSize(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__SetPageSize = slot;
}

bool QPrinter_virtualbase_SetPageSize(void* self, QPageSize* pageSize) {
	return ( (MiqtVirtualQPrinter*)(self) )->virtualbase_SetPageSize(pageSize);
}

void QPrinter_override_virtual_SetPageOrientation(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__SetPageOrientation = slot;
}

bool QPrinter_virtualbase_SetPageOrientation(void* self, int orientation) {
	return ( (MiqtVirtualQPrinter*)(self) )->virtualbase_SetPageOrientation(orientation);
}

void QPrinter_override_virtual_SetPageMargins(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__SetPageMargins = slot;
}

bool QPrinter_virtualbase_SetPageMargins(void* self, QMarginsF* margins, int units) {
	return ( (MiqtVirtualQPrinter*)(self) )->virtualbase_SetPageMargins(margins, units);
}

void QPrinter_override_virtual_SetPageRanges(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrinter*>( (QPrinter*)(self) )->handle__SetPageRanges = slot;
}

void QPrinter_virtualbase_SetPageRanges(void* self, QPageRanges* ranges) {
	( (MiqtVirtualQPrinter*)(self) )->virtualbase_SetPageRanges(ranges);
}

void QPrinter_Delete(QPrinter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPrinter*>( self );
	} else {
		delete self;
	}
}
