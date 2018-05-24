cmake_minimum_required(VERSION 3.0)
include("cmake/HunterGate.cmake")
HunterGate(
	URL "https://github.com/ruslo/hunter/archive/v0.21.13.tar.gz"
	SHA1 "331bf20601aed3369e1d00388d08af8169b2e6c2"
	LOCAL
)

project(rapidjson-testing)

hunter_add_package(RapidJSON)
find_package(RapidJSON CONFIG REQUIRED)

add_executable(rapidjson-sax sax.cpp)
target_link_libraries(rapidjson-sax RapidJSON::rapidjson)

add_executable(rapidjson-structure structure.cpp)
target_link_libraries(rapidjson-structure RapidJSON::rapidjson)

