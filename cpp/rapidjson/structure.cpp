#include "rapidjson/filereadstream.h"
#include "rapidjson/document.h"

using namespace rapidjson;



int main(int argc, char *argv[]) {
	char readBuffer[65536];
	FileReadStream is(stdin, readBuffer, sizeof(readBuffer));
	if(document.Parse<kParseValidateEncodingFlag>(is).HasParseError()) return 1;
	return 0;
}

