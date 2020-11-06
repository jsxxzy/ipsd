build-macos:
	@echo "build macos"
	go build -o build/ipsd -ldflags "-s -w" .
	upx -9 build/ipsd

clean:
	@echo "clean build"
	rm -rf build