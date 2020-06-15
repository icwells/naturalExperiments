#!/bin/bash

##############################################################################
#	Manages tests on naturalExperiments
#
#	Required:	go 1.10+
#
#	Usage:		./test.sh {help/...}
##############################################################################
WD=$(pwd)
SRC="$WD/src/*.go"

whiteBoxTests () {
	echo ""
	echo "Running white box tests..."
	go test $SRC
}

checkSource () {
	# Runs go fmt/vet on source files (vet won't run in loop)
	echo ""
	echo "Running go $1..."
	go $1 $SRC
}

helpText () {
	echo ""
	echo "Runs test scripts for compOncDB."
	echo "Usage: ./test.sh {all/whitebox/blackbox/parse/db/fmt/vet}"
	echo ""
	echo "all		Runs all tests."
	echo "fmt		Runs go fmt on all source files."
	echo "vet		Runs go vet on all source files."
	echo "help		Prints help text."
}

if [ $# -eq 0 ]; then
	helpText
elif [ $1 = "all" ]; then
	whiteBoxTests
elif [ $1 = "fmt" ]; then
	checkSource $1
elif [ $1 = "vet" ]; then
	checkSource $1
elif [ $1 = "help" ]; then
	helpText
else
	helpText
fi
echo ""
