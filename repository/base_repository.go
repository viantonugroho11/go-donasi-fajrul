package repository

import (
	"context"
	// "database/sql"
	"donasi/config"

	"donasi/model"

	"github.com/midtrans/midtrans-go/snap"
)

type databaseRepository struct {
	confDb config.DatabaseConfig
	confMidtrans config.MidtransConfig
}

func NewDatabaseRepository(confDb config.DatabaseConfig, confMd config.MidtransConfig) DatabaseRepository {
	return &databaseRepository{confDb, confMd}
}

type DatabaseRepository interface {
	//donasi
	GetAllDonasi(ctx context.Context) (string, error)
	// PublishNotification(ctx context.Context, user *model.PayloadNotificationRequest) (string, error)
	// ConsumeNotificationFirebase() (result model.PayloadNotificationRequest, err error)

	//transaksi
	PostTranscationDonasi(ctx context.Context, payload *model.PayloadTransaksiRequest) (result *snap.Response, err error)
	StoreTranscationDonasi(ctx context.Context, payload *model.PayloadTransaksiRequest, donasiId string) (string, error)
}