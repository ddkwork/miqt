#pragma once
#ifndef MIQT_QT_GEN_QCOLOR_H
#define MIQT_QT_GEN_QCOLOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QColor QColor;
typedef struct QRgba64 QRgba64;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QColor* QColor_new();
extern __declspec(dllexport) QColor* QColor_new2(int color);
extern __declspec(dllexport) QColor* QColor_new3(int r, int g, int b);
extern __declspec(dllexport) QColor* QColor_new4(unsigned int rgb);
extern __declspec(dllexport) QColor* QColor_new5(QRgba64* rgba64);
extern __declspec(dllexport) QColor* QColor_new6(struct miqt_string name);
extern __declspec(dllexport) QColor* QColor_new7(const char* aname);
extern __declspec(dllexport) QColor* QColor_new8(Spec spec);
extern __declspec(dllexport) QColor* QColor_new9(Spec spec, uint16_t a1, uint16_t a2, uint16_t a3, uint16_t a4);
extern __declspec(dllexport) QColor* QColor_new10(QColor* param1);
extern __declspec(dllexport) QColor* QColor_new11(int r, int g, int b, int a);
extern __declspec(dllexport) QColor* QColor_new12(Spec spec, uint16_t a1, uint16_t a2, uint16_t a3, uint16_t a4, uint16_t a5);
extern __declspec(dllexport) QColor* QColor_FromString(QAnyStringView* name);
extern __declspec(dllexport) void QColor_OperatorAssign(QColor* self, int color);
extern __declspec(dllexport) bool QColor_IsValid(const QColor* self);
extern __declspec(dllexport) struct miqt_string QColor_Name(const QColor* self);
extern __declspec(dllexport) void QColor_SetNamedColor(QColor* self, struct miqt_string name);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QColor_ColorNames();
extern __declspec(dllexport) Spec QColor_Spec(const QColor* self);
extern __declspec(dllexport) int QColor_Alpha(const QColor* self);
extern __declspec(dllexport) void QColor_SetAlpha(QColor* self, int alpha);
extern __declspec(dllexport) float QColor_AlphaF(const QColor* self);
extern __declspec(dllexport) void QColor_SetAlphaF(QColor* self, float alpha);
extern __declspec(dllexport) int QColor_Red(const QColor* self);
extern __declspec(dllexport) int QColor_Green(const QColor* self);
extern __declspec(dllexport) int QColor_Blue(const QColor* self);
extern __declspec(dllexport) void QColor_SetRed(QColor* self, int red);
extern __declspec(dllexport) void QColor_SetGreen(QColor* self, int green);
extern __declspec(dllexport) void QColor_SetBlue(QColor* self, int blue);
extern __declspec(dllexport) float QColor_RedF(const QColor* self);
extern __declspec(dllexport) float QColor_GreenF(const QColor* self);
extern __declspec(dllexport) float QColor_BlueF(const QColor* self);
extern __declspec(dllexport) void QColor_SetRedF(QColor* self, float red);
extern __declspec(dllexport) void QColor_SetGreenF(QColor* self, float green);
extern __declspec(dllexport) void QColor_SetBlueF(QColor* self, float blue);
extern __declspec(dllexport) void QColor_GetRgb(const QColor* self, int* r, int* g, int* b);
extern __declspec(dllexport) void QColor_SetRgb(QColor* self, int r, int g, int b);
extern __declspec(dllexport) void QColor_GetRgbF(const QColor* self, float* r, float* g, float* b);
extern __declspec(dllexport) void QColor_SetRgbF(QColor* self, float r, float g, float b);
extern __declspec(dllexport) QRgba64* QColor_Rgba64(const QColor* self);
extern __declspec(dllexport) void QColor_SetRgba64(QColor* self, QRgba64* rgba);
extern __declspec(dllexport) unsigned int QColor_Rgba(const QColor* self);
extern __declspec(dllexport) void QColor_SetRgba(QColor* self, unsigned int rgba);
extern __declspec(dllexport) unsigned int QColor_Rgb(const QColor* self);
extern __declspec(dllexport) void QColor_SetRgbWithRgb(QColor* self, unsigned int rgb);
extern __declspec(dllexport) int QColor_Hue(const QColor* self);
extern __declspec(dllexport) int QColor_Saturation(const QColor* self);
extern __declspec(dllexport) int QColor_HsvHue(const QColor* self);
extern __declspec(dllexport) int QColor_HsvSaturation(const QColor* self);
extern __declspec(dllexport) int QColor_Value(const QColor* self);
extern __declspec(dllexport) float QColor_HueF(const QColor* self);
extern __declspec(dllexport) float QColor_SaturationF(const QColor* self);
extern __declspec(dllexport) float QColor_HsvHueF(const QColor* self);
extern __declspec(dllexport) float QColor_HsvSaturationF(const QColor* self);
extern __declspec(dllexport) float QColor_ValueF(const QColor* self);
extern __declspec(dllexport) void QColor_GetHsv(const QColor* self, int* h, int* s, int* v);
extern __declspec(dllexport) void QColor_SetHsv(QColor* self, int h, int s, int v);
extern __declspec(dllexport) void QColor_GetHsvF(const QColor* self, float* h, float* s, float* v);
extern __declspec(dllexport) void QColor_SetHsvF(QColor* self, float h, float s, float v);
extern __declspec(dllexport) int QColor_Cyan(const QColor* self);
extern __declspec(dllexport) int QColor_Magenta(const QColor* self);
extern __declspec(dllexport) int QColor_Yellow(const QColor* self);
extern __declspec(dllexport) int QColor_Black(const QColor* self);
extern __declspec(dllexport) float QColor_CyanF(const QColor* self);
extern __declspec(dllexport) float QColor_MagentaF(const QColor* self);
extern __declspec(dllexport) float QColor_YellowF(const QColor* self);
extern __declspec(dllexport) float QColor_BlackF(const QColor* self);
extern __declspec(dllexport) void QColor_GetCmyk(const QColor* self, int* c, int* m, int* y, int* k);
extern __declspec(dllexport) void QColor_SetCmyk(QColor* self, int c, int m, int y, int k);
extern __declspec(dllexport) void QColor_GetCmykF(const QColor* self, float* c, float* m, float* y, float* k);
extern __declspec(dllexport) void QColor_SetCmykF(QColor* self, float c, float m, float y, float k);
extern __declspec(dllexport) int QColor_HslHue(const QColor* self);
extern __declspec(dllexport) int QColor_HslSaturation(const QColor* self);
extern __declspec(dllexport) int QColor_Lightness(const QColor* self);
extern __declspec(dllexport) float QColor_HslHueF(const QColor* self);
extern __declspec(dllexport) float QColor_HslSaturationF(const QColor* self);
extern __declspec(dllexport) float QColor_LightnessF(const QColor* self);
extern __declspec(dllexport) void QColor_GetHsl(const QColor* self, int* h, int* s, int* l);
extern __declspec(dllexport) void QColor_SetHsl(QColor* self, int h, int s, int l);
extern __declspec(dllexport) void QColor_GetHslF(const QColor* self, float* h, float* s, float* l);
extern __declspec(dllexport) void QColor_SetHslF(QColor* self, float h, float s, float l);
extern __declspec(dllexport) QColor* QColor_ToRgb(const QColor* self);
extern __declspec(dllexport) QColor* QColor_ToHsv(const QColor* self);
extern __declspec(dllexport) QColor* QColor_ToCmyk(const QColor* self);
extern __declspec(dllexport) QColor* QColor_ToHsl(const QColor* self);
extern __declspec(dllexport) QColor* QColor_ToExtendedRgb(const QColor* self);
extern __declspec(dllexport) QColor* QColor_ConvertTo(const QColor* self, Spec colorSpec);
extern __declspec(dllexport) QColor* QColor_FromRgb(unsigned int rgb);
extern __declspec(dllexport) QColor* QColor_FromRgba(unsigned int rgba);
extern __declspec(dllexport) QColor* QColor_FromRgb2(int r, int g, int b);
extern __declspec(dllexport) QColor* QColor_FromRgbF(float r, float g, float b);
extern __declspec(dllexport) QColor* QColor_FromRgba64(uint16_t r, uint16_t g, uint16_t b);
extern __declspec(dllexport) QColor* QColor_FromRgba64WithRgba(QRgba64* rgba);
extern __declspec(dllexport) QColor* QColor_FromHsv(int h, int s, int v);
extern __declspec(dllexport) QColor* QColor_FromHsvF(float h, float s, float v);
extern __declspec(dllexport) QColor* QColor_FromCmyk(int c, int m, int y, int k);
extern __declspec(dllexport) QColor* QColor_FromCmykF(float c, float m, float y, float k);
extern __declspec(dllexport) QColor* QColor_FromHsl(int h, int s, int l);
extern __declspec(dllexport) QColor* QColor_FromHslF(float h, float s, float l);
extern __declspec(dllexport) QColor* QColor_Lighter(const QColor* self);
extern __declspec(dllexport) QColor* QColor_Darker(const QColor* self);
extern __declspec(dllexport) bool QColor_OperatorEqual(const QColor* self, QColor* c);
extern __declspec(dllexport) bool QColor_OperatorNotEqual(const QColor* self, QColor* c);
extern __declspec(dllexport) bool QColor_IsValidColor(struct miqt_string name);
extern __declspec(dllexport) bool QColor_IsValidColorName(QAnyStringView* param1);
extern __declspec(dllexport) struct miqt_string QColor_Name1(const QColor* self, NameFormat format);
extern __declspec(dllexport) void QColor_GetRgb4(const QColor* self, int* r, int* g, int* b, int* a);
extern __declspec(dllexport) void QColor_SetRgb4(QColor* self, int r, int g, int b, int a);
extern __declspec(dllexport) void QColor_GetRgbF4(const QColor* self, float* r, float* g, float* b, float* a);
extern __declspec(dllexport) void QColor_SetRgbF4(QColor* self, float r, float g, float b, float a);
extern __declspec(dllexport) void QColor_GetHsv4(const QColor* self, int* h, int* s, int* v, int* a);
extern __declspec(dllexport) void QColor_SetHsv4(QColor* self, int h, int s, int v, int a);
extern __declspec(dllexport) void QColor_GetHsvF4(const QColor* self, float* h, float* s, float* v, float* a);
extern __declspec(dllexport) void QColor_SetHsvF4(QColor* self, float h, float s, float v, float a);
extern __declspec(dllexport) void QColor_GetCmyk5(const QColor* self, int* c, int* m, int* y, int* k, int* a);
extern __declspec(dllexport) void QColor_SetCmyk5(QColor* self, int c, int m, int y, int k, int a);
extern __declspec(dllexport) void QColor_GetCmykF5(const QColor* self, float* c, float* m, float* y, float* k, float* a);
extern __declspec(dllexport) void QColor_SetCmykF5(QColor* self, float c, float m, float y, float k, float a);
extern __declspec(dllexport) void QColor_GetHsl4(const QColor* self, int* h, int* s, int* l, int* a);
extern __declspec(dllexport) void QColor_SetHsl4(QColor* self, int h, int s, int l, int a);
extern __declspec(dllexport) void QColor_GetHslF4(const QColor* self, float* h, float* s, float* l, float* a);
extern __declspec(dllexport) void QColor_SetHslF4(QColor* self, float h, float s, float l, float a);
extern __declspec(dllexport) QColor* QColor_FromRgb4(int r, int g, int b, int a);
extern __declspec(dllexport) QColor* QColor_FromRgbF4(float r, float g, float b, float a);
extern __declspec(dllexport) QColor* QColor_FromRgba644(uint16_t r, uint16_t g, uint16_t b, uint16_t a);
extern __declspec(dllexport) QColor* QColor_FromHsv4(int h, int s, int v, int a);
extern __declspec(dllexport) QColor* QColor_FromHsvF4(float h, float s, float v, float a);
extern __declspec(dllexport) QColor* QColor_FromCmyk5(int c, int m, int y, int k, int a);
extern __declspec(dllexport) QColor* QColor_FromCmykF5(float c, float m, float y, float k, float a);
extern __declspec(dllexport) QColor* QColor_FromHsl4(int h, int s, int l, int a);
extern __declspec(dllexport) QColor* QColor_FromHslF4(float h, float s, float l, float a);
extern __declspec(dllexport) QColor* QColor_Lighter1(const QColor* self, int f);
extern __declspec(dllexport) QColor* QColor_Darker1(const QColor* self, int f);
extern __declspec(dllexport) void QColor_Delete(QColor* self, bool isSubclass);

} 
