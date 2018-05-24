#include "rapidjson/filereadstream.h"
#include "rapidjson/document.h"

using namespace rapidjson;



int main(int argc, char *argv[]) {
	char readBuffer[65536];
	FileReadStream is(stdin, readBuffer, sizeof(readBuffer));
	if(!document.Parse(is)) return 1;
	if(document.size() < 1) return 2;
	return 0;
}

