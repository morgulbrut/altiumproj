 █████╗               ██╗      ██╗                                            
██╔══██╗██╗  ████████╗██║██╗   ██║███╗   ███╗██████╗ ██████╗  ██████╗      ██╗
██║  ██║██║  ╚══██╔══╝██║██║   ██║████╗ ████║██╔══██╗██╔══██╗██╔═══██╗     ██║
███████║██║     ██║   ██║██║   ██║██╔████╔██║██████╔╝██████╔╝██║   ██║     ██║
██╔══██║██║     ██║   ██║██║   ██║██║╚██╔╝██║██╔═══╝ ██╔══██╗██║   ██║     ██║
██║  ██║███████╗██║   ██║╚██████╔╝██║ ╚═╝ ██║██║     ██║  ██║╚██████╔╝██   ██║
██║  ╚═╝╚══════╝██║   ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝     ██║  ╚═╝ ╚═════╝ ╚█████╔╝
╚═╝             ╚═╝                                  ╚═╝               ╚════╝ 

Command line tool to intialize and rename Altium Designer projects.

## Templates
At the moment there are templates for 2, 4 and 6 Layer. There are zips called twolayer.zip, fourlayer.zip, sixlayer.zip.

You could use your own templates, zip them and rename them to fit the scheme.

Theres a `templates_src` containing the Altium Designer template projects as well as a workspace for them.

## Building
The template zips can be embedded into the binary, for convenience and stuff. I use packr, so the instead of `go build` you should use `packr2 build`. Be sure your in the right directory, if using VS code, open the directory directly, don't cd into it in the terminal.

Or simply run the `build.py` script.

## Usage
Simple as `altiumproj init twolayer <PROJECTNAME>`, `altiumproj init fourlayer <PROJECTNAME>`
