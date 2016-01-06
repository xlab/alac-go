all:
	go build

osx:
	go build
	install_name_tool -change /usr/local/lib/libportaudio.2.dylib libportaudio.2.dylib alac-player
