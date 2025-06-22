# notes

A notes organizer.

## Installation

### go install

```sh
go install lesiw.io/notes@latest
```

### curl

```sh
curl lesiw.io/notes | sh
```

## Usage

`notes` opens up the `NOTES` file in the local directory. If `NOTES` is not
present in the current directory, `notes` will continue to search directories
above for the nearest `NOTES` file, if present.

`notes -i` or `notes --init` will force the creation of a `NOTES` file in the
current directory.

Notes are edited using the standard `EDITOR` variable. To override this, set
`NOTESEDITOR`.

Set `NOTESOVERLAY` to one or more overlay directories if you want to keep your
notes in a separate space. The format is `upper:lower::upper2:lower2`, where
the upper directory contains the `NOTES` files and the lower directory is the
one being overlaid on top of.
