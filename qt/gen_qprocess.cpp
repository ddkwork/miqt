// +build ignore

#include <QByteArray>
#include <QIODevice>
#include <QIODeviceBase>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QProcess>
#define WORKAROUND_INNER_CLASS_DEFINITION_QProcess__CreateProcessArguments
#include <QProcessEnvironment>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qprocess.h>
#include "gen_qprocess.h"
QProcessEnvironment* QProcessEnvironment_new() {
return new QProcessEnvironment();
}
QProcessEnvironment* QProcessEnvironment_new2(Initialization param1) {
return new QProcessEnvironment(param1);
}
QProcessEnvironment* QProcessEnvironment_new3(QProcessEnvironment* other) {
return new QProcessEnvironment(*other);
}
void QProcessEnvironment_OperatorAssign(QProcessEnvironment* self, QProcessEnvironment* other) {
	self->operator=(*other);
}
void QProcessEnvironment_Swap(QProcessEnvironment* self, QProcessEnvironment* other) {
	self->swap(*other);
}
bool QProcessEnvironment_IsEmpty(const QProcessEnvironment* self) {
	return self->isEmpty();
}
bool QProcessEnvironment_InheritsFromParent(const QProcessEnvironment* self) {
	return self->inheritsFromParent();
}
void QProcessEnvironment_Clear(QProcessEnvironment* self) {
	self->clear();
}
bool QProcessEnvironment_Contains(const QProcessEnvironment* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return self->contains(name_QString);
}
void QProcessEnvironment_Insert(QProcessEnvironment* self, struct miqt_string name, struct miqt_string value) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString value_QString = QString::fromUtf8(value.data, value.len);
	self->insert(name_QString, value_QString);
}
void QProcessEnvironment_Remove(QProcessEnvironment* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->remove(name_QString);
}
struct miqt_string QProcessEnvironment_Value(const QProcessEnvironment* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString _ret = self->value(name_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_array /* of struct miqt_string */  QProcessEnvironment_ToStringList(const QProcessEnvironment* self) {
	QStringList _ret = self->toStringList();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_array /* of struct miqt_string */  QProcessEnvironment_Keys(const QProcessEnvironment* self) {
	QStringList _ret = self->keys();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QProcessEnvironment_InsertWithQProcessEnvironment(QProcessEnvironment* self, QProcessEnvironment* e) {
	self->insert(*e);
}
QProcessEnvironment* QProcessEnvironment_SystemEnvironment() {
	return new QProcessEnvironment(QProcessEnvironment::systemEnvironment());
}
struct miqt_string QProcessEnvironment_Value2(const QProcessEnvironment* self, struct miqt_string name, struct miqt_string defaultValue) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString defaultValue_QString = QString::fromUtf8(defaultValue.data, defaultValue.len);
	QString _ret = self->value(name_QString, defaultValue_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QProcessEnvironment_Delete(QProcessEnvironment* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QProcessEnvironment*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQProcess : public virtual QProcess {
public:
MiqtVirtualQProcess(): QProcess() {};
MiqtVirtualQProcess(QObject* parent): QProcess(parent) {};

virtual ~MiqtVirtualQProcess() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QProcess::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QProcess_MetaObject(const_cast<MiqtVirtualQProcess*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QProcess::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QProcess::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QProcess_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QProcess::qt_metacast(param1);

	}
};
QProcess* QProcess_new() {
return new MiqtVirtualQProcess();
}
QProcess* QProcess_new2(QObject* parent) {
return new MiqtVirtualQProcess(parent);
}
void QProcess_virtbase(QProcess* src
, QIODevice** outptr_QIODevice
) {
*outptr_QIODevice = static_cast<QIODevice*>(src);
}
QMetaObject* QProcess_MetaObject(const QProcess* self) {
	return (QMetaObject*) self->metaObject();
}
void* QProcess_Metacast(QProcess* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QProcess_Tr(const char* s) {
	QString _ret = QProcess::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QProcess_Start(QProcess* self, struct miqt_string program) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	self->start(program_QString);
}
void QProcess_Start2(QProcess* self) {
	self->start();
}
void QProcess_StartCommand(QProcess* self, struct miqt_string command) {
	QString command_QString = QString::fromUtf8(command.data, command.len);
	self->startCommand(command_QString);
}
bool QProcess_StartDetached(QProcess* self) {
	return self->startDetached();
}
bool QProcess_Open(QProcess* self, OpenMode mode) {
	return self->open(mode);
}
struct miqt_string QProcess_Program(const QProcess* self) {
	QString _ret = self->program();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QProcess_SetProgram(QProcess* self, struct miqt_string program) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	self->setProgram(program_QString);
}
struct miqt_array /* of struct miqt_string */  QProcess_Arguments(const QProcess* self) {
	QStringList _ret = self->arguments();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QProcess_SetArguments(QProcess* self, struct miqt_array /* of struct miqt_string */  arguments) {
	QStringList arguments_QList;
	arguments_QList.reserve(arguments.len);
	struct miqt_string* arguments_arr = static_cast<struct miqt_string*>(arguments.data);
	for(size_t i = 0; i < arguments.len; ++i) {
		QString arguments_arr_i_QString = QString::fromUtf8(arguments_arr[i].data, arguments_arr[i].len);
		arguments_QList.push_back(arguments_arr_i_QString);
	}
	self->setArguments(arguments_QList);
}
ProcessChannelMode QProcess_ProcessChannelMode(const QProcess* self) {
	return self->processChannelMode();
}
void QProcess_SetProcessChannelMode(QProcess* self, ProcessChannelMode mode) {
	self->setProcessChannelMode(mode);
}
InputChannelMode QProcess_InputChannelMode(const QProcess* self) {
	return self->inputChannelMode();
}
void QProcess_SetInputChannelMode(QProcess* self, InputChannelMode mode) {
	self->setInputChannelMode(mode);
}
ProcessChannel QProcess_ReadChannel(const QProcess* self) {
	return self->readChannel();
}
void QProcess_SetReadChannel(QProcess* self, ProcessChannel channel) {
	self->setReadChannel(channel);
}
void QProcess_CloseReadChannel(QProcess* self, ProcessChannel channel) {
	self->closeReadChannel(channel);
}
void QProcess_CloseWriteChannel(QProcess* self) {
	self->closeWriteChannel();
}
void QProcess_SetStandardInputFile(QProcess* self, struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->setStandardInputFile(fileName_QString);
}
void QProcess_SetStandardOutputFile(QProcess* self, struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->setStandardOutputFile(fileName_QString);
}
void QProcess_SetStandardErrorFile(QProcess* self, struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->setStandardErrorFile(fileName_QString);
}
void QProcess_SetStandardOutputProcess(QProcess* self, QProcess* destination) {
	self->setStandardOutputProcess(destination);
}
struct miqt_string QProcess_NativeArguments(const QProcess* self) {
	QString _ret = self->nativeArguments();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QProcess_SetNativeArguments(QProcess* self, struct miqt_string arguments) {
	QString arguments_QString = QString::fromUtf8(arguments.data, arguments.len);
	self->setNativeArguments(arguments_QString);
}
CreateProcessArgumentModifier QProcess_CreateProcessArgumentsModifier(const QProcess* self) {
	return self->createProcessArgumentsModifier();
}
void QProcess_SetCreateProcessArgumentsModifier(QProcess* self, CreateProcessArgumentModifier modifier) {
	self->setCreateProcessArgumentsModifier(modifier);
}
struct miqt_string QProcess_WorkingDirectory(const QProcess* self) {
	QString _ret = self->workingDirectory();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QProcess_SetWorkingDirectory(QProcess* self, struct miqt_string dir) {
	QString dir_QString = QString::fromUtf8(dir.data, dir.len);
	self->setWorkingDirectory(dir_QString);
}
void QProcess_SetEnvironment(QProcess* self, struct miqt_array /* of struct miqt_string */  environment) {
	QStringList environment_QList;
	environment_QList.reserve(environment.len);
	struct miqt_string* environment_arr = static_cast<struct miqt_string*>(environment.data);
	for(size_t i = 0; i < environment.len; ++i) {
		QString environment_arr_i_QString = QString::fromUtf8(environment_arr[i].data, environment_arr[i].len);
		environment_QList.push_back(environment_arr_i_QString);
	}
	self->setEnvironment(environment_QList);
}
struct miqt_array /* of struct miqt_string */  QProcess_Environment(const QProcess* self) {
	QStringList _ret = self->environment();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QProcess_SetProcessEnvironment(QProcess* self, QProcessEnvironment* environment) {
	self->setProcessEnvironment(*environment);
}
QProcessEnvironment* QProcess_ProcessEnvironment(const QProcess* self) {
	return new QProcessEnvironment(self->processEnvironment());
}
int QProcess_Error(const QProcess* self) {
	QProcess::ProcessError _ret = self->error();
	return static_cast<int>(_ret);
}
int QProcess_State(const QProcess* self) {
	QProcess::ProcessState _ret = self->state();
	return static_cast<int>(_ret);
}
long long QProcess_ProcessId(const QProcess* self) {
	qint64 _ret = self->processId();
	return static_cast<long long>(_ret);
}
bool QProcess_WaitForStarted(QProcess* self) {
	return self->waitForStarted();
}
bool QProcess_WaitForReadyRead(QProcess* self, int msecs) {
	return self->waitForReadyRead(static_cast<int>(msecs));
}
bool QProcess_WaitForBytesWritten(QProcess* self, int msecs) {
	return self->waitForBytesWritten(static_cast<int>(msecs));
}
bool QProcess_WaitForFinished(QProcess* self) {
	return self->waitForFinished();
}
struct miqt_string QProcess_ReadAllStandardOutput(QProcess* self) {
	QByteArray _qb = self->readAllStandardOutput();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
struct miqt_string QProcess_ReadAllStandardError(QProcess* self) {
	QByteArray _qb = self->readAllStandardError();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
int QProcess_ExitCode(const QProcess* self) {
	return self->exitCode();
}
int QProcess_ExitStatus(const QProcess* self) {
	QProcess::ExitStatus _ret = self->exitStatus();
	return static_cast<int>(_ret);
}
long long QProcess_BytesToWrite(const QProcess* self) {
	qint64 _ret = self->bytesToWrite();
	return static_cast<long long>(_ret);
}
bool QProcess_IsSequential(const QProcess* self) {
	return self->isSequential();
}
void QProcess_Close(QProcess* self) {
	self->close();
}
int QProcess_Execute(struct miqt_string program) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	return QProcess::execute(program_QString);
}
bool QProcess_StartDetachedWithProgram(struct miqt_string program) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	return QProcess::startDetached(program_QString);
}
struct miqt_array /* of struct miqt_string */  QProcess_SystemEnvironment() {
	QStringList _ret = QProcess::systemEnvironment();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_string QProcess_NullDevice() {
	QString _ret = QProcess::nullDevice();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QProcess_Terminate(QProcess* self) {
	self->terminate();
}
void QProcess_Kill(QProcess* self) {
	self->kill();
}
void QProcess_Finished(QProcess* self, int exitCode) {
	self->finished(static_cast<int>(exitCode));
}
void QProcess_connect_Finished(QProcess* self, intptr_t slot) {
	MiqtVirtualQProcess::connect(self, static_cast<void (QProcess::*)(int, QProcess::ExitStatus)>(&QProcess::finished), self, [=](int exitCode) {
		int sigval1 = exitCode;
		miqt_exec_callback_QProcess_Finished(slot, sigval1);
	});
}
void QProcess_ErrorOccurred(QProcess* self, int error) {
	self->errorOccurred(static_cast<QProcess::ProcessError>(error));
}
void QProcess_connect_ErrorOccurred(QProcess* self, intptr_t slot) {
	MiqtVirtualQProcess::connect(self, static_cast<void (QProcess::*)(QProcess::ProcessError)>(&QProcess::errorOccurred), self, [=](QProcess::ProcessError error) {
		QProcess::ProcessError error_ret = error;
		int sigval1 = static_cast<int>(error_ret);
		miqt_exec_callback_QProcess_ErrorOccurred(slot, sigval1);
	});
}
struct miqt_string QProcess_Tr2(const char* s, const char* c) {
	QString _ret = QProcess::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QProcess_Tr3(const char* s, const char* c, int n) {
	QString _ret = QProcess::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QProcess_Start22(QProcess* self, struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	QStringList arguments_QList;
	arguments_QList.reserve(arguments.len);
	struct miqt_string* arguments_arr = static_cast<struct miqt_string*>(arguments.data);
	for(size_t i = 0; i < arguments.len; ++i) {
		QString arguments_arr_i_QString = QString::fromUtf8(arguments_arr[i].data, arguments_arr[i].len);
		arguments_QList.push_back(arguments_arr_i_QString);
	}
	self->start(program_QString, arguments_QList);
}
void QProcess_Start3(QProcess* self, struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments, OpenMode mode) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	QStringList arguments_QList;
	arguments_QList.reserve(arguments.len);
	struct miqt_string* arguments_arr = static_cast<struct miqt_string*>(arguments.data);
	for(size_t i = 0; i < arguments.len; ++i) {
		QString arguments_arr_i_QString = QString::fromUtf8(arguments_arr[i].data, arguments_arr[i].len);
		arguments_QList.push_back(arguments_arr_i_QString);
	}
	self->start(program_QString, arguments_QList, mode);
}
void QProcess_Start1(QProcess* self, OpenMode mode) {
	self->start(mode);
}
void QProcess_StartCommand2(QProcess* self, struct miqt_string command, OpenMode mode) {
	QString command_QString = QString::fromUtf8(command.data, command.len);
	self->startCommand(command_QString, mode);
}
bool QProcess_StartDetached1(QProcess* self, long long* pid) {
	return self->startDetached(static_cast<qint64*>(pid));
}
void QProcess_SetStandardOutputFile2(QProcess* self, struct miqt_string fileName, OpenMode mode) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->setStandardOutputFile(fileName_QString, mode);
}
void QProcess_SetStandardErrorFile2(QProcess* self, struct miqt_string fileName, OpenMode mode) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->setStandardErrorFile(fileName_QString, mode);
}
bool QProcess_WaitForStarted1(QProcess* self, int msecs) {
	return self->waitForStarted(static_cast<int>(msecs));
}
bool QProcess_WaitForFinished1(QProcess* self, int msecs) {
	return self->waitForFinished(static_cast<int>(msecs));
}
int QProcess_Execute2(struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	QStringList arguments_QList;
	arguments_QList.reserve(arguments.len);
	struct miqt_string* arguments_arr = static_cast<struct miqt_string*>(arguments.data);
	for(size_t i = 0; i < arguments.len; ++i) {
		QString arguments_arr_i_QString = QString::fromUtf8(arguments_arr[i].data, arguments_arr[i].len);
		arguments_QList.push_back(arguments_arr_i_QString);
	}
	return QProcess::execute(program_QString, arguments_QList);
}
bool QProcess_StartDetached2(struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	QStringList arguments_QList;
	arguments_QList.reserve(arguments.len);
	struct miqt_string* arguments_arr = static_cast<struct miqt_string*>(arguments.data);
	for(size_t i = 0; i < arguments.len; ++i) {
		QString arguments_arr_i_QString = QString::fromUtf8(arguments_arr[i].data, arguments_arr[i].len);
		arguments_QList.push_back(arguments_arr_i_QString);
	}
	return QProcess::startDetached(program_QString, arguments_QList);
}
bool QProcess_StartDetached3(struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments, struct miqt_string workingDirectory) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	QStringList arguments_QList;
	arguments_QList.reserve(arguments.len);
	struct miqt_string* arguments_arr = static_cast<struct miqt_string*>(arguments.data);
	for(size_t i = 0; i < arguments.len; ++i) {
		QString arguments_arr_i_QString = QString::fromUtf8(arguments_arr[i].data, arguments_arr[i].len);
		arguments_QList.push_back(arguments_arr_i_QString);
	}
	QString workingDirectory_QString = QString::fromUtf8(workingDirectory.data, workingDirectory.len);
	return QProcess::startDetached(program_QString, arguments_QList, workingDirectory_QString);
}
bool QProcess_StartDetached4(struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments, struct miqt_string workingDirectory, long long* pid) {
	QString program_QString = QString::fromUtf8(program.data, program.len);
	QStringList arguments_QList;
	arguments_QList.reserve(arguments.len);
	struct miqt_string* arguments_arr = static_cast<struct miqt_string*>(arguments.data);
	for(size_t i = 0; i < arguments.len; ++i) {
		QString arguments_arr_i_QString = QString::fromUtf8(arguments_arr[i].data, arguments_arr[i].len);
		arguments_QList.push_back(arguments_arr_i_QString);
	}
	QString workingDirectory_QString = QString::fromUtf8(workingDirectory.data, workingDirectory.len);
	return QProcess::startDetached(program_QString, arguments_QList, workingDirectory_QString, static_cast<qint64*>(pid));
}
void QProcess_Finished2(QProcess* self, int exitCode, int exitStatus) {
	self->finished(static_cast<int>(exitCode), static_cast<QProcess::ExitStatus>(exitStatus));
}
void QProcess_connect_Finished2(QProcess* self, intptr_t slot) {
	MiqtVirtualQProcess::connect(self, static_cast<void (QProcess::*)(int, QProcess::ExitStatus)>(&QProcess::finished), self, [=](int exitCode, QProcess::ExitStatus exitStatus) {
		int sigval1 = exitCode;
		QProcess::ExitStatus exitStatus_ret = exitStatus;
		int sigval2 = static_cast<int>(exitStatus_ret);
		miqt_exec_callback_QProcess_Finished2(slot, sigval1, sigval2);
	});
}
void QProcess_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQProcess*>( (QProcess*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QProcess_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQProcess*)(self) )->virtualbase_MetaObject();
}
void QProcess_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQProcess*>( (QProcess*)(self) )->handle__Metacast = slot;
}
void* QProcess_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQProcess*)(self) )->virtualbase_Metacast(param1);
}
void QProcess_Delete(QProcess* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQProcess*>( self );
	} else {
		delete self;
	}
}
void QProcess__CreateProcessArguments_Delete(QProcess__CreateProcessArguments* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QProcess::CreateProcessArguments*>( self );
	} else {
		delete self;
	}
}
