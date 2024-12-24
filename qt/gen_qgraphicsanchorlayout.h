#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSANCHORLAYOUT_H
#define MIQT_QT_GEN_QGRAPHICSANCHORLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QGraphicsAnchor QGraphicsAnchor;
typedef struct QGraphicsAnchorLayout QGraphicsAnchorLayout;
typedef struct QGraphicsLayout QGraphicsLayout;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
void QGraphicsAnchor_virtbase(QGraphicsAnchor* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QGraphicsAnchor_MetaObject(const QGraphicsAnchor* self);
extern __declspec(dllexport) 
void* QGraphicsAnchor_Metacast(QGraphicsAnchor* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGraphicsAnchor_Tr(const char* s);
extern __declspec(dllexport) 
void QGraphicsAnchor_SetSpacing(QGraphicsAnchor* self, double spacing);
extern __declspec(dllexport) 
void QGraphicsAnchor_UnsetSpacing(QGraphicsAnchor* self);
extern __declspec(dllexport) 
double QGraphicsAnchor_Spacing(const QGraphicsAnchor* self);
extern __declspec(dllexport) 
void QGraphicsAnchor_SetSizePolicy(QGraphicsAnchor* self, int policy);
extern __declspec(dllexport) 
int QGraphicsAnchor_SizePolicy(const QGraphicsAnchor* self);
extern __declspec(dllexport) 
struct miqt_string QGraphicsAnchor_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGraphicsAnchor_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QGraphicsAnchor_Delete(QGraphicsAnchor* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsAnchorLayout* QGraphicsAnchorLayout_new();
extern __declspec(dllexport) 
QGraphicsAnchorLayout* QGraphicsAnchorLayout_new2(QGraphicsLayoutItem* parent);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_virtbase(QGraphicsAnchorLayout* src
, QGraphicsLayout** outptr_QGraphicsLayout
);
extern __declspec(dllexport) 
QGraphicsAnchor* QGraphicsAnchorLayout_AddAnchor(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstEdge, QGraphicsLayoutItem* secondItem, int secondEdge);
extern __declspec(dllexport) 
QGraphicsAnchor* QGraphicsAnchorLayout_Anchor(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstEdge, QGraphicsLayoutItem* secondItem, int secondEdge);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_AddCornerAnchors(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstCorner, QGraphicsLayoutItem* secondItem, int secondCorner);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_AddAnchors(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, QGraphicsLayoutItem* secondItem);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_SetHorizontalSpacing(QGraphicsAnchorLayout* self, double spacing);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_SetVerticalSpacing(QGraphicsAnchorLayout* self, double spacing);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_SetSpacing(QGraphicsAnchorLayout* self, double spacing);
extern __declspec(dllexport) 
double QGraphicsAnchorLayout_HorizontalSpacing(const QGraphicsAnchorLayout* self);
extern __declspec(dllexport) 
double QGraphicsAnchorLayout_VerticalSpacing(const QGraphicsAnchorLayout* self);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_RemoveAt(QGraphicsAnchorLayout* self, int index);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_SetGeometry(QGraphicsAnchorLayout* self, QRectF* rect);
extern __declspec(dllexport) 
int QGraphicsAnchorLayout_Count(const QGraphicsAnchorLayout* self);
extern __declspec(dllexport) 
QGraphicsLayoutItem* QGraphicsAnchorLayout_ItemAt(const QGraphicsAnchorLayout* self, int index);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_Invalidate(QGraphicsAnchorLayout* self);
extern __declspec(dllexport) 
QSizeF* QGraphicsAnchorLayout_SizeHint(const QGraphicsAnchorLayout* self, int which, QSizeF* constraint);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_AddAnchors3(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, QGraphicsLayoutItem* secondItem, int orientations);
extern __declspec(dllexport) 
void QGraphicsAnchorLayout_Delete(QGraphicsAnchorLayout* self, bool isSubclass);

}
