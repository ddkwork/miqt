package qt

import (
	"unsafe"
)

type QFileDialog__ViewMode int

const (
	QFileDialog__Detail QFileDialog__ViewMode = 0
	QFileDialog__List   QFileDialog__ViewMode = 1
)

type QFileDialog__FileMode int

const (
	QFileDialog__AnyFile       QFileDialog__FileMode = 0
	QFileDialog__ExistingFile  QFileDialog__FileMode = 1
	QFileDialog__Directory     QFileDialog__FileMode = 2
	QFileDialog__ExistingFiles QFileDialog__FileMode = 3
)

type QFileDialog__AcceptMode int

const (
	QFileDialog__AcceptOpen QFileDialog__AcceptMode = 0
	QFileDialog__AcceptSave QFileDialog__AcceptMode = 1
)

type QFileDialog__DialogLabel int

const (
	QFileDialog__LookIn   QFileDialog__DialogLabel = 0
	QFileDialog__FileName QFileDialog__DialogLabel = 1
	QFileDialog__FileType QFileDialog__DialogLabel = 2
	QFileDialog__Accept   QFileDialog__DialogLabel = 3
	QFileDialog__Reject   QFileDialog__DialogLabel = 4
)

type QFileDialog__Option int

const (
	QFileDialog__ShowDirsOnly                QFileDialog__Option = 1
	QFileDialog__DontResolveSymlinks         QFileDialog__Option = 2
	QFileDialog__DontConfirmOverwrite        QFileDialog__Option = 4
	QFileDialog__DontUseNativeDialog         QFileDialog__Option = 8
	QFileDialog__ReadOnly                    QFileDialog__Option = 16
	QFileDialog__HideNameFilterDetails       QFileDialog__Option = 32
	QFileDialog__DontUseCustomDirectoryIcons QFileDialog__Option = 64
)

type QFileDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQFileDialog constructs a new QFileDialog object.
func NewQFileDialog(parent *QWidget) *QFileDialog {
	g := newQFileDialog(QFileDialog_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQFileDialog2 constructs a new QFileDialog object.
func NewQFileDialog2(parent *QWidget, f WindowType) *QFileDialog {
	g := newQFileDialog(QFileDialog_new2(parent.cPointer(), (int)(f)))
	g.isSubclass = true
	return g
}

// NewQFileDialog3 constructs a new QFileDialog object.
func NewQFileDialog3() *QFileDialog {
	g := newQFileDialog(QFileDialog_new3())
	g.isSubclass = true
	return g
}

// NewQFileDialog4 constructs a new QFileDialog object.
func NewQFileDialog4(parent *QWidget, caption string) *QFileDialog {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	g := newQFileDialog(QFileDialog_new4(parent.cPointer(), caption_ms))
	g.isSubclass = true
	return g
}

// NewQFileDialog5 constructs a new QFileDialog object.
func NewQFileDialog5(parent *QWidget, caption string, directory string) *QFileDialog {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	g := newQFileDialog(QFileDialog_new5(parent.cPointer(), caption_ms, directory_ms))
	g.isSubclass = true
	return g
}

// NewQFileDialog6 constructs a new QFileDialog object.
func NewQFileDialog6(parent *QWidget, caption string, directory string, filter string) *QFileDialog {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	g := newQFileDialog(QFileDialog_new6(parent.cPointer(), caption_ms, directory_ms, filter_ms))
	g.isSubclass = true
	return g
}

func (this *QFileDialog) MetaObject() *QMetaObject {
	return newQMetaObject(QFileDialog_MetaObject(this.h))
}

func (this *QFileDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFileDialog_Metacast(this.h, param1_Cstring))
}

func QFileDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFileDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileDialog) SetDirectory(directory string) {
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	QFileDialog_SetDirectory(this.h, directory_ms)
}

func (this *QFileDialog) SetDirectoryWithDirectory(directory *QDir) {
	QFileDialog_SetDirectoryWithDirectory(this.h, directory.cPointer())
}

func (this *QFileDialog) Directory() *QDir {
	_goptr := newQDir(QFileDialog_Directory(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileDialog) SetDirectoryUrl(directory *QUrl) {
	QFileDialog_SetDirectoryUrl(this.h, directory.cPointer())
}

func (this *QFileDialog) DirectoryUrl() *QUrl {
	_goptr := newQUrl(QFileDialog_DirectoryUrl(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileDialog) SelectFile(filename string) {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	QFileDialog_SelectFile(this.h, filename_ms)
}

func (this *QFileDialog) SelectedFiles() []string {
	var _ma struct_miqt_array = QFileDialog_SelectedFiles(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QFileDialog) SelectUrl(url *QUrl) {
	QFileDialog_SelectUrl(this.h, url.cPointer())
}

func (this *QFileDialog) SelectedUrls() []QUrl {
	var _ma struct_miqt_array = QFileDialog_SelectedUrls(this.h)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QFileDialog) SetNameFilter(filter string) {
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	QFileDialog_SetNameFilter(this.h, filter_ms)
}

func (this *QFileDialog) SetNameFilters(filters []string) {
	filters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(filters))))
	defer free(unsafe.Pointer(filters_CArray))
	for i := range filters {
		filters_i_ms := struct_miqt_string{}
		filters_i_ms.data = CString(filters[i])
		filters_i_ms.len = size_t(len(filters[i]))
		defer free(unsafe.Pointer(filters_i_ms.data))
		filters_CArray[i] = filters_i_ms
	}
	filters_ma := struct_miqt_array{len: size_t(len(filters)), data: unsafe.Pointer(filters_CArray)}
	QFileDialog_SetNameFilters(this.h, filters_ma)
}

func (this *QFileDialog) NameFilters() []string {
	var _ma struct_miqt_array = QFileDialog_NameFilters(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QFileDialog) SelectNameFilter(filter string) {
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	QFileDialog_SelectNameFilter(this.h, filter_ms)
}

func (this *QFileDialog) SelectedMimeTypeFilter() string {
	var _ms struct_miqt_string = QFileDialog_SelectedMimeTypeFilter(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileDialog) SelectedNameFilter() string {
	var _ms struct_miqt_string = QFileDialog_SelectedNameFilter(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileDialog) SetMimeTypeFilters(filters []string) {
	filters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(filters))))
	defer free(unsafe.Pointer(filters_CArray))
	for i := range filters {
		filters_i_ms := struct_miqt_string{}
		filters_i_ms.data = CString(filters[i])
		filters_i_ms.len = size_t(len(filters[i]))
		defer free(unsafe.Pointer(filters_i_ms.data))
		filters_CArray[i] = filters_i_ms
	}
	filters_ma := struct_miqt_array{len: size_t(len(filters)), data: unsafe.Pointer(filters_CArray)}
	QFileDialog_SetMimeTypeFilters(this.h, filters_ma)
}

func (this *QFileDialog) MimeTypeFilters() []string {
	var _ma struct_miqt_array = QFileDialog_MimeTypeFilters(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QFileDialog) SelectMimeTypeFilter(filter string) {
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	QFileDialog_SelectMimeTypeFilter(this.h, filter_ms)
}

func (this *QFileDialog) Filter() Filter {
	return (Filter)(QFileDialog_Filter(this.h))
}

func (this *QFileDialog) SetFilter(filters Filter) {
	QFileDialog_SetFilter(this.h, (int)(filters))
}

func (this *QFileDialog) SetViewMode(mode ViewMode) {
	QFileDialog_SetViewMode(this.h, mode)
}

func (this *QFileDialog) ViewMode() ViewMode {
	xxxxxxxxx
}

func (this *QFileDialog) SetFileMode(mode FileMode) {
	QFileDialog_SetFileMode(this.h, mode)
}

func (this *QFileDialog) FileMode() FileMode {
	xxxxxxxxx
}

func (this *QFileDialog) SetAcceptMode(mode AcceptMode) {
	QFileDialog_SetAcceptMode(this.h, mode)
}

func (this *QFileDialog) AcceptMode() AcceptMode {
	xxxxxxxxx
}

func (this *QFileDialog) SetSidebarUrls(urls []QUrl) {
	urls_CArray := (*[0xffff]*QUrl)(malloc(size_t(8 * len(urls))))
	defer free(unsafe.Pointer(urls_CArray))
	for i := range urls {
		urls_CArray[i] = urls[i].cPointer()
	}
	urls_ma := struct_miqt_array{len: size_t(len(urls)), data: unsafe.Pointer(urls_CArray)}
	QFileDialog_SetSidebarUrls(this.h, urls_ma)
}

func (this *QFileDialog) SidebarUrls() []QUrl {
	var _ma struct_miqt_array = QFileDialog_SidebarUrls(this.h)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QFileDialog) SaveState() []byte {
	var _bytearray struct_miqt_string = QFileDialog_SaveState(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QFileDialog) RestoreState(state []byte) bool {
	state_alias := struct_miqt_string{}
	state_alias.data = (char)(unsafe.Pointer(&state[0]))
	state_alias.len = size_t(len(state))
	return (bool)(QFileDialog_RestoreState(this.h, state_alias))
}

func (this *QFileDialog) SetDefaultSuffix(suffix string) {
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	QFileDialog_SetDefaultSuffix(this.h, suffix_ms)
}

func (this *QFileDialog) DefaultSuffix() string {
	var _ms struct_miqt_string = QFileDialog_DefaultSuffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileDialog) SetHistory(paths []string) {
	paths_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(paths))))
	defer free(unsafe.Pointer(paths_CArray))
	for i := range paths {
		paths_i_ms := struct_miqt_string{}
		paths_i_ms.data = CString(paths[i])
		paths_i_ms.len = size_t(len(paths[i]))
		defer free(unsafe.Pointer(paths_i_ms.data))
		paths_CArray[i] = paths_i_ms
	}
	paths_ma := struct_miqt_array{len: size_t(len(paths)), data: unsafe.Pointer(paths_CArray)}
	QFileDialog_SetHistory(this.h, paths_ma)
}

func (this *QFileDialog) History() []string {
	var _ma struct_miqt_array = QFileDialog_History(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QFileDialog) SetItemDelegate(delegate *QAbstractItemDelegate) {
	QFileDialog_SetItemDelegate(this.h, delegate.cPointer())
}

func (this *QFileDialog) ItemDelegate() *QAbstractItemDelegate {
	return newQAbstractItemDelegate(QFileDialog_ItemDelegate(this.h))
}

func (this *QFileDialog) SetIconProvider(provider *QAbstractFileIconProvider) {
	QFileDialog_SetIconProvider(this.h, provider.cPointer())
}

func (this *QFileDialog) IconProvider() *QAbstractFileIconProvider {
	return newQAbstractFileIconProvider(QFileDialog_IconProvider(this.h))
}

func (this *QFileDialog) SetLabelText(label DialogLabel, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QFileDialog_SetLabelText(this.h, label, text_ms)
}

func (this *QFileDialog) LabelText(label DialogLabel) string {
	var _ms struct_miqt_string = QFileDialog_LabelText(this.h, label)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileDialog) SetSupportedSchemes(schemes []string) {
	schemes_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(schemes))))
	defer free(unsafe.Pointer(schemes_CArray))
	for i := range schemes {
		schemes_i_ms := struct_miqt_string{}
		schemes_i_ms.data = CString(schemes[i])
		schemes_i_ms.len = size_t(len(schemes[i]))
		defer free(unsafe.Pointer(schemes_i_ms.data))
		schemes_CArray[i] = schemes_i_ms
	}
	schemes_ma := struct_miqt_array{len: size_t(len(schemes)), data: unsafe.Pointer(schemes_CArray)}
	QFileDialog_SetSupportedSchemes(this.h, schemes_ma)
}

func (this *QFileDialog) SupportedSchemes() []string {
	var _ma struct_miqt_array = QFileDialog_SupportedSchemes(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QFileDialog) SetProxyModel(model *QAbstractProxyModel) {
	QFileDialog_SetProxyModel(this.h, model.cPointer())
}

func (this *QFileDialog) ProxyModel() *QAbstractProxyModel {
	return newQAbstractProxyModel(QFileDialog_ProxyModel(this.h))
}

func (this *QFileDialog) SetOption(option Option) {
	QFileDialog_SetOption(this.h, option)
}

func (this *QFileDialog) TestOption(option Option) bool {
	return (bool)(QFileDialog_TestOption(this.h, option))
}

func (this *QFileDialog) SetOptions(options Options) {
	QFileDialog_SetOptions(this.h, options)
}

func (this *QFileDialog) Options() Options {
	xxxxxxxxx
}

func (this *QFileDialog) SetVisible(visible bool) {
	QFileDialog_SetVisible(this.h, (bool)(visible))
}

func (this *QFileDialog) FileSelected(file string) {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	QFileDialog_FileSelected(this.h, file_ms)
}

func (this *QFileDialog) OnFileSelected(slot func(file string)) {
	QFileDialog_connect_FileSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_FileSelected
func miqt_exec_callback_QFileDialog_FileSelected(cb intptr_t, file struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(file string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var file_ms struct_miqt_string = file
	file_ret := GoStringN(file_ms.data, int(int64(file_ms.len)))
	free(unsafe.Pointer(file_ms.data))
	slotval1 := file_ret

	gofunc(slotval1)
}

func (this *QFileDialog) FilesSelected(files []string) {
	files_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(files))))
	defer free(unsafe.Pointer(files_CArray))
	for i := range files {
		files_i_ms := struct_miqt_string{}
		files_i_ms.data = CString(files[i])
		files_i_ms.len = size_t(len(files[i]))
		defer free(unsafe.Pointer(files_i_ms.data))
		files_CArray[i] = files_i_ms
	}
	files_ma := struct_miqt_array{len: size_t(len(files)), data: unsafe.Pointer(files_CArray)}
	QFileDialog_FilesSelected(this.h, files_ma)
}

func (this *QFileDialog) OnFilesSelected(slot func(files []string)) {
	QFileDialog_connect_FilesSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_FilesSelected
func miqt_exec_callback_QFileDialog_FilesSelected(cb intptr_t, files struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(files []string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var files_ma struct_miqt_array = files
	files_ret := make([]string, int(files_ma.len))
	files_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(files_ma.data)) // hey ya
	for i := 0; i < int(files_ma.len); i++ {
		var files_lv_ms struct_miqt_string = files_outCast[i]
		files_lv_ret := GoStringN(files_lv_ms.data, int(int64(files_lv_ms.len)))
		free(unsafe.Pointer(files_lv_ms.data))
		files_ret[i] = files_lv_ret
	}
	slotval1 := files_ret

	gofunc(slotval1)
}

func (this *QFileDialog) CurrentChanged(path string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QFileDialog_CurrentChanged(this.h, path_ms)
}

func (this *QFileDialog) OnCurrentChanged(slot func(path string)) {
	QFileDialog_connect_CurrentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_CurrentChanged
func miqt_exec_callback_QFileDialog_CurrentChanged(cb intptr_t, path struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(path string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var path_ms struct_miqt_string = path
	path_ret := GoStringN(path_ms.data, int(int64(path_ms.len)))
	free(unsafe.Pointer(path_ms.data))
	slotval1 := path_ret

	gofunc(slotval1)
}

func (this *QFileDialog) DirectoryEntered(directory string) {
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	QFileDialog_DirectoryEntered(this.h, directory_ms)
}

func (this *QFileDialog) OnDirectoryEntered(slot func(directory string)) {
	QFileDialog_connect_DirectoryEntered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_DirectoryEntered
func miqt_exec_callback_QFileDialog_DirectoryEntered(cb intptr_t, directory struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(directory string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var directory_ms struct_miqt_string = directory
	directory_ret := GoStringN(directory_ms.data, int(int64(directory_ms.len)))
	free(unsafe.Pointer(directory_ms.data))
	slotval1 := directory_ret

	gofunc(slotval1)
}

func (this *QFileDialog) UrlSelected(url *QUrl) {
	QFileDialog_UrlSelected(this.h, url.cPointer())
}

func (this *QFileDialog) OnUrlSelected(slot func(url *QUrl)) {
	QFileDialog_connect_UrlSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_UrlSelected
func miqt_exec_callback_QFileDialog_UrlSelected(cb intptr_t, url *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(url *QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUrl(url)

	gofunc(slotval1)
}

func (this *QFileDialog) UrlsSelected(urls []QUrl) {
	urls_CArray := (*[0xffff]*QUrl)(malloc(size_t(8 * len(urls))))
	defer free(unsafe.Pointer(urls_CArray))
	for i := range urls {
		urls_CArray[i] = urls[i].cPointer()
	}
	urls_ma := struct_miqt_array{len: size_t(len(urls)), data: unsafe.Pointer(urls_CArray)}
	QFileDialog_UrlsSelected(this.h, urls_ma)
}

func (this *QFileDialog) OnUrlsSelected(slot func(urls []QUrl)) {
	QFileDialog_connect_UrlsSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_UrlsSelected
func miqt_exec_callback_QFileDialog_UrlsSelected(cb intptr_t, urls struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(urls []QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var urls_ma struct_miqt_array = urls
	urls_ret := make([]QUrl, int(urls_ma.len))
	urls_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(urls_ma.data)) // hey ya
	for i := 0; i < int(urls_ma.len); i++ {
		urls_lv_goptr := newQUrl(urls_outCast[i])
		urls_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		urls_ret[i] = *urls_lv_goptr
	}
	slotval1 := urls_ret

	gofunc(slotval1)
}

func (this *QFileDialog) CurrentUrlChanged(url *QUrl) {
	QFileDialog_CurrentUrlChanged(this.h, url.cPointer())
}

func (this *QFileDialog) OnCurrentUrlChanged(slot func(url *QUrl)) {
	QFileDialog_connect_CurrentUrlChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_CurrentUrlChanged
func miqt_exec_callback_QFileDialog_CurrentUrlChanged(cb intptr_t, url *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(url *QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUrl(url)

	gofunc(slotval1)
}

func (this *QFileDialog) DirectoryUrlEntered(directory *QUrl) {
	QFileDialog_DirectoryUrlEntered(this.h, directory.cPointer())
}

func (this *QFileDialog) OnDirectoryUrlEntered(slot func(directory *QUrl)) {
	QFileDialog_connect_DirectoryUrlEntered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_DirectoryUrlEntered
func miqt_exec_callback_QFileDialog_DirectoryUrlEntered(cb intptr_t, directory *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(directory *QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUrl(directory)

	gofunc(slotval1)
}

func (this *QFileDialog) FilterSelected(filter string) {
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	QFileDialog_FilterSelected(this.h, filter_ms)
}

func (this *QFileDialog) OnFilterSelected(slot func(filter string)) {
	QFileDialog_connect_FilterSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_FilterSelected
func miqt_exec_callback_QFileDialog_FilterSelected(cb intptr_t, filter struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(filter string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var filter_ms struct_miqt_string = filter
	filter_ret := GoStringN(filter_ms.data, int(int64(filter_ms.len)))
	free(unsafe.Pointer(filter_ms.data))
	slotval1 := filter_ret

	gofunc(slotval1)
}

func QFileDialog_GetOpenFileName() string {
	var _ms struct_miqt_string = QFileDialog_GetOpenFileName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetOpenFileUrl() *QUrl {
	_goptr := newQUrl(QFileDialog_GetOpenFileUrl())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetSaveFileName() string {
	var _ms struct_miqt_string = QFileDialog_GetSaveFileName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetSaveFileUrl() *QUrl {
	_goptr := newQUrl(QFileDialog_GetSaveFileUrl())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetExistingDirectory() string {
	var _ms struct_miqt_string = QFileDialog_GetExistingDirectory()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetExistingDirectoryUrl() *QUrl {
	_goptr := newQUrl(QFileDialog_GetExistingDirectoryUrl())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetOpenFileNames() []string {
	var _ma struct_miqt_array = QFileDialog_GetOpenFileNames()
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QFileDialog_GetOpenFileUrls() []QUrl {
	var _ma struct_miqt_array = QFileDialog_GetOpenFileUrls()
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QFileDialog_SaveFileContent(fileContent []byte, fileNameHint string) {
	fileContent_alias := struct_miqt_string{}
	fileContent_alias.data = (char)(unsafe.Pointer(&fileContent[0]))
	fileContent_alias.len = size_t(len(fileContent))
	fileNameHint_ms := struct_miqt_string{}
	fileNameHint_ms.data = CString(fileNameHint)
	fileNameHint_ms.len = size_t(len(fileNameHint))
	defer free(unsafe.Pointer(fileNameHint_ms.data))
	QFileDialog_SaveFileContent(fileContent_alias, fileNameHint_ms)
}

func QFileDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileDialog) SetOption2(option Option, on bool) {
	QFileDialog_SetOption2(this.h, option, (bool)(on))
}

func QFileDialog_GetOpenFileName1(parent *QWidget) string {
	var _ms struct_miqt_string = QFileDialog_GetOpenFileName1(parent.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetOpenFileName2(parent *QWidget, caption string) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetOpenFileName2(parent.cPointer(), caption_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetOpenFileName3(parent *QWidget, caption string, dir string) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetOpenFileName3(parent.cPointer(), caption_ms, dir_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetOpenFileName4(parent *QWidget, caption string, dir string, filter string) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetOpenFileName4(parent.cPointer(), caption_ms, dir_ms, filter_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetOpenFileUrl1(parent *QWidget) *QUrl {
	_goptr := newQUrl(QFileDialog_GetOpenFileUrl1(parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetOpenFileUrl2(parent *QWidget, caption string) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	_goptr := newQUrl(QFileDialog_GetOpenFileUrl2(parent.cPointer(), caption_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetOpenFileUrl3(parent *QWidget, caption string, dir *QUrl) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	_goptr := newQUrl(QFileDialog_GetOpenFileUrl3(parent.cPointer(), caption_ms, dir.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetOpenFileUrl4(parent *QWidget, caption string, dir *QUrl, filter string) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	_goptr := newQUrl(QFileDialog_GetOpenFileUrl4(parent.cPointer(), caption_ms, dir.cPointer(), filter_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetSaveFileName1(parent *QWidget) string {
	var _ms struct_miqt_string = QFileDialog_GetSaveFileName1(parent.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetSaveFileName2(parent *QWidget, caption string) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetSaveFileName2(parent.cPointer(), caption_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetSaveFileName3(parent *QWidget, caption string, dir string) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetSaveFileName3(parent.cPointer(), caption_ms, dir_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetSaveFileName4(parent *QWidget, caption string, dir string, filter string) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetSaveFileName4(parent.cPointer(), caption_ms, dir_ms, filter_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetSaveFileUrl1(parent *QWidget) *QUrl {
	_goptr := newQUrl(QFileDialog_GetSaveFileUrl1(parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetSaveFileUrl2(parent *QWidget, caption string) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	_goptr := newQUrl(QFileDialog_GetSaveFileUrl2(parent.cPointer(), caption_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetSaveFileUrl3(parent *QWidget, caption string, dir *QUrl) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	_goptr := newQUrl(QFileDialog_GetSaveFileUrl3(parent.cPointer(), caption_ms, dir.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetSaveFileUrl4(parent *QWidget, caption string, dir *QUrl, filter string) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	_goptr := newQUrl(QFileDialog_GetSaveFileUrl4(parent.cPointer(), caption_ms, dir.cPointer(), filter_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetExistingDirectory1(parent *QWidget) string {
	var _ms struct_miqt_string = QFileDialog_GetExistingDirectory1(parent.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetExistingDirectory2(parent *QWidget, caption string) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetExistingDirectory2(parent.cPointer(), caption_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetExistingDirectory3(parent *QWidget, caption string, dir string) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetExistingDirectory3(parent.cPointer(), caption_ms, dir_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetExistingDirectory4(parent *QWidget, caption string, dir string, options Options) string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	var _ms struct_miqt_string = QFileDialog_GetExistingDirectory4(parent.cPointer(), caption_ms, dir_ms, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileDialog_GetExistingDirectoryUrl1(parent *QWidget) *QUrl {
	_goptr := newQUrl(QFileDialog_GetExistingDirectoryUrl1(parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetExistingDirectoryUrl2(parent *QWidget, caption string) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	_goptr := newQUrl(QFileDialog_GetExistingDirectoryUrl2(parent.cPointer(), caption_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetExistingDirectoryUrl3(parent *QWidget, caption string, dir *QUrl) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	_goptr := newQUrl(QFileDialog_GetExistingDirectoryUrl3(parent.cPointer(), caption_ms, dir.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetExistingDirectoryUrl4(parent *QWidget, caption string, dir *QUrl, options Options) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	_goptr := newQUrl(QFileDialog_GetExistingDirectoryUrl4(parent.cPointer(), caption_ms, dir.cPointer(), options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetExistingDirectoryUrl5(parent *QWidget, caption string, dir *QUrl, options Options, supportedSchemes []string) *QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	supportedSchemes_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(supportedSchemes))))
	defer free(unsafe.Pointer(supportedSchemes_CArray))
	for i := range supportedSchemes {
		supportedSchemes_i_ms := struct_miqt_string{}
		supportedSchemes_i_ms.data = CString(supportedSchemes[i])
		supportedSchemes_i_ms.len = size_t(len(supportedSchemes[i]))
		defer free(unsafe.Pointer(supportedSchemes_i_ms.data))
		supportedSchemes_CArray[i] = supportedSchemes_i_ms
	}
	supportedSchemes_ma := struct_miqt_array{len: size_t(len(supportedSchemes)), data: unsafe.Pointer(supportedSchemes_CArray)}
	_goptr := newQUrl(QFileDialog_GetExistingDirectoryUrl5(parent.cPointer(), caption_ms, dir.cPointer(), options, supportedSchemes_ma))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFileDialog_GetOpenFileNames1(parent *QWidget) []string {
	var _ma struct_miqt_array = QFileDialog_GetOpenFileNames1(parent.cPointer())
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QFileDialog_GetOpenFileNames2(parent *QWidget, caption string) []string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	var _ma struct_miqt_array = QFileDialog_GetOpenFileNames2(parent.cPointer(), caption_ms)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QFileDialog_GetOpenFileNames3(parent *QWidget, caption string, dir string) []string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	var _ma struct_miqt_array = QFileDialog_GetOpenFileNames3(parent.cPointer(), caption_ms, dir_ms)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QFileDialog_GetOpenFileNames4(parent *QWidget, caption string, dir string, filter string) []string {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	dir_ms := struct_miqt_string{}
	dir_ms.data = CString(dir)
	dir_ms.len = size_t(len(dir))
	defer free(unsafe.Pointer(dir_ms.data))
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	var _ma struct_miqt_array = QFileDialog_GetOpenFileNames4(parent.cPointer(), caption_ms, dir_ms, filter_ms)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QFileDialog_GetOpenFileUrls1(parent *QWidget) []QUrl {
	var _ma struct_miqt_array = QFileDialog_GetOpenFileUrls1(parent.cPointer())
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QFileDialog_GetOpenFileUrls2(parent *QWidget, caption string) []QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	var _ma struct_miqt_array = QFileDialog_GetOpenFileUrls2(parent.cPointer(), caption_ms)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QFileDialog_GetOpenFileUrls3(parent *QWidget, caption string, dir *QUrl) []QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	var _ma struct_miqt_array = QFileDialog_GetOpenFileUrls3(parent.cPointer(), caption_ms, dir.cPointer())
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QFileDialog_GetOpenFileUrls4(parent *QWidget, caption string, dir *QUrl, filter string) []QUrl {
	caption_ms := struct_miqt_string{}
	caption_ms.data = CString(caption)
	caption_ms.len = size_t(len(caption))
	defer free(unsafe.Pointer(caption_ms.data))
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	var _ma struct_miqt_array = QFileDialog_GetOpenFileUrls4(parent.cPointer(), caption_ms, dir.cPointer(), filter_ms)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QFileDialog_SaveFileContent3(fileContent []byte, fileNameHint string, parent *QWidget) {
	fileContent_alias := struct_miqt_string{}
	fileContent_alias.data = (char)(unsafe.Pointer(&fileContent[0]))
	fileContent_alias.len = size_t(len(fileContent))
	fileNameHint_ms := struct_miqt_string{}
	fileNameHint_ms.data = CString(fileNameHint)
	fileNameHint_ms.len = size_t(len(fileNameHint))
	defer free(unsafe.Pointer(fileNameHint_ms.data))
	QFileDialog_SaveFileContent3(fileContent_alias, fileNameHint_ms, parent.cPointer())
}

func (this *QFileDialog) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFileDialog_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFileDialog) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileDialog_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_MetaObject
func miqt_exec_callback_QFileDialog_MetaObject(self QFileDialog, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileDialog{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFileDialog) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFileDialog_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFileDialog) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileDialog_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileDialog_Metacast
func miqt_exec_callback_QFileDialog_Metacast(self QFileDialog, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QFileDialog{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
