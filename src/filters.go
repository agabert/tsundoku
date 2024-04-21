/*
   Copyright 2024 Alexander Gabert

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"strings"
)

func skipPath(filePath string) bool {
	i := 0

	for range embargo {
		if strings.Contains(strings.ToLower(filePath), strings.ToLower(embargo[i])) {
			return true
		}
		i++
	}

	return false
}

// true means we do not want the file
func skipExtension(extension string) bool {

	if extension == "" {
		return true
	}

	if extension == "jpeg" {
		extension = "jpg"
	}

	if extension == "tiff" {
		extension = "tif"
	}

	switch strings.ToLower(extension) {
	case
		"m4v",
		"wmv",
		"heic",
		"tga",
		"tif",
		"tiff",
		"gif",
		"png",
		"jpg",
		"jpeg",
		"bmp",
		"xcf",
		"psd",
		"pcx",
		"rle",
		"mp3",
		"wav",
		"ogg",
		"avi",
		"flv",
		"mov",
		"mpg",
		"mpeg",
		"mp4":
		return false
	}

	return true
}
