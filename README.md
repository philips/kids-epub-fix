## Kid epub Fix

Many early reader kid books have pages that are single large image that has been hand drawn and lettered. The publishers then attempt to use a number of hacky CSS/HTML solutions in their epubs to overlay and then hide text on top. 

This tool strips all of the complex CSS and HTML away and just displays the images one per page.

Usage:

```
go run kid-epub-fix pug.epub
```

This will output pub-fix.epub
