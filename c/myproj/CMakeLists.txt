cmake_minimum_required(VERSION 3.16)
project(myproj C)

set(CMAKE_C_STANDARD 11)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/include)

add_library(mylib
  src/math.c
  src/util.c
  src/internal/private_util.c
)

add_executable(myapp
  src/main.c
)

target_link_libraries(myapp PRIVATE mylib)

enable_testing()
add_subdirectory(tests)
