This reads the brightness off the filesystem, adjusts it as needed, and
writes it back out.

## installation
```
go install github.com/joshgordon/backlight_brightness
sudo cp $GOPATH/brightness /usr/local/bin/
sudo chmod o+s /usr/local/bin/brightness
```
You may have to edit the source and specify the filename to the brightness
file.
