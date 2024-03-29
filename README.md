Install
=======

You need [Go](https://golang.org/).

Then call

    go install github.com/gonutz/command_line_programs/...@latest

Programs
========

```
append_to_lines - append the given string to all lines in standard input and print
                  the new lines to standard output.

base64 - Convert standard input to base 64 and write the result to standard output.

count - Print the number of bytes/lines/letters (use one of these as the argument).
        Reads standard input and writes the number to standard output. In case of
        an error it writes 0 to standard output and the error to standard error.

count_lines - Print the number of lines. Ignores the last line if it is empty.

dedup - De-duplicate lines: print lines from standard input only once, the next time
        a line is encountered it is not printed again.

field - Select a space separated field from standard input by index.

first - Output the first N bytes of the input. Outputs all if no N is given or if the
        argument is not a number. Call like this:
		first 5 < file
		first byte < file
		first 5 bytes < file
		first line < file
		first 5 lines < file

last - Output the last N lines of standard input. Call like this:
       last line < file
	   last 5 lines < file

line - Print the line at a given index, starting at 0. line 0 prints the first line
       from standard input, line 1 prints the second line, line 2 the third, etc.

line_index - Print the index of the first matching line in standard input. Prints -1
             if the given line cannot be found on standard input.

non_empty - Print only lines that are not empty.
            Lines containing spaces are not empty.

png - Convert a JPEG/BMP/PNG/GIF image from standard input to PNG and write the
      result to standard output.

replace_all - replaces the first argument with the second for the standard input.

skip - discard the first N bytes of the input and output the rest.

sleep - Wait for the given duration.

sort_by_number_at_line_start - Extract the longest possible integer number (base 10)
                               from the start of a line and sort lines by it.
                               0 is assumed for lines that do not start with a number.

timer - Simple timing tool, press ENTER to take times, quit with Ctrl+C.

trim_space - Print every line with starting and trailing white space removed.

unzip - Unpack a zip archive.

zip - Pack a file or folder into a zip archive.
```
