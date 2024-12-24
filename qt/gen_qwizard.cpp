// +build ignore

#include <QAbstractButton>
#include <QByteArray>
#include <QDialog>
#include <QEvent>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPixmap>
#include <QResizeEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <QWidget>
#include <QWizard>
#include <QWizardPage>
#include <qwizard.h>
#include "gen_qwizard.h"
class MiqtVirtualQWizard : public virtual QWizard {
public:
MiqtVirtualQWizard(QWidget* parent): QWizard(parent) {};
MiqtVirtualQWizard(): QWizard() {};
MiqtVirtualQWizard(QWidget* parent, Qt::WindowFlags flags): QWizard(parent, flags) {};

virtual ~MiqtVirtualQWizard() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QWizard::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QWizard_MetaObject(const_cast<MiqtVirtualQWizard*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QWizard::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QWizard::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QWizard_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QWizard::qt_metacast(param1);

	}
};
QWizard* QWizard_new(QWidget* parent) {
return new MiqtVirtualQWizard(parent);
}
QWizard* QWizard_new2() {
return new MiqtVirtualQWizard();
}
QWizard* QWizard_new3(QWidget* parent, int flags) {
return new MiqtVirtualQWizard(parent, static_cast<Qt::WindowFlags>(flags));
}
void QWizard_virtbase(QWizard* src
, QDialog** outptr_QDialog
) {
*outptr_QDialog = static_cast<QDialog*>(src);
}
QMetaObject* QWizard_MetaObject(const QWizard* self) {
	return (QMetaObject*) self->metaObject();
}
void* QWizard_Metacast(QWizard* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QWizard_Tr(const char* s) {
	QString _ret = QWizard::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QWizard_AddPage(QWizard* self, QWizardPage* page) {
	return self->addPage(page);
}
void QWizard_SetPage(QWizard* self, int id, QWizardPage* page) {
	self->setPage(static_cast<int>(id), page);
}
void QWizard_RemovePage(QWizard* self, int id) {
	self->removePage(static_cast<int>(id));
}
QWizardPage* QWizard_Page(const QWizard* self, int id) {
	return self->page(static_cast<int>(id));
}
bool QWizard_HasVisitedPage(const QWizard* self, int id) {
	return self->hasVisitedPage(static_cast<int>(id));
}
struct miqt_array /* of int */  QWizard_VisitedIds(const QWizard* self) {
	QList<int> _ret = self->visitedIds();
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
struct miqt_array /* of int */  QWizard_PageIds(const QWizard* self) {
	QList<int> _ret = self->pageIds();
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
void QWizard_SetStartId(QWizard* self, int id) {
	self->setStartId(static_cast<int>(id));
}
int QWizard_StartId(const QWizard* self) {
	return self->startId();
}
QWizardPage* QWizard_CurrentPage(const QWizard* self) {
	return self->currentPage();
}
int QWizard_CurrentId(const QWizard* self) {
	return self->currentId();
}
bool QWizard_ValidateCurrentPage(QWizard* self) {
	return self->validateCurrentPage();
}
int QWizard_NextId(const QWizard* self) {
	return self->nextId();
}
void QWizard_SetField(QWizard* self, struct miqt_string name, QVariant* value) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setField(name_QString, *value);
}
QVariant* QWizard_Field(const QWizard* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new QVariant(self->field(name_QString));
}
void QWizard_SetWizardStyle(QWizard* self, WizardStyle style) {
	self->setWizardStyle(style);
}
WizardStyle QWizard_WizardStyle(const QWizard* self) {
	return self->wizardStyle();
}
void QWizard_SetOption(QWizard* self, WizardOption option) {
	self->setOption(option);
}
bool QWizard_TestOption(const QWizard* self, WizardOption option) {
	return self->testOption(option);
}
void QWizard_SetOptions(QWizard* self, WizardOptions options) {
	self->setOptions(options);
}
WizardOptions QWizard_Options(const QWizard* self) {
	return self->options();
}
void QWizard_SetButtonText(QWizard* self, WizardButton which, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setButtonText(which, text_QString);
}
struct miqt_string QWizard_ButtonText(const QWizard* self, WizardButton which) {
	QString _ret = self->buttonText(which);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWizard_SetButtonLayout(QWizard* self, struct miqt_array /* of WizardButton */  layout) {
	QList<WizardButton> layout_QList;
	layout_QList.reserve(layout.len);
	WizardButton* layout_arr = static_cast<WizardButton*>(layout.data);
	for(size_t i = 0; i < layout.len; ++i) {
		layout_QList.push_back(layout_arr[i]);
	}
	self->setButtonLayout(layout_QList);
}
void QWizard_SetButton(QWizard* self, WizardButton which, QAbstractButton* button) {
	self->setButton(which, button);
}
QAbstractButton* QWizard_Button(const QWizard* self, WizardButton which) {
	return self->button(which);
}
void QWizard_SetTitleFormat(QWizard* self, int format) {
	self->setTitleFormat(static_cast<Qt::TextFormat>(format));
}
int QWizard_TitleFormat(const QWizard* self) {
	Qt::TextFormat _ret = self->titleFormat();
	return static_cast<int>(_ret);
}
void QWizard_SetSubTitleFormat(QWizard* self, int format) {
	self->setSubTitleFormat(static_cast<Qt::TextFormat>(format));
}
int QWizard_SubTitleFormat(const QWizard* self) {
	Qt::TextFormat _ret = self->subTitleFormat();
	return static_cast<int>(_ret);
}
void QWizard_SetPixmap(QWizard* self, WizardPixmap which, QPixmap* pixmap) {
	self->setPixmap(which, *pixmap);
}
QPixmap* QWizard_Pixmap(const QWizard* self, WizardPixmap which) {
	return new QPixmap(self->pixmap(which));
}
void QWizard_SetSideWidget(QWizard* self, QWidget* widget) {
	self->setSideWidget(widget);
}
QWidget* QWizard_SideWidget(const QWizard* self) {
	return self->sideWidget();
}
void QWizard_SetDefaultProperty(QWizard* self, const char* className, const char* property, const char* changedSignal) {
	self->setDefaultProperty(className, property, changedSignal);
}
void QWizard_SetVisible(QWizard* self, bool visible) {
	self->setVisible(visible);
}
QSize* QWizard_SizeHint(const QWizard* self) {
	return new QSize(self->sizeHint());
}
void QWizard_CurrentIdChanged(QWizard* self, int id) {
	self->currentIdChanged(static_cast<int>(id));
}
void QWizard_connect_CurrentIdChanged(QWizard* self, intptr_t slot) {
	MiqtVirtualQWizard::connect(self, static_cast<void (QWizard::*)(int)>(&QWizard::currentIdChanged), self, [=](int id) {
		int sigval1 = id;
		miqt_exec_callback_QWizard_CurrentIdChanged(slot, sigval1);
	});
}
void QWizard_HelpRequested(QWizard* self) {
	self->helpRequested();
}
void QWizard_connect_HelpRequested(QWizard* self, intptr_t slot) {
	MiqtVirtualQWizard::connect(self, static_cast<void (QWizard::*)()>(&QWizard::helpRequested), self, [=]() {
		miqt_exec_callback_QWizard_HelpRequested(slot);
	});
}
void QWizard_CustomButtonClicked(QWizard* self, int which) {
	self->customButtonClicked(static_cast<int>(which));
}
void QWizard_connect_CustomButtonClicked(QWizard* self, intptr_t slot) {
	MiqtVirtualQWizard::connect(self, static_cast<void (QWizard::*)(int)>(&QWizard::customButtonClicked), self, [=](int which) {
		int sigval1 = which;
		miqt_exec_callback_QWizard_CustomButtonClicked(slot, sigval1);
	});
}
void QWizard_PageAdded(QWizard* self, int id) {
	self->pageAdded(static_cast<int>(id));
}
void QWizard_connect_PageAdded(QWizard* self, intptr_t slot) {
	MiqtVirtualQWizard::connect(self, static_cast<void (QWizard::*)(int)>(&QWizard::pageAdded), self, [=](int id) {
		int sigval1 = id;
		miqt_exec_callback_QWizard_PageAdded(slot, sigval1);
	});
}
void QWizard_PageRemoved(QWizard* self, int id) {
	self->pageRemoved(static_cast<int>(id));
}
void QWizard_connect_PageRemoved(QWizard* self, intptr_t slot) {
	MiqtVirtualQWizard::connect(self, static_cast<void (QWizard::*)(int)>(&QWizard::pageRemoved), self, [=](int id) {
		int sigval1 = id;
		miqt_exec_callback_QWizard_PageRemoved(slot, sigval1);
	});
}
void QWizard_Back(QWizard* self) {
	self->back();
}
void QWizard_Next(QWizard* self) {
	self->next();
}
void QWizard_SetCurrentId(QWizard* self, int id) {
	self->setCurrentId(static_cast<int>(id));
}
void QWizard_Restart(QWizard* self) {
	self->restart();
}
struct miqt_string QWizard_Tr2(const char* s, const char* c) {
	QString _ret = QWizard::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QWizard_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWizard::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWizard_SetOption2(QWizard* self, WizardOption option, bool on) {
	self->setOption(option, on);
}
void QWizard_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWizard*>( (QWizard*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QWizard_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQWizard*)(self) )->virtualbase_MetaObject();
}
void QWizard_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWizard*>( (QWizard*)(self) )->handle__Metacast = slot;
}
void* QWizard_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQWizard*)(self) )->virtualbase_Metacast(param1);
}
void QWizard_Delete(QWizard* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWizard*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQWizardPage : public virtual QWizardPage {
public:
MiqtVirtualQWizardPage(QWidget* parent): QWizardPage(parent) {};
MiqtVirtualQWizardPage(): QWizardPage() {};

virtual ~MiqtVirtualQWizardPage() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QWizardPage::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QWizardPage_MetaObject(const_cast<MiqtVirtualQWizardPage*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QWizardPage::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QWizardPage::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QWizardPage_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QWizardPage::qt_metacast(param1);

	}
};
QWizardPage* QWizardPage_new(QWidget* parent) {
return new MiqtVirtualQWizardPage(parent);
}
QWizardPage* QWizardPage_new2() {
return new MiqtVirtualQWizardPage();
}
void QWizardPage_virtbase(QWizardPage* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QWizardPage_MetaObject(const QWizardPage* self) {
	return (QMetaObject*) self->metaObject();
}
void* QWizardPage_Metacast(QWizardPage* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QWizardPage_Tr(const char* s) {
	QString _ret = QWizardPage::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWizardPage_SetTitle(QWizardPage* self, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	self->setTitle(title_QString);
}
struct miqt_string QWizardPage_Title(const QWizardPage* self) {
	QString _ret = self->title();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWizardPage_SetSubTitle(QWizardPage* self, struct miqt_string subTitle) {
	QString subTitle_QString = QString::fromUtf8(subTitle.data, subTitle.len);
	self->setSubTitle(subTitle_QString);
}
struct miqt_string QWizardPage_SubTitle(const QWizardPage* self) {
	QString _ret = self->subTitle();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWizardPage_SetPixmap(QWizardPage* self, int which, QPixmap* pixmap) {
	self->setPixmap(static_cast<QWizard::WizardPixmap>(which), *pixmap);
}
QPixmap* QWizardPage_Pixmap(const QWizardPage* self, int which) {
	return new QPixmap(self->pixmap(static_cast<QWizard::WizardPixmap>(which)));
}
void QWizardPage_SetFinalPage(QWizardPage* self, bool finalPage) {
	self->setFinalPage(finalPage);
}
bool QWizardPage_IsFinalPage(const QWizardPage* self) {
	return self->isFinalPage();
}
void QWizardPage_SetCommitPage(QWizardPage* self, bool commitPage) {
	self->setCommitPage(commitPage);
}
bool QWizardPage_IsCommitPage(const QWizardPage* self) {
	return self->isCommitPage();
}
void QWizardPage_SetButtonText(QWizardPage* self, int which, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setButtonText(static_cast<QWizard::WizardButton>(which), text_QString);
}
struct miqt_string QWizardPage_ButtonText(const QWizardPage* self, int which) {
	QString _ret = self->buttonText(static_cast<QWizard::WizardButton>(which));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWizardPage_InitializePage(QWizardPage* self) {
	self->initializePage();
}
void QWizardPage_CleanupPage(QWizardPage* self) {
	self->cleanupPage();
}
bool QWizardPage_ValidatePage(QWizardPage* self) {
	return self->validatePage();
}
bool QWizardPage_IsComplete(const QWizardPage* self) {
	return self->isComplete();
}
int QWizardPage_NextId(const QWizardPage* self) {
	return self->nextId();
}
void QWizardPage_CompleteChanged(QWizardPage* self) {
	self->completeChanged();
}
void QWizardPage_connect_CompleteChanged(QWizardPage* self, intptr_t slot) {
	MiqtVirtualQWizardPage::connect(self, static_cast<void (QWizardPage::*)()>(&QWizardPage::completeChanged), self, [=]() {
		miqt_exec_callback_QWizardPage_CompleteChanged(slot);
	});
}
struct miqt_string QWizardPage_Tr2(const char* s, const char* c) {
	QString _ret = QWizardPage::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QWizardPage_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWizardPage::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWizardPage_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWizardPage*>( (QWizardPage*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QWizardPage_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQWizardPage*)(self) )->virtualbase_MetaObject();
}
void QWizardPage_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWizardPage*>( (QWizardPage*)(self) )->handle__Metacast = slot;
}
void* QWizardPage_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQWizardPage*)(self) )->virtualbase_Metacast(param1);
}
void QWizardPage_Delete(QWizardPage* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWizardPage*>( self );
	} else {
		delete self;
	}
}
