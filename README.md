## Kid epub Fix

Many early reader kid books have pages that are single large image that has been hand drawn and lettered. The publishers then attempt to use a number of hacky CSS/HTML solutions in their epubs to overlay and then hide text on top. 

This tool strips all of the complex CSS and HTML away and just displays the images one per page.

Usage:

```
go run main.go pug.epub
```

This will output pub-fix.epub

## Tested Files

I tested this on 7 epubs, listed below, all purchased from Kobo.com

```
Diary of a Pug #1_ Pug Blasts Off! - Kyla May-fix.epub
Diary of a Pug_ Paws for a Cause - Kyla May-fix.epub
Dragon #1_ A Friend for Dragon - Dav Pilkey-fix.epub
Hello, Hedgehog!_ Do You Like My Bike_ - Norm Feuti-fix.epub
Pug's Snow Day_ A Branches Book (Diary of a Pug #2) - Kyla May-fix.epub
Racing Ace #1_ Drive it! Fix it! - Larry Dane Brimner-fix.epub
Unicorn and Yeti_ Sparkly New Friends - Heather Ayris Burnell-fix.epub
```

**Disclosure** This project was vibe coded in a late evening while being disapointed by koreaders limited rendering engine AND publishers horrible epub formatting. I love koreader and understand the contraints of an embedded CSS/HTML rendering engine- it can't parse all the world's crazy stuff.
