// Teal.Finance/Server is an opinionated boilerplate API and website server.
// Copyright (C) 2021 Teal.Finance contributors
//
// This file is part of Teal.Finance/Server, licensed under LGPL-3.0-or-later.
//
// Teal.Finance/Server is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License
// either version 3 of the License, or (at your option) any later version.
//
// Teal.Finance/Server is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.

package fileserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/teal-finance/server/reserr"
)

type FileServer struct {
	Dir    string
	ResErr reserr.ResErr
}

// ServeFile handles one specific file (and its specific Content-Type).
func (fs FileServer) ServeFile(urlPath, contentType string) func(w http.ResponseWriter, r *http.Request) {
	absPath := path.Join(fs.Dir, urlPath)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		fs.send(w, r, absPath)
	}
}

// ServeDir handles the static files using the same Content-Type.
func (fs FileServer) ServeDir(contentType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if validPath(w, r) {
			w.Header().Set("Content-Type", contentType)

			absPath := path.Join(fs.Dir, r.URL.Path)
			fs.send(w, r, absPath)
		}
	}
}

// ServeImages detects the Content-Type depending on the image extension.
func (fs FileServer) ServeImages() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if validPath(w, r) {
			absPath, contentType := fs.imagePathAndType(r)
			if contentType != "" {
				w.Header().Set("Content-Type", contentType)
			}

			fs.send(w, r, absPath)
		}
	}
}

// validPath returns a HTTP error if the path is invalid.
func validPath(w http.ResponseWriter, r *http.Request) bool {
	if strings.Contains(r.URL.Path, "..") {
		log.Print("WARN Fileserver: reject path with '..' ", r.URL.Path)
		reserr.Write(w, r, http.StatusBadRequest, "Invalid URL Path Containing '..'")

		return false
	}

	return true
}

func (fs FileServer) send(w http.ResponseWriter, r *http.Request, absPath string) {
	var file *os.File

	// if client (browser) supports Brotli and the *.br file is present
	// => send the *.br file
	if strings.Contains(r.Header.Get("Accept-Encoding"), "br") {
		brotli := absPath + ".br"

		f, err := os.Open(brotli)
		if err == nil {
			absPath = brotli

			w.Header().Set("Content-Encoding", "br")
		}

		file = f
	}

	if file == nil {
		f, err := os.Open(absPath)
		if err != nil {
			fs.ResErr.Write(w, r, http.StatusNotFound, "Page not found")
			log.Printf("WARN Fileserver: Open(%v) %v", absPath, err)

			return
		}

		file = f
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Print("WARN Fileserver: Close() ", err)
		}
	}()

	fi, err := file.Stat()
	if err != nil {
		log.Printf("WARN Fileserver: Stat(%v) %v", absPath, err)
		fs.ResErr.Write(w, r, http.StatusInternalServerError, "Internal Server Error")

		return
	}

	w.Header().Set("Content-Length", strconv.FormatInt(fi.Size(), 10))
	w.Header().Set("Last-Modified", fi.ModTime().UTC().Format(http.TimeFormat))
	// We do not manage PartialContent because too much stuff
	// to handle the headers Range If-Range Etag and Content-Range.

	if n, err := io.Copy(w, file); err != nil {
		log.Printf("WARN Fileserver: Copy(%v) %v", absPath, err)
	} else {
		log.Printf("Fileserver sent %v %v", absPath, IEC64(n))
	}
}

// IEC64 converts bytes into KiB (1024 bytes), MiB, GiB…
// as defined within the International System of Quantities (ISQ)
// standardized by the ISO/IEC 80000 and published in 2008.
func IEC64(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// imagePathAndType returns the path/filename and the Content-Type of the image.
// If the client (browser) supports AVIF, imagePathAndType replaces the requested image by the AVIF one.
func (fs FileServer) imagePathAndType(r *http.Request) (absPath, contentType string) {
	extPos := extIndex(r.URL.Path)

	// We only check the first Header "Accept":
	// We do not care to miss an "image/avif" within the second Header "Accept",
	// because we do not break anything: we send the image requested by the client.
	scheme := r.Header.Get("Accept")

	// We perform a stupid search to be fast,
	// but we hope there is no Content-Type such as "image/avifuck"
	const avifContentType = "image/avif"
	if strings.Contains(scheme, avifContentType) {
		avifPath := r.URL.Path[:extPos] + "avif"
		absPath := path.Join(fs.Dir, avifPath)

		_, err := os.Stat(absPath)
		if err == nil {
			return absPath, avifContentType
		}

		log.Printf("Fileserver supports Content-Type=%q but cannot access %q: %v",
			avifContentType, absPath, err)
	}

	absPath = path.Join(fs.Dir, r.URL.Path)

	ext := r.URL.Path[extPos:]
	contentType = imageContentType(ext)

	return absPath, contentType
}

// extIndex returns the position of the extension within the the urlPath.
// If no dot, returns the ending position.
func extIndex(urlPath string) int {
	for i := len(urlPath) - 1; i >= 0 && urlPath[i] != '/'; i-- {
		if urlPath[i] == '.' {
			return i + 1
		}
	}

	return len(urlPath)
}

// imageContentType determines the Content-Type depending on the file extension.
func imageContentType(ext string) (contentType string) {
	// Only the most popular image extensions
	switch ext {
	case "png":
		return "image/png"
	case "jpg":
		return "image/jpeg"
	case "svg":
		return "image/svg+xml"
	default:
		log.Print("Fileserver does not support image extension: ", ext)

		return ""
	}
}

// Extension  MIME types
// ---------  --------------------------------
//  "html":   "text/html; charset=utf-8",
//  "css":    "text/css; charset=utf-8",
//  "js":     "text/javascript; charset=utf-8",
//  "woff2":  "font/woff2",
//  "csv":    "text/csv",
//  "json":   "application/json",
//  "md":     "text/markdown",
//  "pdf":    "application/pdf",
//  "xml":    "text/xml; charset=utf-8",
//  "yaml":   "text/x-yaml",
//  "svg":    "image/svg+xml",
//  "avif":   "image/avif",
//  "gif":    "image/gif",
//  "webp":   "image/webp",
//  "png":    "image/png",
//  "jpg":    "image/jpeg",
