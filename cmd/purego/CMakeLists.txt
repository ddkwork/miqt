cmake_minimum_required(VERSION 3.30)
project(purego)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

if(MSVC)
        set(CMAKE_BUILD_TYPE Release)
        set(CMAKE_MSVC_RUNTIME_LIBRARY MultiThreaded)
endif()
set(CMAKE_PREFIX_PATH "D:/111111111111111111111111/!!!!!!!!!!!!!!!!!!/qt6/sdk/qt_static/lib/cmake")



find_package(Qt6 COMPONENTS
        Core
        Gui
        Widgets
        REQUIRED)

add_library(purego  SHARED
        binding.cpp
#        bindings_test/testapp.cpp
)

target_link_libraries(purego
        Qt::Core
        Qt::Gui
        Qt::Widgets
)

