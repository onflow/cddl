// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CDDLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var cddllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cddllexerLexerInit() {
	staticData := &cddllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "TAG", "MAJOR", "RANGEOP", "CTLOP", "OCCUR",
		"UINT", "VALUE", "INT", "NUMBER", "HEXFLOAT", "FRACTION", "EXPONENT",
		"TEXT", "SCHAR", "SESC", "BYTES", "BCHAR", "BSQUAL", "ID", "ALPHA",
		"EALPHA", "DIGIT", "DIGIT1", "HEXDIG", "BINDIG", "S", "WS", "SP", "NL",
		"COMMENT", "PCHAR", "CRLF",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 367, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 155, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		3, 21, 161, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 168, 8, 22,
		1, 23, 1, 23, 1, 23, 1, 24, 3, 24, 174, 8, 24, 1, 24, 1, 24, 3, 24, 178,
		8, 24, 1, 24, 3, 24, 181, 8, 24, 1, 25, 1, 25, 5, 25, 185, 8, 25, 10, 25,
		12, 25, 188, 9, 25, 1, 25, 1, 25, 1, 25, 1, 25, 4, 25, 194, 8, 25, 11,
		25, 12, 25, 195, 1, 25, 1, 25, 1, 25, 1, 25, 4, 25, 202, 8, 25, 11, 25,
		12, 25, 203, 1, 25, 3, 25, 207, 8, 25, 1, 26, 1, 26, 1, 26, 3, 26, 212,
		8, 26, 1, 27, 3, 27, 215, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 3, 28, 223, 8, 28, 1, 28, 1, 28, 3, 28, 227, 8, 28, 3, 28, 229, 8,
		28, 1, 29, 3, 29, 232, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 4, 29, 238, 8,
		29, 11, 29, 12, 29, 239, 1, 29, 1, 29, 4, 29, 244, 8, 29, 11, 29, 12, 29,
		245, 3, 29, 248, 8, 29, 1, 29, 1, 29, 1, 29, 1, 30, 4, 30, 254, 8, 30,
		11, 30, 12, 30, 255, 1, 31, 3, 31, 259, 8, 31, 1, 31, 4, 31, 262, 8, 31,
		11, 31, 12, 31, 263, 1, 32, 1, 32, 5, 32, 268, 8, 32, 10, 32, 12, 32, 271,
		9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33, 277, 8, 33, 1, 34, 1, 34, 1,
		34, 1, 35, 3, 35, 283, 8, 35, 1, 35, 1, 35, 5, 35, 287, 8, 35, 10, 35,
		12, 35, 290, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 3, 36, 297, 8, 36,
		1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 303, 8, 37, 1, 38, 1, 38, 5, 38, 307,
		8, 38, 10, 38, 12, 38, 310, 9, 38, 1, 38, 1, 38, 3, 38, 314, 8, 38, 5,
		38, 316, 8, 38, 10, 38, 12, 38, 319, 9, 38, 1, 39, 1, 39, 1, 40, 1, 40,
		3, 40, 325, 8, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 3, 43, 333,
		8, 43, 1, 44, 1, 44, 1, 45, 4, 45, 338, 8, 45, 11, 45, 12, 45, 339, 1,
		45, 1, 45, 1, 46, 1, 46, 3, 46, 346, 8, 46, 1, 47, 1, 47, 1, 48, 1, 48,
		3, 48, 352, 8, 48, 1, 49, 1, 49, 5, 49, 356, 8, 49, 10, 49, 12, 49, 359,
		9, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 3, 51, 366, 8, 51, 0, 0, 52,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 0, 53, 26, 55, 0, 57, 0, 59,
		0, 61, 0, 63, 0, 65, 0, 67, 0, 69, 0, 71, 0, 73, 0, 75, 0, 77, 27, 79,
		0, 81, 0, 83, 0, 85, 0, 87, 0, 89, 0, 91, 28, 93, 0, 95, 0, 97, 0, 99,
		0, 101, 0, 103, 0, 1, 0, 14, 2, 0, 43, 43, 63, 63, 2, 0, 43, 43, 45, 45,
		1, 0, 34, 34, 4, 0, 32, 33, 35, 91, 93, 126, 128, 1114109, 2, 0, 32, 126,
		128, 1114109, 1, 0, 39, 39, 3, 0, 32, 38, 40, 91, 93, 1114109, 2, 0, 65,
		90, 97, 122, 3, 0, 36, 36, 64, 64, 95, 95, 1, 0, 48, 57, 1, 0, 49, 57,
		1, 0, 48, 49, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 383, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 91, 1, 0,
		0, 0, 1, 105, 1, 0, 0, 0, 3, 107, 1, 0, 0, 0, 5, 110, 1, 0, 0, 0, 7, 114,
		1, 0, 0, 0, 9, 116, 1, 0, 0, 0, 11, 118, 1, 0, 0, 0, 13, 120, 1, 0, 0,
		0, 15, 122, 1, 0, 0, 0, 17, 124, 1, 0, 0, 0, 19, 126, 1, 0, 0, 0, 21, 128,
		1, 0, 0, 0, 23, 130, 1, 0, 0, 0, 25, 132, 1, 0, 0, 0, 27, 134, 1, 0, 0,
		0, 29, 136, 1, 0, 0, 0, 31, 138, 1, 0, 0, 0, 33, 140, 1, 0, 0, 0, 35, 143,
		1, 0, 0, 0, 37, 145, 1, 0, 0, 0, 39, 148, 1, 0, 0, 0, 41, 150, 1, 0, 0,
		0, 43, 156, 1, 0, 0, 0, 45, 167, 1, 0, 0, 0, 47, 169, 1, 0, 0, 0, 49, 180,
		1, 0, 0, 0, 51, 206, 1, 0, 0, 0, 53, 211, 1, 0, 0, 0, 55, 214, 1, 0, 0,
		0, 57, 228, 1, 0, 0, 0, 59, 231, 1, 0, 0, 0, 61, 253, 1, 0, 0, 0, 63, 258,
		1, 0, 0, 0, 65, 265, 1, 0, 0, 0, 67, 276, 1, 0, 0, 0, 69, 278, 1, 0, 0,
		0, 71, 282, 1, 0, 0, 0, 73, 296, 1, 0, 0, 0, 75, 302, 1, 0, 0, 0, 77, 304,
		1, 0, 0, 0, 79, 320, 1, 0, 0, 0, 81, 324, 1, 0, 0, 0, 83, 326, 1, 0, 0,
		0, 85, 328, 1, 0, 0, 0, 87, 332, 1, 0, 0, 0, 89, 334, 1, 0, 0, 0, 91, 337,
		1, 0, 0, 0, 93, 345, 1, 0, 0, 0, 95, 347, 1, 0, 0, 0, 97, 351, 1, 0, 0,
		0, 99, 353, 1, 0, 0, 0, 101, 362, 1, 0, 0, 0, 103, 365, 1, 0, 0, 0, 105,
		106, 5, 61, 0, 0, 106, 2, 1, 0, 0, 0, 107, 108, 5, 47, 0, 0, 108, 109,
		5, 61, 0, 0, 109, 4, 1, 0, 0, 0, 110, 111, 5, 47, 0, 0, 111, 112, 5, 47,
		0, 0, 112, 113, 5, 61, 0, 0, 113, 6, 1, 0, 0, 0, 114, 115, 5, 60, 0, 0,
		115, 8, 1, 0, 0, 0, 116, 117, 5, 44, 0, 0, 117, 10, 1, 0, 0, 0, 118, 119,
		5, 62, 0, 0, 119, 12, 1, 0, 0, 0, 120, 121, 5, 47, 0, 0, 121, 14, 1, 0,
		0, 0, 122, 123, 5, 40, 0, 0, 123, 16, 1, 0, 0, 0, 124, 125, 5, 41, 0, 0,
		125, 18, 1, 0, 0, 0, 126, 127, 5, 123, 0, 0, 127, 20, 1, 0, 0, 0, 128,
		129, 5, 125, 0, 0, 129, 22, 1, 0, 0, 0, 130, 131, 5, 91, 0, 0, 131, 24,
		1, 0, 0, 0, 132, 133, 5, 93, 0, 0, 133, 26, 1, 0, 0, 0, 134, 135, 5, 126,
		0, 0, 135, 28, 1, 0, 0, 0, 136, 137, 5, 38, 0, 0, 137, 30, 1, 0, 0, 0,
		138, 139, 5, 35, 0, 0, 139, 32, 1, 0, 0, 0, 140, 141, 5, 47, 0, 0, 141,
		142, 5, 47, 0, 0, 142, 34, 1, 0, 0, 0, 143, 144, 5, 94, 0, 0, 144, 36,
		1, 0, 0, 0, 145, 146, 5, 61, 0, 0, 146, 147, 5, 62, 0, 0, 147, 38, 1, 0,
		0, 0, 148, 149, 5, 58, 0, 0, 149, 40, 1, 0, 0, 0, 150, 151, 5, 35, 0, 0,
		151, 154, 5, 54, 0, 0, 152, 153, 5, 46, 0, 0, 153, 155, 3, 51, 25, 0, 154,
		152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 42, 1, 0, 0, 0, 156, 157, 5,
		35, 0, 0, 157, 160, 3, 83, 41, 0, 158, 159, 5, 46, 0, 0, 159, 161, 3, 51,
		25, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 44, 1, 0, 0, 0,
		162, 163, 5, 46, 0, 0, 163, 164, 5, 46, 0, 0, 164, 168, 5, 46, 0, 0, 165,
		166, 5, 46, 0, 0, 166, 168, 5, 46, 0, 0, 167, 162, 1, 0, 0, 0, 167, 165,
		1, 0, 0, 0, 168, 46, 1, 0, 0, 0, 169, 170, 5, 46, 0, 0, 170, 171, 3, 77,
		38, 0, 171, 48, 1, 0, 0, 0, 172, 174, 3, 51, 25, 0, 173, 172, 1, 0, 0,
		0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 5, 42, 0, 0, 176,
		178, 3, 51, 25, 0, 177, 176, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 181,
		1, 0, 0, 0, 179, 181, 7, 0, 0, 0, 180, 173, 1, 0, 0, 0, 180, 179, 1, 0,
		0, 0, 181, 50, 1, 0, 0, 0, 182, 186, 3, 85, 42, 0, 183, 185, 3, 83, 41,
		0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186,
		187, 1, 0, 0, 0, 187, 207, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 190,
		5, 48, 0, 0, 190, 191, 5, 120, 0, 0, 191, 193, 1, 0, 0, 0, 192, 194, 3,
		87, 43, 0, 193, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 193, 1, 0,
		0, 0, 195, 196, 1, 0, 0, 0, 196, 207, 1, 0, 0, 0, 197, 198, 5, 48, 0, 0,
		198, 199, 5, 98, 0, 0, 199, 201, 1, 0, 0, 0, 200, 202, 3, 89, 44, 0, 201,
		200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204,
		1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 207, 5, 48, 0, 0, 206, 182, 1, 0,
		0, 0, 206, 189, 1, 0, 0, 0, 206, 197, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0,
		207, 52, 1, 0, 0, 0, 208, 212, 3, 57, 28, 0, 209, 212, 3, 65, 32, 0, 210,
		212, 3, 71, 35, 0, 211, 208, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 210,
		1, 0, 0, 0, 212, 54, 1, 0, 0, 0, 213, 215, 5, 45, 0, 0, 214, 213, 1, 0,
		0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 3, 51, 25,
		0, 217, 56, 1, 0, 0, 0, 218, 229, 3, 59, 29, 0, 219, 222, 3, 55, 27, 0,
		220, 221, 5, 46, 0, 0, 221, 223, 3, 61, 30, 0, 222, 220, 1, 0, 0, 0, 222,
		223, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 225, 5, 101, 0, 0, 225, 227,
		3, 63, 31, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 229, 1,
		0, 0, 0, 228, 218, 1, 0, 0, 0, 228, 219, 1, 0, 0, 0, 229, 58, 1, 0, 0,
		0, 230, 232, 5, 45, 0, 0, 231, 230, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232,
		233, 1, 0, 0, 0, 233, 234, 5, 48, 0, 0, 234, 235, 5, 120, 0, 0, 235, 237,
		1, 0, 0, 0, 236, 238, 3, 87, 43, 0, 237, 236, 1, 0, 0, 0, 238, 239, 1,
		0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 247, 1, 0, 0,
		0, 241, 243, 5, 46, 0, 0, 242, 244, 3, 87, 43, 0, 243, 242, 1, 0, 0, 0,
		244, 245, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246,
		248, 1, 0, 0, 0, 247, 241, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 250, 5, 112, 0, 0, 250, 251, 3, 63, 31, 0, 251, 60, 1,
		0, 0, 0, 252, 254, 3, 83, 41, 0, 253, 252, 1, 0, 0, 0, 254, 255, 1, 0,
		0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 62, 1, 0, 0, 0,
		257, 259, 7, 1, 0, 0, 258, 257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259,
		261, 1, 0, 0, 0, 260, 262, 3, 83, 41, 0, 261, 260, 1, 0, 0, 0, 262, 263,
		1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 64, 1, 0,
		0, 0, 265, 269, 7, 2, 0, 0, 266, 268, 3, 67, 33, 0, 267, 266, 1, 0, 0,
		0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270,
		272, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 273, 7, 2, 0, 0, 273, 66, 1,
		0, 0, 0, 274, 277, 7, 3, 0, 0, 275, 277, 3, 69, 34, 0, 276, 274, 1, 0,
		0, 0, 276, 275, 1, 0, 0, 0, 277, 68, 1, 0, 0, 0, 278, 279, 5, 92, 0, 0,
		279, 280, 7, 4, 0, 0, 280, 70, 1, 0, 0, 0, 281, 283, 3, 75, 37, 0, 282,
		281, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 288,
		7, 5, 0, 0, 285, 287, 3, 73, 36, 0, 286, 285, 1, 0, 0, 0, 287, 290, 1,
		0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 1, 0, 0,
		0, 290, 288, 1, 0, 0, 0, 291, 292, 7, 5, 0, 0, 292, 72, 1, 0, 0, 0, 293,
		297, 7, 6, 0, 0, 294, 297, 3, 69, 34, 0, 295, 297, 3, 103, 51, 0, 296,
		293, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296, 295, 1, 0, 0, 0, 297, 74, 1,
		0, 0, 0, 298, 303, 5, 104, 0, 0, 299, 300, 5, 98, 0, 0, 300, 301, 5, 54,
		0, 0, 301, 303, 5, 52, 0, 0, 302, 298, 1, 0, 0, 0, 302, 299, 1, 0, 0, 0,
		303, 76, 1, 0, 0, 0, 304, 317, 3, 81, 40, 0, 305, 307, 2, 45, 46, 0, 306,
		305, 1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309,
		1, 0, 0, 0, 309, 313, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 314, 3, 81,
		40, 0, 312, 314, 3, 83, 41, 0, 313, 311, 1, 0, 0, 0, 313, 312, 1, 0, 0,
		0, 314, 316, 1, 0, 0, 0, 315, 308, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317,
		315, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 78, 1, 0, 0, 0, 319, 317, 1,
		0, 0, 0, 320, 321, 7, 7, 0, 0, 321, 80, 1, 0, 0, 0, 322, 325, 3, 79, 39,
		0, 323, 325, 7, 8, 0, 0, 324, 322, 1, 0, 0, 0, 324, 323, 1, 0, 0, 0, 325,
		82, 1, 0, 0, 0, 326, 327, 7, 9, 0, 0, 327, 84, 1, 0, 0, 0, 328, 329, 7,
		10, 0, 0, 329, 86, 1, 0, 0, 0, 330, 333, 3, 83, 41, 0, 331, 333, 2, 65,
		70, 0, 332, 330, 1, 0, 0, 0, 332, 331, 1, 0, 0, 0, 333, 88, 1, 0, 0, 0,
		334, 335, 7, 11, 0, 0, 335, 90, 1, 0, 0, 0, 336, 338, 3, 93, 46, 0, 337,
		336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 339, 340,
		1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 342, 6, 45, 0, 0, 342, 92, 1, 0,
		0, 0, 343, 346, 3, 95, 47, 0, 344, 346, 3, 97, 48, 0, 345, 343, 1, 0, 0,
		0, 345, 344, 1, 0, 0, 0, 346, 94, 1, 0, 0, 0, 347, 348, 7, 12, 0, 0, 348,
		96, 1, 0, 0, 0, 349, 352, 3, 99, 49, 0, 350, 352, 3, 103, 51, 0, 351, 349,
		1, 0, 0, 0, 351, 350, 1, 0, 0, 0, 352, 98, 1, 0, 0, 0, 353, 357, 5, 59,
		0, 0, 354, 356, 3, 101, 50, 0, 355, 354, 1, 0, 0, 0, 356, 359, 1, 0, 0,
		0, 357, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 360, 1, 0, 0, 0, 359,
		357, 1, 0, 0, 0, 360, 361, 3, 103, 51, 0, 361, 100, 1, 0, 0, 0, 362, 363,
		7, 4, 0, 0, 363, 102, 1, 0, 0, 0, 364, 366, 7, 13, 0, 0, 365, 364, 1, 0,
		0, 0, 366, 104, 1, 0, 0, 0, 39, 0, 154, 160, 167, 173, 177, 180, 186, 195,
		203, 206, 211, 214, 222, 226, 228, 231, 239, 245, 247, 255, 258, 263, 269,
		276, 282, 288, 296, 302, 308, 313, 317, 324, 332, 339, 345, 351, 357, 365,
		1, 6, 0, 0,
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

// CDDLLexerInit initializes any static state used to implement CDDLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCDDLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CDDLLexerInit() {
	staticData := &cddllexerLexerStaticData
	staticData.once.Do(cddllexerLexerInit)
}

// NewCDDLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCDDLLexer(input antlr.CharStream) *CDDLLexer {
	CDDLLexerInit()
	l := new(CDDLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &cddllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "CDDL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CDDLLexer tokens.
const (
	CDDLLexerT__0    = 1
	CDDLLexerT__1    = 2
	CDDLLexerT__2    = 3
	CDDLLexerT__3    = 4
	CDDLLexerT__4    = 5
	CDDLLexerT__5    = 6
	CDDLLexerT__6    = 7
	CDDLLexerT__7    = 8
	CDDLLexerT__8    = 9
	CDDLLexerT__9    = 10
	CDDLLexerT__10   = 11
	CDDLLexerT__11   = 12
	CDDLLexerT__12   = 13
	CDDLLexerT__13   = 14
	CDDLLexerT__14   = 15
	CDDLLexerT__15   = 16
	CDDLLexerT__16   = 17
	CDDLLexerT__17   = 18
	CDDLLexerT__18   = 19
	CDDLLexerT__19   = 20
	CDDLLexerTAG     = 21
	CDDLLexerMAJOR   = 22
	CDDLLexerRANGEOP = 23
	CDDLLexerCTLOP   = 24
	CDDLLexerOCCUR   = 25
	CDDLLexerVALUE   = 26
	CDDLLexerID      = 27
	CDDLLexerS       = 28
)
