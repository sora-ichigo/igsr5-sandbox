add_executable(test_math test_math.c)
target_link_libraries(test_math PRIVATE mylib)

add_test(NAME math_test COMMAND test_math)
