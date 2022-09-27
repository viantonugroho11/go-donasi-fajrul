package service

import (
	"context"
	"donasi/model"
	"donasi/repository"
)

type TransaksiService interface {
	PostTranscationDonasi(ctx context.Context, payload *model.PayloadTransaksiRequest)(string, error)
}

type transaksiService struct {
	dbRepo repository.DatabaseRepository
}

func NewTransaksiService(dbRepo repository.DatabaseRepository) TransaksiService {
	return &transaksiService{dbRepo}
}

func(transSrv *transaksiService) PostTranscationDonasi(ctx context.Context, payload *model.PayloadTransaksiRequest)(string, error){

	resultMd , err := transSrv.dbRepo.PostTranscationDonasi(ctx, payload)
	if err != nil {
		return "", err
	}		
	resultDb, err := transSrv.dbRepo.StoreTranscationDonasi(ctx, payload, resultMd.Token)
	if err != nil {
		return "", err
	}
	return resultDb, nil
}