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

type CommentLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var commentlexerLexerStaticData struct {
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

func commentlexerLexerInit() {
	staticData := &commentlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "'null'", "'('",
		"')'", "'{'", "'}'", "'['", "']'", "';'", "','", "'.'", "'='", "'>'",
		"'<'", "'!'", "'~'", "'?'", "':'", "'=='", "'<='", "'>='", "'!='", "'&&'",
		"'||'", "'++'", "'--'", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'^'",
		"'%'", "'+='", "'-='", "'*='", "'/='", "'&='", "'|='", "'^='", "'%='",
		"'<<='", "'>>='", "'>>>='", "", "'->'", "'::'", "'@'", "'...'",
	}
	staticData.symbolicNames = []string{
		"", "COMMENT", "LINE_COMMENT", "PYTHON_COMMENT", "DECIMAL_LITERAL",
		"HEX_LITERAL", "OCT_LITERAL", "BINARY_LITERAL", "FLOAT_LITERAL", "HEX_FLOAT_LITERAL",
		"BOOL_LITERAL", "CHAR_LITERAL", "STRING_LITERAL", "NULL_LITERAL", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT",
		"ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "COLON", "EQUAL",
		"LE", "GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL",
		"DIV", "BITAND", "BITOR", "CARET", "MOD", "ADD_ASSIGN", "SUB_ASSIGN",
		"MUL_ASSIGN", "DIV_ASSIGN", "AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN",
		"MOD_ASSIGN", "LSHIFT_ASSIGN", "RSHIFT_ASSIGN", "URSHIFT_ASSIGN", "TemplateStringLiteral",
		"ARROW", "COLONCOLON", "AT", "ELLIPSIS", "WS", "IDENTIFIER",
	}
	staticData.ruleNames = []string{
		"COMMENT", "LINE_COMMENT", "PYTHON_COMMENT", "DECIMAL_LITERAL", "HEX_LITERAL",
		"OCT_LITERAL", "BINARY_LITERAL", "FLOAT_LITERAL", "HEX_FLOAT_LITERAL",
		"BOOL_LITERAL", "CHAR_LITERAL", "STRING_LITERAL", "NULL_LITERAL", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT",
		"ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "COLON", "EQUAL",
		"LE", "GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL",
		"DIV", "BITAND", "BITOR", "CARET", "MOD", "ADD_ASSIGN", "SUB_ASSIGN",
		"MUL_ASSIGN", "DIV_ASSIGN", "AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN",
		"MOD_ASSIGN", "LSHIFT_ASSIGN", "RSHIFT_ASSIGN", "URSHIFT_ASSIGN", "TemplateStringLiteral",
		"ARROW", "COLONCOLON", "AT", "ELLIPSIS", "WS", "IDENTIFIER", "ExponentPart",
		"EscapeSequence", "HexDigits", "HexDigit", "Digits", "LetterOrDigit",
		"Letter",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 63, 530, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 146, 8, 0, 10,
		0, 12, 0, 149, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 5, 1, 160, 8, 1, 10, 1, 12, 1, 163, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5,
		2, 169, 8, 2, 10, 2, 12, 2, 172, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3,
		3, 179, 8, 3, 1, 3, 4, 3, 182, 8, 3, 11, 3, 12, 3, 183, 1, 3, 3, 3, 187,
		8, 3, 3, 3, 189, 8, 3, 1, 3, 3, 3, 192, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5,
		4, 198, 8, 4, 10, 4, 12, 4, 201, 9, 4, 1, 4, 3, 4, 204, 8, 4, 1, 4, 3,
		4, 207, 8, 4, 1, 5, 1, 5, 5, 5, 211, 8, 5, 10, 5, 12, 5, 214, 9, 5, 1,
		5, 1, 5, 5, 5, 218, 8, 5, 10, 5, 12, 5, 221, 9, 5, 1, 5, 3, 5, 224, 8,
		5, 1, 5, 3, 5, 227, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 233, 8, 6, 10,
		6, 12, 6, 236, 9, 6, 1, 6, 3, 6, 239, 8, 6, 1, 6, 3, 6, 242, 8, 6, 1, 7,
		1, 7, 1, 7, 3, 7, 247, 8, 7, 1, 7, 1, 7, 3, 7, 251, 8, 7, 1, 7, 3, 7, 254,
		8, 7, 1, 7, 3, 7, 257, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 262, 8, 7, 1, 7, 3,
		7, 265, 8, 7, 3, 7, 267, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 273, 8, 8,
		1, 8, 3, 8, 276, 8, 8, 1, 8, 1, 8, 3, 8, 280, 8, 8, 1, 8, 1, 8, 3, 8, 284,
		8, 8, 1, 8, 1, 8, 3, 8, 288, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 299, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10, 304, 8,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 311, 8, 11, 10, 11, 12, 11,
		314, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 49, 1,
		49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52,
		1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 5, 56, 436, 8, 56, 10, 56,
		12, 56, 439, 9, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1,
		58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 60, 1, 61, 4, 61, 456, 8, 61,
		11, 61, 12, 61, 457, 1, 61, 1, 61, 1, 62, 1, 62, 5, 62, 464, 8, 62, 10,
		62, 12, 62, 467, 9, 62, 1, 63, 1, 63, 3, 63, 471, 8, 63, 1, 63, 1, 63,
		1, 64, 1, 64, 1, 64, 1, 64, 3, 64, 479, 8, 64, 1, 64, 3, 64, 482, 8, 64,
		1, 64, 1, 64, 1, 64, 4, 64, 487, 8, 64, 11, 64, 12, 64, 488, 1, 64, 1,
		64, 1, 64, 1, 64, 1, 64, 3, 64, 496, 8, 64, 1, 65, 1, 65, 1, 65, 5, 65,
		501, 8, 65, 10, 65, 12, 65, 504, 9, 65, 1, 65, 3, 65, 507, 8, 65, 1, 66,
		1, 66, 1, 67, 1, 67, 5, 67, 513, 8, 67, 10, 67, 12, 67, 516, 9, 67, 1,
		67, 3, 67, 519, 8, 67, 1, 68, 1, 68, 3, 68, 523, 8, 68, 1, 69, 1, 69, 1,
		69, 1, 69, 3, 69, 529, 8, 69, 1, 147, 0, 70, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50,
		101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58,
		117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 0, 129, 0, 131, 0, 133,
		0, 135, 0, 137, 0, 139, 0, 1, 0, 28, 3, 0, 10, 10, 13, 13, 8232, 8233,
		2, 0, 10, 10, 12, 13, 1, 0, 49, 57, 2, 0, 76, 76, 108, 108, 2, 0, 88, 88,
		120, 120, 3, 0, 48, 57, 65, 70, 97, 102, 4, 0, 48, 57, 65, 70, 95, 95,
		97, 102, 1, 0, 48, 55, 2, 0, 48, 55, 95, 95, 2, 0, 66, 66, 98, 98, 1, 0,
		48, 49, 2, 0, 48, 49, 95, 95, 4, 0, 68, 68, 70, 70, 100, 100, 102, 102,
		2, 0, 80, 80, 112, 112, 2, 0, 43, 43, 45, 45, 4, 0, 10, 10, 13, 13, 39,
		39, 92, 92, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 1, 0, 96, 96, 3, 0, 9,
		10, 12, 13, 32, 32, 2, 0, 69, 69, 101, 101, 8, 0, 34, 34, 39, 39, 92, 92,
		98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 1, 0, 48, 51, 1, 0, 48,
		57, 2, 0, 48, 57, 95, 95, 4, 0, 36, 36, 65, 90, 95, 95, 97, 122, 2, 0,
		0, 127, 55296, 56319, 1, 0, 55296, 56319, 1, 0, 56320, 57343, 574, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1,
		0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101,
		1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0,
		0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1,
		0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0,
		123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 1, 141, 1, 0, 0, 0, 3, 155, 1, 0,
		0, 0, 5, 166, 1, 0, 0, 0, 7, 188, 1, 0, 0, 0, 9, 193, 1, 0, 0, 0, 11, 208,
		1, 0, 0, 0, 13, 228, 1, 0, 0, 0, 15, 266, 1, 0, 0, 0, 17, 268, 1, 0, 0,
		0, 19, 298, 1, 0, 0, 0, 21, 300, 1, 0, 0, 0, 23, 307, 1, 0, 0, 0, 25, 317,
		1, 0, 0, 0, 27, 322, 1, 0, 0, 0, 29, 324, 1, 0, 0, 0, 31, 326, 1, 0, 0,
		0, 33, 328, 1, 0, 0, 0, 35, 330, 1, 0, 0, 0, 37, 332, 1, 0, 0, 0, 39, 334,
		1, 0, 0, 0, 41, 336, 1, 0, 0, 0, 43, 338, 1, 0, 0, 0, 45, 340, 1, 0, 0,
		0, 47, 342, 1, 0, 0, 0, 49, 344, 1, 0, 0, 0, 51, 346, 1, 0, 0, 0, 53, 348,
		1, 0, 0, 0, 55, 350, 1, 0, 0, 0, 57, 352, 1, 0, 0, 0, 59, 354, 1, 0, 0,
		0, 61, 357, 1, 0, 0, 0, 63, 360, 1, 0, 0, 0, 65, 363, 1, 0, 0, 0, 67, 366,
		1, 0, 0, 0, 69, 369, 1, 0, 0, 0, 71, 372, 1, 0, 0, 0, 73, 375, 1, 0, 0,
		0, 75, 378, 1, 0, 0, 0, 77, 380, 1, 0, 0, 0, 79, 382, 1, 0, 0, 0, 81, 384,
		1, 0, 0, 0, 83, 386, 1, 0, 0, 0, 85, 388, 1, 0, 0, 0, 87, 390, 1, 0, 0,
		0, 89, 392, 1, 0, 0, 0, 91, 394, 1, 0, 0, 0, 93, 397, 1, 0, 0, 0, 95, 400,
		1, 0, 0, 0, 97, 403, 1, 0, 0, 0, 99, 406, 1, 0, 0, 0, 101, 409, 1, 0, 0,
		0, 103, 412, 1, 0, 0, 0, 105, 415, 1, 0, 0, 0, 107, 418, 1, 0, 0, 0, 109,
		422, 1, 0, 0, 0, 111, 426, 1, 0, 0, 0, 113, 431, 1, 0, 0, 0, 115, 442,
		1, 0, 0, 0, 117, 445, 1, 0, 0, 0, 119, 448, 1, 0, 0, 0, 121, 450, 1, 0,
		0, 0, 123, 455, 1, 0, 0, 0, 125, 461, 1, 0, 0, 0, 127, 468, 1, 0, 0, 0,
		129, 495, 1, 0, 0, 0, 131, 497, 1, 0, 0, 0, 133, 508, 1, 0, 0, 0, 135,
		510, 1, 0, 0, 0, 137, 522, 1, 0, 0, 0, 139, 528, 1, 0, 0, 0, 141, 142,
		5, 47, 0, 0, 142, 143, 5, 42, 0, 0, 143, 147, 1, 0, 0, 0, 144, 146, 9,
		0, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 148, 1, 0, 0,
		0, 147, 145, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150,
		151, 5, 42, 0, 0, 151, 152, 5, 47, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154,
		6, 0, 0, 0, 154, 2, 1, 0, 0, 0, 155, 156, 5, 47, 0, 0, 156, 157, 5, 47,
		0, 0, 157, 161, 1, 0, 0, 0, 158, 160, 8, 0, 0, 0, 159, 158, 1, 0, 0, 0,
		160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162,
		164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165, 6, 1, 0, 0, 165, 4, 1,
		0, 0, 0, 166, 170, 5, 35, 0, 0, 167, 169, 8, 1, 0, 0, 168, 167, 1, 0, 0,
		0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171,
		173, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 174, 6, 2, 0, 0, 174, 6, 1,
		0, 0, 0, 175, 189, 5, 48, 0, 0, 176, 186, 7, 2, 0, 0, 177, 179, 3, 135,
		67, 0, 178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 187, 1, 0, 0, 0,
		180, 182, 5, 95, 0, 0, 181, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183,
		181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187,
		3, 135, 67, 0, 186, 178, 1, 0, 0, 0, 186, 181, 1, 0, 0, 0, 187, 189, 1,
		0, 0, 0, 188, 175, 1, 0, 0, 0, 188, 176, 1, 0, 0, 0, 189, 191, 1, 0, 0,
		0, 190, 192, 7, 3, 0, 0, 191, 190, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192,
		8, 1, 0, 0, 0, 193, 194, 5, 48, 0, 0, 194, 195, 7, 4, 0, 0, 195, 203, 7,
		5, 0, 0, 196, 198, 7, 6, 0, 0, 197, 196, 1, 0, 0, 0, 198, 201, 1, 0, 0,
		0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201,
		199, 1, 0, 0, 0, 202, 204, 7, 5, 0, 0, 203, 199, 1, 0, 0, 0, 203, 204,
		1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 207, 7, 3, 0, 0, 206, 205, 1, 0,
		0, 0, 206, 207, 1, 0, 0, 0, 207, 10, 1, 0, 0, 0, 208, 212, 5, 48, 0, 0,
		209, 211, 5, 95, 0, 0, 210, 209, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212,
		210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 215, 1, 0, 0, 0, 214, 212,
		1, 0, 0, 0, 215, 223, 7, 7, 0, 0, 216, 218, 7, 8, 0, 0, 217, 216, 1, 0,
		0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0,
		220, 222, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 224, 7, 7, 0, 0, 223,
		219, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 226, 1, 0, 0, 0, 225, 227,
		7, 3, 0, 0, 226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 12, 1, 0,
		0, 0, 228, 229, 5, 48, 0, 0, 229, 230, 7, 9, 0, 0, 230, 238, 7, 10, 0,
		0, 231, 233, 7, 11, 0, 0, 232, 231, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234,
		232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 237, 1, 0, 0, 0, 236, 234,
		1, 0, 0, 0, 237, 239, 7, 10, 0, 0, 238, 234, 1, 0, 0, 0, 238, 239, 1, 0,
		0, 0, 239, 241, 1, 0, 0, 0, 240, 242, 7, 3, 0, 0, 241, 240, 1, 0, 0, 0,
		241, 242, 1, 0, 0, 0, 242, 14, 1, 0, 0, 0, 243, 244, 3, 135, 67, 0, 244,
		246, 5, 46, 0, 0, 245, 247, 3, 135, 67, 0, 246, 245, 1, 0, 0, 0, 246, 247,
		1, 0, 0, 0, 247, 251, 1, 0, 0, 0, 248, 249, 5, 46, 0, 0, 249, 251, 3, 135,
		67, 0, 250, 243, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0,
		252, 254, 3, 127, 63, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254,
		256, 1, 0, 0, 0, 255, 257, 7, 12, 0, 0, 256, 255, 1, 0, 0, 0, 256, 257,
		1, 0, 0, 0, 257, 267, 1, 0, 0, 0, 258, 264, 3, 135, 67, 0, 259, 261, 3,
		127, 63, 0, 260, 262, 7, 12, 0, 0, 261, 260, 1, 0, 0, 0, 261, 262, 1, 0,
		0, 0, 262, 265, 1, 0, 0, 0, 263, 265, 7, 12, 0, 0, 264, 259, 1, 0, 0, 0,
		264, 263, 1, 0, 0, 0, 265, 267, 1, 0, 0, 0, 266, 250, 1, 0, 0, 0, 266,
		258, 1, 0, 0, 0, 267, 16, 1, 0, 0, 0, 268, 269, 5, 48, 0, 0, 269, 279,
		7, 4, 0, 0, 270, 272, 3, 131, 65, 0, 271, 273, 5, 46, 0, 0, 272, 271, 1,
		0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 280, 1, 0, 0, 0, 274, 276, 3, 131,
		65, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0,
		277, 278, 5, 46, 0, 0, 278, 280, 3, 131, 65, 0, 279, 270, 1, 0, 0, 0, 279,
		275, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 283, 7, 13, 0, 0, 282, 284,
		7, 14, 0, 0, 283, 282, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285, 1, 0,
		0, 0, 285, 287, 3, 135, 67, 0, 286, 288, 7, 12, 0, 0, 287, 286, 1, 0, 0,
		0, 287, 288, 1, 0, 0, 0, 288, 18, 1, 0, 0, 0, 289, 290, 5, 116, 0, 0, 290,
		291, 5, 114, 0, 0, 291, 292, 5, 117, 0, 0, 292, 299, 5, 101, 0, 0, 293,
		294, 5, 102, 0, 0, 294, 295, 5, 97, 0, 0, 295, 296, 5, 108, 0, 0, 296,
		297, 5, 115, 0, 0, 297, 299, 5, 101, 0, 0, 298, 289, 1, 0, 0, 0, 298, 293,
		1, 0, 0, 0, 299, 20, 1, 0, 0, 0, 300, 303, 5, 39, 0, 0, 301, 304, 8, 15,
		0, 0, 302, 304, 3, 129, 64, 0, 303, 301, 1, 0, 0, 0, 303, 302, 1, 0, 0,
		0, 304, 305, 1, 0, 0, 0, 305, 306, 5, 39, 0, 0, 306, 22, 1, 0, 0, 0, 307,
		312, 5, 34, 0, 0, 308, 311, 8, 16, 0, 0, 309, 311, 3, 129, 64, 0, 310,
		308, 1, 0, 0, 0, 310, 309, 1, 0, 0, 0, 311, 314, 1, 0, 0, 0, 312, 310,
		1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 315, 1, 0, 0, 0, 314, 312, 1, 0,
		0, 0, 315, 316, 5, 34, 0, 0, 316, 24, 1, 0, 0, 0, 317, 318, 5, 110, 0,
		0, 318, 319, 5, 117, 0, 0, 319, 320, 5, 108, 0, 0, 320, 321, 5, 108, 0,
		0, 321, 26, 1, 0, 0, 0, 322, 323, 5, 40, 0, 0, 323, 28, 1, 0, 0, 0, 324,
		325, 5, 41, 0, 0, 325, 30, 1, 0, 0, 0, 326, 327, 5, 123, 0, 0, 327, 32,
		1, 0, 0, 0, 328, 329, 5, 125, 0, 0, 329, 34, 1, 0, 0, 0, 330, 331, 5, 91,
		0, 0, 331, 36, 1, 0, 0, 0, 332, 333, 5, 93, 0, 0, 333, 38, 1, 0, 0, 0,
		334, 335, 5, 59, 0, 0, 335, 40, 1, 0, 0, 0, 336, 337, 5, 44, 0, 0, 337,
		42, 1, 0, 0, 0, 338, 339, 5, 46, 0, 0, 339, 44, 1, 0, 0, 0, 340, 341, 5,
		61, 0, 0, 341, 46, 1, 0, 0, 0, 342, 343, 5, 62, 0, 0, 343, 48, 1, 0, 0,
		0, 344, 345, 5, 60, 0, 0, 345, 50, 1, 0, 0, 0, 346, 347, 5, 33, 0, 0, 347,
		52, 1, 0, 0, 0, 348, 349, 5, 126, 0, 0, 349, 54, 1, 0, 0, 0, 350, 351,
		5, 63, 0, 0, 351, 56, 1, 0, 0, 0, 352, 353, 5, 58, 0, 0, 353, 58, 1, 0,
		0, 0, 354, 355, 5, 61, 0, 0, 355, 356, 5, 61, 0, 0, 356, 60, 1, 0, 0, 0,
		357, 358, 5, 60, 0, 0, 358, 359, 5, 61, 0, 0, 359, 62, 1, 0, 0, 0, 360,
		361, 5, 62, 0, 0, 361, 362, 5, 61, 0, 0, 362, 64, 1, 0, 0, 0, 363, 364,
		5, 33, 0, 0, 364, 365, 5, 61, 0, 0, 365, 66, 1, 0, 0, 0, 366, 367, 5, 38,
		0, 0, 367, 368, 5, 38, 0, 0, 368, 68, 1, 0, 0, 0, 369, 370, 5, 124, 0,
		0, 370, 371, 5, 124, 0, 0, 371, 70, 1, 0, 0, 0, 372, 373, 5, 43, 0, 0,
		373, 374, 5, 43, 0, 0, 374, 72, 1, 0, 0, 0, 375, 376, 5, 45, 0, 0, 376,
		377, 5, 45, 0, 0, 377, 74, 1, 0, 0, 0, 378, 379, 5, 43, 0, 0, 379, 76,
		1, 0, 0, 0, 380, 381, 5, 45, 0, 0, 381, 78, 1, 0, 0, 0, 382, 383, 5, 42,
		0, 0, 383, 80, 1, 0, 0, 0, 384, 385, 5, 47, 0, 0, 385, 82, 1, 0, 0, 0,
		386, 387, 5, 38, 0, 0, 387, 84, 1, 0, 0, 0, 388, 389, 5, 124, 0, 0, 389,
		86, 1, 0, 0, 0, 390, 391, 5, 94, 0, 0, 391, 88, 1, 0, 0, 0, 392, 393, 5,
		37, 0, 0, 393, 90, 1, 0, 0, 0, 394, 395, 5, 43, 0, 0, 395, 396, 5, 61,
		0, 0, 396, 92, 1, 0, 0, 0, 397, 398, 5, 45, 0, 0, 398, 399, 5, 61, 0, 0,
		399, 94, 1, 0, 0, 0, 400, 401, 5, 42, 0, 0, 401, 402, 5, 61, 0, 0, 402,
		96, 1, 0, 0, 0, 403, 404, 5, 47, 0, 0, 404, 405, 5, 61, 0, 0, 405, 98,
		1, 0, 0, 0, 406, 407, 5, 38, 0, 0, 407, 408, 5, 61, 0, 0, 408, 100, 1,
		0, 0, 0, 409, 410, 5, 124, 0, 0, 410, 411, 5, 61, 0, 0, 411, 102, 1, 0,
		0, 0, 412, 413, 5, 94, 0, 0, 413, 414, 5, 61, 0, 0, 414, 104, 1, 0, 0,
		0, 415, 416, 5, 37, 0, 0, 416, 417, 5, 61, 0, 0, 417, 106, 1, 0, 0, 0,
		418, 419, 5, 60, 0, 0, 419, 420, 5, 60, 0, 0, 420, 421, 5, 61, 0, 0, 421,
		108, 1, 0, 0, 0, 422, 423, 5, 62, 0, 0, 423, 424, 5, 62, 0, 0, 424, 425,
		5, 61, 0, 0, 425, 110, 1, 0, 0, 0, 426, 427, 5, 62, 0, 0, 427, 428, 5,
		62, 0, 0, 428, 429, 5, 62, 0, 0, 429, 430, 5, 61, 0, 0, 430, 112, 1, 0,
		0, 0, 431, 437, 5, 96, 0, 0, 432, 433, 5, 92, 0, 0, 433, 436, 5, 96, 0,
		0, 434, 436, 8, 17, 0, 0, 435, 432, 1, 0, 0, 0, 435, 434, 1, 0, 0, 0, 436,
		439, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 440,
		1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 440, 441, 5, 96, 0, 0, 441, 114, 1, 0,
		0, 0, 442, 443, 5, 45, 0, 0, 443, 444, 5, 62, 0, 0, 444, 116, 1, 0, 0,
		0, 445, 446, 5, 58, 0, 0, 446, 447, 5, 58, 0, 0, 447, 118, 1, 0, 0, 0,
		448, 449, 5, 64, 0, 0, 449, 120, 1, 0, 0, 0, 450, 451, 5, 46, 0, 0, 451,
		452, 5, 46, 0, 0, 452, 453, 5, 46, 0, 0, 453, 122, 1, 0, 0, 0, 454, 456,
		7, 18, 0, 0, 455, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 455, 1, 0,
		0, 0, 457, 458, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 460, 6, 61, 0, 0,
		460, 124, 1, 0, 0, 0, 461, 465, 3, 139, 69, 0, 462, 464, 3, 137, 68, 0,
		463, 462, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0, 465,
		466, 1, 0, 0, 0, 466, 126, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 468, 470,
		7, 19, 0, 0, 469, 471, 7, 14, 0, 0, 470, 469, 1, 0, 0, 0, 470, 471, 1,
		0, 0, 0, 471, 472, 1, 0, 0, 0, 472, 473, 3, 135, 67, 0, 473, 128, 1, 0,
		0, 0, 474, 475, 5, 92, 0, 0, 475, 496, 7, 20, 0, 0, 476, 481, 5, 92, 0,
		0, 477, 479, 7, 21, 0, 0, 478, 477, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479,
		480, 1, 0, 0, 0, 480, 482, 7, 7, 0, 0, 481, 478, 1, 0, 0, 0, 481, 482,
		1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 496, 7, 7, 0, 0, 484, 486, 5, 92,
		0, 0, 485, 487, 5, 117, 0, 0, 486, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0,
		0, 488, 486, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490,
		491, 3, 133, 66, 0, 491, 492, 3, 133, 66, 0, 492, 493, 3, 133, 66, 0, 493,
		494, 3, 133, 66, 0, 494, 496, 1, 0, 0, 0, 495, 474, 1, 0, 0, 0, 495, 476,
		1, 0, 0, 0, 495, 484, 1, 0, 0, 0, 496, 130, 1, 0, 0, 0, 497, 506, 3, 133,
		66, 0, 498, 501, 3, 133, 66, 0, 499, 501, 5, 95, 0, 0, 500, 498, 1, 0,
		0, 0, 500, 499, 1, 0, 0, 0, 501, 504, 1, 0, 0, 0, 502, 500, 1, 0, 0, 0,
		502, 503, 1, 0, 0, 0, 503, 505, 1, 0, 0, 0, 504, 502, 1, 0, 0, 0, 505,
		507, 3, 133, 66, 0, 506, 502, 1, 0, 0, 0, 506, 507, 1, 0, 0, 0, 507, 132,
		1, 0, 0, 0, 508, 509, 7, 5, 0, 0, 509, 134, 1, 0, 0, 0, 510, 518, 7, 22,
		0, 0, 511, 513, 7, 23, 0, 0, 512, 511, 1, 0, 0, 0, 513, 516, 1, 0, 0, 0,
		514, 512, 1, 0, 0, 0, 514, 515, 1, 0, 0, 0, 515, 517, 1, 0, 0, 0, 516,
		514, 1, 0, 0, 0, 517, 519, 7, 22, 0, 0, 518, 514, 1, 0, 0, 0, 518, 519,
		1, 0, 0, 0, 519, 136, 1, 0, 0, 0, 520, 523, 3, 139, 69, 0, 521, 523, 7,
		22, 0, 0, 522, 520, 1, 0, 0, 0, 522, 521, 1, 0, 0, 0, 523, 138, 1, 0, 0,
		0, 524, 529, 7, 24, 0, 0, 525, 529, 8, 25, 0, 0, 526, 527, 7, 26, 0, 0,
		527, 529, 7, 27, 0, 0, 528, 524, 1, 0, 0, 0, 528, 525, 1, 0, 0, 0, 528,
		526, 1, 0, 0, 0, 529, 140, 1, 0, 0, 0, 51, 0, 147, 161, 170, 178, 183,
		186, 188, 191, 199, 203, 206, 212, 219, 223, 226, 234, 238, 241, 246, 250,
		253, 256, 261, 264, 266, 272, 275, 279, 283, 287, 298, 303, 310, 312, 435,
		437, 457, 465, 470, 478, 481, 488, 495, 500, 502, 506, 514, 518, 522, 528,
		1, 0, 1, 0,
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

// CommentLexerInit initializes any static state used to implement CommentLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCommentLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CommentLexerInit() {
	staticData := &commentlexerLexerStaticData
	staticData.once.Do(commentlexerLexerInit)
}

// NewCommentLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCommentLexer(input antlr.CharStream) *CommentLexer {
	CommentLexerInit()
	l := new(CommentLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &commentlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "CommentLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CommentLexer tokens.
const (
	CommentLexerCOMMENT               = 1
	CommentLexerLINE_COMMENT          = 2
	CommentLexerPYTHON_COMMENT        = 3
	CommentLexerDECIMAL_LITERAL       = 4
	CommentLexerHEX_LITERAL           = 5
	CommentLexerOCT_LITERAL           = 6
	CommentLexerBINARY_LITERAL        = 7
	CommentLexerFLOAT_LITERAL         = 8
	CommentLexerHEX_FLOAT_LITERAL     = 9
	CommentLexerBOOL_LITERAL          = 10
	CommentLexerCHAR_LITERAL          = 11
	CommentLexerSTRING_LITERAL        = 12
	CommentLexerNULL_LITERAL          = 13
	CommentLexerLPAREN                = 14
	CommentLexerRPAREN                = 15
	CommentLexerLBRACE                = 16
	CommentLexerRBRACE                = 17
	CommentLexerLBRACK                = 18
	CommentLexerRBRACK                = 19
	CommentLexerSEMI                  = 20
	CommentLexerCOMMA                 = 21
	CommentLexerDOT                   = 22
	CommentLexerASSIGN                = 23
	CommentLexerGT                    = 24
	CommentLexerLT                    = 25
	CommentLexerBANG                  = 26
	CommentLexerTILDE                 = 27
	CommentLexerQUESTION              = 28
	CommentLexerCOLON                 = 29
	CommentLexerEQUAL                 = 30
	CommentLexerLE                    = 31
	CommentLexerGE                    = 32
	CommentLexerNOTEQUAL              = 33
	CommentLexerAND                   = 34
	CommentLexerOR                    = 35
	CommentLexerINC                   = 36
	CommentLexerDEC                   = 37
	CommentLexerADD                   = 38
	CommentLexerSUB                   = 39
	CommentLexerMUL                   = 40
	CommentLexerDIV                   = 41
	CommentLexerBITAND                = 42
	CommentLexerBITOR                 = 43
	CommentLexerCARET                 = 44
	CommentLexerMOD                   = 45
	CommentLexerADD_ASSIGN            = 46
	CommentLexerSUB_ASSIGN            = 47
	CommentLexerMUL_ASSIGN            = 48
	CommentLexerDIV_ASSIGN            = 49
	CommentLexerAND_ASSIGN            = 50
	CommentLexerOR_ASSIGN             = 51
	CommentLexerXOR_ASSIGN            = 52
	CommentLexerMOD_ASSIGN            = 53
	CommentLexerLSHIFT_ASSIGN         = 54
	CommentLexerRSHIFT_ASSIGN         = 55
	CommentLexerURSHIFT_ASSIGN        = 56
	CommentLexerTemplateStringLiteral = 57
	CommentLexerARROW                 = 58
	CommentLexerCOLONCOLON            = 59
	CommentLexerAT                    = 60
	CommentLexerELLIPSIS              = 61
	CommentLexerWS                    = 62
	CommentLexerIDENTIFIER            = 63
)