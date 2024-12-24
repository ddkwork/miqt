#pragma once
#ifndef MIQT_QT_GEN_QPALETTE_H
#define MIQT_QT_GEN_QPALETTE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBrush QBrush;
typedef struct QColor QColor;
typedef struct QPalette QPalette;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPalette* QPalette_new();
extern __declspec(dllexport) QPalette* QPalette_new2(QColor* button);
extern __declspec(dllexport) QPalette* QPalette_new3(int button);
extern __declspec(dllexport) QPalette* QPalette_new4(QColor* button, QColor* window);
extern __declspec(dllexport) QPalette* QPalette_new5(QBrush* windowText, QBrush* button, QBrush* light, QBrush* dark, QBrush* mid, QBrush* text, QBrush* bright_text, QBrush* base, QBrush* window);
extern __declspec(dllexport) QPalette* QPalette_new6(QColor* windowText, QColor* window, QColor* light, QColor* dark, QColor* mid, QColor* text, QColor* base);
extern __declspec(dllexport) QPalette* QPalette_new7(QPalette* palette);
extern __declspec(dllexport) void QPalette_OperatorAssign(QPalette* self, QPalette* palette);
extern __declspec(dllexport) void QPalette_Swap(QPalette* self, QPalette* other);
extern __declspec(dllexport) ColorGroup QPalette_CurrentColorGroup(const QPalette* self);
extern __declspec(dllexport) void QPalette_SetCurrentColorGroup(QPalette* self, ColorGroup cg);
extern __declspec(dllexport) QColor* QPalette_Color(const QPalette* self, ColorGroup cg, ColorRole cr);
extern __declspec(dllexport) QBrush* QPalette_Brush(const QPalette* self, ColorGroup cg, ColorRole cr);
extern __declspec(dllexport) void QPalette_SetColor(QPalette* self, ColorGroup cg, ColorRole cr, QColor* color);
extern __declspec(dllexport) void QPalette_SetColor2(QPalette* self, ColorRole cr, QColor* color);
extern __declspec(dllexport) void QPalette_SetBrush(QPalette* self, ColorRole cr, QBrush* brush);
extern __declspec(dllexport) bool QPalette_IsBrushSet(const QPalette* self, ColorGroup cg, ColorRole cr);
extern __declspec(dllexport) void QPalette_SetBrush2(QPalette* self, ColorGroup cg, ColorRole cr, QBrush* brush);
extern __declspec(dllexport) void QPalette_SetColorGroup(QPalette* self, ColorGroup cr, QBrush* windowText, QBrush* button, QBrush* light, QBrush* dark, QBrush* mid, QBrush* text, QBrush* bright_text, QBrush* base, QBrush* window);
extern __declspec(dllexport) bool QPalette_IsEqual(const QPalette* self, ColorGroup cr1, ColorGroup cr2);
extern __declspec(dllexport) QColor* QPalette_ColorWithCr(const QPalette* self, ColorRole cr);
extern __declspec(dllexport) QBrush* QPalette_BrushWithCr(const QPalette* self, ColorRole cr);
extern __declspec(dllexport) QBrush* QPalette_WindowText(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Button(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Light(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Dark(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Mid(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Text(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Base(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_AlternateBase(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_ToolTipBase(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_ToolTipText(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Window(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Midlight(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_BrightText(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_ButtonText(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Shadow(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Highlight(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_HighlightedText(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Link(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_LinkVisited(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_PlaceholderText(const QPalette* self);
extern __declspec(dllexport) QBrush* QPalette_Accent(const QPalette* self);
extern __declspec(dllexport) bool QPalette_OperatorEqual(const QPalette* self, QPalette* p);
extern __declspec(dllexport) bool QPalette_OperatorNotEqual(const QPalette* self, QPalette* p);
extern __declspec(dllexport) bool QPalette_IsCopyOf(const QPalette* self, QPalette* p);
extern __declspec(dllexport) long long QPalette_CacheKey(const QPalette* self);
extern __declspec(dllexport) QPalette* QPalette_Resolve(const QPalette* self, QPalette* other);
extern __declspec(dllexport) ResolveMask QPalette_ResolveMask(const QPalette* self);
extern __declspec(dllexport) void QPalette_SetResolveMask(QPalette* self, ResolveMask mask);
extern __declspec(dllexport) void QPalette_Delete(QPalette* self, bool isSubclass);

} 
