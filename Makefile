OUT=./bin/dkinit
MAINSRC=./main/dkinit.go

binary:
	go build -o $(OUT) $(MAINSRC)

run: binary
	$(OUT) 

test: binary run

image: linux64 dockerbuild

clean:
	rm $(OUT)
