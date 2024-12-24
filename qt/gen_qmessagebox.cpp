// +build ignore

#include <QAbstractButton>
#include <QCheckBox>
#include <QCloseEvent>
#include <QDialog>
#include <QEvent>
#include <QKeyEvent>
#include <QList>
#include <QMessageBox>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPixmap>
#include <QPushButton>
#include <QResizeEvent>
#include <QShowEvent>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qmessagebox.h>
#include "gen_qmessagebox.h"
class MiqtVirtualQMessageBox : public virtual QMessageBox {
public:
MiqtVirtualQMessageBox(QWidget* parent): QMessageBox(parent) {};
MiqtVirtualQMessageBox(): QMessageBox() {};
MiqtVirtualQMessageBox(Icon icon, const QString& title, const QString& text): QMessageBox(icon, title, text) {};
MiqtVirtualQMessageBox(const QString& title, const QString& text, Icon icon, int button0, int button1, int button2): QMessageBox(title, text, icon, button0, button1, button2) {};
MiqtVirtualQMessageBox(Icon icon, const QString& title, const QString& text, StandardButtons buttons): QMessageBox(icon, title, text, buttons) {};
MiqtVirtualQMessageBox(Icon icon, const QString& title, const QString& text, StandardButtons buttons, QWidget* parent): QMessageBox(icon, title, text, buttons, parent) {};
MiqtVirtualQMessageBox(Icon icon, const QString& title, const QString& text, StandardButtons buttons, QWidget* parent, Qt::WindowFlags flags): QMessageBox(icon, title, text, buttons, parent, flags) {};
MiqtVirtualQMessageBox(const QString& title, const QString& text, Icon icon, int button0, int button1, int button2, QWidget* parent): QMessageBox(title, text, icon, button0, button1, button2, parent) {};
MiqtVirtualQMessageBox(const QString& title, const QString& text, Icon icon, int button0, int button1, int button2, QWidget* parent, Qt::WindowFlags f): QMessageBox(title, text, icon, button0, button1, button2, parent, f) {};

virtual ~MiqtVirtualQMessageBox() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QMessageBox::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QMessageBox_MetaObject(const_cast<MiqtVirtualQMessageBox*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QMessageBox::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QMessageBox::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QMessageBox_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QMessageBox::qt_metacast(param1);

	}
};
QMessageBox* QMessageBox_new(QWidget* parent) {
return new MiqtVirtualQMessageBox(parent);
}
QMessageBox* QMessageBox_new2() {
return new MiqtVirtualQMessageBox();
}
QMessageBox* QMessageBox_new3(Icon icon, struct miqt_string title, struct miqt_string text) {
QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQMessageBox(icon, title_QString, text_QString);
}
QMessageBox* QMessageBox_new4(struct miqt_string title, struct miqt_string text, Icon icon, int button0, int button1, int button2) {
QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQMessageBox(title_QString, text_QString, icon, static_cast<int>(button0), static_cast<int>(button1), static_cast<int>(button2));
}
QMessageBox* QMessageBox_new5(Icon icon, struct miqt_string title, struct miqt_string text, StandardButtons buttons) {
QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQMessageBox(icon, title_QString, text_QString, buttons);
}
QMessageBox* QMessageBox_new6(Icon icon, struct miqt_string title, struct miqt_string text, StandardButtons buttons, QWidget* parent) {
QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQMessageBox(icon, title_QString, text_QString, buttons, parent);
}
QMessageBox* QMessageBox_new7(Icon icon, struct miqt_string title, struct miqt_string text, StandardButtons buttons, QWidget* parent, int flags) {
QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQMessageBox(icon, title_QString, text_QString, buttons, parent, static_cast<Qt::WindowFlags>(flags));
}
QMessageBox* QMessageBox_new8(struct miqt_string title, struct miqt_string text, Icon icon, int button0, int button1, int button2, QWidget* parent) {
QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQMessageBox(title_QString, text_QString, icon, static_cast<int>(button0), static_cast<int>(button1), static_cast<int>(button2), parent);
}
QMessageBox* QMessageBox_new9(struct miqt_string title, struct miqt_string text, Icon icon, int button0, int button1, int button2, QWidget* parent, int f) {
QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQMessageBox(title_QString, text_QString, icon, static_cast<int>(button0), static_cast<int>(button1), static_cast<int>(button2), parent, static_cast<Qt::WindowFlags>(f));
}
void QMessageBox_virtbase(QMessageBox* src
, QDialog** outptr_QDialog
) {
*outptr_QDialog = static_cast<QDialog*>(src);
}
QMetaObject* QMessageBox_MetaObject(const QMessageBox* self) {
	return (QMetaObject*) self->metaObject();
}
void* QMessageBox_Metacast(QMessageBox* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QMessageBox_Tr(const char* s) {
	QString _ret = QMessageBox::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMessageBox_AddButton(QMessageBox* self, QAbstractButton* button, ButtonRole role) {
	self->addButton(button, role);
}
QPushButton* QMessageBox_AddButton2(QMessageBox* self, struct miqt_string text, ButtonRole role) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addButton(text_QString, role);
}
QPushButton* QMessageBox_AddButtonWithButton(QMessageBox* self, StandardButton button) {
	return self->addButton(button);
}
void QMessageBox_RemoveButton(QMessageBox* self, QAbstractButton* button) {
	self->removeButton(button);
}
struct miqt_array /* of QAbstractButton* */  QMessageBox_Buttons(const QMessageBox* self) {
	QList<QAbstractButton *> _ret = self->buttons();
	// Convert QList<> from C++ memory to manually-managed C memory
	QAbstractButton** _arr = static_cast<QAbstractButton**>(malloc(sizeof(QAbstractButton*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
ButtonRole QMessageBox_ButtonRole(const QMessageBox* self, QAbstractButton* button) {
	return self->buttonRole(button);
}
void QMessageBox_SetStandardButtons(QMessageBox* self, StandardButtons buttons) {
	self->setStandardButtons(buttons);
}
StandardButtons QMessageBox_StandardButtons(const QMessageBox* self) {
	return self->standardButtons();
}
StandardButton QMessageBox_StandardButton(const QMessageBox* self, QAbstractButton* button) {
	return self->standardButton(button);
}
QAbstractButton* QMessageBox_Button(const QMessageBox* self, StandardButton which) {
	return self->button(which);
}
QPushButton* QMessageBox_DefaultButton(const QMessageBox* self) {
	return self->defaultButton();
}
void QMessageBox_SetDefaultButton(QMessageBox* self, QPushButton* button) {
	self->setDefaultButton(button);
}
void QMessageBox_SetDefaultButtonWithButton(QMessageBox* self, StandardButton button) {
	self->setDefaultButton(button);
}
QAbstractButton* QMessageBox_EscapeButton(const QMessageBox* self) {
	return self->escapeButton();
}
void QMessageBox_SetEscapeButton(QMessageBox* self, QAbstractButton* button) {
	self->setEscapeButton(button);
}
void QMessageBox_SetEscapeButtonWithButton(QMessageBox* self, StandardButton button) {
	self->setEscapeButton(button);
}
QAbstractButton* QMessageBox_ClickedButton(const QMessageBox* self) {
	return self->clickedButton();
}
struct miqt_string QMessageBox_Text(const QMessageBox* self) {
	QString _ret = self->text();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMessageBox_SetText(QMessageBox* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setText(text_QString);
}
Icon QMessageBox_Icon(const QMessageBox* self) {
	return self->icon();
}
void QMessageBox_SetIcon(QMessageBox* self, Icon icon) {
	self->setIcon(icon);
}
QPixmap* QMessageBox_IconPixmap(const QMessageBox* self) {
	return new QPixmap(self->iconPixmap());
}
void QMessageBox_SetIconPixmap(QMessageBox* self, QPixmap* pixmap) {
	self->setIconPixmap(*pixmap);
}
int QMessageBox_TextFormat(const QMessageBox* self) {
	Qt::TextFormat _ret = self->textFormat();
	return static_cast<int>(_ret);
}
void QMessageBox_SetTextFormat(QMessageBox* self, int format) {
	self->setTextFormat(static_cast<Qt::TextFormat>(format));
}
void QMessageBox_SetTextInteractionFlags(QMessageBox* self, int flags) {
	self->setTextInteractionFlags(static_cast<Qt::TextInteractionFlags>(flags));
}
int QMessageBox_TextInteractionFlags(const QMessageBox* self) {
	Qt::TextInteractionFlags _ret = self->textInteractionFlags();
	return static_cast<int>(_ret);
}
void QMessageBox_SetCheckBox(QMessageBox* self, QCheckBox* cb) {
	self->setCheckBox(cb);
}
QCheckBox* QMessageBox_CheckBox(const QMessageBox* self) {
	return self->checkBox();
}
void QMessageBox_SetOption(QMessageBox* self, Option option) {
	self->setOption(option);
}
bool QMessageBox_TestOption(const QMessageBox* self, Option option) {
	return self->testOption(option);
}
void QMessageBox_SetOptions(QMessageBox* self, Options options) {
	self->setOptions(options);
}
Options QMessageBox_Options(const QMessageBox* self) {
	return self->options();
}
StandardButton QMessageBox_Information(QWidget* parent, struct miqt_string title, struct miqt_string text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::information(parent, title_QString, text_QString);
}
StandardButton QMessageBox_Information2(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::information(parent, title_QString, text_QString, button0);
}
StandardButton QMessageBox_Question(QWidget* parent, struct miqt_string title, struct miqt_string text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::question(parent, title_QString, text_QString);
}
int QMessageBox_Question2(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0, StandardButton button1) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::question(parent, title_QString, text_QString, button0, button1);
}
StandardButton QMessageBox_Warning(QWidget* parent, struct miqt_string title, struct miqt_string text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::warning(parent, title_QString, text_QString);
}
int QMessageBox_Warning2(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0, StandardButton button1) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, button0, button1);
}
StandardButton QMessageBox_Critical(QWidget* parent, struct miqt_string title, struct miqt_string text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::critical(parent, title_QString, text_QString);
}
int QMessageBox_Critical2(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0, StandardButton button1) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, button0, button1);
}
void QMessageBox_About(QWidget* parent, struct miqt_string title, struct miqt_string text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QMessageBox::about(parent, title_QString, text_QString);
}
void QMessageBox_AboutQt(QWidget* parent) {
	QMessageBox::aboutQt(parent);
}
int QMessageBox_Information3(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::information(parent, title_QString, text_QString, static_cast<int>(button0));
}
int QMessageBox_Information4(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	return QMessageBox::information(parent, title_QString, text_QString, button0Text_QString);
}
int QMessageBox_Question3(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::question(parent, title_QString, text_QString, static_cast<int>(button0));
}
int QMessageBox_Question4(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	return QMessageBox::question(parent, title_QString, text_QString, button0Text_QString);
}
int QMessageBox_Warning3(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, static_cast<int>(button0), static_cast<int>(button1));
}
int QMessageBox_Warning4(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, button0Text_QString);
}
int QMessageBox_Critical3(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, static_cast<int>(button0), static_cast<int>(button1));
}
int QMessageBox_Critical4(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, button0Text_QString);
}
struct miqt_string QMessageBox_ButtonText(const QMessageBox* self, int button) {
	QString _ret = self->buttonText(static_cast<int>(button));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMessageBox_SetButtonText(QMessageBox* self, int button, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setButtonText(static_cast<int>(button), text_QString);
}
struct miqt_string QMessageBox_InformativeText(const QMessageBox* self) {
	QString _ret = self->informativeText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMessageBox_SetInformativeText(QMessageBox* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setInformativeText(text_QString);
}
struct miqt_string QMessageBox_DetailedText(const QMessageBox* self) {
	QString _ret = self->detailedText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMessageBox_SetDetailedText(QMessageBox* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setDetailedText(text_QString);
}
void QMessageBox_SetWindowTitle(QMessageBox* self, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	self->setWindowTitle(title_QString);
}
void QMessageBox_SetWindowModality(QMessageBox* self, int windowModality) {
	self->setWindowModality(static_cast<Qt::WindowModality>(windowModality));
}
QPixmap* QMessageBox_StandardIcon(Icon icon) {
	return new QPixmap(QMessageBox::standardIcon(icon));
}
void QMessageBox_ButtonClicked(QMessageBox* self, QAbstractButton* button) {
	self->buttonClicked(button);
}
void QMessageBox_connect_ButtonClicked(QMessageBox* self, intptr_t slot) {
	MiqtVirtualQMessageBox::connect(self, static_cast<void (QMessageBox::*)(QAbstractButton*)>(&QMessageBox::buttonClicked), self, [=](QAbstractButton* button) {
		QAbstractButton* sigval1 = button;
		miqt_exec_callback_QMessageBox_ButtonClicked(slot, sigval1);
	});
}
struct miqt_string QMessageBox_Tr2(const char* s, const char* c) {
	QString _ret = QMessageBox::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMessageBox_Tr3(const char* s, const char* c, int n) {
	QString _ret = QMessageBox::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMessageBox_SetOption2(QMessageBox* self, Option option, bool on) {
	self->setOption(option, on);
}
StandardButton QMessageBox_Information42(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::information(parent, title_QString, text_QString, buttons);
}
StandardButton QMessageBox_Information5(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons, StandardButton defaultButton) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::information(parent, title_QString, text_QString, buttons, defaultButton);
}
StandardButton QMessageBox_Information52(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0, StandardButton button1) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::information(parent, title_QString, text_QString, button0, button1);
}
StandardButton QMessageBox_Question42(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::question(parent, title_QString, text_QString, buttons);
}
StandardButton QMessageBox_Question5(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons, StandardButton defaultButton) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::question(parent, title_QString, text_QString, buttons, defaultButton);
}
StandardButton QMessageBox_Warning42(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, buttons);
}
StandardButton QMessageBox_Warning5(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons, StandardButton defaultButton) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, buttons, defaultButton);
}
StandardButton QMessageBox_Critical42(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, buttons);
}
StandardButton QMessageBox_Critical5(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons, StandardButton defaultButton) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, buttons, defaultButton);
}
void QMessageBox_AboutQt2(QWidget* parent, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QMessageBox::aboutQt(parent, title_QString);
}
int QMessageBox_Information53(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::information(parent, title_QString, text_QString, static_cast<int>(button0), static_cast<int>(button1));
}
int QMessageBox_Information6(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1, int button2) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::information(parent, title_QString, text_QString, static_cast<int>(button0), static_cast<int>(button1), static_cast<int>(button2));
}
int QMessageBox_Information54(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	return QMessageBox::information(parent, title_QString, text_QString, button0Text_QString, button1Text_QString);
}
int QMessageBox_Information62(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::information(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString);
}
int QMessageBox_Information7(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::information(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString, static_cast<int>(defaultButtonNumber));
}
int QMessageBox_Information8(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber, int escapeButtonNumber) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::information(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString, static_cast<int>(defaultButtonNumber), static_cast<int>(escapeButtonNumber));
}
int QMessageBox_Question52(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::question(parent, title_QString, text_QString, static_cast<int>(button0), static_cast<int>(button1));
}
int QMessageBox_Question6(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1, int button2) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::question(parent, title_QString, text_QString, static_cast<int>(button0), static_cast<int>(button1), static_cast<int>(button2));
}
int QMessageBox_Question53(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	return QMessageBox::question(parent, title_QString, text_QString, button0Text_QString, button1Text_QString);
}
int QMessageBox_Question62(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::question(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString);
}
int QMessageBox_Question7(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::question(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString, static_cast<int>(defaultButtonNumber));
}
int QMessageBox_Question8(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber, int escapeButtonNumber) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::question(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString, static_cast<int>(defaultButtonNumber), static_cast<int>(escapeButtonNumber));
}
int QMessageBox_Warning6(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1, int button2) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, static_cast<int>(button0), static_cast<int>(button1), static_cast<int>(button2));
}
int QMessageBox_Warning52(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, button0Text_QString, button1Text_QString);
}
int QMessageBox_Warning62(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString);
}
int QMessageBox_Warning7(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString, static_cast<int>(defaultButtonNumber));
}
int QMessageBox_Warning8(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber, int escapeButtonNumber) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::warning(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString, static_cast<int>(defaultButtonNumber), static_cast<int>(escapeButtonNumber));
}
int QMessageBox_Critical6(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1, int button2) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, static_cast<int>(button0), static_cast<int>(button1), static_cast<int>(button2));
}
int QMessageBox_Critical52(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, button0Text_QString, button1Text_QString);
}
int QMessageBox_Critical62(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString);
}
int QMessageBox_Critical7(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString, static_cast<int>(defaultButtonNumber));
}
int QMessageBox_Critical8(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber, int escapeButtonNumber) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	QString text_QString = QString::fromUtf8(text.data, text.len);
	QString button0Text_QString = QString::fromUtf8(button0Text.data, button0Text.len);
	QString button1Text_QString = QString::fromUtf8(button1Text.data, button1Text.len);
	QString button2Text_QString = QString::fromUtf8(button2Text.data, button2Text.len);
	return QMessageBox::critical(parent, title_QString, text_QString, button0Text_QString, button1Text_QString, button2Text_QString, static_cast<int>(defaultButtonNumber), static_cast<int>(escapeButtonNumber));
}
void QMessageBox_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMessageBox*>( (QMessageBox*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QMessageBox_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQMessageBox*)(self) )->virtualbase_MetaObject();
}
void QMessageBox_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMessageBox*>( (QMessageBox*)(self) )->handle__Metacast = slot;
}
void* QMessageBox_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQMessageBox*)(self) )->virtualbase_Metacast(param1);
}
void QMessageBox_Delete(QMessageBox* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQMessageBox*>( self );
	} else {
		delete self;
	}
}
