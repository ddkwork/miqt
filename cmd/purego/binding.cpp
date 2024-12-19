//+build ignore

#include "binding.h"
#include <QApplication>
#include <QWidget>
#include <QPushButton>

void miqt_exec_callback(CallbackFunc cb, int argc, void* args) {
    if (cb == NULL) {
        fprintf(stderr, "miqt: callback is NULL (heap corruption?)\n");
        return; // 如果回调函数为空，返回错误
    }
    cb(argc, args);
}

PQApplication QApplication_new(int* argc, char** argv) {
    // QApplication takes these parameters byref, not by value
    return new QApplication(*argc, argv);
}

PQWidget QWidget_new() {
    return new QWidget();
}

void QWidget_show(PQWidget self) {
    static_cast<QWidget*>(self)->show();    
}

PQPushButton QPushButton_new(const char* label, PQWidget parent) {
    return new QPushButton(label, static_cast<QWidget*>(parent));
}

void QPushButton_show(PQPushButton self) {
    static_cast<QPushButton*>(self)->show();
}

void QPushButton_connect_pressed(PQPushButton self, intptr_t cb) {
    QPushButton::connect(static_cast<QPushButton*>(self), &QPushButton::pressed, [=]() {
        miqt_exec_callback(reinterpret_cast<CallbackFunc>(cb), 0, nullptr);
    });
}

int QApplication_exec(PQApplication self) {
    return static_cast<QApplication*>(self)->exec();
}
