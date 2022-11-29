// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // CDDL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CDDLParser struct {
	*antlr.BaseParser
}

var cddlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cddlParserInit() {
	staticData := &cddlParserStaticData
	staticData.literalNames = []string{
		"", "'='", "'/='", "'//='", "'<'", "','", "'>'", "'/'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'~'", "'&'", "'#'", "'//'", "'^'", "'=>'",
		"':'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "TAG1", "TAG2", "RANGEOP", "CTLOP", "OCCUR", "VALUE",
		"ID", "S",
	}
	staticData.ruleNames = []string{
		"cddl", "rule", "typeRule", "groupRule", "assignType", "assignGroup",
		"genericParam", "genericArg", "type", "type1", "type2", "group", "groupChoice",
		"groupEntry", "memberKey", "optComma",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 330, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 3, 0, 34, 8, 0, 1, 0, 1, 0, 3, 0, 38, 8, 0, 4, 0, 40, 8, 0, 11, 0,
		12, 0, 41, 1, 0, 3, 0, 45, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 51, 8, 1,
		1, 2, 1, 2, 3, 2, 55, 8, 2, 1, 2, 3, 2, 58, 8, 2, 1, 2, 1, 2, 3, 2, 62,
		8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 68, 8, 3, 1, 3, 3, 3, 71, 8, 3, 1,
		3, 1, 3, 3, 3, 75, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		3, 6, 85, 8, 6, 1, 6, 1, 6, 3, 6, 89, 8, 6, 1, 6, 1, 6, 3, 6, 93, 8, 6,
		1, 6, 1, 6, 3, 6, 97, 8, 6, 5, 6, 99, 8, 6, 10, 6, 12, 6, 102, 9, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 3, 7, 108, 8, 7, 1, 7, 1, 7, 3, 7, 112, 8, 7, 1, 7,
		1, 7, 3, 7, 116, 8, 7, 1, 7, 1, 7, 3, 7, 120, 8, 7, 5, 7, 122, 8, 7, 10,
		7, 12, 7, 125, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 131, 8, 8, 1, 8, 1,
		8, 3, 8, 135, 8, 8, 1, 8, 5, 8, 138, 8, 8, 10, 8, 12, 8, 141, 9, 8, 1,
		9, 1, 9, 3, 9, 145, 8, 9, 1, 9, 1, 9, 3, 9, 149, 8, 9, 1, 9, 3, 9, 152,
		8, 9, 1, 10, 1, 10, 1, 10, 3, 10, 157, 8, 10, 1, 10, 1, 10, 3, 10, 161,
		8, 10, 1, 10, 1, 10, 3, 10, 165, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 171, 8, 10, 1, 10, 1, 10, 3, 10, 175, 8, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 181, 8, 10, 1, 10, 1, 10, 3, 10, 185, 8, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 191, 8, 10, 1, 10, 1, 10, 3, 10, 195, 8, 10, 1, 10, 1,
		10, 3, 10, 199, 8, 10, 1, 10, 1, 10, 3, 10, 203, 8, 10, 1, 10, 1, 10, 3,
		10, 207, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 213, 8, 10, 1, 10, 1,
		10, 3, 10, 217, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 222, 8, 10, 1, 10, 1,
		10, 3, 10, 226, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 232, 8, 10, 1,
		11, 1, 11, 3, 11, 236, 8, 11, 1, 11, 1, 11, 3, 11, 240, 8, 11, 1, 11, 5,
		11, 243, 8, 11, 10, 11, 12, 11, 246, 9, 11, 1, 12, 1, 12, 1, 12, 5, 12,
		251, 8, 12, 10, 12, 12, 12, 254, 9, 12, 1, 13, 1, 13, 3, 13, 258, 8, 13,
		3, 13, 260, 8, 13, 1, 13, 1, 13, 3, 13, 264, 8, 13, 3, 13, 266, 8, 13,
		1, 13, 1, 13, 1, 13, 3, 13, 271, 8, 13, 3, 13, 273, 8, 13, 1, 13, 1, 13,
		3, 13, 277, 8, 13, 1, 13, 1, 13, 3, 13, 281, 8, 13, 3, 13, 283, 8, 13,
		1, 13, 1, 13, 3, 13, 287, 8, 13, 1, 13, 1, 13, 3, 13, 291, 8, 13, 1, 13,
		1, 13, 3, 13, 295, 8, 13, 1, 14, 1, 14, 3, 14, 299, 8, 14, 1, 14, 1, 14,
		3, 14, 303, 8, 14, 3, 14, 305, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14,
		311, 8, 14, 1, 14, 1, 14, 1, 14, 3, 14, 316, 8, 14, 1, 14, 3, 14, 319,
		8, 14, 1, 15, 3, 15, 322, 8, 15, 1, 15, 3, 15, 325, 8, 15, 1, 15, 3, 15,
		328, 8, 15, 1, 15, 0, 0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 0, 3, 1, 0, 1, 2, 2, 0, 1, 1, 3, 3, 1, 0, 23, 24, 393,
		0, 33, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 52, 1, 0, 0, 0, 6, 65, 1, 0, 0,
		0, 8, 78, 1, 0, 0, 0, 10, 80, 1, 0, 0, 0, 12, 82, 1, 0, 0, 0, 14, 105,
		1, 0, 0, 0, 16, 128, 1, 0, 0, 0, 18, 142, 1, 0, 0, 0, 20, 231, 1, 0, 0,
		0, 22, 233, 1, 0, 0, 0, 24, 252, 1, 0, 0, 0, 26, 294, 1, 0, 0, 0, 28, 318,
		1, 0, 0, 0, 30, 321, 1, 0, 0, 0, 32, 34, 5, 28, 0, 0, 33, 32, 1, 0, 0,
		0, 33, 34, 1, 0, 0, 0, 34, 39, 1, 0, 0, 0, 35, 37, 3, 2, 1, 0, 36, 38,
		5, 28, 0, 0, 37, 36, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0,
		39, 35, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1,
		0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 45, 5, 28, 0, 0, 44, 43, 1, 0, 0, 0, 44,
		45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 5, 0, 0, 1, 47, 1, 1, 0, 0,
		0, 48, 51, 3, 4, 2, 0, 49, 51, 3, 6, 3, 0, 50, 48, 1, 0, 0, 0, 50, 49,
		1, 0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 54, 5, 27, 0, 0, 53, 55, 3, 12, 6, 0,
		54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0, 0, 56, 58, 5,
		28, 0, 0, 57, 56, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59,
		61, 3, 8, 4, 0, 60, 62, 5, 28, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0,
		0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 3, 16, 8, 0, 64, 5, 1, 0, 0, 0, 65, 67,
		5, 27, 0, 0, 66, 68, 3, 12, 6, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0,
		0, 68, 70, 1, 0, 0, 0, 69, 71, 5, 28, 0, 0, 70, 69, 1, 0, 0, 0, 70, 71,
		1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74, 3, 10, 5, 0, 73, 75, 5, 28, 0,
		0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77,
		3, 26, 13, 0, 77, 7, 1, 0, 0, 0, 78, 79, 7, 0, 0, 0, 79, 9, 1, 0, 0, 0,
		80, 81, 7, 1, 0, 0, 81, 11, 1, 0, 0, 0, 82, 84, 5, 4, 0, 0, 83, 85, 5,
		28, 0, 0, 84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86,
		88, 5, 27, 0, 0, 87, 89, 5, 28, 0, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0,
		0, 0, 89, 100, 1, 0, 0, 0, 90, 92, 5, 5, 0, 0, 91, 93, 5, 28, 0, 0, 92,
		91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 5, 27,
		0, 0, 95, 97, 5, 28, 0, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97,
		99, 1, 0, 0, 0, 98, 90, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0,
		0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0,
		103, 104, 5, 6, 0, 0, 104, 13, 1, 0, 0, 0, 105, 107, 5, 4, 0, 0, 106, 108,
		5, 28, 0, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0,
		0, 0, 109, 111, 3, 18, 9, 0, 110, 112, 5, 28, 0, 0, 111, 110, 1, 0, 0,
		0, 111, 112, 1, 0, 0, 0, 112, 123, 1, 0, 0, 0, 113, 115, 5, 5, 0, 0, 114,
		116, 5, 28, 0, 0, 115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 119, 3, 18, 9, 0, 118, 120, 5, 28, 0, 0, 119, 118, 1,
		0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121, 113, 1, 0, 0,
		0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124,
		126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 6, 0, 0, 127, 15, 1,
		0, 0, 0, 128, 139, 3, 18, 9, 0, 129, 131, 5, 28, 0, 0, 130, 129, 1, 0,
		0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 5, 7, 0, 0,
		133, 135, 5, 28, 0, 0, 134, 133, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135,
		136, 1, 0, 0, 0, 136, 138, 3, 18, 9, 0, 137, 130, 1, 0, 0, 0, 138, 141,
		1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 17, 1, 0,
		0, 0, 141, 139, 1, 0, 0, 0, 142, 151, 3, 20, 10, 0, 143, 145, 5, 28, 0,
		0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		148, 7, 2, 0, 0, 147, 149, 5, 28, 0, 0, 148, 147, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 3, 20, 10, 0, 151, 144, 1,
		0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 19, 1, 0, 0, 0, 153, 232, 5, 26, 0,
		0, 154, 156, 5, 27, 0, 0, 155, 157, 3, 14, 7, 0, 156, 155, 1, 0, 0, 0,
		156, 157, 1, 0, 0, 0, 157, 232, 1, 0, 0, 0, 158, 160, 5, 8, 0, 0, 159,
		161, 5, 28, 0, 0, 160, 159, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162,
		1, 0, 0, 0, 162, 164, 3, 16, 8, 0, 163, 165, 5, 28, 0, 0, 164, 163, 1,
		0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 5, 9, 0,
		0, 167, 232, 1, 0, 0, 0, 168, 170, 5, 10, 0, 0, 169, 171, 5, 28, 0, 0,
		170, 169, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172,
		174, 3, 22, 11, 0, 173, 175, 5, 28, 0, 0, 174, 173, 1, 0, 0, 0, 174, 175,
		1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 5, 11, 0, 0, 177, 232, 1, 0,
		0, 0, 178, 180, 5, 12, 0, 0, 179, 181, 5, 28, 0, 0, 180, 179, 1, 0, 0,
		0, 180, 181, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 3, 22, 11, 0,
		183, 185, 5, 28, 0, 0, 184, 183, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185,
		186, 1, 0, 0, 0, 186, 187, 5, 13, 0, 0, 187, 232, 1, 0, 0, 0, 188, 190,
		5, 14, 0, 0, 189, 191, 5, 28, 0, 0, 190, 189, 1, 0, 0, 0, 190, 191, 1,
		0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 5, 27, 0, 0, 193, 195, 3, 14,
		7, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 232, 1, 0, 0, 0,
		196, 198, 5, 15, 0, 0, 197, 199, 5, 28, 0, 0, 198, 197, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 202, 5, 8, 0, 0, 201, 203,
		5, 28, 0, 0, 202, 201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 1, 0,
		0, 0, 204, 206, 3, 22, 11, 0, 205, 207, 5, 28, 0, 0, 206, 205, 1, 0, 0,
		0, 206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 5, 9, 0, 0, 209,
		232, 1, 0, 0, 0, 210, 212, 5, 15, 0, 0, 211, 213, 5, 28, 0, 0, 212, 211,
		1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216, 5, 27,
		0, 0, 215, 217, 3, 14, 7, 0, 216, 215, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0,
		217, 232, 1, 0, 0, 0, 218, 219, 5, 21, 0, 0, 219, 221, 5, 8, 0, 0, 220,
		222, 5, 28, 0, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223,
		1, 0, 0, 0, 223, 225, 3, 16, 8, 0, 224, 226, 5, 28, 0, 0, 225, 224, 1,
		0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 5, 9, 0,
		0, 228, 232, 1, 0, 0, 0, 229, 232, 5, 22, 0, 0, 230, 232, 5, 16, 0, 0,
		231, 153, 1, 0, 0, 0, 231, 154, 1, 0, 0, 0, 231, 158, 1, 0, 0, 0, 231,
		168, 1, 0, 0, 0, 231, 178, 1, 0, 0, 0, 231, 188, 1, 0, 0, 0, 231, 196,
		1, 0, 0, 0, 231, 210, 1, 0, 0, 0, 231, 218, 1, 0, 0, 0, 231, 229, 1, 0,
		0, 0, 231, 230, 1, 0, 0, 0, 232, 21, 1, 0, 0, 0, 233, 244, 3, 24, 12, 0,
		234, 236, 5, 28, 0, 0, 235, 234, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236,
		237, 1, 0, 0, 0, 237, 239, 5, 17, 0, 0, 238, 240, 5, 28, 0, 0, 239, 238,
		1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 243, 3, 24,
		12, 0, 242, 235, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0,
		244, 245, 1, 0, 0, 0, 245, 23, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 247, 248,
		3, 26, 13, 0, 248, 249, 3, 30, 15, 0, 249, 251, 1, 0, 0, 0, 250, 247, 1,
		0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0,
		0, 253, 25, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 257, 5, 25, 0, 0, 256,
		258, 5, 28, 0, 0, 257, 256, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260,
		1, 0, 0, 0, 259, 255, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 265, 1, 0,
		0, 0, 261, 263, 3, 28, 14, 0, 262, 264, 5, 28, 0, 0, 263, 262, 1, 0, 0,
		0, 263, 264, 1, 0, 0, 0, 264, 266, 1, 0, 0, 0, 265, 261, 1, 0, 0, 0, 265,
		266, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 295, 3, 16, 8, 0, 268, 270,
		5, 25, 0, 0, 269, 271, 5, 28, 0, 0, 270, 269, 1, 0, 0, 0, 270, 271, 1,
		0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 268, 1, 0, 0, 0, 272, 273, 1, 0, 0,
		0, 273, 274, 1, 0, 0, 0, 274, 276, 5, 27, 0, 0, 275, 277, 3, 14, 7, 0,
		276, 275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 295, 1, 0, 0, 0, 278,
		280, 5, 25, 0, 0, 279, 281, 5, 28, 0, 0, 280, 279, 1, 0, 0, 0, 280, 281,
		1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282, 278, 1, 0, 0, 0, 282, 283, 1, 0,
		0, 0, 283, 284, 1, 0, 0, 0, 284, 286, 5, 8, 0, 0, 285, 287, 5, 28, 0, 0,
		286, 285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288,
		290, 3, 22, 11, 0, 289, 291, 5, 28, 0, 0, 290, 289, 1, 0, 0, 0, 290, 291,
		1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293, 5, 9, 0, 0, 293, 295, 1, 0,
		0, 0, 294, 259, 1, 0, 0, 0, 294, 272, 1, 0, 0, 0, 294, 282, 1, 0, 0, 0,
		295, 27, 1, 0, 0, 0, 296, 298, 3, 18, 9, 0, 297, 299, 5, 28, 0, 0, 298,
		297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 304, 1, 0, 0, 0, 300, 302,
		5, 18, 0, 0, 301, 303, 5, 28, 0, 0, 302, 301, 1, 0, 0, 0, 302, 303, 1,
		0, 0, 0, 303, 305, 1, 0, 0, 0, 304, 300, 1, 0, 0, 0, 304, 305, 1, 0, 0,
		0, 305, 306, 1, 0, 0, 0, 306, 307, 5, 19, 0, 0, 307, 319, 1, 0, 0, 0, 308,
		310, 5, 27, 0, 0, 309, 311, 5, 28, 0, 0, 310, 309, 1, 0, 0, 0, 310, 311,
		1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 319, 5, 20, 0, 0, 313, 315, 5, 26,
		0, 0, 314, 316, 5, 28, 0, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0,
		316, 317, 1, 0, 0, 0, 317, 319, 5, 20, 0, 0, 318, 296, 1, 0, 0, 0, 318,
		308, 1, 0, 0, 0, 318, 313, 1, 0, 0, 0, 319, 29, 1, 0, 0, 0, 320, 322, 5,
		28, 0, 0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 324, 1, 0, 0,
		0, 323, 325, 5, 5, 0, 0, 324, 323, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325,
		327, 1, 0, 0, 0, 326, 328, 5, 28, 0, 0, 327, 326, 1, 0, 0, 0, 327, 328,
		1, 0, 0, 0, 328, 31, 1, 0, 0, 0, 69, 33, 37, 41, 44, 50, 54, 57, 61, 67,
		70, 74, 84, 88, 92, 96, 100, 107, 111, 115, 119, 123, 130, 134, 139, 144,
		148, 151, 156, 160, 164, 170, 174, 180, 184, 190, 194, 198, 202, 206, 212,
		216, 221, 225, 231, 235, 239, 244, 252, 257, 259, 263, 265, 270, 272, 276,
		280, 282, 286, 290, 294, 298, 302, 304, 310, 315, 318, 321, 324, 327,
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

// CDDLParserInit initializes any static state used to implement CDDLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCDDLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CDDLParserInit() {
	staticData := &cddlParserStaticData
	staticData.once.Do(cddlParserInit)
}

// NewCDDLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCDDLParser(input antlr.TokenStream) *CDDLParser {
	CDDLParserInit()
	this := new(CDDLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cddlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// CDDLParser tokens.
const (
	CDDLParserEOF     = antlr.TokenEOF
	CDDLParserT__0    = 1
	CDDLParserT__1    = 2
	CDDLParserT__2    = 3
	CDDLParserT__3    = 4
	CDDLParserT__4    = 5
	CDDLParserT__5    = 6
	CDDLParserT__6    = 7
	CDDLParserT__7    = 8
	CDDLParserT__8    = 9
	CDDLParserT__9    = 10
	CDDLParserT__10   = 11
	CDDLParserT__11   = 12
	CDDLParserT__12   = 13
	CDDLParserT__13   = 14
	CDDLParserT__14   = 15
	CDDLParserT__15   = 16
	CDDLParserT__16   = 17
	CDDLParserT__17   = 18
	CDDLParserT__18   = 19
	CDDLParserT__19   = 20
	CDDLParserTAG1    = 21
	CDDLParserTAG2    = 22
	CDDLParserRANGEOP = 23
	CDDLParserCTLOP   = 24
	CDDLParserOCCUR   = 25
	CDDLParserVALUE   = 26
	CDDLParserID      = 27
	CDDLParserS       = 28
)

// CDDLParser rules.
const (
	CDDLParserRULE_cddl         = 0
	CDDLParserRULE_rule         = 1
	CDDLParserRULE_typeRule     = 2
	CDDLParserRULE_groupRule    = 3
	CDDLParserRULE_assignType   = 4
	CDDLParserRULE_assignGroup  = 5
	CDDLParserRULE_genericParam = 6
	CDDLParserRULE_genericArg   = 7
	CDDLParserRULE_type         = 8
	CDDLParserRULE_type1        = 9
	CDDLParserRULE_type2        = 10
	CDDLParserRULE_group        = 11
	CDDLParserRULE_groupChoice  = 12
	CDDLParserRULE_groupEntry   = 13
	CDDLParserRULE_memberKey    = 14
	CDDLParserRULE_optComma     = 15
)

// ICddlContext is an interface to support dynamic dispatch.
type ICddlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCddlContext differentiates from other interfaces.
	IsCddlContext()
}

type CddlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCddlContext() *CddlContext {
	var p = new(CddlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_cddl
	return p
}

func (*CddlContext) IsCddlContext() {}

func NewCddlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CddlContext {
	var p = new(CddlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_cddl

	return p
}

func (s *CddlContext) GetParser() antlr.Parser { return s.parser }

func (s *CddlContext) EOF() antlr.TerminalNode {
	return s.GetToken(CDDLParserEOF, 0)
}

func (s *CddlContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *CddlContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *CddlContext) AllRule_() []IRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleContext); ok {
			len++
		}
	}

	tst := make([]IRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleContext); ok {
			tst[i] = t.(IRuleContext)
			i++
		}
	}

	return tst
}

func (s *CddlContext) Rule_(i int) IRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleContext); ok {
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

	return t.(IRuleContext)
}

func (s *CddlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CddlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CddlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterCddl(s)
	}
}

func (s *CddlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitCddl(s)
	}
}

func (p *CDDLParser) Cddl() (localctx ICddlContext) {
	this := p
	_ = this

	localctx = NewCddlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CDDLParserRULE_cddl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(32)
			p.Match(CDDLParserS)
		}

	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CDDLParserID {
		{
			p.SetState(35)
			p.Rule_()
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(36)
				p.Match(CDDLParserS)
			}

		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(43)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(46)
		p.Match(CDDLParserEOF)
	}

	return localctx
}

// IRuleContext is an interface to support dynamic dispatch.
type IRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleContext differentiates from other interfaces.
	IsRuleContext()
}

type RuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContext() *RuleContext {
	var p = new(RuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_rule
	return p
}

func (*RuleContext) IsRuleContext() {}

func NewRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContext {
	var p = new(RuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_rule

	return p
}

func (s *RuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContext) TypeRule() ITypeRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *RuleContext) GroupRule() IGroupRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupRuleContext)
}

func (s *RuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterRule(s)
	}
}

func (s *RuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitRule(s)
	}
}

func (p *CDDLParser) Rule_() (localctx IRuleContext) {
	this := p
	_ = this

	localctx = NewRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CDDLParserRULE_rule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.TypeRule()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.GroupRule()
		}

	}

	return localctx
}

// ITypeRuleContext is an interface to support dynamic dispatch.
type ITypeRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRuleContext differentiates from other interfaces.
	IsTypeRuleContext()
}

type TypeRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRuleContext() *TypeRuleContext {
	var p = new(TypeRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_typeRule
	return p
}

func (*TypeRuleContext) IsTypeRuleContext() {}

func NewTypeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_typeRule

	return p
}

func (s *TypeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRuleContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *TypeRuleContext) AssignType() IAssignTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignTypeContext)
}

func (s *TypeRuleContext) Type_() ITypeContext {
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

func (s *TypeRuleContext) GenericParam() IGenericParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericParamContext)
}

func (s *TypeRuleContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *TypeRuleContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterTypeRule(s)
	}
}

func (s *TypeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitTypeRule(s)
	}
}

func (p *CDDLParser) TypeRule() (localctx ITypeRuleContext) {
	this := p
	_ = this

	localctx = NewTypeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CDDLParserRULE_typeRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(CDDLParserID)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserT__3 {
		{
			p.SetState(53)
			p.GenericParam()
		}

	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(56)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(59)
		p.AssignType()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(60)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(63)
		p.Type_()
	}

	return localctx
}

// IGroupRuleContext is an interface to support dynamic dispatch.
type IGroupRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupRuleContext differentiates from other interfaces.
	IsGroupRuleContext()
}

type GroupRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupRuleContext() *GroupRuleContext {
	var p = new(GroupRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_groupRule
	return p
}

func (*GroupRuleContext) IsGroupRuleContext() {}

func NewGroupRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupRuleContext {
	var p = new(GroupRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_groupRule

	return p
}

func (s *GroupRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupRuleContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *GroupRuleContext) AssignGroup() IAssignGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignGroupContext)
}

func (s *GroupRuleContext) GroupEntry() IGroupEntryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupEntryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupEntryContext)
}

func (s *GroupRuleContext) GenericParam() IGenericParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericParamContext)
}

func (s *GroupRuleContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *GroupRuleContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *GroupRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterGroupRule(s)
	}
}

func (s *GroupRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitGroupRule(s)
	}
}

func (p *CDDLParser) GroupRule() (localctx IGroupRuleContext) {
	this := p
	_ = this

	localctx = NewGroupRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CDDLParserRULE_groupRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(CDDLParserID)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserT__3 {
		{
			p.SetState(66)
			p.GenericParam()
		}

	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(69)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(72)
		p.AssignGroup()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(73)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(76)
		p.GroupEntry()
	}

	return localctx
}

// IAssignTypeContext is an interface to support dynamic dispatch.
type IAssignTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignTypeContext differentiates from other interfaces.
	IsAssignTypeContext()
}

type AssignTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignTypeContext() *AssignTypeContext {
	var p = new(AssignTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_assignType
	return p
}

func (*AssignTypeContext) IsAssignTypeContext() {}

func NewAssignTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignTypeContext {
	var p = new(AssignTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_assignType

	return p
}

func (s *AssignTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterAssignType(s)
	}
}

func (s *AssignTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitAssignType(s)
	}
}

func (p *CDDLParser) AssignType() (localctx IAssignTypeContext) {
	this := p
	_ = this

	localctx = NewAssignTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CDDLParserRULE_assignType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CDDLParserT__0 || _la == CDDLParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignGroupContext is an interface to support dynamic dispatch.
type IAssignGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignGroupContext differentiates from other interfaces.
	IsAssignGroupContext()
}

type AssignGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignGroupContext() *AssignGroupContext {
	var p = new(AssignGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_assignGroup
	return p
}

func (*AssignGroupContext) IsAssignGroupContext() {}

func NewAssignGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignGroupContext {
	var p = new(AssignGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_assignGroup

	return p
}

func (s *AssignGroupContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterAssignGroup(s)
	}
}

func (s *AssignGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitAssignGroup(s)
	}
}

func (p *CDDLParser) AssignGroup() (localctx IAssignGroupContext) {
	this := p
	_ = this

	localctx = NewAssignGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CDDLParserRULE_assignGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CDDLParserT__0 || _la == CDDLParserT__2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGenericParamContext is an interface to support dynamic dispatch.
type IGenericParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericParamContext differentiates from other interfaces.
	IsGenericParamContext()
}

type GenericParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericParamContext() *GenericParamContext {
	var p = new(GenericParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_genericParam
	return p
}

func (*GenericParamContext) IsGenericParamContext() {}

func NewGenericParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericParamContext {
	var p = new(GenericParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_genericParam

	return p
}

func (s *GenericParamContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserID)
}

func (s *GenericParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserID, i)
}

func (s *GenericParamContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *GenericParamContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *GenericParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterGenericParam(s)
	}
}

func (s *GenericParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitGenericParam(s)
	}
}

func (p *CDDLParser) GenericParam() (localctx IGenericParamContext) {
	this := p
	_ = this

	localctx = NewGenericParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CDDLParserRULE_genericParam)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(CDDLParserT__3)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(83)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(86)
		p.Match(CDDLParserID)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(87)
			p.Match(CDDLParserS)
		}

	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDDLParserT__4 {
		{
			p.SetState(90)
			p.Match(CDDLParserT__4)
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(91)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(94)
			p.Match(CDDLParserID)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(95)
				p.Match(CDDLParserS)
			}

		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(CDDLParserT__5)
	}

	return localctx
}

// IGenericArgContext is an interface to support dynamic dispatch.
type IGenericArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericArgContext differentiates from other interfaces.
	IsGenericArgContext()
}

type GenericArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericArgContext() *GenericArgContext {
	var p = new(GenericArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_genericArg
	return p
}

func (*GenericArgContext) IsGenericArgContext() {}

func NewGenericArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericArgContext {
	var p = new(GenericArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_genericArg

	return p
}

func (s *GenericArgContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericArgContext) AllType1() []IType1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType1Context); ok {
			len++
		}
	}

	tst := make([]IType1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType1Context); ok {
			tst[i] = t.(IType1Context)
			i++
		}
	}

	return tst
}

func (s *GenericArgContext) Type1(i int) IType1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType1Context); ok {
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

	return t.(IType1Context)
}

func (s *GenericArgContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *GenericArgContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *GenericArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterGenericArg(s)
	}
}

func (s *GenericArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitGenericArg(s)
	}
}

func (p *CDDLParser) GenericArg() (localctx IGenericArgContext) {
	this := p
	_ = this

	localctx = NewGenericArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CDDLParserRULE_genericArg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(CDDLParserT__3)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(106)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(109)
		p.Type1()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(110)
			p.Match(CDDLParserS)
		}

	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDDLParserT__4 {
		{
			p.SetState(113)
			p.Match(CDDLParserT__4)
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(114)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(117)
			p.Type1()
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(118)
				p.Match(CDDLParserS)
			}

		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(CDDLParserT__5)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) AllType1() []IType1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType1Context); ok {
			len++
		}
	}

	tst := make([]IType1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType1Context); ok {
			tst[i] = t.(IType1Context)
			i++
		}
	}

	return tst
}

func (s *TypeContext) Type1(i int) IType1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType1Context); ok {
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

	return t.(IType1Context)
}

func (s *TypeContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *TypeContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *CDDLParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CDDLParserRULE_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Type1()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(129)
					p.Match(CDDLParserS)
				}

			}
			{
				p.SetState(132)
				p.Match(CDDLParserT__6)
			}
			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(133)
					p.Match(CDDLParserS)
				}

			}
			{
				p.SetState(136)
				p.Type1()
			}

		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IType1Context is an interface to support dynamic dispatch.
type IType1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType1Context differentiates from other interfaces.
	IsType1Context()
}

type Type1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType1Context() *Type1Context {
	var p = new(Type1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_type1
	return p
}

func (*Type1Context) IsType1Context() {}

func NewType1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type1Context {
	var p = new(Type1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_type1

	return p
}

func (s *Type1Context) GetParser() antlr.Parser { return s.parser }

func (s *Type1Context) AllType2() []IType2Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType2Context); ok {
			len++
		}
	}

	tst := make([]IType2Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType2Context); ok {
			tst[i] = t.(IType2Context)
			i++
		}
	}

	return tst
}

func (s *Type1Context) Type2(i int) IType2Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType2Context); ok {
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

	return t.(IType2Context)
}

func (s *Type1Context) RANGEOP() antlr.TerminalNode {
	return s.GetToken(CDDLParserRANGEOP, 0)
}

func (s *Type1Context) CTLOP() antlr.TerminalNode {
	return s.GetToken(CDDLParserCTLOP, 0)
}

func (s *Type1Context) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *Type1Context) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *Type1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterType1(s)
	}
}

func (s *Type1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitType1(s)
	}
}

func (p *CDDLParser) Type1() (localctx IType1Context) {
	this := p
	_ = this

	localctx = NewType1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CDDLParserRULE_type1)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Type2()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(143)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(146)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CDDLParserRANGEOP || _la == CDDLParserCTLOP) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(147)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(150)
			p.Type2()
		}

	}

	return localctx
}

// IType2Context is an interface to support dynamic dispatch.
type IType2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType2Context differentiates from other interfaces.
	IsType2Context()
}

type Type2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType2Context() *Type2Context {
	var p = new(Type2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_type2
	return p
}

func (*Type2Context) IsType2Context() {}

func NewType2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type2Context {
	var p = new(Type2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_type2

	return p
}

func (s *Type2Context) GetParser() antlr.Parser { return s.parser }

func (s *Type2Context) VALUE() antlr.TerminalNode {
	return s.GetToken(CDDLParserVALUE, 0)
}

func (s *Type2Context) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *Type2Context) GenericArg() IGenericArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericArgContext)
}

func (s *Type2Context) Type_() ITypeContext {
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

func (s *Type2Context) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *Type2Context) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *Type2Context) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *Type2Context) TAG1() antlr.TerminalNode {
	return s.GetToken(CDDLParserTAG1, 0)
}

func (s *Type2Context) TAG2() antlr.TerminalNode {
	return s.GetToken(CDDLParserTAG2, 0)
}

func (s *Type2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterType2(s)
	}
}

func (s *Type2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitType2(s)
	}
}

func (p *CDDLParser) Type2() (localctx IType2Context) {
	this := p
	_ = this

	localctx = NewType2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CDDLParserRULE_type2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(CDDLParserVALUE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(CDDLParserID)
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(155)
				p.GenericArg()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)
			p.Match(CDDLParserT__7)
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(159)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(162)
			p.Type_()
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(163)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(166)
			p.Match(CDDLParserT__8)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(168)
			p.Match(CDDLParserT__9)
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(169)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(172)
			p.Group()
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(173)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(176)
			p.Match(CDDLParserT__10)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(178)
			p.Match(CDDLParserT__11)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(179)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(182)
			p.Group()
		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(183)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(186)
			p.Match(CDDLParserT__12)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(188)
			p.Match(CDDLParserT__13)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(189)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(192)
			p.Match(CDDLParserID)
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(193)
				p.GenericArg()
			}

		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(196)
			p.Match(CDDLParserT__14)
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(197)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(200)
			p.Match(CDDLParserT__7)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(201)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(204)
			p.Group()
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(205)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(208)
			p.Match(CDDLParserT__8)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(210)
			p.Match(CDDLParserT__14)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(211)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(214)
			p.Match(CDDLParserID)
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(215)
				p.GenericArg()
			}

		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(218)
			p.Match(CDDLParserTAG1)
		}
		{
			p.SetState(219)
			p.Match(CDDLParserT__7)
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(220)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(223)
			p.Type_()
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(224)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(227)
			p.Match(CDDLParserT__8)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(229)
			p.Match(CDDLParserTAG2)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(230)
			p.Match(CDDLParserT__15)
		}

	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) AllGroupChoice() []IGroupChoiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupChoiceContext); ok {
			len++
		}
	}

	tst := make([]IGroupChoiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupChoiceContext); ok {
			tst[i] = t.(IGroupChoiceContext)
			i++
		}
	}

	return tst
}

func (s *GroupContext) GroupChoice(i int) IGroupChoiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupChoiceContext); ok {
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

	return t.(IGroupChoiceContext)
}

func (s *GroupContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *GroupContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *CDDLParser) Group() (localctx IGroupContext) {
	this := p
	_ = this

	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CDDLParserRULE_group)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.GroupChoice()
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(234)
					p.Match(CDDLParserS)
				}

			}
			{
				p.SetState(237)
				p.Match(CDDLParserT__16)
			}
			p.SetState(239)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(238)
					p.Match(CDDLParserS)
				}

			}
			{
				p.SetState(241)
				p.GroupChoice()
			}

		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}

	return localctx
}

// IGroupChoiceContext is an interface to support dynamic dispatch.
type IGroupChoiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupChoiceContext differentiates from other interfaces.
	IsGroupChoiceContext()
}

type GroupChoiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupChoiceContext() *GroupChoiceContext {
	var p = new(GroupChoiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_groupChoice
	return p
}

func (*GroupChoiceContext) IsGroupChoiceContext() {}

func NewGroupChoiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupChoiceContext {
	var p = new(GroupChoiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_groupChoice

	return p
}

func (s *GroupChoiceContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupChoiceContext) AllGroupEntry() []IGroupEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupEntryContext); ok {
			len++
		}
	}

	tst := make([]IGroupEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupEntryContext); ok {
			tst[i] = t.(IGroupEntryContext)
			i++
		}
	}

	return tst
}

func (s *GroupChoiceContext) GroupEntry(i int) IGroupEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupEntryContext); ok {
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

	return t.(IGroupEntryContext)
}

func (s *GroupChoiceContext) AllOptComma() []IOptCommaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptCommaContext); ok {
			len++
		}
	}

	tst := make([]IOptCommaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptCommaContext); ok {
			tst[i] = t.(IOptCommaContext)
			i++
		}
	}

	return tst
}

func (s *GroupChoiceContext) OptComma(i int) IOptCommaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptCommaContext); ok {
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

	return t.(IOptCommaContext)
}

func (s *GroupChoiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupChoiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupChoiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterGroupChoice(s)
	}
}

func (s *GroupChoiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitGroupChoice(s)
	}
}

func (p *CDDLParser) GroupChoice() (localctx IGroupChoiceContext) {
	this := p
	_ = this

	localctx = NewGroupChoiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CDDLParserRULE_groupChoice)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&241292544) != 0 {
		{
			p.SetState(247)
			p.GroupEntry()
		}
		{
			p.SetState(248)
			p.OptComma()
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroupEntryContext is an interface to support dynamic dispatch.
type IGroupEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupEntryContext differentiates from other interfaces.
	IsGroupEntryContext()
}

type GroupEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupEntryContext() *GroupEntryContext {
	var p = new(GroupEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_groupEntry
	return p
}

func (*GroupEntryContext) IsGroupEntryContext() {}

func NewGroupEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupEntryContext {
	var p = new(GroupEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_groupEntry

	return p
}

func (s *GroupEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupEntryContext) Type_() ITypeContext {
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

func (s *GroupEntryContext) OCCUR() antlr.TerminalNode {
	return s.GetToken(CDDLParserOCCUR, 0)
}

func (s *GroupEntryContext) MemberKey() IMemberKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberKeyContext)
}

func (s *GroupEntryContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *GroupEntryContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *GroupEntryContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *GroupEntryContext) GenericArg() IGenericArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericArgContext)
}

func (s *GroupEntryContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *GroupEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterGroupEntry(s)
	}
}

func (s *GroupEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitGroupEntry(s)
	}
}

func (p *CDDLParser) GroupEntry() (localctx IGroupEntryContext) {
	this := p
	_ = this

	localctx = NewGroupEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CDDLParserRULE_groupEntry)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(255)
				p.Match(CDDLParserOCCUR)
			}
			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(256)
					p.Match(CDDLParserS)
				}

			}

		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(261)
				p.MemberKey()
			}
			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(262)
					p.Match(CDDLParserS)
				}

			}

		}
		{
			p.SetState(267)
			p.Type_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(268)
				p.Match(CDDLParserOCCUR)
			}
			p.SetState(270)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(269)
					p.Match(CDDLParserS)
				}

			}

		}
		{
			p.SetState(274)
			p.Match(CDDLParserID)
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(275)
				p.GenericArg()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(278)
				p.Match(CDDLParserOCCUR)
			}
			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(279)
					p.Match(CDDLParserS)
				}

			}

		}
		{
			p.SetState(284)
			p.Match(CDDLParserT__7)
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(285)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(288)
			p.Group()
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(289)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(292)
			p.Match(CDDLParserT__8)
		}

	}

	return localctx
}

// IMemberKeyContext is an interface to support dynamic dispatch.
type IMemberKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberKeyContext differentiates from other interfaces.
	IsMemberKeyContext()
}

type MemberKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberKeyContext() *MemberKeyContext {
	var p = new(MemberKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_memberKey
	return p
}

func (*MemberKeyContext) IsMemberKeyContext() {}

func NewMemberKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberKeyContext {
	var p = new(MemberKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_memberKey

	return p
}

func (s *MemberKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberKeyContext) Type1() IType1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType1Context)
}

func (s *MemberKeyContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *MemberKeyContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *MemberKeyContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *MemberKeyContext) VALUE() antlr.TerminalNode {
	return s.GetToken(CDDLParserVALUE, 0)
}

func (s *MemberKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterMemberKey(s)
	}
}

func (s *MemberKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitMemberKey(s)
	}
}

func (p *CDDLParser) MemberKey() (localctx IMemberKeyContext) {
	this := p
	_ = this

	localctx = NewMemberKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CDDLParserRULE_memberKey)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(296)
			p.Type1()
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(297)
				p.Match(CDDLParserS)
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__17 {
			{
				p.SetState(300)
				p.Match(CDDLParserT__17)
			}
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(301)
					p.Match(CDDLParserS)
				}

			}

		}
		{
			p.SetState(306)
			p.Match(CDDLParserT__18)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Match(CDDLParserID)
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(309)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(312)
			p.Match(CDDLParserT__19)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(313)
			p.Match(CDDLParserVALUE)
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(314)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(317)
			p.Match(CDDLParserT__19)
		}

	}

	return localctx
}

// IOptCommaContext is an interface to support dynamic dispatch.
type IOptCommaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptCommaContext differentiates from other interfaces.
	IsOptCommaContext()
}

type OptCommaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptCommaContext() *OptCommaContext {
	var p = new(OptCommaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_optComma
	return p
}

func (*OptCommaContext) IsOptCommaContext() {}

func NewOptCommaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptCommaContext {
	var p = new(OptCommaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_optComma

	return p
}

func (s *OptCommaContext) GetParser() antlr.Parser { return s.parser }

func (s *OptCommaContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *OptCommaContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *OptCommaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptCommaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptCommaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.EnterOptComma(s)
	}
}

func (s *OptCommaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDDLListener); ok {
		listenerT.ExitOptComma(s)
	}
}

func (p *CDDLParser) OptComma() (localctx IOptCommaContext) {
	this := p
	_ = this

	localctx = NewOptCommaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CDDLParserRULE_optComma)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(321)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(320)
			p.Match(CDDLParserS)
		}

	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserT__4 {
		{
			p.SetState(323)
			p.Match(CDDLParserT__4)
		}

	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(326)
			p.Match(CDDLParserS)
		}

	}

	return localctx
}
