package server

import (
	"github.com/JResearchLabs/Flutechain/consensus"
	consensusDev "github.com/JResearchLabs/Flutechain/consensus/dev"
	consensusDummy "github.com/JResearchLabs/Flutechain/consensus/dummy"
	consensusIBFT "github.com/JResearchLabs/Flutechain/consensus/ibft"
	"github.com/JResearchLabs/Flutechain/secrets"
	"github.com/JResearchLabs/Flutechain/secrets/awsssm"
	"github.com/JResearchLabs/Flutechain/secrets/gcpssm"
	"github.com/JResearchLabs/Flutechain/secrets/hashicorpvault"
	"github.com/JResearchLabs/Flutechain/secrets/local"
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
