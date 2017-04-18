# Forgetful

*A little Go program that will help you make sure your code is ready for production*

## How it Works
The program will recursively scan through files and folders from a specified starting point and look for "// TODO" comments by default, or something else if you specify it in the CLI.

## Installation
    go get github.com/multapply/forgetful

## Usage
    go run main.go <startpath> <optional target> <optional -C flag>

The `-C` flag means "case-sensitive" and will treat your target as such, making sure things like "SKIP" aren't returned if your specified target is "skip"

## Todo:
 + Make installation better lol
 + Additional CLI argument for directories you want to ignore
 + Colorful output