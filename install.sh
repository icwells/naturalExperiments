#!/bin/bash

##############################################################################
# This script will install scripts for the naturalExperiment package.
# 
# Required programs:	Go 1.11+
##############################################################################

MAIN="naturalExperiments"

installMain () {
	# compOncDB 
	echo "Building $MAIN..."
	go build -i -o $GOBIN/$MAIN src/*.go
	echo ""
}

echo ""
echo "Preparing compOncDB package..."
echo "GOPATH identified as $GOPATH"
echo ""

installMain

echo "Finished"
echo ""
