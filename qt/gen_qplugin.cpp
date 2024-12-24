// +build ignore

#include <QJsonObject>
#include <QPluginMetaData>
#define WORKAROUND_INNER_CLASS_DEFINITION_QPluginMetaData__ElfNoteHeader
#define WORKAROUND_INNER_CLASS_DEFINITION_QPluginMetaData__Header
#define WORKAROUND_INNER_CLASS_DEFINITION_QPluginMetaData__MagicHeader
#include <QStaticPlugin>
#include <qplugin.h>
#include "gen_qplugin.h"
unsigned char QPluginMetaData_ArchRequirements() {
	quint8 _ret = QPluginMetaData::archRequirements();
	return static_cast<unsigned char>(_ret);
}
void QPluginMetaData_Delete(QPluginMetaData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPluginMetaData*>( self );
	} else {
		delete self;
	}
}
QStaticPlugin* QStaticPlugin_new(QStaticPlugin* param1) {
return new QStaticPlugin(*param1);
}
QJsonObject* QStaticPlugin_MetaData(const QStaticPlugin* self) {
	return new QJsonObject(self->metaData());
}
void QStaticPlugin_Delete(QStaticPlugin* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QStaticPlugin*>( self );
	} else {
		delete self;
	}
}
QPluginMetaData__Header* QPluginMetaData__Header_new(const Header* param1) {
return new QPluginMetaData::Header(*param1);
}
void QPluginMetaData__Header_Delete(QPluginMetaData__Header* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPluginMetaData::Header*>( self );
	} else {
		delete self;
	}
}
QPluginMetaData__MagicHeader* QPluginMetaData__MagicHeader_new() {
return new QPluginMetaData::MagicHeader();
}
void QPluginMetaData__MagicHeader_Delete(QPluginMetaData__MagicHeader* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPluginMetaData::MagicHeader*>( self );
	} else {
		delete self;
	}
}
QPluginMetaData__ElfNoteHeader* QPluginMetaData__ElfNoteHeader_new(unsigned int payloadSize) {
return new QPluginMetaData::ElfNoteHeader(static_cast<quint32>(payloadSize));
}
QPluginMetaData__ElfNoteHeader* QPluginMetaData__ElfNoteHeader_new2(const ElfNoteHeader* param1) {
return new QPluginMetaData::ElfNoteHeader(*param1);
}
void QPluginMetaData__ElfNoteHeader_Delete(QPluginMetaData__ElfNoteHeader* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPluginMetaData::ElfNoteHeader*>( self );
	} else {
		delete self;
	}
}
