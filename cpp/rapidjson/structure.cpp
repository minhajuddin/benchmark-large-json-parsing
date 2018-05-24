#include "rapidjson/filereadstream.h"
#include "rapidjson/document.h"

using namespace rapidjson;



int main(int argc, char *argv[]) {
	char readBuffer[65536];
	FileReadStream is(stdin, readBuffer, sizeof(readBuffer));
	Document document;
	if(document.ParseStream<kParseValidateEncodingFlag>(is).HasParseError()) return 1;
	return 0;
}

