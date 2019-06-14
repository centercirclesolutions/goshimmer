package request

import (
	"github.com/iotaledger/goshimmer/packages/identity"
	"github.com/iotaledger/goshimmer/plugins/autopeering/types/peer"
)

const (
	MARSHALLED_PACKET_HEADER = 0xBE

	PACKET_HEADER_START        = 0
	MARSHALLED_ISSUER_START    = PACKET_HEADER_END
	MARSHALLED_SIGNATURE_START = MARSHALLED_ISSUER_END

	PACKET_HEADER_END        = PACKET_HEADER_START + PACKET_HEADER_SIZE
	MARSHALLED_ISSUER_END    = MARSHALLED_ISSUER_START + MARSHALLED_ISSUER_SIZE
	MARSHALLED_SIGNATURE_END = MARSHALLED_SIGNATURE_START + MARSHALLED_SIGNATURE_SIZE

	PACKET_HEADER_SIZE        = 1
	MARSHALLED_ISSUER_SIZE    = peer.MARSHALLED_TOTAL_SIZE
	MARSHALLED_SIGNATURE_SIZE = identity.SIGNATURE_BYTE_LENGTH

	MARSHALLED_TOTAL_SIZE = MARSHALLED_SIGNATURE_END
)
