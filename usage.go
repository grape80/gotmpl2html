package gotmpl2html

const usage string = `
Usage:
    gotmpl2html [commands] [options] <argument>

Description:
    gotmpl2html is a cli tool to convert gotemplate to html.

Commands:
    version    Print version.
    help       Print usage.

Options:
    --basedir <directory>    Specify the base directory of the root relative path.

Argument:
    Specify a gotemplate file.

Examples:
    gotmpl2html input.gotmpl > output.html
    gotmpl2html version
    gotmpl2html help

Authors:
    [hachirog](https://github.com/hachirog)

`
