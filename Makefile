all:
	cgogen alac.yml

clean:
	rm -f alac/cgo_helpers.go alac/cgo_helpers.h alac/doc.go alac/types.go
	rm -f alac/alac.go

test:
	cd alac && go build