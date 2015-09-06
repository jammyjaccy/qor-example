package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/media_library"
	"github.com/qor/slug"
)

type Product struct {
	gorm.Model
	Name           string
	NameWithSlug   slug.Slug
	Code           string
	Price          float32
	MadeCountry    string
	Description    string `sql:"size:2000"`
	Images         []ProductImage
	ColorVariation []ColorVariation
	Category       Category
}

type ProductImage struct {
	gorm.Model
	Image media_library.FileSystem
}

type ColorVariation struct {
	gorm.Model
	ProductID      uint
	Color          Color
	SizeVariations []SizeVariation
}

type SizeVariation struct {
	gorm.Model
	ColorVariationID  uint
	Size              Size
	AvailableQuantity uint
}
