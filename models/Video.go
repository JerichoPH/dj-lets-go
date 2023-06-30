package models

type (
	// VideoFile 视频模型
	VideoFile struct {
		GormModel
		Title             string               `gorm:"type:varchar(256);unique;not null;comment:视频标题;" json:"title"`
		ArtName           string               `gorm:"type:varchar(256);not null;default:'';comment:视频演员;" json:"art_name"`
		OriginalFilename  string               `gorm:"type:varchar(256);not null;comment:原始文件名;" json:"original_filename"`
		OriginalExtension string               `gorm:"varchar(64);not null;default:'';comment:原始扩展名;" json:"original_extension"`
		OriginalMime      string               `gorm:"varchar(256);not null;default:'';comment:原始文件类型;" json:"original_mime"`
		Filename          string               `gorm:"type:text;comment:文件存放路径;" json:"filename"`
		UploaderUuid      string               `gorm:"type:varchar(36);index;not null;default:'';comment:所属上传人uuid;" json:"uploader_uuid"`
		Uploader          AuthorizationAccount `gorm:"foreignKey:uploader_uuid;references:uuid;comment:所属上传人;" json:"uploader"`
		AlbumUuid         string               `gorm:"type:varchar(36);index;not null;default:'';comment:所属专辑uuid;" json:"album_uuid"`
		Album             *VideoAlbum          `gorm:"foreignKey:uuid;references:album_uuid;comment:所属专辑;" json:"album"`
		Tags              []*VideoTag          `gorm:"many2many:video_pivot_post_and_tags;foreignKey:uuid;joinForeignKey:video_tag_uuid;references:uuid;joinReferences:video_file_uuid;" json:"tags"`
	}

	// VideoAlbum 视频专辑
	VideoAlbum struct {
		GormModel
		Name  string       `gorm:"type:varchar(64);unique;not null;comment:视频专辑;" json:"name"`
		Files []*VideoFile `gorm:"foreignKey:uuid;references:album_uuid;comment:相关视频;" json:"files"`
	}

	// VideoTag 视频标签
	VideoTag struct {
		GormModel
		Name  string       `gorm:"type:varchar(64);unique;not null;comment:视频标签;" json:"name"`
		Files []*VideoFile `gorm:"many2many:video_pivot_post_and_tags;foreignKey:uuid;joinForeignKey:video_file_uuid;references:uuid;joinReferences:video_tag_uuid;" json:"files"`
	}
)

// TableName 视频表名称
func (receiver VideoFile) TableName() string {
	return "video_files"
}

// TableName 视频专辑表名称
func (receiver VideoAlbum) TableName() string {
	return "video_albums"
}

// TableName 视频标签表名称
func (receiver VideoTag) TableName() string {
	return "video_tags"
}
