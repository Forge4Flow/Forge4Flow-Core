package flow

import (
	"context"
	"encoding/hex"

	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/service"

	"github.com/onflow/cadence"
	flowSDK "github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access/http"
	"github.com/rs/zerolog/log"
)

type FlowService struct {
	service.BaseService
	Config     config.Forge4FlowConfig
	FlowClient *http.Client
}

func NewService(env service.Env, cfg config.Forge4FlowConfig) *FlowService {
	flowClient, err := http.NewClient(cfg.GetFlowNetwork())
	if err != nil {
		log.Fatal().Err(err).Msg("Could not initialize and connect to the configured Flow Blockchain. Shutting down.")
	}

	svc := &FlowService{
		BaseService: service.NewBaseService(env),
		Config:      cfg,
		FlowClient:  flowClient,
	}

	return svc
}

func (svc FlowService) ID() string {
	return service.FlowService
}

func (svc *FlowService) VerifyAccountProof(ctx context.Context, accountProof AccountProofSpec) (bool, error) {
	address := flowSDK.HexToAddress(accountProof.Address)
	message, err := flowSDK.EncodeAccountProofMessage(address, svc.Config.GetAppIdentifier(), accountProof.Nonce)
	if err != nil {
		return false, err
	}

	var signaturesArr []string
	var keyIndices []int

	for _, el := range accountProof.Signatures {
		signaturesArr = append(signaturesArr, el.Signature)
		keyIndices = append(keyIndices, el.KeyId)
	}

	scriptBytes := []byte(getVerifyAccountProofScript(svc.Config.GetFlowNetwork()))

	cadenceSignaturesArr, err := InputToCadence(signaturesArr, resolver)
	if err != nil {
		return false, err
	}

	cadenceKeyIndices, err := InputToCadence(keyIndices, resolver)
	if err != nil {
		return false, err
	}

	args := []cadence.Value{
		cadence.BytesToAddress(address.Bytes()),
		cadence.String(hex.EncodeToString(message)),
		cadenceKeyIndices,
		cadenceSignaturesArr,
	}

	value, err := svc.FlowClient.ExecuteScriptAtLatestBlock(ctx, scriptBytes, args)
	if err != nil {
		return false, err
	}

	return value.ToGoValue().(bool), nil
}

func resolver(str string) (string, error) {
	return str, nil
}
