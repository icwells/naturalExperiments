# Identifies species that have recently diverged and have different cancer rates  

Copyright 2020 by Shawn Rupp

1. [Description](#Description)
2. [Installation](#Installation)  
3. [Usage](#Usage)  

## Description  
naturalExperiments can be used to identify pairs of species that have recently diverged but have different cancer rates.  
It takes a newick tree file (must have branch lengths) and the cancer rate output of the Comparative Oncology Database as input.

## Installation  

### Dependencies  
[Go version 1.11 or higher](https://golang.org/doc/install)  

### Installing Go and Setting Paths  
Go requires a GOPATH environment variable to set to install packages, an compOncDB requires the GOBIN variable to be set as well.  
Follow the directions [here](https://github.com/golang/go/wiki/SettingGOPATH) to set your GOPATH. Before you close your .bashrc or 
similar file, add the following lines after you deifne you GOPATH:  

	export GOBIN=$GOPATH/bin  
	export PATH=$PATH:$GOBIN   

### Download  
Download the repository into correct Go src directory (required for package imports):  

	cd $GOPATH/src
	mkdir -p github.com/icwells/
	cd github.com/icwells/
	git clone https://github.com/icwells/naturalExperiments.git  

### Compiling scripts:
Any missing Go packages will be downloaded and installed when running install.sh.  

	cd naturalExperiments/  
	./install.sh  

### Testing the Installation  
Run the following in a terminal:

	./test.sh all

All of the output from the test scripts should begin with "ok".  

## Commands  
usage: naturalExperiments --infile=INFILE --treefile=TREEFILE [<flags>]

	--help					Show context-sensitive help.
	-i, --infile=INFILE		Path to input cancer rates file.
	--malignant				Examine malignancy rates (examines neoplasia rate by default).
	--max=10.0				The maximum divergeance allowed to compare species.
	--min=0.2				The minimum difference in cancer rates to report results.
	--records=50			The minimum number of records for a species required for examination.
	-o, --outfile="nil"		Name of output file.
	-t, --treefile=TREEFILE	Path to newick tree file.
