package qt

import (
	"unsafe"
)

type QUrl__ParsingMode int

const (
	QUrl__TolerantMode QUrl__ParsingMode = 0
	QUrl__StrictMode   QUrl__ParsingMode = 1
	QUrl__DecodedMode  QUrl__ParsingMode = 2
)

type QUrl__UrlFormattingOption uint

const (
	QUrl__None                  QUrl__UrlFormattingOption = 0
	QUrl__RemoveScheme          QUrl__UrlFormattingOption = 1
	QUrl__RemovePassword        QUrl__UrlFormattingOption = 2
	QUrl__RemoveUserInfo        QUrl__UrlFormattingOption = 6
	QUrl__RemovePort            QUrl__UrlFormattingOption = 8
	QUrl__RemoveAuthority       QUrl__UrlFormattingOption = 30
	QUrl__RemovePath            QUrl__UrlFormattingOption = 32
	QUrl__RemoveQuery           QUrl__UrlFormattingOption = 64
	QUrl__RemoveFragment        QUrl__UrlFormattingOption = 128
	QUrl__PreferLocalFile       QUrl__UrlFormattingOption = 512
	QUrl__StripTrailingSlash    QUrl__UrlFormattingOption = 1024
	QUrl__RemoveFilename        QUrl__UrlFormattingOption = 2048
	QUrl__NormalizePathSegments QUrl__UrlFormattingOption = 4096
)

type QUrl__ComponentFormattingOption uint

const (
	QUrl__PrettyDecoded    QUrl__ComponentFormattingOption = 0
	QUrl__EncodeSpaces     QUrl__ComponentFormattingOption = 1048576
	QUrl__EncodeUnicode    QUrl__ComponentFormattingOption = 2097152
	QUrl__EncodeDelimiters QUrl__ComponentFormattingOption = 12582912
	QUrl__EncodeReserved   QUrl__ComponentFormattingOption = 16777216
	QUrl__DecodeReserved   QUrl__ComponentFormattingOption = 33554432
	QUrl__FullyEncoded     QUrl__ComponentFormattingOption = 32505856
	QUrl__FullyDecoded     QUrl__ComponentFormattingOption = 133169152
)

type QUrl__UserInputResolutionOption int

const (
	QUrl__DefaultResolution QUrl__UserInputResolutionOption = 0
	QUrl__AssumeLocalFile   QUrl__UserInputResolutionOption = 1
)

type QUrl__AceProcessingOption uint

const (
	QUrl__IgnoreIDNWhitelist        QUrl__AceProcessingOption = 1
	QUrl__AceTransitionalProcessing QUrl__AceProcessingOption = 2
)

type QUrl struct {
	h          uintptr
	isSubclass bool
}

// NewQUrl constructs a new QUrl object.
func NewQUrl() *QUrl {
	g := newQUrl(QUrl_new())
	g.isSubclass = true
	return g
}

// NewQUrl2 constructs a new QUrl object.
func NewQUrl2(copyVal *QUrl) *QUrl {
	g := newQUrl(QUrl_new2(copyVal.cPointer()))
	g.isSubclass = true
	return g
}

// NewQUrl3 constructs a new QUrl object.
func NewQUrl3(url string) *QUrl {
	url_ms := struct_miqt_string{}
	url_ms.data = CString(url)
	url_ms.len = size_t(len(url))
	defer free(unsafe.Pointer(url_ms.data))
	g := newQUrl(QUrl_new3(url_ms))
	g.isSubclass = true
	return g
}

// NewQUrl4 constructs a new QUrl object.
func NewQUrl4(url string, mode ParsingMode) *QUrl {
	url_ms := struct_miqt_string{}
	url_ms.data = CString(url)
	url_ms.len = size_t(len(url))
	defer free(unsafe.Pointer(url_ms.data))
	g := newQUrl(QUrl_new4(url_ms, mode))
	g.isSubclass = true
	return g
}

func (this *QUrl) OperatorAssign(copyVal *QUrl) {
	QUrl_OperatorAssign(this.h, copyVal.cPointer())
}

func (this *QUrl) OperatorAssignWithUrl(url string) {
	url_ms := struct_miqt_string{}
	url_ms.data = CString(url)
	url_ms.len = size_t(len(url))
	defer free(unsafe.Pointer(url_ms.data))
	QUrl_OperatorAssignWithUrl(this.h, url_ms)
}

func (this *QUrl) Swap(other *QUrl) {
	QUrl_Swap(this.h, other.cPointer())
}

func (this *QUrl) SetUrl(url string) {
	url_ms := struct_miqt_string{}
	url_ms.data = CString(url)
	url_ms.len = size_t(len(url))
	defer free(unsafe.Pointer(url_ms.data))
	QUrl_SetUrl(this.h, url_ms)
}

func (this *QUrl) Url() string {
	var _ms struct_miqt_string = QUrl_Url(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) ToString() string {
	var _ms struct_miqt_string = QUrl_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) ToDisplayString() string {
	var _ms struct_miqt_string = QUrl_ToDisplayString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) Adjusted(options FormattingOptions) *QUrl {
	_goptr := newQUrl(QUrl_Adjusted(this.h, options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) ToEncoded() []byte {
	var _bytearray struct_miqt_string = QUrl_ToEncoded(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_FromEncoded(input QByteArrayView) *QUrl {
	_goptr := newQUrl(QUrl_FromEncoded(input.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QUrl_FromUserInput(userInput string) *QUrl {
	userInput_ms := struct_miqt_string{}
	userInput_ms.data = CString(userInput)
	userInput_ms.len = size_t(len(userInput))
	defer free(unsafe.Pointer(userInput_ms.data))
	_goptr := newQUrl(QUrl_FromUserInput(userInput_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) IsValid() bool {
	return (bool)(QUrl_IsValid(this.h))
}

func (this *QUrl) ErrorString() string {
	var _ms struct_miqt_string = QUrl_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) IsEmpty() bool {
	return (bool)(QUrl_IsEmpty(this.h))
}

func (this *QUrl) Clear() {
	QUrl_Clear(this.h)
}

func (this *QUrl) SetScheme(scheme string) {
	scheme_ms := struct_miqt_string{}
	scheme_ms.data = CString(scheme)
	scheme_ms.len = size_t(len(scheme))
	defer free(unsafe.Pointer(scheme_ms.data))
	QUrl_SetScheme(this.h, scheme_ms)
}

func (this *QUrl) Scheme() string {
	var _ms struct_miqt_string = QUrl_Scheme(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetAuthority(authority string) {
	authority_ms := struct_miqt_string{}
	authority_ms.data = CString(authority)
	authority_ms.len = size_t(len(authority))
	defer free(unsafe.Pointer(authority_ms.data))
	QUrl_SetAuthority(this.h, authority_ms)
}

func (this *QUrl) Authority() string {
	var _ms struct_miqt_string = QUrl_Authority(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetUserInfo(userInfo string) {
	userInfo_ms := struct_miqt_string{}
	userInfo_ms.data = CString(userInfo)
	userInfo_ms.len = size_t(len(userInfo))
	defer free(unsafe.Pointer(userInfo_ms.data))
	QUrl_SetUserInfo(this.h, userInfo_ms)
}

func (this *QUrl) UserInfo() string {
	var _ms struct_miqt_string = QUrl_UserInfo(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetUserName(userName string) {
	userName_ms := struct_miqt_string{}
	userName_ms.data = CString(userName)
	userName_ms.len = size_t(len(userName))
	defer free(unsafe.Pointer(userName_ms.data))
	QUrl_SetUserName(this.h, userName_ms)
}

func (this *QUrl) UserName() string {
	var _ms struct_miqt_string = QUrl_UserName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetPassword(password string) {
	password_ms := struct_miqt_string{}
	password_ms.data = CString(password)
	password_ms.len = size_t(len(password))
	defer free(unsafe.Pointer(password_ms.data))
	QUrl_SetPassword(this.h, password_ms)
}

func (this *QUrl) Password() string {
	var _ms struct_miqt_string = QUrl_Password(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetHost(host string) {
	host_ms := struct_miqt_string{}
	host_ms.data = CString(host)
	host_ms.len = size_t(len(host))
	defer free(unsafe.Pointer(host_ms.data))
	QUrl_SetHost(this.h, host_ms)
}

func (this *QUrl) Host() string {
	var _ms struct_miqt_string = QUrl_Host(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetPort(port int) {
	QUrl_SetPort(this.h, (int)(port))
}

func (this *QUrl) Port() int {
	return (int)(QUrl_Port(this.h))
}

func (this *QUrl) SetPath(path string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QUrl_SetPath(this.h, path_ms)
}

func (this *QUrl) Path() string {
	var _ms struct_miqt_string = QUrl_Path(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) FileName() string {
	var _ms struct_miqt_string = QUrl_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) HasQuery() bool {
	return (bool)(QUrl_HasQuery(this.h))
}

func (this *QUrl) SetQuery(query string) {
	query_ms := struct_miqt_string{}
	query_ms.data = CString(query)
	query_ms.len = size_t(len(query))
	defer free(unsafe.Pointer(query_ms.data))
	QUrl_SetQuery(this.h, query_ms)
}

func (this *QUrl) SetQueryWithQuery(query *QUrlQuery) {
	QUrl_SetQueryWithQuery(this.h, query.cPointer())
}

func (this *QUrl) Query() string {
	var _ms struct_miqt_string = QUrl_Query(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) HasFragment() bool {
	return (bool)(QUrl_HasFragment(this.h))
}

func (this *QUrl) Fragment() string {
	var _ms struct_miqt_string = QUrl_Fragment(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetFragment(fragment string) {
	fragment_ms := struct_miqt_string{}
	fragment_ms.data = CString(fragment)
	fragment_ms.len = size_t(len(fragment))
	defer free(unsafe.Pointer(fragment_ms.data))
	QUrl_SetFragment(this.h, fragment_ms)
}

func (this *QUrl) Resolved(relative *QUrl) *QUrl {
	_goptr := newQUrl(QUrl_Resolved(this.h, relative.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) IsRelative() bool {
	return (bool)(QUrl_IsRelative(this.h))
}

func (this *QUrl) IsParentOf(url *QUrl) bool {
	return (bool)(QUrl_IsParentOf(this.h, url.cPointer()))
}

func (this *QUrl) IsLocalFile() bool {
	return (bool)(QUrl_IsLocalFile(this.h))
}

func QUrl_FromLocalFile(localfile string) *QUrl {
	localfile_ms := struct_miqt_string{}
	localfile_ms.data = CString(localfile)
	localfile_ms.len = size_t(len(localfile))
	defer free(unsafe.Pointer(localfile_ms.data))
	_goptr := newQUrl(QUrl_FromLocalFile(localfile_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) ToLocalFile() string {
	var _ms struct_miqt_string = QUrl_ToLocalFile(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) Detach() {
	QUrl_Detach(this.h)
}

func (this *QUrl) IsDetached() bool {
	return (bool)(QUrl_IsDetached(this.h))
}

func (this *QUrl) Matches(url *QUrl, options FormattingOptions) bool {
	return (bool)(QUrl_Matches(this.h, url.cPointer(), options))
}

func QUrl_FromPercentEncoding(param1 []byte) string {
	param1_alias := struct_miqt_string{}
	param1_alias.data = (char)(unsafe.Pointer(&param1[0]))
	param1_alias.len = size_t(len(param1))
	var _ms struct_miqt_string = QUrl_FromPercentEncoding(param1_alias)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUrl_ToPercentEncoding(param1 string) []byte {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	var _bytearray struct_miqt_string = QUrl_ToPercentEncoding(param1_ms)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_FromAce(domain []byte) string {
	domain_alias := struct_miqt_string{}
	domain_alias.data = (char)(unsafe.Pointer(&domain[0]))
	domain_alias.len = size_t(len(domain))
	var _ms struct_miqt_string = QUrl_FromAce(domain_alias)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUrl_ToAce(domain string) []byte {
	domain_ms := struct_miqt_string{}
	domain_ms.data = CString(domain)
	domain_ms.len = size_t(len(domain))
	defer free(unsafe.Pointer(domain_ms.data))
	var _bytearray struct_miqt_string = QUrl_ToAce(domain_ms)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_IdnWhitelist() []string {
	var _ma struct_miqt_array = QUrl_IdnWhitelist()
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

func QUrl_ToStringList(uris []QUrl) []string {
	uris_CArray := (*[0xffff]*QUrl)(malloc(size_t(8 * len(uris))))
	defer free(unsafe.Pointer(uris_CArray))
	for i := range uris {
		uris_CArray[i] = uris[i].cPointer()
	}
	uris_ma := struct_miqt_array{len: size_t(len(uris)), data: unsafe.Pointer(uris_CArray)}
	var _ma struct_miqt_array = QUrl_ToStringList(uris_ma)
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

func QUrl_FromStringList(uris []string) []QUrl {
	uris_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(uris))))
	defer free(unsafe.Pointer(uris_CArray))
	for i := range uris {
		uris_i_ms := struct_miqt_string{}
		uris_i_ms.data = CString(uris[i])
		uris_i_ms.len = size_t(len(uris[i]))
		defer free(unsafe.Pointer(uris_i_ms.data))
		uris_CArray[i] = uris_i_ms
	}
	uris_ma := struct_miqt_array{len: size_t(len(uris)), data: unsafe.Pointer(uris_CArray)}
	var _ma struct_miqt_array = QUrl_FromStringList(uris_ma)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QUrl_SetIdnWhitelist(idnWhitelist []string) {
	idnWhitelist_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(idnWhitelist))))
	defer free(unsafe.Pointer(idnWhitelist_CArray))
	for i := range idnWhitelist {
		idnWhitelist_i_ms := struct_miqt_string{}
		idnWhitelist_i_ms.data = CString(idnWhitelist[i])
		idnWhitelist_i_ms.len = size_t(len(idnWhitelist[i]))
		defer free(unsafe.Pointer(idnWhitelist_i_ms.data))
		idnWhitelist_CArray[i] = idnWhitelist_i_ms
	}
	idnWhitelist_ma := struct_miqt_array{len: size_t(len(idnWhitelist)), data: unsafe.Pointer(idnWhitelist_CArray)}
	QUrl_SetIdnWhitelist(idnWhitelist_ma)
}

func (this *QUrl) DataPtr() *DataPtr {
	xxxxxxxxx
}

func (this *QUrl) SetUrl2(url string, mode ParsingMode) {
	url_ms := struct_miqt_string{}
	url_ms.data = CString(url)
	url_ms.len = size_t(len(url))
	defer free(unsafe.Pointer(url_ms.data))
	QUrl_SetUrl2(this.h, url_ms, mode)
}

func (this *QUrl) Url1(options FormattingOptions) string {
	var _ms struct_miqt_string = QUrl_Url1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) ToString1(options FormattingOptions) string {
	var _ms struct_miqt_string = QUrl_ToString1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) ToDisplayString1(options FormattingOptions) string {
	var _ms struct_miqt_string = QUrl_ToDisplayString1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) ToEncoded1(options FormattingOptions) []byte {
	var _bytearray struct_miqt_string = QUrl_ToEncoded1(this.h, options)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_FromEncoded2(input QByteArrayView, mode ParsingMode) *QUrl {
	_goptr := newQUrl(QUrl_FromEncoded2(input.cPointer(), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QUrl_FromUserInput2(userInput string, workingDirectory string) *QUrl {
	userInput_ms := struct_miqt_string{}
	userInput_ms.data = CString(userInput)
	userInput_ms.len = size_t(len(userInput))
	defer free(unsafe.Pointer(userInput_ms.data))
	workingDirectory_ms := struct_miqt_string{}
	workingDirectory_ms.data = CString(workingDirectory)
	workingDirectory_ms.len = size_t(len(workingDirectory))
	defer free(unsafe.Pointer(workingDirectory_ms.data))
	_goptr := newQUrl(QUrl_FromUserInput2(userInput_ms, workingDirectory_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QUrl_FromUserInput3(userInput string, workingDirectory string, options UserInputResolutionOptions) *QUrl {
	userInput_ms := struct_miqt_string{}
	userInput_ms.data = CString(userInput)
	userInput_ms.len = size_t(len(userInput))
	defer free(unsafe.Pointer(userInput_ms.data))
	workingDirectory_ms := struct_miqt_string{}
	workingDirectory_ms.data = CString(workingDirectory)
	workingDirectory_ms.len = size_t(len(workingDirectory))
	defer free(unsafe.Pointer(workingDirectory_ms.data))
	_goptr := newQUrl(QUrl_FromUserInput3(userInput_ms, workingDirectory_ms, options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) SetAuthority2(authority string, mode ParsingMode) {
	authority_ms := struct_miqt_string{}
	authority_ms.data = CString(authority)
	authority_ms.len = size_t(len(authority))
	defer free(unsafe.Pointer(authority_ms.data))
	QUrl_SetAuthority2(this.h, authority_ms, mode)
}

func (this *QUrl) Authority1(options ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_Authority1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetUserInfo2(userInfo string, mode ParsingMode) {
	userInfo_ms := struct_miqt_string{}
	userInfo_ms.data = CString(userInfo)
	userInfo_ms.len = size_t(len(userInfo))
	defer free(unsafe.Pointer(userInfo_ms.data))
	QUrl_SetUserInfo2(this.h, userInfo_ms, mode)
}

func (this *QUrl) UserInfo1(options ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_UserInfo1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetUserName2(userName string, mode ParsingMode) {
	userName_ms := struct_miqt_string{}
	userName_ms.data = CString(userName)
	userName_ms.len = size_t(len(userName))
	defer free(unsafe.Pointer(userName_ms.data))
	QUrl_SetUserName2(this.h, userName_ms, mode)
}

func (this *QUrl) UserName1(options ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_UserName1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetPassword2(password string, mode ParsingMode) {
	password_ms := struct_miqt_string{}
	password_ms.data = CString(password)
	password_ms.len = size_t(len(password))
	defer free(unsafe.Pointer(password_ms.data))
	QUrl_SetPassword2(this.h, password_ms, mode)
}

func (this *QUrl) Password1(param1 ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_Password1(this.h, param1)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetHost2(host string, mode ParsingMode) {
	host_ms := struct_miqt_string{}
	host_ms.data = CString(host)
	host_ms.len = size_t(len(host))
	defer free(unsafe.Pointer(host_ms.data))
	QUrl_SetHost2(this.h, host_ms, mode)
}

func (this *QUrl) Host1(param1 ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_Host1(this.h, param1)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) Port1(defaultPort int) int {
	return (int)(QUrl_Port1(this.h, (int)(defaultPort)))
}

func (this *QUrl) SetPath2(path string, mode ParsingMode) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QUrl_SetPath2(this.h, path_ms, mode)
}

func (this *QUrl) Path1(options ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_Path1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) FileName1(options ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_FileName1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetQuery2(query string, mode ParsingMode) {
	query_ms := struct_miqt_string{}
	query_ms.data = CString(query)
	query_ms.len = size_t(len(query))
	defer free(unsafe.Pointer(query_ms.data))
	QUrl_SetQuery2(this.h, query_ms, mode)
}

func (this *QUrl) Query1(param1 ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_Query1(this.h, param1)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) Fragment1(options ComponentFormattingOptions) string {
	var _ms struct_miqt_string = QUrl_Fragment1(this.h, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetFragment2(fragment string, mode ParsingMode) {
	fragment_ms := struct_miqt_string{}
	fragment_ms.data = CString(fragment)
	fragment_ms.len = size_t(len(fragment))
	defer free(unsafe.Pointer(fragment_ms.data))
	QUrl_SetFragment2(this.h, fragment_ms, mode)
}

func QUrl_ToPercentEncoding2(param1 string, exclude []byte) []byte {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	exclude_alias := struct_miqt_string{}
	exclude_alias.data = (char)(unsafe.Pointer(&exclude[0]))
	exclude_alias.len = size_t(len(exclude))
	var _bytearray struct_miqt_string = QUrl_ToPercentEncoding2(param1_ms, exclude_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_ToPercentEncoding3(param1 string, exclude []byte, include []byte) []byte {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	exclude_alias := struct_miqt_string{}
	exclude_alias.data = (char)(unsafe.Pointer(&exclude[0]))
	exclude_alias.len = size_t(len(exclude))
	include_alias := struct_miqt_string{}
	include_alias.data = (char)(unsafe.Pointer(&include[0]))
	include_alias.len = size_t(len(include))
	var _bytearray struct_miqt_string = QUrl_ToPercentEncoding3(param1_ms, exclude_alias, include_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_FromAce2(domain []byte, options AceProcessingOptions) string {
	domain_alias := struct_miqt_string{}
	domain_alias.data = (char)(unsafe.Pointer(&domain[0]))
	domain_alias.len = size_t(len(domain))
	var _ms struct_miqt_string = QUrl_FromAce2(domain_alias, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUrl_ToAce2(domain string, options AceProcessingOptions) []byte {
	domain_ms := struct_miqt_string{}
	domain_ms.data = CString(domain)
	domain_ms.len = size_t(len(domain))
	defer free(unsafe.Pointer(domain_ms.data))
	var _bytearray struct_miqt_string = QUrl_ToAce2(domain_ms, options)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_ToStringList2(uris []QUrl, options FormattingOptions) []string {
	uris_CArray := (*[0xffff]*QUrl)(malloc(size_t(8 * len(uris))))
	defer free(unsafe.Pointer(uris_CArray))
	for i := range uris {
		uris_CArray[i] = uris[i].cPointer()
	}
	uris_ma := struct_miqt_array{len: size_t(len(uris)), data: unsafe.Pointer(uris_CArray)}
	var _ma struct_miqt_array = QUrl_ToStringList2(uris_ma, options)
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

func QUrl_FromStringList2(uris []string, mode ParsingMode) []QUrl {
	uris_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(uris))))
	defer free(unsafe.Pointer(uris_CArray))
	for i := range uris {
		uris_i_ms := struct_miqt_string{}
		uris_i_ms.data = CString(uris[i])
		uris_i_ms.len = size_t(len(uris[i]))
		defer free(unsafe.Pointer(uris_i_ms.data))
		uris_CArray[i] = uris_i_ms
	}
	uris_ma := struct_miqt_array{len: size_t(len(uris)), data: unsafe.Pointer(uris_CArray)}
	var _ma struct_miqt_array = QUrl_FromStringList2(uris_ma, mode)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}
