#include <QAudioDevice>
#include <QAudioInput>
#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <qaudioinput.h>
#include "gen_qaudioinput.h"

#ifdef __cplusplus
extern "C" {
#endif

void miqt_exec_callback_QAudioInput_deviceChanged(intptr_t);
void miqt_exec_callback_QAudioInput_volumeChanged(intptr_t, float);
void miqt_exec_callback_QAudioInput_mutedChanged(intptr_t, bool);
bool miqt_exec_callback_QAudioInput_event(QAudioInput*, intptr_t, QEvent*);
bool miqt_exec_callback_QAudioInput_eventFilter(QAudioInput*, intptr_t, QObject*, QEvent*);
void miqt_exec_callback_QAudioInput_timerEvent(QAudioInput*, intptr_t, QTimerEvent*);
void miqt_exec_callback_QAudioInput_childEvent(QAudioInput*, intptr_t, QChildEvent*);
void miqt_exec_callback_QAudioInput_customEvent(QAudioInput*, intptr_t, QEvent*);
void miqt_exec_callback_QAudioInput_connectNotify(QAudioInput*, intptr_t, QMetaMethod*);
void miqt_exec_callback_QAudioInput_disconnectNotify(QAudioInput*, intptr_t, QMetaMethod*);
#ifdef __cplusplus
} /* extern C */
#endif

class MiqtVirtualQAudioInput final : public QAudioInput {
public:

	MiqtVirtualQAudioInput(): QAudioInput() {};
	MiqtVirtualQAudioInput(const QAudioDevice& deviceInfo): QAudioInput(deviceInfo) {};
	MiqtVirtualQAudioInput(QObject* parent): QAudioInput(parent) {};
	MiqtVirtualQAudioInput(const QAudioDevice& deviceInfo, QObject* parent): QAudioInput(deviceInfo, parent) {};

	virtual ~MiqtVirtualQAudioInput() override = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__event == 0) {
			return QAudioInput::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QAudioInput_event(this, handle__event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_event(QEvent* event) {

		return QAudioInput::event(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__eventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__eventFilter == 0) {
			return QAudioInput::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QAudioInput_eventFilter(this, handle__eventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_eventFilter(QObject* watched, QEvent* event) {

		return QAudioInput::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__timerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__timerEvent == 0) {
			QAudioInput::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QAudioInput_timerEvent(this, handle__timerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_timerEvent(QTimerEvent* event) {

		QAudioInput::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__childEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__childEvent == 0) {
			QAudioInput::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QAudioInput_childEvent(this, handle__childEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_childEvent(QChildEvent* event) {

		QAudioInput::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__customEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__customEvent == 0) {
			QAudioInput::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QAudioInput_customEvent(this, handle__customEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_customEvent(QEvent* event) {

		QAudioInput::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__connectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__connectNotify == 0) {
			QAudioInput::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QAudioInput_connectNotify(this, handle__connectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_connectNotify(QMetaMethod* signal) {

		QAudioInput::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__disconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__disconnectNotify == 0) {
			QAudioInput::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QAudioInput_disconnectNotify(this, handle__disconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_disconnectNotify(QMetaMethod* signal) {

		QAudioInput::disconnectNotify(*signal);

	}

	// Wrappers to allow calling protected methods:
	friend QObject* QAudioInput_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
	friend int QAudioInput_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
	friend int QAudioInput_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
	friend bool QAudioInput_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);
};

QAudioInput* QAudioInput_new() {
	return new MiqtVirtualQAudioInput();
}

QAudioInput* QAudioInput_new2(QAudioDevice* deviceInfo) {
	return new MiqtVirtualQAudioInput(*deviceInfo);
}

QAudioInput* QAudioInput_new3(QObject* parent) {
	return new MiqtVirtualQAudioInput(parent);
}

QAudioInput* QAudioInput_new4(QAudioDevice* deviceInfo, QObject* parent) {
	return new MiqtVirtualQAudioInput(*deviceInfo, parent);
}

void QAudioInput_virtbase(QAudioInput* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QAudioInput_metaObject(const QAudioInput* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAudioInput_metacast(QAudioInput* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAudioInput_tr(const char* s) {
	QString _ret = QAudioInput::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QAudioDevice* QAudioInput_device(const QAudioInput* self) {
	return new QAudioDevice(self->device());
}

float QAudioInput_volume(const QAudioInput* self) {
	return self->volume();
}

bool QAudioInput_isMuted(const QAudioInput* self) {
	return self->isMuted();
}

void QAudioInput_setDevice(QAudioInput* self, QAudioDevice* device) {
	self->setDevice(*device);
}

void QAudioInput_setVolume(QAudioInput* self, float volume) {
	self->setVolume(static_cast<float>(volume));
}

void QAudioInput_setMuted(QAudioInput* self, bool muted) {
	self->setMuted(muted);
}

void QAudioInput_deviceChanged(QAudioInput* self) {
	self->deviceChanged();
}

void QAudioInput_connect_deviceChanged(QAudioInput* self, intptr_t slot) {
	MiqtVirtualQAudioInput::connect(self, static_cast<void (QAudioInput::*)()>(&QAudioInput::deviceChanged), self, [=]() {
		miqt_exec_callback_QAudioInput_deviceChanged(slot);
	});
}

void QAudioInput_volumeChanged(QAudioInput* self, float volume) {
	self->volumeChanged(static_cast<float>(volume));
}

void QAudioInput_connect_volumeChanged(QAudioInput* self, intptr_t slot) {
	MiqtVirtualQAudioInput::connect(self, static_cast<void (QAudioInput::*)(float)>(&QAudioInput::volumeChanged), self, [=](float volume) {
		float sigval1 = volume;
		miqt_exec_callback_QAudioInput_volumeChanged(slot, sigval1);
	});
}

void QAudioInput_mutedChanged(QAudioInput* self, bool muted) {
	self->mutedChanged(muted);
}

void QAudioInput_connect_mutedChanged(QAudioInput* self, intptr_t slot) {
	MiqtVirtualQAudioInput::connect(self, static_cast<void (QAudioInput::*)(bool)>(&QAudioInput::mutedChanged), self, [=](bool muted) {
		bool sigval1 = muted;
		miqt_exec_callback_QAudioInput_mutedChanged(slot, sigval1);
	});
}

struct miqt_string QAudioInput_tr2(const char* s, const char* c) {
	QString _ret = QAudioInput::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioInput_tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioInput::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QAudioInput_override_virtual_event(void* self, intptr_t slot) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__event = slot;
	return true;
}

bool QAudioInput_virtualbase_event(void* self, QEvent* event) {
	return ( (MiqtVirtualQAudioInput*)(self) )->virtualbase_event(event);
}

bool QAudioInput_override_virtual_eventFilter(void* self, intptr_t slot) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__eventFilter = slot;
	return true;
}

bool QAudioInput_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQAudioInput*)(self) )->virtualbase_eventFilter(watched, event);
}

bool QAudioInput_override_virtual_timerEvent(void* self, intptr_t slot) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__timerEvent = slot;
	return true;
}

void QAudioInput_virtualbase_timerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQAudioInput*)(self) )->virtualbase_timerEvent(event);
}

bool QAudioInput_override_virtual_childEvent(void* self, intptr_t slot) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__childEvent = slot;
	return true;
}

void QAudioInput_virtualbase_childEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQAudioInput*)(self) )->virtualbase_childEvent(event);
}

bool QAudioInput_override_virtual_customEvent(void* self, intptr_t slot) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__customEvent = slot;
	return true;
}

void QAudioInput_virtualbase_customEvent(void* self, QEvent* event) {
	( (MiqtVirtualQAudioInput*)(self) )->virtualbase_customEvent(event);
}

bool QAudioInput_override_virtual_connectNotify(void* self, intptr_t slot) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__connectNotify = slot;
	return true;
}

void QAudioInput_virtualbase_connectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQAudioInput*)(self) )->virtualbase_connectNotify(signal);
}

bool QAudioInput_override_virtual_disconnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__disconnectNotify = slot;
	return true;
}

void QAudioInput_virtualbase_disconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQAudioInput*)(self) )->virtualbase_disconnectNotify(signal);
}

QObject* QAudioInput_protectedbase_sender(bool* _dynamic_cast_ok, const void* self) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return nullptr;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->sender();

}

int QAudioInput_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return 0;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->senderSignalIndex();

}

int QAudioInput_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return 0;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->receivers(signal);

}

bool QAudioInput_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal) {
	MiqtVirtualQAudioInput* self_cast = dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return false;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->isSignalConnected(*signal);

}

void QAudioInput_delete(QAudioInput* self) {
	delete self;
}

