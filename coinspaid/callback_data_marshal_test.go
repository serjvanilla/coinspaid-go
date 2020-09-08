package coinspaid_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/serjvanilla/coinspaid-go/coinspaid"

	"github.com/stretchr/testify/assert"
)

func TestCallbackDataMarshaling(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{
			name: "Deposit BTC",
			data: []byte(`{"id":1,"type":"deposit","crypto_address":{"id":1,"currency":"BTC","address":"39mFf3X46YzUtfdwVQpYXPCMydc74ccbAZ","foreign_id":"user-id:2048","tag":null},"currency_sent":{"currency":"BTC","amount":"6.53157512"},"currency_received":{"currency":"BTC","amount":"6.53157512","amount_minus_fee":"6.51198"},"transactions":[{"id":1,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"39mFf3X46YzUtfdwVQpYXPCMydc74ccbAZ","tag":null,"amount":"6.53157512","txid":"3950ad8149421a850d01dff88f024810e363ac18c9e8dd9bc0b9116e7937ad93","riskscore":"0.5","confirmations":"3"}],"fees":[{"type":"deposit","currency":"BTC","amount":"0.01959472"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Deposit ETH",
			data: []byte(`{"id":2686563,"type":"deposit","crypto_address":{"id":381738,"currency":"ETH","address":"0xd61180ff0cf74dc3ee8e264751f18c47060729b9","tag":null,"foreign_id":"991904"},"currency_sent":{"currency":"ETH","amount":"0.01"},"currency_received":{"currency":"ETH","amount":"0.01","amount_minus_fee":"0.009395"},"transactions":[{"id":714657,"currency":"ETH","transaction_type":"blockchain","type":"deposit","address":"0xd61180ff0cf74dc3ee8e264751f18c47060729b9","tag":null,"amount":"0.01","txid":"0x6b353da88a8ba2df4926c1ccc58991f484a683ba57ec3dde70e812b5c8c7fa1d","confirmations":"9"}],"fees":[{"type":"transfer","currency":"ETH","amount":"0.000105"},{"type":"deposit","currency":"ETH","amount":"0.0005"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Deposit ERC20",
			data: []byte(`{"id":2686567,"type":"deposit","crypto_address":{"id":382357,"currency":"NNM","address":"0x19722627885da3fff134dd36de0a14898c8e053b","tag":null,"foreign_id":"31"},"currency_sent":{"currency":"NNM","amount":"0.06"},"currency_received":{"currency":"NNM","amount":"0.06","amount_minus_fee":"0.0591"},"transactions":[{"id":714662,"currency":"NNM","transaction_type":"blockchain","type":"deposit","address":"0x19722627885da3fff134dd36de0a14898c8e053b","tag":null,"amount":"0.06","txid":"0x835c9f286a311bc9df8c94c2771cadf1d91e6967039358a44e24c56115dc59a1","confirmations":"7"}],"fees":[{"type":"transfer","currency":"ETH","amount":"0.00037336"},{"type":"deposit","currency":"NNM","amount":"0.0009"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Deposit BTC with exchange to EUR",
			data: []byte(`{"id":2686510,"type":"deposit_exchange","crypto_address":{"id":382270,"currency":"BTC","convert_to":"EUR","address":"31vnLqxVJ1iShJ5Ly586q8XKucECx12bZS","tag":null,"foreign_id":"13a"},"currency_sent":{"currency":"BTC","amount":"0.01"},"currency_received":{"currency":"EUR","amount":"84.17070222","amount_minus_fee":"79.96216711"},"transactions":[{"id":714576,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"31vnLqxVJ1iShJ5Ly586q8XKucECx12bZS","tag":null,"amount":"0.01","txid":"3a491da90a1ce5a318d0aeff6867ab98a03219abae29ed68d702291703c3538b","riskscore":"0.42","confirmations":"1"},{"id":714577,"currency":"BTC","currency_to":"EUR","transaction_type":"exchange","type":"exchange","tag":null,"amount":"0.01","amount_to":"84.17070222","txid":null,"confirmations":"0"}],"fees":[{"type":"exchange","currency":"EUR","amount":"4.20853511"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Withdraw BTC",
			data: []byte(`{"id":1,"foreign_id":"10","type":"withdrawal","crypto_address":{"id":1,"currency":"BTC","address":"115Mn1jCjBh1CNqug7yAB21Hq2rw8PfmTA","tag":null},"currency_sent":{"currency":"BTC","amount":"0.02"},"currency_received":{"currency":"BTC","amount":"0.02"},"transactions":[{"id":1,"currency":"BTC","transaction_type":"blockchain","type":"withdrawal","address":"115Mn1jCjBh1CNqug7yAB21Hq2rw8PfmTA","tag":null,"amount":"0.02","txid":"bb040d895ef7141ea0b06b04227d8f5dd4ee12d5b890e6e5633f6439393a666b","confirmations":"3"}],"fees":[{"type":"mining","currency":"BTC","amount":"0.00007402"},{"type":"withdrawal","currency":"BTC","amount":"0.00002"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Withdraw ETH",
			data: []byte(`{"id":2686565,"foreign_id":"23","type":"withdrawal","crypto_address":{"id":381460,"currency":"ETH","address":"0x2D6CA312567986C08CC4eF3F706136D1c9eF0321","tag":null},"currency_sent":{"currency":"ETH","amount":"0.01"},"currency_received":{"currency":"ETH","amount":"0.01"},"transactions":[{"id":714660,"currency":"ETH","transaction_type":"blockchain","type":"withdrawal","address":"0x2D6CA312567986C08CC4eF3F706136D1c9eF0321","tag":null,"amount":"0.01","txid":"0x1ccb9fa0ef5e8cf7d4cc2a23fe8119170a2a3d08fba36036665a12c88d7bcccb","confirmations":"0"}],"fees":[{"type":"mining","currency":"ETH","amount":"0.000105"},{"type":"withdrawal","currency":"ETH","amount":"0.0001"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Withdraw ERC20",
			data: []byte(`{"id":2686572,"foreign_id":"5","type":"withdrawal","crypto_address":{"id":381734,"currency":"NNM","address":"0x2D6CA312567986C08CC4eF3F706136D1c9eF0321","tag":null},"currency_sent":{"currency":"NNM","amount":"4.456"},"currency_received":{"currency":"NNM","amount":"4.456"},"transactions":[{"id":714670,"currency":"NNM","transaction_type":"blockchain","type":"withdrawal","address":"0x2D6CA312567986C08CC4eF3F706136D1c9eF0321","tag":null,"amount":"4.456","txid":"0x884ad10cc60dfe0d6fdc776b541c5c5efce6151886a88994bb4ba41aa9575563","confirmations":"0"}],"fees":[{"type":"mining","currency":"ETH","amount":"0.0000748"},{"type":"withdrawal","currency":"NNM","amount":"0.04456"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Withdraw EUR with exchange to BTC",
			data: []byte(`{"id":1,"foreign_id":"20","type":"withdrawal_exchange","crypto_address":{"id":1,"currency":"EUR","convert_to":"BTC","address":"1K2btnZ8cqNFBPhaq729Mdj8W6G3w2nBbL","tag":null},"currency_sent":{"currency":"EUR","amount":"381"},"currency_received":{"currency":"BTC","amount":"0.108823"},"transactions":[{"id":1,"currency":"EUR","currency_to":"BTC","transaction_type":"exchange","type":"exchange","tag":null,"amount":"381","amount_to":"0.108823","txid":null,"confirmations":"0"},{"id":1,"currency":"BTC","transaction_type":"blockchain","type":"withdrawal","address":"1K2btnZ8cqNFBPhaq729Mdj8W6G3w2nBbL","tag":null,"amount":"0.108823","txid":"aa3345b96389e126f1ce88a670d1b1e38f2c3f73fb3ecfff8d9da1b1ce6964a6","confirmations":"3"}],"fees":[{"type":"exchange","currency":"EUR","amount":"3.048"},{"type":"mining","currency":"EUR","amount":"0.0448978"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Buy BTC for EUR",
			data: []byte(`{"id":2686900,"type":"exchange","currency_sent":{"currency":"EUR","amount":"26.75248865"},"currency_received":{"currency":"BTC","amount":"0.003"},"transactions":[{"id":715072,"currency":"EUR","currency_to":"BTC","transaction_type":"exchange","type":"exchange","tag":null,"amount":"26.75248865","amount_to":"0.003","txid":null,"confirmations":"0"}],"fees":[{"type":"exchange","currency":"EUR","amount":"0.26752489"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Sell BTC for USD",
			data: []byte(`{"id":2686901,"type":"exchange","currency_sent":{"currency":"BTC","amount":"0.003"},"currency_received":{"currency":"USD","amount":"29.50408884"},"transactions":[{"id":715072,"currency":"BTC","currency_to":"USD","transaction_type":"exchange","type":"exchange","tag":null,"amount":"0.003","amount_to":"29.50408884","txid":null,"confirmations":"0"}],"fees":[{"type":"exchange","currency":"USD","amount":"0.29504089"}],"error":"","status":"confirmed"}`),
		},
		{
			name: "Buy BTC for ETH future",
			data: []byte(`{"id":2688873,"type":"deposit_exchange","crypto_address":{"id":384708,"currency":"ETH","convert_to":"BTC","address":"0x4b41a526d3d12de36bdf969e7b70fd0bd2e0d263","tag":null,"foreign_id":"hfjs781"},"currency_sent":{"currency":"ETH","amount":"0.11129846"},"currency_received":{"currency":"BTC","amount":"0.003","amount_minus_fee":"0.00299944"},"transactions":[{"id":717555,"currency":"ETH","transaction_type":"blockchain","type":"deposit","address":"0x4b41a526d3d12de36bdf969e7b70fd0bd2e0d263","tag":null,"amount":"0.11129846","txid":"0x19f8094e12dfc6cb14910d6057269d10f39dfdc7c8b0d0e22b789c3e5d03b9e5","confirmations":"13"},{"id":717556,"currency":"ETH","currency_to":"BTC","transaction_type":"exchange","type":"exchange","tag":null,"amount":"0.11129846","amount_to":"0.003","txid":null,"confirmations":"0"}],"fees":[{"type":"transfer","currency":"BTC","amount":"0.00000056"},{"type":"exchange","currency":"ETH","amount":"0.00667791"}],"error":"","status":"confirmed","futures_id":95,"transaction_id":2688873}`),
		},
		{
			name: "Invoice payment by installments",
			data: []byte(`{"id":588,"foreign_id":"8FW1KI7LesB9yxWcK1K","type":"invoice","crypto_address":{"id":386897,"currency":"BTC","address":"2Mvo6FMduhHz1BTDHsQ5GyRoifcG3y4ycpk","tag":null},"currency_sent":{"currency":"BTC","amount":"0.02","remaining_amount":"0.01"},"currency_received":{"currency":"BTC","amount":"0.02"},"transactions":[{"id":750294,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"2Mvo6FMduhHz1BTDHsQ5GyRoifcG3y4ycpk","tag":null,"amount":"0.01","txid":"528dcda13270f8590853405600bf5634d53aa66d2ce5d3a873006a670f9da788","confirmations":"0"}],"fees":[],"error":"","status":"pending","fixed_at":1592307241,"expires_at":1592308141}`),
		},
		{
			name: "Transaction is in the mempool",
			data: []byte(`{"id":22,"foreign_id":"229-hdsa","type":"invoice","crypto_address":{"id":1845,"currency":"BTC","convert_to":"EUR","address":"2N8pvnVKEFjVCP4TteJVCPTbg9Aqibfws8C","tag":null},"currency_sent":{"currency":"BTC","amount":"0.00309556","remaining_amount":"0"},"currency_received":{"currency":"EUR","amount":"26"},"transactions":[{"id":1504,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"2N8pvnVKEFjVCP4TteJVCPTbg9Aqibfws8C","tag":null,"amount":"0.00309556","txid":"3e68744626ba23d591dc1b548b8b9b126ab0f49e6da020958f66ee64c38ec7ed","confirmations":"0"}],"fees":[],"error":"","status":"processing","fixed_at":1592308917,"expires_at":1592395375}`),
		},
		{
			name: "Successful payment",
			data: []byte(`{"id":588,"foreign_id":"8FW1KI7LesB9yxWcK1K","type":"invoice","crypto_address":{"id":386897,"currency":"BTC","address":"2Mvo6FMduhHz1BTDHsQ5GyRoifcG3y4ycpk","tag":null},"currency_sent":{"currency":"BTC","amount":"0.02","remaining_amount":"0"},"currency_received":{"currency":"BTC","amount":"0.02"},"transactions":[{"id":750294,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"2Mvo6FMduhHz1BTDHsQ5GyRoifcG3y4ycpk","tag":null,"amount":"0.01","txid":"6647cf5cae701507b7076b32ca12f19d8e9fe037407c02e09b04abdaede99fd0","confirmations":"1"},{"id":750384,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"2Mvo6FMduhHz1BTDHsQ5GyRoifcG3y4ycpk","tag":null,"amount":"0.01","txid":"528dcda13270f8590853405600bf5634d53aa66d2ce5d3a873006a670f9da788","confirmations":"1"}],"fees":[{"type":"fee_crypto_deposit","currency":"BTC","amount":"0.0008"}],"error":"","status":"confirmed","fixed_at":1592307241,"expires_at":1592394104}`),
		},
		{
			name: "Timer expired",
			data: []byte(`{"id":23,"foreign_id":"77xa2pd","type":"invoice","crypto_address":{"id":1846,"currency":"BTC","address":"2MyL2ftBdNsmpsx1EKggPzNVuoiZCWdVaLX","tag":null},"currency_sent":{"currency":"BTC","amount":"0.02","remaining_amount":"0.02"},"currency_received":{"currency":"BTC","amount":"0.02"},"transactions":[],"fees":[],"error":"Timer expired. User not paid.","status":"failed","fixed_at":1592309817,"expires_at":1592310717}`),
		}, {
			name: "Processing status for more than 24 hours",
			data: []byte(`{"id":21,"foreign_id":"88smaan2","type":"invoice","crypto_address":{"id":1844,"currency":"BTC","address":"2NCfH4keMq2SZj3H73VeQYtb8yBgTuUrACf","tag":null},"currency_sent":{"currency":"BTC","amount":"0.01","remaining_amount":"0.0099"},"currency_received":{"currency":"BTC","amount":"0.01"},"transactions":[{"id":1505,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"2NCfH4keMq2SZj3H73VeQYtb8yBgTuUrACf","tag":null,"amount":"0.0001","txid":"6dd68265eb2b3a1be0dd4a40975345155d0fbbe244c84efcc7f6b910629a38f9","confirmations":"0"}],"fees":[],"error":"Timer expired. Transactions were in status processing too long.","status":"failed","fixed_at":1592308709,"expires_at":1592309609}`),
		},
		{
			name: "Payment is lesser than invoice amount",
			data: []byte(`{"id":21,"foreign_id":"88smaan2","type":"invoice","crypto_address":{"id":1844,"currency":"BTC","address":"2NCfH4keMq2SZj3H73VeQYtb8yBgTuUrACf","tag":null},"currency_sent":{"currency":"BTC","amount":"0.01","remaining_amount":"0.0099"},"currency_received":{"currency":"BTC","amount":"0.01"},"transactions":[{"id":1505,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"2NCfH4keMq2SZj3H73VeQYtb8yBgTuUrACf","tag":null,"amount":"0.0001","txid":"6dd68265eb2b3a1be0dd4a40975345155d0fbbe244c84efcc7f6b910629a38f9","confirmations":"2"}],"fees":[],"error":"Timer expired. User paid less than requested.","status":"failed","fixed_at":1592308709,"expires_at":1592309609}`),
		},
		{
			name: "Not confirmed",
			data: []byte(`{"id":2686579,"type":"deposit","crypto_address":{"id":381711,"currency":"BTC","address":"2N9zXNdiT8ucZp7zZSrucqYGCD6xYF8F3di","tag":null,"foreign_id":"991904"},"currency_sent":{"currency":"BTC","amount":"0.01"},"currency_received":{"currency":"BTC","amount":"0.01","amount_minus_fee":"0.01"},"transactions":[{"id":714680,"currency":"BTC","transaction_type":"blockchain","type":"deposit","address":"2N9zXNdiT8ucZp7zZSrucqYGCD6xYF8F3di","tag":null,"amount":"0.01","txid":"998c4d9bb7145aafd88658b292f41fe05973c217f7adcd6052bcafe2309e7e02","confirmations":"0"}],"fees":[],"error":"","status":"not_confirmed"}`),
		},
		{
			name: "Cancelled",
			data: []byte(`{"id":2686580,"foreign_id":"123","type":"withdrawal","crypto_address":{"id":382362,"currency":"ETH","address":"12345","tag":null},"transactions":[{"id":714681,"currency":"ETH","transaction_type":"blockchain","type":"withdrawal","address":"12345","tag":null,"amount":"1","txid":null,"confirmations":"0"}],"fees":[],"error":"Invalid params: expected a hex-encoded hash with 0x prefix.","status":"cancelled"}`),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			indented := new(bytes.Buffer)
			err := json.Indent(indented, test.data, "", "\t")
			if err != nil {
				t.Errorf("src indent error: %v", err)
			}

			var cb coinspaid.CallbackData
			err = json.Unmarshal(test.data, &cb)
			if err != nil {
				t.Errorf("unmarshal error: %v", err)
			}

			newData, err := json.MarshalIndent(cb, "", "\t")
			if err != nil {
				t.Errorf("marshal error: %v", err)
			}

			assert.JSONEq(t, indented.String(), string(newData), "json should be equal")
			// if diff := cmp.Diff(indented.String(), string(newData)); diff != "" {
			// 	t.Errorf("json mismatch (-want +got):\n%s", diff)
			// }
		})
	}
}
