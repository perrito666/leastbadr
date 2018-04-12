package features

type MetaData struct {
// FileName contains the name of the file relative to the folder, should be unique unless deletion happens (ask if name changes sha)
FileName string
// Title contains a user defined title for the picture.
Title string
// Description contains a user defined description for the title.
Description string
// Author contains the name of the author (credit where it's due)
Author string
// Tags contains an arbitrary list of words, in the future it might be useful.
Tags []string
// Exif contains a string representing the data of the Exif
Exif string // Store this as raw string
// Sha contains the sha of the original image, if another image with same name but different sha appears and this one dissapears assume it moved.
Sha string
}