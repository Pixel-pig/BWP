package util

import (
	"BTCproject/entlity"
	"fmt"
	"net/http"
)

//获取最新区块的hash
func GetBestBlockHash() string {
	reqbytes := PrepareJSON(entlity.GETBESTBLOCKHASH)
	response, body := DoPost(reqbytes)
	if response == nil {
		return ""
	}
	if response.StatusCode == http.StatusOK {
		besthash := Analysis(response, body)
		fmt.Println("最新区块hash:", besthash.Result)
		return besthash.Result.(string)
	}
	return ""
}

//该方法用于获取比特币当前节点的区块数
func GetBlockCount() float64 {
	reqbytes := PrepareJSON(entlity.GETBLOCKCOUNT)
	response, body := DoPost(reqbytes)
	if response == nil {
		return -1
	}
	if response.StatusCode == http.StatusOK {
		count := Analysis(response, body)
		fmt.Println("区块的总数:", count.Result)
		return count.Result.(float64)
	}
	return -1
}

//根据区块hash获取区块信息
func GetBlock(params interface{}) *entlity.Block {
	reqbytes := PrepareJSON(entlity.GETBLOCK, params)
	response, body := DoPost(reqbytes)
	if response == nil {
		return nil
	}
	if response.StatusCode == http.StatusOK {
		data := Analysis(response, body)
		blockMap := data.Result.(map[string]interface{})
		// mapstructure.Decode(blockMap,entlity.Block{})
		block := entlity.Block{
			Bits:          blockMap["bits"].(string),
			Hash:          blockMap["hash"].(string),
			Confirmations: blockMap["confirmations"].(float64),
			Strippedsize:  blockMap["strippedsize"].(float64),
			Size:          blockMap["size"].(float64),
			Weight:        blockMap["weight"].(float64),
			Height:        blockMap["height"].(float64),
			Version:       blockMap["version"].(float64),
			//VersionHex:    blockMap["version_hex"].(string),
			Merkleroot: blockMap["merkleroot"].(string),
			Time:       blockMap["time"].(float64),
			Mediantime: blockMap["mediantime"].(float64),
			Nonce:      blockMap["nonce"].(float64),
			Difficulty: blockMap["difficulty"].(float64),
			Chainwork:  blockMap["chainwork"].(string),
			NTx:        blockMap["nTx"].(float64),
			//			Nextblockhash: blockMap["nextblockhash"].(string),
		}
		return &block
	}
	return nil
}

//根据指定高度获取该高度区块的hash
func GetBlockHash(height interface{}) string {
	reqbytes := PrepareJSON(entlity.GETBLOCKHASH, height)
	response, body := DoPost(reqbytes)
	fmt.Println("===111", response)
	if response == nil {
		return ""
	}
	if response.StatusCode == http.StatusOK {
		hash := Analysis(response, body)
		fmt.Println("指定区块的hash:", hash.Result)
		return hash.Result.(string)
	}
	return ""
}

//获取区块链节点的信息
func GetBlockChainInfo()*entlity.Blockinfo {
	reqbytes := PrepareJSON(entlity.GETBLOCKCHAININFO)
	response, body := DoPost(reqbytes)
	if response == nil {
		return nil
	}
	if response.StatusCode == http.StatusOK {
		blockinfo := Analysis(response, body)
		fmt.Println("区块链节点的信息:", blockinfo.Result)
		blockInfoMap := blockinfo.Result.(map[string]interface{})
		info := entlity.Blockinfo{
			Headers:              blockInfoMap["headers"].(float64),
			Chain:                blockInfoMap["chain"].(string),
			SizeOnDisk:           blockInfoMap["size_on_disk"].(float64),
			Mediantime:           blockInfoMap["mediantime"].(float64),
			Blocks:               blockInfoMap["blocks"].(float64),
			Pruned:               blockInfoMap["pruned"].(bool),
			Warnings:             blockInfoMap["warnings"].(string),
			Difficulty:           blockInfoMap["difficulty"].(float64),
			Chainwork:            blockInfoMap["chainwork"].(string),
			Verificationprogress: blockInfoMap["verificationprogress"].(float64),
			Bestblockhash:        blockInfoMap["bestblockhash"].(string),
			Initialblockdownload: blockInfoMap["initialblockdownload"].(bool),
		}
		return &info
	}
	return nil
}
//获取当前区块难度
func Getdifficulty() float64 {
	reqbytes := PrepareJSON(entlity.GETDIFFICULTY)
	response, body := DoPost(reqbytes)
	if response == nil {
		return -1
	}
	if response.StatusCode == http.StatusOK {
		difficulty := Analysis(response, body)
		return difficulty.Result.(float64)
	}
	return -1
}