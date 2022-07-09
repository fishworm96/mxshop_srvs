package model

// 在开发中，尽量不要设置null
// https://zhuanlan.zhihu.com/p/73997266
// 商品表
type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:'商品名称'"`
	// 在类型转换中经常要用int32或者int64为了方便定义为int32
	ParentCategoryID int32     `gorm:"comment:'自关联id'"`
	ParentCategory   *Category `gorm:"comment:'自关联商品'"`
	Level            int32     `gorm:"type:int;not null;default:1;comment:'1代表1级类目，2代表二级类目，3代表三级类目'"`
	IsTab            bool      `gorm:"default:false;not null;comment:'是否在tab栏展示'"`
}

// 品牌表
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:'品牌名称'"`
	Logo string `gorm:"type:varchar(200);default:'';not null;comment:'品牌logo图片'"`
}

// 品牌和商品关联表
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"` // 联合唯一索引，解决一个数据重复添加2次
	Category Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// 轮播图表
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url string `gorm:"type:varcahr(200);not null"`
	Index int32 `gorm:"type:int;default:1;not null"`
}