

package params



const (
	// BloomBitsBlocks is the number of blocks a single bloom bit section vector
	// contains on the server side.
	BloomBitsBlocks uint64 = 4096

	// BloomBitsBlocksClient is the number of blocks a single bloom bit section vector
	// contains on the light client side
	BloomBitsBlocksClient uint64 = 32768

	// BloomConfirms is the number of confirmation blocks before a bloom section is
	// considered probably final and its rotated bits are calculated.
	BloomConfirms = 256

	// CHTFrequency is the block frequency for creating CHTs
	CHTFrequency = 32768

	// BloomTrieFrequency is the block frequency for creating BloomTrie on both
	// server/client sides.
	BloomTrieFrequency = 32768

	// HelperTrieConfirmations is the number of confirmations before a client is expected
	// to have the given HelperTrie available.
	HelperTrieConfirmations = 2048

	// HelperTrieProcessConfirmations is the number of confirmations before a HelperTrie
	// is generated
	HelperTrieProcessConfirmations = 256

	// CheckpointFrequency is the block frequency for creating checkpoint
	CheckpointFrequency = 32768

	// CheckpointProcessConfirmations is the number before a checkpoint is generated
	CheckpointProcessConfirmations = 256

	// FullImmutabilityThreshold is the number of blocks after which a chain segment is
	// considered immutable (i.e. soft finality). It is used by the downloader as a
	// hard limit against deep ancestors, by the blockchain against deep reorgs, by
	// the freezer as the cutoff threshold and by clique as the snapshot trust limit.
	FullImmutabilityThreshold = 90000

	// LightImmutabilityThreshold is the number of blocks after which a header chain
	// segment is considered immutable for light client(i.e. soft finality). It is used by
	// the downloader as a hard limit against deep ancestors, by the blockchain against deep
	// reorgs, by the light pruner as the pruning validity guarantee.
	LightImmutabilityThreshold = 30000
)
