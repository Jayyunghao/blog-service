package service

import (
	"Practice/go-programming-tour-book/blog-service/global"
	"Practice/go-programming-tour-book/blog-service/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
	//dao *dao.DaoEtcd
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	//svc.dao = dao.NewDaoEtcd("etcd")
	return svc
}
