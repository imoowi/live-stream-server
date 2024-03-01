/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package repos

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/imoowi/live-stream-server/apps/user/models"
	"github.com/imoowi/comer/components"
	"github.com/imoowi/live-stream-server/global" 
	"github.com/imoowi/comer/utils/maker" 
	"github.com/imoowi/comer/utils/password" 
	"github.com/imoowi/comer/utils/request" 
	"github.com/imoowi/comer/utils/response"
)

type UserRepo struct {
	Db *components.MysqlODM
}

func newUserRepo() *UserRepo {
	return &UserRepo{
		Db: global.MysqlDb,
	}
}
func init() {
	global.DigContainer.Provide(newUserRepo)
}
func (r *UserRepo) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	db := r.Db.Client
	var users []*models.User

	if req.SearchKey != `` {
		db = db.Where(`name LIKE ?`, `%`+req.SearchKey+`%`)
	}
	offset := (req.Page - 1) * req.PageSize
	db = db.Offset(int(offset)).Limit(int(req.PageSize))
	// db=db.Order(`name desc`)
	err = db.Find(&users).Error

	var count int64
	db.Offset(-1).Limit(-1).Count(&count)

	res = &response.PageList{
		List:  users,
		Pages: response.MakePages(count, req.Page, req.PageSize),
	}
	return
}

func (r *UserRepo) One(c *gin.Context, id uint) (user *models.User, err error) {
	db := r.Db.Client
	err = db.Where(`id=?`, id).First(&user).Error
	return
}

func (r *UserRepo) OneByUsername(c *gin.Context, username string) (user *models.User, err error) {
	db := r.Db.Client
	err = db.Where(`username=?`, username).First(&user).Error
	return
}
func (r *UserRepo) Add(c *gin.Context, model *models.User) (newId uint, err error) {
	db := r.Db.Client
	model.Salt = maker.MakeRandStr(6)
	model.Passwd = password.GeneratePassword(model.Passwd + model.Salt)
	db = db.Create(&model)
	err = db.Error

	newId = model.ID
	return
}

func (r *UserRepo) Update(c *gin.Context, model *models.User, id uint) (updated bool, err error) {
	if id == 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	model.ID = uint(id)
	db := r.Db.Client
	err = db.Omit(`created_at`).Save(&model).Error
	if err == nil {
		updated = true
	}
	return
}

func (r *UserRepo) PatchUpdate(c *gin.Context, patchData map[string]any, id uint) (updated bool, err error) {
	if id == 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	model, err := r.One(c, id)
	if err != nil {
		return
	}
	if model == nil {
		err = errors.New(`no data existed`)
		return
	}

	patchDataBytes, err := json.Marshal(patchData)
	if err != nil {
		return
	}
	err = json.Unmarshal(patchDataBytes, &model)
	if err != nil {
		return
	}

	db := r.Db.Client
	err = db.Omit(`created_at`).Save(&model).Error
	if err == nil {
		updated = true
	}
	return
}

func (r *UserRepo) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	if id == 0 {
		deleted = false
		err = errors.New(`pls input id`)
		return
	}
	db := r.Db.Client
	model, err := r.One(c, id)
	if err != nil {
		return
	}
	if model.ID == 0 {
		err = errors.New(`obj not existe`)
		return
	}
	err = db.Delete(&model).Error
	if err == nil {
		deleted = true
	}
	return
}

func (r *UserRepo) Login(c *gin.Context, login *models.UserLogin) (*models.User, error) {
	user, err := r.OneByUsername(c, login.Username)
	if err != nil {
		return nil, err
	}
	if user.ID > 0 {
		if user.Passwd == password.GeneratePassword(login.Passwd+user.Salt) {
			return user, nil
		} else {
			return nil, errors.New(`密码错误`)
		}
	}
	return nil, err
}

func (r *UserRepo) ChgPwd(c *gin.Context, userChgPwd *models.UserChgPwd) (ok bool, err error) {
	user, err := r.One(c, userChgPwd.UserId)
	if err != nil {
		return
	}
	if user.ID <= 0 {
		err = errors.New(`用户不存在`)
		return
	}
	if userChgPwd.NewPwd != userChgPwd.ConfirmPwd {
		err = errors.New(`两次输入的新密码不一致`)
		return
	}
	if user.Passwd != password.GeneratePassword(userChgPwd.OriginPwd+user.Salt) {
		err = errors.New(`原始密码错误`)
		return
	}
	user.Salt = maker.MakeRandStr(6)
	user.Passwd = password.GeneratePassword(userChgPwd.NewPwd + user.Salt)
	db := r.Db.Client
	err = db.Omit(`created_at`).Save(&user).Error
	if err != nil {
		ok = false
	}
	ok = true
	return
}
