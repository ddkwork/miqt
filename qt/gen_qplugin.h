#pragma once
#ifndef MIQT_QT_GEN_QPLUGIN_H
#define MIQT_QT_GEN_QPLUGIN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPluginMetaData__ElfNoteHeader)
typedef QPluginMetaData::ElfNoteHeader QPluginMetaData__ElfNoteHeader;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPluginMetaData__Header)
typedef QPluginMetaData::Header QPluginMetaData__Header;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPluginMetaData__MagicHeader)
typedef QPluginMetaData::MagicHeader QPluginMetaData__MagicHeader;
typedef struct QJsonObject QJsonObject;
typedef struct QPluginMetaData QPluginMetaData;
typedef struct QPluginMetaData__ElfNoteHeader QPluginMetaData__ElfNoteHeader;
typedef struct QPluginMetaData__Header QPluginMetaData__Header;
typedef struct QPluginMetaData__MagicHeader QPluginMetaData__MagicHeader;
typedef struct QStaticPlugin QStaticPlugin;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) unsigned char QPluginMetaData_ArchRequirements();
extern __declspec(dllexport) void QPluginMetaData_Delete(QPluginMetaData* self, bool isSubclass);

extern __declspec(dllexport) QStaticPlugin* QStaticPlugin_new(QStaticPlugin* param1);
extern __declspec(dllexport) QJsonObject* QStaticPlugin_MetaData(const QStaticPlugin* self);
extern __declspec(dllexport) void QStaticPlugin_Delete(QStaticPlugin* self, bool isSubclass);

extern __declspec(dllexport) QPluginMetaData__Header* QPluginMetaData__Header_new(const Header* param1);
extern __declspec(dllexport) void QPluginMetaData__Header_Delete(QPluginMetaData__Header* self, bool isSubclass);

extern __declspec(dllexport) QPluginMetaData__MagicHeader* QPluginMetaData__MagicHeader_new();
extern __declspec(dllexport) void QPluginMetaData__MagicHeader_Delete(QPluginMetaData__MagicHeader* self, bool isSubclass);

extern __declspec(dllexport) QPluginMetaData__ElfNoteHeader* QPluginMetaData__ElfNoteHeader_new(unsigned int payloadSize);
extern __declspec(dllexport) QPluginMetaData__ElfNoteHeader* QPluginMetaData__ElfNoteHeader_new2(const ElfNoteHeader* param1);
extern __declspec(dllexport) void QPluginMetaData__ElfNoteHeader_Delete(QPluginMetaData__ElfNoteHeader* self, bool isSubclass);

} 
