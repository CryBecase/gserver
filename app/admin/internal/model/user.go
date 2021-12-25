package model

import "time"

// User 用户表
type User struct {
	Id        int       `gorm:"column:id;primaryKey" json:"id"`         //
	Username  string    `gorm:"column:username"      json:"username"`   //
	Password  string    `gorm:"column:password"      json:"password"`   //
	Telephone string    `gorm:"column:telephone"     json:"telephone"`  //
	Birthday  string    `gorm:"column:birthday"      json:"birthday"`   //
	Gender    int       `gorm:"column:gender"        json:"gender"`     //
	CreatedAt time.Time `gorm:"column:created_at"    json:"created_at"` //
	UpdatedAt time.Time `gorm:"column:updated_at"    json:"updated_at"` //
}

func (*User) TableName() string {
	return "user"
}

type UserSlice []*User

func (u *UserSlice) IdMap() map[int]*User {
	uni := make(map[int]*User)
	for _, item := range *u {
		uni[item.Id] = item
	}
	return uni
}

func (u *UserSlice) GroupByUsername() map[string]UserSlice {
	res := make(map[string]UserSlice)
	for _, item := range *u {
		res[item.Username] = append(res[item.Username], item)
	}
	return res
}

func (u *UserSlice) GroupByPassword() map[string]UserSlice {
	res := make(map[string]UserSlice)
	for _, item := range *u {
		res[item.Password] = append(res[item.Password], item)
	}
	return res
}

func (u *UserSlice) GroupByTelephone() map[string]UserSlice {
	res := make(map[string]UserSlice)
	for _, item := range *u {
		res[item.Telephone] = append(res[item.Telephone], item)
	}
	return res
}

func (u *UserSlice) GroupByBirthday() map[string]UserSlice {
	res := make(map[string]UserSlice)
	for _, item := range *u {
		res[item.Birthday] = append(res[item.Birthday], item)
	}
	return res
}

func (u *UserSlice) GroupByGender() map[int]UserSlice {
	res := make(map[int]UserSlice)
	for _, item := range *u {
		res[item.Gender] = append(res[item.Gender], item)
	}
	return res
}

func (u *UserSlice) GroupByCreatedAt() map[time.Time]UserSlice {
	res := make(map[time.Time]UserSlice)
	for _, item := range *u {
		res[item.CreatedAt] = append(res[item.CreatedAt], item)
	}
	return res
}

func (u *UserSlice) GroupByUpdatedAt() map[time.Time]UserSlice {
	res := make(map[time.Time]UserSlice)
	for _, item := range *u {
		res[item.UpdatedAt] = append(res[item.UpdatedAt], item)
	}
	return res
}

func (u *UserSlice) PluckId() []int {
	res := make([]int, 0, len(*u))
	for _, item := range *u {
		res = append(res, item.Id)
	}
	return res
}

func (u *UserSlice) PluckUsername() []string {
	res := make([]string, 0, len(*u))
	for _, item := range *u {
		res = append(res, item.Username)
	}
	return res
}

func (u *UserSlice) PluckPassword() []string {
	res := make([]string, 0, len(*u))
	for _, item := range *u {
		res = append(res, item.Password)
	}
	return res
}

func (u *UserSlice) PluckTelephone() []string {
	res := make([]string, 0, len(*u))
	for _, item := range *u {
		res = append(res, item.Telephone)
	}
	return res
}

func (u *UserSlice) PluckBirthday() []string {
	res := make([]string, 0, len(*u))
	for _, item := range *u {
		res = append(res, item.Birthday)
	}
	return res
}

func (u *UserSlice) PluckGender() []int {
	res := make([]int, 0, len(*u))
	for _, item := range *u {
		res = append(res, item.Gender)
	}
	return res
}

func (u *UserSlice) PluckCreatedAt() []time.Time {
	res := make([]time.Time, 0, len(*u))
	for _, item := range *u {
		res = append(res, item.CreatedAt)
	}
	return res
}

func (u *UserSlice) PluckUpdatedAt() []time.Time {
	res := make([]time.Time, 0, len(*u))
	for _, item := range *u {
		res = append(res, item.UpdatedAt)
	}
	return res
}

func (u *UserSlice) UniqueId() []int {
	uni := make(map[int]struct{})
	res := make([]int, 0)
	for _, item := range *u {
		uni[item.Id] = struct{}{}
	}
	for key := range uni {
		res = append(res, key)
	}
	return res
}

func (u *UserSlice) UniqueUsername() []string {
	uni := make(map[string]struct{})
	res := make([]string, 0)
	for _, item := range *u {
		uni[item.Username] = struct{}{}
	}
	for key := range uni {
		res = append(res, key)
	}
	return res
}

func (u *UserSlice) UniquePassword() []string {
	uni := make(map[string]struct{})
	res := make([]string, 0)
	for _, item := range *u {
		uni[item.Password] = struct{}{}
	}
	for key := range uni {
		res = append(res, key)
	}
	return res
}

func (u *UserSlice) UniqueTelephone() []string {
	uni := make(map[string]struct{})
	res := make([]string, 0)
	for _, item := range *u {
		uni[item.Telephone] = struct{}{}
	}
	for key := range uni {
		res = append(res, key)
	}
	return res
}

func (u *UserSlice) UniqueBirthday() []string {
	uni := make(map[string]struct{})
	res := make([]string, 0)
	for _, item := range *u {
		uni[item.Birthday] = struct{}{}
	}
	for key := range uni {
		res = append(res, key)
	}
	return res
}

func (u *UserSlice) UniqueGender() []int {
	uni := make(map[int]struct{})
	res := make([]int, 0)
	for _, item := range *u {
		uni[item.Gender] = struct{}{}
	}
	for key := range uni {
		res = append(res, key)
	}
	return res
}

func (u *UserSlice) UniqueCreatedAt() []time.Time {
	uni := make(map[time.Time]struct{})
	res := make([]time.Time, 0)
	for _, item := range *u {
		uni[item.CreatedAt] = struct{}{}
	}
	for key := range uni {
		res = append(res, key)
	}
	return res
}

func (u *UserSlice) UniqueUpdatedAt() []time.Time {
	uni := make(map[time.Time]struct{})
	res := make([]time.Time, 0)
	for _, item := range *u {
		uni[item.UpdatedAt] = struct{}{}
	}
	for key := range uni {
		res = append(res, key)
	}
	return res
}
