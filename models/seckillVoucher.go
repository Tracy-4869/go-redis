package models

import (
	"go-redis/utils"
	"time"
)

type SeckillVoucher struct {
	VoucherId  int              `json:"voucher_id"`
	Stock      int              `json:"stock"`
	CreateTime time.Time        `json:"create_time"`
	BeginTime  *utils.LocalTime `json:"begin_time"`
	EndTime    *utils.LocalTime `json:"end_time"`
	UpdateTime time.Time        `json:"update_time"`
}

func (SeckillVoucher) TableName() string {
	return "tb_seckill_voucher"
}

func AddSkillVoucher(skillVoucher SeckillVoucher) error {
	return db.Create(&skillVoucher).Error
}