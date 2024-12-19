cmake_minimum_required(VERSION 3.30)
project(demo)

set(CMAKE_CXX_STANDARD 23)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)
if(MSVC)
    set(CMAKE_BUILD_TYPE Release)
    set(CMAKE_MSVC_RUNTIME_LIBRARY MultiThreaded)
endif()
set(CMAKE_PREFIX_PATH "D:/qt6/qt_static/lib/cmake")

find_package(Qt6 COMPONENTS
        Core
        Gui
        Widgets
        REQUIRED)

add_executable(demo main.cpp)
target_link_libraries(demo
        Qt::Core
        Qt::Gui        Qt::Widgets
)
