package config

import "ice-mall/schema/annex"

// FileMaps 文件类型
var FileMaps = map[int][]string{
	annex.ImageFileType: {
		"image/png",     // .png
		"image/gif",     // .gif
		"image/webp",    // .webp
		"image/jpeg",    // .jpeg
		"image/svg+xml", // .svg
	},
	annex.VideoFileType: {
		"video/mp4",   // .mp4
		"video/x-flv", // .flv
	},
	annex.AudioFileType: {
		"audio/mpeg", // .mp3
		"audio/wave", // .wav
		"audio/mp4",  // .m4a
	},
	annex.DocumentFileType: {
		"application/pdf",    // .pdf
		"application/msword", // .doc
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document", // .docx
		"application/vnd.ms-excel", // .xls
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", // .xlsx
		"text/csv", // .csv
		"application/vnd.openxmlformats-officedocument.presentationml.presentation", // .pptx,
		"text/plain",      // .txt
		"application/xml", // .xml
	},
	annex.ZipFileType: {
		"application/zip",     // .zip
		"application/vnd.rar", // .rar
	},
	annex.DesignFileType: {
		"image/vnd.adobe.photoshop", // .psd
	},
}
