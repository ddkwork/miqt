#ifndef GEN_QANIMATIONGROUP_H
#define GEN_QANIMATIONGROUP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractAnimation;
class QAnimationGroup;
class QMetaObject;
#else
typedef struct QAbstractAnimation QAbstractAnimation;
typedef struct QAnimationGroup QAnimationGroup;
typedef struct QMetaObject QMetaObject;
#endif

QMetaObject* QAnimationGroup_MetaObject(const QAnimationGroup* self);
void QAnimationGroup_Tr(const char* s, char** _out, int* _out_Strlen);
void QAnimationGroup_TrUtf8(const char* s, char** _out, int* _out_Strlen);
QAbstractAnimation* QAnimationGroup_AnimationAt(const QAnimationGroup* self, int index);
int QAnimationGroup_AnimationCount(const QAnimationGroup* self);
int QAnimationGroup_IndexOfAnimation(const QAnimationGroup* self, QAbstractAnimation* animation);
void QAnimationGroup_AddAnimation(QAnimationGroup* self, QAbstractAnimation* animation);
void QAnimationGroup_InsertAnimation(QAnimationGroup* self, int index, QAbstractAnimation* animation);
void QAnimationGroup_RemoveAnimation(QAnimationGroup* self, QAbstractAnimation* animation);
QAbstractAnimation* QAnimationGroup_TakeAnimation(QAnimationGroup* self, int index);
void QAnimationGroup_Clear(QAnimationGroup* self);
void QAnimationGroup_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QAnimationGroup_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QAnimationGroup_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QAnimationGroup_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QAnimationGroup_Delete(QAnimationGroup* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
