package annex

type FileType = int

const (
	UnknownFileType  FileType = iota //未知文件
	ImageFileType                    // 图片文件
	VideoFileType                    // 视频文件
	AudioFileType                    // 音频文件
	DocumentFileType                 // 文档文件
	ZipFileType                      // 压缩包文件
	DesignFileType                   // 设计文件
)
