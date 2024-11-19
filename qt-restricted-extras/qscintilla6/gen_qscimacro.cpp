#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <qscimacro.h>
#include "gen_qscimacro.h"
#include "_cgo_export.h"

class MiqtVirtualQsciMacro : public virtual QsciMacro {
public:

	MiqtVirtualQsciMacro(QsciScintilla* parent): QsciMacro(parent) {};
	MiqtVirtualQsciMacro(const QString& asc, QsciScintilla* parent): QsciMacro(asc, parent) {};

	virtual ~MiqtVirtualQsciMacro() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Play = 0;

	// Subclass to allow providing a Go implementation
	virtual void play() override {
		if (handle__Play == 0) {
			QsciMacro::play();
			return;
		}
		

		miqt_exec_callback_QsciMacro_Play(this, handle__Play);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Play() {

		QsciMacro::play();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__StartRecording = 0;

	// Subclass to allow providing a Go implementation
	virtual void startRecording() override {
		if (handle__StartRecording == 0) {
			QsciMacro::startRecording();
			return;
		}
		

		miqt_exec_callback_QsciMacro_StartRecording(this, handle__StartRecording);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_StartRecording() {

		QsciMacro::startRecording();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EndRecording = 0;

	// Subclass to allow providing a Go implementation
	virtual void endRecording() override {
		if (handle__EndRecording == 0) {
			QsciMacro::endRecording();
			return;
		}
		

		miqt_exec_callback_QsciMacro_EndRecording(this, handle__EndRecording);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_EndRecording() {

		QsciMacro::endRecording();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__Event == 0) {
			return QsciMacro::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QsciMacro_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* event) {

		return QsciMacro::event(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__EventFilter == 0) {
			return QsciMacro::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QsciMacro_EventFilter(this, handle__EventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EventFilter(QObject* watched, QEvent* event) {

		return QsciMacro::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__TimerEvent == 0) {
			QsciMacro::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QsciMacro_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* event) {

		QsciMacro::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChildEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__ChildEvent == 0) {
			QsciMacro::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QsciMacro_ChildEvent(this, handle__ChildEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChildEvent(QChildEvent* event) {

		QsciMacro::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CustomEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__CustomEvent == 0) {
			QsciMacro::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QsciMacro_CustomEvent(this, handle__CustomEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CustomEvent(QEvent* event) {

		QsciMacro::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__ConnectNotify == 0) {
			QsciMacro::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QsciMacro_ConnectNotify(this, handle__ConnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ConnectNotify(QMetaMethod* signal) {

		QsciMacro::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DisconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__DisconnectNotify == 0) {
			QsciMacro::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QsciMacro_DisconnectNotify(this, handle__DisconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DisconnectNotify(QMetaMethod* signal) {

		QsciMacro::disconnectNotify(*signal);

	}

};

void QsciMacro_new(QsciScintilla* parent, QsciMacro** outptr_QsciMacro, QObject** outptr_QObject) {
	MiqtVirtualQsciMacro* ret = new MiqtVirtualQsciMacro(parent);
	*outptr_QsciMacro = ret;
	*outptr_QObject = static_cast<QObject*>(ret);
}

void QsciMacro_new2(struct miqt_string asc, QsciScintilla* parent, QsciMacro** outptr_QsciMacro, QObject** outptr_QObject) {
	QString asc_QString = QString::fromUtf8(asc.data, asc.len);
	MiqtVirtualQsciMacro* ret = new MiqtVirtualQsciMacro(asc_QString, parent);
	*outptr_QsciMacro = ret;
	*outptr_QObject = static_cast<QObject*>(ret);
}

QMetaObject* QsciMacro_MetaObject(const QsciMacro* self) {
	return (QMetaObject*) self->metaObject();
}

void* QsciMacro_Metacast(QsciMacro* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QsciMacro_Tr(const char* s) {
	QString _ret = QsciMacro::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QsciMacro_Clear(QsciMacro* self) {
	self->clear();
}

bool QsciMacro_Load(QsciMacro* self, struct miqt_string asc) {
	QString asc_QString = QString::fromUtf8(asc.data, asc.len);
	return self->load(asc_QString);
}

struct miqt_string QsciMacro_Save(const QsciMacro* self) {
	QString _ret = self->save();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QsciMacro_Play(QsciMacro* self) {
	self->play();
}

void QsciMacro_StartRecording(QsciMacro* self) {
	self->startRecording();
}

void QsciMacro_EndRecording(QsciMacro* self) {
	self->endRecording();
}

struct miqt_string QsciMacro_Tr2(const char* s, const char* c) {
	QString _ret = QsciMacro::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QsciMacro_Tr3(const char* s, const char* c, int n) {
	QString _ret = QsciMacro::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QsciMacro_override_virtual_Play(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__Play = slot;
}

void QsciMacro_virtualbase_Play(void* self) {
	( (MiqtVirtualQsciMacro*)(self) )->virtualbase_Play();
}

void QsciMacro_override_virtual_StartRecording(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__StartRecording = slot;
}

void QsciMacro_virtualbase_StartRecording(void* self) {
	( (MiqtVirtualQsciMacro*)(self) )->virtualbase_StartRecording();
}

void QsciMacro_override_virtual_EndRecording(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__EndRecording = slot;
}

void QsciMacro_virtualbase_EndRecording(void* self) {
	( (MiqtVirtualQsciMacro*)(self) )->virtualbase_EndRecording();
}

void QsciMacro_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__Event = slot;
}

bool QsciMacro_virtualbase_Event(void* self, QEvent* event) {
	return ( (MiqtVirtualQsciMacro*)(self) )->virtualbase_Event(event);
}

void QsciMacro_override_virtual_EventFilter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__EventFilter = slot;
}

bool QsciMacro_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQsciMacro*)(self) )->virtualbase_EventFilter(watched, event);
}

void QsciMacro_override_virtual_TimerEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__TimerEvent = slot;
}

void QsciMacro_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQsciMacro*)(self) )->virtualbase_TimerEvent(event);
}

void QsciMacro_override_virtual_ChildEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__ChildEvent = slot;
}

void QsciMacro_virtualbase_ChildEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQsciMacro*)(self) )->virtualbase_ChildEvent(event);
}

void QsciMacro_override_virtual_CustomEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__CustomEvent = slot;
}

void QsciMacro_virtualbase_CustomEvent(void* self, QEvent* event) {
	( (MiqtVirtualQsciMacro*)(self) )->virtualbase_CustomEvent(event);
}

void QsciMacro_override_virtual_ConnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__ConnectNotify = slot;
}

void QsciMacro_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQsciMacro*)(self) )->virtualbase_ConnectNotify(signal);
}

void QsciMacro_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciMacro*>( (QsciMacro*)(self) )->handle__DisconnectNotify = slot;
}

void QsciMacro_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQsciMacro*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QsciMacro_Delete(QsciMacro* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQsciMacro*>( self );
	} else {
		delete self;
	}
}

