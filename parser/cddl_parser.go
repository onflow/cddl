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
		"", "", "", "", "TAG", "MAJOR", "RANGEOP", "CTLOP", "OCCUR", "VALUE",
		"ID", "S",
	}
	staticData.ruleNames = []string{
		"cddl", "rule", "assignType", "assignGroup", "genericParam", "genericArg",
		"type", "type1", "type2", "group", "groupChoice", "groupEntry", "memberKey",
		"optComma",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 324, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 3, 0, 30, 8, 0, 1,
		0, 1, 0, 3, 0, 34, 8, 0, 4, 0, 36, 8, 0, 11, 0, 12, 0, 37, 1, 0, 3, 0,
		41, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 47, 8, 1, 1, 1, 3, 1, 50, 8, 1,
		1, 1, 1, 1, 3, 1, 54, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 60, 8, 1, 1,
		1, 3, 1, 63, 8, 1, 1, 1, 1, 1, 3, 1, 67, 8, 1, 1, 1, 1, 1, 3, 1, 71, 8,
		1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 79, 8, 4, 1, 4, 1, 4, 3, 4,
		83, 8, 4, 1, 4, 1, 4, 3, 4, 87, 8, 4, 1, 4, 1, 4, 3, 4, 91, 8, 4, 5, 4,
		93, 8, 4, 10, 4, 12, 4, 96, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 102, 8,
		5, 1, 5, 1, 5, 3, 5, 106, 8, 5, 1, 5, 1, 5, 3, 5, 110, 8, 5, 1, 5, 1, 5,
		3, 5, 114, 8, 5, 5, 5, 116, 8, 5, 10, 5, 12, 5, 119, 9, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 3, 6, 125, 8, 6, 1, 6, 1, 6, 3, 6, 129, 8, 6, 1, 6, 5, 6, 132,
		8, 6, 10, 6, 12, 6, 135, 9, 6, 1, 7, 1, 7, 3, 7, 139, 8, 7, 1, 7, 1, 7,
		3, 7, 143, 8, 7, 1, 7, 3, 7, 146, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 151, 8,
		8, 1, 8, 1, 8, 3, 8, 155, 8, 8, 1, 8, 1, 8, 3, 8, 159, 8, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 165, 8, 8, 1, 8, 1, 8, 3, 8, 169, 8, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 3, 8, 175, 8, 8, 1, 8, 1, 8, 3, 8, 179, 8, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 185, 8, 8, 1, 8, 1, 8, 3, 8, 189, 8, 8, 1, 8, 1, 8, 3, 8, 193,
		8, 8, 1, 8, 1, 8, 3, 8, 197, 8, 8, 1, 8, 1, 8, 3, 8, 201, 8, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 207, 8, 8, 1, 8, 1, 8, 3, 8, 211, 8, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 216, 8, 8, 1, 8, 1, 8, 3, 8, 220, 8, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 226, 8, 8, 1, 9, 1, 9, 3, 9, 230, 8, 9, 1, 9, 1, 9, 3, 9, 234,
		8, 9, 1, 9, 5, 9, 237, 8, 9, 10, 9, 12, 9, 240, 9, 9, 1, 10, 1, 10, 1,
		10, 5, 10, 245, 8, 10, 10, 10, 12, 10, 248, 9, 10, 1, 11, 3, 11, 251, 8,
		11, 1, 11, 3, 11, 254, 8, 11, 1, 11, 3, 11, 257, 8, 11, 1, 11, 3, 11, 260,
		8, 11, 1, 11, 1, 11, 3, 11, 264, 8, 11, 1, 11, 3, 11, 267, 8, 11, 1, 11,
		1, 11, 3, 11, 271, 8, 11, 1, 11, 3, 11, 274, 8, 11, 1, 11, 3, 11, 277,
		8, 11, 1, 11, 1, 11, 3, 11, 281, 8, 11, 1, 11, 1, 11, 3, 11, 285, 8, 11,
		1, 11, 1, 11, 3, 11, 289, 8, 11, 1, 12, 1, 12, 3, 12, 293, 8, 12, 1, 12,
		3, 12, 296, 8, 12, 1, 12, 3, 12, 299, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 305, 8, 12, 1, 12, 1, 12, 1, 12, 3, 12, 310, 8, 12, 1, 12, 3, 12,
		313, 8, 12, 1, 13, 3, 13, 316, 8, 13, 1, 13, 3, 13, 319, 8, 13, 1, 13,
		3, 13, 322, 8, 13, 1, 13, 0, 0, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 0, 3, 1, 0, 1, 2, 2, 0, 1, 1, 3, 3, 1, 0, 23, 24, 389,
		0, 29, 1, 0, 0, 0, 2, 70, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 74, 1, 0, 0,
		0, 8, 76, 1, 0, 0, 0, 10, 99, 1, 0, 0, 0, 12, 122, 1, 0, 0, 0, 14, 136,
		1, 0, 0, 0, 16, 225, 1, 0, 0, 0, 18, 227, 1, 0, 0, 0, 20, 246, 1, 0, 0,
		0, 22, 288, 1, 0, 0, 0, 24, 312, 1, 0, 0, 0, 26, 315, 1, 0, 0, 0, 28, 30,
		5, 28, 0, 0, 29, 28, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 35, 1, 0, 0, 0,
		31, 33, 3, 2, 1, 0, 32, 34, 5, 28, 0, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1,
		0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 31, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37,
		35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 41, 5, 28,
		0, 0, 40, 39, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43,
		5, 0, 0, 1, 43, 1, 1, 0, 0, 0, 44, 46, 5, 27, 0, 0, 45, 47, 3, 8, 4, 0,
		46, 45, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 50, 5,
		28, 0, 0, 49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51,
		53, 3, 4, 2, 0, 52, 54, 5, 28, 0, 0, 53, 52, 1, 0, 0, 0, 53, 54, 1, 0,
		0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 3, 12, 6, 0, 56, 71, 1, 0, 0, 0, 57,
		59, 5, 27, 0, 0, 58, 60, 3, 8, 4, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0,
		0, 0, 60, 62, 1, 0, 0, 0, 61, 63, 5, 28, 0, 0, 62, 61, 1, 0, 0, 0, 62,
		63, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 66, 3, 6, 3, 0, 65, 67, 5, 28,
		0, 0, 66, 65, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69,
		3, 22, 11, 0, 69, 71, 1, 0, 0, 0, 70, 44, 1, 0, 0, 0, 70, 57, 1, 0, 0,
		0, 71, 3, 1, 0, 0, 0, 72, 73, 7, 0, 0, 0, 73, 5, 1, 0, 0, 0, 74, 75, 7,
		1, 0, 0, 75, 7, 1, 0, 0, 0, 76, 78, 5, 4, 0, 0, 77, 79, 5, 28, 0, 0, 78,
		77, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 5, 27,
		0, 0, 81, 83, 5, 28, 0, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83,
		94, 1, 0, 0, 0, 84, 86, 5, 5, 0, 0, 85, 87, 5, 28, 0, 0, 86, 85, 1, 0,
		0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 5, 27, 0, 0, 89,
		91, 5, 28, 0, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0,
		0, 0, 92, 84, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 98, 5, 6, 0, 0,
		98, 9, 1, 0, 0, 0, 99, 101, 5, 4, 0, 0, 100, 102, 5, 28, 0, 0, 101, 100,
		1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 3, 14,
		7, 0, 104, 106, 5, 28, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0,
		106, 117, 1, 0, 0, 0, 107, 109, 5, 5, 0, 0, 108, 110, 5, 28, 0, 0, 109,
		108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113,
		3, 14, 7, 0, 112, 114, 5, 28, 0, 0, 113, 112, 1, 0, 0, 0, 113, 114, 1,
		0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 107, 1, 0, 0, 0, 116, 119, 1, 0, 0,
		0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119,
		117, 1, 0, 0, 0, 120, 121, 5, 6, 0, 0, 121, 11, 1, 0, 0, 0, 122, 133, 3,
		14, 7, 0, 123, 125, 5, 28, 0, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0,
		0, 0, 125, 126, 1, 0, 0, 0, 126, 128, 5, 7, 0, 0, 127, 129, 5, 28, 0, 0,
		128, 127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130,
		132, 3, 14, 7, 0, 131, 124, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131,
		1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 13, 1, 0, 0, 0, 135, 133, 1, 0,
		0, 0, 136, 145, 3, 16, 8, 0, 137, 139, 5, 28, 0, 0, 138, 137, 1, 0, 0,
		0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 142, 7, 2, 0, 0, 141,
		143, 5, 28, 0, 0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144,
		1, 0, 0, 0, 144, 146, 3, 16, 8, 0, 145, 138, 1, 0, 0, 0, 145, 146, 1, 0,
		0, 0, 146, 15, 1, 0, 0, 0, 147, 226, 5, 26, 0, 0, 148, 150, 5, 27, 0, 0,
		149, 151, 3, 10, 5, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		226, 1, 0, 0, 0, 152, 154, 5, 8, 0, 0, 153, 155, 5, 28, 0, 0, 154, 153,
		1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 158, 3, 12,
		6, 0, 157, 159, 5, 28, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0,
		159, 160, 1, 0, 0, 0, 160, 161, 5, 9, 0, 0, 161, 226, 1, 0, 0, 0, 162,
		164, 5, 10, 0, 0, 163, 165, 5, 28, 0, 0, 164, 163, 1, 0, 0, 0, 164, 165,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 3, 18, 9, 0, 167, 169, 5, 28,
		0, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0,
		170, 171, 5, 11, 0, 0, 171, 226, 1, 0, 0, 0, 172, 174, 5, 12, 0, 0, 173,
		175, 5, 28, 0, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176,
		1, 0, 0, 0, 176, 178, 3, 18, 9, 0, 177, 179, 5, 28, 0, 0, 178, 177, 1,
		0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 5, 13, 0,
		0, 181, 226, 1, 0, 0, 0, 182, 184, 5, 14, 0, 0, 183, 185, 5, 28, 0, 0,
		184, 183, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186,
		188, 5, 27, 0, 0, 187, 189, 3, 10, 5, 0, 188, 187, 1, 0, 0, 0, 188, 189,
		1, 0, 0, 0, 189, 226, 1, 0, 0, 0, 190, 192, 5, 15, 0, 0, 191, 193, 5, 28,
		0, 0, 192, 191, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0,
		194, 196, 5, 8, 0, 0, 195, 197, 5, 28, 0, 0, 196, 195, 1, 0, 0, 0, 196,
		197, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 3, 18, 9, 0, 199, 201,
		5, 28, 0, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0,
		0, 0, 202, 203, 5, 9, 0, 0, 203, 226, 1, 0, 0, 0, 204, 206, 5, 15, 0, 0,
		205, 207, 5, 28, 0, 0, 206, 205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 210, 5, 27, 0, 0, 209, 211, 3, 10, 5, 0, 210, 209,
		1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 226, 1, 0, 0, 0, 212, 213, 5, 21,
		0, 0, 213, 215, 5, 8, 0, 0, 214, 216, 5, 28, 0, 0, 215, 214, 1, 0, 0, 0,
		215, 216, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 3, 12, 6, 0, 218,
		220, 5, 28, 0, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221,
		1, 0, 0, 0, 221, 222, 5, 9, 0, 0, 222, 226, 1, 0, 0, 0, 223, 226, 5, 22,
		0, 0, 224, 226, 5, 16, 0, 0, 225, 147, 1, 0, 0, 0, 225, 148, 1, 0, 0, 0,
		225, 152, 1, 0, 0, 0, 225, 162, 1, 0, 0, 0, 225, 172, 1, 0, 0, 0, 225,
		182, 1, 0, 0, 0, 225, 190, 1, 0, 0, 0, 225, 204, 1, 0, 0, 0, 225, 212,
		1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 224, 1, 0, 0, 0, 226, 17, 1, 0,
		0, 0, 227, 238, 3, 20, 10, 0, 228, 230, 5, 28, 0, 0, 229, 228, 1, 0, 0,
		0, 229, 230, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 5, 17, 0, 0, 232,
		234, 5, 28, 0, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235,
		1, 0, 0, 0, 235, 237, 3, 20, 10, 0, 236, 229, 1, 0, 0, 0, 237, 240, 1,
		0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 19, 1, 0, 0,
		0, 240, 238, 1, 0, 0, 0, 241, 242, 3, 22, 11, 0, 242, 243, 3, 26, 13, 0,
		243, 245, 1, 0, 0, 0, 244, 241, 1, 0, 0, 0, 245, 248, 1, 0, 0, 0, 246,
		244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 21, 1, 0, 0, 0, 248, 246, 1,
		0, 0, 0, 249, 251, 5, 25, 0, 0, 250, 249, 1, 0, 0, 0, 250, 251, 1, 0, 0,
		0, 251, 253, 1, 0, 0, 0, 252, 254, 5, 28, 0, 0, 253, 252, 1, 0, 0, 0, 253,
		254, 1, 0, 0, 0, 254, 256, 1, 0, 0, 0, 255, 257, 3, 24, 12, 0, 256, 255,
		1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 259, 1, 0, 0, 0, 258, 260, 5, 28,
		0, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0,
		261, 289, 3, 12, 6, 0, 262, 264, 5, 25, 0, 0, 263, 262, 1, 0, 0, 0, 263,
		264, 1, 0, 0, 0, 264, 266, 1, 0, 0, 0, 265, 267, 5, 28, 0, 0, 266, 265,
		1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 270, 5, 27,
		0, 0, 269, 271, 3, 10, 5, 0, 270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0,
		271, 289, 1, 0, 0, 0, 272, 274, 5, 25, 0, 0, 273, 272, 1, 0, 0, 0, 273,
		274, 1, 0, 0, 0, 274, 276, 1, 0, 0, 0, 275, 277, 5, 28, 0, 0, 276, 275,
		1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 280, 5, 8,
		0, 0, 279, 281, 5, 28, 0, 0, 280, 279, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0,
		281, 282, 1, 0, 0, 0, 282, 284, 3, 18, 9, 0, 283, 285, 5, 28, 0, 0, 284,
		283, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 287,
		5, 9, 0, 0, 287, 289, 1, 0, 0, 0, 288, 250, 1, 0, 0, 0, 288, 263, 1, 0,
		0, 0, 288, 273, 1, 0, 0, 0, 289, 23, 1, 0, 0, 0, 290, 292, 3, 14, 7, 0,
		291, 293, 5, 28, 0, 0, 292, 291, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293,
		295, 1, 0, 0, 0, 294, 296, 5, 18, 0, 0, 295, 294, 1, 0, 0, 0, 295, 296,
		1, 0, 0, 0, 296, 298, 1, 0, 0, 0, 297, 299, 5, 28, 0, 0, 298, 297, 1, 0,
		0, 0, 298, 299, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 301, 5, 19, 0, 0,
		301, 313, 1, 0, 0, 0, 302, 304, 5, 27, 0, 0, 303, 305, 5, 28, 0, 0, 304,
		303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 313,
		5, 20, 0, 0, 307, 309, 5, 26, 0, 0, 308, 310, 5, 28, 0, 0, 309, 308, 1,
		0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 313, 5, 20, 0,
		0, 312, 290, 1, 0, 0, 0, 312, 302, 1, 0, 0, 0, 312, 307, 1, 0, 0, 0, 313,
		25, 1, 0, 0, 0, 314, 316, 5, 28, 0, 0, 315, 314, 1, 0, 0, 0, 315, 316,
		1, 0, 0, 0, 316, 318, 1, 0, 0, 0, 317, 319, 5, 5, 0, 0, 318, 317, 1, 0,
		0, 0, 318, 319, 1, 0, 0, 0, 319, 321, 1, 0, 0, 0, 320, 322, 5, 28, 0, 0,
		321, 320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 27, 1, 0, 0, 0, 69, 29,
		33, 37, 40, 46, 49, 53, 59, 62, 66, 70, 78, 82, 86, 90, 94, 101, 105, 109,
		113, 117, 124, 128, 133, 138, 142, 145, 150, 154, 158, 164, 168, 174, 178,
		184, 188, 192, 196, 200, 206, 210, 215, 219, 225, 229, 233, 238, 246, 250,
		253, 256, 259, 263, 266, 270, 273, 276, 280, 284, 288, 292, 295, 298, 304,
		309, 312, 315, 318, 321,
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
	CDDLParserTAG     = 21
	CDDLParserMAJOR   = 22
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
	CDDLParserRULE_assignType   = 2
	CDDLParserRULE_assignGroup  = 3
	CDDLParserRULE_genericParam = 4
	CDDLParserRULE_genericArg   = 5
	CDDLParserRULE_type         = 6
	CDDLParserRULE_type1        = 7
	CDDLParserRULE_type2        = 8
	CDDLParserRULE_group        = 9
	CDDLParserRULE_groupChoice  = 10
	CDDLParserRULE_groupEntry   = 11
	CDDLParserRULE_memberKey    = 12
	CDDLParserRULE_optComma     = 13
)

// ICddlContext is an interface to support dynamic dispatch.
type ICddlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_rule returns the _rule rule contexts.
	Get_rule() IRuleContext

	// Set_rule sets the _rule rule contexts.
	Set_rule(IRuleContext)

	// GetRules returns the rules rule context list.
	GetRules() []IRuleContext

	// SetRules sets the rules rule context list.
	SetRules([]IRuleContext)

	// IsCddlContext differentiates from other interfaces.
	IsCddlContext()
}

type CddlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_rule  IRuleContext
	rules  []IRuleContext
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

func (s *CddlContext) Get_rule() IRuleContext { return s._rule }

func (s *CddlContext) Set_rule(v IRuleContext) { s._rule = v }

func (s *CddlContext) GetRules() []IRuleContext { return s.rules }

func (s *CddlContext) SetRules(v []IRuleContext) { s.rules = v }

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

func (s *CddlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitCddl(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(28)
			p.Match(CDDLParserS)
		}

	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CDDLParserID {
		{
			p.SetState(31)

			var _x = p.Rule_()

			localctx.(*CddlContext)._rule = _x
		}
		localctx.(*CddlContext).rules = append(localctx.(*CddlContext).rules, localctx.(*CddlContext)._rule)
		p.SetState(33)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(32)
				p.Match(CDDLParserS)
			}

		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(39)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(42)
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

func (s *RuleContext) CopyFrom(ctx *RuleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupRuleContext struct {
	*RuleContext
}

func NewGroupRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupRuleContext {
	var p = new(GroupRuleContext)

	p.RuleContext = NewEmptyRuleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RuleContext))

	return p
}

func (s *GroupRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

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

func (s *GroupRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroupRule(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeRuleContext struct {
	*RuleContext
}

func NewTypeRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.RuleContext = NewEmptyRuleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RuleContext))

	return p
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

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

func (s *TypeRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitTypeRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Rule_() (localctx IRuleContext) {
	this := p
	_ = this

	localctx = NewRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CDDLParserRULE_rule)
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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Match(CDDLParserID)
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(45)
				p.GenericParam()
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(48)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(51)
			p.AssignType()
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(52)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(55)
			p.Type_()
		}

	case 2:
		localctx = NewGroupRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(CDDLParserID)
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(58)
				p.GenericParam()
			}

		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(61)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(64)
			p.AssignGroup()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(65)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(68)
			p.GroupEntry()
		}

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

func (s *AssignTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitAssignType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) AssignType() (localctx IAssignTypeContext) {
	this := p
	_ = this

	localctx = NewAssignTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CDDLParserRULE_assignType)
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
		p.SetState(72)
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

func (s *AssignGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitAssignGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) AssignGroup() (localctx IAssignGroupContext) {
	this := p
	_ = this

	localctx = NewAssignGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CDDLParserRULE_assignGroup)
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
		p.SetState(74)
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

func (s *GenericParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGenericParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) GenericParam() (localctx IGenericParamContext) {
	this := p
	_ = this

	localctx = NewGenericParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CDDLParserRULE_genericParam)
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
		p.SetState(76)
		p.Match(CDDLParserT__3)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(77)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(80)
		p.Match(CDDLParserID)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(81)
			p.Match(CDDLParserS)
		}

	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDDLParserT__4 {
		{
			p.SetState(84)
			p.Match(CDDLParserT__4)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(85)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(88)
			p.Match(CDDLParserID)
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(89)
				p.Match(CDDLParserS)
			}

		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
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

func (s *GenericArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGenericArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) GenericArg() (localctx IGenericArgContext) {
	this := p
	_ = this

	localctx = NewGenericArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CDDLParserRULE_genericArg)
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
		p.SetState(99)
		p.Match(CDDLParserT__3)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(100)
			p.Match(CDDLParserS)
		}

	}
	{
		p.SetState(103)
		p.Type1()
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserS {
		{
			p.SetState(104)
			p.Match(CDDLParserS)
		}

	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDDLParserT__4 {
		{
			p.SetState(107)
			p.Match(CDDLParserT__4)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(108)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(111)
			p.Type1()
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(112)
				p.Match(CDDLParserS)
			}

		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(120)
		p.Match(CDDLParserT__5)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_type1 returns the _type1 rule contexts.
	Get_type1() IType1Context

	// Set_type1 sets the _type1 rule contexts.
	Set_type1(IType1Context)

	// GetTypes returns the types rule context list.
	GetTypes() []IType1Context

	// SetTypes sets the types rule context list.
	SetTypes([]IType1Context)

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_type1 IType1Context
	types  []IType1Context
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

func (s *TypeContext) Get_type1() IType1Context { return s._type1 }

func (s *TypeContext) Set_type1(v IType1Context) { s._type1 = v }

func (s *TypeContext) GetTypes() []IType1Context { return s.types }

func (s *TypeContext) SetTypes(v []IType1Context) { s.types = v }

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

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CDDLParserRULE_type)
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
		p.SetState(122)

		var _x = p.Type1()

		localctx.(*TypeContext)._type1 = _x
	}
	localctx.(*TypeContext).types = append(localctx.(*TypeContext).types, localctx.(*TypeContext)._type1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(123)
					p.Match(CDDLParserS)
				}

			}
			{
				p.SetState(126)
				p.Match(CDDLParserT__6)
			}
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(127)
					p.Match(CDDLParserS)
				}

			}
			{
				p.SetState(130)

				var _x = p.Type1()

				localctx.(*TypeContext)._type1 = _x
			}
			localctx.(*TypeContext).types = append(localctx.(*TypeContext).types, localctx.(*TypeContext)._type1)

		}
		p.SetState(135)
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

func (s *Type1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitType1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Type1() (localctx IType1Context) {
	this := p
	_ = this

	localctx = NewType1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CDDLParserRULE_type1)
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
		p.SetState(136)
		p.Type2()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(137)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(140)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CDDLParserRANGEOP || _la == CDDLParserCTLOP) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(141)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(144)
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

func (s *Type2Context) CopyFrom(ctx *Type2Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Type2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MapExprContext struct {
	*Type2Context
}

func NewMapExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapExprContext {
	var p = new(MapExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *MapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExprContext) Group() IGroupContext {
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

func (s *MapExprContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *MapExprContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *MapExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitMapExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ChoiceExprContext struct {
	*Type2Context
}

func NewChoiceExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChoiceExprContext {
	var p = new(ChoiceExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *ChoiceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceExprContext) Group() IGroupContext {
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

func (s *ChoiceExprContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *ChoiceExprContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *ChoiceExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitChoiceExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueExprContext struct {
	*Type2Context
	value antlr.Token
}

func NewValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExprContext {
	var p = new(ValueExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *ValueExprContext) GetValue() antlr.Token { return s.value }

func (s *ValueExprContext) SetValue(v antlr.Token) { s.value = v }

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) VALUE() antlr.TerminalNode {
	return s.GetToken(CDDLParserVALUE, 0)
}

func (s *ValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupExprContext struct {
	*Type2Context
}

func NewGroupExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupExprContext {
	var p = new(GroupExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *GroupExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupExprContext) Type_() ITypeContext {
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

func (s *GroupExprContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *GroupExprContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *GroupExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroupExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayExprContext struct {
	*Type2Context
}

func NewArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExprContext {
	var p = new(ArrayExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) Group() IGroupContext {
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

func (s *ArrayExprContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *ArrayExprContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *ArrayExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitArrayExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MajorExprContext struct {
	*Type2Context
	major antlr.Token
}

func NewMajorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MajorExprContext {
	var p = new(MajorExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *MajorExprContext) GetMajor() antlr.Token { return s.major }

func (s *MajorExprContext) SetMajor(v antlr.Token) { s.major = v }

func (s *MajorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MajorExprContext) MAJOR() antlr.TerminalNode {
	return s.GetToken(CDDLParserMAJOR, 0)
}

func (s *MajorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitMajorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	*Type2Context
	id  antlr.Token
	arg IGenericArgContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *IdExprContext) GetId() antlr.Token { return s.id }

func (s *IdExprContext) SetId(v antlr.Token) { s.id = v }

func (s *IdExprContext) GetArg() IGenericArgContext { return s.arg }

func (s *IdExprContext) SetArg(v IGenericArgContext) { s.arg = v }

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *IdExprContext) GenericArg() IGenericArgContext {
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

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AnyExprContext struct {
	*Type2Context
}

func NewAnyExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnyExprContext {
	var p = new(AnyExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *AnyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitAnyExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnwrapExprContext struct {
	*Type2Context
	id  antlr.Token
	arg IGenericArgContext
}

func NewUnwrapExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnwrapExprContext {
	var p = new(UnwrapExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *UnwrapExprContext) GetId() antlr.Token { return s.id }

func (s *UnwrapExprContext) SetId(v antlr.Token) { s.id = v }

func (s *UnwrapExprContext) GetArg() IGenericArgContext { return s.arg }

func (s *UnwrapExprContext) SetArg(v IGenericArgContext) { s.arg = v }

func (s *UnwrapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnwrapExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *UnwrapExprContext) S() antlr.TerminalNode {
	return s.GetToken(CDDLParserS, 0)
}

func (s *UnwrapExprContext) GenericArg() IGenericArgContext {
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

func (s *UnwrapExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitUnwrapExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ChoiceIDExprContext struct {
	*Type2Context
	id  antlr.Token
	arg IGenericArgContext
}

func NewChoiceIDExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChoiceIDExprContext {
	var p = new(ChoiceIDExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *ChoiceIDExprContext) GetId() antlr.Token { return s.id }

func (s *ChoiceIDExprContext) SetId(v antlr.Token) { s.id = v }

func (s *ChoiceIDExprContext) GetArg() IGenericArgContext { return s.arg }

func (s *ChoiceIDExprContext) SetArg(v IGenericArgContext) { s.arg = v }

func (s *ChoiceIDExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceIDExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *ChoiceIDExprContext) S() antlr.TerminalNode {
	return s.GetToken(CDDLParserS, 0)
}

func (s *ChoiceIDExprContext) GenericArg() IGenericArgContext {
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

func (s *ChoiceIDExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitChoiceIDExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TaggedExprContext struct {
	*Type2Context
	tag antlr.Token
}

func NewTaggedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TaggedExprContext {
	var p = new(TaggedExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *TaggedExprContext) GetTag() antlr.Token { return s.tag }

func (s *TaggedExprContext) SetTag(v antlr.Token) { s.tag = v }

func (s *TaggedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaggedExprContext) Type_() ITypeContext {
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

func (s *TaggedExprContext) TAG() antlr.TerminalNode {
	return s.GetToken(CDDLParserTAG, 0)
}

func (s *TaggedExprContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *TaggedExprContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *TaggedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitTaggedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Type2() (localctx IType2Context) {
	this := p
	_ = this

	localctx = NewType2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CDDLParserRULE_type2)
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

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)

			var _m = p.Match(CDDLParserVALUE)

			localctx.(*ValueExprContext).value = _m
		}

	case 2:
		localctx = NewIdExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)

			var _m = p.Match(CDDLParserID)

			localctx.(*IdExprContext).id = _m
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(149)

				var _x = p.GenericArg()

				localctx.(*IdExprContext).arg = _x
			}

		}

	case 3:
		localctx = NewGroupExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
			p.Match(CDDLParserT__7)
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(153)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(156)
			p.Type_()
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(157)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(160)
			p.Match(CDDLParserT__8)
		}

	case 4:
		localctx = NewMapExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.Match(CDDLParserT__9)
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(163)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(166)
			p.Group()
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(167)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(170)
			p.Match(CDDLParserT__10)
		}

	case 5:
		localctx = NewArrayExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(172)
			p.Match(CDDLParserT__11)
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(173)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(176)
			p.Group()
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(177)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(180)
			p.Match(CDDLParserT__12)
		}

	case 6:
		localctx = NewUnwrapExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(182)
			p.Match(CDDLParserT__13)
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

			var _m = p.Match(CDDLParserID)

			localctx.(*UnwrapExprContext).id = _m
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(187)

				var _x = p.GenericArg()

				localctx.(*UnwrapExprContext).arg = _x
			}

		}

	case 7:
		localctx = NewChoiceExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(190)
			p.Match(CDDLParserT__14)
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(191)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(194)
			p.Match(CDDLParserT__7)
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(195)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(198)
			p.Group()
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(199)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(202)
			p.Match(CDDLParserT__8)
		}

	case 8:
		localctx = NewChoiceIDExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(204)
			p.Match(CDDLParserT__14)
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

			var _m = p.Match(CDDLParserID)

			localctx.(*ChoiceIDExprContext).id = _m
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(209)

				var _x = p.GenericArg()

				localctx.(*ChoiceIDExprContext).arg = _x
			}

		}

	case 9:
		localctx = NewTaggedExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(212)

			var _m = p.Match(CDDLParserTAG)

			localctx.(*TaggedExprContext).tag = _m
		}
		{
			p.SetState(213)
			p.Match(CDDLParserT__7)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(214)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(217)
			p.Type_()
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(218)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(221)
			p.Match(CDDLParserT__8)
		}

	case 10:
		localctx = NewMajorExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(223)

			var _m = p.Match(CDDLParserMAJOR)

			localctx.(*MajorExprContext).major = _m
		}

	case 11:
		localctx = NewAnyExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(224)
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

	// Get_groupChoice returns the _groupChoice rule contexts.
	Get_groupChoice() IGroupChoiceContext

	// Set_groupChoice sets the _groupChoice rule contexts.
	Set_groupChoice(IGroupChoiceContext)

	// GetGroups returns the groups rule context list.
	GetGroups() []IGroupChoiceContext

	// SetGroups sets the groups rule context list.
	SetGroups([]IGroupChoiceContext)

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_groupChoice IGroupChoiceContext
	groups       []IGroupChoiceContext
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

func (s *GroupContext) Get_groupChoice() IGroupChoiceContext { return s._groupChoice }

func (s *GroupContext) Set_groupChoice(v IGroupChoiceContext) { s._groupChoice = v }

func (s *GroupContext) GetGroups() []IGroupChoiceContext { return s.groups }

func (s *GroupContext) SetGroups(v []IGroupChoiceContext) { s.groups = v }

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

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Group() (localctx IGroupContext) {
	this := p
	_ = this

	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CDDLParserRULE_group)
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
		p.SetState(227)

		var _x = p.GroupChoice()

		localctx.(*GroupContext)._groupChoice = _x
	}
	localctx.(*GroupContext).groups = append(localctx.(*GroupContext).groups, localctx.(*GroupContext)._groupChoice)
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CDDLParserS {
				{
					p.SetState(228)
					p.Match(CDDLParserS)
				}

			}
			{
				p.SetState(231)
				p.Match(CDDLParserT__16)
			}
			p.SetState(233)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(232)
					p.Match(CDDLParserS)
				}

			}
			{
				p.SetState(235)

				var _x = p.GroupChoice()

				localctx.(*GroupContext)._groupChoice = _x
			}
			localctx.(*GroupContext).groups = append(localctx.(*GroupContext).groups, localctx.(*GroupContext)._groupChoice)

		}
		p.SetState(240)
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

	// GetEntry returns the entry rule contexts.
	GetEntry() IGroupEntryContext

	// SetEntry sets the entry rule contexts.
	SetEntry(IGroupEntryContext)

	// IsGroupChoiceContext differentiates from other interfaces.
	IsGroupChoiceContext()
}

type GroupChoiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	entry  IGroupEntryContext
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

func (s *GroupChoiceContext) GetEntry() IGroupEntryContext { return s.entry }

func (s *GroupChoiceContext) SetEntry(v IGroupEntryContext) { s.entry = v }

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

func (s *GroupChoiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupChoiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupChoiceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroupChoice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) GroupChoice() (localctx IGroupChoiceContext) {
	this := p
	_ = this

	localctx = NewGroupChoiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CDDLParserRULE_groupChoice)

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
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(241)

				var _x = p.GroupEntry()

				localctx.(*GroupChoiceContext).entry = _x
			}
			{
				p.SetState(242)
				p.OptComma()
			}

		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
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

func (s *GroupEntryContext) CopyFrom(ctx *GroupEntryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *GroupEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupsEntryContext struct {
	*GroupEntryContext
}

func NewGroupsEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupsEntryContext {
	var p = new(GroupsEntryContext)

	p.GroupEntryContext = NewEmptyGroupEntryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GroupEntryContext))

	return p
}

func (s *GroupsEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupsEntryContext) Group() IGroupContext {
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

func (s *GroupsEntryContext) OCCUR() antlr.TerminalNode {
	return s.GetToken(CDDLParserOCCUR, 0)
}

func (s *GroupsEntryContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *GroupsEntryContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *GroupsEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroupsEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameEntryContext struct {
	*GroupEntryContext
	id  antlr.Token
	arg IGenericArgContext
}

func NewNameEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameEntryContext {
	var p = new(NameEntryContext)

	p.GroupEntryContext = NewEmptyGroupEntryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GroupEntryContext))

	return p
}

func (s *NameEntryContext) GetId() antlr.Token { return s.id }

func (s *NameEntryContext) SetId(v antlr.Token) { s.id = v }

func (s *NameEntryContext) GetArg() IGenericArgContext { return s.arg }

func (s *NameEntryContext) SetArg(v IGenericArgContext) { s.arg = v }

func (s *NameEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameEntryContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *NameEntryContext) OCCUR() antlr.TerminalNode {
	return s.GetToken(CDDLParserOCCUR, 0)
}

func (s *NameEntryContext) S() antlr.TerminalNode {
	return s.GetToken(CDDLParserS, 0)
}

func (s *NameEntryContext) GenericArg() IGenericArgContext {
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

func (s *NameEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitNameEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeEntryContext struct {
	*GroupEntryContext
}

func NewTypeEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeEntryContext {
	var p = new(TypeEntryContext)

	p.GroupEntryContext = NewEmptyGroupEntryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GroupEntryContext))

	return p
}

func (s *TypeEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeEntryContext) Type_() ITypeContext {
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

func (s *TypeEntryContext) OCCUR() antlr.TerminalNode {
	return s.GetToken(CDDLParserOCCUR, 0)
}

func (s *TypeEntryContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *TypeEntryContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *TypeEntryContext) MemberKey() IMemberKeyContext {
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

func (s *TypeEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitTypeEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) GroupEntry() (localctx IGroupEntryContext) {
	this := p
	_ = this

	localctx = NewGroupEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CDDLParserRULE_groupEntry)
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

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(249)
				p.Match(CDDLParserOCCUR)
			}

		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(252)
				p.Match(CDDLParserS)
			}

		}
		p.SetState(256)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(255)
				p.MemberKey()
			}

		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(258)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(261)
			p.Type_()
		}

	case 2:
		localctx = NewNameEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(262)
				p.Match(CDDLParserOCCUR)
			}

		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(265)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(268)

			var _m = p.Match(CDDLParserID)

			localctx.(*NameEntryContext).id = _m
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(269)

				var _x = p.GenericArg()

				localctx.(*NameEntryContext).arg = _x
			}

		}

	case 3:
		localctx = NewGroupsEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(272)
				p.Match(CDDLParserOCCUR)
			}

		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(275)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(278)
			p.Match(CDDLParserT__7)
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(279)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(282)
			p.Group()
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(283)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(286)
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

func (s *MemberKeyContext) CopyFrom(ctx *MemberKeyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MemberKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeMemberContext struct {
	*MemberKeyContext
	cut antlr.Token
}

func NewTypeMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeMemberContext {
	var p = new(TypeMemberContext)

	p.MemberKeyContext = NewEmptyMemberKeyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MemberKeyContext))

	return p
}

func (s *TypeMemberContext) GetCut() antlr.Token { return s.cut }

func (s *TypeMemberContext) SetCut(v antlr.Token) { s.cut = v }

func (s *TypeMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeMemberContext) Type1() IType1Context {
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

func (s *TypeMemberContext) AllS() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserS)
}

func (s *TypeMemberContext) S(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserS, i)
}

func (s *TypeMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitTypeMember(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueMemberContext struct {
	*MemberKeyContext
	value antlr.Token
}

func NewValueMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueMemberContext {
	var p = new(ValueMemberContext)

	p.MemberKeyContext = NewEmptyMemberKeyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MemberKeyContext))

	return p
}

func (s *ValueMemberContext) GetValue() antlr.Token { return s.value }

func (s *ValueMemberContext) SetValue(v antlr.Token) { s.value = v }

func (s *ValueMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueMemberContext) VALUE() antlr.TerminalNode {
	return s.GetToken(CDDLParserVALUE, 0)
}

func (s *ValueMemberContext) S() antlr.TerminalNode {
	return s.GetToken(CDDLParserS, 0)
}

func (s *ValueMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitValueMember(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameMemberContext struct {
	*MemberKeyContext
	id antlr.Token
}

func NewNameMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameMemberContext {
	var p = new(NameMemberContext)

	p.MemberKeyContext = NewEmptyMemberKeyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MemberKeyContext))

	return p
}

func (s *NameMemberContext) GetId() antlr.Token { return s.id }

func (s *NameMemberContext) SetId(v antlr.Token) { s.id = v }

func (s *NameMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameMemberContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *NameMemberContext) S() antlr.TerminalNode {
	return s.GetToken(CDDLParserS, 0)
}

func (s *NameMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitNameMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) MemberKey() (localctx IMemberKeyContext) {
	this := p
	_ = this

	localctx = NewMemberKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CDDLParserRULE_memberKey)
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

	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeMemberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Type1()
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(291)
				p.Match(CDDLParserS)
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__17 {
			{
				p.SetState(294)

				var _m = p.Match(CDDLParserT__17)

				localctx.(*TypeMemberContext).cut = _m
			}

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
		{
			p.SetState(300)
			p.Match(CDDLParserT__18)
		}

	case 2:
		localctx = NewNameMemberContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)

			var _m = p.Match(CDDLParserID)

			localctx.(*NameMemberContext).id = _m
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(303)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(306)
			p.Match(CDDLParserT__19)
		}

	case 3:
		localctx = NewValueMemberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(307)

			var _m = p.Match(CDDLParserVALUE)

			localctx.(*ValueMemberContext).value = _m
		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserS {
			{
				p.SetState(308)
				p.Match(CDDLParserS)
			}

		}
		{
			p.SetState(311)
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

func (s *OptCommaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitOptComma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) OptComma() (localctx IOptCommaContext) {
	this := p
	_ = this

	localctx = NewOptCommaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CDDLParserRULE_optComma)
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
	p.SetState(315)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(314)
			p.Match(CDDLParserS)
		}

	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserT__4 {
		{
			p.SetState(317)
			p.Match(CDDLParserT__4)
		}

	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(320)
			p.Match(CDDLParserS)
		}

	}

	return localctx
}
