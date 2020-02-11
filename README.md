# Altiumproj 


Command line tool to intialize and rename Altium Designer projects.

## Templates
In `build.py` in line 3 and 4 you can set the path to your templates.

Templates should be in folders and as clean as possible (no history and no output folders.) The of those folders will later by the names of you templates.


## Building
The template zips can be embedded into the binary, for convenience and stuff. I use packr, so the instead of `go build` you should use `packr2 build`. Be sure your in the right directory, if using VS code, open the directory directly, don't cd into it in the terminal.

Or simply run the `build.py` script.

## Usage
Simple as `altiumproj init <TEMPLATENAME> <PROJECTNAME>`, `altiumproj rename <OLDNAME> <NEWNAME>`
