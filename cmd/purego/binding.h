#ifndef BINDING_H
#define BINDING_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

typedef void* PQApplication;
typedef void* PQPushButton;
typedef void* PQWidget;
typedef void (*CallbackFunc)(int argc, void* args);
extern __declspec(dllexport) void miqt_exec_callback(CallbackFunc cb, int argc, void* args);
extern __declspec(dllexport)  PQApplication QApplication_new(int* argc, char** argv);
extern __declspec(dllexport) PQWidget QWidget_new();
extern __declspec(dllexport) void QWidget_show(PQWidget self);
extern __declspec(dllexport) PQPushButton QPushButton_new(const char* label, PQWidget parent);
extern __declspec(dllexport) void QPushButton_show(PQPushButton self);
extern __declspec(dllexport) void QPushButton_connect_pressed(PQPushButton self, intptr_t cb);
extern __declspec(dllexport) int QApplication_exec(PQApplication self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
