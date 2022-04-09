package models

type Music struct {
	ID         uint
	AlbumTypes []uint
	Name       string
	Info       string
	Image      string
	FilePath   string
	FileExt    string
}
