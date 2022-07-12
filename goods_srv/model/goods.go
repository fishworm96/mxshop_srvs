package model

// 在开发中，尽量不要设置null
// https://zhuanlan.zhihu.com/p/73997266
// 分类表
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

// 品牌和分类关联表
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
	Image string `gorm:"type:varchar(200);not null;"`
	Url string `gorm:"type:varchar(200);not null"`
	Index int32 `gorm:"type:int;default:1;not null"`
}

// 商品表
type Goods struct {
	BaseModel

	// 这里的唯一键是由商品id或商品名称来确定
	CategoryID int32 `gorm:"type:int;not null;comment:'分类id'"`
	Category Category
	BrandsID int32 `gorm:"type:int;not null;comment:'商品id'"`
	Brands Brands

	OnSale bool `gorm:"default:false;not null;comment:'是否上架'"`
	ShipFree bool `gorm:"default:false;not null;comment:'是否免运费'"`
	IsNew bool `gorm:"default:false;not null;comment:'是否新品'"`
	IsHot bool `gorm:"default:false;not null;comment:'是否热门商品'"`

	Name string `gorm:"type:varchar(50);not null;comment:'商品名称'"`
	GoodsSn string `gorm:"type:varchar(50);not null;comment:'商品编号'"`
	ClickNum int32 `gorm:"type:int;default:0;not null;comment:'点击数'"`
	SoldNum int32 `gorm:"type:int;default:0;not null;comment:'购买数'"`
	FavNum int32 `gorm:"type:int;default:0;not null;comment:'收藏数'"`
	MarketPrice float32 `gorm:"not null;comment:'市场价'"`
	ShopPrice float32 `gorm:"not null;comment:'本店价格'"`
	GoodsBrief string `gorm:"type:varchar(100);not null;comment:'商品简介'"`
	// 另外再建一张表存储的话，通过join后会有性能问题，所以使用gorm的自定义类型来出来。
	Images GormList `gorm:"type:varchar(1000);not null;comment:'商品展示图片'"`
	DescImages GormList `gorm:"type:varchar(1000);not null;comment:'商品内容图片'"`
	GoodsFrontImage string `gorm:"type:varchar(200);not null;comment:'商品封面图片'"`
}