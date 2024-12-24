// +build ignore

#include <QActionEvent>
#include <QByteArray>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEnterEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QImage>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QResizeEvent>
#include <QRhiWidget>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTabletEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <qrhiwidget.h>
#include "gen_qrhiwidget.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQRhiWidget : public virtual QRhiWidget {
public:

	MiqtVirtualQRhiWidget(QWidget* parent): QRhiWidget(parent) {};
	MiqtVirtualQRhiWidget(): QRhiWidget() {};
	MiqtVirtualQRhiWidget(QWidget* parent, Qt::WindowFlags f): QRhiWidget(parent, f) {};

	virtual ~MiqtVirtualQRhiWidget() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Initialize = 0;

	// Subclass to allow providing a Go implementation
	virtual void initialize(QRhiCommandBuffer* cb) override {
		if (handle__Initialize == 0) {
			QRhiWidget::initialize(cb);
			return;
		}
		
		QRhiCommandBuffer* sigval1 = cb;

		miqt_exec_callback_QRhiWidget_Initialize(this, handle__Initialize, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Initialize(QRhiCommandBuffer* cb) {

		QRhiWidget::initialize(cb);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Render = 0;

	// Subclass to allow providing a Go implementation
	virtual void render(QRhiCommandBuffer* cb) override {
		if (handle__Render == 0) {
			QRhiWidget::render(cb);
			return;
		}
		
		QRhiCommandBuffer* sigval1 = cb;

		miqt_exec_callback_QRhiWidget_Render(this, handle__Render, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Render(QRhiCommandBuffer* cb) {

		QRhiWidget::render(cb);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ReleaseResources = 0;

	// Subclass to allow providing a Go implementation
	virtual void releaseResources() override {
		if (handle__ReleaseResources == 0) {
			QRhiWidget::releaseResources();
			return;
		}
		

		miqt_exec_callback_QRhiWidget_ReleaseResources(this, handle__ReleaseResources);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ReleaseResources() {

		QRhiWidget::releaseResources();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ResizeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void resizeEvent(QResizeEvent* e) override {
		if (handle__ResizeEvent == 0) {
			QRhiWidget::resizeEvent(e);
			return;
		}
		
		QResizeEvent* sigval1 = e;

		miqt_exec_callback_QRhiWidget_ResizeEvent(this, handle__ResizeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ResizeEvent(QResizeEvent* e) {

		QRhiWidget::resizeEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void paintEvent(QPaintEvent* e) override {
		if (handle__PaintEvent == 0) {
			QRhiWidget::paintEvent(e);
			return;
		}
		
		QPaintEvent* sigval1 = e;

		miqt_exec_callback_QRhiWidget_PaintEvent(this, handle__PaintEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_PaintEvent(QPaintEvent* e) {

		QRhiWidget::paintEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* e) override {
		if (handle__Event == 0) {
			return QRhiWidget::event(e);
		}
		
		QEvent* sigval1 = e;

		bool callback_return_value = miqt_exec_callback_QRhiWidget_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* e) {

		return QRhiWidget::event(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DevType = 0;

	// Subclass to allow providing a Go implementation
	virtual int devType() const override {
		if (handle__DevType == 0) {
			return QRhiWidget::devType();
		}
		

		int callback_return_value = miqt_exec_callback_QRhiWidget_DevType(const_cast<MiqtVirtualQRhiWidget*>(this), handle__DevType);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_DevType() const {

		return QRhiWidget::devType();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetVisible = 0;

	// Subclass to allow providing a Go implementation
	virtual void setVisible(bool visible) override {
		if (handle__SetVisible == 0) {
			QRhiWidget::setVisible(visible);
			return;
		}
		
		bool sigval1 = visible;

		miqt_exec_callback_QRhiWidget_SetVisible(this, handle__SetVisible, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetVisible(bool visible) {

		QRhiWidget::setVisible(visible);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__SizeHint == 0) {
			return QRhiWidget::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QRhiWidget_SizeHint(const_cast<MiqtVirtualQRhiWidget*>(this), handle__SizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeHint() const {

		return new QSize(QRhiWidget::sizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumSizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSizeHint() const override {
		if (handle__MinimumSizeHint == 0) {
			return QRhiWidget::minimumSizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QRhiWidget_MinimumSizeHint(const_cast<MiqtVirtualQRhiWidget*>(this), handle__MinimumSizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MinimumSizeHint() const {

		return new QSize(QRhiWidget::minimumSizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int heightForWidth(int param1) const override {
		if (handle__HeightForWidth == 0) {
			return QRhiWidget::heightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QRhiWidget_HeightForWidth(const_cast<MiqtVirtualQRhiWidget*>(this), handle__HeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_HeightForWidth(int param1) const {

		return QRhiWidget::heightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HasHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasHeightForWidth() const override {
		if (handle__HasHeightForWidth == 0) {
			return QRhiWidget::hasHeightForWidth();
		}
		

		bool callback_return_value = miqt_exec_callback_QRhiWidget_HasHeightForWidth(const_cast<MiqtVirtualQRhiWidget*>(this), handle__HasHeightForWidth);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HasHeightForWidth() const {

		return QRhiWidget::hasHeightForWidth();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEngine = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintEngine* paintEngine() const override {
		if (handle__PaintEngine == 0) {
			return QRhiWidget::paintEngine();
		}
		

		QPaintEngine* callback_return_value = miqt_exec_callback_QRhiWidget_PaintEngine(const_cast<MiqtVirtualQRhiWidget*>(this), handle__PaintEngine);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintEngine* virtualbase_PaintEngine() const {

		return QRhiWidget::paintEngine();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MousePressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mousePressEvent(QMouseEvent* event) override {
		if (handle__MousePressEvent == 0) {
			QRhiWidget::mousePressEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_MousePressEvent(this, handle__MousePressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MousePressEvent(QMouseEvent* event) {

		QRhiWidget::mousePressEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseReleaseEvent(QMouseEvent* event) override {
		if (handle__MouseReleaseEvent == 0) {
			QRhiWidget::mouseReleaseEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_MouseReleaseEvent(this, handle__MouseReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseReleaseEvent(QMouseEvent* event) {

		QRhiWidget::mouseReleaseEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseDoubleClickEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseDoubleClickEvent(QMouseEvent* event) override {
		if (handle__MouseDoubleClickEvent == 0) {
			QRhiWidget::mouseDoubleClickEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_MouseDoubleClickEvent(this, handle__MouseDoubleClickEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseDoubleClickEvent(QMouseEvent* event) {

		QRhiWidget::mouseDoubleClickEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseMoveEvent(QMouseEvent* event) override {
		if (handle__MouseMoveEvent == 0) {
			QRhiWidget::mouseMoveEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_MouseMoveEvent(this, handle__MouseMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseMoveEvent(QMouseEvent* event) {

		QRhiWidget::mouseMoveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WheelEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void wheelEvent(QWheelEvent* event) override {
		if (handle__WheelEvent == 0) {
			QRhiWidget::wheelEvent(event);
			return;
		}
		
		QWheelEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_WheelEvent(this, handle__WheelEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_WheelEvent(QWheelEvent* event) {

		QRhiWidget::wheelEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyPressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyPressEvent(QKeyEvent* event) override {
		if (handle__KeyPressEvent == 0) {
			QRhiWidget::keyPressEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_KeyPressEvent(this, handle__KeyPressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyPressEvent(QKeyEvent* event) {

		QRhiWidget::keyPressEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyReleaseEvent(QKeyEvent* event) override {
		if (handle__KeyReleaseEvent == 0) {
			QRhiWidget::keyReleaseEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_KeyReleaseEvent(this, handle__KeyReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyReleaseEvent(QKeyEvent* event) {

		QRhiWidget::keyReleaseEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusInEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusInEvent(QFocusEvent* event) override {
		if (handle__FocusInEvent == 0) {
			QRhiWidget::focusInEvent(event);
			return;
		}
		
		QFocusEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_FocusInEvent(this, handle__FocusInEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusInEvent(QFocusEvent* event) {

		QRhiWidget::focusInEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusOutEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusOutEvent(QFocusEvent* event) override {
		if (handle__FocusOutEvent == 0) {
			QRhiWidget::focusOutEvent(event);
			return;
		}
		
		QFocusEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_FocusOutEvent(this, handle__FocusOutEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusOutEvent(QFocusEvent* event) {

		QRhiWidget::focusOutEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void enterEvent(QEnterEvent* event) override {
		if (handle__EnterEvent == 0) {
			QRhiWidget::enterEvent(event);
			return;
		}
		
		QEnterEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_EnterEvent(this, handle__EnterEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_EnterEvent(QEnterEvent* event) {

		QRhiWidget::enterEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__LeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void leaveEvent(QEvent* event) override {
		if (handle__LeaveEvent == 0) {
			QRhiWidget::leaveEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_LeaveEvent(this, handle__LeaveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_LeaveEvent(QEvent* event) {

		QRhiWidget::leaveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void moveEvent(QMoveEvent* event) override {
		if (handle__MoveEvent == 0) {
			QRhiWidget::moveEvent(event);
			return;
		}
		
		QMoveEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_MoveEvent(this, handle__MoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MoveEvent(QMoveEvent* event) {

		QRhiWidget::moveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CloseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void closeEvent(QCloseEvent* event) override {
		if (handle__CloseEvent == 0) {
			QRhiWidget::closeEvent(event);
			return;
		}
		
		QCloseEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_CloseEvent(this, handle__CloseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CloseEvent(QCloseEvent* event) {

		QRhiWidget::closeEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ContextMenuEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void contextMenuEvent(QContextMenuEvent* event) override {
		if (handle__ContextMenuEvent == 0) {
			QRhiWidget::contextMenuEvent(event);
			return;
		}
		
		QContextMenuEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_ContextMenuEvent(this, handle__ContextMenuEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ContextMenuEvent(QContextMenuEvent* event) {

		QRhiWidget::contextMenuEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TabletEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void tabletEvent(QTabletEvent* event) override {
		if (handle__TabletEvent == 0) {
			QRhiWidget::tabletEvent(event);
			return;
		}
		
		QTabletEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_TabletEvent(this, handle__TabletEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TabletEvent(QTabletEvent* event) {

		QRhiWidget::tabletEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ActionEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void actionEvent(QActionEvent* event) override {
		if (handle__ActionEvent == 0) {
			QRhiWidget::actionEvent(event);
			return;
		}
		
		QActionEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_ActionEvent(this, handle__ActionEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ActionEvent(QActionEvent* event) {

		QRhiWidget::actionEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragEnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragEnterEvent(QDragEnterEvent* event) override {
		if (handle__DragEnterEvent == 0) {
			QRhiWidget::dragEnterEvent(event);
			return;
		}
		
		QDragEnterEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_DragEnterEvent(this, handle__DragEnterEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragEnterEvent(QDragEnterEvent* event) {

		QRhiWidget::dragEnterEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragMoveEvent(QDragMoveEvent* event) override {
		if (handle__DragMoveEvent == 0) {
			QRhiWidget::dragMoveEvent(event);
			return;
		}
		
		QDragMoveEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_DragMoveEvent(this, handle__DragMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragMoveEvent(QDragMoveEvent* event) {

		QRhiWidget::dragMoveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragLeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragLeaveEvent(QDragLeaveEvent* event) override {
		if (handle__DragLeaveEvent == 0) {
			QRhiWidget::dragLeaveEvent(event);
			return;
		}
		
		QDragLeaveEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_DragLeaveEvent(this, handle__DragLeaveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragLeaveEvent(QDragLeaveEvent* event) {

		QRhiWidget::dragLeaveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DropEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dropEvent(QDropEvent* event) override {
		if (handle__DropEvent == 0) {
			QRhiWidget::dropEvent(event);
			return;
		}
		
		QDropEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_DropEvent(this, handle__DropEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DropEvent(QDropEvent* event) {

		QRhiWidget::dropEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ShowEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void showEvent(QShowEvent* event) override {
		if (handle__ShowEvent == 0) {
			QRhiWidget::showEvent(event);
			return;
		}
		
		QShowEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_ShowEvent(this, handle__ShowEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ShowEvent(QShowEvent* event) {

		QRhiWidget::showEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HideEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void hideEvent(QHideEvent* event) override {
		if (handle__HideEvent == 0) {
			QRhiWidget::hideEvent(event);
			return;
		}
		
		QHideEvent* sigval1 = event;

		miqt_exec_callback_QRhiWidget_HideEvent(this, handle__HideEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_HideEvent(QHideEvent* event) {

		QRhiWidget::hideEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__NativeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		if (handle__NativeEvent == 0) {
			return QRhiWidget::nativeEvent(eventType, message, result);
		}
		
		const QByteArray eventType_qb = eventType;
		struct miqt_string eventType_ms;
		eventType_ms.len = eventType_qb.length();
		eventType_ms.data = static_cast<char*>(malloc(eventType_ms.len));
		memcpy(eventType_ms.data, eventType_qb.data(), eventType_ms.len);
		struct miqt_string sigval1 = eventType_ms;
		void* sigval2 = message;
		qintptr* result_ret = result;
		intptr_t* sigval3 = (intptr_t*)(result_ret);

		bool callback_return_value = miqt_exec_callback_QRhiWidget_NativeEvent(this, handle__NativeEvent, sigval1, sigval2, sigval3);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_NativeEvent(struct miqt_string eventType, void* message, intptr_t* result) {
		QByteArray eventType_QByteArray(eventType.data, eventType.len);

		return QRhiWidget::nativeEvent(eventType_QByteArray, message, (qintptr*)(result));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChangeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void changeEvent(QEvent* param1) override {
		if (handle__ChangeEvent == 0) {
			QRhiWidget::changeEvent(param1);
			return;
		}
		
		QEvent* sigval1 = param1;

		miqt_exec_callback_QRhiWidget_ChangeEvent(this, handle__ChangeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChangeEvent(QEvent* param1) {

		QRhiWidget::changeEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metric = 0;

	// Subclass to allow providing a Go implementation
	virtual int metric(PaintDeviceMetric param1) const override {
		if (handle__Metric == 0) {
			return QRhiWidget::metric(param1);
		}
		
		PaintDeviceMetric sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QRhiWidget_Metric(const_cast<MiqtVirtualQRhiWidget*>(this), handle__Metric, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Metric(PaintDeviceMetric param1) const {

		return QRhiWidget::metric(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InitPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual void initPainter(QPainter* painter) const override {
		if (handle__InitPainter == 0) {
			QRhiWidget::initPainter(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QRhiWidget_InitPainter(const_cast<MiqtVirtualQRhiWidget*>(this), handle__InitPainter, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_InitPainter(QPainter* painter) const {

		QRhiWidget::initPainter(painter);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Redirected = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintDevice* redirected(QPoint* offset) const override {
		if (handle__Redirected == 0) {
			return QRhiWidget::redirected(offset);
		}
		
		QPoint* sigval1 = offset;

		QPaintDevice* callback_return_value = miqt_exec_callback_QRhiWidget_Redirected(const_cast<MiqtVirtualQRhiWidget*>(this), handle__Redirected, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintDevice* virtualbase_Redirected(QPoint* offset) const {

		return QRhiWidget::redirected(offset);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SharedPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual QPainter* sharedPainter() const override {
		if (handle__SharedPainter == 0) {
			return QRhiWidget::sharedPainter();
		}
		

		QPainter* callback_return_value = miqt_exec_callback_QRhiWidget_SharedPainter(const_cast<MiqtVirtualQRhiWidget*>(this), handle__SharedPainter);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPainter* virtualbase_SharedPainter() const {

		return QRhiWidget::sharedPainter();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InputMethodEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void inputMethodEvent(QInputMethodEvent* param1) override {
		if (handle__InputMethodEvent == 0) {
			QRhiWidget::inputMethodEvent(param1);
			return;
		}
		
		QInputMethodEvent* sigval1 = param1;

		miqt_exec_callback_QRhiWidget_InputMethodEvent(this, handle__InputMethodEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_InputMethodEvent(QInputMethodEvent* param1) {

		QRhiWidget::inputMethodEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InputMethodQuery = 0;

	// Subclass to allow providing a Go implementation
	virtual QVariant inputMethodQuery(Qt::InputMethodQuery param1) const override {
		if (handle__InputMethodQuery == 0) {
			return QRhiWidget::inputMethodQuery(param1);
		}
		
		Qt::InputMethodQuery param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		QVariant* callback_return_value = miqt_exec_callback_QRhiWidget_InputMethodQuery(const_cast<MiqtVirtualQRhiWidget*>(this), handle__InputMethodQuery, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QVariant* virtualbase_InputMethodQuery(int param1) const {

		return new QVariant(QRhiWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(param1)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusNextPrevChild = 0;

	// Subclass to allow providing a Go implementation
	virtual bool focusNextPrevChild(bool next) override {
		if (handle__FocusNextPrevChild == 0) {
			return QRhiWidget::focusNextPrevChild(next);
		}
		
		bool sigval1 = next;

		bool callback_return_value = miqt_exec_callback_QRhiWidget_FocusNextPrevChild(this, handle__FocusNextPrevChild, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_FocusNextPrevChild(bool next) {

		return QRhiWidget::focusNextPrevChild(next);

	}

};

QRhiWidget* QRhiWidget_new(QWidget* parent) {
	return new MiqtVirtualQRhiWidget(parent);
}

QRhiWidget* QRhiWidget_new2() {
	return new MiqtVirtualQRhiWidget();
}

QRhiWidget* QRhiWidget_new3(QWidget* parent, int f) {
	return new MiqtVirtualQRhiWidget(parent, static_cast<Qt::WindowFlags>(f));
}

void QRhiWidget_virtbase(QRhiWidget* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QRhiWidget_MetaObject(const QRhiWidget* self) {
	return (QMetaObject*) self->metaObject();
}

void* QRhiWidget_Metacast(QRhiWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QRhiWidget_Tr(const char* s) {
	QString _ret = QRhiWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

Api QRhiWidget_Api(const QRhiWidget* self) {
	return self->api();
}

void QRhiWidget_SetApi(QRhiWidget* self, Api api) {
	self->setApi(api);
}

bool QRhiWidget_IsDebugLayerEnabled(const QRhiWidget* self) {
	return self->isDebugLayerEnabled();
}

void QRhiWidget_SetDebugLayerEnabled(QRhiWidget* self, bool enable) {
	self->setDebugLayerEnabled(enable);
}

int QRhiWidget_SampleCount(const QRhiWidget* self) {
	return self->sampleCount();
}

void QRhiWidget_SetSampleCount(QRhiWidget* self, int samples) {
	self->setSampleCount(static_cast<int>(samples));
}

TextureFormat QRhiWidget_ColorBufferFormat(const QRhiWidget* self) {
	return self->colorBufferFormat();
}

void QRhiWidget_SetColorBufferFormat(QRhiWidget* self, TextureFormat format) {
	self->setColorBufferFormat(format);
}

QSize* QRhiWidget_FixedColorBufferSize(const QRhiWidget* self) {
	return new QSize(self->fixedColorBufferSize());
}

void QRhiWidget_SetFixedColorBufferSize(QRhiWidget* self, QSize* pixelSize) {
	self->setFixedColorBufferSize(*pixelSize);
}

void QRhiWidget_SetFixedColorBufferSize2(QRhiWidget* self, int w, int h) {
	self->setFixedColorBufferSize(static_cast<int>(w), static_cast<int>(h));
}

bool QRhiWidget_IsMirrorVerticallyEnabled(const QRhiWidget* self) {
	return self->isMirrorVerticallyEnabled();
}

void QRhiWidget_SetMirrorVertically(QRhiWidget* self, bool enabled) {
	self->setMirrorVertically(enabled);
}

QImage* QRhiWidget_GrabFramebuffer(const QRhiWidget* self) {
	return new QImage(self->grabFramebuffer());
}

void QRhiWidget_FrameSubmitted(QRhiWidget* self) {
	self->frameSubmitted();
}

void QRhiWidget_connect_FrameSubmitted(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)()>(&QRhiWidget::frameSubmitted), self, [=]() {
		miqt_exec_callback_QRhiWidget_FrameSubmitted(slot);
	});
}

void QRhiWidget_RenderFailed(QRhiWidget* self) {
	self->renderFailed();
}

void QRhiWidget_connect_RenderFailed(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)()>(&QRhiWidget::renderFailed), self, [=]() {
		miqt_exec_callback_QRhiWidget_RenderFailed(slot);
	});
}

void QRhiWidget_SampleCountChanged(QRhiWidget* self, int samples) {
	self->sampleCountChanged(static_cast<int>(samples));
}

void QRhiWidget_connect_SampleCountChanged(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)(int)>(&QRhiWidget::sampleCountChanged), self, [=](int samples) {
		int sigval1 = samples;
		miqt_exec_callback_QRhiWidget_SampleCountChanged(slot, sigval1);
	});
}

void QRhiWidget_ColorBufferFormatChanged(QRhiWidget* self, TextureFormat format) {
	self->colorBufferFormatChanged(format);
}

void QRhiWidget_connect_ColorBufferFormatChanged(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)(TextureFormat)>(&QRhiWidget::colorBufferFormatChanged), self, [=](TextureFormat format) {
		TextureFormat sigval1 = format;
		miqt_exec_callback_QRhiWidget_ColorBufferFormatChanged(slot, sigval1);
	});
}

void QRhiWidget_FixedColorBufferSizeChanged(QRhiWidget* self, QSize* pixelSize) {
	self->fixedColorBufferSizeChanged(*pixelSize);
}

void QRhiWidget_connect_FixedColorBufferSizeChanged(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)(const QSize&)>(&QRhiWidget::fixedColorBufferSizeChanged), self, [=](const QSize& pixelSize) {
		const QSize& pixelSize_ret = pixelSize;
		// Cast returned reference into pointer
		QSize* sigval1 = const_cast<QSize*>(&pixelSize_ret);
		miqt_exec_callback_QRhiWidget_FixedColorBufferSizeChanged(slot, sigval1);
	});
}

void QRhiWidget_MirrorVerticallyChanged(QRhiWidget* self, bool enabled) {
	self->mirrorVerticallyChanged(enabled);
}

void QRhiWidget_connect_MirrorVerticallyChanged(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)(bool)>(&QRhiWidget::mirrorVerticallyChanged), self, [=](bool enabled) {
		bool sigval1 = enabled;
		miqt_exec_callback_QRhiWidget_MirrorVerticallyChanged(slot, sigval1);
	});
}

struct miqt_string QRhiWidget_Tr2(const char* s, const char* c) {
	QString _ret = QRhiWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRhiWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRhiWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QRhiWidget_override_virtual_Initialize(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__Initialize = slot;
}

void QRhiWidget_virtualbase_Initialize(void* self, QRhiCommandBuffer* cb) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_Initialize(cb);
}

void QRhiWidget_override_virtual_Render(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__Render = slot;
}

void QRhiWidget_virtualbase_Render(void* self, QRhiCommandBuffer* cb) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_Render(cb);
}

void QRhiWidget_override_virtual_ReleaseResources(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__ReleaseResources = slot;
}

void QRhiWidget_virtualbase_ReleaseResources(void* self) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_ReleaseResources();
}

void QRhiWidget_override_virtual_ResizeEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__ResizeEvent = slot;
}

void QRhiWidget_virtualbase_ResizeEvent(void* self, QResizeEvent* e) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_ResizeEvent(e);
}

void QRhiWidget_override_virtual_PaintEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__PaintEvent = slot;
}

void QRhiWidget_virtualbase_PaintEvent(void* self, QPaintEvent* e) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_PaintEvent(e);
}

void QRhiWidget_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__Event = slot;
}

bool QRhiWidget_virtualbase_Event(void* self, QEvent* e) {
	return ( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_Event(e);
}

void QRhiWidget_override_virtual_DevType(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__DevType = slot;
}

int QRhiWidget_virtualbase_DevType(const void* self) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_DevType();
}

void QRhiWidget_override_virtual_SetVisible(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__SetVisible = slot;
}

void QRhiWidget_virtualbase_SetVisible(void* self, bool visible) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_SetVisible(visible);
}

void QRhiWidget_override_virtual_SizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__SizeHint = slot;
}

QSize* QRhiWidget_virtualbase_SizeHint(const void* self) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_SizeHint();
}

void QRhiWidget_override_virtual_MinimumSizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__MinimumSizeHint = slot;
}

QSize* QRhiWidget_virtualbase_MinimumSizeHint(const void* self) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_MinimumSizeHint();
}

void QRhiWidget_override_virtual_HeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__HeightForWidth = slot;
}

int QRhiWidget_virtualbase_HeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_HeightForWidth(param1);
}

void QRhiWidget_override_virtual_HasHeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__HasHeightForWidth = slot;
}

bool QRhiWidget_virtualbase_HasHeightForWidth(const void* self) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_HasHeightForWidth();
}

void QRhiWidget_override_virtual_PaintEngine(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__PaintEngine = slot;
}

QPaintEngine* QRhiWidget_virtualbase_PaintEngine(const void* self) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_PaintEngine();
}

void QRhiWidget_override_virtual_MousePressEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__MousePressEvent = slot;
}

void QRhiWidget_virtualbase_MousePressEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_MousePressEvent(event);
}

void QRhiWidget_override_virtual_MouseReleaseEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__MouseReleaseEvent = slot;
}

void QRhiWidget_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_MouseReleaseEvent(event);
}

void QRhiWidget_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__MouseDoubleClickEvent = slot;
}

void QRhiWidget_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_MouseDoubleClickEvent(event);
}

void QRhiWidget_override_virtual_MouseMoveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__MouseMoveEvent = slot;
}

void QRhiWidget_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_MouseMoveEvent(event);
}

void QRhiWidget_override_virtual_WheelEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__WheelEvent = slot;
}

void QRhiWidget_virtualbase_WheelEvent(void* self, QWheelEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_WheelEvent(event);
}

void QRhiWidget_override_virtual_KeyPressEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__KeyPressEvent = slot;
}

void QRhiWidget_virtualbase_KeyPressEvent(void* self, QKeyEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_KeyPressEvent(event);
}

void QRhiWidget_override_virtual_KeyReleaseEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__KeyReleaseEvent = slot;
}

void QRhiWidget_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_KeyReleaseEvent(event);
}

void QRhiWidget_override_virtual_FocusInEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__FocusInEvent = slot;
}

void QRhiWidget_virtualbase_FocusInEvent(void* self, QFocusEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_FocusInEvent(event);
}

void QRhiWidget_override_virtual_FocusOutEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__FocusOutEvent = slot;
}

void QRhiWidget_virtualbase_FocusOutEvent(void* self, QFocusEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_FocusOutEvent(event);
}

void QRhiWidget_override_virtual_EnterEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__EnterEvent = slot;
}

void QRhiWidget_virtualbase_EnterEvent(void* self, QEnterEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_EnterEvent(event);
}

void QRhiWidget_override_virtual_LeaveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__LeaveEvent = slot;
}

void QRhiWidget_virtualbase_LeaveEvent(void* self, QEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_LeaveEvent(event);
}

void QRhiWidget_override_virtual_MoveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__MoveEvent = slot;
}

void QRhiWidget_virtualbase_MoveEvent(void* self, QMoveEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_MoveEvent(event);
}

void QRhiWidget_override_virtual_CloseEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__CloseEvent = slot;
}

void QRhiWidget_virtualbase_CloseEvent(void* self, QCloseEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_CloseEvent(event);
}

void QRhiWidget_override_virtual_ContextMenuEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__ContextMenuEvent = slot;
}

void QRhiWidget_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_ContextMenuEvent(event);
}

void QRhiWidget_override_virtual_TabletEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__TabletEvent = slot;
}

void QRhiWidget_virtualbase_TabletEvent(void* self, QTabletEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_TabletEvent(event);
}

void QRhiWidget_override_virtual_ActionEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__ActionEvent = slot;
}

void QRhiWidget_virtualbase_ActionEvent(void* self, QActionEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_ActionEvent(event);
}

void QRhiWidget_override_virtual_DragEnterEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__DragEnterEvent = slot;
}

void QRhiWidget_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_DragEnterEvent(event);
}

void QRhiWidget_override_virtual_DragMoveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__DragMoveEvent = slot;
}

void QRhiWidget_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_DragMoveEvent(event);
}

void QRhiWidget_override_virtual_DragLeaveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__DragLeaveEvent = slot;
}

void QRhiWidget_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_DragLeaveEvent(event);
}

void QRhiWidget_override_virtual_DropEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__DropEvent = slot;
}

void QRhiWidget_virtualbase_DropEvent(void* self, QDropEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_DropEvent(event);
}

void QRhiWidget_override_virtual_ShowEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__ShowEvent = slot;
}

void QRhiWidget_virtualbase_ShowEvent(void* self, QShowEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_ShowEvent(event);
}

void QRhiWidget_override_virtual_HideEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__HideEvent = slot;
}

void QRhiWidget_virtualbase_HideEvent(void* self, QHideEvent* event) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_HideEvent(event);
}

void QRhiWidget_override_virtual_NativeEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__NativeEvent = slot;
}

bool QRhiWidget_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result) {
	return ( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_NativeEvent(eventType, message, result);
}

void QRhiWidget_override_virtual_ChangeEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__ChangeEvent = slot;
}

void QRhiWidget_virtualbase_ChangeEvent(void* self, QEvent* param1) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_ChangeEvent(param1);
}

void QRhiWidget_override_virtual_Metric(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__Metric = slot;
}

int QRhiWidget_virtualbase_Metric(const void* self, PaintDeviceMetric param1) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_Metric(param1);
}

void QRhiWidget_override_virtual_InitPainter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__InitPainter = slot;
}

void QRhiWidget_virtualbase_InitPainter(const void* self, QPainter* painter) {
	( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_InitPainter(painter);
}

void QRhiWidget_override_virtual_Redirected(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__Redirected = slot;
}

QPaintDevice* QRhiWidget_virtualbase_Redirected(const void* self, QPoint* offset) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_Redirected(offset);
}

void QRhiWidget_override_virtual_SharedPainter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__SharedPainter = slot;
}

QPainter* QRhiWidget_virtualbase_SharedPainter(const void* self) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_SharedPainter();
}

void QRhiWidget_override_virtual_InputMethodEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__InputMethodEvent = slot;
}

void QRhiWidget_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1) {
	( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_InputMethodEvent(param1);
}

void QRhiWidget_override_virtual_InputMethodQuery(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__InputMethodQuery = slot;
}

QVariant* QRhiWidget_virtualbase_InputMethodQuery(const void* self, int param1) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_InputMethodQuery(param1);
}

void QRhiWidget_override_virtual_FocusNextPrevChild(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__FocusNextPrevChild = slot;
}

bool QRhiWidget_virtualbase_FocusNextPrevChild(void* self, bool next) {
	return ( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_FocusNextPrevChild(next);
}

void QRhiWidget_Delete(QRhiWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRhiWidget*>( self );
	} else {
		delete self;
	}
}

