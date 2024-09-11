#include <QByteArray>
#include <QFont>
#include <QFontDatabase>
#include <QFontInfo>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include "qfontdatabase.h"

#include "gen_qfontdatabase.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QFontDatabase* QFontDatabase_new() {
	return new QFontDatabase();
}

void QFontDatabase_StandardSizes(int** _out, size_t* _out_len) {
	QList<int> ret = QFontDatabase::standardSizes();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* __out = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		__out[i] = ret[i];
	}
	*_out = __out;
	*_out_len = ret.length();
}

void QFontDatabase_WritingSystems(const QFontDatabase* self, uintptr_t** _out, size_t* _out_len) {
	QList<QFontDatabase::WritingSystem> ret = self->writingSystems();
	// Convert QList<> from C++ memory to manually-managed C memory
	uintptr_t* __out = static_cast<uintptr_t*>(malloc(sizeof(uintptr_t) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		__out[i] = ret[i];
	}
	*_out = __out;
	*_out_len = ret.length();
}

void QFontDatabase_WritingSystemsWithFamily(const QFontDatabase* self, const char* family, size_t family_Strlen, uintptr_t** _out, size_t* _out_len) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QList<QFontDatabase::WritingSystem> ret = self->writingSystems(family_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	uintptr_t* __out = static_cast<uintptr_t*>(malloc(sizeof(uintptr_t) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		__out[i] = ret[i];
	}
	*_out = __out;
	*_out_len = ret.length();
}

void QFontDatabase_Families(const QFontDatabase* self, char*** _out, int** _out_Lengths, size_t* _out_len) {
	QStringList ret = self->families();
	// Convert QStringList from C++ memory to manually-managed C memory
	char** __out = static_cast<char**>(malloc(sizeof(char*) * ret.length()));
	int* __out_Lengths = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray b = ret[i].toUtf8();
		__out[i] = static_cast<char*>(malloc(b.length()));
		memcpy(__out[i], b.data(), b.length());
		__out_Lengths[i] = b.length();
	}
	*_out = __out;
	*_out_Lengths = __out_Lengths;
	*_out_len = ret.length();
}

void QFontDatabase_Styles(const QFontDatabase* self, const char* family, size_t family_Strlen, char*** _out, int** _out_Lengths, size_t* _out_len) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QStringList ret = self->styles(family_QString);
	// Convert QStringList from C++ memory to manually-managed C memory
	char** __out = static_cast<char**>(malloc(sizeof(char*) * ret.length()));
	int* __out_Lengths = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray b = ret[i].toUtf8();
		__out[i] = static_cast<char*>(malloc(b.length()));
		memcpy(__out[i], b.data(), b.length());
		__out_Lengths[i] = b.length();
	}
	*_out = __out;
	*_out_Lengths = __out_Lengths;
	*_out_len = ret.length();
}

void QFontDatabase_PointSizes(QFontDatabase* self, const char* family, size_t family_Strlen, int** _out, size_t* _out_len) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QList<int> ret = self->pointSizes(family_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	int* __out = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		__out[i] = ret[i];
	}
	*_out = __out;
	*_out_len = ret.length();
}

void QFontDatabase_SmoothSizes(QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen, int** _out, size_t* _out_len) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	QList<int> ret = self->smoothSizes(family_QString, style_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	int* __out = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		__out[i] = ret[i];
	}
	*_out = __out;
	*_out_len = ret.length();
}

void QFontDatabase_StyleString(QFontDatabase* self, QFont* font, char** _out, int* _out_Strlen) {
	QString ret = self->styleString(*font);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QFontDatabase_StyleStringWithFontInfo(QFontDatabase* self, QFontInfo* fontInfo, char** _out, int* _out_Strlen) {
	QString ret = self->styleString(*fontInfo);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

QFont* QFontDatabase_Font(const QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen, int pointSize) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	QFont ret = self->font(family_QString, style_QString, static_cast<int>(pointSize));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QFont*>(new QFont(ret));
}

bool QFontDatabase_IsBitmapScalable(const QFontDatabase* self, const char* family, size_t family_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	return self->isBitmapScalable(family_QString);
}

bool QFontDatabase_IsSmoothlyScalable(const QFontDatabase* self, const char* family, size_t family_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	return self->isSmoothlyScalable(family_QString);
}

bool QFontDatabase_IsScalable(const QFontDatabase* self, const char* family, size_t family_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	return self->isScalable(family_QString);
}

bool QFontDatabase_IsFixedPitch(const QFontDatabase* self, const char* family, size_t family_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	return self->isFixedPitch(family_QString);
}

bool QFontDatabase_Italic(const QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	return self->italic(family_QString, style_QString);
}

bool QFontDatabase_Bold(const QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	return self->bold(family_QString, style_QString);
}

int QFontDatabase_Weight(const QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	return self->weight(family_QString, style_QString);
}

bool QFontDatabase_HasFamily(const QFontDatabase* self, const char* family, size_t family_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	return self->hasFamily(family_QString);
}

bool QFontDatabase_IsPrivateFamily(const QFontDatabase* self, const char* family, size_t family_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	return self->isPrivateFamily(family_QString);
}

void QFontDatabase_WritingSystemName(uintptr_t writingSystem, char** _out, int* _out_Strlen) {
	QString ret = QFontDatabase::writingSystemName(static_cast<QFontDatabase::WritingSystem>(writingSystem));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QFontDatabase_WritingSystemSample(uintptr_t writingSystem, char** _out, int* _out_Strlen) {
	QString ret = QFontDatabase::writingSystemSample(static_cast<QFontDatabase::WritingSystem>(writingSystem));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

int QFontDatabase_AddApplicationFont(const char* fileName, size_t fileName_Strlen) {
	QString fileName_QString = QString::fromUtf8(fileName, fileName_Strlen);
	return QFontDatabase::addApplicationFont(fileName_QString);
}

int QFontDatabase_AddApplicationFontFromData(QByteArray* fontData) {
	return QFontDatabase::addApplicationFontFromData(*fontData);
}

void QFontDatabase_ApplicationFontFamilies(int id, char*** _out, int** _out_Lengths, size_t* _out_len) {
	QStringList ret = QFontDatabase::applicationFontFamilies(static_cast<int>(id));
	// Convert QStringList from C++ memory to manually-managed C memory
	char** __out = static_cast<char**>(malloc(sizeof(char*) * ret.length()));
	int* __out_Lengths = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray b = ret[i].toUtf8();
		__out[i] = static_cast<char*>(malloc(b.length()));
		memcpy(__out[i], b.data(), b.length());
		__out_Lengths[i] = b.length();
	}
	*_out = __out;
	*_out_Lengths = __out_Lengths;
	*_out_len = ret.length();
}

bool QFontDatabase_RemoveApplicationFont(int id) {
	return QFontDatabase::removeApplicationFont(static_cast<int>(id));
}

bool QFontDatabase_RemoveAllApplicationFonts() {
	return QFontDatabase::removeAllApplicationFonts();
}

bool QFontDatabase_SupportsThreadedFontRendering() {
	return QFontDatabase::supportsThreadedFontRendering();
}

QFont* QFontDatabase_SystemFont(uintptr_t typeVal) {
	QFont ret = QFontDatabase::systemFont(static_cast<QFontDatabase::SystemFont>(typeVal));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QFont*>(new QFont(ret));
}

void QFontDatabase_Families1(const QFontDatabase* self, uintptr_t writingSystem, char*** _out, int** _out_Lengths, size_t* _out_len) {
	QStringList ret = self->families(static_cast<QFontDatabase::WritingSystem>(writingSystem));
	// Convert QStringList from C++ memory to manually-managed C memory
	char** __out = static_cast<char**>(malloc(sizeof(char*) * ret.length()));
	int* __out_Lengths = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray b = ret[i].toUtf8();
		__out[i] = static_cast<char*>(malloc(b.length()));
		memcpy(__out[i], b.data(), b.length());
		__out_Lengths[i] = b.length();
	}
	*_out = __out;
	*_out_Lengths = __out_Lengths;
	*_out_len = ret.length();
}

void QFontDatabase_PointSizes2(QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen, int** _out, size_t* _out_len) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	QList<int> ret = self->pointSizes(family_QString, style_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	int* __out = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		__out[i] = ret[i];
	}
	*_out = __out;
	*_out_len = ret.length();
}

bool QFontDatabase_IsBitmapScalable2(const QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	return self->isBitmapScalable(family_QString, style_QString);
}

bool QFontDatabase_IsSmoothlyScalable2(const QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	return self->isSmoothlyScalable(family_QString, style_QString);
}

bool QFontDatabase_IsScalable2(const QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	return self->isScalable(family_QString, style_QString);
}

bool QFontDatabase_IsFixedPitch2(const QFontDatabase* self, const char* family, size_t family_Strlen, const char* style, size_t style_Strlen) {
	QString family_QString = QString::fromUtf8(family, family_Strlen);
	QString style_QString = QString::fromUtf8(style, style_Strlen);
	return self->isFixedPitch(family_QString, style_QString);
}

void QFontDatabase_Delete(QFontDatabase* self) {
	delete self;
}

