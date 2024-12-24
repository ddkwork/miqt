package qt

import (
	"unsafe"
)

type QXmlStreamReader__TokenType int

const (
	QXmlStreamReader__NoToken               QXmlStreamReader__TokenType = 0
	QXmlStreamReader__Invalid               QXmlStreamReader__TokenType = 1
	QXmlStreamReader__StartDocument         QXmlStreamReader__TokenType = 2
	QXmlStreamReader__EndDocument           QXmlStreamReader__TokenType = 3
	QXmlStreamReader__StartElement          QXmlStreamReader__TokenType = 4
	QXmlStreamReader__EndElement            QXmlStreamReader__TokenType = 5
	QXmlStreamReader__Characters            QXmlStreamReader__TokenType = 6
	QXmlStreamReader__Comment               QXmlStreamReader__TokenType = 7
	QXmlStreamReader__DTD                   QXmlStreamReader__TokenType = 8
	QXmlStreamReader__EntityReference       QXmlStreamReader__TokenType = 9
	QXmlStreamReader__ProcessingInstruction QXmlStreamReader__TokenType = 10
)

type QXmlStreamReader__ReadElementTextBehaviour int

const (
	QXmlStreamReader__ErrorOnUnexpectedElement QXmlStreamReader__ReadElementTextBehaviour = 0
	QXmlStreamReader__IncludeChildElements     QXmlStreamReader__ReadElementTextBehaviour = 1
	QXmlStreamReader__SkipChildElements        QXmlStreamReader__ReadElementTextBehaviour = 2
)

type QXmlStreamReader__Error int

const (
	QXmlStreamReader__NoError                     QXmlStreamReader__Error = 0
	QXmlStreamReader__UnexpectedElementError      QXmlStreamReader__Error = 1
	QXmlStreamReader__CustomError                 QXmlStreamReader__Error = 2
	QXmlStreamReader__NotWellFormedError          QXmlStreamReader__Error = 3
	QXmlStreamReader__PrematureEndOfDocumentError QXmlStreamReader__Error = 4
)

type QXmlStreamAttribute struct {
	h          uintptr
	isSubclass bool
}

// NewQXmlStreamAttribute constructs a new QXmlStreamAttribute object.
func NewQXmlStreamAttribute() *QXmlStreamAttribute {
	ret := newQXmlStreamAttribute(QXmlStreamAttribute_new())
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamAttribute2 constructs a new QXmlStreamAttribute object.
func NewQXmlStreamAttribute2(qualifiedName string, value string) *QXmlStreamAttribute {
	qualifiedName_ms := struct_miqt_string{}
	qualifiedName_ms.data = CString(qualifiedName)
	qualifiedName_ms.len = size_t(len(qualifiedName))
	defer free(unsafe.Pointer(qualifiedName_ms.data))
	value_ms := struct_miqt_string{}
	value_ms.data = CString(value)
	value_ms.len = size_t(len(value))
	defer free(unsafe.Pointer(value_ms.data))

	ret := newQXmlStreamAttribute(QXmlStreamAttribute_new2(qualifiedName_ms, value_ms))
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamAttribute3 constructs a new QXmlStreamAttribute object.
func NewQXmlStreamAttribute3(namespaceUri string, name string, value string) *QXmlStreamAttribute {
	namespaceUri_ms := struct_miqt_string{}
	namespaceUri_ms.data = CString(namespaceUri)
	namespaceUri_ms.len = size_t(len(namespaceUri))
	defer free(unsafe.Pointer(namespaceUri_ms.data))
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	value_ms := struct_miqt_string{}
	value_ms.data = CString(value)
	value_ms.len = size_t(len(value))
	defer free(unsafe.Pointer(value_ms.data))

	ret := newQXmlStreamAttribute(QXmlStreamAttribute_new3(namespaceUri_ms, name_ms, value_ms))
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamAttribute4 constructs a new QXmlStreamAttribute object.
func NewQXmlStreamAttribute4(param1 *QXmlStreamAttribute) *QXmlStreamAttribute {
	ret := newQXmlStreamAttribute(QXmlStreamAttribute_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QXmlStreamAttribute) IsDefault() bool {
	return (bool)(QXmlStreamAttribute_IsDefault(this.h))
}

type QXmlStreamNamespaceDeclaration struct {
	h          uintptr
	isSubclass bool
}

// NewQXmlStreamNamespaceDeclaration constructs a new QXmlStreamNamespaceDeclaration object.
func NewQXmlStreamNamespaceDeclaration() *QXmlStreamNamespaceDeclaration {
	ret := newQXmlStreamNamespaceDeclaration(QXmlStreamNamespaceDeclaration_new())
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamNamespaceDeclaration2 constructs a new QXmlStreamNamespaceDeclaration object.
func NewQXmlStreamNamespaceDeclaration2(prefix string, namespaceUri string) *QXmlStreamNamespaceDeclaration {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	namespaceUri_ms := struct_miqt_string{}
	namespaceUri_ms.data = CString(namespaceUri)
	namespaceUri_ms.len = size_t(len(namespaceUri))
	defer free(unsafe.Pointer(namespaceUri_ms.data))

	ret := newQXmlStreamNamespaceDeclaration(QXmlStreamNamespaceDeclaration_new2(prefix_ms, namespaceUri_ms))
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamNamespaceDeclaration3 constructs a new QXmlStreamNamespaceDeclaration object.
func NewQXmlStreamNamespaceDeclaration3(param1 *QXmlStreamNamespaceDeclaration) *QXmlStreamNamespaceDeclaration {
	ret := newQXmlStreamNamespaceDeclaration(QXmlStreamNamespaceDeclaration_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

type QXmlStreamNotationDeclaration struct {
	h          uintptr
	isSubclass bool
}

// NewQXmlStreamNotationDeclaration constructs a new QXmlStreamNotationDeclaration object.
func NewQXmlStreamNotationDeclaration() *QXmlStreamNotationDeclaration {
	ret := newQXmlStreamNotationDeclaration(QXmlStreamNotationDeclaration_new())
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamNotationDeclaration2 constructs a new QXmlStreamNotationDeclaration object.
func NewQXmlStreamNotationDeclaration2(param1 *QXmlStreamNotationDeclaration) *QXmlStreamNotationDeclaration {
	ret := newQXmlStreamNotationDeclaration(QXmlStreamNotationDeclaration_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

type QXmlStreamEntityDeclaration struct {
	h          uintptr
	isSubclass bool
}

// NewQXmlStreamEntityDeclaration constructs a new QXmlStreamEntityDeclaration object.
func NewQXmlStreamEntityDeclaration() *QXmlStreamEntityDeclaration {
	ret := newQXmlStreamEntityDeclaration(QXmlStreamEntityDeclaration_new())
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamEntityDeclaration2 constructs a new QXmlStreamEntityDeclaration object.
func NewQXmlStreamEntityDeclaration2(param1 *QXmlStreamEntityDeclaration) *QXmlStreamEntityDeclaration {
	ret := newQXmlStreamEntityDeclaration(QXmlStreamEntityDeclaration_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

type QXmlStreamEntityResolver struct {
	h          uintptr
	isSubclass bool
}

// NewQXmlStreamEntityResolver constructs a new QXmlStreamEntityResolver object.
func NewQXmlStreamEntityResolver() *QXmlStreamEntityResolver {
	ret := newQXmlStreamEntityResolver(QXmlStreamEntityResolver_new())
	ret.isSubclass = true
	return ret
}

func (this *QXmlStreamEntityResolver) ResolveEntity(publicId string, systemId string) string {
	publicId_ms := struct_miqt_string{}
	publicId_ms.data = CString(publicId)
	publicId_ms.len = size_t(len(publicId))
	defer free(unsafe.Pointer(publicId_ms.data))
	systemId_ms := struct_miqt_string{}
	systemId_ms.data = CString(systemId)
	systemId_ms.len = size_t(len(systemId))
	defer free(unsafe.Pointer(systemId_ms.data))
	var _ms struct_miqt_string = QXmlStreamEntityResolver_ResolveEntity(this.h, publicId_ms, systemId_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QXmlStreamEntityResolver) ResolveUndeclaredEntity(name string) string {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	var _ms struct_miqt_string = QXmlStreamEntityResolver_ResolveUndeclaredEntity(this.h, name_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QXmlStreamReader struct {
	h          uintptr
	isSubclass bool
}

// NewQXmlStreamReader constructs a new QXmlStreamReader object.
func NewQXmlStreamReader() *QXmlStreamReader {
	ret := newQXmlStreamReader(QXmlStreamReader_new())
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamReader2 constructs a new QXmlStreamReader object.
func NewQXmlStreamReader2(device *QIODevice) *QXmlStreamReader {
	ret := newQXmlStreamReader(QXmlStreamReader_new2(device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamReader3 constructs a new QXmlStreamReader object.
func NewQXmlStreamReader3(data QAnyStringView) *QXmlStreamReader {
	ret := newQXmlStreamReader(QXmlStreamReader_new3(data.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QXmlStreamReader) SetDevice(device *QIODevice) {
	QXmlStreamReader_SetDevice(this.h, device.cPointer())
}

func (this *QXmlStreamReader) Device() *QIODevice {
	return newQIODevice(QXmlStreamReader_Device(this.h))
}

func (this *QXmlStreamReader) AddData(data QAnyStringView) {
	QXmlStreamReader_AddData(this.h, data.cPointer())
}

func (this *QXmlStreamReader) Clear() {
	QXmlStreamReader_Clear(this.h)
}

func (this *QXmlStreamReader) AtEnd() bool {
	return (bool)(QXmlStreamReader_AtEnd(this.h))
}

func (this *QXmlStreamReader) ReadNext() TokenType {
	xxxxxxxxx
}

func (this *QXmlStreamReader) ReadNextStartElement() bool {
	return (bool)(QXmlStreamReader_ReadNextStartElement(this.h))
}

func (this *QXmlStreamReader) SkipCurrentElement() {
	QXmlStreamReader_SkipCurrentElement(this.h)
}

func (this *QXmlStreamReader) TokenType() TokenType {
	xxxxxxxxx
}

func (this *QXmlStreamReader) TokenString() string {
	var _ms struct_miqt_string = QXmlStreamReader_TokenString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QXmlStreamReader) SetNamespaceProcessing(namespaceProcessing bool) {
	QXmlStreamReader_SetNamespaceProcessing(this.h, (bool)(namespaceProcessing))
}

func (this *QXmlStreamReader) NamespaceProcessing() bool {
	return (bool)(QXmlStreamReader_NamespaceProcessing(this.h))
}

func (this *QXmlStreamReader) IsStartDocument() bool {
	return (bool)(QXmlStreamReader_IsStartDocument(this.h))
}

func (this *QXmlStreamReader) IsEndDocument() bool {
	return (bool)(QXmlStreamReader_IsEndDocument(this.h))
}

func (this *QXmlStreamReader) IsStartElement() bool {
	return (bool)(QXmlStreamReader_IsStartElement(this.h))
}

func (this *QXmlStreamReader) IsEndElement() bool {
	return (bool)(QXmlStreamReader_IsEndElement(this.h))
}

func (this *QXmlStreamReader) IsCharacters() bool {
	return (bool)(QXmlStreamReader_IsCharacters(this.h))
}

func (this *QXmlStreamReader) IsWhitespace() bool {
	return (bool)(QXmlStreamReader_IsWhitespace(this.h))
}

func (this *QXmlStreamReader) IsCDATA() bool {
	return (bool)(QXmlStreamReader_IsCDATA(this.h))
}

func (this *QXmlStreamReader) IsComment() bool {
	return (bool)(QXmlStreamReader_IsComment(this.h))
}

func (this *QXmlStreamReader) IsDTD() bool {
	return (bool)(QXmlStreamReader_IsDTD(this.h))
}

func (this *QXmlStreamReader) IsEntityReference() bool {
	return (bool)(QXmlStreamReader_IsEntityReference(this.h))
}

func (this *QXmlStreamReader) IsProcessingInstruction() bool {
	return (bool)(QXmlStreamReader_IsProcessingInstruction(this.h))
}

func (this *QXmlStreamReader) IsStandaloneDocument() bool {
	return (bool)(QXmlStreamReader_IsStandaloneDocument(this.h))
}

func (this *QXmlStreamReader) HasStandaloneDeclaration() bool {
	return (bool)(QXmlStreamReader_HasStandaloneDeclaration(this.h))
}

func (this *QXmlStreamReader) LineNumber() int64 {
	return (int64)(QXmlStreamReader_LineNumber(this.h))
}

func (this *QXmlStreamReader) ColumnNumber() int64 {
	return (int64)(QXmlStreamReader_ColumnNumber(this.h))
}

func (this *QXmlStreamReader) CharacterOffset() int64 {
	return (int64)(QXmlStreamReader_CharacterOffset(this.h))
}

func (this *QXmlStreamReader) ReadElementText() string {
	var _ms struct_miqt_string = QXmlStreamReader_ReadElementText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QXmlStreamReader) NamespaceDeclarations() []QXmlStreamNamespaceDeclaration {
	var _ma struct_miqt_array = QXmlStreamReader_NamespaceDeclarations(this.h)
	_ret := make([]QXmlStreamNamespaceDeclaration, int(_ma.len))
	_outCast := (*[0xffff]*QXmlStreamNamespaceDeclaration)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQXmlStreamNamespaceDeclaration(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QXmlStreamReader) AddExtraNamespaceDeclaration(extraNamespaceDeclaraction *QXmlStreamNamespaceDeclaration) {
	QXmlStreamReader_AddExtraNamespaceDeclaration(this.h, extraNamespaceDeclaraction.cPointer())
}

func (this *QXmlStreamReader) AddExtraNamespaceDeclarations(extraNamespaceDeclaractions []QXmlStreamNamespaceDeclaration) {
	extraNamespaceDeclaractions_CArray := (*[0xffff]*QXmlStreamNamespaceDeclaration)(malloc(size_t(8 * len(extraNamespaceDeclaractions))))
	defer free(unsafe.Pointer(extraNamespaceDeclaractions_CArray))
	for i := range extraNamespaceDeclaractions {
		extraNamespaceDeclaractions_CArray[i] = extraNamespaceDeclaractions[i].cPointer()
	}
	extraNamespaceDeclaractions_ma := struct_miqt_array{len: size_t(len(extraNamespaceDeclaractions)), data: unsafe.Pointer(extraNamespaceDeclaractions_CArray)}
	QXmlStreamReader_AddExtraNamespaceDeclarations(this.h, extraNamespaceDeclaractions_ma)
}

func (this *QXmlStreamReader) NotationDeclarations() []QXmlStreamNotationDeclaration {
	var _ma struct_miqt_array = QXmlStreamReader_NotationDeclarations(this.h)
	_ret := make([]QXmlStreamNotationDeclaration, int(_ma.len))
	_outCast := (*[0xffff]*QXmlStreamNotationDeclaration)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQXmlStreamNotationDeclaration(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QXmlStreamReader) EntityDeclarations() []QXmlStreamEntityDeclaration {
	var _ma struct_miqt_array = QXmlStreamReader_EntityDeclarations(this.h)
	_ret := make([]QXmlStreamEntityDeclaration, int(_ma.len))
	_outCast := (*[0xffff]*QXmlStreamEntityDeclaration)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQXmlStreamEntityDeclaration(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QXmlStreamReader) EntityExpansionLimit() int {
	return (int)(QXmlStreamReader_EntityExpansionLimit(this.h))
}

func (this *QXmlStreamReader) SetEntityExpansionLimit(limit int) {
	QXmlStreamReader_SetEntityExpansionLimit(this.h, (int)(limit))
}

func (this *QXmlStreamReader) RaiseError() {
	QXmlStreamReader_RaiseError(this.h)
}

func (this *QXmlStreamReader) ErrorString() string {
	var _ms struct_miqt_string = QXmlStreamReader_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QXmlStreamReader) Error() Error {
	xxxxxxxxx
}

func (this *QXmlStreamReader) HasError() bool {
	return (bool)(QXmlStreamReader_HasError(this.h))
}

func (this *QXmlStreamReader) SetEntityResolver(resolver *QXmlStreamEntityResolver) {
	QXmlStreamReader_SetEntityResolver(this.h, resolver.cPointer())
}

func (this *QXmlStreamReader) EntityResolver() *QXmlStreamEntityResolver {
	return newQXmlStreamEntityResolver(QXmlStreamReader_EntityResolver(this.h))
}

func (this *QXmlStreamReader) ReadElementText1(behaviour ReadElementTextBehaviour) string {
	var _ms struct_miqt_string = QXmlStreamReader_ReadElementText1(this.h, behaviour)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QXmlStreamReader) RaiseError1(message string) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QXmlStreamReader_RaiseError1(this.h, message_ms)
}

type QXmlStreamWriter struct {
	h          uintptr
	isSubclass bool
}

// NewQXmlStreamWriter constructs a new QXmlStreamWriter object.
func NewQXmlStreamWriter() *QXmlStreamWriter {
	ret := newQXmlStreamWriter(QXmlStreamWriter_new())
	ret.isSubclass = true
	return ret
}

// NewQXmlStreamWriter2 constructs a new QXmlStreamWriter object.
func NewQXmlStreamWriter2(device *QIODevice) *QXmlStreamWriter {
	ret := newQXmlStreamWriter(QXmlStreamWriter_new2(device.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QXmlStreamWriter) SetDevice(device *QIODevice) {
	QXmlStreamWriter_SetDevice(this.h, device.cPointer())
}

func (this *QXmlStreamWriter) Device() *QIODevice {
	return newQIODevice(QXmlStreamWriter_Device(this.h))
}

func (this *QXmlStreamWriter) SetAutoFormatting(autoFormatting bool) {
	QXmlStreamWriter_SetAutoFormatting(this.h, (bool)(autoFormatting))
}

func (this *QXmlStreamWriter) AutoFormatting() bool {
	return (bool)(QXmlStreamWriter_AutoFormatting(this.h))
}

func (this *QXmlStreamWriter) SetAutoFormattingIndent(spacesOrTabs int) {
	QXmlStreamWriter_SetAutoFormattingIndent(this.h, (int)(spacesOrTabs))
}

func (this *QXmlStreamWriter) AutoFormattingIndent() int {
	return (int)(QXmlStreamWriter_AutoFormattingIndent(this.h))
}

func (this *QXmlStreamWriter) WriteAttribute(qualifiedName QAnyStringView, value QAnyStringView) {
	QXmlStreamWriter_WriteAttribute(this.h, qualifiedName.cPointer(), value.cPointer())
}

func (this *QXmlStreamWriter) WriteAttribute2(namespaceUri QAnyStringView, name QAnyStringView, value QAnyStringView) {
	QXmlStreamWriter_WriteAttribute2(this.h, namespaceUri.cPointer(), name.cPointer(), value.cPointer())
}

func (this *QXmlStreamWriter) WriteAttributeWithAttribute(attribute *QXmlStreamAttribute) {
	QXmlStreamWriter_WriteAttributeWithAttribute(this.h, attribute.cPointer())
}

func (this *QXmlStreamWriter) WriteCDATA(text QAnyStringView) {
	QXmlStreamWriter_WriteCDATA(this.h, text.cPointer())
}

func (this *QXmlStreamWriter) WriteCharacters(text QAnyStringView) {
	QXmlStreamWriter_WriteCharacters(this.h, text.cPointer())
}

func (this *QXmlStreamWriter) WriteComment(text QAnyStringView) {
	QXmlStreamWriter_WriteComment(this.h, text.cPointer())
}

func (this *QXmlStreamWriter) WriteDTD(dtd QAnyStringView) {
	QXmlStreamWriter_WriteDTD(this.h, dtd.cPointer())
}

func (this *QXmlStreamWriter) WriteEmptyElement(qualifiedName QAnyStringView) {
	QXmlStreamWriter_WriteEmptyElement(this.h, qualifiedName.cPointer())
}

func (this *QXmlStreamWriter) WriteEmptyElement2(namespaceUri QAnyStringView, name QAnyStringView) {
	QXmlStreamWriter_WriteEmptyElement2(this.h, namespaceUri.cPointer(), name.cPointer())
}

func (this *QXmlStreamWriter) WriteTextElement(qualifiedName QAnyStringView, text QAnyStringView) {
	QXmlStreamWriter_WriteTextElement(this.h, qualifiedName.cPointer(), text.cPointer())
}

func (this *QXmlStreamWriter) WriteTextElement2(namespaceUri QAnyStringView, name QAnyStringView, text QAnyStringView) {
	QXmlStreamWriter_WriteTextElement2(this.h, namespaceUri.cPointer(), name.cPointer(), text.cPointer())
}

func (this *QXmlStreamWriter) WriteEndDocument() {
	QXmlStreamWriter_WriteEndDocument(this.h)
}

func (this *QXmlStreamWriter) WriteEndElement() {
	QXmlStreamWriter_WriteEndElement(this.h)
}

func (this *QXmlStreamWriter) WriteEntityReference(name QAnyStringView) {
	QXmlStreamWriter_WriteEntityReference(this.h, name.cPointer())
}

func (this *QXmlStreamWriter) WriteNamespace(namespaceUri QAnyStringView) {
	QXmlStreamWriter_WriteNamespace(this.h, namespaceUri.cPointer())
}

func (this *QXmlStreamWriter) WriteDefaultNamespace(namespaceUri QAnyStringView) {
	QXmlStreamWriter_WriteDefaultNamespace(this.h, namespaceUri.cPointer())
}

func (this *QXmlStreamWriter) WriteProcessingInstruction(target QAnyStringView) {
	QXmlStreamWriter_WriteProcessingInstruction(this.h, target.cPointer())
}

func (this *QXmlStreamWriter) WriteStartDocument() {
	QXmlStreamWriter_WriteStartDocument(this.h)
}

func (this *QXmlStreamWriter) WriteStartDocumentWithVersion(version QAnyStringView) {
	QXmlStreamWriter_WriteStartDocumentWithVersion(this.h, version.cPointer())
}

func (this *QXmlStreamWriter) WriteStartDocument2(version QAnyStringView, standalone bool) {
	QXmlStreamWriter_WriteStartDocument2(this.h, version.cPointer(), (bool)(standalone))
}

func (this *QXmlStreamWriter) WriteStartElement(qualifiedName QAnyStringView) {
	QXmlStreamWriter_WriteStartElement(this.h, qualifiedName.cPointer())
}

func (this *QXmlStreamWriter) WriteStartElement2(namespaceUri QAnyStringView, name QAnyStringView) {
	QXmlStreamWriter_WriteStartElement2(this.h, namespaceUri.cPointer(), name.cPointer())
}

func (this *QXmlStreamWriter) WriteCurrentToken(reader *QXmlStreamReader) {
	QXmlStreamWriter_WriteCurrentToken(this.h, reader.cPointer())
}

func (this *QXmlStreamWriter) HasError() bool {
	return (bool)(QXmlStreamWriter_HasError(this.h))
}

func (this *QXmlStreamWriter) WriteNamespace2(namespaceUri QAnyStringView, prefix QAnyStringView) {
	QXmlStreamWriter_WriteNamespace2(this.h, namespaceUri.cPointer(), prefix.cPointer())
}

func (this *QXmlStreamWriter) WriteProcessingInstruction2(target QAnyStringView, data QAnyStringView) {
	QXmlStreamWriter_WriteProcessingInstruction2(this.h, target.cPointer(), data.cPointer())
}
