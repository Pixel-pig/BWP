package util

import (
	"CryptoHashCode/Base58"
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func Address(){
	//1.生成私钥 ECC
	curve := elliptic.P256()
	_, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println("私钥：",pri)
	//公钥拼接，拼接x和y生成公钥
	//pubkey := append(x.Bytes(), y.Bytes()...)
	publicKey:=elliptic.Marshal(curve,x,y)
	//sha256哈希计算
	sha256Hash := sha256.New()
	sha256Hash.Write(publicKey)
	pub256Hash := sha256Hash.Sum(nil)
	//ripemd160计算(github下载ripemd160库)
	ripemd160 := ripemd160.New()
	ripemd160.Write(pub256Hash)
	pubRpmd160 := ripemd160.Sum(nil)
	//2.添加版本号
	versionPubRpmd160 := append([]byte{0x00}, pubRpmd160...)
	//3.计算校验位
	sha256Hash.Reset() //重置
	sha256Hash.Write(versionPubRpmd160)
	hash1 := sha256Hash.Sum(nil)
	sha256Hash.Reset()
	sha256Hash.Write(hash1)
	hash2 := sha256Hash.Sum(nil)
	//取前四位字节,形式：[:],规则：前闭后开
	check := hash2[:4]
	//4.地址拼接
	addBytes := append(versionPubRpmd160, check...)
	//5.形式变化：地址的拼接结果形式进行base58编码
	//fmt.Println("地址拼接:",addBytes)
	address := Base58.Encode(addBytes)
	fmt.Println("新生成的比特币地址:", address)
	//6.校验比特币地址是否有效
	//6.1对地址base58解码
	deaddBytes := Base58.Decode(address)
	//6.2取后4位作为校验位
	check2 := deaddBytes[len(deaddBytes)-4:]
	//6.3双hash计算
	//6.3.a.deAddBytesExpelCheck获取 地址反编码 去除四个字节 以后的内容
	deAddBytesExpelCheck := deaddBytes[:len(deaddBytes)-len(check2)]
	//6.3.b.deHash1:=sha256=sha256(deAddBytesExpelCheck)
	sha256Hash.Reset()
	sha256Hash.Write(deAddBytesExpelCheck)
	deHash1 := sha256Hash.Sum(nil)
	//6.3.c.deHash2"=sha256(deHash2)
	sha256Hash.Reset()
	sha256Hash.Write(deHash1)
	deHash2 := sha256Hash.Sum(nil)
	//6.4去前四个字节
	//deCheck = deHash2[前四个字节]
	deCheck := deHash2[:4]
	//7.将check2与deCheck进行比较
	if bytes.Compare(deCheck,check2)==0{
		fmt.Println("该地址校验成功")
	}else {
		fmt.Println("该地址校验失败")
	}
}


