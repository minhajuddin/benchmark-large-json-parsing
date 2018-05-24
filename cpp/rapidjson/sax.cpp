#include "rapidjson/reader.h"
#include "rapidjson/filereadstream.h"

using namespace rapidjson;


struct Handler : public BaseReaderHandler<UTF8<>, Handler> {};

int main(int argc, char *argv[]) {
	Reader reader;
	char readBuffer[65536];
	FileReadStream is(stdin, readBuffer, sizeof(readBuffer));
	Handler handler;
	if(!reader.Parse(is, handler)) return 1;
	return 0;
}

