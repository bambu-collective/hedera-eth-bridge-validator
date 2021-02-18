/*
 * Copyright 2021 LimeChain Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ethereum

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/limechain/hedera-eth-bridge-validator/app/clients/ethereum/contracts/bridge"
	"github.com/limechain/hedera-eth-bridge-validator/app/helper"
	"github.com/limechain/hedera-eth-bridge-validator/proto"
	"math/big"
	"strconv"
	"strings"
)

const (
	MintFunctionParameterAmount        = "amount"
	MintFunctionParameterFee           = "fee"
	MintFunctionParameterReceiver      = "receiver"
	MintFunctionParameterSignatures    = "signatures"
	MintFunctionParameterTransactionId = "transactionId"
)

const (
	MintFunction                = "mint"
	MintFunctionParametersCount = 5
)

const (
	InvalidMintFunctionPropertyMessage = "invalid Mint function property: [%s]."
)

var (
	ErrorInvalidMintFunctionParameters = errors.New("invalid mint function parameters length")
)

func generateArguments() (abi.Arguments, error) {
	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, err
	}

	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}

	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}

	return abi.Arguments{
		{
			Type: bytesType,
		},
		{
			Type: addressType,
		},
		{
			Type: uint256Type,
		},
		{
			Type: uint256Type,
		}}, nil
}

func EncodeData(ctm *proto.CryptoTransferMessage) ([]byte, error) {
	args, err := generateArguments()
	if err != nil {
		return nil, err
	}

	amountBn, err := helper.ToBigInt(strconv.Itoa(int(ctm.Amount)))
	if err != nil {
		return nil, err
	}
	feeBn, err := helper.ToBigInt(ctm.Fee)
	if err != nil {
		return nil, err
	}

	return args.Pack(
		[]byte(ctm.TransactionId),
		common.HexToAddress(ctm.EthAddress),
		amountBn,
		feeBn)
}

func DecodeSignature(signature string) (decodedSignature []byte, ethSignature string, err error) {
	decodedSig, err := hex.DecodeString(signature)
	if err != nil {
		return nil, "", err
	}

	return switchSignatureValueV(decodedSig)
}

func DecodeBridgeMintFunction(data []byte) (transferMessage *proto.CryptoTransferMessage, signatures [][]byte, err error) {
	bridgeAbi, err := abi.JSON(strings.NewReader(bridge.BridgeABI))
	if err != nil {
		return nil, nil, err
	}

	// bytes transactionId, address receiver, uint256 amount, uint256 fee, bytes[] signatures
	decodedParameters := make(map[string]interface{})
	err = bridgeAbi.Methods[MintFunction].Inputs.UnpackIntoMap(decodedParameters, data[4:])
	if err != nil {
		return nil, nil, err
	}

	if len(decodedParameters) != MintFunctionParametersCount {
		return nil, nil, ErrorInvalidMintFunctionParameters
	}

	transactionId, ok := decodedParameters[MintFunctionParameterTransactionId].([]byte)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf(InvalidMintFunctionPropertyMessage, MintFunctionParameterTransactionId))
	}

	receiver, ok := decodedParameters[MintFunctionParameterReceiver].(common.Address)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf(InvalidMintFunctionPropertyMessage, MintFunctionParameterReceiver))
	}

	amount, ok := decodedParameters[MintFunctionParameterAmount].(*big.Int)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf(InvalidMintFunctionPropertyMessage, MintFunctionParameterAmount))
	}

	fee, ok := decodedParameters[MintFunctionParameterFee].(*big.Int)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf(InvalidMintFunctionPropertyMessage, MintFunctionParameterFee))
	}

	signatures, ok = decodedParameters[MintFunctionParameterSignatures].([][]byte)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf(InvalidMintFunctionPropertyMessage, MintFunctionParameterAmount))
	}

	var decodedSignatures [][]byte
	for _, sig := range signatures {
		decodedSig, _, err := switchSignatureValueV(sig)
		if err != nil {
			return nil, nil, err
		}
		decodedSignatures = append(decodedSignatures, decodedSig)
	}

	transferMessage = &proto.CryptoTransferMessage{
		TransactionId: string(transactionId),
		EthAddress:    receiver.String(),
		Amount:        amount.Uint64(),
		Fee:           fee.String(),
	}

	return transferMessage, signatures, nil
}

func GetAddressBySignature(hash []byte, signature []byte) (string, error) {
	key, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		return "", err
	}

	pubKey, err := crypto.UnmarshalPubkey(key)
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*pubKey).String(), nil
}

func switchSignatureValueV(decodedSig []byte) (decodedSignature []byte, ethSignature string, err error) {
	if len(decodedSig) != 65 {
		return nil, "", errors.New("invalid signature length")
	}

	// note: https://github.com/ethereum/go-ethereum/issues/19751
	ethSig := make([]byte, len(decodedSig))
	copy(ethSig, decodedSig)

	if decodedSig[64] == 0 || decodedSig[64] == 1 {
		ethSig[64] += 27
	} else if decodedSig[64] == 27 || decodedSig[64] == 28 {
		decodedSig[64] -= 27
	}

	return decodedSig, hex.EncodeToString(ethSig), nil
}
