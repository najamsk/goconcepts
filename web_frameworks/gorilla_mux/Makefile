ver := najamsk/gorilla_mux_api_8000:1.0.3
lat := najamsk/gorilla_mux_api_8000:latest
img := najamsk/gorilla_mux_api_8000

dexe:
	GOOS=linux GOARCH=amd64 go build -o main 1.go

dimg:
	docker build -t ${ver} -t ${lat} .

drun:
	docker run -p 8000:8000 ${lat}

dpush:
	docker image push  ${lat}
	docker image push  ${ver}
