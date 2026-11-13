package product

type ProductSize struct {
	ID           uint     `gorm:"primaryKey;autoIncrement"`
	ProductID    uint     `gorm:"not null;uniqueIndex:idx_product_size"`
	Label        string   `gorm:"type:text;not null;uniqueIndex:idx_product_size"`
	SortOrder    int      `gorm:"not null"`
	ChestCM      *float64 `gorm:"type:numeric(5,1)"`
	ShoulderCM   *float64 `gorm:"type:numeric(5,1)"`
	WaistCM      *float64 `gorm:"type:numeric(5,1)"`
	HipCM        *float64 `gorm:"type:numeric(5,1)"`
	LengthCM     *float64 `gorm:"type:numeric(5,1)"`
	InseamCM     *float64 `gorm:"type:numeric(5,1)"`
	FootLengthCM *float64 `gorm:"type:numeric(4,1)"`

	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	//  SKUs    []SKU   `gorm:"foreignKey:SizeID;constraint:OnDelete:CASCADE"`
}

func (ProductSize) TableName() string {
	return "product_sizes"
}
