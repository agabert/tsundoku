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
	"fmt"
	"github.com/Sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var sourceDirectory = "/space"
var destinationDirectory = "/space/tsundoku/trunk"

func main() {
	fileCache := make(map[string]string)

	if _, isDebug := os.LookupEnv("DEBUG"); isDebug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug logging enabled.")
	}

	logrus.Info("main(): starting.")

	var fileWalker = func(fullName string, fileInfo os.FileInfo, fileError error) error {
		if fileError != nil {
			logrus.Fatal(fileError)
			return fileError
		}

		if fileInfo.IsDir() {
			logrus.Debug("Looking at directory: ", fullName)
		} else {
			/*
			 * (I) check if the file is a normal file
			 */
			sourceFileStat, err := os.Stat(fullName)
			if err != nil {
				return nil
			}
			if !sourceFileStat.Mode().IsRegular() {
				return nil
			}
			sourceLinkStat, err := os.Lstat(fullName)
			if sourceLinkStat.Mode()&os.ModeSymlink != 0 {
				return nil
			}

			/*
			 * (II) check if the file has an extension that we are looking for
			 */
			extension := strings.TrimPrefix(strings.ToLower(filepath.Ext(fullName)), ".")

			if skipExtension(extension) {
				logrus.Debug("Skipping extension: ", fullName)
				return nil
			}
			if skipPath(fullName) {
				logrus.Debug("Skipping path: ", fullName)
				return nil
			}

			/*
			 * (III) check for minimum and maximum file size
			 */
			fhandle, err := os.Stat(fullName)
			if err != nil {
				return nil
			}
			if fhandle.Size() < 40000 {
				return nil
			}

			if fhandle.Size() > 4000000000 {
				return nil
			}

			/*
			 * (IV) prepare the fingerprint by reading the file once
			 */
			fingerPrint := getFileChecksum(fullName)
			if fingerPrint == "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
				// empty file
				return nil
			}
			if fingerPrint == "a8b7b3b3de85c0309fa72d15a91876e8b4638e8c3412bc783d94cf422025917d" {
				// NSFW ;)
				return nil
			}

			/*
			 * (V) Add the file to the hash map with fingerprint as key
			 */
			logrus.Debug("Adding/updating file in memory cache: [", fullName, "] as [", fingerPrint, ".", extension, "]")
			fileCache[fmt.Sprintf("%s.%s", fingerPrint, extension)] = fullName
		}

		return nil
	}

	if err := filepath.Walk(sourceDirectory, fileWalker); err != nil {
		logrus.Fatal(err)
	}

	if _, isDebug := os.LookupEnv("ONLYSCAN"); isDebug {
		logrus.Info("main(): scan finished.")
	} else {
		logrus.Info("main(): writing files.")
		for k, v := range fileCache {
			dstFile := fmt.Sprintf("%s/%s", destinationDirectory, k)
			sourceFileHandle, err := os.Open(v)
			if err == nil {
				destinationFileHandle, err := os.Create(dstFile)
				if err == nil {
					logrus.Debug("Creating: [", dstFile, "] from [", v, "]")
					io.Copy(destinationFileHandle, sourceFileHandle)
					destinationFileHandle.Close()
					sourceFileHandle.Close()
				}
			}
		}
	}

	logrus.Info("main(): exiting.")
	os.Exit(0)
}
