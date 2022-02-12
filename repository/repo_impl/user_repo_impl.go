package repo_impl

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"lienquanMess/banana"
	"lienquanMess/db"
	"lienquanMess/log"
	"lienquanMess/model"
	"lienquanMess/model/req"
	"lienquanMess/repository"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (n *UserRepoImpl) SaveUser(context context.Context, user model.User, userInfo model.AllUser, userProfile model.UserProfile) (model.User, error) {
	statement := `
		INSERT INTO users(id, full_name, phone,password, created_at, updated_at)
		VALUES(:id, :full_name, :phone, :password, :created_at, :updated_at)
	`
	statementUserInfo := `
		INSERT INTO userinfo(id, full_name)
		VALUES(:id, :full_name)
	`
	statementUserProfile := `
		INSERT INTO userprofile(id, full_name, created_at, updated_at)
		VALUES(:id, :full_name, :created_at, :updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	userProfile.CreatedAt = time.Now()
	userProfile.UpdatedAt = time.Now()

	_, err := n.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		print("a")
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFail
	}
	n.sql.Db.NamedExecContext(context, statementUserInfo, userInfo)
	n.sql.Db.NamedExecContext(context, statementUserProfile, userProfile)
	return user, nil
}

func (n *UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqSignIp) (model.User, error) {
	var user = model.User{}
	err := n.sql.Db.GetContext(context, &user, "SELECT *FROM users WHERE phone=$1", loginReq.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}

func (n *UserRepoImpl) SelectUserProfileById(context context.Context, userId string) (model.UserProfile, error) {
	var user model.UserProfile
	err := n.sql.Db.GetContext(context, &user,
		"SELECT * FROM  userprofile WHERE id = $1", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (n *UserRepoImpl) SelectAllUserById(context context.Context, start string, userId string) ([]model.AllUser, error) {
	var user []model.AllUser

	err := n.sql.Db.SelectContext(context, &user,
		"SELECT * FROM  userinfo WHERE id != $2 offset $1 limit 100", start, userId)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}

func (n UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	sqlStatement := `
		UPDATE users
		SET 
			full_name  = (CASE WHEN LENGTH(:full_name) = 0 THEN full_name ELSE :full_name END),
			sex = (CASE WHEN LENGTH(:sex) = 0 THEN sex ELSE :sex END),
		    age = (CASE WHEN LENGTH(:age) = 0 THEN age ELSE :age END),
		    height = (CASE WHEN LENGTH(:height) = 0 THEN height ELSE :height END),
		    job = (CASE WHEN LENGTH(:job) = 0 THEN job ELSE :job END),
		    income = (CASE WHEN LENGTH(:income) = 0 THEN income ELSE :income END),
		    marriage = (CASE WHEN LENGTH(:marriage) = 0 THEN marriage ELSE :marriage END),
		    children = (CASE WHEN LENGTH(:children) = 0 THEN children ELSE :children END),
		    home = (CASE WHEN LENGTH(:home) = 0 THEN home ELSE :home END),
		    zodiac = (CASE WHEN LENGTH(:zodiac) = 0 THEN zodiac ELSE :zodiac END),
		    status = (CASE WHEN LENGTH(:status) = 0 THEN status ELSE :status END),
		    formality = (CASE WHEN LENGTH(:formality) = 0 THEN formality ELSE :formality END),
		    link_fb = (CASE WHEN LENGTH(:link_fb) = 0 THEN link_fb ELSE :link_fb END),
		    link_is = (CASE WHEN LENGTH(:link_is) = 0 THEN link_is ELSE :link_is END),
		    zalo = (CASE WHEN LENGTH(:zalo) = 0 THEN zalo ELSE :zalo END),
		    address = (CASE WHEN LENGTH(:address) = 0 THEN address ELSE :address END),
		    target = (CASE WHEN LENGTH(:target) = 0 THEN target ELSE :target END),
		    about = (CASE WHEN LENGTH(:about) = 0 THEN about ELSE :about END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE user_id    = :user_id
	`

	user.UpdatedAt = time.Now()

	result, err := n.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, banana.UserNotUpdated
	}
	if count == 0 {
		return user, banana.UserNotUpdated
	}

	return user, nil
}

func (n UserRepoImpl) UpdateLineUser(context context.Context, user model.UserProfile) (model.UserProfile, error) {
	sqlStatement := `
		UPDATE userprofile
		SET 
			avatar_url  = (CASE WHEN LENGTH(:avatar_url) = 0 THEN avatar_url ELSE :avatar_url END),
			full_name = (CASE WHEN LENGTH(:full_name) = 0 THEN full_name ELSE :full_name END),
		    age = (CASE WHEN LENGTH(:age) = 0 THEN age ELSE :age END),
		    sex = (CASE WHEN LENGTH(:sex) = 0 THEN sex ELSE :sex END),
		    height = (CASE WHEN :height = 0 THEN height ELSE :height END),
		    job = (CASE WHEN LENGTH(:job) = 0 THEN job ELSE :job END),
		    income = (CASE WHEN LENGTH(:income) = 0 THEN income ELSE :income END),
		    marriage = (CASE WHEN LENGTH(:marriage) = 0 THEN marriage ELSE :marriage END),
		    children = (CASE WHEN LENGTH(:children) = 0 THEN children ELSE :children END),
		    home = (CASE WHEN LENGTH(:home) = 0 THEN home ELSE :home END),
		    zodiac = (CASE WHEN LENGTH(:zodiac) = 0 THEN zodiac ELSE :zodiac END),
		    formality = (CASE WHEN LENGTH(:formality) = 0 THEN formality ELSE :formality END),
		    link_fb = (CASE WHEN LENGTH(:link_fb) = 0 THEN link_fb ELSE :link_fb END),
		    link_is = (CASE WHEN LENGTH(:link_is) = 0 THEN link_is ELSE :link_is END),
		    zl_phone = (CASE WHEN LENGTH(:zl_phone) = 0 THEN zl_phone ELSE :zl_phone END),
		    address = (CASE WHEN LENGTH(:address) = 0 THEN address ELSE :address END),
		    target = (CASE WHEN LENGTH(:target) = 0 THEN target ELSE :target END),
		    about = (CASE WHEN LENGTH(:about) = 0 THEN about ELSE :about END),
		    count_follower = (CASE WHEN LENGTH(:count_follower) = 0 THEN count_follower ELSE :count_follower END),
		    count_following = (CASE WHEN LENGTH(:count_following) = 0 THEN count_following ELSE :count_following END),
		    count_like = (CASE WHEN LENGTH(:count_like) = 0 THEN count_like ELSE :count_like END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE id    = :id
	`
	user.UpdatedAt = time.Now()

	result, err := n.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, banana.UserNotUpdated
	}
	if count == 0 {
		return user, banana.UserNotUpdated
	}

	return user, nil
}
