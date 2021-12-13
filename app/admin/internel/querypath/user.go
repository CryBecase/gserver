package querypath

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"gserver/pkg/cache/redis"
)

type User struct {
	*querypath
}

func NewUser(db *gorm.DB) *User {
	return &User{
		querypath: &querypath{
			db: db,
			p:  nil,
		},
	}
}

// ---------------
// -----自定义-----
// ---------------

// ---------------
// -----WHERE-----
// ---------------

func (u *User) WhIdEq(v int) *User {
	u.db.Where("id = ?", v)
	return u
}

func (u *User) ZWhIdEq(v int) *User {
	if v == 0 {
		return u
	}

	u.db.Where("id = ?", v)
	return u
}

func (u *User) WhIdNotEq(v int) *User {
	u.db.Where("id != ?", v)
	return u
}

func (u *User) ZWhIdNotEq(v int) *User {
	if v == 0 {
		return u
	}

	u.db.Where("id != ?", v)
	return u
}

func (u *User) WhIdIn(v []int) *User {
	u.db.Where("id IN ?", v)
	return u
}

func (u *User) ZWhIdIn(v []int) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("id IN ?", v)
	return u
}

func (u *User) WhUsernameEq(v string) *User {
	u.db.Where("username = ?", v)
	return u
}

func (u *User) ZWhUsernameEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("username = ?", v)
	return u
}

func (u *User) WhUsernameNotEq(v string) *User {
	u.db.Where("username != ?", v)
	return u
}

func (u *User) ZWhUsernameNotEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("username != ?", v)
	return u
}

func (u *User) WhUsernameIn(v []string) *User {
	u.db.Where("username IN ?", v)
	return u
}

func (u *User) ZWhUsernameIn(v []string) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("username IN ?", v)
	return u
}

func (u *User) WhUsernameLike(v string) *User {
	u.db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) ZWhUsernameLike(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) WhPasswordEq(v string) *User {
	u.db.Where("password = ?", v)
	return u
}

func (u *User) ZWhPasswordEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("password = ?", v)
	return u
}

func (u *User) WhPasswordNotEq(v string) *User {
	u.db.Where("password != ?", v)
	return u
}

func (u *User) ZWhPasswordNotEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("password != ?", v)
	return u
}

func (u *User) WhPasswordIn(v []string) *User {
	u.db.Where("password IN ?", v)
	return u
}

func (u *User) ZWhPasswordIn(v []string) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("password IN ?", v)
	return u
}

func (u *User) WhPasswordLike(v string) *User {
	u.db.Where("password LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) ZWhPasswordLike(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("password LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) WhTelephoneEq(v string) *User {
	u.db.Where("telephone = ?", v)
	return u
}

func (u *User) ZWhTelephoneEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("telephone = ?", v)
	return u
}

func (u *User) WhTelephoneNotEq(v string) *User {
	u.db.Where("telephone != ?", v)
	return u
}

func (u *User) ZWhTelephoneNotEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("telephone != ?", v)
	return u
}

func (u *User) WhTelephoneIn(v []string) *User {
	u.db.Where("telephone IN ?", v)
	return u
}

func (u *User) ZWhTelephoneIn(v []string) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("telephone IN ?", v)
	return u
}

func (u *User) WhTelephoneLike(v string) *User {
	u.db.Where("telephone LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) ZWhTelephoneLike(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("telephone LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) WhBirthdayEq(v string) *User {
	u.db.Where("birthday = ?", v)
	return u
}

func (u *User) ZWhBirthdayEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("birthday = ?", v)
	return u
}

func (u *User) WhBirthdayNotEq(v string) *User {
	u.db.Where("birthday != ?", v)
	return u
}

func (u *User) ZWhBirthdayNotEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("birthday != ?", v)
	return u
}

func (u *User) WhBirthdayIn(v []string) *User {
	u.db.Where("birthday IN ?", v)
	return u
}

func (u *User) ZWhBirthdayIn(v []string) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("birthday IN ?", v)
	return u
}

func (u *User) WhBirthdayLike(v string) *User {
	u.db.Where("birthday LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) ZWhBirthdayLike(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("birthday LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) WhGenderEq(v int) *User {
	u.db.Where("gender = ?", v)
	return u
}

func (u *User) ZWhGenderEq(v int) *User {
	if v == 0 {
		return u
	}

	u.db.Where("gender = ?", v)
	return u
}

func (u *User) WhGenderNotEq(v int) *User {
	u.db.Where("gender != ?", v)
	return u
}

func (u *User) ZWhGenderNotEq(v int) *User {
	if v == 0 {
		return u
	}

	u.db.Where("gender != ?", v)
	return u
}

func (u *User) WhGenderIn(v []int) *User {
	u.db.Where("gender IN ?", v)
	return u
}

func (u *User) ZWhGenderIn(v []int) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("gender IN ?", v)
	return u
}

func (u *User) WhCreatedAtEq(v time.Time) *User {
	u.db.Where("created_at = ?", v)
	return u
}

func (u *User) ZWhCreatedAtEq(v time.Time) *User {
	if v.IsZero() {
		return u
	}

	u.db.Where("created_at = ?", v)
	return u
}

func (u *User) WhCreatedAtNotEq(v time.Time) *User {
	u.db.Where("created_at != ?", v)
	return u
}

func (u *User) ZWhCreatedAtNotEq(v time.Time) *User {
	if v.IsZero() {
		return u
	}

	u.db.Where("created_at != ?", v)
	return u
}

func (u *User) WhCreatedAtIn(v []time.Time) *User {
	u.db.Where("created_at IN ?", v)
	return u
}

func (u *User) ZWhCreatedAtIn(v []time.Time) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("created_at IN ?", v)
	return u
}

func (u *User) WhUpdatedAtEq(v time.Time) *User {
	u.db.Where("updated_at = ?", v)
	return u
}

func (u *User) ZWhUpdatedAtEq(v time.Time) *User {
	if v.IsZero() {
		return u
	}

	u.db.Where("updated_at = ?", v)
	return u
}

func (u *User) WhUpdatedAtNotEq(v time.Time) *User {
	u.db.Where("updated_at != ?", v)
	return u
}

func (u *User) ZWhUpdatedAtNotEq(v time.Time) *User {
	if v.IsZero() {
		return u
	}

	u.db.Where("updated_at != ?", v)
	return u
}

func (u *User) WhUpdatedAtIn(v []time.Time) *User {
	u.db.Where("updated_at IN ?", v)
	return u
}

func (u *User) ZWhUpdatedAtIn(v []time.Time) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("updated_at IN ?", v)
	return u
}

func (u *User) Where(query interface{}, args ...interface{}) *User {
	u.db.Where(query, args...)
	return u
}

// ------------
// -----OR-----
// ------------

func (u *User) OrIdEq(v int) *User {
	u.db.Where("id = ?", v)
	return u
}

func (u *User) ZOrIdEq(v int) *User {
	if v == 0 {
		return u
	}

	u.db.Where("id = ?", v)
	return u
}

func (u *User) OrIdNotEq(v int) *User {
	u.db.Where("id != ?", v)
	return u
}

func (u *User) ZOrIdNotEq(v int) *User {
	if v == 0 {
		return u
	}

	u.db.Where("id != ?", v)
	return u
}

func (u *User) OrIdIn(v []int) *User {
	u.db.Where("id IN ?", v)
	return u
}

func (u *User) ZOrIdIn(v []int) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("id IN ?", v)
	return u
}

func (u *User) OrUsernameEq(v string) *User {
	u.db.Where("username = ?", v)
	return u
}

func (u *User) ZOrUsernameEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("username = ?", v)
	return u
}

func (u *User) OrUsernameNotEq(v string) *User {
	u.db.Where("username != ?", v)
	return u
}

func (u *User) ZOrUsernameNotEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("username != ?", v)
	return u
}

func (u *User) OrUsernameIn(v []string) *User {
	u.db.Where("username IN ?", v)
	return u
}

func (u *User) ZOrUsernameIn(v []string) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("username IN ?", v)
	return u
}

func (u *User) OrUsernameLike(v string) *User {
	u.db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) ZOrUsernameLike(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) OrPasswordEq(v string) *User {
	u.db.Where("password = ?", v)
	return u
}

func (u *User) ZOrPasswordEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("password = ?", v)
	return u
}

func (u *User) OrPasswordNotEq(v string) *User {
	u.db.Where("password != ?", v)
	return u
}

func (u *User) ZOrPasswordNotEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("password != ?", v)
	return u
}

func (u *User) OrPasswordIn(v []string) *User {
	u.db.Where("password IN ?", v)
	return u
}

func (u *User) ZOrPasswordIn(v []string) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("password IN ?", v)
	return u
}

func (u *User) OrPasswordLike(v string) *User {
	u.db.Where("password LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) ZOrPasswordLike(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("password LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) OrTelephoneEq(v string) *User {
	u.db.Where("telephone = ?", v)
	return u
}

func (u *User) ZOrTelephoneEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("telephone = ?", v)
	return u
}

func (u *User) OrTelephoneNotEq(v string) *User {
	u.db.Where("telephone != ?", v)
	return u
}

func (u *User) ZOrTelephoneNotEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("telephone != ?", v)
	return u
}

func (u *User) OrTelephoneIn(v []string) *User {
	u.db.Where("telephone IN ?", v)
	return u
}

func (u *User) ZOrTelephoneIn(v []string) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("telephone IN ?", v)
	return u
}

func (u *User) OrTelephoneLike(v string) *User {
	u.db.Where("telephone LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) ZOrTelephoneLike(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("telephone LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) OrBirthdayEq(v string) *User {
	u.db.Where("birthday = ?", v)
	return u
}

func (u *User) ZOrBirthdayEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("birthday = ?", v)
	return u
}

func (u *User) OrBirthdayNotEq(v string) *User {
	u.db.Where("birthday != ?", v)
	return u
}

func (u *User) ZOrBirthdayNotEq(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("birthday != ?", v)
	return u
}

func (u *User) OrBirthdayIn(v []string) *User {
	u.db.Where("birthday IN ?", v)
	return u
}

func (u *User) ZOrBirthdayIn(v []string) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("birthday IN ?", v)
	return u
}

func (u *User) OrBirthdayLike(v string) *User {
	u.db.Where("birthday LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) ZOrBirthdayLike(v string) *User {
	if v == "" {
		return u
	}

	u.db.Where("birthday LIKE ?", fmt.Sprintf("%%%s%%", v))
	return u
}

func (u *User) OrGenderEq(v int) *User {
	u.db.Where("gender = ?", v)
	return u
}

func (u *User) ZOrGenderEq(v int) *User {
	if v == 0 {
		return u
	}

	u.db.Where("gender = ?", v)
	return u
}

func (u *User) OrGenderNotEq(v int) *User {
	u.db.Where("gender != ?", v)
	return u
}

func (u *User) ZOrGenderNotEq(v int) *User {
	if v == 0 {
		return u
	}

	u.db.Where("gender != ?", v)
	return u
}

func (u *User) OrGenderIn(v []int) *User {
	u.db.Where("gender IN ?", v)
	return u
}

func (u *User) ZOrGenderIn(v []int) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("gender IN ?", v)
	return u
}

func (u *User) OrCreatedAtEq(v time.Time) *User {
	u.db.Where("created_at = ?", v)
	return u
}

func (u *User) ZOrCreatedAtEq(v time.Time) *User {
	if v.IsZero() {
		return u
	}

	u.db.Where("created_at = ?", v)
	return u
}

func (u *User) OrCreatedAtNotEq(v time.Time) *User {
	u.db.Where("created_at != ?", v)
	return u
}

func (u *User) ZOrCreatedAtNotEq(v time.Time) *User {
	if v.IsZero() {
		return u
	}

	u.db.Where("created_at != ?", v)
	return u
}

func (u *User) OrCreatedAtIn(v []time.Time) *User {
	u.db.Where("created_at IN ?", v)
	return u
}

func (u *User) ZOrCreatedAtIn(v []time.Time) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("created_at IN ?", v)
	return u
}

func (u *User) OrUpdatedAtEq(v time.Time) *User {
	u.db.Where("updated_at = ?", v)
	return u
}

func (u *User) ZOrUpdatedAtEq(v time.Time) *User {
	if v.IsZero() {
		return u
	}

	u.db.Where("updated_at = ?", v)
	return u
}

func (u *User) OrUpdatedAtNotEq(v time.Time) *User {
	u.db.Where("updated_at != ?", v)
	return u
}

func (u *User) ZOrUpdatedAtNotEq(v time.Time) *User {
	if v.IsZero() {
		return u
	}

	u.db.Where("updated_at != ?", v)
	return u
}

func (u *User) OrUpdatedAtIn(v []time.Time) *User {
	u.db.Where("updated_at IN ?", v)
	return u
}

func (u *User) ZOrUpdatedAtIn(v []time.Time) *User {
	if len(v) == 0 {
		return u
	}

	u.db.Where("updated_at IN ?", v)
	return u
}

// ------------------
// -----ORDER BY-----
// ------------------

func (u *User) OrderByIdDesc() *User {
	u.db.Order("id DESC")
	return u
}

func (u *User) OrderByIdAsc() *User {
	u.db.Order("id ASC")
	return u
}

func (u *User) OrderByUsernameDesc() *User {
	u.db.Order("username DESC")
	return u
}

func (u *User) OrderByUsernameAsc() *User {
	u.db.Order("username ASC")
	return u
}

func (u *User) OrderByPasswordDesc() *User {
	u.db.Order("password DESC")
	return u
}

func (u *User) OrderByPasswordAsc() *User {
	u.db.Order("password ASC")
	return u
}

func (u *User) OrderByTelephoneDesc() *User {
	u.db.Order("telephone DESC")
	return u
}

func (u *User) OrderByTelephoneAsc() *User {
	u.db.Order("telephone ASC")
	return u
}

func (u *User) OrderByBirthdayDesc() *User {
	u.db.Order("birthday DESC")
	return u
}

func (u *User) OrderByBirthdayAsc() *User {
	u.db.Order("birthday ASC")
	return u
}

func (u *User) OrderByGenderDesc() *User {
	u.db.Order("gender DESC")
	return u
}

func (u *User) OrderByGenderAsc() *User {
	u.db.Order("gender ASC")
	return u
}

func (u *User) OrderByCreatedAtDesc() *User {
	u.db.Order("created_at DESC")
	return u
}

func (u *User) OrderByCreatedAtAsc() *User {
	u.db.Order("created_at ASC")
	return u
}

func (u *User) OrderByUpdatedAtDesc() *User {
	u.db.Order("updated_at DESC")
	return u
}

func (u *User) OrderByUpdatedAtAsc() *User {
	u.db.Order("updated_at ASC")
	return u
}

func (u *User) Raw(sql string, values ...interface{}) *User {
	u.db.Raw(sql, values...)
	return u
}

func (u *User) SetPaginate(p *Paginate) *User {
	u.p = p
	return u
}

func (u *User) WithCache(redis *redis.Redis, key string, expire time.Duration) *User {
	u.redis = &redisWrapper{redis, key, expire}
	return u
}
