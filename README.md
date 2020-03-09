Install
=======

You need [Go](https://golang.org/).

Then call

    go get -u github.com/gonutz/command_line_programs/...

The option `-u` is to pull the latest online version.


Programs
========

```
non_empty - Print only lines that are not empty.
            Lines containing spaces are not empty.

trim_space - Print every line with starting and trailing white space removed.

sort_by_number_at_line_start - Extract the longest possible integer number (base 10) from the start of a
                               line and sort lines by it.
                               0 is assumed for lines that do not start with a number.

count_lines - Print the number of lines. Ignores the last line if it is empty.
```
