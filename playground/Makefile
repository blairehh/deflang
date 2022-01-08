install:
	rm -f ./def
	rm -f ./def-cmd
	rm -f ./deflang
	rm -f ./deflang-cmd
	rm -f ./defvar
	cd ../defvar && go build && mv defvar ../playground
	cd ../deflang && go build && mv deflang ../playground/deflang-cmd
	cd ../def && go build && mv def ../playground/def-cmd
	cp ../deflang/deflang.sh deflang