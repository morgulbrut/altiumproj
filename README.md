# Altiumproj

Command line tool to intialize and rename Altium Designer projects.

## Templates
By default there are a twolayer, a fourlayer and a sixlayer template installed.

Theres a `templates_src` containing the Altium Designer template projects as well as a workspace for them.

### Own templates

Just put a cleaned up (deleted History and stuff) Altium Designer project into the `template_src` folder and build it using `build.py`.


## Building
The template zips can be embedded into the binary, for convenience and stuff. I use packr, so the instead of `go build` you should use `packr2 build`. Be sure your in the right directory, if using VS code, open the directory directly, don't cd into it in the terminal.

Or simply run the `build.py` script.

## Usage
Simple as `altiumproj init <TEMPLATENAME> <PROJECTNAME>`, `altiumproj rename <OLDNAME> <NEWNAME>`
