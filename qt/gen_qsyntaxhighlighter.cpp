// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QSyntaxHighlighter>
#include <QTextBlock>
#include <QTextDocument>
#include <qsyntaxhighlighter.h>
#include "gen_qsyntaxhighlighter.h"
class MiqtVirtualQSyntaxHighlighter : public virtual QSyntaxHighlighter {
public:
MiqtVirtualQSyntaxHighlighter(QObject* parent): QSyntaxHighlighter(parent) {};
MiqtVirtualQSyntaxHighlighter(QTextDocument* parent): QSyntaxHighlighter(parent) {};

virtual ~MiqtVirtualQSyntaxHighlighter() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSyntaxHighlighter::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSyntaxHighlighter_MetaObject(const_cast<MiqtVirtualQSyntaxHighlighter*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSyntaxHighlighter::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSyntaxHighlighter::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSyntaxHighlighter_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSyntaxHighlighter::qt_metacast(param1);

	}
};
QSyntaxHighlighter* QSyntaxHighlighter_new(QObject* parent) {
return new MiqtVirtualQSyntaxHighlighter(parent);
}
QSyntaxHighlighter* QSyntaxHighlighter_new2(QTextDocument* parent) {
return new MiqtVirtualQSyntaxHighlighter(parent);
}
void QSyntaxHighlighter_virtbase(QSyntaxHighlighter* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QSyntaxHighlighter_MetaObject(const QSyntaxHighlighter* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSyntaxHighlighter_Metacast(QSyntaxHighlighter* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSyntaxHighlighter_Tr(const char* s) {
	QString _ret = QSyntaxHighlighter::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSyntaxHighlighter_SetDocument(QSyntaxHighlighter* self, QTextDocument* doc) {
	self->setDocument(doc);
}
QTextDocument* QSyntaxHighlighter_Document(const QSyntaxHighlighter* self) {
	return self->document();
}
void QSyntaxHighlighter_Rehighlight(QSyntaxHighlighter* self) {
	self->rehighlight();
}
void QSyntaxHighlighter_RehighlightBlock(QSyntaxHighlighter* self, QTextBlock* block) {
	self->rehighlightBlock(*block);
}
struct miqt_string QSyntaxHighlighter_Tr2(const char* s, const char* c) {
	QString _ret = QSyntaxHighlighter::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSyntaxHighlighter_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSyntaxHighlighter::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSyntaxHighlighter_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSyntaxHighlighter*>( (QSyntaxHighlighter*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSyntaxHighlighter_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSyntaxHighlighter*)(self) )->virtualbase_MetaObject();
}
void QSyntaxHighlighter_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSyntaxHighlighter*>( (QSyntaxHighlighter*)(self) )->handle__Metacast = slot;
}
void* QSyntaxHighlighter_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSyntaxHighlighter*)(self) )->virtualbase_Metacast(param1);
}
void QSyntaxHighlighter_Delete(QSyntaxHighlighter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSyntaxHighlighter*>( self );
	} else {
		delete self;
	}
}
