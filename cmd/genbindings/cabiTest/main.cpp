#include "gen_qapplication.h"
#include "gen_qwidget.h"

int main(int argc, char** argv) {

    QApplication* app = QApplication_new(&argc, argv);

    QWidget* w = QWidget_new2();//重载造成的绑定真是一言难尽，参数不明确
    QWidget_Show(w);

    return QApplication_Exec(app);
}