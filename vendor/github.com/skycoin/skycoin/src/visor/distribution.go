package visor

import "github.com/skycoin/skycoin/src/coin"

const (
	// Maximum supply of skycoins
	MaxCoinSupply uint64 = 3e8 // 100,000,000 million

	// Number of distribution addresses
	DistributionAddressesTotal uint64 = 100

	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// Initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100

	// Number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5

	// Unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

func init() {
	if MaxCoinSupply%DistributionAddressesTotal != 0 {
		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
	}
}

// Returns a copy of the hardcoded distribution addresses array.
// Each address has 1,000,000 coins. There are 100 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (30) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (30) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (30) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// Returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"hEXXeSCzxSiRvLA5dyF52dtBCBZyFxUK77",
	"AyPmfqbdxiSMXstvN6vsFHjwKfS1wBfhHd",
	"y77vYGuABPQ9aYKQQPHggUf4UMAyM7NJZc",
	"2fgBXhctAcfhFURqo7YvkocxRfdkJKUGmep",
	"2SRnQQ7W7ctJH29bnkwdL84dSZmkuYHty5q",
	"27eMsgc7vMBnNxATHqHhqndtPbBqy6SCE9S",
	"e3e52T5sY8iURaxF88AJXNpYQY2cdVDd4y",
	"2Qi6qvdrp4wNW2ziiK7H7LtrHEfyGkibQRU",
	"2fRi3e3G6aJxqLtUZCDZRmwtDe8t4odfUfK",
	"2DrzmC3h3vZJ74Mz2ar9Dk5e9KyMLKe3wNU",
	"2XQCszSpi9qSA7j2G1Hrnin1FohCsvuwrXM",
	"2F1DZcqZvjfs2BAUfgJEfP8XmRdmhFsCkMU",
	"21NbaFfajDLJJ6fH8CK26usaVjBotG95JWM",
	"28nDBkktniC8HLWpTMzDCdTSb91ai7KwDG",
	"2bATyRqkHC3QspYP6ggEUy9wGCZKoY98bvi",
	"UXfpFqnVHjNfabadjXZCTBsFwjomKvbpPT",
	"h9cLi2XvC12iYPdRYxrr2bCEVztQya9o8",
	"9hnctyuThX1VNwmmBFnohhdYr8xViUsTfY",
	"hSDoGxsu6NKTUTkFK7jFPopUXtuTUbJ8kf",
	"2TGvP1JYxqartjgLArmmsDvSogfsgzxLeYx",
	"27Qj8puyVQhhoPj4QJtqEpgTyfivSRf9KCF",
	"TCD1E9rWcZxEWKhrLgtABBevu34okJSLZp",
	"AiHBC557Dhz4xFnC11DN3nJkdGbzj3fQQC",
	"5kSeJHMMm96HPdNJhk4F7d2cVrXgJKNZSQ",
	"26uPwFHXEZZKgUBx8dhBdkcULGDkYAPsYAf",
	"2NHq2sM3yBJzDysyRq9WsKqPAqEciBnyXcd",
	"bgNhaJi4uh5G7Z2EocQrynr4BcdX1VvnW",
	"XC8oTMkNcZYQQWoquk7BixsCdvGph7tyqw",
	"2euHDW7Cn1HCqtVBvkWL1EtcyRv8MU1Evnj",
	"2JxBJZj9cJBJSXpsGbZRpibi8TPZLU9TJSN",
	"77t4D18uTxnV1yzBfQvUH95Uj9DAeTZGxJ",
	"MrVLrr6psQHN3Zt3y3VcshRQo3R6AajF5C",
	"2DYoTyTsyGSjwD7j5VA6zRvrikgL7XLr1d1",
	"WHLKGp2evtaK9PXEp5NqTFMcAQo8XhYaoA",
	"2YHZPCeYihjPFYeQst2mX72PQadj8nkm7id",
	"P3WbBepU8ecy7UDtBWGpQWhtQ3QhLSwoLm",
	"S8t4AtsDteHth68AT3nMQKPfSFfWkAp9JA",
	"LufYMyit4hFzAxNu61vcpYFap5AZsKXxtt",
	"8oTvsaCTuafTsmhKnQx2aLNo4L1VgoTeZo",
	"qGYfjMgKSkBpTJHggStRh87G1a5y1eqST4",
	"2VzZDhdgCVF8k8Nj36TR5jufoHtbDDY2A14",
	"2BXoiBmYUXL4WP4fNzDios1892d5UTpQWmv",
	"2dexNtxop2foiKHSKPpw1vxKs98niLBC31N",
	"CQqyx9qcz2DE76BmKviQUdfANwLQG1knwr",
	"2NYmkgznnPDBGWJnjuZxb8mMyQpCuFpMY2K",
	"Dkpc8uBHrNTbxnZdTqNzYSVPbAqWbXpkTR",
	"2edWeRwqeeoTMKLjvNPpKp27GBQujNVYzzZ",
	"2P7BxrYp4HLk2ap9VHr1UgZDcqH2GetADbx",
	"2m5NLJtuAEY8aqJko9VN7V58viAMU1u2KeW",
	"RovxoH629KGcmJmEf9CyFTXY7CPtWyaezc",
	"Ani9fuagBsW5kEUE2YfJGK5yA3ZXKppCNC",
	"2j1XA366Ai2v7K3E8TD29cXnfpgasHAXPfD",
	"2Jzdkj2tGbrvJH5tyhc7z4FmQff925QWD14",
	"gNAAtTRfUJuzmmiFP1q5BrAwWPozCHHfoA",
	"2YfTTReUSEC5aLnftVCJnxbZR5M9rMv7Mq2",
	"o2JnrAJMfFhYMRCfqXMy9Y6FF9xyhRepNE",
	"2XXamzqRMfn8HRvNWa52suEFF3QndDzZN4n",
	"2Zsi8XXSHJx3U5zJEKYRpjEEym4c2bEunHg",
	"wjSeapuqXZ9yWTd4S2SzT4RHo9gFDfuxCd",
	"s6TkMYUzUscsk4vsZjSNKZtVPKNFLZG7X5",
	"6Sonz4dsmR255fpX2v2eBh2KLsYwXT9Wut",
	"2XB95TGw9YDwGhxeR9kC7HzY3yckyoYCh5L",
	"drW9kN1i7EDyv5wLm3bestYrqFdTi1KkGr",
	"uDTMyWQZg34gLLxLM8LzYfRDqjoAXzdiyA",
	"mF2Lhic9TEw6ByhL2J7eL78LHbnnkbCP5T",
	"Z4m8k2Cwevdk9yQ1froe1KjUurHXjrY57N",
	"2NCRniWkTHkKLBZWLnZZ2Ua36sDyxc9KEF9",
	"TLtgJYeGTZ1dhuDLxhe1bUv6iLxPmvGC2f",
	"2E3F6TJsqg5XniApQbVVXyXEJXNbwNSxgXe",
	"2ZQodyGgnhzatgPj5RLcNKZdzDsZSNcBRyP",
	"2AFFgWvHvTtrLTMQm1pZtL6qfVspqqKmQnJ",
	"jbBgr6derPEwvTp8acLqBjHr9iUFFFSCCe",
	"21dTGB1xuPCuAAYkN1Ld7wGuFaw4fUErjp2",
	"vDR9aW9k769ZWWNYPDJ7u4QQgqNVBckABH",
	"xGMSX1Xr1byrgwuDwZ4JCj6ij3wydbkfqE",
	"2RemJq8DsC2U7DzNcXgBMjpwLyUfetcN67q",
	"241gYPPpAgQYxizQdwqdwGahjQvtFxHZpAj",
	"2iT2b9eKVxDtZcS1hh1BFfoDNSEsGDGit8d",
	"2k8c5mfG9az1g2yJTer1pbHyRyjzk64ccMv",
	"7F9wRaJtyt3drNNer73Aen7CtBDSUKDaAp",
	"kZ1aFrgogN3GENp8jmHekXf12G2JrvbWzd",
	"22Swg1FBF4EpXSsyQxNwiHTijKx7jbT7x2q",
	"2jawfyd6ucDGaCQE1Ay1rsHjcVkSqGKdWPs",
	"frgScGwzLdrz9qoMFCsrDuoiwY9ZvYfvxH",
	"21okSv2mC2Cs6dUEmGXPfe5aR9vkAXtQ7Wj",
	"2CQGWyrSBmho9y7DczzYPjeB9eGGjt16HaF",
	"pHnYvkZa36y39yFoN1KoXwSy4h88pjdDsW",
	"2BndgDYdJEJQwYUV93EBABqrKLVcH6KVkgE",
	"Gc8hHrXCjmiPFF6Tifh8Y6PhPtrV8rGegX",
	"22r3fNondEXWLXnHzP53XCSTtioPi6tSFQU",
	"PaK9AxvqPkCmiwRfiT6xgUueD5yMRUsajN",
	"j9eKjL7CPmktDHRpANDZ5qmXrzycbhyJjy",
	"2hUGbnBFVYjdGrUpyDzV2UAqSstFU9Fmj4T",
	"tw7CPNuAoJn15vU2gz1f5vBsaGAPaYYrb1",
	"2Uw21WjojURavuxgqJtCqKH3sqMMB139nxA",
	"2C8jjSqrZoUhZnMr6TqM8QUjEGQgaKsHvsW",
	"22doUCc71aTfqnhP9THuz8zkbp2J5LRJLWG",
	"q9WFaaYeiePmFZnwxk9qMAwX4rCEFM2QnM",
	"V65ZH7sxS7jeGWgtygm3z6m8RPZbAFPusJ",
	"wPnZYVVsoZUiQov85gZyt9XR71GY9rnfjY",
}
