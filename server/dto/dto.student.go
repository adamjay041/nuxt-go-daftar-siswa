package dto

type CreateRequestStudent struct {
	Npm      string `json:"npm" form:"npm" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Kelas    string `json:"kelas" form:"kelas" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Periode  string `json:"periode" form:"periode" binding:"required"`
	L_pict   string `json:"l_pict" form:"l_pict"`
	Email    string `json:"email" form:"email" binding:"required"`
	Status   string `json:"status" form:"status" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
	Major    string `json:"Major" gorm:"type:varchar;not null"`
}

type EditRequestStudent struct {
	// Npm      string `json:"npm" form:"npm" binding:"required"`
	Name     string `json:"name" form:"name"`
	Kelas    string `json:"kelas" form:"kelas"`
	Password string `json:"password" form:"password"`
	// Periode  string `json:"periode" form:"periode"`
	L_pict string `json:"l_pict" form:"l_pict"`
	Email  string `json:"email" form:"email"`
	// Status   string `json:"status" form:"status"`
	// Role     string `json:"role" form:"role"`
	Major string `json:"Major" gorm:"type:varchar;not null"`
}

type Login struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
