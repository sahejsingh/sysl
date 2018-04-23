// Generated from SyslLexer.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "encoding/json"
import "strings"

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 68, 560,
	8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5,
	9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4,
	11, 9, 11, 4, 12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16,
	9, 16, 4, 17, 9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9,
	21, 4, 22, 9, 22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26,
	4, 27, 9, 27, 4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4,
	32, 9, 32, 4, 33, 9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37,
	9, 37, 4, 38, 9, 38, 4, 39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9,
	42, 4, 43, 9, 43, 4, 44, 9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47,
	4, 48, 9, 48, 4, 49, 9, 49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4,
	53, 9, 53, 4, 54, 9, 54, 4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58,
	9, 58, 4, 59, 9, 59, 4, 60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9,
	63, 4, 64, 9, 64, 4, 65, 9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68,
	4, 69, 9, 69, 4, 70, 9, 70, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 5, 2, 178, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 201, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 6, 8, 232, 10, 8, 13,
	8, 14, 8, 233, 3, 9, 3, 9, 6, 9, 238, 10, 9, 13, 9, 14, 9, 239, 3, 9, 3,
	9, 3, 9, 6, 9, 245, 10, 9, 13, 9, 14, 9, 246, 5, 9, 249, 10, 9, 3, 9, 7,
	9, 252, 10, 9, 12, 9, 14, 9, 255, 11, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 5, 47, 395, 10, 47, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 6, 52, 416, 10,
	52, 13, 52, 14, 52, 417, 3, 52, 3, 52, 3, 52, 3, 53, 6, 53, 424, 10, 53,
	13, 53, 14, 53, 425, 3, 53, 3, 53, 6, 53, 430, 10, 53, 13, 53, 14, 53,
	431, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 7, 54, 439, 10, 54, 12, 54, 14,
	54, 442, 11, 54, 3, 55, 7, 55, 445, 10, 55, 12, 55, 14, 55, 448, 11, 55,
	3, 56, 7, 56, 451, 10, 56, 12, 56, 14, 56, 454, 11, 56, 3, 57, 3, 57, 3,
	57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 464, 10, 57, 3, 57, 3, 57,
	3, 58, 5, 58, 469, 10, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 59, 3,
	59, 3, 59, 3, 59, 3, 59, 3, 60, 6, 60, 482, 10, 60, 13, 60, 14, 60, 483,
	3, 61, 3, 61, 3, 61, 6, 61, 489, 10, 61, 13, 61, 14, 61, 490, 3, 61, 6,
	61, 494, 10, 61, 13, 61, 14, 61, 495, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62,
	3, 62, 7, 62, 504, 10, 62, 12, 62, 14, 62, 507, 11, 62, 3, 63, 6, 63, 510,
	10, 63, 13, 63, 14, 63, 511, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 6, 64,
	519, 10, 64, 13, 64, 14, 64, 520, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3,
	65, 3, 66, 3, 66, 6, 66, 531, 10, 66, 13, 66, 14, 66, 532, 3, 66, 3, 66,
	3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 7, 68, 544, 10, 68, 12,
	68, 14, 68, 547, 11, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 6, 70,
	555, 10, 70, 13, 70, 14, 70, 556, 3, 70, 3, 70, 3, 431, 2, 71, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 2, 19, 2, 21, 10, 23, 11, 25, 12, 27, 13, 29,
	14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22, 47,
	23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 61, 30, 63, 31, 65,
	32, 67, 33, 69, 34, 71, 35, 73, 36, 75, 37, 77, 38, 79, 39, 81, 40, 83,
	41, 85, 42, 87, 43, 89, 44, 91, 45, 93, 46, 95, 47, 97, 48, 99, 49, 101,
	50, 103, 51, 105, 52, 107, 53, 109, 54, 111, 55, 113, 2, 115, 2, 117, 56,
	119, 57, 121, 58, 123, 2, 125, 59, 127, 60, 129, 61, 131, 62, 133, 63,
	135, 64, 137, 65, 139, 66, 141, 67, 143, 68, 7, 2, 3, 4, 5, 6, 30, 8, 2,
	11, 12, 15, 15, 34, 34, 49, 49, 60, 60, 94, 94, 4, 2, 11, 11, 34, 34, 4,
	2, 84, 84, 116, 116, 4, 2, 71, 71, 103, 103, 4, 2, 86, 86, 118, 118, 4,
	2, 87, 87, 119, 119, 4, 2, 80, 80, 112, 112, 4, 2, 75, 75, 107, 107, 4,
	2, 72, 72, 104, 104, 4, 2, 78, 78, 110, 110, 4, 2, 85, 85, 117, 117, 4,
	2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 3, 2, 36, 36, 3, 2, 41, 41,
	4, 2, 12, 12, 15, 15, 3, 2, 12, 12, 3, 2, 50, 59, 5, 2, 12, 12, 15, 15,
	36, 36, 5, 2, 12, 12, 15, 15, 41, 41, 14, 2, 12, 12, 15, 15, 34, 37, 40,
	43, 47, 47, 49, 49, 60, 60, 62, 63, 65, 66, 93, 93, 95, 95, 125, 127, 5,
	2, 34, 34, 40, 40, 47, 47, 4, 2, 67, 92, 99, 124, 7, 2, 47, 47, 50, 59,
	67, 92, 97, 97, 99, 124, 3, 2, 34, 34, 8, 2, 12, 12, 15, 15, 60, 60, 62,
	62, 93, 93, 95, 95, 7, 2, 47, 48, 50, 59, 67, 92, 97, 97, 99, 124, 7, 2,
	12, 12, 15, 15, 60, 60, 93, 93, 95, 95, 2, 582, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2,
	105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2,
	2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 125,
	3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 3, 131, 3, 2, 2, 2,
	4, 133, 3, 2, 2, 2, 4, 135, 3, 2, 2, 2, 5, 137, 3, 2, 2, 2, 5, 139, 3,
	2, 2, 2, 6, 141, 3, 2, 2, 2, 6, 143, 3, 2, 2, 2, 7, 177, 3, 2, 2, 2, 9,
	200, 3, 2, 2, 2, 11, 204, 3, 2, 2, 2, 13, 210, 3, 2, 2, 2, 15, 217, 3,
	2, 2, 2, 17, 223, 3, 2, 2, 2, 19, 231, 3, 2, 2, 2, 21, 235, 3, 2, 2, 2,
	23, 259, 3, 2, 2, 2, 25, 268, 3, 2, 2, 2, 27, 273, 3, 2, 2, 2, 29, 278,
	3, 2, 2, 2, 31, 284, 3, 2, 2, 2, 33, 289, 3, 2, 2, 2, 35, 293, 3, 2, 2,
	2, 37, 296, 3, 2, 2, 2, 39, 303, 3, 2, 2, 2, 41, 310, 3, 2, 2, 2, 43, 314,
	3, 2, 2, 2, 45, 320, 3, 2, 2, 2, 47, 325, 3, 2, 2, 2, 49, 328, 3, 2, 2,
	2, 51, 333, 3, 2, 2, 2, 53, 336, 3, 2, 2, 2, 55, 346, 3, 2, 2, 2, 57, 348,
	3, 2, 2, 2, 59, 350, 3, 2, 2, 2, 61, 352, 3, 2, 2, 2, 63, 354, 3, 2, 2,
	2, 65, 356, 3, 2, 2, 2, 67, 358, 3, 2, 2, 2, 69, 360, 3, 2, 2, 2, 71, 362,
	3, 2, 2, 2, 73, 364, 3, 2, 2, 2, 75, 366, 3, 2, 2, 2, 77, 368, 3, 2, 2,
	2, 79, 370, 3, 2, 2, 2, 81, 372, 3, 2, 2, 2, 83, 376, 3, 2, 2, 2, 85, 378,
	3, 2, 2, 2, 87, 381, 3, 2, 2, 2, 89, 384, 3, 2, 2, 2, 91, 386, 3, 2, 2,
	2, 93, 388, 3, 2, 2, 2, 95, 390, 3, 2, 2, 2, 97, 392, 3, 2, 2, 2, 99, 402,
	3, 2, 2, 2, 101, 406, 3, 2, 2, 2, 103, 410, 3, 2, 2, 2, 105, 412, 3, 2,
	2, 2, 107, 415, 3, 2, 2, 2, 109, 423, 3, 2, 2, 2, 111, 436, 3, 2, 2, 2,
	113, 446, 3, 2, 2, 2, 115, 452, 3, 2, 2, 2, 117, 463, 3, 2, 2, 2, 119,
	468, 3, 2, 2, 2, 121, 475, 3, 2, 2, 2, 123, 481, 3, 2, 2, 2, 125, 485,
	3, 2, 2, 2, 127, 501, 3, 2, 2, 2, 129, 509, 3, 2, 2, 2, 131, 518, 3, 2,
	2, 2, 133, 524, 3, 2, 2, 2, 135, 528, 3, 2, 2, 2, 137, 536, 3, 2, 2, 2,
	139, 541, 3, 2, 2, 2, 141, 548, 3, 2, 2, 2, 143, 552, 3, 2, 2, 2, 145,
	146, 7, 107, 2, 2, 146, 147, 7, 112, 2, 2, 147, 178, 7, 118, 2, 2, 148,
	149, 7, 117, 2, 2, 149, 150, 7, 118, 2, 2, 150, 151, 7, 116, 2, 2, 151,
	152, 7, 107, 2, 2, 152, 153, 7, 112, 2, 2, 153, 178, 7, 105, 2, 2, 154,
	155, 7, 102, 2, 2, 155, 156, 7, 99, 2, 2, 156, 157, 7, 118, 2, 2, 157,
	178, 7, 103, 2, 2, 158, 159, 7, 100, 2, 2, 159, 160, 7, 113, 2, 2, 160,
	161, 7, 113, 2, 2, 161, 178, 7, 110, 2, 2, 162, 163, 7, 102, 2, 2, 163,
	164, 7, 103, 2, 2, 164, 165, 7, 101, 2, 2, 165, 166, 7, 107, 2, 2, 166,
	167, 7, 111, 2, 2, 167, 168, 7, 99, 2, 2, 168, 178, 7, 110, 2, 2, 169,
	170, 7, 102, 2, 2, 170, 171, 7, 99, 2, 2, 171, 172, 7, 118, 2, 2, 172,
	173, 7, 103, 2, 2, 173, 174, 7, 118, 2, 2, 174, 175, 7, 107, 2, 2, 175,
	176, 7, 111, 2, 2, 176, 178, 7, 103, 2, 2, 177, 145, 3, 2, 2, 2, 177, 148,
	3, 2, 2, 2, 177, 154, 3, 2, 2, 2, 177, 158, 3, 2, 2, 2, 177, 162, 3, 2,
	2, 2, 177, 169, 3, 2, 2, 2, 178, 8, 3, 2, 2, 2, 179, 180, 7, 73, 2, 2,
	180, 181, 7, 71, 2, 2, 181, 201, 7, 86, 2, 2, 182, 183, 7, 82, 2, 2, 183,
	184, 7, 81, 2, 2, 184, 185, 7, 85, 2, 2, 185, 201, 7, 86, 2, 2, 186, 187,
	7, 70, 2, 2, 187, 188, 7, 71, 2, 2, 188, 189, 7, 78, 2, 2, 189, 190, 7,
	71, 2, 2, 190, 191, 7, 86, 2, 2, 191, 201, 7, 71, 2, 2, 192, 193, 7, 82,
	2, 2, 193, 194, 7, 87, 2, 2, 194, 201, 7, 86, 2, 2, 195, 196, 7, 82, 2,
	2, 196, 197, 7, 67, 2, 2, 197, 198, 7, 86, 2, 2, 198, 199, 7, 69, 2, 2,
	199, 201, 7, 74, 2, 2, 200, 179, 3, 2, 2, 2, 200, 182, 3, 2, 2, 2, 200,
	186, 3, 2, 2, 2, 200, 192, 3, 2, 2, 2, 200, 195, 3, 2, 2, 2, 201, 202,
	3, 2, 2, 2, 202, 203, 8, 3, 2, 2, 203, 10, 3, 2, 2, 2, 204, 205, 7, 35,
	2, 2, 205, 206, 7, 121, 2, 2, 206, 207, 7, 116, 2, 2, 207, 208, 7, 99,
	2, 2, 208, 209, 7, 114, 2, 2, 209, 12, 3, 2, 2, 2, 210, 211, 7, 35, 2,
	2, 211, 212, 7, 118, 2, 2, 212, 213, 7, 99, 2, 2, 213, 214, 7, 100, 2,
	2, 214, 215, 7, 110, 2, 2, 215, 216, 7, 103, 2, 2, 216, 14, 3, 2, 2, 2,
	217, 218, 7, 35, 2, 2, 218, 219, 7, 118, 2, 2, 219, 220, 7, 123, 2, 2,
	220, 221, 7, 114, 2, 2, 221, 222, 7, 103, 2, 2, 222, 16, 3, 2, 2, 2, 223,
	224, 7, 107, 2, 2, 224, 225, 7, 111, 2, 2, 225, 226, 7, 114, 2, 2, 226,
	227, 7, 113, 2, 2, 227, 228, 7, 116, 2, 2, 228, 229, 7, 118, 2, 2, 229,
	18, 3, 2, 2, 2, 230, 232, 10, 2, 2, 2, 231, 230, 3, 2, 2, 2, 232, 233,
	3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 20, 3, 2,
	2, 2, 235, 237, 5, 17, 7, 2, 236, 238, 7, 34, 2, 2, 237, 236, 3, 2, 2,
	2, 238, 239, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240,
	248, 3, 2, 2, 2, 241, 249, 5, 19, 8, 2, 242, 243, 7, 49, 2, 2, 243, 245,
	5, 19, 8, 2, 244, 242, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 244, 3, 2,
	2, 2, 246, 247, 3, 2, 2, 2, 247, 249, 3, 2, 2, 2, 248, 241, 3, 2, 2, 2,
	248, 244, 3, 2, 2, 2, 249, 253, 3, 2, 2, 2, 250, 252, 9, 3, 2, 2, 251,
	250, 3, 2, 2, 2, 252, 255, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254,
	3, 2, 2, 2, 254, 256, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 256, 257, 5, 119,
	58, 2, 257, 258, 8, 9, 3, 2, 258, 22, 3, 2, 2, 2, 259, 260, 9, 4, 2, 2,
	260, 261, 9, 5, 2, 2, 261, 262, 9, 6, 2, 2, 262, 263, 9, 7, 2, 2, 263,
	264, 9, 4, 2, 2, 264, 265, 9, 8, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267,
	8, 10, 4, 2, 267, 24, 3, 2, 2, 2, 268, 269, 9, 9, 2, 2, 269, 270, 9, 10,
	2, 2, 270, 271, 3, 2, 2, 2, 271, 272, 8, 11, 5, 2, 272, 26, 3, 2, 2, 2,
	273, 274, 9, 5, 2, 2, 274, 275, 9, 11, 2, 2, 275, 276, 9, 12, 2, 2, 276,
	277, 9, 5, 2, 2, 277, 28, 3, 2, 2, 2, 278, 279, 9, 10, 2, 2, 279, 280,
	9, 13, 2, 2, 280, 281, 9, 4, 2, 2, 281, 282, 3, 2, 2, 2, 282, 283, 8, 13,
	5, 2, 283, 30, 3, 2, 2, 2, 284, 285, 9, 11, 2, 2, 285, 286, 9, 13, 2, 2,
	286, 287, 9, 13, 2, 2, 287, 288, 9, 14, 2, 2, 288, 32, 3, 2, 2, 2, 289,
	290, 7, 48, 2, 2, 290, 291, 7, 48, 2, 2, 291, 292, 7, 48, 2, 2, 292, 34,
	3, 2, 2, 2, 293, 294, 7, 48, 2, 2, 294, 295, 7, 48, 2, 2, 295, 36, 3, 2,
	2, 2, 296, 297, 7, 117, 2, 2, 297, 298, 7, 103, 2, 2, 298, 299, 7, 118,
	2, 2, 299, 300, 7, 34, 2, 2, 300, 301, 7, 113, 2, 2, 301, 302, 7, 104,
	2, 2, 302, 38, 3, 2, 2, 2, 303, 304, 9, 13, 2, 2, 304, 305, 7, 112, 2,
	2, 305, 306, 7, 103, 2, 2, 306, 307, 7, 34, 2, 2, 307, 308, 7, 113, 2,
	2, 308, 309, 7, 104, 2, 2, 309, 40, 3, 2, 2, 2, 310, 311, 7, 47, 2, 2,
	311, 312, 7, 126, 2, 2, 312, 313, 7, 64, 2, 2, 313, 42, 3, 2, 2, 2, 314,
	315, 7, 62, 2, 2, 315, 316, 7, 47, 2, 2, 316, 317, 7, 64, 2, 2, 317, 318,
	3, 2, 2, 2, 318, 319, 8, 20, 6, 2, 319, 44, 3, 2, 2, 2, 320, 321, 7, 60,
	2, 2, 321, 322, 7, 60, 2, 2, 322, 323, 3, 2, 2, 2, 323, 324, 8, 21, 5,
	2, 324, 46, 3, 2, 2, 2, 325, 326, 7, 62, 2, 2, 326, 327, 7, 60, 2, 2, 327,
	48, 3, 2, 2, 2, 328, 329, 7, 62, 2, 2, 329, 330, 7, 47, 2, 2, 330, 331,
	3, 2, 2, 2, 331, 332, 8, 23, 5, 2, 332, 50, 3, 2, 2, 2, 333, 334, 7, 47,
	2, 2, 334, 335, 7, 64, 2, 2, 335, 52, 3, 2, 2, 2, 336, 337, 7, 48, 2, 2,
	337, 338, 7, 48, 2, 2, 338, 339, 7, 34, 2, 2, 339, 340, 7, 44, 2, 2, 340,
	341, 7, 34, 2, 2, 341, 342, 7, 62, 2, 2, 342, 343, 7, 47, 2, 2, 343, 344,
	7, 34, 2, 2, 344, 345, 7, 44, 2, 2, 345, 54, 3, 2, 2, 2, 346, 347, 7, 45,
	2, 2, 347, 56, 3, 2, 2, 2, 348, 349, 7, 128, 2, 2, 349, 58, 3, 2, 2, 2,
	350, 351, 7, 46, 2, 2, 351, 60, 3, 2, 2, 2, 352, 353, 7, 63, 2, 2, 353,
	62, 3, 2, 2, 2, 354, 355, 7, 38, 2, 2, 355, 64, 3, 2, 2, 2, 356, 357, 7,
	49, 2, 2, 357, 66, 3, 2, 2, 2, 358, 359, 7, 47, 2, 2, 359, 68, 3, 2, 2,
	2, 360, 361, 7, 44, 2, 2, 361, 70, 3, 2, 2, 2, 362, 363, 7, 60, 2, 2, 363,
	72, 3, 2, 2, 2, 364, 365, 7, 39, 2, 2, 365, 74, 3, 2, 2, 2, 366, 367, 7,
	48, 2, 2, 367, 76, 3, 2, 2, 2, 368, 369, 7, 35, 2, 2, 369, 78, 3, 2, 2,
	2, 370, 371, 7, 65, 2, 2, 371, 80, 3, 2, 2, 2, 372, 373, 7, 66, 2, 2, 373,
	374, 3, 2, 2, 2, 374, 375, 8, 39, 7, 2, 375, 82, 3, 2, 2, 2, 376, 377,
	7, 40, 2, 2, 377, 84, 3, 2, 2, 2, 378, 379, 7, 93, 2, 2, 379, 380, 8, 41,
	8, 2, 380, 86, 3, 2, 2, 2, 381, 382, 7, 95, 2, 2, 382, 383, 8, 42, 9, 2,
	383, 88, 3, 2, 2, 2, 384, 385, 7, 125, 2, 2, 385, 90, 3, 2, 2, 2, 386,
	387, 7, 127, 2, 2, 387, 92, 3, 2, 2, 2, 388, 389, 7, 42, 2, 2, 389, 94,
	3, 2, 2, 2, 390, 391, 7, 43, 2, 2, 391, 96, 3, 2, 2, 2, 392, 394, 7, 37,
	2, 2, 393, 395, 7, 15, 2, 2, 394, 393, 3, 2, 2, 2, 394, 395, 3, 2, 2, 2,
	395, 396, 3, 2, 2, 2, 396, 397, 7, 12, 2, 2, 397, 398, 3, 2, 2, 2, 398,
	399, 8, 47, 10, 2, 399, 400, 3, 2, 2, 2, 400, 401, 8, 47, 11, 2, 401, 98,
	3, 2, 2, 2, 402, 403, 7, 37, 2, 2, 403, 404, 3, 2, 2, 2, 404, 405, 8, 48,
	4, 2, 405, 100, 3, 2, 2, 2, 406, 407, 7, 126, 2, 2, 407, 408, 3, 2, 2,
	2, 408, 409, 8, 49, 4, 2, 409, 102, 3, 2, 2, 2, 410, 411, 9, 15, 2, 2,
	411, 104, 3, 2, 2, 2, 412, 413, 9, 16, 2, 2, 413, 106, 3, 2, 2, 2, 414,
	416, 9, 3, 2, 2, 415, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 415,
	3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420, 9, 17,
	2, 2, 420, 421, 8, 52, 12, 2, 421, 108, 3, 2, 2, 2, 422, 424, 9, 3, 2,
	2, 423, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 425,
	426, 3, 2, 2, 2, 426, 427, 3, 2, 2, 2, 427, 429, 7, 37, 2, 2, 428, 430,
	10, 18, 2, 2, 429, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 432, 3,
	2, 2, 2, 431, 429, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 434, 7, 12, 2,
	2, 434, 435, 8, 53, 13, 2, 435, 110, 3, 2, 2, 2, 436, 440, 9, 19, 2, 2,
	437, 439, 9, 19, 2, 2, 438, 437, 3, 2, 2, 2, 439, 442, 3, 2, 2, 2, 440,
	438, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 112, 3, 2, 2, 2, 442, 440,
	3, 2, 2, 2, 443, 445, 10, 20, 2, 2, 444, 443, 3, 2, 2, 2, 445, 448, 3,
	2, 2, 2, 446, 444, 3, 2, 2, 2, 446, 447, 3, 2, 2, 2, 447, 114, 3, 2, 2,
	2, 448, 446, 3, 2, 2, 2, 449, 451, 10, 21, 2, 2, 450, 449, 3, 2, 2, 2,
	451, 454, 3, 2, 2, 2, 452, 450, 3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453,
	116, 3, 2, 2, 2, 454, 452, 3, 2, 2, 2, 455, 456, 5, 103, 50, 2, 456, 457,
	5, 113, 55, 2, 457, 458, 5, 103, 50, 2, 458, 464, 3, 2, 2, 2, 459, 460,
	5, 105, 51, 2, 460, 461, 5, 115, 56, 2, 461, 462, 5, 105, 51, 2, 462, 464,
	3, 2, 2, 2, 463, 455, 3, 2, 2, 2, 463, 459, 3, 2, 2, 2, 464, 465, 3, 2,
	2, 2, 465, 466, 8, 57, 14, 2, 466, 118, 3, 2, 2, 2, 467, 469, 7, 15, 2,
	2, 468, 467, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470,
	471, 7, 12, 2, 2, 471, 472, 8, 58, 15, 2, 472, 473, 3, 2, 2, 2, 473, 474,
	8, 58, 11, 2, 474, 120, 3, 2, 2, 2, 475, 476, 5, 99, 48, 2, 476, 477, 5,
	131, 64, 2, 477, 478, 3, 2, 2, 2, 478, 479, 8, 59, 11, 2, 479, 122, 3,
	2, 2, 2, 480, 482, 10, 22, 2, 2, 481, 480, 3, 2, 2, 2, 482, 483, 3, 2,
	2, 2, 483, 481, 3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 484, 124, 3, 2, 2, 2,
	485, 486, 6, 61, 2, 2, 486, 493, 5, 123, 60, 2, 487, 489, 9, 23, 2, 2,
	488, 487, 3, 2, 2, 2, 489, 490, 3, 2, 2, 2, 490, 488, 3, 2, 2, 2, 490,
	491, 3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492, 494, 5, 123, 60, 2, 493, 488,
	3, 2, 2, 2, 494, 495, 3, 2, 2, 2, 495, 493, 3, 2, 2, 2, 495, 496, 3, 2,
	2, 2, 496, 497, 3, 2, 2, 2, 497, 498, 6, 61, 3, 2, 498, 499, 6, 61, 4,
	2, 499, 500, 6, 61, 5, 2, 500, 126, 3, 2, 2, 2, 501, 505, 9, 24, 2, 2,
	502, 504, 9, 25, 2, 2, 503, 502, 3, 2, 2, 2, 504, 507, 3, 2, 2, 2, 505,
	503, 3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 128, 3, 2, 2, 2, 507, 505,
	3, 2, 2, 2, 508, 510, 9, 3, 2, 2, 509, 508, 3, 2, 2, 2, 510, 511, 3, 2,
	2, 2, 511, 509, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2,
	513, 514, 8, 63, 16, 2, 514, 515, 3, 2, 2, 2, 515, 516, 8, 63, 11, 2, 516,
	130, 3, 2, 2, 2, 517, 519, 10, 17, 2, 2, 518, 517, 3, 2, 2, 2, 519, 520,
	3, 2, 2, 2, 520, 518, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 522, 3, 2,
	2, 2, 522, 523, 8, 64, 17, 2, 523, 132, 3, 2, 2, 2, 524, 525, 9, 26, 2,
	2, 525, 526, 3, 2, 2, 2, 526, 527, 8, 65, 18, 2, 527, 134, 3, 2, 2, 2,
	528, 530, 10, 26, 2, 2, 529, 531, 10, 27, 2, 2, 530, 529, 3, 2, 2, 2, 531,
	532, 3, 2, 2, 2, 532, 530, 3, 2, 2, 2, 532, 533, 3, 2, 2, 2, 533, 534,
	3, 2, 2, 2, 534, 535, 8, 66, 17, 2, 535, 136, 3, 2, 2, 2, 536, 537, 9,
	26, 2, 2, 537, 538, 3, 2, 2, 2, 538, 539, 8, 67, 18, 2, 539, 540, 8, 67,
	17, 2, 540, 138, 3, 2, 2, 2, 541, 545, 9, 24, 2, 2, 542, 544, 9, 28, 2,
	2, 543, 542, 3, 2, 2, 2, 544, 547, 3, 2, 2, 2, 545, 543, 3, 2, 2, 2, 545,
	546, 3, 2, 2, 2, 546, 140, 3, 2, 2, 2, 547, 545, 3, 2, 2, 2, 548, 549,
	9, 26, 2, 2, 549, 550, 3, 2, 2, 2, 550, 551, 8, 69, 18, 2, 551, 142, 3,
	2, 2, 2, 552, 554, 10, 26, 2, 2, 553, 555, 10, 29, 2, 2, 554, 553, 3, 2,
	2, 2, 555, 556, 3, 2, 2, 2, 556, 554, 3, 2, 2, 2, 556, 557, 3, 2, 2, 2,
	557, 558, 3, 2, 2, 2, 558, 559, 8, 70, 17, 2, 559, 144, 3, 2, 2, 2, 32,
	2, 3, 4, 5, 6, 177, 200, 233, 239, 246, 248, 253, 394, 417, 425, 431, 440,
	446, 452, 463, 468, 483, 490, 495, 505, 511, 520, 532, 545, 556, 19, 3,
	3, 2, 3, 9, 3, 7, 3, 2, 7, 4, 2, 7, 6, 2, 7, 5, 2, 3, 41, 4, 3, 42, 5,
	3, 47, 6, 2, 3, 2, 3, 52, 7, 3, 53, 8, 3, 57, 9, 3, 58, 10, 3, 63, 11,
	6, 2, 2, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "NOT_NEWLINE", "FREE_TEXT_NAME", "AT_VAR_DECL", "EVENT_NAME_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "'!wrap'", "'!table'", "'!type'", "", "", "", "", "",
	"", "'...'", "'..'", "'set of'", "", "", "'<->'", "'::'", "'<:'", "'<-'",
	"'->'", "'.. * <- *'", "'+'", "'~'", "','", "'='", "'$'", "'/'", "'-'",
	"'*'", "':'", "'%'", "'.'", "'!'", "'?'", "'@'", "'&'", "'['", "']'", "'{'",
	"'}'", "'('", "')'", "", "'#'", "'|'",
}

var lexerSymbolicNames = []string{
	"", "INDENT", "DEDENT", "NativeDataTypes", "HTTP_VERBS", "WRAP", "TABLE",
	"TYPE", "IMPORT", "RETURN", "IF", "ELSE", "FOR", "LOOP", "WHATEVER", "DOTDOT",
	"SET_OF", "ONE_OF", "MIXIN", "DISTANCE", "NAME_SEP", "LESS_COLON", "ARROW_LEFT",
	"ARROW_RIGHT", "COLLECTOR", "PLUS", "TILDE", "COMMA", "EQ", "DOLLAR", "FORWARD_SLASH",
	"MINUS", "STAR", "COLON", "PERCENT", "DOT", "EXCLAIM", "QN", "AT", "AMP",
	"SQ_OPEN", "SQ_CLOSE", "CURLY_OPEN", "CURLY_CLOSE", "OPEN_PAREN", "CLOSE_PAREN",
	"EMPTY_COMMENT", "HASH", "PIPE", "DBL_QT", "SINGLE_QT", "EMPTY_LINE", "INDENTED_COMMENT",
	"DIGITS", "QSTRING", "NEWLINE", "SYSL_COMMENT", "TEXT_LINE", "Name", "WS",
	"TEXT", "SKIP_WS", "TEXT_NAME", "POP_WS", "VAR_NAME", "SKIP_WS_2", "EVENT_NAME",
}

var lexerRuleNames = []string{
	"NativeDataTypes", "HTTP_VERBS", "WRAP", "TABLE", "TYPE", "IMPORT_KEY",
	"SUB_PATH_NAME", "IMPORT", "RETURN", "IF", "ELSE", "FOR", "LOOP", "WHATEVER",
	"DOTDOT", "SET_OF", "ONE_OF", "MIXIN", "DISTANCE", "NAME_SEP", "LESS_COLON",
	"ARROW_LEFT", "ARROW_RIGHT", "COLLECTOR", "PLUS", "TILDE", "COMMA", "EQ",
	"DOLLAR", "FORWARD_SLASH", "MINUS", "STAR", "COLON", "PERCENT", "DOT",
	"EXCLAIM", "QN", "AT", "AMP", "SQ_OPEN", "SQ_CLOSE", "CURLY_OPEN", "CURLY_CLOSE",
	"OPEN_PAREN", "CLOSE_PAREN", "EMPTY_COMMENT", "HASH", "PIPE", "DBL_QT",
	"SINGLE_QT", "EMPTY_LINE", "INDENTED_COMMENT", "DIGITS", "WITHIN_DBL_QTS",
	"WITHIN_SNGL_QTS", "QSTRING", "NEWLINE", "SYSL_COMMENT", "PRINTABLE", "TEXT_LINE",
	"Name", "WS", "TEXT", "SKIP_WS", "TEXT_NAME", "POP_WS", "VAR_NAME", "SKIP_WS_2",
	"EVENT_NAME",
}

type SyslLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSyslLexer(input antlr.CharStream) *SyslLexer {

	l := new(SyslLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SyslLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SyslLexer tokens.
const (
	SyslLexerINDENT           = 1
	SyslLexerDEDENT           = 2
	SyslLexerNativeDataTypes  = 3
	SyslLexerHTTP_VERBS       = 4
	SyslLexerWRAP             = 5
	SyslLexerTABLE            = 6
	SyslLexerTYPE             = 7
	SyslLexerIMPORT           = 8
	SyslLexerRETURN           = 9
	SyslLexerIF               = 10
	SyslLexerELSE             = 11
	SyslLexerFOR              = 12
	SyslLexerLOOP             = 13
	SyslLexerWHATEVER         = 14
	SyslLexerDOTDOT           = 15
	SyslLexerSET_OF           = 16
	SyslLexerONE_OF           = 17
	SyslLexerMIXIN            = 18
	SyslLexerDISTANCE         = 19
	SyslLexerNAME_SEP         = 20
	SyslLexerLESS_COLON       = 21
	SyslLexerARROW_LEFT       = 22
	SyslLexerARROW_RIGHT      = 23
	SyslLexerCOLLECTOR        = 24
	SyslLexerPLUS             = 25
	SyslLexerTILDE            = 26
	SyslLexerCOMMA            = 27
	SyslLexerEQ               = 28
	SyslLexerDOLLAR           = 29
	SyslLexerFORWARD_SLASH    = 30
	SyslLexerMINUS            = 31
	SyslLexerSTAR             = 32
	SyslLexerCOLON            = 33
	SyslLexerPERCENT          = 34
	SyslLexerDOT              = 35
	SyslLexerEXCLAIM          = 36
	SyslLexerQN               = 37
	SyslLexerAT               = 38
	SyslLexerAMP              = 39
	SyslLexerSQ_OPEN          = 40
	SyslLexerSQ_CLOSE         = 41
	SyslLexerCURLY_OPEN       = 42
	SyslLexerCURLY_CLOSE      = 43
	SyslLexerOPEN_PAREN       = 44
	SyslLexerCLOSE_PAREN      = 45
	SyslLexerEMPTY_COMMENT    = 46
	SyslLexerHASH             = 47
	SyslLexerPIPE             = 48
	SyslLexerDBL_QT           = 49
	SyslLexerSINGLE_QT        = 50
	SyslLexerEMPTY_LINE       = 51
	SyslLexerINDENTED_COMMENT = 52
	SyslLexerDIGITS           = 53
	SyslLexerQSTRING          = 54
	SyslLexerNEWLINE          = 55
	SyslLexerSYSL_COMMENT     = 56
	SyslLexerTEXT_LINE        = 57
	SyslLexerName             = 58
	SyslLexerWS               = 59
	SyslLexerTEXT             = 60
	SyslLexerSKIP_WS          = 61
	SyslLexerTEXT_NAME        = 62
	SyslLexerPOP_WS           = 63
	SyslLexerVAR_NAME         = 64
	SyslLexerSKIP_WS_2        = 65
	SyslLexerEVENT_NAME       = 66
)

// SyslLexer modes.
const (
	SyslLexerNOT_NEWLINE = iota + 1
	SyslLexerFREE_TEXT_NAME
	SyslLexerAT_VAR_DECL
	SyslLexerEVENT_NAME_MODE
)

var spaces int
var linenum int
var in_sq_brackets int

var gotNewLine bool
var gotHttpVerb bool

var prevTokenIndex = -1

func (l *SyslLexer) NextToken() antlr.Token {
	return GetNextToken(l)
}

func (l *SyslLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 1:
		l.HTTP_VERBS_Action(localctx, actionIndex)

	case 7:
		l.IMPORT_Action(localctx, actionIndex)

	case 39:
		l.SQ_OPEN_Action(localctx, actionIndex)

	case 40:
		l.SQ_CLOSE_Action(localctx, actionIndex)

	case 45:
		l.EMPTY_COMMENT_Action(localctx, actionIndex)

	case 50:
		l.EMPTY_LINE_Action(localctx, actionIndex)

	case 51:
		l.INDENTED_COMMENT_Action(localctx, actionIndex)

	case 55:
		l.QSTRING_Action(localctx, actionIndex)

	case 56:
		l.NEWLINE_Action(localctx, actionIndex)

	case 61:
		l.WS_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *SyslLexer) HTTP_VERBS_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		gotHttpVerb = true

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) IMPORT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:
		gotNewLine = true
		spaces = 0
		gotHttpVerb = false
		linenum++

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) SQ_OPEN_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 2:
		in_sq_brackets++

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) SQ_CLOSE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 3:
		in_sq_brackets--

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) EMPTY_COMMENT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 4:
		gotNewLine = true
		spaces = 0
		gotHttpVerb = false
		linenum++

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) EMPTY_LINE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 5:
		l.Skip()
		spaces = 0
		linenum++

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) INDENTED_COMMENT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 6:
		l.Skip()
		spaces = 0
		linenum++

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) QSTRING_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 7:

		var val string
		if json.Unmarshal([]byte(l.GetText()), &val) == nil {
			l.SetText(val)
		} else {
			l.SetText(strings.Trim(l.GetText(), "'"))
		}

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) NEWLINE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 8:
		gotNewLine = true
		gotHttpVerb = false
		spaces = 0
		linenum++

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SyslLexer) WS_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 9:
		spaces = calcSpaces(l.GetText())

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}

func (l *SyslLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 59:
		return l.TEXT_LINE_Sempred(localctx, predIndex)

	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SyslLexer) TEXT_LINE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return !gotHttpVerb

	case 1:
		return in_sq_brackets == 0

	case 2:
		return !gotHttpVerb

	case 3:
		return startsWithKeyword(p.GetText()) == false

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
