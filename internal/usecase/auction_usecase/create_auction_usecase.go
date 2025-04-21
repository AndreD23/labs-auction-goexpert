package auction_usecase

import (
	"context"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/entity/bid_entity"
	"fullcycle-auction_go/internal/internal_error"
	"fullcycle-auction_go/internal/usecase/bid_usecase"
	"os"
	"strconv"
	"time"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=2"`
	Description string           `json:"description" binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition" binding:"oneof=0 1 2"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

type AuctionUseCaseInterface interface {
	CreateAuction(
		ctx context.Context,
		auctionInput AuctionInputDTO) *internal_error.InternalError

	FindAuctionById(
		ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)

	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)

	FindWinningBidByAuctionId(
		ctx context.Context,
		auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError)
}

func NewAuctionUseCase(
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface,
	bidRepositoryInterface bid_entity.BidEntityRepository) AuctionUseCaseInterface {

	interval := getBatchAuctionInterval()
	batchSize := getBatchSize()

	auctionusecase := &AuctionUseCase{
		auctionRepositoryInterface: auctionRepositoryInterface,
		bidRepositoryInterface:     bidRepositoryInterface,
		closeTimeInterval:          interval,
		timer:                      time.NewTimer(interval),
		maxAuctionBatchSize:        batchSize,
		auctionChannel:             make(chan auction_entity.Auction),
	}

	go auctionusecase.handleCloseAuctionsRoutine(context.Background())
	return auctionusecase
}

var auctionBatch []auction_entity.Auction

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface     bid_entity.BidEntityRepository
	timer                      *time.Timer
	maxAuctionBatchSize        int
	closeTimeInterval          time.Duration
	auctionChannel             chan auction_entity.Auction
}

func (au *AuctionUseCase) CreateAuction(
	ctx context.Context,
	auctionInput AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionInput.ProductName,
		auctionInput.Category,
		auctionInput.Description,
		auction_entity.ProductCondition(auctionInput.Condition))
	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(
		ctx, auction); err != nil {
		return err
	}

	return nil
}

func (au *AuctionUseCase) handleCloseAuctionsRoutine(ctx context.Context) {
	defer close(au.auctionChannel)

	for {
		select {
		case auctionEntity, ok := <-au.auctionChannel:
			if !ok {
				// Se o tamanho do batch for maior que 0, fecha os audits
				if len(auctionBatch) > 0 {
					if err := au.auctionRepositoryInterface.CloseAuctions(ctx, auctionBatch); err != nil {
						logger.Error("error on close batch", err)
					}
				}
				return
			}

			auctionBatch = append(auctionBatch, auctionEntity)

			if len(auctionBatch) >= au.maxAuctionBatchSize {
				if err := au.auctionRepositoryInterface.CloseAuctions(ctx, auctionBatch); err != nil {
					logger.Error("error on close batch", err)
				}

				// Zera a lista
				auctionBatch = nil

				// Reinicia o timer
				au.timer.Reset(au.closeTimeInterval)
			}
		case <-au.timer.C:
			if err := au.auctionRepositoryInterface.CloseAuctions(ctx, auctionBatch); err != nil {
				logger.Error("error on close batch", err)
			}

			// Zera a lista
			auctionBatch = nil

			// Reinicia o timer
			au.timer.Reset(au.closeTimeInterval)
		}

	}
}

func getBatchAuctionInterval() time.Duration {
	interval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return 3 * time.Minute
	}

	return duration
}

func getBatchSize() int {
	batchSize, err := strconv.Atoi(os.Getenv("AUCTION_BATCH_SIZE"))
	if err != nil {
		return 10
	}

	return batchSize
}
