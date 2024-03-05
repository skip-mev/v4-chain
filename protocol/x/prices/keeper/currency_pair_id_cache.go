package keeper

// CurrencyPairIDCache handles the caching logic of currency-pairs to their corresponding IDs
type CurrencyPairIDCache struct {
	// ID -> CurrencyPair
	idToCurrencyPair map[uint64]string // ID -> CurrencyPair
	// CurrencyPair -> ID
	currencyPairToID map[string]uint64 // CurrencyPair -> ID
}

// NewCurrencyPairIDCache creates a new CurrencyPairIDCache
func NewCurrencyPairIDCache() *CurrencyPairIDCache {
	return &CurrencyPairIDCache{
		idToCurrencyPair: make(map[uint64]string),
		currencyPairToID: make(map[string]uint64),
	}
}

// AddCurrencyPair adds a currency pair to the cache
func (c *CurrencyPairIDCache) AddCurrencyPair(id uint64, currencyPair string) {
	c.idToCurrencyPair[id] = currencyPair
	c.currencyPairToID[currencyPair] = id
}

// GetCurrencyPairFromID returns the currency pair from the cache
func (c *CurrencyPairIDCache) GetCurrencyPairFromID(id uint64) (string, bool) {
	currencyPair, found := c.idToCurrencyPair[id]
	return currencyPair, found
}

// GetIDForCurrencyPair returns the ID for the currency pair from the cache
func (c *CurrencyPairIDCache) GetIDForCurrencyPair(currencyPair string) (uint64, bool) {
	id, found := c.currencyPairToID[currencyPair]
	return id, found
}
