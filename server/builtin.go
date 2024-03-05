package server

import (
	"github.com/JResearchLabs/aigisos/consensus"
	consensusDev "github.com/JResearchLabs/aigisos/consensus/dev"
	consensusDummy "github.com/JResearchLabs/aigisos/consensus/dummy"
	consensusIBFT "github.com/JResearchLabs/aigisos/consensus/ibft"
	"github.com/JResearchLabs/aigisos/secrets"
	"github.com/JResearchLabs/aigisos/secrets/awsssm"
	"github.com/JResearchLabs/aigisos/secrets/gcpssm"
	"github.com/JResearchLabs/aigisos/secrets/hashicorpvault"
	"github.com/JResearchLabs/aigisos/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
