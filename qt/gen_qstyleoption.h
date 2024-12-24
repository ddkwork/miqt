#pragma once
#ifndef MIQT_QT_GEN_QSTYLEOPTION_H
#define MIQT_QT_GEN_QSTYLEOPTION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QStyleHintReturn;
class QStyleHintReturnMask;
class QStyleHintReturnVariant;
class QStyleOption;
class QStyleOptionButton;
class QStyleOptionComboBox;
class QStyleOptionComplex;
class QStyleOptionDockWidget;
class QStyleOptionFocusRect;
class QStyleOptionFrame;
class QStyleOptionGraphicsItem;
class QStyleOptionGroupBox;
class QStyleOptionHeader;
class QStyleOptionHeaderV2;
class QStyleOptionMenuItem;
class QStyleOptionProgressBar;
class QStyleOptionRubberBand;
class QStyleOptionSizeGrip;
class QStyleOptionSlider;
class QStyleOptionSpinBox;
class QStyleOptionTab;
class QStyleOptionTabBarBase;
class QStyleOptionTabWidgetFrame;
class QStyleOptionTitleBar;
class QStyleOptionToolBar;
class QStyleOptionToolBox;
class QStyleOptionToolButton;
class QStyleOptionViewItem;
class QTransform;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QStyleHintReturn QStyleHintReturn;
typedef struct QStyleHintReturnMask QStyleHintReturnMask;
typedef struct QStyleHintReturnVariant QStyleHintReturnVariant;
typedef struct QStyleOption QStyleOption;
typedef struct QStyleOptionButton QStyleOptionButton;
typedef struct QStyleOptionComboBox QStyleOptionComboBox;
typedef struct QStyleOptionComplex QStyleOptionComplex;
typedef struct QStyleOptionDockWidget QStyleOptionDockWidget;
typedef struct QStyleOptionFocusRect QStyleOptionFocusRect;
typedef struct QStyleOptionFrame QStyleOptionFrame;
typedef struct QStyleOptionGraphicsItem QStyleOptionGraphicsItem;
typedef struct QStyleOptionGroupBox QStyleOptionGroupBox;
typedef struct QStyleOptionHeader QStyleOptionHeader;
typedef struct QStyleOptionHeaderV2 QStyleOptionHeaderV2;
typedef struct QStyleOptionMenuItem QStyleOptionMenuItem;
typedef struct QStyleOptionProgressBar QStyleOptionProgressBar;
typedef struct QStyleOptionRubberBand QStyleOptionRubberBand;
typedef struct QStyleOptionSizeGrip QStyleOptionSizeGrip;
typedef struct QStyleOptionSlider QStyleOptionSlider;
typedef struct QStyleOptionSpinBox QStyleOptionSpinBox;
typedef struct QStyleOptionTab QStyleOptionTab;
typedef struct QStyleOptionTabBarBase QStyleOptionTabBarBase;
typedef struct QStyleOptionTabWidgetFrame QStyleOptionTabWidgetFrame;
typedef struct QStyleOptionTitleBar QStyleOptionTitleBar;
typedef struct QStyleOptionToolBar QStyleOptionToolBar;
typedef struct QStyleOptionToolBox QStyleOptionToolBox;
typedef struct QStyleOptionToolButton QStyleOptionToolButton;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QTransform QTransform;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStyleOption* QStyleOption_new();
extern __declspec(dllexport) QStyleOption* QStyleOption_new2(QStyleOption* other);
extern __declspec(dllexport) QStyleOption* QStyleOption_new3(int version);
extern __declspec(dllexport) QStyleOption* QStyleOption_new4(int version, int typeVal);
extern __declspec(dllexport) void QStyleOption_InitFrom(QStyleOption* self, QWidget* w);
extern __declspec(dllexport) void QStyleOption_OperatorAssign(QStyleOption* self, QStyleOption* other);
extern __declspec(dllexport) void QStyleOption_Delete(QStyleOption* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionFocusRect* QStyleOptionFocusRect_new();
extern __declspec(dllexport) QStyleOptionFocusRect* QStyleOptionFocusRect_new2(QStyleOptionFocusRect* other);
extern __declspec(dllexport) void QStyleOptionFocusRect_virtbase(QStyleOptionFocusRect* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionFocusRect_Delete(QStyleOptionFocusRect* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionFrame* QStyleOptionFrame_new();
extern __declspec(dllexport) QStyleOptionFrame* QStyleOptionFrame_new2(QStyleOptionFrame* other);
extern __declspec(dllexport) void QStyleOptionFrame_virtbase(QStyleOptionFrame* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionFrame_Delete(QStyleOptionFrame* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionTabWidgetFrame* QStyleOptionTabWidgetFrame_new();
extern __declspec(dllexport) QStyleOptionTabWidgetFrame* QStyleOptionTabWidgetFrame_new2(QStyleOptionTabWidgetFrame* other);
extern __declspec(dllexport) void QStyleOptionTabWidgetFrame_virtbase(QStyleOptionTabWidgetFrame* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionTabWidgetFrame_Delete(QStyleOptionTabWidgetFrame* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionTabBarBase* QStyleOptionTabBarBase_new();
extern __declspec(dllexport) QStyleOptionTabBarBase* QStyleOptionTabBarBase_new2(QStyleOptionTabBarBase* other);
extern __declspec(dllexport) void QStyleOptionTabBarBase_virtbase(QStyleOptionTabBarBase* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionTabBarBase_Delete(QStyleOptionTabBarBase* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionHeader* QStyleOptionHeader_new();
extern __declspec(dllexport) QStyleOptionHeader* QStyleOptionHeader_new2(QStyleOptionHeader* other);
extern __declspec(dllexport) void QStyleOptionHeader_virtbase(QStyleOptionHeader* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionHeader_Delete(QStyleOptionHeader* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionHeaderV2* QStyleOptionHeaderV2_new();
extern __declspec(dllexport) QStyleOptionHeaderV2* QStyleOptionHeaderV2_new2(QStyleOptionHeaderV2* other);
extern __declspec(dllexport) void QStyleOptionHeaderV2_virtbase(QStyleOptionHeaderV2* src, QStyleOptionHeader** outptr_QStyleOptionHeader);
extern __declspec(dllexport) void QStyleOptionHeaderV2_Delete(QStyleOptionHeaderV2* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionButton* QStyleOptionButton_new();
extern __declspec(dllexport) QStyleOptionButton* QStyleOptionButton_new2(QStyleOptionButton* other);
extern __declspec(dllexport) void QStyleOptionButton_virtbase(QStyleOptionButton* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionButton_Delete(QStyleOptionButton* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionTab* QStyleOptionTab_new();
extern __declspec(dllexport) QStyleOptionTab* QStyleOptionTab_new2(QStyleOptionTab* other);
extern __declspec(dllexport) void QStyleOptionTab_virtbase(QStyleOptionTab* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionTab_Delete(QStyleOptionTab* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionToolBar* QStyleOptionToolBar_new();
extern __declspec(dllexport) QStyleOptionToolBar* QStyleOptionToolBar_new2(QStyleOptionToolBar* other);
extern __declspec(dllexport) void QStyleOptionToolBar_virtbase(QStyleOptionToolBar* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionToolBar_Delete(QStyleOptionToolBar* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionProgressBar* QStyleOptionProgressBar_new();
extern __declspec(dllexport) QStyleOptionProgressBar* QStyleOptionProgressBar_new2(QStyleOptionProgressBar* other);
extern __declspec(dllexport) void QStyleOptionProgressBar_virtbase(QStyleOptionProgressBar* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionProgressBar_Delete(QStyleOptionProgressBar* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionMenuItem* QStyleOptionMenuItem_new();
extern __declspec(dllexport) QStyleOptionMenuItem* QStyleOptionMenuItem_new2(QStyleOptionMenuItem* other);
extern __declspec(dllexport) void QStyleOptionMenuItem_virtbase(QStyleOptionMenuItem* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionMenuItem_Delete(QStyleOptionMenuItem* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionDockWidget* QStyleOptionDockWidget_new();
extern __declspec(dllexport) QStyleOptionDockWidget* QStyleOptionDockWidget_new2(QStyleOptionDockWidget* other);
extern __declspec(dllexport) void QStyleOptionDockWidget_virtbase(QStyleOptionDockWidget* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionDockWidget_Delete(QStyleOptionDockWidget* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionViewItem* QStyleOptionViewItem_new();
extern __declspec(dllexport) QStyleOptionViewItem* QStyleOptionViewItem_new2(QStyleOptionViewItem* other);
extern __declspec(dllexport) void QStyleOptionViewItem_virtbase(QStyleOptionViewItem* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionViewItem_Delete(QStyleOptionViewItem* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionToolBox* QStyleOptionToolBox_new();
extern __declspec(dllexport) QStyleOptionToolBox* QStyleOptionToolBox_new2(QStyleOptionToolBox* other);
extern __declspec(dllexport) void QStyleOptionToolBox_virtbase(QStyleOptionToolBox* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionToolBox_Delete(QStyleOptionToolBox* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionRubberBand* QStyleOptionRubberBand_new();
extern __declspec(dllexport) QStyleOptionRubberBand* QStyleOptionRubberBand_new2(QStyleOptionRubberBand* other);
extern __declspec(dllexport) void QStyleOptionRubberBand_virtbase(QStyleOptionRubberBand* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionRubberBand_Delete(QStyleOptionRubberBand* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionComplex* QStyleOptionComplex_new();
extern __declspec(dllexport) QStyleOptionComplex* QStyleOptionComplex_new2(QStyleOptionComplex* other);
extern __declspec(dllexport) QStyleOptionComplex* QStyleOptionComplex_new3(int version);
extern __declspec(dllexport) QStyleOptionComplex* QStyleOptionComplex_new4(int version, int typeVal);
extern __declspec(dllexport) void QStyleOptionComplex_virtbase(QStyleOptionComplex* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) void QStyleOptionComplex_Delete(QStyleOptionComplex* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionSlider* QStyleOptionSlider_new();
extern __declspec(dllexport) QStyleOptionSlider* QStyleOptionSlider_new2(QStyleOptionSlider* other);
extern __declspec(dllexport) void QStyleOptionSlider_virtbase(QStyleOptionSlider* src, QStyleOptionComplex** outptr_QStyleOptionComplex);
extern __declspec(dllexport) void QStyleOptionSlider_Delete(QStyleOptionSlider* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionSpinBox* QStyleOptionSpinBox_new();
extern __declspec(dllexport) QStyleOptionSpinBox* QStyleOptionSpinBox_new2(QStyleOptionSpinBox* other);
extern __declspec(dllexport) void QStyleOptionSpinBox_virtbase(QStyleOptionSpinBox* src, QStyleOptionComplex** outptr_QStyleOptionComplex);
extern __declspec(dllexport) void QStyleOptionSpinBox_Delete(QStyleOptionSpinBox* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionToolButton* QStyleOptionToolButton_new();
extern __declspec(dllexport) QStyleOptionToolButton* QStyleOptionToolButton_new2(QStyleOptionToolButton* other);
extern __declspec(dllexport) void QStyleOptionToolButton_virtbase(QStyleOptionToolButton* src, QStyleOptionComplex** outptr_QStyleOptionComplex);
extern __declspec(dllexport) void QStyleOptionToolButton_Delete(QStyleOptionToolButton* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionComboBox* QStyleOptionComboBox_new();
extern __declspec(dllexport) QStyleOptionComboBox* QStyleOptionComboBox_new2(QStyleOptionComboBox* other);
extern __declspec(dllexport) void QStyleOptionComboBox_virtbase(QStyleOptionComboBox* src, QStyleOptionComplex** outptr_QStyleOptionComplex);
extern __declspec(dllexport) void QStyleOptionComboBox_Delete(QStyleOptionComboBox* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionTitleBar* QStyleOptionTitleBar_new();
extern __declspec(dllexport) QStyleOptionTitleBar* QStyleOptionTitleBar_new2(QStyleOptionTitleBar* other);
extern __declspec(dllexport) void QStyleOptionTitleBar_virtbase(QStyleOptionTitleBar* src, QStyleOptionComplex** outptr_QStyleOptionComplex);
extern __declspec(dllexport) void QStyleOptionTitleBar_Delete(QStyleOptionTitleBar* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionGroupBox* QStyleOptionGroupBox_new();
extern __declspec(dllexport) QStyleOptionGroupBox* QStyleOptionGroupBox_new2(QStyleOptionGroupBox* other);
extern __declspec(dllexport) void QStyleOptionGroupBox_virtbase(QStyleOptionGroupBox* src, QStyleOptionComplex** outptr_QStyleOptionComplex);
extern __declspec(dllexport) void QStyleOptionGroupBox_Delete(QStyleOptionGroupBox* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionSizeGrip* QStyleOptionSizeGrip_new();
extern __declspec(dllexport) QStyleOptionSizeGrip* QStyleOptionSizeGrip_new2(QStyleOptionSizeGrip* other);
extern __declspec(dllexport) void QStyleOptionSizeGrip_virtbase(QStyleOptionSizeGrip* src, QStyleOptionComplex** outptr_QStyleOptionComplex);
extern __declspec(dllexport) void QStyleOptionSizeGrip_Delete(QStyleOptionSizeGrip* self, bool isSubclass);

extern __declspec(dllexport) QStyleOptionGraphicsItem* QStyleOptionGraphicsItem_new();
extern __declspec(dllexport) QStyleOptionGraphicsItem* QStyleOptionGraphicsItem_new2(QStyleOptionGraphicsItem* other);
extern __declspec(dllexport) void QStyleOptionGraphicsItem_virtbase(QStyleOptionGraphicsItem* src, QStyleOption** outptr_QStyleOption);
extern __declspec(dllexport) double QStyleOptionGraphicsItem_LevelOfDetailFromTransform(QTransform* worldTransform);
extern __declspec(dllexport) void QStyleOptionGraphicsItem_Delete(QStyleOptionGraphicsItem* self, bool isSubclass);

extern __declspec(dllexport) QStyleHintReturn* QStyleHintReturn_new();
extern __declspec(dllexport) QStyleHintReturn* QStyleHintReturn_new2(QStyleHintReturn* param1);
extern __declspec(dllexport) QStyleHintReturn* QStyleHintReturn_new3(int version);
extern __declspec(dllexport) QStyleHintReturn* QStyleHintReturn_new4(int version, int typeVal);
extern __declspec(dllexport) void QStyleHintReturn_OperatorAssign(QStyleHintReturn* self, QStyleHintReturn* param1);
extern __declspec(dllexport) void QStyleHintReturn_Delete(QStyleHintReturn* self, bool isSubclass);

extern __declspec(dllexport) QStyleHintReturnMask* QStyleHintReturnMask_new();
extern __declspec(dllexport) QStyleHintReturnMask* QStyleHintReturnMask_new2(QStyleHintReturnMask* param1);
extern __declspec(dllexport) void QStyleHintReturnMask_virtbase(QStyleHintReturnMask* src, QStyleHintReturn** outptr_QStyleHintReturn);
extern __declspec(dllexport) void QStyleHintReturnMask_OperatorAssign(QStyleHintReturnMask* self, QStyleHintReturnMask* param1);
extern __declspec(dllexport) void QStyleHintReturnMask_Delete(QStyleHintReturnMask* self, bool isSubclass);

extern __declspec(dllexport) QStyleHintReturnVariant* QStyleHintReturnVariant_new();
extern __declspec(dllexport) QStyleHintReturnVariant* QStyleHintReturnVariant_new2(QStyleHintReturnVariant* param1);
extern __declspec(dllexport) void QStyleHintReturnVariant_virtbase(QStyleHintReturnVariant* src, QStyleHintReturn** outptr_QStyleHintReturn);
extern __declspec(dllexport) void QStyleHintReturnVariant_OperatorAssign(QStyleHintReturnVariant* self, QStyleHintReturnVariant* param1);
extern __declspec(dllexport) void QStyleHintReturnVariant_Delete(QStyleHintReturnVariant* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
