
# atool/varsion

## atools

atools are little cross-platform programs that solve try to solve a single problem.

### other atools?

- varsion - currently this is all I got

## varsion

varsion is a go program that attempts to handle source code versioning universally by placing a VERSION file at the root of the source code.
it uses this VERSION file along with a .atoolconfig.json configuration file to handle [Semantic Versioning](https://semver.org/) for the developers.
the user just provides what segment of the version to change and it will handle everything else.
varsion allows for prefixes and suffixes to be added to the version but other applications that might use the VERSION file might fail due to this.

### download

the only way to download the executable is to download the source code.
the compiled binaries can be seen in the .bin folder where each version will be listed along with a zip of it.
the binaries are the compiled program from my Windows machine using arm64.
other platform binaries will be released later on
you may also compile the source code yourself using the golang compiler.
a real download method will be developed and the binaries will be removed.

### compile

to compile the source code one must have the golang compiler installed.
compile like a normal go program.

### usage

the program can be used in two different ways. 
the first way is by providing the segment to upgrade as a console argument.
example:

- version.exe major
- version.exe ma
- version.exe minor
- version.exe mi
- version.exe patch
- version.exe p

the second way is by being prompted by the application to provide what segment to upgrade.
example:

- version.exe
- segment to change (major, minor, patch): major

## folder structure

the folder and project structure was designed when I was learning many different system design patterns.
this heavily influenced the decisions I made in this project and I may have been too eager to try some.
the structure should be rethought and redesigned when a bigger picture can be seen.

### .bin

the temporary location of the compiled binaries.
each version can be found here alongside a zip version of each.

### data

the data access layer of the application.
this application only needs to access the file system to read the VERSION or config file

### domain

the domain/business logic layer of the application.
this application only needed one object to hold the internal workings of [Semantic Versioning](https://semver.org/).

### service

the service layer of the application.
a configuration service is used to handle configuration throughout the application
a version handler is used to handle the VERSION files as objects and abstract the internal workings of [Semantic Versioning](https://semver.org/).

## files

### VERSION

the file that can be universally used to hold the VERSION of the source code or project.

### .atoolconfig.json

the file used for configuration in the varsion atool.
any user can edit the configuration file and see the effects in the VERSION file.

### VARSION

an alternative to the VERSION file.
some people change the name to VARSION to know that the atool is being used.
the file that can be universally used to hold the VERSION of the source code or project.

## initial project plans

automated-version
autoversion
command line interface tool that will attempt to keep track of the global project version
user must provide what version segment to increase

project version stored in VERSION file
VERSION file can be used for anything to do with version

version type
examples:
version major
version minor
version patch
will increase the given segment by 1 making changes to the others as needed

Semantic Versioning 2.0.0
https://semver.org/
