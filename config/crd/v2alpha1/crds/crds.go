/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by crds_generate.go; DO NOT EDIT.

package crds

import (
	"bytes"
	"compress/gzip"
	"io"

	apiextinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

var rawCRDs = [][]byte{
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcYɒ\x1b\xb9\x11\xbd\xf7Wd\xc8\a]D\xf6Ȟp8x\x93\xd8vD\x87G-ư\xddwT!Y\xc4\b\x05\xc0Xȡ\x97\x7f\x9fH\xa0\xaaX\v\xb8\xce\xc2\x1b\x81D\xe2\xe5\x82\xcc\a\xd4l6{`F\xbc\xa1uB\xab\x050#\xf0g\x8f\x8a\xfe\xb9\xf9\xb7\xbf\xb9\xb9Џ\xbb\x8f\x0f߄\xe2\vX\x06\xe7u\xfd#:\x1dl\x89O\xb8\x11Jx\xa1\xd5C\x8d\x9eq\xe6\xd9\xe2\x01\x80)\xa5=\xa3aG\x7f\x01J\xad\xbc\xd5R\xa2\x9dU\xa8\xe6\xdfB\x81E\x10\x92\xa3\x8d\xcaۭw\xdf\xcd?~?\xff\xee\x01@\xb1\x1a\x17@\xfa\xb8\xde+\xa9\x19w\xf3\x1dJ\xb4z.\xf4\x833X\x92\xe2\xca\xea`\x16p\x9cH\v\x9bM\x13\xe0'\xe6\xd9S\xa3#\x0eK\xe1\xfc?'S?\b\xe7㴑\xc129\xda;\xce8\xa1\xaa \x99\x1d\xce=\x00\xb8R\x1b\\\xc0\vmmX\x894\xd6\xd8\x14\xa1̀q\x1e\xbd\xc4\xe4\xca\n\xe5\xd1.\xb5\fu\xeb\x9d\x19pt\xa5\x15\xc6G/\xf4a\x81\xf3\xcc\a\a.\x94[`\x0e^p\xff\xf8\xacVVW\x16]\x82\x05\xf0\x93\xd3j\xc5\xfcv\x01\xf3$>7[氙M\xae\\ǉf\xc8\x1f\b\xaf\xf3V\xa8*\x87\xe0U\xd4\b<\xd8\x18B\xb2\xbbD\xf0[\xe1\x86\xd0\xf6\xcc\x11<둟\x04\x12\xe7I\x9d\xf3\xac6cD\xbd\xa5\t\x12g\x1es\x80\x96\xba6\x12=r(\x0e\x1e[36\xda\xd6\xcc/@(\xff\xd7\xefO\xfb\xa2q\xd6<.}\xd2j\xe8\x98\xcf4\n\xbdᄄ\xa2T\xa1\xcdzG{&\x7f\r\x10O\n>\xf7\xd6'$Io\x7f\xfc\"\x14J9\xd0\x1b\xf0[\x84Ϭ\xfc\x16\f\xac\xbd\xb6\xacB\xf8A\x97)|\xfb-Z\x8c\x12E\x92\xa0\xec\x05A\xb1\xd36\x1b:\x83\xe5<\xc96\xcaZ]\xa3\xf8\r7\xfa\xcds\xab\xb4Ȳ\xb9Ֆ\x9ay\x94\x10Z\xe5\x13\xecS\x85W%W߉Js\xecyl\x80I80V\x97\xe8ܙ\x84'\x05\x03\x14/ǁ\x89k\x92\xc4\xee\xcfL\x9a-\xfb\x98\x8aL\xb9Ś-\x9a\x15ڠ\xfa\xb4z~\xfb\xcbz0\fg\n\x06+\xbd\xa3JA\xf0\x8d\xd5^\x97ZB\x81~\x8f\xa8R\xe8k\xbdCKu\xae\x12\xcau\x1a\xa9j\xf3\xbe\xc0\xb1fS~G}4\x9b&-\xc6\xec!\x80\xb6\x1f}\xa0=\rZ/\xda*\xdc\xe8>6\x98\xde\xe8Ȏ\xff\xcd\x06s\x00dzZ\x05\x9c:\r&\xb3\x9aڊ\xbc\xf1V\n\x9ep`\xd1Xt\xa8R\xef\xa1a\xa6@\x17?a\xe9\xe7#\xd5k\xb4\xa4\x06\xdcV\a\xc9\xc9\xd8\x1dZ\x0f\x16K])\xf1\x9fN\xb7\x03\xaf㦒yt>\x1eF\xab\x98\x84\x1d\x93\x01?\x90\xd3F\x9akv\x00\x8b\xb4'\x04\xd5\xd3\x17\x17\xb81\x8e/\xe4E\xa16z\x01[\xef\x8d[<>V·m\xb7\xd4u\x1d\x94\xf0\x87\xc7\x18\rQ\x04\xaf\xad{\xe4\xb8C\xf9\xe8D5c\xb6\xdc\n\x8f\xa5\x0f\x16\x1f\x99\x11\xb3h\x88\x8a\xadw^\xf3?٦Q\xbb\xc1\xb6\x93DL\xbf\xd80o\b\x0fuQ:\x15\xacQ\x95L<F\x81\x86\xc8u?\xfe}\xfd\n-\x92\x14\xa9\x14\x94\xa3\xe8\xc4/m|țBmЦu\x1b\xab\xeb\xa8\x13\x157Z(\x1f\xff\x94R\xa0\xf2\xe0BQ\vOi\xf0\xef\x80\xceS\xe8\xc6j\x97\x91\x9a@\x81\x10\f\xd5\x03>\x16xV\xb0d5\xca%s\xf8\aǊ\xa2\xe2f\x14\x84\xab\xa2\xd5'\\c\xe1\xe4\xde\xdeD˘N\x84\xb6_A\xd6\x06K\x8a*9\x96\x96\x89\x8dh:\t\x95\x016\x90\x1dz(\x7f\xf4\xe9\x97\xed&c\xa1K\xe9F\xbf\xcf9E-Z\xd5+\xe4M\xafsM\x93\x92\xc3&\xd5\xffM\xfa\xa3E\xa3\x9d\xf0\xda\x1e\x8e]r\x9c\n'\xa3B\xbf\x92\xa9\x12\xe5=\xe6-\xe3J\x10\x8a\x93ϱKe*BIk\x04\xaaU\xa5\xe9p\rB\x01Ϟd(\xb7\x1d\xfa\xbc\xa1*\xdbՄ\x82#\xa7\x84>w\x1c\x9b[h-\x91\x8d\xbdHY\xf8\x85\xda\xc2R\xab\x8d\xa8\xa6\x86\xf7\xe9\xef\xa9\x14\xb9\xe0\xd3L\xc2\xf6\xb6$+(;\t\xc9,v\xa8Y\x9b\xbaT\xda7\xa2\n\xf6T\xfc7\x02%\x9fԟ\x93'\xa958\xeerO\x8c;\xe8\xed\xe9j\xbaZ\xaf\xf5z\x1d+\x94\x8b|\xb7\x97\x9aS\x90\x00ϛ\x9eF\xe1\xe0\xdd;\xd0\x16ޥ;ѻ\x0fiu\x10\xd2\xcfĠ\xff\uf154\xed.7e71\x9c\xaf\xeb\v\x96\xbfD!\xc2\xf3u}+\xb7\x9a\xa2A\x15\xea\xe9\x863`\xc1\xeb̰\x14*\xfc\x9c\x19\xdf\v\xc5\xf5\xde\xddbl\xc7o\x88b\xea\xe0\xef\t\xf8ב\x8eQ\xdc=\x11\xe2\x18k\xafa\xcfD\x8fct\xbb\xbb\x0f\x19\xbd\x05n\xa8!Y\xf4\xc1**\ah-Uh\x17U\xea0\xe1<g-u\x8a\x19\xb7\xd5\xfe\xf9邍\xebN\xb0\xad\xbb\xcfOm\x88\xdfb\xd6uŷ\x91\x84L\x94\b~\xcb\"yl뷡\x8d\\\xa2\xbbq\xdf\x13\x96\xf5PEk\x8c\xb6\xa2\x12\xe4|\xd5\xcd\x1csvG\xd7\xf6(J&\"\x87`N`\a*\xc7D^\n\x04.6\x1b\xb4\xc4P\"}I\x1b\xafޖ\xef]o\x13\xb1\xe9\xff\xa1\xca_3c\x90\xd3]\x89\x82\xdb\xf8\xea&/yf+\xf4o\x11\xf4\x05\x17\xbd\xf6D[W\x10ա\x8bmåc\xb2F1X\xbd-3̗~\xab\xb7)\xc2Ӽ\x00\x9aKЉ NPN\xa2\xd5\xe0\xe9tdU\x9cm+\x00fw\xc5Ϋ\xb7\x1c\xcb\xe8\xdc\x01~\xcb<I4\x97V(\x0eY\x9d\xd0\x1e\x91&\x9c\xf7\xe1-\xaf\x02\xbc<\x8bx9\x86|\x02oq\xf8Ր\x89\xc4\b\x8b<W\xc2OGn\x06f\x97\x1d,\xafo\xd5\xf9\x9dgy>:\x92\x19\x97\xfe\xd1\xf4\xb1^\x8e'\x86ue4\xdb?\x92W\x11\xf7\xf8\xacp-uO\x8f\x85M\xd8\xcb`c\xd1i\x9e\x10\xe96|\x17yge\x89\xc6#\xff|\xa0\xae~E\xe3'\x00\xea\xfc\xa3ʿ̱\xed\xa3a\xb72\xec\x16R\xf7\xf0sO\x03\xf84V\x12o\xff\x96\xf7\xda\xf2\x14n\xa2f\xa7A\x03\xbcҽ)\xde^ߧNL\xcbb\x7f'\x86:\xd9t\xa2\xa1}L\xa4\xeb\xe9\x8c\xd6O$T\x90\x92\x15\x12\x17\xe0m8E\xd5\xf37\x93\xf4\x8e\xda\x7f2\xbb\xeb\x9a2U3\xf5\x1d\xeb\x1e\x89\xe2c^\xfb\x82\x9bs\xd9Q_簤\x0e9\xe0\x0e\x15\xd0\xe5\x93\t\x89\xbcՙ\xe1\xeb\x97<\x9f\x01=\xa5\x82\xbf\xa7\xf3kt\x8eU\x97\x0eЗ$\x95\xdeU\x9a%\xc0\n\xe2\x8dc\xd6\xfc\xde5g\xfbf\xfe\xfe\xdb\x1c\xe2+\xd9\xfb\x19,\xf1\xaey\x01̊dr5\xad\x83v\xba\xa8\xc1\x99\xcb\xc3\v\xee3\xa3\xed\xf9\xccL\xad\x9aC\x9f\x99\x9a|\x92\xe9O\xa6K}\xae1\xb6sY\x9d\xdd7\x8f\xcc\xdc?\xe2a\xb8\xc9\xd3\r\xbe{\x8e{\xf74\xb0ղ=\xe1\xf1[\x85\nu\x81\x96\xc2\x10\xbf\x86\xb4\xf1\xe8x?S\xbc\x1f\xb5\x1c\xf9\xeb4tw\x81\xa8j\x0e\xaf[\xa2&\xe9=\xa3\xbd\x1dq\xe1\x8cd\x87Θ>C\xcd(?\x9e\x9a\xc9s\xf5\xad$\xb5\xfbv\x94g^\xb9\x0f@\xc3\xdf\xf4S\xceh\xbe\xfb&\xf4\xfb\xecp\xe65c\xf8\x8d\uebab\xd4@åV\xd0|3\xbc\xbd\x82\x0f\xb7\xf9#\x8bw\xd6{\x93\xc1\x88\x9c\xf7t7\xaf\x8f\xfd\x91PtO\xf2\v\xf8\xef\xff\x1f~\t\x00\x00\xff\xff\xd9ߐm~\x1f\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcZIs\xe3\xb8\x15\xbe\xfbW\xbc\xea\x1c\xe6b\xc9\xd3\xc9T*\xa5[[N\xaa\\\x99v\xbbZ\x8e\xef\x10\xf9Db\f\x02\f\x16i\x9c\xe5\xbfO=\x00\xa4@\x12Zg\xbay\xe8jayx\x1b\xbe\xb7\xc0\xb3\xd9솵\xfc\x15\xb5\xe1J.\x80\xb5\x1c\x7f\xb5(闙\xbf\xfd\xcd̹\xba\xdb~\xbcy\xe3\xb2\\\xc0\xd2\x19\xab\x9a\xafh\x94\xd3\x05>\xe0\x86Kn\xb9\x927\rZV2\xcb\x167\x00LJe\x19\r\x1b\xfa\tP(i\xb5\x12\x02\xf5\xacB9\x7fsk\\;.JԞxw\xf4\xf6\xc7\xf9ǟ\xe6?\xde\x00H\xd6\xe0\x02\x88\x9ek\x85b\xa5\x99oQ\xa0Vs\xaenL\x8b\x05\x91\xad\xb4r\xed\x02\xf6\x13a[<2\xb0\xfb\xc0,\xfb\x97\xa7\xe0\a\x057\xf6\x9f\xa3\x89\x9f\xb9\xb1~\xb2\x15N318Տ\x1b.+'\x98Ngn\x00L\xa1Z\\\xc0\x13\x1dٲ\x02i,J\xe2Y\x98\x01+K\xaf\x1b&\x9e5\x97\x16\xf5R\t\xd7t:\x99A\x89\xa6м\xb5^\xf6=C`,\xb3\u0380qE\r\xcc\xc0\x13\xee\xee\x1e\xe5\xb3V\x95F\x13X\x02\xf8\xc5(\xf9\xccl\xbd\x80yX>okf0\xce\x06\xf5\xad\xfcD\x1c\xb2\xefĭ\xb1\x9a\xcb*w\xfe\vo\x10J\xa7\xbd\xd9H\xe6\x02\xc1\xd6ܤ\x8c\xed\x98!\xe6\xb4\xc5\xf2 \x1b~\x9e\x88\x19˚v\xccO\xb250T2\x8b9v\x96\xaai\x05Z,a\xfdn\xb1\x13b\xa3t\xc3\xec\x02\xb8\xb4\x7f\xfd\xe9\xb0&\xa2\xaa\xe6~냒C\xb5\xdc\xd3($Á\x13\xb2P\x85:\xab\x1be\x99\xf8=\x8cX\"p\x9f\xec\x0f\x9c\x04\xba\xe9\xf8IV\xc8\xdd@m\xc0\xd6\b\xf7\xacxs-\xac\xacҬB\xf8Y\x15\xc1x\xbb\x1au4\xde:,1\xb5r\xa2\x84u'1\x80\xb1Jg\xad\xd8b1\x0f\xbb\"ݎ\xecȔ\xc33\xff`'+4\xb2\xac\x93u(3\xf7+\xb8\x92yO\xfbT\xe1Y^\x96jS\xaa\x12{\xd5a\xca\x117\xd0jU\xa01G\xfc\x9e\xb6\x0fxx\xda\x0fL\xd4\x12Vl\xff\xccD[\xb3\x8f\x01e\x8a\x1a\x1b\xb6\x88;T\x8b\xf2\xd3\xf3\xe3\xeb_V\x83a8\x88\x19\xac\xb0\x86\xc0\x82Xo\xb5\xb2\xaaP\x02\xd6hw\x88\xd2\xe3\x164j\x8b\x9a@\xae\xe2\xd2\x00\x93eO\x13\xd2\x05{\xa8&'\xf7\xf4h6LFwR-\xea\xd4\xec@G\xb6\xa8-\xef\xd07|IXIFGB\xfco6\x98\x03 \xb9\xc3.()\xbe`\x90*b+\x96QU\xc1n܀\xc6V\xa3A\x19\"\x0e\r3\tj\xfd\v\x16v>\"\xbdBMd\xba\xfbP(\xb9EmAc\xa1*\xc9\xff\xd3\xd36`\x95?T0\x8b\xc6\xfa\v\xa9%\x13\xb0e\xc2\xe1\xedH{\xf45\xec\x1d4ҙ\xe0dB\xcfo0c>>+\x8d\xc0\xe5F-\xa0\xb6\xb65\x8b\xbb\xbb\x8a\xdb.\xd8\x16\xaai\x9c\xe4\xf6\xfd\xce\x1b\x83\xaf\x9dU\xdaܕ\xb8Eqgx5c\xba\xa8\xb9\xc5\xc2:\x8dw\xac\xe53/\x88\xf4\x01wޔ\x7f\xd21<\x9b\xc1\xb1\x13/\f\x9f\x0f\x94\x17\x98\x87\xe2']\t\x16I\x05\x11\xf7V\xa0!R\xdd\u05ff\xaf^\xa0\xe3$X*\x18e\xbft\xa2\x97\xce>\xa4M.7\xa8þ\x8dV\x8d\xa7\x89\xb2l\x15\x97\xd6\xff(\x04Gi\xc1\xb8u\xc3-\xb9\xc1\xbf\x1d\x1aK\xa6\x1b\x93]\xfa\x84\x04\xd6\b\xae%((\xc7\v\x1e%,Y\x83b\xc9\f~g[\x91Ǔ\x8cp\x96\xb5\xd24k\xbc8\xa87\x99\xe82\xa5\x03\xa6\xdd\xc3ǪłlJj\xa5M|\xc3c,!\f`\xc9ʡv\xf2מ\xbel\b\x19/:\xe5j\xf4\xdd\xe7\bu\xbc\xca\x04\xbf\xbbP\x17#\x93\x18F\xa6\xf4ۃ|ܣ\xb1U\x86[\xa5߉p\b\x8dc78h\x11\xfa\n&\v\x14\u05c8\xb7\xf4;\x81˒4\x8e\xbd\x1b\x13\x00\x05\xaa\x9eQ%+E\x17+1\x04<ZZA^m\xd0\xe6Ŕ\x99P\xc6%\xec\xb3IH\xb3Ʊ\xa8k\xa5\x04\xb2\xb1\x06\v\xc3W\x92\xb5\xa6V\xf6\x84\xc0\x8f\x1b\xe8V\xbe\xbc\xb7H\x87/W\x8f\xb7\xf4O7N\x1e\xb4\xe5e\x84x\xbae\x94W\xe5\xcd\x16\xed\xbc\\=\x82\x89ۧF\x92N\b\xb6\x16\xb8\x00\xab\xddT\xb0\xc3\x0eK_Gv)\x98\xc9.\x18\t\xb8J\xd7\xe7|\xb2#\b\x85_ak\x963\x94\xd78E8*\x0f\x92M\xbcO\x84`\xc7m\x9d\xddy\xc4)!\xa6y\xac³\x05J\x96g剗+\x88\xa36G\x84y~]zyOIF\xd8~\x8dd\x81\xe4aO\x9c\xc8\xf6:ؐ\x93n\xc4\xe5!\xe1\x14]9B\x0e,\xc1\xb5\x97\xf3N7\x9ck,\xa7<\xcf\x06\xf6\xcaL\x0f\x85>pm'a\x00b\x86\xf7\x99r\xb8\xa5\x92\x1b^M\xcfN\x8b\xd5cw\xe4\xa8h\x93\xf0\x92\x1cI\x1a\xa7hB\x9c\xcc|:9\xebB\r%b\x1b^9}\xe8\xeao8\x8ar\x92-\x9c\xbc\xed'\xf4ᙸ\x06\xb4{ɺ`\x19\xf1+I\xa3\x83\x978\xe3\v\xd8$\xd6Le\x00\xc2\xc9=En\xe0\xc3\aP\x1a>\x84\xc6Ƈ۰\xdbqag|\x90\xcb\xef\xb8\x10\xdd)\x17\x85\xab>\x7f\xa7\xeaI\xb9S8\x9e\xd5\xc1\x97\x11\x8d\x91*,Uz^|\xab`\xc7x\x92C\xf7\xa7\x9b\xdb\f\xdd5n(\xe1\xd2h\x9d\x96\x14\xf2Pk\xcaA\x8c'\xa9\\\x06\xf3\x8fHj\x92\xf8sB\xcaq\xa8\xf2R\xd0\xff\xc7X\x9e\x02@F\x80\x9c\x8d\x8fq\xe8\xf3㾋t\x8d)VC\x12\x1d\xf3J\xf3\x8a\x93\xc2e?\xb3\xcf|\"\xd6\xc5\x16\x81G2\x0f\xc5Y\xff\xec\xd1\xd2\x10Z\xee\xc9\xd1u\x0e\x87\x13\xda3Y\xfa\xe0\xdcϗ\xf1\xeae.\xeeI\x85<\xbf.O٫?8\x03\xe54\xbc\xabyQ\x0fMǧ\xa0\n`\xd9\x1b\xfaD\xf7\x026\xf3\x18>˧\xbd\xa35\xe3\xdb7\x9aN]v<54tv\xf6\xf9uyVi\xe0\xbb\x16\xe7\x15\a\xa1\x1d\x19\xb5\\8\xad}\xd9\x15F\xa9ھ\xa2<`E\x81\xad\xc5\xf2\xfe\xfdI\x95\xa7\x9c\xfe\xd3`11\"\xcf\xe9\xdbdL\xed;9زK\xf3\xfb\x8eݾ\xdbt\xcd5\xfd4&\xe2\xfb\x0e\xbaL\x00s\x9a\xad\a\xb09\xcc4\xc0\v9\xb8\xaf\x9b\x7f\b\x18I\xdb<\xf2\xd2\xf5\x9c\x1c:\xa1е2\xa90\x9e\xd1\xfe\xeb\xa2l\xbe.\n]ܴOwU\x914%3\xd5\x1d\xeb\xaa9\xdf@\xec\xda\xc79\x8d\xed\xc9\xf5\xfa\n\u0530\x04ܢ\x04\xaa{\x19\x17\x14\xbb=\xc9\f\x80\x1d\xa7\x12\x83Xx+\xe8\x1a\"]\xf3,ۙ:mɌ\x12\xa6h\xf6-\x8d٧\x90_\xd18\x91I\x1a\xbea\n\x19\x8e\f\xa5\xb9ɦ\x90\xc7kGf\x80\x81\x0eD\"n\x1c\x02\xad\xb3\x95\x94\xcd+\x1b4\x86U\xa7\x10\xedsX\x15\xdahq\v\xb05\xa5QC\xd6~0\x11h/\x82+y\x1aS/B\xd2A\a\xfcbN\xbe\xac\xce\xe0\xe5ˊ\x0e\xf9\xb2\xfa\xbd\xbc\xa0tM\xae\xb0bΪ̰\xe0\xd2\xfd\x9a\x19\xdfqY\xaa\xdd\xf4~\x1d\x11\xb5e\xb6>!\xe83\xb3u\x17G7N\b\xbfg\x92_\xc6\xd4l\x8d\x04\x1c\x7fT\x9a\xe9\xfb<\xa7أ5\xb98\x8f\xe7ܙC\x9a\x7f\xc2]f\xb4\x8bK\x99\xa9\xe7\x18\xec2S\x93g\xd0t2\xb4\xd2r\x98\xd2\xcdei\xf6/\x8d\x99\xb9\x7f\xf8(p\x91\x9e#\x7fׄ\xb9\xbe)W+\xd1E6\xffB(]\xb3FMF\xf0o\x90\xa3~\x05\xe5݉\xc52\x84\x93\xfd}\xb2\xef)\xcd\xe1\xa5\xe6\xa6k#v\xe5Z\xc9M+\xd8{/\xcb)l\xedqk\xfc<4u\x92\xe3\xfd\xb7\xfe\xbd6\xdf\xce\xc9=\xba\x0e\xbf\xe9\xf3\xe9h\xbe\x7f\x87\xfd6'\x1c\t\f\xdd\xf5~|8\xb3\x0e}|\xe8\xae\"/QZ*\xad\xf7Or\xfb\x8aƷxs\xba\x1c\xb7\xb6/+\xc2\x06\xaf\xf8W\x15\xa5\x03\n'ҵ\xf8G\x05\xb9\xa4hE`@\x10\xe4\x1f\x81\x96\xe3g\xdf\xdb\xfe\x15\x99\xd9\xf8\x12U\xd4LV\x98\xab\xf4\x94\xa4\x1c\xc0\xe7\x10\x97\xe7_C\x81\xbeg\xea\x95\xf5\xaaɠ\xe7\xbcLh\xc7^b:\xe2\xd6\xfd\xd3\xe0\x02\xfe\xfb\xff\x9b\xdf\x02\x00\x00\xff\xff\xcc\b\u008b\xfc#\x00\x00"),
}

var CRDs = crds()

func crds() []*apiextv1.CustomResourceDefinition {
	apiextinstall.Install(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	var objs []*apiextv1.CustomResourceDefinition
	for _, crd := range rawCRDs {
		gzr, err := gzip.NewReader(bytes.NewReader(crd))
		if err != nil {
			panic(err)
		}
		bytes, err := io.ReadAll(gzr)
		if err != nil {
			panic(err)
		}
		gzr.Close()

		obj, _, err := decode(bytes, nil, nil)
		if err != nil {
			panic(err)
		}
		objs = append(objs, obj.(*apiextv1.CustomResourceDefinition))
	}
	return objs
}
