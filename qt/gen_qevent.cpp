// +build ignore

#include <QAction>
#include <QActionEvent>
#include <QApplicationStateChangeEvent>
#include <QChildWindowEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEnterEvent>
#include <QEvent>
#include <QEventPoint>
#include <QExposeEvent>
#include <QFile>
#include <QFileOpenEvent>
#include <QFocusEvent>
#include <QHelpEvent>
#include <QHideEvent>
#include <QHoverEvent>
#include <QIconDragEvent>
#include <QInputDevice>
#include <QInputEvent>
#include <QInputMethodEvent>
#define WORKAROUND_INNER_CLASS_DEFINITION_QInputMethodEvent__Attribute
#include <QInputMethodQueryEvent>
#include <QKeyCombination>
#include <QKeyEvent>
#include <QKeySequence>
#include <QList>
#include <QMimeData>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QNativeGestureEvent>
#include <QObject>
#include <QPaintEvent>
#include <QPoint>
#include <QPointF>
#include <QPointerEvent>
#include <QPointingDevice>
#include <QRect>
#include <QRectF>
#include <QRegion>
#include <QResizeEvent>
#include <QScreen>
#include <QScreenOrientationChangeEvent>
#include <QScrollEvent>
#include <QScrollPrepareEvent>
#include <QShortcut>
#include <QShortcutEvent>
#include <QShowEvent>
#include <QSinglePointEvent>
#include <QSize>
#include <QSizeF>
#include <QStatusTipEvent>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTabletEvent>
#include <QToolBarChangeEvent>
#include <QTouchEvent>
#include <QUrl>
#include <QVariant>
#include <QWhatsThisClickedEvent>
#include <QWheelEvent>
#include <QWindow>
#include <QWindowStateChangeEvent>
#include <qevent.h>
#include "gen_qevent.h"
QInputEvent* QInputEvent_new(Type typeVal, QInputDevice* m_dev) {
return new QInputEvent(typeVal, m_dev);
}
QInputEvent* QInputEvent_new2(Type typeVal, QInputDevice* m_dev, int modifiers) {
return new QInputEvent(typeVal, m_dev, static_cast<Qt::KeyboardModifiers>(modifiers));
}
void QInputEvent_virtbase(QInputEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QInputEvent* QInputEvent_Clone(const QInputEvent* self) {
	return self->clone();
}
QInputDevice* QInputEvent_Device(const QInputEvent* self) {
	return (QInputDevice*) self->device();
}
int QInputEvent_DeviceType(const QInputEvent* self) {
	QInputDevice::DeviceType _ret = self->deviceType();
	return static_cast<int>(_ret);
}
int QInputEvent_Modifiers(const QInputEvent* self) {
	Qt::KeyboardModifiers _ret = self->modifiers();
	return static_cast<int>(_ret);
}
void QInputEvent_SetModifiers(QInputEvent* self, int modifiers) {
	self->setModifiers(static_cast<Qt::KeyboardModifiers>(modifiers));
}
unsigned long long QInputEvent_Timestamp(const QInputEvent* self) {
	quint64 _ret = self->timestamp();
	return static_cast<unsigned long long>(_ret);
}
void QInputEvent_SetTimestamp(QInputEvent* self, unsigned long long timestamp) {
	self->setTimestamp(static_cast<quint64>(timestamp));
}
void QInputEvent_Delete(QInputEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QInputEvent*>( self );
	} else {
		delete self;
	}
}
QPointerEvent* QPointerEvent_new(Type typeVal, QPointingDevice* dev) {
return new QPointerEvent(typeVal, dev);
}
QPointerEvent* QPointerEvent_new2(Type typeVal, QPointingDevice* dev, int modifiers) {
return new QPointerEvent(typeVal, dev, static_cast<Qt::KeyboardModifiers>(modifiers));
}
QPointerEvent* QPointerEvent_new3(Type typeVal, QPointingDevice* dev, int modifiers, struct miqt_array /* of QEventPoint* */  points) {
QList<QEventPoint> points_QList;
	points_QList.reserve(points.len);
	QEventPoint** points_arr = static_cast<QEventPoint**>(points.data);
	for(size_t i = 0; i < points.len; ++i) {
		points_QList.push_back(*(points_arr[i]));
	}
	return new QPointerEvent(typeVal, dev, static_cast<Qt::KeyboardModifiers>(modifiers), points_QList);
}
void QPointerEvent_virtbase(QPointerEvent* src
, QInputEvent** outptr_QInputEvent
) {
*outptr_QInputEvent = static_cast<QInputEvent*>(src);
}
QPointerEvent* QPointerEvent_Clone(const QPointerEvent* self) {
	return self->clone();
}
QPointingDevice* QPointerEvent_PointingDevice(const QPointerEvent* self) {
	return (QPointingDevice*) self->pointingDevice();
}
int QPointerEvent_PointerType(const QPointerEvent* self) {
	QPointingDevice::PointerType _ret = self->pointerType();
	return static_cast<int>(_ret);
}
void QPointerEvent_SetTimestamp(QPointerEvent* self, unsigned long long timestamp) {
	self->setTimestamp(static_cast<quint64>(timestamp));
}
ptrdiff_t QPointerEvent_PointCount(const QPointerEvent* self) {
	qsizetype _ret = self->pointCount();
	return static_cast<ptrdiff_t>(_ret);
}
QEventPoint* QPointerEvent_Point(QPointerEvent* self, ptrdiff_t i) {
	QEventPoint& _ret = self->point((qsizetype)(i));
	// Cast returned reference into pointer
	return &_ret;
}
struct miqt_array /* of QEventPoint* */  QPointerEvent_Points(const QPointerEvent* self) {
	const QList<QEventPoint>& _ret = self->points();
	// Convert QList<> from C++ memory to manually-managed C memory
	QEventPoint** _arr = static_cast<QEventPoint**>(malloc(sizeof(QEventPoint*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QEventPoint(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
QEventPoint* QPointerEvent_PointById(QPointerEvent* self, int id) {
	return self->pointById(static_cast<int>(id));
}
bool QPointerEvent_AllPointsGrabbed(const QPointerEvent* self) {
	return self->allPointsGrabbed();
}
bool QPointerEvent_IsBeginEvent(const QPointerEvent* self) {
	return self->isBeginEvent();
}
bool QPointerEvent_IsUpdateEvent(const QPointerEvent* self) {
	return self->isUpdateEvent();
}
bool QPointerEvent_IsEndEvent(const QPointerEvent* self) {
	return self->isEndEvent();
}
bool QPointerEvent_AllPointsAccepted(const QPointerEvent* self) {
	return self->allPointsAccepted();
}
void QPointerEvent_SetAccepted(QPointerEvent* self, bool accepted) {
	self->setAccepted(accepted);
}
QObject* QPointerEvent_ExclusiveGrabber(const QPointerEvent* self, QEventPoint* point) {
	return self->exclusiveGrabber(*point);
}
void QPointerEvent_SetExclusiveGrabber(QPointerEvent* self, QEventPoint* point, QObject* exclusiveGrabber) {
	self->setExclusiveGrabber(*point, exclusiveGrabber);
}
void QPointerEvent_ClearPassiveGrabbers(QPointerEvent* self, QEventPoint* point) {
	self->clearPassiveGrabbers(*point);
}
bool QPointerEvent_AddPassiveGrabber(QPointerEvent* self, QEventPoint* point, QObject* grabber) {
	return self->addPassiveGrabber(*point, grabber);
}
bool QPointerEvent_RemovePassiveGrabber(QPointerEvent* self, QEventPoint* point, QObject* grabber) {
	return self->removePassiveGrabber(*point, grabber);
}
void QPointerEvent_Delete(QPointerEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPointerEvent*>( self );
	} else {
		delete self;
	}
}
void QSinglePointEvent_virtbase(QSinglePointEvent* src
, QPointerEvent** outptr_QPointerEvent
) {
*outptr_QPointerEvent = static_cast<QPointerEvent*>(src);
}
QSinglePointEvent* QSinglePointEvent_Clone(const QSinglePointEvent* self) {
	return self->clone();
}
int QSinglePointEvent_Button(const QSinglePointEvent* self) {
	Qt::MouseButton _ret = self->button();
	return static_cast<int>(_ret);
}
int QSinglePointEvent_Buttons(const QSinglePointEvent* self) {
	Qt::MouseButtons _ret = self->buttons();
	return static_cast<int>(_ret);
}
QPointF* QSinglePointEvent_Position(const QSinglePointEvent* self) {
	return new QPointF(self->position());
}
QPointF* QSinglePointEvent_ScenePosition(const QSinglePointEvent* self) {
	return new QPointF(self->scenePosition());
}
QPointF* QSinglePointEvent_GlobalPosition(const QSinglePointEvent* self) {
	return new QPointF(self->globalPosition());
}
bool QSinglePointEvent_IsBeginEvent(const QSinglePointEvent* self) {
	return self->isBeginEvent();
}
bool QSinglePointEvent_IsUpdateEvent(const QSinglePointEvent* self) {
	return self->isUpdateEvent();
}
bool QSinglePointEvent_IsEndEvent(const QSinglePointEvent* self) {
	return self->isEndEvent();
}
QObject* QSinglePointEvent_ExclusivePointGrabber(const QSinglePointEvent* self) {
	return self->exclusivePointGrabber();
}
void QSinglePointEvent_SetExclusivePointGrabber(QSinglePointEvent* self, QObject* exclusiveGrabber) {
	self->setExclusivePointGrabber(exclusiveGrabber);
}
void QSinglePointEvent_Delete(QSinglePointEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QSinglePointEvent*>( self );
	} else {
		delete self;
	}
}
QEnterEvent* QEnterEvent_new(QPointF* localPos, QPointF* scenePos, QPointF* globalPos) {
return new QEnterEvent(*localPos, *scenePos, *globalPos);
}
QEnterEvent* QEnterEvent_new2(QPointF* localPos, QPointF* scenePos, QPointF* globalPos, QPointingDevice* device) {
return new QEnterEvent(*localPos, *scenePos, *globalPos, device);
}
void QEnterEvent_virtbase(QEnterEvent* src
, QSinglePointEvent** outptr_QSinglePointEvent
) {
*outptr_QSinglePointEvent = static_cast<QSinglePointEvent*>(src);
}
QEnterEvent* QEnterEvent_Clone(const QEnterEvent* self) {
	return self->clone();
}
QPoint* QEnterEvent_Pos(const QEnterEvent* self) {
	return new QPoint(self->pos());
}
QPoint* QEnterEvent_GlobalPos(const QEnterEvent* self) {
	return new QPoint(self->globalPos());
}
int QEnterEvent_X(const QEnterEvent* self) {
	return self->x();
}
int QEnterEvent_Y(const QEnterEvent* self) {
	return self->y();
}
int QEnterEvent_GlobalX(const QEnterEvent* self) {
	return self->globalX();
}
int QEnterEvent_GlobalY(const QEnterEvent* self) {
	return self->globalY();
}
QPointF* QEnterEvent_LocalPos(const QEnterEvent* self) {
	return new QPointF(self->localPos());
}
QPointF* QEnterEvent_WindowPos(const QEnterEvent* self) {
	return new QPointF(self->windowPos());
}
QPointF* QEnterEvent_ScreenPos(const QEnterEvent* self) {
	return new QPointF(self->screenPos());
}
void QEnterEvent_Delete(QEnterEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QEnterEvent*>( self );
	} else {
		delete self;
	}
}
QMouseEvent* QMouseEvent_new(Type typeVal, QPointF* localPos, int button, int buttons, int modifiers) {
return new QMouseEvent(typeVal, *localPos, static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers));
}
QMouseEvent* QMouseEvent_new2(Type typeVal, QPointF* localPos, QPointF* globalPos, int button, int buttons, int modifiers) {
return new QMouseEvent(typeVal, *localPos, *globalPos, static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers));
}
QMouseEvent* QMouseEvent_new3(Type typeVal, QPointF* localPos, QPointF* scenePos, QPointF* globalPos, int button, int buttons, int modifiers) {
return new QMouseEvent(typeVal, *localPos, *scenePos, *globalPos, static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers));
}
QMouseEvent* QMouseEvent_new4(Type typeVal, QPointF* localPos, QPointF* scenePos, QPointF* globalPos, int button, int buttons, int modifiers, int source) {
return new QMouseEvent(typeVal, *localPos, *scenePos, *globalPos, static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<Qt::MouseEventSource>(source));
}
QMouseEvent* QMouseEvent_new5(Type typeVal, QPointF* localPos, int button, int buttons, int modifiers, QPointingDevice* device) {
return new QMouseEvent(typeVal, *localPos, static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), device);
}
QMouseEvent* QMouseEvent_new6(Type typeVal, QPointF* localPos, QPointF* globalPos, int button, int buttons, int modifiers, QPointingDevice* device) {
return new QMouseEvent(typeVal, *localPos, *globalPos, static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), device);
}
QMouseEvent* QMouseEvent_new7(Type typeVal, QPointF* localPos, QPointF* scenePos, QPointF* globalPos, int button, int buttons, int modifiers, QPointingDevice* device) {
return new QMouseEvent(typeVal, *localPos, *scenePos, *globalPos, static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), device);
}
QMouseEvent* QMouseEvent_new8(Type typeVal, QPointF* localPos, QPointF* scenePos, QPointF* globalPos, int button, int buttons, int modifiers, int source, QPointingDevice* device) {
return new QMouseEvent(typeVal, *localPos, *scenePos, *globalPos, static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<Qt::MouseEventSource>(source), device);
}
void QMouseEvent_virtbase(QMouseEvent* src
, QSinglePointEvent** outptr_QSinglePointEvent
) {
*outptr_QSinglePointEvent = static_cast<QSinglePointEvent*>(src);
}
QMouseEvent* QMouseEvent_Clone(const QMouseEvent* self) {
	return self->clone();
}
QPoint* QMouseEvent_Pos(const QMouseEvent* self) {
	return new QPoint(self->pos());
}
QPoint* QMouseEvent_GlobalPos(const QMouseEvent* self) {
	return new QPoint(self->globalPos());
}
int QMouseEvent_X(const QMouseEvent* self) {
	return self->x();
}
int QMouseEvent_Y(const QMouseEvent* self) {
	return self->y();
}
int QMouseEvent_GlobalX(const QMouseEvent* self) {
	return self->globalX();
}
int QMouseEvent_GlobalY(const QMouseEvent* self) {
	return self->globalY();
}
QPointF* QMouseEvent_LocalPos(const QMouseEvent* self) {
	return new QPointF(self->localPos());
}
QPointF* QMouseEvent_WindowPos(const QMouseEvent* self) {
	return new QPointF(self->windowPos());
}
QPointF* QMouseEvent_ScreenPos(const QMouseEvent* self) {
	return new QPointF(self->screenPos());
}
int QMouseEvent_Source(const QMouseEvent* self) {
	Qt::MouseEventSource _ret = self->source();
	return static_cast<int>(_ret);
}
int QMouseEvent_Flags(const QMouseEvent* self) {
	Qt::MouseEventFlags _ret = self->flags();
	return static_cast<int>(_ret);
}
void QMouseEvent_Delete(QMouseEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QMouseEvent*>( self );
	} else {
		delete self;
	}
}
QHoverEvent* QHoverEvent_new(Type typeVal, QPointF* scenePos, QPointF* globalPos, QPointF* oldPos) {
return new QHoverEvent(typeVal, *scenePos, *globalPos, *oldPos);
}
QHoverEvent* QHoverEvent_new2(Type typeVal, QPointF* pos, QPointF* oldPos) {
return new QHoverEvent(typeVal, *pos, *oldPos);
}
QHoverEvent* QHoverEvent_new3(Type typeVal, QPointF* scenePos, QPointF* globalPos, QPointF* oldPos, int modifiers) {
return new QHoverEvent(typeVal, *scenePos, *globalPos, *oldPos, static_cast<Qt::KeyboardModifiers>(modifiers));
}
QHoverEvent* QHoverEvent_new4(Type typeVal, QPointF* scenePos, QPointF* globalPos, QPointF* oldPos, int modifiers, QPointingDevice* device) {
return new QHoverEvent(typeVal, *scenePos, *globalPos, *oldPos, static_cast<Qt::KeyboardModifiers>(modifiers), device);
}
QHoverEvent* QHoverEvent_new5(Type typeVal, QPointF* pos, QPointF* oldPos, int modifiers) {
return new QHoverEvent(typeVal, *pos, *oldPos, static_cast<Qt::KeyboardModifiers>(modifiers));
}
QHoverEvent* QHoverEvent_new6(Type typeVal, QPointF* pos, QPointF* oldPos, int modifiers, QPointingDevice* device) {
return new QHoverEvent(typeVal, *pos, *oldPos, static_cast<Qt::KeyboardModifiers>(modifiers), device);
}
void QHoverEvent_virtbase(QHoverEvent* src
, QSinglePointEvent** outptr_QSinglePointEvent
) {
*outptr_QSinglePointEvent = static_cast<QSinglePointEvent*>(src);
}
QHoverEvent* QHoverEvent_Clone(const QHoverEvent* self) {
	return self->clone();
}
QPoint* QHoverEvent_Pos(const QHoverEvent* self) {
	return new QPoint(self->pos());
}
QPointF* QHoverEvent_PosF(const QHoverEvent* self) {
	return new QPointF(self->posF());
}
bool QHoverEvent_IsUpdateEvent(const QHoverEvent* self) {
	return self->isUpdateEvent();
}
QPoint* QHoverEvent_OldPos(const QHoverEvent* self) {
	return new QPoint(self->oldPos());
}
QPointF* QHoverEvent_OldPosF(const QHoverEvent* self) {
	return new QPointF(self->oldPosF());
}
void QHoverEvent_Delete(QHoverEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QHoverEvent*>( self );
	} else {
		delete self;
	}
}
QWheelEvent* QWheelEvent_new(QPointF* pos, QPointF* globalPos, QPoint* pixelDelta, QPoint* angleDelta, int buttons, int modifiers, int phase, bool inverted) {
return new QWheelEvent(*pos, *globalPos, *pixelDelta, *angleDelta, static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<Qt::ScrollPhase>(phase), inverted);
}
QWheelEvent* QWheelEvent_new2(QPointF* pos, QPointF* globalPos, QPoint* pixelDelta, QPoint* angleDelta, int buttons, int modifiers, int phase, bool inverted, int source) {
return new QWheelEvent(*pos, *globalPos, *pixelDelta, *angleDelta, static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<Qt::ScrollPhase>(phase), inverted, static_cast<Qt::MouseEventSource>(source));
}
QWheelEvent* QWheelEvent_new3(QPointF* pos, QPointF* globalPos, QPoint* pixelDelta, QPoint* angleDelta, int buttons, int modifiers, int phase, bool inverted, int source, QPointingDevice* device) {
return new QWheelEvent(*pos, *globalPos, *pixelDelta, *angleDelta, static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<Qt::ScrollPhase>(phase), inverted, static_cast<Qt::MouseEventSource>(source), device);
}
void QWheelEvent_virtbase(QWheelEvent* src
, QSinglePointEvent** outptr_QSinglePointEvent
) {
*outptr_QSinglePointEvent = static_cast<QSinglePointEvent*>(src);
}
QWheelEvent* QWheelEvent_Clone(const QWheelEvent* self) {
	return self->clone();
}
QPoint* QWheelEvent_PixelDelta(const QWheelEvent* self) {
	return new QPoint(self->pixelDelta());
}
QPoint* QWheelEvent_AngleDelta(const QWheelEvent* self) {
	return new QPoint(self->angleDelta());
}
int QWheelEvent_Phase(const QWheelEvent* self) {
	Qt::ScrollPhase _ret = self->phase();
	return static_cast<int>(_ret);
}
bool QWheelEvent_Inverted(const QWheelEvent* self) {
	return self->inverted();
}
bool QWheelEvent_IsInverted(const QWheelEvent* self) {
	return self->isInverted();
}
bool QWheelEvent_HasPixelDelta(const QWheelEvent* self) {
	return self->hasPixelDelta();
}
bool QWheelEvent_IsBeginEvent(const QWheelEvent* self) {
	return self->isBeginEvent();
}
bool QWheelEvent_IsUpdateEvent(const QWheelEvent* self) {
	return self->isUpdateEvent();
}
bool QWheelEvent_IsEndEvent(const QWheelEvent* self) {
	return self->isEndEvent();
}
int QWheelEvent_Source(const QWheelEvent* self) {
	Qt::MouseEventSource _ret = self->source();
	return static_cast<int>(_ret);
}
void QWheelEvent_Delete(QWheelEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QWheelEvent*>( self );
	} else {
		delete self;
	}
}
QTabletEvent* QTabletEvent_new(Type t, QPointingDevice* device, QPointF* pos, QPointF* globalPos, double pressure, float xTilt, float yTilt, float tangentialPressure, double rotation, float z, int keyState, int button, int buttons) {
return new QTabletEvent(t, device, *pos, *globalPos, static_cast<qreal>(pressure), static_cast<float>(xTilt), static_cast<float>(yTilt), static_cast<float>(tangentialPressure), static_cast<qreal>(rotation), static_cast<float>(z), static_cast<Qt::KeyboardModifiers>(keyState), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButtons>(buttons));
}
void QTabletEvent_virtbase(QTabletEvent* src
, QSinglePointEvent** outptr_QSinglePointEvent
) {
*outptr_QSinglePointEvent = static_cast<QSinglePointEvent*>(src);
}
QTabletEvent* QTabletEvent_Clone(const QTabletEvent* self) {
	return self->clone();
}
QPoint* QTabletEvent_Pos(const QTabletEvent* self) {
	return new QPoint(self->pos());
}
QPoint* QTabletEvent_GlobalPos(const QTabletEvent* self) {
	return new QPoint(self->globalPos());
}
QPointF* QTabletEvent_PosF(const QTabletEvent* self) {
	return new QPointF(self->posF());
}
QPointF* QTabletEvent_GlobalPosF(const QTabletEvent* self) {
	return new QPointF(self->globalPosF());
}
int QTabletEvent_X(const QTabletEvent* self) {
	return self->x();
}
int QTabletEvent_Y(const QTabletEvent* self) {
	return self->y();
}
int QTabletEvent_GlobalX(const QTabletEvent* self) {
	return self->globalX();
}
int QTabletEvent_GlobalY(const QTabletEvent* self) {
	return self->globalY();
}
double QTabletEvent_HiResGlobalX(const QTabletEvent* self) {
	qreal _ret = self->hiResGlobalX();
	return static_cast<double>(_ret);
}
double QTabletEvent_HiResGlobalY(const QTabletEvent* self) {
	qreal _ret = self->hiResGlobalY();
	return static_cast<double>(_ret);
}
long long QTabletEvent_UniqueId(const QTabletEvent* self) {
	qint64 _ret = self->uniqueId();
	return static_cast<long long>(_ret);
}
double QTabletEvent_Pressure(const QTabletEvent* self) {
	qreal _ret = self->pressure();
	return static_cast<double>(_ret);
}
double QTabletEvent_Rotation(const QTabletEvent* self) {
	qreal _ret = self->rotation();
	return static_cast<double>(_ret);
}
double QTabletEvent_Z(const QTabletEvent* self) {
	qreal _ret = self->z();
	return static_cast<double>(_ret);
}
double QTabletEvent_TangentialPressure(const QTabletEvent* self) {
	qreal _ret = self->tangentialPressure();
	return static_cast<double>(_ret);
}
double QTabletEvent_XTilt(const QTabletEvent* self) {
	qreal _ret = self->xTilt();
	return static_cast<double>(_ret);
}
double QTabletEvent_YTilt(const QTabletEvent* self) {
	qreal _ret = self->yTilt();
	return static_cast<double>(_ret);
}
void QTabletEvent_Delete(QTabletEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTabletEvent*>( self );
	} else {
		delete self;
	}
}
QNativeGestureEvent* QNativeGestureEvent_new(int typeVal, QPointingDevice* dev, QPointF* localPos, QPointF* scenePos, QPointF* globalPos, double value, unsigned long long sequenceId, unsigned long long intArgument) {
return new QNativeGestureEvent(static_cast<Qt::NativeGestureType>(typeVal), dev, *localPos, *scenePos, *globalPos, static_cast<qreal>(value), static_cast<quint64>(sequenceId), static_cast<quint64>(intArgument));
}
QNativeGestureEvent* QNativeGestureEvent_new2(int typeVal, QPointingDevice* dev, int fingerCount, QPointF* localPos, QPointF* scenePos, QPointF* globalPos, double value, QPointF* delta) {
return new QNativeGestureEvent(static_cast<Qt::NativeGestureType>(typeVal), dev, static_cast<int>(fingerCount), *localPos, *scenePos, *globalPos, static_cast<qreal>(value), *delta);
}
QNativeGestureEvent* QNativeGestureEvent_new3(int typeVal, QPointingDevice* dev, int fingerCount, QPointF* localPos, QPointF* scenePos, QPointF* globalPos, double value, QPointF* delta, unsigned long long sequenceId) {
return new QNativeGestureEvent(static_cast<Qt::NativeGestureType>(typeVal), dev, static_cast<int>(fingerCount), *localPos, *scenePos, *globalPos, static_cast<qreal>(value), *delta, static_cast<quint64>(sequenceId));
}
void QNativeGestureEvent_virtbase(QNativeGestureEvent* src
, QSinglePointEvent** outptr_QSinglePointEvent
) {
*outptr_QSinglePointEvent = static_cast<QSinglePointEvent*>(src);
}
QNativeGestureEvent* QNativeGestureEvent_Clone(const QNativeGestureEvent* self) {
	return self->clone();
}
int QNativeGestureEvent_GestureType(const QNativeGestureEvent* self) {
	Qt::NativeGestureType _ret = self->gestureType();
	return static_cast<int>(_ret);
}
int QNativeGestureEvent_FingerCount(const QNativeGestureEvent* self) {
	return self->fingerCount();
}
double QNativeGestureEvent_Value(const QNativeGestureEvent* self) {
	qreal _ret = self->value();
	return static_cast<double>(_ret);
}
QPointF* QNativeGestureEvent_Delta(const QNativeGestureEvent* self) {
	return new QPointF(self->delta());
}
QPoint* QNativeGestureEvent_Pos(const QNativeGestureEvent* self) {
	return new QPoint(self->pos());
}
QPoint* QNativeGestureEvent_GlobalPos(const QNativeGestureEvent* self) {
	return new QPoint(self->globalPos());
}
QPointF* QNativeGestureEvent_LocalPos(const QNativeGestureEvent* self) {
	return new QPointF(self->localPos());
}
QPointF* QNativeGestureEvent_WindowPos(const QNativeGestureEvent* self) {
	return new QPointF(self->windowPos());
}
QPointF* QNativeGestureEvent_ScreenPos(const QNativeGestureEvent* self) {
	return new QPointF(self->screenPos());
}
void QNativeGestureEvent_Delete(QNativeGestureEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QNativeGestureEvent*>( self );
	} else {
		delete self;
	}
}
QKeyEvent* QKeyEvent_new(Type typeVal, int key, int modifiers) {
return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers));
}
QKeyEvent* QKeyEvent_new2(Type typeVal, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers) {
return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<quint32>(nativeScanCode), static_cast<quint32>(nativeVirtualKey), static_cast<quint32>(nativeModifiers));
}
QKeyEvent* QKeyEvent_new3(Type typeVal, int key, int modifiers, struct miqt_string text) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers), text_QString);
}
QKeyEvent* QKeyEvent_new4(Type typeVal, int key, int modifiers, struct miqt_string text, bool autorep) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers), text_QString, autorep);
}
QKeyEvent* QKeyEvent_new5(Type typeVal, int key, int modifiers, struct miqt_string text, bool autorep, uint16_t count) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers), text_QString, autorep, static_cast<quint16>(count));
}
QKeyEvent* QKeyEvent_new6(Type typeVal, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, struct miqt_string text) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<quint32>(nativeScanCode), static_cast<quint32>(nativeVirtualKey), static_cast<quint32>(nativeModifiers), text_QString);
}
QKeyEvent* QKeyEvent_new7(Type typeVal, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, struct miqt_string text, bool autorep) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<quint32>(nativeScanCode), static_cast<quint32>(nativeVirtualKey), static_cast<quint32>(nativeModifiers), text_QString, autorep);
}
QKeyEvent* QKeyEvent_new8(Type typeVal, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, struct miqt_string text, bool autorep, uint16_t count) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<quint32>(nativeScanCode), static_cast<quint32>(nativeVirtualKey), static_cast<quint32>(nativeModifiers), text_QString, autorep, static_cast<quint16>(count));
}
QKeyEvent* QKeyEvent_new9(Type typeVal, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, struct miqt_string text, bool autorep, uint16_t count, QInputDevice* device) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QKeyEvent(typeVal, static_cast<int>(key), static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<quint32>(nativeScanCode), static_cast<quint32>(nativeVirtualKey), static_cast<quint32>(nativeModifiers), text_QString, autorep, static_cast<quint16>(count), device);
}
void QKeyEvent_virtbase(QKeyEvent* src
, QInputEvent** outptr_QInputEvent
) {
*outptr_QInputEvent = static_cast<QInputEvent*>(src);
}
QKeyEvent* QKeyEvent_Clone(const QKeyEvent* self) {
	return self->clone();
}
int QKeyEvent_Key(const QKeyEvent* self) {
	return self->key();
}
bool QKeyEvent_Matches(const QKeyEvent* self, int key) {
	return self->matches(static_cast<QKeySequence::StandardKey>(key));
}
int QKeyEvent_Modifiers(const QKeyEvent* self) {
	Qt::KeyboardModifiers _ret = self->modifiers();
	return static_cast<int>(_ret);
}
QKeyCombination* QKeyEvent_KeyCombination(const QKeyEvent* self) {
	return new QKeyCombination(self->keyCombination());
}
struct miqt_string QKeyEvent_Text(const QKeyEvent* self) {
	QString _ret = self->text();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QKeyEvent_IsAutoRepeat(const QKeyEvent* self) {
	return self->isAutoRepeat();
}
int QKeyEvent_Count(const QKeyEvent* self) {
	return self->count();
}
unsigned int QKeyEvent_NativeScanCode(const QKeyEvent* self) {
	quint32 _ret = self->nativeScanCode();
	return static_cast<unsigned int>(_ret);
}
unsigned int QKeyEvent_NativeVirtualKey(const QKeyEvent* self) {
	quint32 _ret = self->nativeVirtualKey();
	return static_cast<unsigned int>(_ret);
}
unsigned int QKeyEvent_NativeModifiers(const QKeyEvent* self) {
	quint32 _ret = self->nativeModifiers();
	return static_cast<unsigned int>(_ret);
}
void QKeyEvent_Delete(QKeyEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QKeyEvent*>( self );
	} else {
		delete self;
	}
}
QFocusEvent* QFocusEvent_new(Type typeVal) {
return new QFocusEvent(typeVal);
}
QFocusEvent* QFocusEvent_new2(Type typeVal, int reason) {
return new QFocusEvent(typeVal, static_cast<Qt::FocusReason>(reason));
}
void QFocusEvent_virtbase(QFocusEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QFocusEvent* QFocusEvent_Clone(const QFocusEvent* self) {
	return self->clone();
}
bool QFocusEvent_GotFocus(const QFocusEvent* self) {
	return self->gotFocus();
}
bool QFocusEvent_LostFocus(const QFocusEvent* self) {
	return self->lostFocus();
}
int QFocusEvent_Reason(const QFocusEvent* self) {
	Qt::FocusReason _ret = self->reason();
	return static_cast<int>(_ret);
}
void QFocusEvent_Delete(QFocusEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFocusEvent*>( self );
	} else {
		delete self;
	}
}
QPaintEvent* QPaintEvent_new(QRegion* paintRegion) {
return new QPaintEvent(*paintRegion);
}
QPaintEvent* QPaintEvent_new2(QRect* paintRect) {
return new QPaintEvent(*paintRect);
}
void QPaintEvent_virtbase(QPaintEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QPaintEvent* QPaintEvent_Clone(const QPaintEvent* self) {
	return self->clone();
}
QRect* QPaintEvent_Rect(const QPaintEvent* self) {
	const QRect& _ret = self->rect();
	// Cast returned reference into pointer
	return const_cast<QRect*>(&_ret);
}
QRegion* QPaintEvent_Region(const QPaintEvent* self) {
	const QRegion& _ret = self->region();
	// Cast returned reference into pointer
	return const_cast<QRegion*>(&_ret);
}
void QPaintEvent_Delete(QPaintEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPaintEvent*>( self );
	} else {
		delete self;
	}
}
QMoveEvent* QMoveEvent_new(QPoint* pos, QPoint* oldPos) {
return new QMoveEvent(*pos, *oldPos);
}
void QMoveEvent_virtbase(QMoveEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QMoveEvent* QMoveEvent_Clone(const QMoveEvent* self) {
	return self->clone();
}
QPoint* QMoveEvent_Pos(const QMoveEvent* self) {
	const QPoint& _ret = self->pos();
	// Cast returned reference into pointer
	return const_cast<QPoint*>(&_ret);
}
QPoint* QMoveEvent_OldPos(const QMoveEvent* self) {
	const QPoint& _ret = self->oldPos();
	// Cast returned reference into pointer
	return const_cast<QPoint*>(&_ret);
}
void QMoveEvent_Delete(QMoveEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QMoveEvent*>( self );
	} else {
		delete self;
	}
}
QExposeEvent* QExposeEvent_new(QRegion* m_region) {
return new QExposeEvent(*m_region);
}
void QExposeEvent_virtbase(QExposeEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QExposeEvent* QExposeEvent_Clone(const QExposeEvent* self) {
	return self->clone();
}
QRegion* QExposeEvent_Region(const QExposeEvent* self) {
	const QRegion& _ret = self->region();
	// Cast returned reference into pointer
	return const_cast<QRegion*>(&_ret);
}
void QExposeEvent_Delete(QExposeEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QExposeEvent*>( self );
	} else {
		delete self;
	}
}
QPlatformSurfaceEvent* QPlatformSurfaceEvent_new(SurfaceEventType surfaceEventType) {
return new QPlatformSurfaceEvent(surfaceEventType);
}
void QPlatformSurfaceEvent_virtbase(QPlatformSurfaceEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QPlatformSurfaceEvent* QPlatformSurfaceEvent_Clone(const QPlatformSurfaceEvent* self) {
	return self->clone();
}
SurfaceEventType QPlatformSurfaceEvent_SurfaceEventType(const QPlatformSurfaceEvent* self) {
	return self->surfaceEventType();
}
void QPlatformSurfaceEvent_Delete(QPlatformSurfaceEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPlatformSurfaceEvent*>( self );
	} else {
		delete self;
	}
}
QResizeEvent* QResizeEvent_new(QSize* size, QSize* oldSize) {
return new QResizeEvent(*size, *oldSize);
}
void QResizeEvent_virtbase(QResizeEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QResizeEvent* QResizeEvent_Clone(const QResizeEvent* self) {
	return self->clone();
}
QSize* QResizeEvent_Size(const QResizeEvent* self) {
	const QSize& _ret = self->size();
	// Cast returned reference into pointer
	return const_cast<QSize*>(&_ret);
}
QSize* QResizeEvent_OldSize(const QResizeEvent* self) {
	const QSize& _ret = self->oldSize();
	// Cast returned reference into pointer
	return const_cast<QSize*>(&_ret);
}
void QResizeEvent_Delete(QResizeEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QResizeEvent*>( self );
	} else {
		delete self;
	}
}
QCloseEvent* QCloseEvent_new() {
return new QCloseEvent();
}
void QCloseEvent_virtbase(QCloseEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QCloseEvent* QCloseEvent_Clone(const QCloseEvent* self) {
	return self->clone();
}
void QCloseEvent_Delete(QCloseEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCloseEvent*>( self );
	} else {
		delete self;
	}
}
QIconDragEvent* QIconDragEvent_new() {
return new QIconDragEvent();
}
void QIconDragEvent_virtbase(QIconDragEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QIconDragEvent* QIconDragEvent_Clone(const QIconDragEvent* self) {
	return self->clone();
}
void QIconDragEvent_Delete(QIconDragEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QIconDragEvent*>( self );
	} else {
		delete self;
	}
}
QShowEvent* QShowEvent_new() {
return new QShowEvent();
}
void QShowEvent_virtbase(QShowEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QShowEvent* QShowEvent_Clone(const QShowEvent* self) {
	return self->clone();
}
void QShowEvent_Delete(QShowEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QShowEvent*>( self );
	} else {
		delete self;
	}
}
QHideEvent* QHideEvent_new() {
return new QHideEvent();
}
void QHideEvent_virtbase(QHideEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QHideEvent* QHideEvent_Clone(const QHideEvent* self) {
	return self->clone();
}
void QHideEvent_Delete(QHideEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QHideEvent*>( self );
	} else {
		delete self;
	}
}
QContextMenuEvent* QContextMenuEvent_new(Reason reason, QPoint* pos, QPoint* globalPos) {
return new QContextMenuEvent(reason, *pos, *globalPos);
}
QContextMenuEvent* QContextMenuEvent_new2(Reason reason, QPoint* pos) {
return new QContextMenuEvent(reason, *pos);
}
QContextMenuEvent* QContextMenuEvent_new3(Reason reason, QPoint* pos, QPoint* globalPos, int modifiers) {
return new QContextMenuEvent(reason, *pos, *globalPos, static_cast<Qt::KeyboardModifiers>(modifiers));
}
void QContextMenuEvent_virtbase(QContextMenuEvent* src
, QInputEvent** outptr_QInputEvent
) {
*outptr_QInputEvent = static_cast<QInputEvent*>(src);
}
QContextMenuEvent* QContextMenuEvent_Clone(const QContextMenuEvent* self) {
	return self->clone();
}
int QContextMenuEvent_X(const QContextMenuEvent* self) {
	return self->x();
}
int QContextMenuEvent_Y(const QContextMenuEvent* self) {
	return self->y();
}
int QContextMenuEvent_GlobalX(const QContextMenuEvent* self) {
	return self->globalX();
}
int QContextMenuEvent_GlobalY(const QContextMenuEvent* self) {
	return self->globalY();
}
QPoint* QContextMenuEvent_Pos(const QContextMenuEvent* self) {
	const QPoint& _ret = self->pos();
	// Cast returned reference into pointer
	return const_cast<QPoint*>(&_ret);
}
QPoint* QContextMenuEvent_GlobalPos(const QContextMenuEvent* self) {
	const QPoint& _ret = self->globalPos();
	// Cast returned reference into pointer
	return const_cast<QPoint*>(&_ret);
}
Reason QContextMenuEvent_Reason(const QContextMenuEvent* self) {
	return self->reason();
}
void QContextMenuEvent_Delete(QContextMenuEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QContextMenuEvent*>( self );
	} else {
		delete self;
	}
}
QInputMethodEvent* QInputMethodEvent_new() {
return new QInputMethodEvent();
}
QInputMethodEvent* QInputMethodEvent_new2(struct miqt_string preeditText, struct miqt_array /* of Attribute */  attributes) {
QString preeditText_QString = QString::fromUtf8(preeditText.data, preeditText.len);
	QList<Attribute> attributes_QList;
	attributes_QList.reserve(attributes.len);
	Attribute* attributes_arr = static_cast<Attribute*>(attributes.data);
	for(size_t i = 0; i < attributes.len; ++i) {
		attributes_QList.push_back(attributes_arr[i]);
	}
	return new QInputMethodEvent(preeditText_QString, attributes_QList);
}
void QInputMethodEvent_virtbase(QInputMethodEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QInputMethodEvent* QInputMethodEvent_Clone(const QInputMethodEvent* self) {
	return self->clone();
}
void QInputMethodEvent_SetCommitString(QInputMethodEvent* self, struct miqt_string commitString) {
	QString commitString_QString = QString::fromUtf8(commitString.data, commitString.len);
	self->setCommitString(commitString_QString);
}
struct miqt_array /* of Attribute */  QInputMethodEvent_Attributes(const QInputMethodEvent* self) {
	const QList<Attribute>& _ret = self->attributes();
	// Convert QList<> from C++ memory to manually-managed C memory
	Attribute* _arr = static_cast<Attribute*>(malloc(sizeof(Attribute) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_string QInputMethodEvent_PreeditString(const QInputMethodEvent* self) {
	const QString _ret = self->preeditString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QInputMethodEvent_CommitString(const QInputMethodEvent* self) {
	const QString _ret = self->commitString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QInputMethodEvent_ReplacementStart(const QInputMethodEvent* self) {
	return self->replacementStart();
}
int QInputMethodEvent_ReplacementLength(const QInputMethodEvent* self) {
	return self->replacementLength();
}
void QInputMethodEvent_SetCommitString2(QInputMethodEvent* self, struct miqt_string commitString, int replaceFrom) {
	QString commitString_QString = QString::fromUtf8(commitString.data, commitString.len);
	self->setCommitString(commitString_QString, static_cast<int>(replaceFrom));
}
void QInputMethodEvent_SetCommitString3(QInputMethodEvent* self, struct miqt_string commitString, int replaceFrom, int replaceLength) {
	QString commitString_QString = QString::fromUtf8(commitString.data, commitString.len);
	self->setCommitString(commitString_QString, static_cast<int>(replaceFrom), static_cast<int>(replaceLength));
}
void QInputMethodEvent_Delete(QInputMethodEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QInputMethodEvent*>( self );
	} else {
		delete self;
	}
}
QInputMethodQueryEvent* QInputMethodQueryEvent_new(int queries) {
return new QInputMethodQueryEvent(static_cast<Qt::InputMethodQueries>(queries));
}
void QInputMethodQueryEvent_virtbase(QInputMethodQueryEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QInputMethodQueryEvent* QInputMethodQueryEvent_Clone(const QInputMethodQueryEvent* self) {
	return self->clone();
}
int QInputMethodQueryEvent_Queries(const QInputMethodQueryEvent* self) {
	Qt::InputMethodQueries _ret = self->queries();
	return static_cast<int>(_ret);
}
void QInputMethodQueryEvent_SetValue(QInputMethodQueryEvent* self, int query, QVariant* value) {
	self->setValue(static_cast<Qt::InputMethodQuery>(query), *value);
}
QVariant* QInputMethodQueryEvent_Value(const QInputMethodQueryEvent* self, int query) {
	return new QVariant(self->value(static_cast<Qt::InputMethodQuery>(query)));
}
void QInputMethodQueryEvent_Delete(QInputMethodQueryEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QInputMethodQueryEvent*>( self );
	} else {
		delete self;
	}
}
QDropEvent* QDropEvent_new(QPointF* pos, int actions, QMimeData* data, int buttons, int modifiers) {
return new QDropEvent(*pos, static_cast<Qt::DropActions>(actions), data, static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers));
}
QDropEvent* QDropEvent_new2(QPointF* pos, int actions, QMimeData* data, int buttons, int modifiers, Type typeVal) {
return new QDropEvent(*pos, static_cast<Qt::DropActions>(actions), data, static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), typeVal);
}
void QDropEvent_virtbase(QDropEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QDropEvent* QDropEvent_Clone(const QDropEvent* self) {
	return self->clone();
}
QPoint* QDropEvent_Pos(const QDropEvent* self) {
	return new QPoint(self->pos());
}
QPointF* QDropEvent_PosF(const QDropEvent* self) {
	return new QPointF(self->posF());
}
int QDropEvent_MouseButtons(const QDropEvent* self) {
	Qt::MouseButtons _ret = self->mouseButtons();
	return static_cast<int>(_ret);
}
int QDropEvent_KeyboardModifiers(const QDropEvent* self) {
	Qt::KeyboardModifiers _ret = self->keyboardModifiers();
	return static_cast<int>(_ret);
}
QPointF* QDropEvent_Position(const QDropEvent* self) {
	return new QPointF(self->position());
}
int QDropEvent_Buttons(const QDropEvent* self) {
	Qt::MouseButtons _ret = self->buttons();
	return static_cast<int>(_ret);
}
int QDropEvent_Modifiers(const QDropEvent* self) {
	Qt::KeyboardModifiers _ret = self->modifiers();
	return static_cast<int>(_ret);
}
int QDropEvent_PossibleActions(const QDropEvent* self) {
	Qt::DropActions _ret = self->possibleActions();
	return static_cast<int>(_ret);
}
int QDropEvent_ProposedAction(const QDropEvent* self) {
	Qt::DropAction _ret = self->proposedAction();
	return static_cast<int>(_ret);
}
void QDropEvent_AcceptProposedAction(QDropEvent* self) {
	self->acceptProposedAction();
}
int QDropEvent_DropAction(const QDropEvent* self) {
	Qt::DropAction _ret = self->dropAction();
	return static_cast<int>(_ret);
}
void QDropEvent_SetDropAction(QDropEvent* self, int action) {
	self->setDropAction(static_cast<Qt::DropAction>(action));
}
QObject* QDropEvent_Source(const QDropEvent* self) {
	return self->source();
}
QMimeData* QDropEvent_MimeData(const QDropEvent* self) {
	return (QMimeData*) self->mimeData();
}
void QDropEvent_Delete(QDropEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDropEvent*>( self );
	} else {
		delete self;
	}
}
QDragMoveEvent* QDragMoveEvent_new(QPoint* pos, int actions, QMimeData* data, int buttons, int modifiers) {
return new QDragMoveEvent(*pos, static_cast<Qt::DropActions>(actions), data, static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers));
}
QDragMoveEvent* QDragMoveEvent_new2(QPoint* pos, int actions, QMimeData* data, int buttons, int modifiers, Type typeVal) {
return new QDragMoveEvent(*pos, static_cast<Qt::DropActions>(actions), data, static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers), typeVal);
}
void QDragMoveEvent_virtbase(QDragMoveEvent* src
, QDropEvent** outptr_QDropEvent
) {
*outptr_QDropEvent = static_cast<QDropEvent*>(src);
}
QDragMoveEvent* QDragMoveEvent_Clone(const QDragMoveEvent* self) {
	return self->clone();
}
QRect* QDragMoveEvent_AnswerRect(const QDragMoveEvent* self) {
	return new QRect(self->answerRect());
}
void QDragMoveEvent_Accept(QDragMoveEvent* self) {
	self->accept();
}
void QDragMoveEvent_Ignore(QDragMoveEvent* self) {
	self->ignore();
}
void QDragMoveEvent_AcceptWithQRect(QDragMoveEvent* self, QRect* r) {
	self->accept(*r);
}
void QDragMoveEvent_IgnoreWithQRect(QDragMoveEvent* self, QRect* r) {
	self->ignore(*r);
}
void QDragMoveEvent_Delete(QDragMoveEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDragMoveEvent*>( self );
	} else {
		delete self;
	}
}
QDragEnterEvent* QDragEnterEvent_new(QPoint* pos, int actions, QMimeData* data, int buttons, int modifiers) {
return new QDragEnterEvent(*pos, static_cast<Qt::DropActions>(actions), data, static_cast<Qt::MouseButtons>(buttons), static_cast<Qt::KeyboardModifiers>(modifiers));
}
void QDragEnterEvent_virtbase(QDragEnterEvent* src
, QDragMoveEvent** outptr_QDragMoveEvent
) {
*outptr_QDragMoveEvent = static_cast<QDragMoveEvent*>(src);
}
QDragEnterEvent* QDragEnterEvent_Clone(const QDragEnterEvent* self) {
	return self->clone();
}
void QDragEnterEvent_Delete(QDragEnterEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDragEnterEvent*>( self );
	} else {
		delete self;
	}
}
QDragLeaveEvent* QDragLeaveEvent_new() {
return new QDragLeaveEvent();
}
void QDragLeaveEvent_virtbase(QDragLeaveEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QDragLeaveEvent* QDragLeaveEvent_Clone(const QDragLeaveEvent* self) {
	return self->clone();
}
void QDragLeaveEvent_Delete(QDragLeaveEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDragLeaveEvent*>( self );
	} else {
		delete self;
	}
}
QHelpEvent* QHelpEvent_new(Type typeVal, QPoint* pos, QPoint* globalPos) {
return new QHelpEvent(typeVal, *pos, *globalPos);
}
void QHelpEvent_virtbase(QHelpEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QHelpEvent* QHelpEvent_Clone(const QHelpEvent* self) {
	return self->clone();
}
int QHelpEvent_X(const QHelpEvent* self) {
	return self->x();
}
int QHelpEvent_Y(const QHelpEvent* self) {
	return self->y();
}
int QHelpEvent_GlobalX(const QHelpEvent* self) {
	return self->globalX();
}
int QHelpEvent_GlobalY(const QHelpEvent* self) {
	return self->globalY();
}
QPoint* QHelpEvent_Pos(const QHelpEvent* self) {
	const QPoint& _ret = self->pos();
	// Cast returned reference into pointer
	return const_cast<QPoint*>(&_ret);
}
QPoint* QHelpEvent_GlobalPos(const QHelpEvent* self) {
	const QPoint& _ret = self->globalPos();
	// Cast returned reference into pointer
	return const_cast<QPoint*>(&_ret);
}
void QHelpEvent_Delete(QHelpEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QHelpEvent*>( self );
	} else {
		delete self;
	}
}
QStatusTipEvent* QStatusTipEvent_new(struct miqt_string tip) {
QString tip_QString = QString::fromUtf8(tip.data, tip.len);
	return new QStatusTipEvent(tip_QString);
}
void QStatusTipEvent_virtbase(QStatusTipEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QStatusTipEvent* QStatusTipEvent_Clone(const QStatusTipEvent* self) {
	return self->clone();
}
struct miqt_string QStatusTipEvent_Tip(const QStatusTipEvent* self) {
	QString _ret = self->tip();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStatusTipEvent_Delete(QStatusTipEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QStatusTipEvent*>( self );
	} else {
		delete self;
	}
}
QWhatsThisClickedEvent* QWhatsThisClickedEvent_new(struct miqt_string href) {
QString href_QString = QString::fromUtf8(href.data, href.len);
	return new QWhatsThisClickedEvent(href_QString);
}
void QWhatsThisClickedEvent_virtbase(QWhatsThisClickedEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QWhatsThisClickedEvent* QWhatsThisClickedEvent_Clone(const QWhatsThisClickedEvent* self) {
	return self->clone();
}
struct miqt_string QWhatsThisClickedEvent_Href(const QWhatsThisClickedEvent* self) {
	QString _ret = self->href();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWhatsThisClickedEvent_Delete(QWhatsThisClickedEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QWhatsThisClickedEvent*>( self );
	} else {
		delete self;
	}
}
QActionEvent* QActionEvent_new(int typeVal, QAction* action) {
return new QActionEvent(static_cast<int>(typeVal), action);
}
QActionEvent* QActionEvent_new2(int typeVal, QAction* action, QAction* before) {
return new QActionEvent(static_cast<int>(typeVal), action, before);
}
void QActionEvent_virtbase(QActionEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QActionEvent* QActionEvent_Clone(const QActionEvent* self) {
	return self->clone();
}
QAction* QActionEvent_Action(const QActionEvent* self) {
	return self->action();
}
QAction* QActionEvent_Before(const QActionEvent* self) {
	return self->before();
}
void QActionEvent_Delete(QActionEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QActionEvent*>( self );
	} else {
		delete self;
	}
}
QFileOpenEvent* QFileOpenEvent_new(struct miqt_string file) {
QString file_QString = QString::fromUtf8(file.data, file.len);
	return new QFileOpenEvent(file_QString);
}
QFileOpenEvent* QFileOpenEvent_new2(QUrl* url) {
return new QFileOpenEvent(*url);
}
void QFileOpenEvent_virtbase(QFileOpenEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QFileOpenEvent* QFileOpenEvent_Clone(const QFileOpenEvent* self) {
	return self->clone();
}
struct miqt_string QFileOpenEvent_File(const QFileOpenEvent* self) {
	QString _ret = self->file();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QUrl* QFileOpenEvent_Url(const QFileOpenEvent* self) {
	return new QUrl(self->url());
}
bool QFileOpenEvent_OpenFile(const QFileOpenEvent* self, QFile* file, int flags) {
	return self->openFile(*file, static_cast<QIODevice::OpenMode>(flags));
}
void QFileOpenEvent_Delete(QFileOpenEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFileOpenEvent*>( self );
	} else {
		delete self;
	}
}
QToolBarChangeEvent* QToolBarChangeEvent_new(bool t) {
return new QToolBarChangeEvent(t);
}
void QToolBarChangeEvent_virtbase(QToolBarChangeEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QToolBarChangeEvent* QToolBarChangeEvent_Clone(const QToolBarChangeEvent* self) {
	return self->clone();
}
bool QToolBarChangeEvent_Toggle(const QToolBarChangeEvent* self) {
	return self->toggle();
}
void QToolBarChangeEvent_Delete(QToolBarChangeEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QToolBarChangeEvent*>( self );
	} else {
		delete self;
	}
}
QShortcutEvent* QShortcutEvent_new(QKeySequence* key, int id) {
return new QShortcutEvent(*key, static_cast<int>(id));
}
QShortcutEvent* QShortcutEvent_new2(QKeySequence* key) {
return new QShortcutEvent(*key);
}
QShortcutEvent* QShortcutEvent_new3(QKeySequence* key, int id, bool ambiguous) {
return new QShortcutEvent(*key, static_cast<int>(id), ambiguous);
}
QShortcutEvent* QShortcutEvent_new4(QKeySequence* key, QShortcut* shortcut) {
return new QShortcutEvent(*key, shortcut);
}
QShortcutEvent* QShortcutEvent_new5(QKeySequence* key, QShortcut* shortcut, bool ambiguous) {
return new QShortcutEvent(*key, shortcut, ambiguous);
}
void QShortcutEvent_virtbase(QShortcutEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QShortcutEvent* QShortcutEvent_Clone(const QShortcutEvent* self) {
	return self->clone();
}
QKeySequence* QShortcutEvent_Key(const QShortcutEvent* self) {
	const QKeySequence& _ret = self->key();
	// Cast returned reference into pointer
	return const_cast<QKeySequence*>(&_ret);
}
int QShortcutEvent_ShortcutId(const QShortcutEvent* self) {
	return self->shortcutId();
}
bool QShortcutEvent_IsAmbiguous(const QShortcutEvent* self) {
	return self->isAmbiguous();
}
void QShortcutEvent_Delete(QShortcutEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QShortcutEvent*>( self );
	} else {
		delete self;
	}
}
QWindowStateChangeEvent* QWindowStateChangeEvent_new(int oldState) {
return new QWindowStateChangeEvent(static_cast<Qt::WindowStates>(oldState));
}
QWindowStateChangeEvent* QWindowStateChangeEvent_new2(int oldState, bool isOverride) {
return new QWindowStateChangeEvent(static_cast<Qt::WindowStates>(oldState), isOverride);
}
void QWindowStateChangeEvent_virtbase(QWindowStateChangeEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QWindowStateChangeEvent* QWindowStateChangeEvent_Clone(const QWindowStateChangeEvent* self) {
	return self->clone();
}
int QWindowStateChangeEvent_OldState(const QWindowStateChangeEvent* self) {
	Qt::WindowStates _ret = self->oldState();
	return static_cast<int>(_ret);
}
bool QWindowStateChangeEvent_IsOverride(const QWindowStateChangeEvent* self) {
	return self->isOverride();
}
void QWindowStateChangeEvent_Delete(QWindowStateChangeEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QWindowStateChangeEvent*>( self );
	} else {
		delete self;
	}
}
QTouchEvent* QTouchEvent_new(int eventType) {
return new QTouchEvent(static_cast<QEvent::Type>(eventType));
}
QTouchEvent* QTouchEvent_new2(int eventType, QPointingDevice* device, int modifiers, int touchPointStates) {
return new QTouchEvent(static_cast<QEvent::Type>(eventType), device, static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<QEventPoint::States>(touchPointStates));
}
QTouchEvent* QTouchEvent_new3(int eventType, QPointingDevice* device) {
return new QTouchEvent(static_cast<QEvent::Type>(eventType), device);
}
QTouchEvent* QTouchEvent_new4(int eventType, QPointingDevice* device, int modifiers) {
return new QTouchEvent(static_cast<QEvent::Type>(eventType), device, static_cast<Qt::KeyboardModifiers>(modifiers));
}
QTouchEvent* QTouchEvent_new5(int eventType, QPointingDevice* device, int modifiers, struct miqt_array /* of QEventPoint* */  touchPoints) {
QList<QEventPoint> touchPoints_QList;
	touchPoints_QList.reserve(touchPoints.len);
	QEventPoint** touchPoints_arr = static_cast<QEventPoint**>(touchPoints.data);
	for(size_t i = 0; i < touchPoints.len; ++i) {
		touchPoints_QList.push_back(*(touchPoints_arr[i]));
	}
	return new QTouchEvent(static_cast<QEvent::Type>(eventType), device, static_cast<Qt::KeyboardModifiers>(modifiers), touchPoints_QList);
}
QTouchEvent* QTouchEvent_new6(int eventType, QPointingDevice* device, int modifiers, int touchPointStates, struct miqt_array /* of QEventPoint* */  touchPoints) {
QList<QEventPoint> touchPoints_QList;
	touchPoints_QList.reserve(touchPoints.len);
	QEventPoint** touchPoints_arr = static_cast<QEventPoint**>(touchPoints.data);
	for(size_t i = 0; i < touchPoints.len; ++i) {
		touchPoints_QList.push_back(*(touchPoints_arr[i]));
	}
	return new QTouchEvent(static_cast<QEvent::Type>(eventType), device, static_cast<Qt::KeyboardModifiers>(modifiers), static_cast<QEventPoint::States>(touchPointStates), touchPoints_QList);
}
void QTouchEvent_virtbase(QTouchEvent* src
, QPointerEvent** outptr_QPointerEvent
) {
*outptr_QPointerEvent = static_cast<QPointerEvent*>(src);
}
QTouchEvent* QTouchEvent_Clone(const QTouchEvent* self) {
	return self->clone();
}
QObject* QTouchEvent_Target(const QTouchEvent* self) {
	return self->target();
}
int QTouchEvent_TouchPointStates(const QTouchEvent* self) {
	QEventPoint::States _ret = self->touchPointStates();
	return static_cast<int>(_ret);
}
struct miqt_array /* of QEventPoint* */  QTouchEvent_TouchPoints(const QTouchEvent* self) {
	const QList<QEventPoint>& _ret = self->touchPoints();
	// Convert QList<> from C++ memory to manually-managed C memory
	QEventPoint** _arr = static_cast<QEventPoint**>(malloc(sizeof(QEventPoint*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QEventPoint(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
bool QTouchEvent_IsBeginEvent(const QTouchEvent* self) {
	return self->isBeginEvent();
}
bool QTouchEvent_IsUpdateEvent(const QTouchEvent* self) {
	return self->isUpdateEvent();
}
bool QTouchEvent_IsEndEvent(const QTouchEvent* self) {
	return self->isEndEvent();
}
void QTouchEvent_Delete(QTouchEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTouchEvent*>( self );
	} else {
		delete self;
	}
}
QScrollPrepareEvent* QScrollPrepareEvent_new(QPointF* startPos) {
return new QScrollPrepareEvent(*startPos);
}
void QScrollPrepareEvent_virtbase(QScrollPrepareEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QScrollPrepareEvent* QScrollPrepareEvent_Clone(const QScrollPrepareEvent* self) {
	return self->clone();
}
QPointF* QScrollPrepareEvent_StartPos(const QScrollPrepareEvent* self) {
	return new QPointF(self->startPos());
}
QSizeF* QScrollPrepareEvent_ViewportSize(const QScrollPrepareEvent* self) {
	return new QSizeF(self->viewportSize());
}
QRectF* QScrollPrepareEvent_ContentPosRange(const QScrollPrepareEvent* self) {
	return new QRectF(self->contentPosRange());
}
QPointF* QScrollPrepareEvent_ContentPos(const QScrollPrepareEvent* self) {
	return new QPointF(self->contentPos());
}
void QScrollPrepareEvent_SetViewportSize(QScrollPrepareEvent* self, QSizeF* size) {
	self->setViewportSize(*size);
}
void QScrollPrepareEvent_SetContentPosRange(QScrollPrepareEvent* self, QRectF* rect) {
	self->setContentPosRange(*rect);
}
void QScrollPrepareEvent_SetContentPos(QScrollPrepareEvent* self, QPointF* pos) {
	self->setContentPos(*pos);
}
void QScrollPrepareEvent_Delete(QScrollPrepareEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QScrollPrepareEvent*>( self );
	} else {
		delete self;
	}
}
QScrollEvent* QScrollEvent_new(QPointF* contentPos, QPointF* overshoot, ScrollState scrollState) {
return new QScrollEvent(*contentPos, *overshoot, scrollState);
}
void QScrollEvent_virtbase(QScrollEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QScrollEvent* QScrollEvent_Clone(const QScrollEvent* self) {
	return self->clone();
}
QPointF* QScrollEvent_ContentPos(const QScrollEvent* self) {
	return new QPointF(self->contentPos());
}
QPointF* QScrollEvent_OvershootDistance(const QScrollEvent* self) {
	return new QPointF(self->overshootDistance());
}
ScrollState QScrollEvent_ScrollState(const QScrollEvent* self) {
	return self->scrollState();
}
void QScrollEvent_Delete(QScrollEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QScrollEvent*>( self );
	} else {
		delete self;
	}
}
QScreenOrientationChangeEvent* QScreenOrientationChangeEvent_new(QScreen* screen, int orientation) {
return new QScreenOrientationChangeEvent(screen, static_cast<Qt::ScreenOrientation>(orientation));
}
void QScreenOrientationChangeEvent_virtbase(QScreenOrientationChangeEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QScreenOrientationChangeEvent* QScreenOrientationChangeEvent_Clone(const QScreenOrientationChangeEvent* self) {
	return self->clone();
}
QScreen* QScreenOrientationChangeEvent_Screen(const QScreenOrientationChangeEvent* self) {
	return self->screen();
}
int QScreenOrientationChangeEvent_Orientation(const QScreenOrientationChangeEvent* self) {
	Qt::ScreenOrientation _ret = self->orientation();
	return static_cast<int>(_ret);
}
void QScreenOrientationChangeEvent_Delete(QScreenOrientationChangeEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QScreenOrientationChangeEvent*>( self );
	} else {
		delete self;
	}
}
QApplicationStateChangeEvent* QApplicationStateChangeEvent_new(int state) {
return new QApplicationStateChangeEvent(static_cast<Qt::ApplicationState>(state));
}
void QApplicationStateChangeEvent_virtbase(QApplicationStateChangeEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QApplicationStateChangeEvent* QApplicationStateChangeEvent_Clone(const QApplicationStateChangeEvent* self) {
	return self->clone();
}
int QApplicationStateChangeEvent_ApplicationState(const QApplicationStateChangeEvent* self) {
	Qt::ApplicationState _ret = self->applicationState();
	return static_cast<int>(_ret);
}
void QApplicationStateChangeEvent_Delete(QApplicationStateChangeEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QApplicationStateChangeEvent*>( self );
	} else {
		delete self;
	}
}
QChildWindowEvent* QChildWindowEvent_new(Type typeVal, QWindow* childWindow) {
return new QChildWindowEvent(typeVal, childWindow);
}
void QChildWindowEvent_virtbase(QChildWindowEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
QChildWindowEvent* QChildWindowEvent_Clone(const QChildWindowEvent* self) {
	return self->clone();
}
QWindow* QChildWindowEvent_Child(const QChildWindowEvent* self) {
	return self->child();
}
void QChildWindowEvent_Delete(QChildWindowEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QChildWindowEvent*>( self );
	} else {
		delete self;
	}
}
QInputMethodEvent__Attribute* QInputMethodEvent__Attribute_new(AttributeType typ, int s, int l, QVariant* val) {
return new QInputMethodEvent::Attribute(typ, static_cast<int>(s), static_cast<int>(l), *val);
}
QInputMethodEvent__Attribute* QInputMethodEvent__Attribute_new2(AttributeType typ, int s, int l) {
return new QInputMethodEvent::Attribute(typ, static_cast<int>(s), static_cast<int>(l));
}
void QInputMethodEvent__Attribute_OperatorAssign(QInputMethodEvent__Attribute* self, const Attribute* param1) {
	self->operator=(*param1);
}
void QInputMethodEvent__Attribute_Delete(QInputMethodEvent__Attribute* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QInputMethodEvent::Attribute*>( self );
	} else {
		delete self;
	}
}
