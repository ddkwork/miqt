#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSSCENEEVENT_H
#define MIQT_QT_GEN_QGRAPHICSSCENEEVENT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QGraphicsSceneContextMenuEvent QGraphicsSceneContextMenuEvent;
typedef struct QGraphicsSceneDragDropEvent QGraphicsSceneDragDropEvent;
typedef struct QGraphicsSceneEvent QGraphicsSceneEvent;
typedef struct QGraphicsSceneHelpEvent QGraphicsSceneHelpEvent;
typedef struct QGraphicsSceneHoverEvent QGraphicsSceneHoverEvent;
typedef struct QGraphicsSceneMouseEvent QGraphicsSceneMouseEvent;
typedef struct QGraphicsSceneMoveEvent QGraphicsSceneMoveEvent;
typedef struct QGraphicsSceneResizeEvent QGraphicsSceneResizeEvent;
typedef struct QGraphicsSceneWheelEvent QGraphicsSceneWheelEvent;
typedef struct QMimeData QMimeData;
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct QSizeF QSizeF;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneEvent* QGraphicsSceneEvent_new(Type typeVal);
extern __declspec(dllexport) void QGraphicsSceneEvent_virtbase(QGraphicsSceneEvent* src, QEvent** outptr_QEvent);
extern __declspec(dllexport) QWidget* QGraphicsSceneEvent_Widget(const QGraphicsSceneEvent* self);
extern __declspec(dllexport) void QGraphicsSceneEvent_SetWidget(QGraphicsSceneEvent* self, QWidget* widget);
extern __declspec(dllexport) unsigned long long QGraphicsSceneEvent_Timestamp(const QGraphicsSceneEvent* self);
extern __declspec(dllexport) void QGraphicsSceneEvent_SetTimestamp(QGraphicsSceneEvent* self, unsigned long long ts);
extern __declspec(dllexport) void QGraphicsSceneEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QGraphicsSceneEvent_virtualbase_SetAccepted(void* self, bool accepted);
extern __declspec(dllexport) void QGraphicsSceneEvent_override_virtual_Clone(void* self, intptr_t slot);
QEvent* QGraphicsSceneEvent_virtualbase_Clone(const void* self);
extern __declspec(dllexport) void QGraphicsSceneEvent_Delete(QGraphicsSceneEvent* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneMouseEvent* QGraphicsSceneMouseEvent_new();
extern __declspec(dllexport) QGraphicsSceneMouseEvent* QGraphicsSceneMouseEvent_new2(Type typeVal);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_virtbase(QGraphicsSceneMouseEvent* src, QGraphicsSceneEvent** outptr_QGraphicsSceneEvent);
extern __declspec(dllexport) QPointF* QGraphicsSceneMouseEvent_Pos(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetPos(QGraphicsSceneMouseEvent* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneMouseEvent_ScenePos(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetScenePos(QGraphicsSceneMouseEvent* self, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneMouseEvent_ScreenPos(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetScreenPos(QGraphicsSceneMouseEvent* self, QPoint* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneMouseEvent_ButtonDownPos(const QGraphicsSceneMouseEvent* self, int button);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetButtonDownPos(QGraphicsSceneMouseEvent* self, int button, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneMouseEvent_ButtonDownScenePos(const QGraphicsSceneMouseEvent* self, int button);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetButtonDownScenePos(QGraphicsSceneMouseEvent* self, int button, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneMouseEvent_ButtonDownScreenPos(const QGraphicsSceneMouseEvent* self, int button);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetButtonDownScreenPos(QGraphicsSceneMouseEvent* self, int button, QPoint* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneMouseEvent_LastPos(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetLastPos(QGraphicsSceneMouseEvent* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneMouseEvent_LastScenePos(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetLastScenePos(QGraphicsSceneMouseEvent* self, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneMouseEvent_LastScreenPos(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetLastScreenPos(QGraphicsSceneMouseEvent* self, QPoint* pos);
extern __declspec(dllexport) int QGraphicsSceneMouseEvent_Buttons(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetButtons(QGraphicsSceneMouseEvent* self, int buttons);
extern __declspec(dllexport) int QGraphicsSceneMouseEvent_Button(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetButton(QGraphicsSceneMouseEvent* self, int button);
extern __declspec(dllexport) int QGraphicsSceneMouseEvent_Modifiers(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetModifiers(QGraphicsSceneMouseEvent* self, int modifiers);
extern __declspec(dllexport) int QGraphicsSceneMouseEvent_Source(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetSource(QGraphicsSceneMouseEvent* self, int source);
extern __declspec(dllexport) int QGraphicsSceneMouseEvent_Flags(const QGraphicsSceneMouseEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_SetFlags(QGraphicsSceneMouseEvent* self, int flags);
extern __declspec(dllexport) void QGraphicsSceneMouseEvent_Delete(QGraphicsSceneMouseEvent* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneWheelEvent* QGraphicsSceneWheelEvent_new();
extern __declspec(dllexport) QGraphicsSceneWheelEvent* QGraphicsSceneWheelEvent_new2(Type typeVal);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_virtbase(QGraphicsSceneWheelEvent* src, QGraphicsSceneEvent** outptr_QGraphicsSceneEvent);
extern __declspec(dllexport) QPointF* QGraphicsSceneWheelEvent_Pos(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetPos(QGraphicsSceneWheelEvent* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneWheelEvent_ScenePos(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetScenePos(QGraphicsSceneWheelEvent* self, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneWheelEvent_ScreenPos(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetScreenPos(QGraphicsSceneWheelEvent* self, QPoint* pos);
extern __declspec(dllexport) int QGraphicsSceneWheelEvent_Buttons(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetButtons(QGraphicsSceneWheelEvent* self, int buttons);
extern __declspec(dllexport) int QGraphicsSceneWheelEvent_Modifiers(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetModifiers(QGraphicsSceneWheelEvent* self, int modifiers);
extern __declspec(dllexport) int QGraphicsSceneWheelEvent_Delta(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetDelta(QGraphicsSceneWheelEvent* self, int delta);
extern __declspec(dllexport) int QGraphicsSceneWheelEvent_Orientation(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetOrientation(QGraphicsSceneWheelEvent* self, int orientation);
extern __declspec(dllexport) int QGraphicsSceneWheelEvent_Phase(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetPhase(QGraphicsSceneWheelEvent* self, int scrollPhase);
extern __declspec(dllexport) QPoint* QGraphicsSceneWheelEvent_PixelDelta(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetPixelDelta(QGraphicsSceneWheelEvent* self, QPoint* delta);
extern __declspec(dllexport) bool QGraphicsSceneWheelEvent_IsInverted(const QGraphicsSceneWheelEvent* self);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_SetInverted(QGraphicsSceneWheelEvent* self, bool inverted);
extern __declspec(dllexport) void QGraphicsSceneWheelEvent_Delete(QGraphicsSceneWheelEvent* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneContextMenuEvent* QGraphicsSceneContextMenuEvent_new();
extern __declspec(dllexport) QGraphicsSceneContextMenuEvent* QGraphicsSceneContextMenuEvent_new2(Type typeVal);
extern __declspec(dllexport) void QGraphicsSceneContextMenuEvent_virtbase(QGraphicsSceneContextMenuEvent* src, QGraphicsSceneEvent** outptr_QGraphicsSceneEvent);
extern __declspec(dllexport) QPointF* QGraphicsSceneContextMenuEvent_Pos(const QGraphicsSceneContextMenuEvent* self);
extern __declspec(dllexport) void QGraphicsSceneContextMenuEvent_SetPos(QGraphicsSceneContextMenuEvent* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneContextMenuEvent_ScenePos(const QGraphicsSceneContextMenuEvent* self);
extern __declspec(dllexport) void QGraphicsSceneContextMenuEvent_SetScenePos(QGraphicsSceneContextMenuEvent* self, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneContextMenuEvent_ScreenPos(const QGraphicsSceneContextMenuEvent* self);
extern __declspec(dllexport) void QGraphicsSceneContextMenuEvent_SetScreenPos(QGraphicsSceneContextMenuEvent* self, QPoint* pos);
extern __declspec(dllexport) int QGraphicsSceneContextMenuEvent_Modifiers(const QGraphicsSceneContextMenuEvent* self);
extern __declspec(dllexport) void QGraphicsSceneContextMenuEvent_SetModifiers(QGraphicsSceneContextMenuEvent* self, int modifiers);
extern __declspec(dllexport) Reason QGraphicsSceneContextMenuEvent_Reason(const QGraphicsSceneContextMenuEvent* self);
extern __declspec(dllexport) void QGraphicsSceneContextMenuEvent_SetReason(QGraphicsSceneContextMenuEvent* self, Reason reason);
extern __declspec(dllexport) void QGraphicsSceneContextMenuEvent_Delete(QGraphicsSceneContextMenuEvent* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneHoverEvent* QGraphicsSceneHoverEvent_new();
extern __declspec(dllexport) QGraphicsSceneHoverEvent* QGraphicsSceneHoverEvent_new2(Type typeVal);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_virtbase(QGraphicsSceneHoverEvent* src, QGraphicsSceneEvent** outptr_QGraphicsSceneEvent);
extern __declspec(dllexport) QPointF* QGraphicsSceneHoverEvent_Pos(const QGraphicsSceneHoverEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_SetPos(QGraphicsSceneHoverEvent* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneHoverEvent_ScenePos(const QGraphicsSceneHoverEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_SetScenePos(QGraphicsSceneHoverEvent* self, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneHoverEvent_ScreenPos(const QGraphicsSceneHoverEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_SetScreenPos(QGraphicsSceneHoverEvent* self, QPoint* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneHoverEvent_LastPos(const QGraphicsSceneHoverEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_SetLastPos(QGraphicsSceneHoverEvent* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneHoverEvent_LastScenePos(const QGraphicsSceneHoverEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_SetLastScenePos(QGraphicsSceneHoverEvent* self, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneHoverEvent_LastScreenPos(const QGraphicsSceneHoverEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_SetLastScreenPos(QGraphicsSceneHoverEvent* self, QPoint* pos);
extern __declspec(dllexport) int QGraphicsSceneHoverEvent_Modifiers(const QGraphicsSceneHoverEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_SetModifiers(QGraphicsSceneHoverEvent* self, int modifiers);
extern __declspec(dllexport) void QGraphicsSceneHoverEvent_Delete(QGraphicsSceneHoverEvent* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneHelpEvent* QGraphicsSceneHelpEvent_new();
extern __declspec(dllexport) QGraphicsSceneHelpEvent* QGraphicsSceneHelpEvent_new2(Type typeVal);
extern __declspec(dllexport) void QGraphicsSceneHelpEvent_virtbase(QGraphicsSceneHelpEvent* src, QGraphicsSceneEvent** outptr_QGraphicsSceneEvent);
extern __declspec(dllexport) QPointF* QGraphicsSceneHelpEvent_ScenePos(const QGraphicsSceneHelpEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHelpEvent_SetScenePos(QGraphicsSceneHelpEvent* self, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneHelpEvent_ScreenPos(const QGraphicsSceneHelpEvent* self);
extern __declspec(dllexport) void QGraphicsSceneHelpEvent_SetScreenPos(QGraphicsSceneHelpEvent* self, QPoint* pos);
extern __declspec(dllexport) void QGraphicsSceneHelpEvent_Delete(QGraphicsSceneHelpEvent* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneDragDropEvent* QGraphicsSceneDragDropEvent_new();
extern __declspec(dllexport) QGraphicsSceneDragDropEvent* QGraphicsSceneDragDropEvent_new2(Type typeVal);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_virtbase(QGraphicsSceneDragDropEvent* src, QGraphicsSceneEvent** outptr_QGraphicsSceneEvent);
extern __declspec(dllexport) QPointF* QGraphicsSceneDragDropEvent_Pos(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetPos(QGraphicsSceneDragDropEvent* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneDragDropEvent_ScenePos(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetScenePos(QGraphicsSceneDragDropEvent* self, QPointF* pos);
extern __declspec(dllexport) QPoint* QGraphicsSceneDragDropEvent_ScreenPos(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetScreenPos(QGraphicsSceneDragDropEvent* self, QPoint* pos);
extern __declspec(dllexport) int QGraphicsSceneDragDropEvent_Buttons(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetButtons(QGraphicsSceneDragDropEvent* self, int buttons);
extern __declspec(dllexport) int QGraphicsSceneDragDropEvent_Modifiers(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetModifiers(QGraphicsSceneDragDropEvent* self, int modifiers);
extern __declspec(dllexport) int QGraphicsSceneDragDropEvent_PossibleActions(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetPossibleActions(QGraphicsSceneDragDropEvent* self, int actions);
extern __declspec(dllexport) int QGraphicsSceneDragDropEvent_ProposedAction(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetProposedAction(QGraphicsSceneDragDropEvent* self, int action);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_AcceptProposedAction(QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) int QGraphicsSceneDragDropEvent_DropAction(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetDropAction(QGraphicsSceneDragDropEvent* self, int action);
extern __declspec(dllexport) QWidget* QGraphicsSceneDragDropEvent_Source(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetSource(QGraphicsSceneDragDropEvent* self, QWidget* source);
extern __declspec(dllexport) QMimeData* QGraphicsSceneDragDropEvent_MimeData(const QGraphicsSceneDragDropEvent* self);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_SetMimeData(QGraphicsSceneDragDropEvent* self, QMimeData* data);
extern __declspec(dllexport) void QGraphicsSceneDragDropEvent_Delete(QGraphicsSceneDragDropEvent* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneResizeEvent* QGraphicsSceneResizeEvent_new();
extern __declspec(dllexport) void QGraphicsSceneResizeEvent_virtbase(QGraphicsSceneResizeEvent* src, QGraphicsSceneEvent** outptr_QGraphicsSceneEvent);
extern __declspec(dllexport) QSizeF* QGraphicsSceneResizeEvent_OldSize(const QGraphicsSceneResizeEvent* self);
extern __declspec(dllexport) void QGraphicsSceneResizeEvent_SetOldSize(QGraphicsSceneResizeEvent* self, QSizeF* size);
extern __declspec(dllexport) QSizeF* QGraphicsSceneResizeEvent_NewSize(const QGraphicsSceneResizeEvent* self);
extern __declspec(dllexport) void QGraphicsSceneResizeEvent_SetNewSize(QGraphicsSceneResizeEvent* self, QSizeF* size);
extern __declspec(dllexport) void QGraphicsSceneResizeEvent_Delete(QGraphicsSceneResizeEvent* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSceneMoveEvent* QGraphicsSceneMoveEvent_new();
extern __declspec(dllexport) void QGraphicsSceneMoveEvent_virtbase(QGraphicsSceneMoveEvent* src, QGraphicsSceneEvent** outptr_QGraphicsSceneEvent);
extern __declspec(dllexport) QPointF* QGraphicsSceneMoveEvent_OldPos(const QGraphicsSceneMoveEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMoveEvent_SetOldPos(QGraphicsSceneMoveEvent* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QGraphicsSceneMoveEvent_NewPos(const QGraphicsSceneMoveEvent* self);
extern __declspec(dllexport) void QGraphicsSceneMoveEvent_SetNewPos(QGraphicsSceneMoveEvent* self, QPointF* pos);
extern __declspec(dllexport) void QGraphicsSceneMoveEvent_Delete(QGraphicsSceneMoveEvent* self, bool isSubclass);

} 
