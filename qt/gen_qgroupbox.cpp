#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QGroupBox>
#include <QHideEvent>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
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
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <qgroupbox.h>
#include "gen_qgroupbox.h"

#ifdef __cplusplus
extern "C" {
#endif

void miqt_exec_callback_QGroupBox_Clicked(intptr_t);
void miqt_exec_callback_QGroupBox_Toggled(intptr_t, bool);
void miqt_exec_callback_QGroupBox_Clicked1(intptr_t, bool);
QSize* miqt_exec_callback_QGroupBox_MinimumSizeHint(const QGroupBox*, intptr_t);
bool miqt_exec_callback_QGroupBox_Event(QGroupBox*, intptr_t, QEvent*);
void miqt_exec_callback_QGroupBox_ChildEvent(QGroupBox*, intptr_t, QChildEvent*);
void miqt_exec_callback_QGroupBox_ResizeEvent(QGroupBox*, intptr_t, QResizeEvent*);
void miqt_exec_callback_QGroupBox_PaintEvent(QGroupBox*, intptr_t, QPaintEvent*);
void miqt_exec_callback_QGroupBox_FocusInEvent(QGroupBox*, intptr_t, QFocusEvent*);
void miqt_exec_callback_QGroupBox_ChangeEvent(QGroupBox*, intptr_t, QEvent*);
void miqt_exec_callback_QGroupBox_MousePressEvent(QGroupBox*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QGroupBox_MouseMoveEvent(QGroupBox*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QGroupBox_MouseReleaseEvent(QGroupBox*, intptr_t, QMouseEvent*);
int miqt_exec_callback_QGroupBox_DevType(const QGroupBox*, intptr_t);
void miqt_exec_callback_QGroupBox_SetVisible(QGroupBox*, intptr_t, bool);
QSize* miqt_exec_callback_QGroupBox_SizeHint(const QGroupBox*, intptr_t);
int miqt_exec_callback_QGroupBox_HeightForWidth(const QGroupBox*, intptr_t, int);
bool miqt_exec_callback_QGroupBox_HasHeightForWidth(const QGroupBox*, intptr_t);
QPaintEngine* miqt_exec_callback_QGroupBox_PaintEngine(const QGroupBox*, intptr_t);
void miqt_exec_callback_QGroupBox_MouseDoubleClickEvent(QGroupBox*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QGroupBox_WheelEvent(QGroupBox*, intptr_t, QWheelEvent*);
void miqt_exec_callback_QGroupBox_KeyPressEvent(QGroupBox*, intptr_t, QKeyEvent*);
void miqt_exec_callback_QGroupBox_KeyReleaseEvent(QGroupBox*, intptr_t, QKeyEvent*);
void miqt_exec_callback_QGroupBox_FocusOutEvent(QGroupBox*, intptr_t, QFocusEvent*);
void miqt_exec_callback_QGroupBox_EnterEvent(QGroupBox*, intptr_t, QEvent*);
void miqt_exec_callback_QGroupBox_LeaveEvent(QGroupBox*, intptr_t, QEvent*);
void miqt_exec_callback_QGroupBox_MoveEvent(QGroupBox*, intptr_t, QMoveEvent*);
void miqt_exec_callback_QGroupBox_CloseEvent(QGroupBox*, intptr_t, QCloseEvent*);
void miqt_exec_callback_QGroupBox_ContextMenuEvent(QGroupBox*, intptr_t, QContextMenuEvent*);
void miqt_exec_callback_QGroupBox_TabletEvent(QGroupBox*, intptr_t, QTabletEvent*);
void miqt_exec_callback_QGroupBox_ActionEvent(QGroupBox*, intptr_t, QActionEvent*);
void miqt_exec_callback_QGroupBox_DragEnterEvent(QGroupBox*, intptr_t, QDragEnterEvent*);
void miqt_exec_callback_QGroupBox_DragMoveEvent(QGroupBox*, intptr_t, QDragMoveEvent*);
void miqt_exec_callback_QGroupBox_DragLeaveEvent(QGroupBox*, intptr_t, QDragLeaveEvent*);
void miqt_exec_callback_QGroupBox_DropEvent(QGroupBox*, intptr_t, QDropEvent*);
void miqt_exec_callback_QGroupBox_ShowEvent(QGroupBox*, intptr_t, QShowEvent*);
void miqt_exec_callback_QGroupBox_HideEvent(QGroupBox*, intptr_t, QHideEvent*);
bool miqt_exec_callback_QGroupBox_NativeEvent(QGroupBox*, intptr_t, struct miqt_string, void*, long*);
int miqt_exec_callback_QGroupBox_Metric(const QGroupBox*, intptr_t, int);
void miqt_exec_callback_QGroupBox_InitPainter(const QGroupBox*, intptr_t, QPainter*);
QPaintDevice* miqt_exec_callback_QGroupBox_Redirected(const QGroupBox*, intptr_t, QPoint*);
QPainter* miqt_exec_callback_QGroupBox_SharedPainter(const QGroupBox*, intptr_t);
void miqt_exec_callback_QGroupBox_InputMethodEvent(QGroupBox*, intptr_t, QInputMethodEvent*);
QVariant* miqt_exec_callback_QGroupBox_InputMethodQuery(const QGroupBox*, intptr_t, int);
bool miqt_exec_callback_QGroupBox_FocusNextPrevChild(QGroupBox*, intptr_t, bool);
bool miqt_exec_callback_QGroupBox_EventFilter(QGroupBox*, intptr_t, QObject*, QEvent*);
void miqt_exec_callback_QGroupBox_TimerEvent(QGroupBox*, intptr_t, QTimerEvent*);
void miqt_exec_callback_QGroupBox_CustomEvent(QGroupBox*, intptr_t, QEvent*);
void miqt_exec_callback_QGroupBox_ConnectNotify(QGroupBox*, intptr_t, QMetaMethod*);
void miqt_exec_callback_QGroupBox_DisconnectNotify(QGroupBox*, intptr_t, QMetaMethod*);
#ifdef __cplusplus
} /* extern C */
#endif

class MiqtVirtualQGroupBox final : public QGroupBox {
public:

	MiqtVirtualQGroupBox(QWidget* parent): QGroupBox(parent) {};
	MiqtVirtualQGroupBox(): QGroupBox() {};
	MiqtVirtualQGroupBox(const QString& title): QGroupBox(title) {};
	MiqtVirtualQGroupBox(const QString& title, QWidget* parent): QGroupBox(title, parent) {};

	virtual ~MiqtVirtualQGroupBox() override = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumSizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSizeHint() const override {
		if (handle__MinimumSizeHint == 0) {
			return QGroupBox::minimumSizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QGroupBox_MinimumSizeHint(this, handle__MinimumSizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MinimumSizeHint() const {

		return new QSize(QGroupBox::minimumSizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__Event == 0) {
			return QGroupBox::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QGroupBox_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* event) {

		return QGroupBox::event(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChildEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__ChildEvent == 0) {
			QGroupBox::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_ChildEvent(this, handle__ChildEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChildEvent(QChildEvent* event) {

		QGroupBox::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ResizeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void resizeEvent(QResizeEvent* event) override {
		if (handle__ResizeEvent == 0) {
			QGroupBox::resizeEvent(event);
			return;
		}
		
		QResizeEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_ResizeEvent(this, handle__ResizeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ResizeEvent(QResizeEvent* event) {

		QGroupBox::resizeEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void paintEvent(QPaintEvent* event) override {
		if (handle__PaintEvent == 0) {
			QGroupBox::paintEvent(event);
			return;
		}
		
		QPaintEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_PaintEvent(this, handle__PaintEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_PaintEvent(QPaintEvent* event) {

		QGroupBox::paintEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusInEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusInEvent(QFocusEvent* event) override {
		if (handle__FocusInEvent == 0) {
			QGroupBox::focusInEvent(event);
			return;
		}
		
		QFocusEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_FocusInEvent(this, handle__FocusInEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusInEvent(QFocusEvent* event) {

		QGroupBox::focusInEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChangeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void changeEvent(QEvent* event) override {
		if (handle__ChangeEvent == 0) {
			QGroupBox::changeEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_ChangeEvent(this, handle__ChangeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChangeEvent(QEvent* event) {

		QGroupBox::changeEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MousePressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mousePressEvent(QMouseEvent* event) override {
		if (handle__MousePressEvent == 0) {
			QGroupBox::mousePressEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_MousePressEvent(this, handle__MousePressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MousePressEvent(QMouseEvent* event) {

		QGroupBox::mousePressEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseMoveEvent(QMouseEvent* event) override {
		if (handle__MouseMoveEvent == 0) {
			QGroupBox::mouseMoveEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_MouseMoveEvent(this, handle__MouseMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseMoveEvent(QMouseEvent* event) {

		QGroupBox::mouseMoveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseReleaseEvent(QMouseEvent* event) override {
		if (handle__MouseReleaseEvent == 0) {
			QGroupBox::mouseReleaseEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_MouseReleaseEvent(this, handle__MouseReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseReleaseEvent(QMouseEvent* event) {

		QGroupBox::mouseReleaseEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DevType = 0;

	// Subclass to allow providing a Go implementation
	virtual int devType() const override {
		if (handle__DevType == 0) {
			return QGroupBox::devType();
		}
		

		int callback_return_value = miqt_exec_callback_QGroupBox_DevType(this, handle__DevType);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_DevType() const {

		return QGroupBox::devType();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetVisible = 0;

	// Subclass to allow providing a Go implementation
	virtual void setVisible(bool visible) override {
		if (handle__SetVisible == 0) {
			QGroupBox::setVisible(visible);
			return;
		}
		
		bool sigval1 = visible;

		miqt_exec_callback_QGroupBox_SetVisible(this, handle__SetVisible, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetVisible(bool visible) {

		QGroupBox::setVisible(visible);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__SizeHint == 0) {
			return QGroupBox::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QGroupBox_SizeHint(this, handle__SizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeHint() const {

		return new QSize(QGroupBox::sizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int heightForWidth(int param1) const override {
		if (handle__HeightForWidth == 0) {
			return QGroupBox::heightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QGroupBox_HeightForWidth(this, handle__HeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_HeightForWidth(int param1) const {

		return QGroupBox::heightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HasHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasHeightForWidth() const override {
		if (handle__HasHeightForWidth == 0) {
			return QGroupBox::hasHeightForWidth();
		}
		

		bool callback_return_value = miqt_exec_callback_QGroupBox_HasHeightForWidth(this, handle__HasHeightForWidth);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HasHeightForWidth() const {

		return QGroupBox::hasHeightForWidth();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEngine = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintEngine* paintEngine() const override {
		if (handle__PaintEngine == 0) {
			return QGroupBox::paintEngine();
		}
		

		QPaintEngine* callback_return_value = miqt_exec_callback_QGroupBox_PaintEngine(this, handle__PaintEngine);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintEngine* virtualbase_PaintEngine() const {

		return QGroupBox::paintEngine();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseDoubleClickEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseDoubleClickEvent(QMouseEvent* event) override {
		if (handle__MouseDoubleClickEvent == 0) {
			QGroupBox::mouseDoubleClickEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_MouseDoubleClickEvent(this, handle__MouseDoubleClickEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseDoubleClickEvent(QMouseEvent* event) {

		QGroupBox::mouseDoubleClickEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WheelEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void wheelEvent(QWheelEvent* event) override {
		if (handle__WheelEvent == 0) {
			QGroupBox::wheelEvent(event);
			return;
		}
		
		QWheelEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_WheelEvent(this, handle__WheelEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_WheelEvent(QWheelEvent* event) {

		QGroupBox::wheelEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyPressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyPressEvent(QKeyEvent* event) override {
		if (handle__KeyPressEvent == 0) {
			QGroupBox::keyPressEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_KeyPressEvent(this, handle__KeyPressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyPressEvent(QKeyEvent* event) {

		QGroupBox::keyPressEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyReleaseEvent(QKeyEvent* event) override {
		if (handle__KeyReleaseEvent == 0) {
			QGroupBox::keyReleaseEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_KeyReleaseEvent(this, handle__KeyReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyReleaseEvent(QKeyEvent* event) {

		QGroupBox::keyReleaseEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusOutEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusOutEvent(QFocusEvent* event) override {
		if (handle__FocusOutEvent == 0) {
			QGroupBox::focusOutEvent(event);
			return;
		}
		
		QFocusEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_FocusOutEvent(this, handle__FocusOutEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusOutEvent(QFocusEvent* event) {

		QGroupBox::focusOutEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void enterEvent(QEvent* event) override {
		if (handle__EnterEvent == 0) {
			QGroupBox::enterEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_EnterEvent(this, handle__EnterEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_EnterEvent(QEvent* event) {

		QGroupBox::enterEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__LeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void leaveEvent(QEvent* event) override {
		if (handle__LeaveEvent == 0) {
			QGroupBox::leaveEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_LeaveEvent(this, handle__LeaveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_LeaveEvent(QEvent* event) {

		QGroupBox::leaveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void moveEvent(QMoveEvent* event) override {
		if (handle__MoveEvent == 0) {
			QGroupBox::moveEvent(event);
			return;
		}
		
		QMoveEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_MoveEvent(this, handle__MoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MoveEvent(QMoveEvent* event) {

		QGroupBox::moveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CloseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void closeEvent(QCloseEvent* event) override {
		if (handle__CloseEvent == 0) {
			QGroupBox::closeEvent(event);
			return;
		}
		
		QCloseEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_CloseEvent(this, handle__CloseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CloseEvent(QCloseEvent* event) {

		QGroupBox::closeEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ContextMenuEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void contextMenuEvent(QContextMenuEvent* event) override {
		if (handle__ContextMenuEvent == 0) {
			QGroupBox::contextMenuEvent(event);
			return;
		}
		
		QContextMenuEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_ContextMenuEvent(this, handle__ContextMenuEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ContextMenuEvent(QContextMenuEvent* event) {

		QGroupBox::contextMenuEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TabletEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void tabletEvent(QTabletEvent* event) override {
		if (handle__TabletEvent == 0) {
			QGroupBox::tabletEvent(event);
			return;
		}
		
		QTabletEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_TabletEvent(this, handle__TabletEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TabletEvent(QTabletEvent* event) {

		QGroupBox::tabletEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ActionEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void actionEvent(QActionEvent* event) override {
		if (handle__ActionEvent == 0) {
			QGroupBox::actionEvent(event);
			return;
		}
		
		QActionEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_ActionEvent(this, handle__ActionEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ActionEvent(QActionEvent* event) {

		QGroupBox::actionEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragEnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragEnterEvent(QDragEnterEvent* event) override {
		if (handle__DragEnterEvent == 0) {
			QGroupBox::dragEnterEvent(event);
			return;
		}
		
		QDragEnterEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_DragEnterEvent(this, handle__DragEnterEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragEnterEvent(QDragEnterEvent* event) {

		QGroupBox::dragEnterEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragMoveEvent(QDragMoveEvent* event) override {
		if (handle__DragMoveEvent == 0) {
			QGroupBox::dragMoveEvent(event);
			return;
		}
		
		QDragMoveEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_DragMoveEvent(this, handle__DragMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragMoveEvent(QDragMoveEvent* event) {

		QGroupBox::dragMoveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragLeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragLeaveEvent(QDragLeaveEvent* event) override {
		if (handle__DragLeaveEvent == 0) {
			QGroupBox::dragLeaveEvent(event);
			return;
		}
		
		QDragLeaveEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_DragLeaveEvent(this, handle__DragLeaveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragLeaveEvent(QDragLeaveEvent* event) {

		QGroupBox::dragLeaveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DropEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dropEvent(QDropEvent* event) override {
		if (handle__DropEvent == 0) {
			QGroupBox::dropEvent(event);
			return;
		}
		
		QDropEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_DropEvent(this, handle__DropEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DropEvent(QDropEvent* event) {

		QGroupBox::dropEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ShowEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void showEvent(QShowEvent* event) override {
		if (handle__ShowEvent == 0) {
			QGroupBox::showEvent(event);
			return;
		}
		
		QShowEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_ShowEvent(this, handle__ShowEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ShowEvent(QShowEvent* event) {

		QGroupBox::showEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HideEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void hideEvent(QHideEvent* event) override {
		if (handle__HideEvent == 0) {
			QGroupBox::hideEvent(event);
			return;
		}
		
		QHideEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_HideEvent(this, handle__HideEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_HideEvent(QHideEvent* event) {

		QGroupBox::hideEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__NativeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual bool nativeEvent(const QByteArray& eventType, void* message, long* result) override {
		if (handle__NativeEvent == 0) {
			return QGroupBox::nativeEvent(eventType, message, result);
		}
		
		const QByteArray eventType_qb = eventType;
		struct miqt_string eventType_ms;
		eventType_ms.len = eventType_qb.length();
		eventType_ms.data = static_cast<char*>(malloc(eventType_ms.len));
		memcpy(eventType_ms.data, eventType_qb.data(), eventType_ms.len);
		struct miqt_string sigval1 = eventType_ms;
		void* sigval2 = message;
		long* sigval3 = result;

		bool callback_return_value = miqt_exec_callback_QGroupBox_NativeEvent(this, handle__NativeEvent, sigval1, sigval2, sigval3);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_NativeEvent(struct miqt_string eventType, void* message, long* result) {
		QByteArray eventType_QByteArray(eventType.data, eventType.len);

		return QGroupBox::nativeEvent(eventType_QByteArray, message, static_cast<long*>(result));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metric = 0;

	// Subclass to allow providing a Go implementation
	virtual int metric(QPaintDevice::PaintDeviceMetric param1) const override {
		if (handle__Metric == 0) {
			return QGroupBox::metric(param1);
		}
		
		QPaintDevice::PaintDeviceMetric param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		int callback_return_value = miqt_exec_callback_QGroupBox_Metric(this, handle__Metric, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Metric(int param1) const {

		return QGroupBox::metric(static_cast<QPaintDevice::PaintDeviceMetric>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InitPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual void initPainter(QPainter* painter) const override {
		if (handle__InitPainter == 0) {
			QGroupBox::initPainter(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QGroupBox_InitPainter(this, handle__InitPainter, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_InitPainter(QPainter* painter) const {

		QGroupBox::initPainter(painter);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Redirected = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintDevice* redirected(QPoint* offset) const override {
		if (handle__Redirected == 0) {
			return QGroupBox::redirected(offset);
		}
		
		QPoint* sigval1 = offset;

		QPaintDevice* callback_return_value = miqt_exec_callback_QGroupBox_Redirected(this, handle__Redirected, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintDevice* virtualbase_Redirected(QPoint* offset) const {

		return QGroupBox::redirected(offset);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SharedPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual QPainter* sharedPainter() const override {
		if (handle__SharedPainter == 0) {
			return QGroupBox::sharedPainter();
		}
		

		QPainter* callback_return_value = miqt_exec_callback_QGroupBox_SharedPainter(this, handle__SharedPainter);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPainter* virtualbase_SharedPainter() const {

		return QGroupBox::sharedPainter();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InputMethodEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void inputMethodEvent(QInputMethodEvent* param1) override {
		if (handle__InputMethodEvent == 0) {
			QGroupBox::inputMethodEvent(param1);
			return;
		}
		
		QInputMethodEvent* sigval1 = param1;

		miqt_exec_callback_QGroupBox_InputMethodEvent(this, handle__InputMethodEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_InputMethodEvent(QInputMethodEvent* param1) {

		QGroupBox::inputMethodEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InputMethodQuery = 0;

	// Subclass to allow providing a Go implementation
	virtual QVariant inputMethodQuery(Qt::InputMethodQuery param1) const override {
		if (handle__InputMethodQuery == 0) {
			return QGroupBox::inputMethodQuery(param1);
		}
		
		Qt::InputMethodQuery param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		QVariant* callback_return_value = miqt_exec_callback_QGroupBox_InputMethodQuery(this, handle__InputMethodQuery, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QVariant* virtualbase_InputMethodQuery(int param1) const {

		return new QVariant(QGroupBox::inputMethodQuery(static_cast<Qt::InputMethodQuery>(param1)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusNextPrevChild = 0;

	// Subclass to allow providing a Go implementation
	virtual bool focusNextPrevChild(bool next) override {
		if (handle__FocusNextPrevChild == 0) {
			return QGroupBox::focusNextPrevChild(next);
		}
		
		bool sigval1 = next;

		bool callback_return_value = miqt_exec_callback_QGroupBox_FocusNextPrevChild(this, handle__FocusNextPrevChild, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_FocusNextPrevChild(bool next) {

		return QGroupBox::focusNextPrevChild(next);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__EventFilter == 0) {
			return QGroupBox::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QGroupBox_EventFilter(this, handle__EventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EventFilter(QObject* watched, QEvent* event) {

		return QGroupBox::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__TimerEvent == 0) {
			QGroupBox::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* event) {

		QGroupBox::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CustomEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__CustomEvent == 0) {
			QGroupBox::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QGroupBox_CustomEvent(this, handle__CustomEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CustomEvent(QEvent* event) {

		QGroupBox::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__ConnectNotify == 0) {
			QGroupBox::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QGroupBox_ConnectNotify(this, handle__ConnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ConnectNotify(QMetaMethod* signal) {

		QGroupBox::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DisconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__DisconnectNotify == 0) {
			QGroupBox::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QGroupBox_DisconnectNotify(this, handle__DisconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DisconnectNotify(QMetaMethod* signal) {

		QGroupBox::disconnectNotify(*signal);

	}

};

QGroupBox* QGroupBox_new(QWidget* parent) {
	return new MiqtVirtualQGroupBox(parent);
}

QGroupBox* QGroupBox_new2() {
	return new MiqtVirtualQGroupBox();
}

QGroupBox* QGroupBox_new3(struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	return new MiqtVirtualQGroupBox(title_QString);
}

QGroupBox* QGroupBox_new4(struct miqt_string title, QWidget* parent) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	return new MiqtVirtualQGroupBox(title_QString, parent);
}

void QGroupBox_virtbase(QGroupBox* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QGroupBox_MetaObject(const QGroupBox* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGroupBox_Metacast(QGroupBox* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGroupBox_Tr(const char* s) {
	QString _ret = QGroupBox::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGroupBox_TrUtf8(const char* s) {
	QString _ret = QGroupBox::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGroupBox_Title(const QGroupBox* self) {
	QString _ret = self->title();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGroupBox_SetTitle(QGroupBox* self, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	self->setTitle(title_QString);
}

int QGroupBox_Alignment(const QGroupBox* self) {
	Qt::Alignment _ret = self->alignment();
	return static_cast<int>(_ret);
}

void QGroupBox_SetAlignment(QGroupBox* self, int alignment) {
	self->setAlignment(static_cast<int>(alignment));
}

QSize* QGroupBox_MinimumSizeHint(const QGroupBox* self) {
	return new QSize(self->minimumSizeHint());
}

bool QGroupBox_IsFlat(const QGroupBox* self) {
	return self->isFlat();
}

void QGroupBox_SetFlat(QGroupBox* self, bool flat) {
	self->setFlat(flat);
}

bool QGroupBox_IsCheckable(const QGroupBox* self) {
	return self->isCheckable();
}

void QGroupBox_SetCheckable(QGroupBox* self, bool checkable) {
	self->setCheckable(checkable);
}

bool QGroupBox_IsChecked(const QGroupBox* self) {
	return self->isChecked();
}

void QGroupBox_SetChecked(QGroupBox* self, bool checked) {
	self->setChecked(checked);
}

void QGroupBox_Clicked(QGroupBox* self) {
	self->clicked();
}

void QGroupBox_connect_Clicked(QGroupBox* self, intptr_t slot) {
	MiqtVirtualQGroupBox::connect(self, static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::clicked), self, [=]() {
		miqt_exec_callback_QGroupBox_Clicked(slot);
	});
}

void QGroupBox_Toggled(QGroupBox* self, bool param1) {
	self->toggled(param1);
}

void QGroupBox_connect_Toggled(QGroupBox* self, intptr_t slot) {
	MiqtVirtualQGroupBox::connect(self, static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::toggled), self, [=](bool param1) {
		bool sigval1 = param1;
		miqt_exec_callback_QGroupBox_Toggled(slot, sigval1);
	});
}

struct miqt_string QGroupBox_Tr2(const char* s, const char* c) {
	QString _ret = QGroupBox::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGroupBox_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGroupBox::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGroupBox_TrUtf82(const char* s, const char* c) {
	QString _ret = QGroupBox::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGroupBox_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QGroupBox::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGroupBox_Clicked1(QGroupBox* self, bool checked) {
	self->clicked(checked);
}

void QGroupBox_connect_Clicked1(QGroupBox* self, intptr_t slot) {
	MiqtVirtualQGroupBox::connect(self, static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::clicked), self, [=](bool checked) {
		bool sigval1 = checked;
		miqt_exec_callback_QGroupBox_Clicked1(slot, sigval1);
	});
}

bool QGroupBox_override_virtual_MinimumSizeHint(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MinimumSizeHint = slot;
	return true;
}

QSize* QGroupBox_virtualbase_MinimumSizeHint(const void* self) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_MinimumSizeHint();
}

bool QGroupBox_override_virtual_Event(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Event = slot;
	return true;
}

bool QGroupBox_virtualbase_Event(void* self, QEvent* event) {
	return ( (MiqtVirtualQGroupBox*)(self) )->virtualbase_Event(event);
}

bool QGroupBox_override_virtual_ChildEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ChildEvent = slot;
	return true;
}

void QGroupBox_virtualbase_ChildEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_ChildEvent(event);
}

bool QGroupBox_override_virtual_ResizeEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ResizeEvent = slot;
	return true;
}

void QGroupBox_virtualbase_ResizeEvent(void* self, QResizeEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_ResizeEvent(event);
}

bool QGroupBox_override_virtual_PaintEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__PaintEvent = slot;
	return true;
}

void QGroupBox_virtualbase_PaintEvent(void* self, QPaintEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_PaintEvent(event);
}

bool QGroupBox_override_virtual_FocusInEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__FocusInEvent = slot;
	return true;
}

void QGroupBox_virtualbase_FocusInEvent(void* self, QFocusEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_FocusInEvent(event);
}

bool QGroupBox_override_virtual_ChangeEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ChangeEvent = slot;
	return true;
}

void QGroupBox_virtualbase_ChangeEvent(void* self, QEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_ChangeEvent(event);
}

bool QGroupBox_override_virtual_MousePressEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MousePressEvent = slot;
	return true;
}

void QGroupBox_virtualbase_MousePressEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_MousePressEvent(event);
}

bool QGroupBox_override_virtual_MouseMoveEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MouseMoveEvent = slot;
	return true;
}

void QGroupBox_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_MouseMoveEvent(event);
}

bool QGroupBox_override_virtual_MouseReleaseEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MouseReleaseEvent = slot;
	return true;
}

void QGroupBox_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_MouseReleaseEvent(event);
}

bool QGroupBox_override_virtual_DevType(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DevType = slot;
	return true;
}

int QGroupBox_virtualbase_DevType(const void* self) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_DevType();
}

bool QGroupBox_override_virtual_SetVisible(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__SetVisible = slot;
	return true;
}

void QGroupBox_virtualbase_SetVisible(void* self, bool visible) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_SetVisible(visible);
}

bool QGroupBox_override_virtual_SizeHint(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__SizeHint = slot;
	return true;
}

QSize* QGroupBox_virtualbase_SizeHint(const void* self) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_SizeHint();
}

bool QGroupBox_override_virtual_HeightForWidth(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__HeightForWidth = slot;
	return true;
}

int QGroupBox_virtualbase_HeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_HeightForWidth(param1);
}

bool QGroupBox_override_virtual_HasHeightForWidth(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__HasHeightForWidth = slot;
	return true;
}

bool QGroupBox_virtualbase_HasHeightForWidth(const void* self) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_HasHeightForWidth();
}

bool QGroupBox_override_virtual_PaintEngine(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__PaintEngine = slot;
	return true;
}

QPaintEngine* QGroupBox_virtualbase_PaintEngine(const void* self) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_PaintEngine();
}

bool QGroupBox_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MouseDoubleClickEvent = slot;
	return true;
}

void QGroupBox_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_MouseDoubleClickEvent(event);
}

bool QGroupBox_override_virtual_WheelEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__WheelEvent = slot;
	return true;
}

void QGroupBox_virtualbase_WheelEvent(void* self, QWheelEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_WheelEvent(event);
}

bool QGroupBox_override_virtual_KeyPressEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__KeyPressEvent = slot;
	return true;
}

void QGroupBox_virtualbase_KeyPressEvent(void* self, QKeyEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_KeyPressEvent(event);
}

bool QGroupBox_override_virtual_KeyReleaseEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__KeyReleaseEvent = slot;
	return true;
}

void QGroupBox_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_KeyReleaseEvent(event);
}

bool QGroupBox_override_virtual_FocusOutEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__FocusOutEvent = slot;
	return true;
}

void QGroupBox_virtualbase_FocusOutEvent(void* self, QFocusEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_FocusOutEvent(event);
}

bool QGroupBox_override_virtual_EnterEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__EnterEvent = slot;
	return true;
}

void QGroupBox_virtualbase_EnterEvent(void* self, QEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_EnterEvent(event);
}

bool QGroupBox_override_virtual_LeaveEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__LeaveEvent = slot;
	return true;
}

void QGroupBox_virtualbase_LeaveEvent(void* self, QEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_LeaveEvent(event);
}

bool QGroupBox_override_virtual_MoveEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MoveEvent = slot;
	return true;
}

void QGroupBox_virtualbase_MoveEvent(void* self, QMoveEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_MoveEvent(event);
}

bool QGroupBox_override_virtual_CloseEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__CloseEvent = slot;
	return true;
}

void QGroupBox_virtualbase_CloseEvent(void* self, QCloseEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_CloseEvent(event);
}

bool QGroupBox_override_virtual_ContextMenuEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ContextMenuEvent = slot;
	return true;
}

void QGroupBox_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_ContextMenuEvent(event);
}

bool QGroupBox_override_virtual_TabletEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__TabletEvent = slot;
	return true;
}

void QGroupBox_virtualbase_TabletEvent(void* self, QTabletEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_TabletEvent(event);
}

bool QGroupBox_override_virtual_ActionEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ActionEvent = slot;
	return true;
}

void QGroupBox_virtualbase_ActionEvent(void* self, QActionEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_ActionEvent(event);
}

bool QGroupBox_override_virtual_DragEnterEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DragEnterEvent = slot;
	return true;
}

void QGroupBox_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_DragEnterEvent(event);
}

bool QGroupBox_override_virtual_DragMoveEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DragMoveEvent = slot;
	return true;
}

void QGroupBox_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_DragMoveEvent(event);
}

bool QGroupBox_override_virtual_DragLeaveEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DragLeaveEvent = slot;
	return true;
}

void QGroupBox_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_DragLeaveEvent(event);
}

bool QGroupBox_override_virtual_DropEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DropEvent = slot;
	return true;
}

void QGroupBox_virtualbase_DropEvent(void* self, QDropEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_DropEvent(event);
}

bool QGroupBox_override_virtual_ShowEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ShowEvent = slot;
	return true;
}

void QGroupBox_virtualbase_ShowEvent(void* self, QShowEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_ShowEvent(event);
}

bool QGroupBox_override_virtual_HideEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__HideEvent = slot;
	return true;
}

void QGroupBox_virtualbase_HideEvent(void* self, QHideEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_HideEvent(event);
}

bool QGroupBox_override_virtual_NativeEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__NativeEvent = slot;
	return true;
}

bool QGroupBox_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, long* result) {
	return ( (MiqtVirtualQGroupBox*)(self) )->virtualbase_NativeEvent(eventType, message, result);
}

bool QGroupBox_override_virtual_Metric(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Metric = slot;
	return true;
}

int QGroupBox_virtualbase_Metric(const void* self, int param1) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_Metric(param1);
}

bool QGroupBox_override_virtual_InitPainter(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__InitPainter = slot;
	return true;
}

void QGroupBox_virtualbase_InitPainter(const void* self, QPainter* painter) {
	( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_InitPainter(painter);
}

bool QGroupBox_override_virtual_Redirected(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Redirected = slot;
	return true;
}

QPaintDevice* QGroupBox_virtualbase_Redirected(const void* self, QPoint* offset) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_Redirected(offset);
}

bool QGroupBox_override_virtual_SharedPainter(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__SharedPainter = slot;
	return true;
}

QPainter* QGroupBox_virtualbase_SharedPainter(const void* self) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_SharedPainter();
}

bool QGroupBox_override_virtual_InputMethodEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__InputMethodEvent = slot;
	return true;
}

void QGroupBox_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_InputMethodEvent(param1);
}

bool QGroupBox_override_virtual_InputMethodQuery(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__InputMethodQuery = slot;
	return true;
}

QVariant* QGroupBox_virtualbase_InputMethodQuery(const void* self, int param1) {
	return ( (const MiqtVirtualQGroupBox*)(self) )->virtualbase_InputMethodQuery(param1);
}

bool QGroupBox_override_virtual_FocusNextPrevChild(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__FocusNextPrevChild = slot;
	return true;
}

bool QGroupBox_virtualbase_FocusNextPrevChild(void* self, bool next) {
	return ( (MiqtVirtualQGroupBox*)(self) )->virtualbase_FocusNextPrevChild(next);
}

bool QGroupBox_override_virtual_EventFilter(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__EventFilter = slot;
	return true;
}

bool QGroupBox_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQGroupBox*)(self) )->virtualbase_EventFilter(watched, event);
}

bool QGroupBox_override_virtual_TimerEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__TimerEvent = slot;
	return true;
}

void QGroupBox_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_TimerEvent(event);
}

bool QGroupBox_override_virtual_CustomEvent(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__CustomEvent = slot;
	return true;
}

void QGroupBox_virtualbase_CustomEvent(void* self, QEvent* event) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_CustomEvent(event);
}

bool QGroupBox_override_virtual_ConnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ConnectNotify = slot;
	return true;
}

void QGroupBox_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_ConnectNotify(signal);
}

bool QGroupBox_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQGroupBox* self_cast = dynamic_cast<MiqtVirtualQGroupBox*>( (QGroupBox*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DisconnectNotify = slot;
	return true;
}

void QGroupBox_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQGroupBox*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QGroupBox_Delete(QGroupBox* self) {
	delete self;
}

