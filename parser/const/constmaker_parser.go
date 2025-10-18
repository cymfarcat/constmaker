// Code generated from ConstMaker.g4 by ANTLR 4.13.2. DO NOT EDIT.

package consts // ConstMaker
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ConstMakerParser struct {
	*antlr.BaseParser
}

var ConstMakerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func constmakerParserInit() {
	staticData := &ConstMakerParserStaticData
	staticData.LiteralNames = []string{
		"", "':'", "'='", "'{'", "'}'", "'['", "']'", "','", "'enum'", "'namespace'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "ENUM", "NAMESPACE", "STRING", "TYPE",
		"DEC_INTEGER", "HEX_INTEGER", "FLOAT", "IDENT", "DOC_COMMENT", "TRIPLE_COMMENT",
		"BLOCK_COMMENT", "LINE_COMMENT", "NL", "WS",
	}
	staticData.RuleNames = []string{
		"file", "rules", "const", "enum", "namespace", "type", "enumValue",
		"identAssignValue", "identAssignValueLoop", "option", "value", "integer",
		"ident", "keywords", "docComment", "tripleComment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 262, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 5, 0, 35, 8, 0, 10, 0, 12, 0, 38, 9, 0, 5, 0, 40, 8, 0, 10,
		0, 12, 0, 43, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 1, 2,
		5, 2, 53, 8, 2, 10, 2, 12, 2, 56, 9, 2, 1, 2, 3, 2, 59, 8, 2, 1, 2, 5,
		2, 62, 8, 2, 10, 2, 12, 2, 65, 9, 2, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2,
		1, 2, 1, 2, 3, 2, 74, 8, 2, 1, 2, 3, 2, 77, 8, 2, 1, 2, 3, 2, 80, 8, 2,
		1, 2, 5, 2, 83, 8, 2, 10, 2, 12, 2, 86, 9, 2, 1, 3, 5, 3, 89, 8, 3, 10,
		3, 12, 3, 92, 9, 3, 1, 3, 3, 3, 95, 8, 3, 1, 3, 5, 3, 98, 8, 3, 10, 3,
		12, 3, 101, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 107, 8, 3, 1, 3, 3, 3,
		110, 8, 3, 1, 3, 5, 3, 113, 8, 3, 10, 3, 12, 3, 116, 9, 3, 1, 3, 1, 3,
		5, 3, 120, 8, 3, 10, 3, 12, 3, 123, 9, 3, 1, 3, 5, 3, 126, 8, 3, 10, 3,
		12, 3, 129, 9, 3, 1, 3, 5, 3, 132, 8, 3, 10, 3, 12, 3, 135, 9, 3, 1, 3,
		1, 3, 5, 3, 139, 8, 3, 10, 3, 12, 3, 142, 9, 3, 1, 4, 5, 4, 145, 8, 4,
		10, 4, 12, 4, 148, 9, 4, 1, 4, 3, 4, 151, 8, 4, 1, 4, 5, 4, 154, 8, 4,
		10, 4, 12, 4, 157, 9, 4, 1, 4, 1, 4, 1, 4, 3, 4, 162, 8, 4, 1, 4, 5, 4,
		165, 8, 4, 10, 4, 12, 4, 168, 9, 4, 1, 4, 1, 4, 5, 4, 172, 8, 4, 10, 4,
		12, 4, 175, 9, 4, 1, 4, 1, 4, 1, 4, 5, 4, 180, 8, 4, 10, 4, 12, 4, 183,
		9, 4, 1, 4, 5, 4, 186, 8, 4, 10, 4, 12, 4, 189, 9, 4, 1, 4, 1, 4, 5, 4,
		193, 8, 4, 10, 4, 12, 4, 196, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 202,
		8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 208, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6,
		213, 8, 6, 1, 6, 3, 6, 216, 8, 6, 1, 6, 3, 6, 219, 8, 6, 1, 6, 5, 6, 222,
		8, 6, 10, 6, 12, 6, 225, 9, 6, 1, 7, 1, 7, 1, 7, 3, 7, 230, 8, 7, 1, 8,
		1, 8, 1, 8, 5, 8, 235, 8, 8, 10, 8, 12, 8, 238, 9, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 248, 8, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 3, 12, 254, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 0, 0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		0, 2, 1, 0, 12, 13, 1, 0, 8, 9, 291, 0, 41, 1, 0, 0, 0, 2, 49, 1, 0, 0,
		0, 4, 54, 1, 0, 0, 0, 6, 90, 1, 0, 0, 0, 8, 146, 1, 0, 0, 0, 10, 207, 1,
		0, 0, 0, 12, 209, 1, 0, 0, 0, 14, 226, 1, 0, 0, 0, 16, 231, 1, 0, 0, 0,
		18, 239, 1, 0, 0, 0, 20, 247, 1, 0, 0, 0, 22, 249, 1, 0, 0, 0, 24, 253,
		1, 0, 0, 0, 26, 255, 1, 0, 0, 0, 28, 257, 1, 0, 0, 0, 30, 259, 1, 0, 0,
		0, 32, 36, 3, 2, 1, 0, 33, 35, 5, 20, 0, 0, 34, 33, 1, 0, 0, 0, 35, 38,
		1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0,
		38, 36, 1, 0, 0, 0, 39, 32, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1,
		0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44,
		45, 5, 0, 0, 1, 45, 1, 1, 0, 0, 0, 46, 50, 3, 4, 2, 0, 47, 50, 3, 6, 3,
		0, 48, 50, 3, 8, 4, 0, 49, 46, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 48,
		1, 0, 0, 0, 50, 3, 1, 0, 0, 0, 51, 53, 5, 20, 0, 0, 52, 51, 1, 0, 0, 0,
		53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 58, 1,
		0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 59, 3, 28, 14, 0, 58, 57, 1, 0, 0, 0,
		58, 59, 1, 0, 0, 0, 59, 63, 1, 0, 0, 0, 60, 62, 5, 20, 0, 0, 61, 60, 1,
		0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64,
		66, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 69, 3, 24, 12, 0, 67, 68, 5, 1,
		0, 0, 68, 70, 3, 10, 5, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70,
		73, 1, 0, 0, 0, 71, 72, 5, 2, 0, 0, 72, 74, 3, 20, 10, 0, 73, 71, 1, 0,
		0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 77, 3, 18, 9, 0, 76,
		75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 80, 3, 30,
		15, 0, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 84, 1, 0, 0, 0, 81,
		83, 5, 20, 0, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0,
		0, 0, 84, 85, 1, 0, 0, 0, 85, 5, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 89,
		5, 20, 0, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0,
		90, 91, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 95, 3,
		28, 14, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 99, 1, 0, 0, 0,
		96, 98, 5, 20, 0, 0, 97, 96, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1,
		0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 102, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0,
		102, 103, 5, 8, 0, 0, 103, 106, 3, 24, 12, 0, 104, 105, 5, 1, 0, 0, 105,
		107, 3, 10, 5, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109,
		1, 0, 0, 0, 108, 110, 3, 18, 9, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 114, 1, 0, 0, 0, 111, 113, 5, 20, 0, 0, 112, 111, 1, 0, 0, 0,
		113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		117, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 121, 5, 3, 0, 0, 118, 120,
		5, 20, 0, 0, 119, 118, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0,
		0, 0, 121, 122, 1, 0, 0, 0, 122, 127, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0,
		124, 126, 3, 12, 6, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 133, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 130, 132, 5, 20, 0, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0,
		0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0,
		135, 133, 1, 0, 0, 0, 136, 140, 5, 4, 0, 0, 137, 139, 5, 20, 0, 0, 138,
		137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 7, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 145, 5, 20,
		0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0,
		146, 147, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149,
		151, 3, 28, 14, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 155,
		1, 0, 0, 0, 152, 154, 5, 20, 0, 0, 153, 152, 1, 0, 0, 0, 154, 157, 1, 0,
		0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0, 0,
		157, 155, 1, 0, 0, 0, 158, 159, 5, 9, 0, 0, 159, 161, 3, 24, 12, 0, 160,
		162, 3, 18, 9, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 166,
		1, 0, 0, 0, 163, 165, 5, 20, 0, 0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0,
		0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0,
		168, 166, 1, 0, 0, 0, 169, 173, 5, 3, 0, 0, 170, 172, 5, 20, 0, 0, 171,
		170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174,
		1, 0, 0, 0, 174, 181, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 180, 3, 8,
		4, 0, 177, 180, 3, 4, 2, 0, 178, 180, 3, 6, 3, 0, 179, 176, 1, 0, 0, 0,
		179, 177, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 187, 1, 0, 0, 0, 183, 181,
		1, 0, 0, 0, 184, 186, 5, 20, 0, 0, 185, 184, 1, 0, 0, 0, 186, 189, 1, 0,
		0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 190, 1, 0, 0, 0,
		189, 187, 1, 0, 0, 0, 190, 194, 5, 4, 0, 0, 191, 193, 5, 20, 0, 0, 192,
		191, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 9, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 198, 5, 5, 0,
		0, 198, 201, 3, 10, 5, 0, 199, 200, 5, 1, 0, 0, 200, 202, 3, 22, 11, 0,
		201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		204, 5, 6, 0, 0, 204, 208, 1, 0, 0, 0, 205, 208, 5, 11, 0, 0, 206, 208,
		3, 24, 12, 0, 207, 197, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 206, 1,
		0, 0, 0, 208, 11, 1, 0, 0, 0, 209, 212, 3, 24, 12, 0, 210, 211, 5, 2, 0,
		0, 211, 213, 3, 22, 11, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0,
		213, 215, 1, 0, 0, 0, 214, 216, 3, 18, 9, 0, 215, 214, 1, 0, 0, 0, 215,
		216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217, 219, 3, 30, 15, 0, 218, 217,
		1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 223, 1, 0, 0, 0, 220, 222, 5, 20,
		0, 0, 221, 220, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0,
		223, 224, 1, 0, 0, 0, 224, 13, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 229,
		3, 24, 12, 0, 227, 228, 5, 1, 0, 0, 228, 230, 3, 20, 10, 0, 229, 227, 1,
		0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 15, 1, 0, 0, 0, 231, 236, 3, 14, 7,
		0, 232, 233, 5, 7, 0, 0, 233, 235, 3, 14, 7, 0, 234, 232, 1, 0, 0, 0, 235,
		238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 17, 1,
		0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 240, 5, 5, 0, 0, 240, 241, 3, 16, 8,
		0, 241, 242, 5, 6, 0, 0, 242, 19, 1, 0, 0, 0, 243, 248, 3, 22, 11, 0, 244,
		248, 5, 14, 0, 0, 245, 248, 5, 10, 0, 0, 246, 248, 3, 24, 12, 0, 247, 243,
		1, 0, 0, 0, 247, 244, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 246, 1, 0,
		0, 0, 248, 21, 1, 0, 0, 0, 249, 250, 7, 0, 0, 0, 250, 23, 1, 0, 0, 0, 251,
		254, 5, 15, 0, 0, 252, 254, 3, 26, 13, 0, 253, 251, 1, 0, 0, 0, 253, 252,
		1, 0, 0, 0, 254, 25, 1, 0, 0, 0, 255, 256, 7, 1, 0, 0, 256, 27, 1, 0, 0,
		0, 257, 258, 5, 16, 0, 0, 258, 29, 1, 0, 0, 0, 259, 260, 5, 17, 0, 0, 260,
		31, 1, 0, 0, 0, 41, 36, 41, 49, 54, 58, 63, 69, 73, 76, 79, 84, 90, 94,
		99, 106, 109, 114, 121, 127, 133, 140, 146, 150, 155, 161, 166, 173, 179,
		181, 187, 194, 201, 207, 212, 215, 218, 223, 229, 236, 247, 253,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ConstMakerParserInit initializes any static state used to implement ConstMakerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConstMakerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConstMakerParserInit() {
	staticData := &ConstMakerParserStaticData
	staticData.once.Do(constmakerParserInit)
}

// NewConstMakerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConstMakerParser(input antlr.TokenStream) *ConstMakerParser {
	ConstMakerParserInit()
	this := new(ConstMakerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ConstMakerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ConstMaker.g4"

	return this
}

// ConstMakerParser tokens.
const (
	ConstMakerParserEOF            = antlr.TokenEOF
	ConstMakerParserT__0           = 1
	ConstMakerParserT__1           = 2
	ConstMakerParserT__2           = 3
	ConstMakerParserT__3           = 4
	ConstMakerParserT__4           = 5
	ConstMakerParserT__5           = 6
	ConstMakerParserT__6           = 7
	ConstMakerParserENUM           = 8
	ConstMakerParserNAMESPACE      = 9
	ConstMakerParserSTRING         = 10
	ConstMakerParserTYPE           = 11
	ConstMakerParserDEC_INTEGER    = 12
	ConstMakerParserHEX_INTEGER    = 13
	ConstMakerParserFLOAT          = 14
	ConstMakerParserIDENT          = 15
	ConstMakerParserDOC_COMMENT    = 16
	ConstMakerParserTRIPLE_COMMENT = 17
	ConstMakerParserBLOCK_COMMENT  = 18
	ConstMakerParserLINE_COMMENT   = 19
	ConstMakerParserNL             = 20
	ConstMakerParserWS             = 21
)

// ConstMakerParser rules.
const (
	ConstMakerParserRULE_file                 = 0
	ConstMakerParserRULE_rules                = 1
	ConstMakerParserRULE_const                = 2
	ConstMakerParserRULE_enum                 = 3
	ConstMakerParserRULE_namespace            = 4
	ConstMakerParserRULE_type                 = 5
	ConstMakerParserRULE_enumValue            = 6
	ConstMakerParserRULE_identAssignValue     = 7
	ConstMakerParserRULE_identAssignValueLoop = 8
	ConstMakerParserRULE_option               = 9
	ConstMakerParserRULE_value                = 10
	ConstMakerParserRULE_integer              = 11
	ConstMakerParserRULE_ident                = 12
	ConstMakerParserRULE_keywords             = 13
	ConstMakerParserRULE_docComment           = 14
	ConstMakerParserRULE_tripleComment        = 15
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllRules() []IRulesContext
	Rules(i int) IRulesContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserEOF, 0)
}

func (s *FileContext) AllRules() []IRulesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRulesContext); ok {
			len++
		}
	}

	tst := make([]IRulesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRulesContext); ok {
			tst[i] = t.(IRulesContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Rules(i int) IRulesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRulesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRulesContext)
}

func (s *FileContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConstMakerParserNL)
}

func (s *FileContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConstMakerParserNL, i)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConstMakerParserRULE_file)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1147648) != 0 {
		{
			p.SetState(32)
			p.Rules()
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(33)
					p.Match(ConstMakerParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
		p.Match(ConstMakerParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_() IConstContext
	Enum() IEnumContext
	Namespace() INamespaceContext

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_rules
	return p
}

func InitEmptyRulesContext(p *RulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_rules
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) Const_() IConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstContext)
}

func (s *RulesContext) Enum() IEnumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumContext)
}

func (s *RulesContext) Namespace() INamespaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitRules(s)
	}
}

func (s *RulesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitRules(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Rules() (localctx IRulesContext) {
	localctx = NewRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConstMakerParserRULE_rules)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Const_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Enum()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Namespace()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstContext is an interface to support dynamic dispatch.
type IConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	DocComment() IDocCommentContext
	Type_() ITypeContext
	Value() IValueContext
	Option() IOptionContext
	TripleComment() ITripleCommentContext

	// IsConstContext differentiates from other interfaces.
	IsConstContext()
}

type ConstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstContext() *ConstContext {
	var p = new(ConstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_const
	return p
}

func InitEmptyConstContext(p *ConstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_const
}

func (*ConstContext) IsConstContext() {}

func NewConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstContext {
	var p = new(ConstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_const

	return p
}

func (s *ConstContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ConstContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConstMakerParserNL)
}

func (s *ConstContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConstMakerParserNL, i)
}

func (s *ConstContext) DocComment() IDocCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocCommentContext)
}

func (s *ConstContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ConstContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ConstContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ConstContext) TripleComment() ITripleCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITripleCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITripleCommentContext)
}

func (s *ConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterConst(s)
	}
}

func (s *ConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitConst(s)
	}
}

func (s *ConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitConst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Const_() (localctx IConstContext) {
	localctx = NewConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConstMakerParserRULE_const)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(51)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserDOC_COMMENT {
		{
			p.SetState(57)
			p.DocComment()
		}

	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConstMakerParserNL {
		{
			p.SetState(60)
			p.Match(ConstMakerParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Ident()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__0 {
		{
			p.SetState(67)
			p.Match(ConstMakerParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Type_()
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__1 {
		{
			p.SetState(71)
			p.Match(ConstMakerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Value()
		}

	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__4 {
		{
			p.SetState(75)
			p.Option()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserTRIPLE_COMMENT {
		{
			p.SetState(78)
			p.TripleComment()
		}

	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(81)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumContext is an interface to support dynamic dispatch.
type IEnumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	Ident() IIdentContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	DocComment() IDocCommentContext
	Type_() ITypeContext
	Option() IOptionContext
	AllEnumValue() []IEnumValueContext
	EnumValue(i int) IEnumValueContext

	// IsEnumContext differentiates from other interfaces.
	IsEnumContext()
}

type EnumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumContext() *EnumContext {
	var p = new(EnumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_enum
	return p
}

func InitEmptyEnumContext(p *EnumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_enum
}

func (*EnumContext) IsEnumContext() {}

func NewEnumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumContext {
	var p = new(EnumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_enum

	return p
}

func (s *EnumContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumContext) ENUM() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserENUM, 0)
}

func (s *EnumContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConstMakerParserNL)
}

func (s *EnumContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConstMakerParserNL, i)
}

func (s *EnumContext) DocComment() IDocCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocCommentContext)
}

func (s *EnumContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *EnumContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *EnumContext) AllEnumValue() []IEnumValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumValueContext); ok {
			len++
		}
	}

	tst := make([]IEnumValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumValueContext); ok {
			tst[i] = t.(IEnumValueContext)
			i++
		}
	}

	return tst
}

func (s *EnumContext) EnumValue(i int) IEnumValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterEnum(s)
	}
}

func (s *EnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitEnum(s)
	}
}

func (s *EnumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitEnum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Enum() (localctx IEnumContext) {
	localctx = NewEnumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConstMakerParserRULE_enum)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(87)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserDOC_COMMENT {
		{
			p.SetState(93)
			p.DocComment()
		}

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConstMakerParserNL {
		{
			p.SetState(96)
			p.Match(ConstMakerParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(102)
		p.Match(ConstMakerParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Ident()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__0 {
		{
			p.SetState(104)
			p.Match(ConstMakerParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Type_()
		}

	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__4 {
		{
			p.SetState(108)
			p.Option()
		}

	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConstMakerParserNL {
		{
			p.SetState(111)
			p.Match(ConstMakerParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(ConstMakerParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(118)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33536) != 0 {
		{
			p.SetState(124)
			p.EnumValue()
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConstMakerParserNL {
		{
			p.SetState(130)
			p.Match(ConstMakerParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(ConstMakerParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(137)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAMESPACE() antlr.TerminalNode
	Ident() IIdentContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	DocComment() IDocCommentContext
	Option() IOptionContext
	AllNamespace() []INamespaceContext
	Namespace(i int) INamespaceContext
	AllConst_() []IConstContext
	Const_(i int) IConstContext
	AllEnum() []IEnumContext
	Enum(i int) IEnumContext

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	var p = new(NamespaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_namespace
	return p
}

func InitEmptyNamespaceContext(p *NamespaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_namespace
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserNAMESPACE, 0)
}

func (s *NamespaceContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *NamespaceContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConstMakerParserNL)
}

func (s *NamespaceContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConstMakerParserNL, i)
}

func (s *NamespaceContext) DocComment() IDocCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocCommentContext)
}

func (s *NamespaceContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *NamespaceContext) AllNamespace() []INamespaceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespaceContext); ok {
			len++
		}
	}

	tst := make([]INamespaceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespaceContext); ok {
			tst[i] = t.(INamespaceContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceContext) Namespace(i int) INamespaceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *NamespaceContext) AllConst_() []IConstContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstContext); ok {
			len++
		}
	}

	tst := make([]IConstContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstContext); ok {
			tst[i] = t.(IConstContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceContext) Const_(i int) IConstContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstContext)
}

func (s *NamespaceContext) AllEnum() []IEnumContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumContext); ok {
			len++
		}
	}

	tst := make([]IEnumContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumContext); ok {
			tst[i] = t.(IEnumContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceContext) Enum(i int) IEnumContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumContext)
}

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (s *NamespaceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitNamespace(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Namespace() (localctx INamespaceContext) {
	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConstMakerParserRULE_namespace)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(143)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserDOC_COMMENT {
		{
			p.SetState(149)
			p.DocComment()
		}

	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConstMakerParserNL {
		{
			p.SetState(152)
			p.Match(ConstMakerParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(158)
		p.Match(ConstMakerParserNAMESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Ident()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__4 {
		{
			p.SetState(160)
			p.Option()
		}

	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConstMakerParserNL {
		{
			p.SetState(163)
			p.Match(ConstMakerParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
		p.Match(ConstMakerParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(170)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(176)
					p.Namespace()
				}

			case 2:
				{
					p.SetState(177)
					p.Const_()
				}

			case 3:
				{
					p.SetState(178)
					p.Enum()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConstMakerParserNL {
		{
			p.SetState(184)
			p.Match(ConstMakerParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
		p.Match(ConstMakerParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(191)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Integer() IIntegerContext
	TYPE() antlr.TerminalNode
	Ident() IIdentContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *TypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserTYPE, 0)
}

func (s *TypeContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConstMakerParserRULE_type)
	var _la int

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ConstMakerParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.Match(ConstMakerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Type_()
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ConstMakerParserT__0 {
			{
				p.SetState(199)
				p.Match(ConstMakerParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.Integer()
			}

		}
		{
			p.SetState(203)
			p.Match(ConstMakerParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ConstMakerParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.Match(ConstMakerParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ConstMakerParserENUM, ConstMakerParserNAMESPACE, ConstMakerParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(206)
			p.Ident()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	Integer() IIntegerContext
	Option() IOptionContext
	TripleComment() ITripleCommentContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_enumValue
	return p
}

func InitEmptyEnumValueContext(p *EnumValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_enumValue
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumValueContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *EnumValueContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *EnumValueContext) TripleComment() ITripleCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITripleCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITripleCommentContext)
}

func (s *EnumValueContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConstMakerParserNL)
}

func (s *EnumValueContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConstMakerParserNL, i)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (s *EnumValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitEnumValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConstMakerParserRULE_enumValue)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Ident()
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__1 {
		{
			p.SetState(210)
			p.Match(ConstMakerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Integer()
		}

	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__4 {
		{
			p.SetState(214)
			p.Option()
		}

	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserTRIPLE_COMMENT {
		{
			p.SetState(217)
			p.TripleComment()
		}

	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(220)
				p.Match(ConstMakerParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentAssignValueContext is an interface to support dynamic dispatch.
type IIdentAssignValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	Value() IValueContext

	// IsIdentAssignValueContext differentiates from other interfaces.
	IsIdentAssignValueContext()
}

type IdentAssignValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentAssignValueContext() *IdentAssignValueContext {
	var p = new(IdentAssignValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_identAssignValue
	return p
}

func InitEmptyIdentAssignValueContext(p *IdentAssignValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_identAssignValue
}

func (*IdentAssignValueContext) IsIdentAssignValueContext() {}

func NewIdentAssignValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentAssignValueContext {
	var p = new(IdentAssignValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_identAssignValue

	return p
}

func (s *IdentAssignValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentAssignValueContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *IdentAssignValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *IdentAssignValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentAssignValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentAssignValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterIdentAssignValue(s)
	}
}

func (s *IdentAssignValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitIdentAssignValue(s)
	}
}

func (s *IdentAssignValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitIdentAssignValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) IdentAssignValue() (localctx IIdentAssignValueContext) {
	localctx = NewIdentAssignValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ConstMakerParserRULE_identAssignValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Ident()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConstMakerParserT__0 {
		{
			p.SetState(227)
			p.Match(ConstMakerParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Value()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentAssignValueLoopContext is an interface to support dynamic dispatch.
type IIdentAssignValueLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentAssignValue() []IIdentAssignValueContext
	IdentAssignValue(i int) IIdentAssignValueContext

	// IsIdentAssignValueLoopContext differentiates from other interfaces.
	IsIdentAssignValueLoopContext()
}

type IdentAssignValueLoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentAssignValueLoopContext() *IdentAssignValueLoopContext {
	var p = new(IdentAssignValueLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_identAssignValueLoop
	return p
}

func InitEmptyIdentAssignValueLoopContext(p *IdentAssignValueLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_identAssignValueLoop
}

func (*IdentAssignValueLoopContext) IsIdentAssignValueLoopContext() {}

func NewIdentAssignValueLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentAssignValueLoopContext {
	var p = new(IdentAssignValueLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_identAssignValueLoop

	return p
}

func (s *IdentAssignValueLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentAssignValueLoopContext) AllIdentAssignValue() []IIdentAssignValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentAssignValueContext); ok {
			len++
		}
	}

	tst := make([]IIdentAssignValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentAssignValueContext); ok {
			tst[i] = t.(IIdentAssignValueContext)
			i++
		}
	}

	return tst
}

func (s *IdentAssignValueLoopContext) IdentAssignValue(i int) IIdentAssignValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentAssignValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentAssignValueContext)
}

func (s *IdentAssignValueLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentAssignValueLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentAssignValueLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterIdentAssignValueLoop(s)
	}
}

func (s *IdentAssignValueLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitIdentAssignValueLoop(s)
	}
}

func (s *IdentAssignValueLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitIdentAssignValueLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) IdentAssignValueLoop() (localctx IIdentAssignValueLoopContext) {
	localctx = NewIdentAssignValueLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ConstMakerParserRULE_identAssignValueLoop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.IdentAssignValue()
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ConstMakerParserT__6 {
		{
			p.SetState(232)
			p.Match(ConstMakerParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.IdentAssignValue()
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentAssignValueLoop() IIdentAssignValueLoopContext

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_option
	return p
}

func InitEmptyOptionContext(p *OptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_option
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) IdentAssignValueLoop() IIdentAssignValueLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentAssignValueLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentAssignValueLoopContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitOption(s)
	}
}

func (s *OptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ConstMakerParserRULE_option)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(ConstMakerParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.IdentAssignValueLoop()
	}
	{
		p.SetState(241)
		p.Match(ConstMakerParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	Ident() IIdentContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserFLOAT, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserSTRING, 0)
}

func (s *ValueContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ConstMakerParserRULE_value)
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ConstMakerParserDEC_INTEGER, ConstMakerParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.Integer()
		}

	case ConstMakerParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.Match(ConstMakerParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ConstMakerParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(245)
			p.Match(ConstMakerParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ConstMakerParserENUM, ConstMakerParserNAMESPACE, ConstMakerParserIDENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)
			p.Ident()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEC_INTEGER() antlr.TerminalNode
	HEX_INTEGER() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DEC_INTEGER() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserDEC_INTEGER, 0)
}

func (s *IntegerContext) HEX_INTEGER() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserHEX_INTEGER, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ConstMakerParserRULE_integer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ConstMakerParserDEC_INTEGER || _la == ConstMakerParserHEX_INTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	Keywords() IKeywordsContext

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserIDENT, 0)
}

func (s *IdentContext) Keywords() IKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ConstMakerParserRULE_ident)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ConstMakerParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			p.Match(ConstMakerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ConstMakerParserENUM, ConstMakerParserNAMESPACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
			p.Keywords()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordsContext is an interface to support dynamic dispatch.
type IKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	NAMESPACE() antlr.TerminalNode

	// IsKeywordsContext differentiates from other interfaces.
	IsKeywordsContext()
}

type KeywordsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordsContext() *KeywordsContext {
	var p = new(KeywordsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_keywords
	return p
}

func InitEmptyKeywordsContext(p *KeywordsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_keywords
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserENUM, 0)
}

func (s *KeywordsContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserNAMESPACE, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterKeywords(s)
	}
}

func (s *KeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitKeywords(s)
	}
}

func (s *KeywordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitKeywords(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) Keywords() (localctx IKeywordsContext) {
	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ConstMakerParserRULE_keywords)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ConstMakerParserENUM || _la == ConstMakerParserNAMESPACE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDocCommentContext is an interface to support dynamic dispatch.
type IDocCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOC_COMMENT() antlr.TerminalNode

	// IsDocCommentContext differentiates from other interfaces.
	IsDocCommentContext()
}

type DocCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocCommentContext() *DocCommentContext {
	var p = new(DocCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_docComment
	return p
}

func InitEmptyDocCommentContext(p *DocCommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_docComment
}

func (*DocCommentContext) IsDocCommentContext() {}

func NewDocCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocCommentContext {
	var p = new(DocCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_docComment

	return p
}

func (s *DocCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocCommentContext) DOC_COMMENT() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserDOC_COMMENT, 0)
}

func (s *DocCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterDocComment(s)
	}
}

func (s *DocCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitDocComment(s)
	}
}

func (s *DocCommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitDocComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) DocComment() (localctx IDocCommentContext) {
	localctx = NewDocCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ConstMakerParserRULE_docComment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(ConstMakerParserDOC_COMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITripleCommentContext is an interface to support dynamic dispatch.
type ITripleCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRIPLE_COMMENT() antlr.TerminalNode

	// IsTripleCommentContext differentiates from other interfaces.
	IsTripleCommentContext()
}

type TripleCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTripleCommentContext() *TripleCommentContext {
	var p = new(TripleCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_tripleComment
	return p
}

func InitEmptyTripleCommentContext(p *TripleCommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConstMakerParserRULE_tripleComment
}

func (*TripleCommentContext) IsTripleCommentContext() {}

func NewTripleCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TripleCommentContext {
	var p = new(TripleCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstMakerParserRULE_tripleComment

	return p
}

func (s *TripleCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *TripleCommentContext) TRIPLE_COMMENT() antlr.TerminalNode {
	return s.GetToken(ConstMakerParserTRIPLE_COMMENT, 0)
}

func (s *TripleCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TripleCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TripleCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.EnterTripleComment(s)
	}
}

func (s *TripleCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstMakerListener); ok {
		listenerT.ExitTripleComment(s)
	}
}

func (s *TripleCommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConstMakerVisitor:
		return t.VisitTripleComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConstMakerParser) TripleComment() (localctx ITripleCommentContext) {
	localctx = NewTripleCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ConstMakerParserRULE_tripleComment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(ConstMakerParserTRIPLE_COMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
