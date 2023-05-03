// Copyright 2023 Dolthub, Inc.
//
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

package encodings

// Latin1_danish_ci_RuneWeight returns the weight of a given rune based on its relational sort order from
// the `latin1_danish_ci` collation.
func Latin1_danish_ci_RuneWeight(r rune) int32 {
	weight, ok := latin1_danish_ci_Weights[r]
	if ok {
		return weight
	} else {
		return 2147483647
	}
}

// latin1_danish_ci_Weights contain a map from rune to weight for the `latin1_danish_ci` collation. The
// map primarily contains mappings that have a random order. Mappings that fit into a sequential range (and are long
// enough) are defined in the calling function to save space.
var latin1_danish_ci_Weights = map[rune]int32{
	0:    0,
	1:    1,
	2:    2,
	3:    3,
	4:    4,
	5:    5,
	6:    6,
	7:    7,
	8:    8,
	9:    9,
	10:   10,
	11:   11,
	12:   12,
	13:   13,
	14:   14,
	15:   15,
	16:   16,
	17:   17,
	18:   18,
	19:   19,
	20:   20,
	21:   21,
	22:   22,
	23:   23,
	24:   24,
	25:   25,
	26:   26,
	27:   27,
	28:   28,
	29:   29,
	30:   30,
	31:   31,
	32:   32,
	33:   33,
	34:   34,
	35:   35,
	36:   36,
	37:   37,
	38:   38,
	39:   39,
	40:   40,
	41:   41,
	42:   42,
	43:   43,
	44:   44,
	45:   45,
	46:   46,
	47:   47,
	48:   48,
	49:   49,
	50:   50,
	51:   51,
	52:   52,
	53:   53,
	54:   54,
	55:   55,
	56:   56,
	57:   57,
	58:   58,
	59:   59,
	60:   60,
	61:   61,
	62:   62,
	63:   63,
	64:   64,
	65:   65,
	97:   65,
	192:  65,
	193:  65,
	194:  65,
	195:  65,
	224:  65,
	225:  65,
	226:  65,
	227:  65,
	66:   66,
	98:   66,
	67:   67,
	99:   67,
	199:  67,
	231:  67,
	68:   68,
	100:  68,
	208:  68,
	240:  68,
	69:   69,
	101:  69,
	200:  69,
	201:  69,
	202:  69,
	203:  69,
	232:  69,
	233:  69,
	234:  69,
	235:  69,
	70:   70,
	102:  70,
	71:   71,
	103:  71,
	72:   72,
	104:  72,
	73:   73,
	105:  73,
	204:  73,
	205:  73,
	206:  73,
	207:  73,
	236:  73,
	237:  73,
	238:  73,
	239:  73,
	74:   74,
	106:  74,
	75:   75,
	107:  75,
	76:   76,
	108:  76,
	77:   77,
	109:  77,
	78:   78,
	110:  78,
	209:  78,
	241:  78,
	79:   79,
	111:  79,
	210:  79,
	211:  79,
	212:  79,
	213:  79,
	242:  79,
	243:  79,
	244:  79,
	245:  79,
	80:   80,
	112:  80,
	81:   81,
	113:  81,
	82:   82,
	114:  82,
	83:   83,
	115:  83,
	84:   84,
	116:  84,
	85:   85,
	117:  85,
	217:  85,
	218:  85,
	219:  85,
	249:  85,
	250:  85,
	251:  85,
	86:   86,
	118:  86,
	87:   87,
	119:  87,
	88:   88,
	120:  88,
	89:   89,
	121:  89,
	220:  89,
	221:  89,
	252:  89,
	253:  89,
	90:   90,
	122:  90,
	91:   91,
	196:  91,
	198:  91,
	228:  91,
	230:  91,
	92:   92,
	214:  92,
	216:  92,
	246:  92,
	248:  92,
	93:   93,
	197:  93,
	229:  93,
	94:   94,
	95:   95,
	96:   96,
	123:  97,
	124:  98,
	125:  99,
	126:  100,
	127:  101,
	8364: 102,
	129:  103,
	8218: 104,
	402:  105,
	8222: 106,
	8230: 107,
	8224: 108,
	8225: 109,
	710:  110,
	8240: 111,
	352:  112,
	8249: 113,
	338:  114,
	141:  115,
	381:  116,
	143:  117,
	144:  118,
	8216: 119,
	8217: 120,
	8220: 121,
	8221: 122,
	8226: 123,
	8211: 124,
	8212: 125,
	732:  126,
	8482: 127,
	353:  128,
	8250: 129,
	339:  130,
	157:  131,
	382:  132,
	376:  133,
	160:  134,
	161:  135,
	162:  136,
	163:  137,
	164:  138,
	165:  139,
	166:  140,
	167:  141,
	168:  142,
	169:  143,
	170:  144,
	171:  145,
	172:  146,
	173:  147,
	174:  148,
	175:  149,
	176:  150,
	177:  151,
	178:  152,
	179:  153,
	180:  154,
	181:  155,
	182:  156,
	183:  157,
	184:  158,
	185:  159,
	186:  160,
	187:  161,
	188:  162,
	189:  163,
	190:  164,
	191:  165,
	215:  166,
	222:  167,
	254:  167,
	223:  168,
	247:  169,
	255:  170,
}
