package bridgeservice

import (
	"context"
	"time"

	fleetpb "github.com/etaques/orb/fleet/pb"
	policiespb "github.com/etaques/orb/policies/pb"
	"github.com/etaques/orb/sinker/config"
	"go.uber.org/zap"
)

type BridgeService interface {
	ExtractAgent(ctx context.Context, channelID string) (*fleetpb.AgentInfoRes, error)
	GetDataSetsFromAgentGroups(ctx context.Context, mfOwnerId string, agentGroupIds []string) (map[string]string, error)
	NotifyActiveSink(ctx context.Context, mfOwnerId, sinkId, state, message string) error
	GetSinkIdsFromPolicyID(ctx context.Context, mfOwnerId string, policyID string) (map[string]string, error)
}

func NewBridgeService(logger *zap.Logger,
	sinkerCache config.ConfigRepo,
	policiesClient policiespb.PolicyServiceClient,
	fleetClient fleetpb.FleetServiceClient) SinkerOtelBridgeService {
	return SinkerOtelBridgeService{
		logger:         logger,
		sinkerCache:    sinkerCache,
		policiesClient: policiesClient,
		fleetClient:    fleetClient,
	}
}

type SinkerOtelBridgeService struct {
	logger         *zap.Logger
	sinkerCache    config.ConfigRepo
	policiesClient policiespb.PolicyServiceClient
	fleetClient    fleetpb.FleetServiceClient
}

func (bs *SinkerOtelBridgeService) NotifyActiveSink(_ context.Context, mfOwnerId, sinkId, newState, message string) error {
	cfgRepo, err := bs.sinkerCache.Get(mfOwnerId, sinkId)
	if err != nil {
		bs.logger.Error("unable to retrieve the sink config", zap.Error(err))
		return err
	}
	err = cfgRepo.State.SetFromString(newState)
	if err != nil {
		bs.logger.Error("unable to set state", zap.String("new_state", newState), zap.Error(err))
		return err
	}
	if cfgRepo.State == config.Error {
		cfgRepo.Msg = message
	} else if cfgRepo.State == config.Active {
		cfgRepo.LastRemoteWrite = time.Now()
	}
	err = bs.sinkerCache.Edit(cfgRepo)
	if err != nil {
		bs.logger.Error("error during update sink cache", zap.String("sinkId", sinkId), zap.Error(err))
		return err
	}

	return nil
}

func (bs *SinkerOtelBridgeService) ExtractAgent(ctx context.Context, channelID string) (*fleetpb.AgentInfoRes, error) {
	agentPb, err := bs.fleetClient.RetrieveAgentInfoByChannelID(ctx, &fleetpb.AgentInfoByChannelIDReq{Channel: channelID})
	if err != nil {
		return nil, err
	}
	return agentPb, nil
}

func (bs *SinkerOtelBridgeService) GetSinkIdsFromPolicyID(ctx context.Context, mfOwnerId string, policyID string) (map[string]string, error) {
	// Here needs to retrieve datasets by policyID
	datasetRes, err := bs.policiesClient.RetrieveDatasetsByPolicy(ctx, &policiespb.DatasetsByPolicyReq{
		PolicyID: policyID,
		OwnerID:  mfOwnerId,
	})
	if err != nil {
		bs.logger.Info("unable to retrieve datasets from policy")
		return nil, err
	}
	mapSinkIdPolicy := make(map[string]string)
	for _, datasetList := range datasetRes.DatasetList {
		for _, sinkId := range datasetList.SinkIds {
			mapSinkIdPolicy[sinkId] = "active"
		}
	}
	return mapSinkIdPolicy, nil
}
