// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package capella_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"

	"github.com/attestantio/go-builder-client/api/capella"
	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func bytestr(input string) []byte {
	res, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}
	return res
}

func TestSignedBuilderBidJSON(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   string
	}{
		{
			name: "Empty",
			err:  "unexpected end of JSON input",
		},
		{
			name:  "JSONBad",
			input: []byte("[]"),
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type capella.signedBuilderBidJSON",
		},
		{
			name:  "MessageMissing",
			input: []byte(`{"signature":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"}`),
			err:   "message missing",
		},
		{
			name:  "MessageWrongType",
			input: []byte(`{"message":true,"signature":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"}`),
			err:   "invalid JSON: invalid JSON: json: cannot unmarshal bool into Go value of type capella.builderBidJSON",
		},
		{
			name:  "MessageInvalid",
			input: []byte(`{"message":{},"signature":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"}`),
			err:   "invalid JSON: header missing",
		},
		{
			name:  "SignatureMissing",
			input: []byte(`{"message":{"header":{"parent_hash":"0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef","fee_recipient":"0x58E809C71e4885cB7B3f1D5c793AB04eD239d779","state_root":"0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d","receipts_root":"0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36","logs_bloom":"0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44","prev_randao":"0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00","block_number":"2983837628677007840","gas_limit":"6738255228996962210","gas_used":"5573520557770513197","timestamp":"1744720080366521389","extra_data":"0xc648","base_fee_per_gas":"88770397543877639215846057887940126737648744594802753726778414602657613619599","block_hash":"0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f","transactions_root":"0x8457d0eb7611a621e7a094059f087415ffcfc91714fc184a1f3c48db06b4d08b","withdrawals_root":"0x5c1a7e3e0eab917a7c0d677c6b692ed55ce05f07d572e09b1c06d558f474ea7a"},"value":"12345","pubkey":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"}}`),
			err:   "signature missing",
		},
		{
			name:  "SignatureWrongType",
			input: []byte(`{"message":{"header":{"parent_hash":"0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef","fee_recipient":"0x58E809C71e4885cB7B3f1D5c793AB04eD239d779","state_root":"0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d","receipts_root":"0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36","logs_bloom":"0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44","prev_randao":"0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00","block_number":"2983837628677007840","gas_limit":"6738255228996962210","gas_used":"5573520557770513197","timestamp":"1744720080366521389","extra_data":"0xc648","base_fee_per_gas":"88770397543877639215846057887940126737648744594802753726778414602657613619599","block_hash":"0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f","transactions_root":"0x8457d0eb7611a621e7a094059f087415ffcfc91714fc184a1f3c48db06b4d08b","withdrawals_root":"0x5c1a7e3e0eab917a7c0d677c6b692ed55ce05f07d572e09b1c06d558f474ea7a"},"value":"12345","pubkey":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"},"signature":true}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field signedBuilderBidJSON.signature of type string",
		},
		{
			name:  "SignatureInvalid",
			input: []byte(`{"message":{"header":{"parent_hash":"0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef","fee_recipient":"0x58E809C71e4885cB7B3f1D5c793AB04eD239d779","state_root":"0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d","receipts_root":"0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36","logs_bloom":"0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44","prev_randao":"0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00","block_number":"2983837628677007840","gas_limit":"6738255228996962210","gas_used":"5573520557770513197","timestamp":"1744720080366521389","extra_data":"0xc648","base_fee_per_gas":"88770397543877639215846057887940126737648744594802753726778414602657613619599","block_hash":"0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f","transactions_root":"0x8457d0eb7611a621e7a094059f087415ffcfc91714fc184a1f3c48db06b4d08b","withdrawals_root":"0x5c1a7e3e0eab917a7c0d677c6b692ed55ce05f07d572e09b1c06d558f474ea7a"},"value":"12345","pubkey":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"},"signature":"invalid"}`),
			err:   "invalid value for signature: encoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "SignatureIncorrectLength",
			input: []byte(`{"message":{"header":{"parent_hash":"0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef","fee_recipient":"0x58E809C71e4885cB7B3f1D5c793AB04eD239d779","state_root":"0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d","receipts_root":"0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36","logs_bloom":"0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44","prev_randao":"0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00","block_number":"2983837628677007840","gas_limit":"6738255228996962210","gas_used":"5573520557770513197","timestamp":"1744720080366521389","extra_data":"0xc648","base_fee_per_gas":"88770397543877639215846057887940126737648744594802753726778414602657613619599","block_hash":"0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f","transactions_root":"0x8457d0eb7611a621e7a094059f087415ffcfc91714fc184a1f3c48db06b4d08b","withdrawals_root":"0x5c1a7e3e0eab917a7c0d677c6b692ed55ce05f07d572e09b1c06d558f474ea7a"},"value":"12345","pubkey":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"},"signature":"0x0101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"}`),
			err:   "incorrect length for signature",
		},
		{
			name:  "Good",
			input: []byte(`{"message":{"header":{"parent_hash":"0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef","fee_recipient":"0x58E809C71e4885cB7B3f1D5c793AB04eD239d779","state_root":"0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d","receipts_root":"0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36","logs_bloom":"0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44","prev_randao":"0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00","block_number":"2983837628677007840","gas_limit":"6738255228996962210","gas_used":"5573520557770513197","timestamp":"1744720080366521389","extra_data":"0xc648","base_fee_per_gas":"88770397543877639215846057887940126737648744594802753726778414602657613619599","block_hash":"0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f","transactions_root":"0x8457d0eb7611a621e7a094059f087415ffcfc91714fc184a1f3c48db06b4d08b","withdrawals_root":"0x5c1a7e3e0eab917a7c0d677c6b692ed55ce05f07d572e09b1c06d558f474ea7a"},"value":"12345","pubkey":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"},"signature":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res capella.SignedBuilderBid
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				assert.Equal(t, string(test.input), string(rt))
			}
		})
	}
}

func TestSignedBuilderBidYAML(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   string
	}{
		{
			name:  "Good",
			input: []byte(`{message: {header: {parent_hash: '0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef', fee_recipient: '0x58E809C71e4885cB7B3f1D5c793AB04eD239d779', state_root: '0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d', receipts_root: '0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36', logs_bloom: '0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44', prev_randao: '0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00', block_number: 2983837628677007840, gas_limit: 6738255228996962210, gas_used: 5573520557770513197, timestamp: 1744720080366521389, extra_data: '0xc648', base_fee_per_gas: '88770397543877639215846057887940126737648744594802753726778414602657613619599', block_hash: '0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f', transactions_root: '0x8457d0eb7611a621e7a094059f087415ffcfc91714fc184a1f3c48db06b4d08b', withdrawals_root: '0x5c1a7e3e0eab917a7c0d677c6b692ed55ce05f07d572e09b1c06d558f474ea7a'}, value: 12345, pubkey: '0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101'}, signature: '0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101'}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res capella.SignedBuilderBid
			err := yaml.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := yaml.Marshal(&res)
				require.NoError(t, err)
				assert.Equal(t, string(rt), res.String())
				rt = bytes.TrimSuffix(rt, []byte("\n"))
				assert.Equal(t, string(test.input), string(rt))
			}
		})
	}
}
