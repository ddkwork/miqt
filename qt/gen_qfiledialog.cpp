#include <QAbstractItemDelegate>
#include <QAbstractProxyModel>
#include <QByteArray>
#include <QDir>
#include <QFileDialog>
#include <QFileIconProvider>
#include <QList>
#include <QMetaObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <QWidget>
#include "qfiledialog.h"
#include "gen_qfiledialog.h"
#include "_cgo_export.h"

QFileDialog* QFileDialog_new(QWidget* parent, int f) {
	return new QFileDialog(parent, static_cast<Qt::WindowFlags>(f));
}

QFileDialog* QFileDialog_new2() {
	return new QFileDialog();
}

QFileDialog* QFileDialog_new3(QWidget* parent) {
	return new QFileDialog(parent);
}

QFileDialog* QFileDialog_new4(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	return new QFileDialog(parent, caption_QString);
}

QFileDialog* QFileDialog_new5(QWidget* parent, struct miqt_string* caption, struct miqt_string* directory) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString directory_QString = QString::fromUtf8(&directory->data, directory->len);
	return new QFileDialog(parent, caption_QString, directory_QString);
}

QFileDialog* QFileDialog_new6(QWidget* parent, struct miqt_string* caption, struct miqt_string* directory, struct miqt_string* filter) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString directory_QString = QString::fromUtf8(&directory->data, directory->len);
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	return new QFileDialog(parent, caption_QString, directory_QString, filter_QString);
}

QMetaObject* QFileDialog_MetaObject(const QFileDialog* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QFileDialog_Tr(const char* s) {
	QString _ret = QFileDialog::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_TrUtf8(const char* s) {
	QString _ret = QFileDialog::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QFileDialog_SetDirectory(QFileDialog* self, struct miqt_string* directory) {
	QString directory_QString = QString::fromUtf8(&directory->data, directory->len);
	self->setDirectory(directory_QString);
}

void QFileDialog_SetDirectoryWithDirectory(QFileDialog* self, QDir* directory) {
	self->setDirectory(*directory);
}

QDir* QFileDialog_Directory(const QFileDialog* self) {
	return new QDir(self->directory());
}

void QFileDialog_SetDirectoryUrl(QFileDialog* self, QUrl* directory) {
	self->setDirectoryUrl(*directory);
}

QUrl* QFileDialog_DirectoryUrl(const QFileDialog* self) {
	return new QUrl(self->directoryUrl());
}

void QFileDialog_SelectFile(QFileDialog* self, struct miqt_string* filename) {
	QString filename_QString = QString::fromUtf8(&filename->data, filename->len);
	self->selectFile(filename_QString);
}

struct miqt_array* QFileDialog_SelectedFiles(const QFileDialog* self) {
	QStringList _ret = self->selectedFiles();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QFileDialog_SelectUrl(QFileDialog* self, QUrl* url) {
	self->selectUrl(*url);
}

struct miqt_array* QFileDialog_SelectedUrls(const QFileDialog* self) {
	QList<QUrl> _ret = self->selectedUrls();
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QFileDialog_SetNameFilterDetailsVisible(QFileDialog* self, bool enabled) {
	self->setNameFilterDetailsVisible(enabled);
}

bool QFileDialog_IsNameFilterDetailsVisible(const QFileDialog* self) {
	return self->isNameFilterDetailsVisible();
}

void QFileDialog_SetNameFilter(QFileDialog* self, struct miqt_string* filter) {
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	self->setNameFilter(filter_QString);
}

void QFileDialog_SetNameFilters(QFileDialog* self, struct miqt_array* /* of struct miqt_string* */ filters) {
	QList<QString> filters_QList;
	filters_QList.reserve(filters->len);
	struct miqt_string** filters_arr = static_cast<struct miqt_string**>(filters->data);
	for(size_t i = 0; i < filters->len; ++i) {
		QString filters_arr_i_QString = QString::fromUtf8(&filters_arr[i]->data, filters_arr[i]->len);
		filters_QList.push_back(filters_arr_i_QString);
	}
	self->setNameFilters(filters_QList);
}

struct miqt_array* QFileDialog_NameFilters(const QFileDialog* self) {
	QStringList _ret = self->nameFilters();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QFileDialog_SelectNameFilter(QFileDialog* self, struct miqt_string* filter) {
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	self->selectNameFilter(filter_QString);
}

struct miqt_string* QFileDialog_SelectedMimeTypeFilter(const QFileDialog* self) {
	QString _ret = self->selectedMimeTypeFilter();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_SelectedNameFilter(const QFileDialog* self) {
	QString _ret = self->selectedNameFilter();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QFileDialog_SetMimeTypeFilters(QFileDialog* self, struct miqt_array* /* of struct miqt_string* */ filters) {
	QList<QString> filters_QList;
	filters_QList.reserve(filters->len);
	struct miqt_string** filters_arr = static_cast<struct miqt_string**>(filters->data);
	for(size_t i = 0; i < filters->len; ++i) {
		QString filters_arr_i_QString = QString::fromUtf8(&filters_arr[i]->data, filters_arr[i]->len);
		filters_QList.push_back(filters_arr_i_QString);
	}
	self->setMimeTypeFilters(filters_QList);
}

struct miqt_array* QFileDialog_MimeTypeFilters(const QFileDialog* self) {
	QStringList _ret = self->mimeTypeFilters();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QFileDialog_SelectMimeTypeFilter(QFileDialog* self, struct miqt_string* filter) {
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	self->selectMimeTypeFilter(filter_QString);
}

int QFileDialog_Filter(const QFileDialog* self) {
	QDir::Filters _ret = self->filter();
	return static_cast<int>(_ret);
}

void QFileDialog_SetFilter(QFileDialog* self, int filters) {
	self->setFilter(static_cast<QDir::Filters>(filters));
}

void QFileDialog_SetViewMode(QFileDialog* self, int mode) {
	self->setViewMode(static_cast<QFileDialog::ViewMode>(mode));
}

int QFileDialog_ViewMode(const QFileDialog* self) {
	QFileDialog::ViewMode _ret = self->viewMode();
	return static_cast<int>(_ret);
}

void QFileDialog_SetFileMode(QFileDialog* self, int mode) {
	self->setFileMode(static_cast<QFileDialog::FileMode>(mode));
}

int QFileDialog_FileMode(const QFileDialog* self) {
	QFileDialog::FileMode _ret = self->fileMode();
	return static_cast<int>(_ret);
}

void QFileDialog_SetAcceptMode(QFileDialog* self, int mode) {
	self->setAcceptMode(static_cast<QFileDialog::AcceptMode>(mode));
}

int QFileDialog_AcceptMode(const QFileDialog* self) {
	QFileDialog::AcceptMode _ret = self->acceptMode();
	return static_cast<int>(_ret);
}

void QFileDialog_SetReadOnly(QFileDialog* self, bool enabled) {
	self->setReadOnly(enabled);
}

bool QFileDialog_IsReadOnly(const QFileDialog* self) {
	return self->isReadOnly();
}

void QFileDialog_SetResolveSymlinks(QFileDialog* self, bool enabled) {
	self->setResolveSymlinks(enabled);
}

bool QFileDialog_ResolveSymlinks(const QFileDialog* self) {
	return self->resolveSymlinks();
}

void QFileDialog_SetSidebarUrls(QFileDialog* self, struct miqt_array* /* of QUrl* */ urls) {
	QList<QUrl> urls_QList;
	urls_QList.reserve(urls->len);
	QUrl** urls_arr = static_cast<QUrl**>(urls->data);
	for(size_t i = 0; i < urls->len; ++i) {
		urls_QList.push_back(*(urls_arr[i]));
	}
	self->setSidebarUrls(urls_QList);
}

struct miqt_array* QFileDialog_SidebarUrls(const QFileDialog* self) {
	QList<QUrl> _ret = self->sidebarUrls();
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

QByteArray* QFileDialog_SaveState(const QFileDialog* self) {
	return new QByteArray(self->saveState());
}

bool QFileDialog_RestoreState(QFileDialog* self, QByteArray* state) {
	return self->restoreState(*state);
}

void QFileDialog_SetConfirmOverwrite(QFileDialog* self, bool enabled) {
	self->setConfirmOverwrite(enabled);
}

bool QFileDialog_ConfirmOverwrite(const QFileDialog* self) {
	return self->confirmOverwrite();
}

void QFileDialog_SetDefaultSuffix(QFileDialog* self, struct miqt_string* suffix) {
	QString suffix_QString = QString::fromUtf8(&suffix->data, suffix->len);
	self->setDefaultSuffix(suffix_QString);
}

struct miqt_string* QFileDialog_DefaultSuffix(const QFileDialog* self) {
	QString _ret = self->defaultSuffix();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QFileDialog_SetHistory(QFileDialog* self, struct miqt_array* /* of struct miqt_string* */ paths) {
	QList<QString> paths_QList;
	paths_QList.reserve(paths->len);
	struct miqt_string** paths_arr = static_cast<struct miqt_string**>(paths->data);
	for(size_t i = 0; i < paths->len; ++i) {
		QString paths_arr_i_QString = QString::fromUtf8(&paths_arr[i]->data, paths_arr[i]->len);
		paths_QList.push_back(paths_arr_i_QString);
	}
	self->setHistory(paths_QList);
}

struct miqt_array* QFileDialog_History(const QFileDialog* self) {
	QStringList _ret = self->history();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QFileDialog_SetItemDelegate(QFileDialog* self, QAbstractItemDelegate* delegate) {
	self->setItemDelegate(delegate);
}

QAbstractItemDelegate* QFileDialog_ItemDelegate(const QFileDialog* self) {
	return self->itemDelegate();
}

void QFileDialog_SetIconProvider(QFileDialog* self, QFileIconProvider* provider) {
	self->setIconProvider(provider);
}

QFileIconProvider* QFileDialog_IconProvider(const QFileDialog* self) {
	return self->iconProvider();
}

void QFileDialog_SetLabelText(QFileDialog* self, int label, struct miqt_string* text) {
	QString text_QString = QString::fromUtf8(&text->data, text->len);
	self->setLabelText(static_cast<QFileDialog::DialogLabel>(label), text_QString);
}

struct miqt_string* QFileDialog_LabelText(const QFileDialog* self, int label) {
	QString _ret = self->labelText(static_cast<QFileDialog::DialogLabel>(label));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QFileDialog_SetSupportedSchemes(QFileDialog* self, struct miqt_array* /* of struct miqt_string* */ schemes) {
	QList<QString> schemes_QList;
	schemes_QList.reserve(schemes->len);
	struct miqt_string** schemes_arr = static_cast<struct miqt_string**>(schemes->data);
	for(size_t i = 0; i < schemes->len; ++i) {
		QString schemes_arr_i_QString = QString::fromUtf8(&schemes_arr[i]->data, schemes_arr[i]->len);
		schemes_QList.push_back(schemes_arr_i_QString);
	}
	self->setSupportedSchemes(schemes_QList);
}

struct miqt_array* QFileDialog_SupportedSchemes(const QFileDialog* self) {
	QStringList _ret = self->supportedSchemes();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QFileDialog_SetProxyModel(QFileDialog* self, QAbstractProxyModel* model) {
	self->setProxyModel(model);
}

QAbstractProxyModel* QFileDialog_ProxyModel(const QFileDialog* self) {
	return self->proxyModel();
}

void QFileDialog_SetOption(QFileDialog* self, int option) {
	self->setOption(static_cast<QFileDialog::Option>(option));
}

bool QFileDialog_TestOption(const QFileDialog* self, int option) {
	return self->testOption(static_cast<QFileDialog::Option>(option));
}

void QFileDialog_SetOptions(QFileDialog* self, int options) {
	self->setOptions(static_cast<QFileDialog::Options>(options));
}

int QFileDialog_Options(const QFileDialog* self) {
	QFileDialog::Options _ret = self->options();
	return static_cast<int>(_ret);
}

void QFileDialog_SetVisible(QFileDialog* self, bool visible) {
	self->setVisible(visible);
}

void QFileDialog_FileSelected(QFileDialog* self, struct miqt_string* file) {
	QString file_QString = QString::fromUtf8(&file->data, file->len);
	self->fileSelected(file_QString);
}

void QFileDialog_connect_FileSelected(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QString&)>(&QFileDialog::fileSelected), self, [=](const QString& file) {
		const QString file_ret = file;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray file_b = file_ret.toUtf8();
		struct miqt_string* sigval1 = miqt_strdup(file_b.data(), file_b.length());
		miqt_exec_callback_QFileDialog_FileSelected(slot, sigval1);
	});
}

void QFileDialog_FilesSelected(QFileDialog* self, struct miqt_array* /* of struct miqt_string* */ files) {
	QList<QString> files_QList;
	files_QList.reserve(files->len);
	struct miqt_string** files_arr = static_cast<struct miqt_string**>(files->data);
	for(size_t i = 0; i < files->len; ++i) {
		QString files_arr_i_QString = QString::fromUtf8(&files_arr[i]->data, files_arr[i]->len);
		files_QList.push_back(files_arr_i_QString);
	}
	self->filesSelected(files_QList);
}

void QFileDialog_connect_FilesSelected(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QStringList&)>(&QFileDialog::filesSelected), self, [=](const QStringList& files) {
		const QStringList& files_ret = files;
		// Convert QList<> from C++ memory to manually-managed C memory
		struct miqt_string** files_arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * files_ret.length()));
		for (size_t i = 0, e = files_ret.length(); i < e; ++i) {
			QString files_lv_ret = files_ret[i];
			// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
			QByteArray files_lv_b = files_lv_ret.toUtf8();
			files_arr[i] = miqt_strdup(files_lv_b.data(), files_lv_b.length());
		}
		struct miqt_array* files_out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
		files_out->len = files_ret.length();
		files_out->data = static_cast<void*>(files_arr);
		struct miqt_array* sigval1 = files_out;
		miqt_exec_callback_QFileDialog_FilesSelected(slot, sigval1);
	});
}

void QFileDialog_CurrentChanged(QFileDialog* self, struct miqt_string* path) {
	QString path_QString = QString::fromUtf8(&path->data, path->len);
	self->currentChanged(path_QString);
}

void QFileDialog_connect_CurrentChanged(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QString&)>(&QFileDialog::currentChanged), self, [=](const QString& path) {
		const QString path_ret = path;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray path_b = path_ret.toUtf8();
		struct miqt_string* sigval1 = miqt_strdup(path_b.data(), path_b.length());
		miqt_exec_callback_QFileDialog_CurrentChanged(slot, sigval1);
	});
}

void QFileDialog_DirectoryEntered(QFileDialog* self, struct miqt_string* directory) {
	QString directory_QString = QString::fromUtf8(&directory->data, directory->len);
	self->directoryEntered(directory_QString);
}

void QFileDialog_connect_DirectoryEntered(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QString&)>(&QFileDialog::directoryEntered), self, [=](const QString& directory) {
		const QString directory_ret = directory;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray directory_b = directory_ret.toUtf8();
		struct miqt_string* sigval1 = miqt_strdup(directory_b.data(), directory_b.length());
		miqt_exec_callback_QFileDialog_DirectoryEntered(slot, sigval1);
	});
}

void QFileDialog_UrlSelected(QFileDialog* self, QUrl* url) {
	self->urlSelected(*url);
}

void QFileDialog_connect_UrlSelected(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QUrl&)>(&QFileDialog::urlSelected), self, [=](const QUrl& url) {
		const QUrl& url_ret = url;
		// Cast returned reference into pointer
		QUrl* sigval1 = const_cast<QUrl*>(&url_ret);
		miqt_exec_callback_QFileDialog_UrlSelected(slot, sigval1);
	});
}

void QFileDialog_UrlsSelected(QFileDialog* self, struct miqt_array* /* of QUrl* */ urls) {
	QList<QUrl> urls_QList;
	urls_QList.reserve(urls->len);
	QUrl** urls_arr = static_cast<QUrl**>(urls->data);
	for(size_t i = 0; i < urls->len; ++i) {
		urls_QList.push_back(*(urls_arr[i]));
	}
	self->urlsSelected(urls_QList);
}

void QFileDialog_connect_UrlsSelected(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QList<QUrl>&)>(&QFileDialog::urlsSelected), self, [=](const QList<QUrl>& urls) {
		const QList<QUrl>& urls_ret = urls;
		// Convert QList<> from C++ memory to manually-managed C memory
		QUrl** urls_arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * urls_ret.length()));
		for (size_t i = 0, e = urls_ret.length(); i < e; ++i) {
			urls_arr[i] = new QUrl(urls_ret[i]);
		}
		struct miqt_array* urls_out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
		urls_out->len = urls_ret.length();
		urls_out->data = static_cast<void*>(urls_arr);
		struct miqt_array* sigval1 = urls_out;
		miqt_exec_callback_QFileDialog_UrlsSelected(slot, sigval1);
	});
}

void QFileDialog_CurrentUrlChanged(QFileDialog* self, QUrl* url) {
	self->currentUrlChanged(*url);
}

void QFileDialog_connect_CurrentUrlChanged(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QUrl&)>(&QFileDialog::currentUrlChanged), self, [=](const QUrl& url) {
		const QUrl& url_ret = url;
		// Cast returned reference into pointer
		QUrl* sigval1 = const_cast<QUrl*>(&url_ret);
		miqt_exec_callback_QFileDialog_CurrentUrlChanged(slot, sigval1);
	});
}

void QFileDialog_DirectoryUrlEntered(QFileDialog* self, QUrl* directory) {
	self->directoryUrlEntered(*directory);
}

void QFileDialog_connect_DirectoryUrlEntered(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QUrl&)>(&QFileDialog::directoryUrlEntered), self, [=](const QUrl& directory) {
		const QUrl& directory_ret = directory;
		// Cast returned reference into pointer
		QUrl* sigval1 = const_cast<QUrl*>(&directory_ret);
		miqt_exec_callback_QFileDialog_DirectoryUrlEntered(slot, sigval1);
	});
}

void QFileDialog_FilterSelected(QFileDialog* self, struct miqt_string* filter) {
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	self->filterSelected(filter_QString);
}

void QFileDialog_connect_FilterSelected(QFileDialog* self, void* slot) {
	QFileDialog::connect(self, static_cast<void (QFileDialog::*)(const QString&)>(&QFileDialog::filterSelected), self, [=](const QString& filter) {
		const QString filter_ret = filter;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray filter_b = filter_ret.toUtf8();
		struct miqt_string* sigval1 = miqt_strdup(filter_b.data(), filter_b.length());
		miqt_exec_callback_QFileDialog_FilterSelected(slot, sigval1);
	});
}

struct miqt_string* QFileDialog_GetOpenFileName() {
	QString _ret = QFileDialog::getOpenFileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QUrl* QFileDialog_GetOpenFileUrl() {
	return new QUrl(QFileDialog::getOpenFileUrl());
}

struct miqt_string* QFileDialog_GetSaveFileName() {
	QString _ret = QFileDialog::getSaveFileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QUrl* QFileDialog_GetSaveFileUrl() {
	return new QUrl(QFileDialog::getSaveFileUrl());
}

struct miqt_string* QFileDialog_GetExistingDirectory() {
	QString _ret = QFileDialog::getExistingDirectory();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QUrl* QFileDialog_GetExistingDirectoryUrl() {
	return new QUrl(QFileDialog::getExistingDirectoryUrl());
}

struct miqt_array* QFileDialog_GetOpenFileNames() {
	QStringList _ret = QFileDialog::getOpenFileNames();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QFileDialog_GetOpenFileUrls() {
	QList<QUrl> _ret = QFileDialog::getOpenFileUrls();
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QFileDialog_SaveFileContent(QByteArray* fileContent) {
	QFileDialog::saveFileContent(*fileContent);
}

struct miqt_string* QFileDialog_Tr2(const char* s, const char* c) {
	QString _ret = QFileDialog::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_Tr3(const char* s, const char* c, int n) {
	QString _ret = QFileDialog::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_TrUtf82(const char* s, const char* c) {
	QString _ret = QFileDialog::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QFileDialog::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QFileDialog_SetOption2(QFileDialog* self, int option, bool on) {
	self->setOption(static_cast<QFileDialog::Option>(option), on);
}

struct miqt_string* QFileDialog_GetOpenFileName1(QWidget* parent) {
	QString _ret = QFileDialog::getOpenFileName(parent);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetOpenFileName2(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString _ret = QFileDialog::getOpenFileName(parent, caption_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetOpenFileName3(QWidget* parent, struct miqt_string* caption, struct miqt_string* dir) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString dir_QString = QString::fromUtf8(&dir->data, dir->len);
	QString _ret = QFileDialog::getOpenFileName(parent, caption_QString, dir_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetOpenFileName4(QWidget* parent, struct miqt_string* caption, struct miqt_string* dir, struct miqt_string* filter) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString dir_QString = QString::fromUtf8(&dir->data, dir->len);
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	QString _ret = QFileDialog::getOpenFileName(parent, caption_QString, dir_QString, filter_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QUrl* QFileDialog_GetOpenFileUrl1(QWidget* parent) {
	return new QUrl(QFileDialog::getOpenFileUrl(parent));
}

QUrl* QFileDialog_GetOpenFileUrl2(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	return new QUrl(QFileDialog::getOpenFileUrl(parent, caption_QString));
}

QUrl* QFileDialog_GetOpenFileUrl3(QWidget* parent, struct miqt_string* caption, QUrl* dir) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	return new QUrl(QFileDialog::getOpenFileUrl(parent, caption_QString, *dir));
}

QUrl* QFileDialog_GetOpenFileUrl4(QWidget* parent, struct miqt_string* caption, QUrl* dir, struct miqt_string* filter) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	return new QUrl(QFileDialog::getOpenFileUrl(parent, caption_QString, *dir, filter_QString));
}

struct miqt_string* QFileDialog_GetSaveFileName1(QWidget* parent) {
	QString _ret = QFileDialog::getSaveFileName(parent);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetSaveFileName2(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString _ret = QFileDialog::getSaveFileName(parent, caption_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetSaveFileName3(QWidget* parent, struct miqt_string* caption, struct miqt_string* dir) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString dir_QString = QString::fromUtf8(&dir->data, dir->len);
	QString _ret = QFileDialog::getSaveFileName(parent, caption_QString, dir_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetSaveFileName4(QWidget* parent, struct miqt_string* caption, struct miqt_string* dir, struct miqt_string* filter) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString dir_QString = QString::fromUtf8(&dir->data, dir->len);
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	QString _ret = QFileDialog::getSaveFileName(parent, caption_QString, dir_QString, filter_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QUrl* QFileDialog_GetSaveFileUrl1(QWidget* parent) {
	return new QUrl(QFileDialog::getSaveFileUrl(parent));
}

QUrl* QFileDialog_GetSaveFileUrl2(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	return new QUrl(QFileDialog::getSaveFileUrl(parent, caption_QString));
}

QUrl* QFileDialog_GetSaveFileUrl3(QWidget* parent, struct miqt_string* caption, QUrl* dir) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	return new QUrl(QFileDialog::getSaveFileUrl(parent, caption_QString, *dir));
}

QUrl* QFileDialog_GetSaveFileUrl4(QWidget* parent, struct miqt_string* caption, QUrl* dir, struct miqt_string* filter) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	return new QUrl(QFileDialog::getSaveFileUrl(parent, caption_QString, *dir, filter_QString));
}

struct miqt_string* QFileDialog_GetExistingDirectory1(QWidget* parent) {
	QString _ret = QFileDialog::getExistingDirectory(parent);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetExistingDirectory2(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString _ret = QFileDialog::getExistingDirectory(parent, caption_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetExistingDirectory3(QWidget* parent, struct miqt_string* caption, struct miqt_string* dir) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString dir_QString = QString::fromUtf8(&dir->data, dir->len);
	QString _ret = QFileDialog::getExistingDirectory(parent, caption_QString, dir_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QFileDialog_GetExistingDirectory4(QWidget* parent, struct miqt_string* caption, struct miqt_string* dir, int options) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString dir_QString = QString::fromUtf8(&dir->data, dir->len);
	QString _ret = QFileDialog::getExistingDirectory(parent, caption_QString, dir_QString, static_cast<QFileDialog::Options>(options));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QUrl* QFileDialog_GetExistingDirectoryUrl1(QWidget* parent) {
	return new QUrl(QFileDialog::getExistingDirectoryUrl(parent));
}

QUrl* QFileDialog_GetExistingDirectoryUrl2(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	return new QUrl(QFileDialog::getExistingDirectoryUrl(parent, caption_QString));
}

QUrl* QFileDialog_GetExistingDirectoryUrl3(QWidget* parent, struct miqt_string* caption, QUrl* dir) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	return new QUrl(QFileDialog::getExistingDirectoryUrl(parent, caption_QString, *dir));
}

QUrl* QFileDialog_GetExistingDirectoryUrl4(QWidget* parent, struct miqt_string* caption, QUrl* dir, int options) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	return new QUrl(QFileDialog::getExistingDirectoryUrl(parent, caption_QString, *dir, static_cast<QFileDialog::Options>(options)));
}

QUrl* QFileDialog_GetExistingDirectoryUrl5(QWidget* parent, struct miqt_string* caption, QUrl* dir, int options, struct miqt_array* /* of struct miqt_string* */ supportedSchemes) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QList<QString> supportedSchemes_QList;
	supportedSchemes_QList.reserve(supportedSchemes->len);
	struct miqt_string** supportedSchemes_arr = static_cast<struct miqt_string**>(supportedSchemes->data);
	for(size_t i = 0; i < supportedSchemes->len; ++i) {
		QString supportedSchemes_arr_i_QString = QString::fromUtf8(&supportedSchemes_arr[i]->data, supportedSchemes_arr[i]->len);
		supportedSchemes_QList.push_back(supportedSchemes_arr_i_QString);
	}
	return new QUrl(QFileDialog::getExistingDirectoryUrl(parent, caption_QString, *dir, static_cast<QFileDialog::Options>(options), supportedSchemes_QList));
}

struct miqt_array* QFileDialog_GetOpenFileNames1(QWidget* parent) {
	QStringList _ret = QFileDialog::getOpenFileNames(parent);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QFileDialog_GetOpenFileNames2(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QStringList _ret = QFileDialog::getOpenFileNames(parent, caption_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QFileDialog_GetOpenFileNames3(QWidget* parent, struct miqt_string* caption, struct miqt_string* dir) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString dir_QString = QString::fromUtf8(&dir->data, dir->len);
	QStringList _ret = QFileDialog::getOpenFileNames(parent, caption_QString, dir_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QFileDialog_GetOpenFileNames4(QWidget* parent, struct miqt_string* caption, struct miqt_string* dir, struct miqt_string* filter) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString dir_QString = QString::fromUtf8(&dir->data, dir->len);
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	QStringList _ret = QFileDialog::getOpenFileNames(parent, caption_QString, dir_QString, filter_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QFileDialog_GetOpenFileUrls1(QWidget* parent) {
	QList<QUrl> _ret = QFileDialog::getOpenFileUrls(parent);
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QFileDialog_GetOpenFileUrls2(QWidget* parent, struct miqt_string* caption) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QList<QUrl> _ret = QFileDialog::getOpenFileUrls(parent, caption_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QFileDialog_GetOpenFileUrls3(QWidget* parent, struct miqt_string* caption, QUrl* dir) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QList<QUrl> _ret = QFileDialog::getOpenFileUrls(parent, caption_QString, *dir);
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QFileDialog_GetOpenFileUrls4(QWidget* parent, struct miqt_string* caption, QUrl* dir, struct miqt_string* filter) {
	QString caption_QString = QString::fromUtf8(&caption->data, caption->len);
	QString filter_QString = QString::fromUtf8(&filter->data, filter->len);
	QList<QUrl> _ret = QFileDialog::getOpenFileUrls(parent, caption_QString, *dir, filter_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QFileDialog_SaveFileContent2(QByteArray* fileContent, struct miqt_string* fileNameHint) {
	QString fileNameHint_QString = QString::fromUtf8(&fileNameHint->data, fileNameHint->len);
	QFileDialog::saveFileContent(*fileContent, fileNameHint_QString);
}

void QFileDialog_Delete(QFileDialog* self) {
	delete self;
}

