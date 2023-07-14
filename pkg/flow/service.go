package flow

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/auth4flow/auth4flow-core/pkg/config"
	"github.com/auth4flow/auth4flow-core/pkg/service"

	"github.com/onflow/cadence"
	flowSDK "github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access/http"
	"github.com/rs/zerolog/log"
)

type FlowService struct {
	service.BaseService
	Config       config.Auth4FlowConfig
	FlowClient   *http.Client
	queue        *Queue
	eventMonitor *EventMonitorService
}

func NewService(env service.Env, cfg config.Auth4FlowConfig) *FlowService {
	flowClient, err := http.NewClient(cfg.GetFlowNetwork())
	if err != nil {
		log.Fatal().Err(err).Msg("Could not initialize and connect to the configured Flow Blockchain. Shutting down.")
	}

	svc := &FlowService{
		BaseService: service.NewBaseService(env),
		Config:      cfg,
		FlowClient:  flowClient,
	}

	svc.queue = newQueue(svc, 25)
	go svc.queue.Start()

	svc.eventMonitor = newEventMonitorService(svc)

	return svc
}

func (svc *FlowService) ID() string {
	return service.FlowService
}

func (svc *FlowService) StartQueue() error {
	if !svc.queue.running {
		go svc.queue.Start()
		return nil
	}

	return errors.New("queue is already running")
}

func (svc *FlowService) StopQueue() error {
	if svc.queue.running {
		svc.queue.Stop()
		return nil
	}

	return errors.New("queue is not running")
}

func (svc *FlowService) Wait() {
	svc.queue.WaitGroup.Wait()
}

func (svc *FlowService) CreateQueueJob(job JobInterface) (string, error) {
	return svc.queue.CreateJob(job)
}

func (svc *FlowService) RemoveQueueJobByID(id string) error {
	return svc.queue.RemoveJobByID(id)
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

	script, err := getVerifyAccountProofScript(svc.Config.GetFlowNetwork())
	if err != nil {
		return false, err
	}

	scriptBytes := []byte(script)

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
