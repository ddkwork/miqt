package qt

import (
	"unsafe"
)

type QPluginMetaData struct {
	h          uintptr
	isSubclass bool
}

func QPluginMetaData_ArchRequirements() byte {
	return (byte)(QPluginMetaData_ArchRequirements())
}

type QStaticPlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQStaticPlugin constructs a new QStaticPlugin object.
func NewQStaticPlugin(param1 *QStaticPlugin) *QStaticPlugin {
	g := newQStaticPlugin(QStaticPlugin_new(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QStaticPlugin) MetaData() *QJsonObject {
	_goptr := newQJsonObject(QStaticPlugin_MetaData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QPluginMetaData__Header struct {
	h          uintptr
	isSubclass bool
}

// NewQPluginMetaData__Header constructs a new QPluginMetaData::Header object.
func NewQPluginMetaData__Header(param1 *Header) *QPluginMetaData__Header {
	g := newQPluginMetaData__Header(QPluginMetaData__Header_new(param1))
	g.isSubclass = true
	return g
}

type QPluginMetaData__MagicHeader struct {
	h          uintptr
	isSubclass bool
}

// NewQPluginMetaData__MagicHeader constructs a new QPluginMetaData::MagicHeader object.
func NewQPluginMetaData__MagicHeader() *QPluginMetaData__MagicHeader {
	g := newQPluginMetaData__MagicHeader(QPluginMetaData__MagicHeader_new())
	g.isSubclass = true
	return g
}

type QPluginMetaData__ElfNoteHeader struct {
	h          uintptr
	isSubclass bool
}

// NewQPluginMetaData__ElfNoteHeader constructs a new QPluginMetaData::ElfNoteHeader object.
func NewQPluginMetaData__ElfNoteHeader(payloadSize uint) *QPluginMetaData__ElfNoteHeader {
	g := newQPluginMetaData__ElfNoteHeader(QPluginMetaData__ElfNoteHeader_new((uint)(payloadSize)))
	g.isSubclass = true
	return g
}

// NewQPluginMetaData__ElfNoteHeader2 constructs a new QPluginMetaData::ElfNoteHeader object.
func NewQPluginMetaData__ElfNoteHeader2(param1 *ElfNoteHeader) *QPluginMetaData__ElfNoteHeader {
	g := newQPluginMetaData__ElfNoteHeader(QPluginMetaData__ElfNoteHeader_new2(param1))
	g.isSubclass = true
	return g
}
