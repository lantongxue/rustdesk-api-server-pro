main_output=build
frontend=soybean-admin
frontend_dist=${frontend}/dist

build: clean
	go build -C backend -o ../${main_output}/
	cd ${frontend} && pnpm build && cp -R dist ${main_output}/

clean:
	rm -rf ${main_output}
	rm -rf ${frontend_dist}
	go clean
