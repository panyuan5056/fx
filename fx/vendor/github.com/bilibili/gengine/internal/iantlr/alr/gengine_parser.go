// Code generated from /Users/renyunyi/go_project/github.com/bilibili/gengine/internal/iantlr/gengine.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // gengine

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 331, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 
	39, 3, 2, 6, 2, 80, 10, 2, 13, 2, 14, 2, 81, 3, 3, 3, 3, 3, 3, 5, 3, 87, 
	10, 3, 3, 3, 5, 3, 90, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 
	3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 7, 8, 106, 10, 8, 12, 8, 14, 
	8, 109, 11, 8, 3, 8, 5, 8, 112, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 
	9, 5, 9, 120, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 128, 
	10, 10, 12, 10, 14, 10, 131, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 
	5, 11, 138, 10, 11, 3, 11, 3, 11, 5, 11, 142, 10, 11, 3, 11, 3, 11, 3, 
	11, 3, 11, 5, 11, 148, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 
	3, 11, 3, 11, 7, 11, 158, 10, 11, 12, 11, 14, 11, 161, 11, 11, 3, 12, 3, 
	12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 169, 10, 12, 3, 12, 3, 12, 3, 12, 
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 179, 10, 12, 12, 12, 14, 12, 
	182, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 190, 10, 
	13, 3, 14, 3, 14, 5, 14, 194, 10, 14, 3, 14, 3, 14, 3, 14, 5, 14, 199, 
	10, 14, 3, 15, 3, 15, 5, 15, 203, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 
	16, 3, 16, 7, 16, 211, 10, 16, 12, 16, 14, 16, 214, 11, 16, 3, 16, 5, 16, 
	217, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 
	3, 19, 5, 19, 239, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 
	20, 5, 20, 248, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 
	3, 20, 5, 20, 258, 10, 20, 7, 20, 260, 10, 20, 12, 20, 14, 20, 263, 11, 
	20, 3, 21, 5, 21, 266, 10, 21, 3, 21, 3, 21, 3, 22, 5, 22, 271, 10, 22, 
	3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 282, 
	10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 5, 26, 289, 10, 26, 3, 26, 3, 
	26, 3, 27, 3, 27, 3, 27, 5, 27, 296, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 
	3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 
	34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 319, 10, 35, 3, 35, 
	3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 2, 
	4, 20, 22, 40, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 
	68, 70, 72, 74, 76, 2, 9, 3, 2, 16, 17, 4, 2, 22, 22, 50, 51, 3, 2, 24, 
	25, 3, 2, 26, 27, 3, 2, 28, 33, 3, 2, 10, 11, 3, 2, 35, 40, 2, 352, 2, 
	79, 3, 2, 2, 2, 4, 83, 3, 2, 2, 2, 6, 95, 3, 2, 2, 2, 8, 97, 3, 2, 2, 2, 
	10, 99, 3, 2, 2, 2, 12, 102, 3, 2, 2, 2, 14, 107, 3, 2, 2, 2, 16, 119, 
	3, 2, 2, 2, 18, 121, 3, 2, 2, 2, 20, 147, 3, 2, 2, 2, 22, 168, 3, 2, 2, 
	2, 24, 189, 3, 2, 2, 2, 26, 193, 3, 2, 2, 2, 28, 200, 3, 2, 2, 2, 30, 204, 
	3, 2, 2, 2, 32, 218, 3, 2, 2, 2, 34, 225, 3, 2, 2, 2, 36, 238, 3, 2, 2, 
	2, 38, 247, 3, 2, 2, 2, 40, 265, 3, 2, 2, 2, 42, 270, 3, 2, 2, 2, 44, 274, 
	3, 2, 2, 2, 46, 276, 3, 2, 2, 2, 48, 278, 3, 2, 2, 2, 50, 285, 3, 2, 2, 
	2, 52, 292, 3, 2, 2, 2, 54, 299, 3, 2, 2, 2, 56, 301, 3, 2, 2, 2, 58, 303, 
	3, 2, 2, 2, 60, 305, 3, 2, 2, 2, 62, 307, 3, 2, 2, 2, 64, 309, 3, 2, 2, 
	2, 66, 311, 3, 2, 2, 2, 68, 313, 3, 2, 2, 2, 70, 322, 3, 2, 2, 2, 72, 324, 
	3, 2, 2, 2, 74, 326, 3, 2, 2, 2, 76, 328, 3, 2, 2, 2, 78, 80, 5, 4, 3, 
	2, 79, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 
	3, 2, 2, 2, 82, 3, 3, 2, 2, 2, 83, 84, 7, 9, 2, 2, 84, 86, 5, 6, 4, 2, 
	85, 87, 5, 8, 5, 2, 86, 85, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 
	2, 2, 2, 88, 90, 5, 10, 6, 2, 89, 88, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 
	91, 3, 2, 2, 2, 91, 92, 7, 20, 2, 2, 92, 93, 5, 12, 7, 2, 93, 94, 7, 21, 
	2, 2, 94, 5, 3, 2, 2, 2, 95, 96, 5, 44, 23, 2, 96, 7, 3, 2, 2, 2, 97, 98, 
	5, 44, 23, 2, 98, 9, 3, 2, 2, 2, 99, 100, 7, 19, 2, 2, 100, 101, 5, 40, 
	21, 2, 101, 11, 3, 2, 2, 2, 102, 103, 5, 14, 8, 2, 103, 13, 3, 2, 2, 2, 
	104, 106, 5, 16, 9, 2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 
	105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 
	3, 2, 2, 2, 110, 112, 5, 28, 15, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 
	2, 2, 2, 112, 15, 3, 2, 2, 2, 113, 120, 5, 30, 16, 2, 114, 120, 5, 48, 
	25, 2, 115, 120, 5, 50, 26, 2, 116, 120, 5, 52, 27, 2, 117, 120, 5, 26, 
	14, 2, 118, 120, 5, 18, 10, 2, 119, 113, 3, 2, 2, 2, 119, 114, 3, 2, 2, 
	2, 119, 115, 3, 2, 2, 2, 119, 116, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 
	118, 3, 2, 2, 2, 120, 17, 3, 2, 2, 2, 121, 122, 7, 12, 2, 2, 122, 129, 
	7, 44, 2, 2, 123, 128, 5, 48, 25, 2, 124, 128, 5, 50, 26, 2, 125, 128, 
	5, 52, 27, 2, 126, 128, 5, 26, 14, 2, 127, 123, 3, 2, 2, 2, 127, 124, 3, 
	2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 126, 3, 2, 2, 2, 128, 131, 3, 2, 2, 
	2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 3, 2, 2, 2, 131, 
	129, 3, 2, 2, 2, 132, 133, 7, 45, 2, 2, 133, 19, 3, 2, 2, 2, 134, 135, 
	8, 11, 1, 2, 135, 148, 5, 22, 12, 2, 136, 138, 5, 66, 34, 2, 137, 136, 
	3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 148, 5, 24, 
	13, 2, 140, 142, 5, 66, 34, 2, 141, 140, 3, 2, 2, 2, 141, 142, 3, 2, 2, 
	2, 142, 143, 3, 2, 2, 2, 143, 144, 7, 46, 2, 2, 144, 145, 5, 20, 11, 2, 
	145, 146, 7, 47, 2, 2, 146, 148, 3, 2, 2, 2, 147, 134, 3, 2, 2, 2, 147, 
	137, 3, 2, 2, 2, 147, 141, 3, 2, 2, 2, 148, 159, 3, 2, 2, 2, 149, 150, 
	12, 6, 2, 2, 150, 151, 5, 60, 31, 2, 151, 152, 5, 20, 11, 7, 152, 158, 
	3, 2, 2, 2, 153, 154, 12, 5, 2, 2, 154, 155, 5, 62, 32, 2, 155, 156, 5, 
	20, 11, 6, 156, 158, 3, 2, 2, 2, 157, 149, 3, 2, 2, 2, 157, 153, 3, 2, 
	2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 
	160, 21, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 162, 163, 8, 12, 1, 2, 163, 
	169, 5, 24, 13, 2, 164, 165, 7, 46, 2, 2, 165, 166, 5, 22, 12, 2, 166, 
	167, 7, 47, 2, 2, 167, 169, 3, 2, 2, 2, 168, 162, 3, 2, 2, 2, 168, 164, 
	3, 2, 2, 2, 169, 180, 3, 2, 2, 2, 170, 171, 12, 6, 2, 2, 171, 172, 5, 58, 
	30, 2, 172, 173, 5, 22, 12, 7, 173, 179, 3, 2, 2, 2, 174, 175, 12, 5, 2, 
	2, 175, 176, 5, 56, 29, 2, 176, 177, 5, 22, 12, 6, 177, 179, 3, 2, 2, 2, 
	178, 170, 3, 2, 2, 2, 178, 174, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 
	178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 23, 3, 2, 2, 2, 182, 180, 3, 
	2, 2, 2, 183, 190, 5, 48, 25, 2, 184, 190, 5, 50, 26, 2, 185, 190, 5, 52, 
	27, 2, 186, 190, 5, 36, 19, 2, 187, 190, 5, 68, 35, 2, 188, 190, 5, 54, 
	28, 2, 189, 183, 3, 2, 2, 2, 189, 184, 3, 2, 2, 2, 189, 185, 3, 2, 2, 2, 
	189, 186, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189, 188, 3, 2, 2, 2, 190, 
	25, 3, 2, 2, 2, 191, 194, 5, 68, 35, 2, 192, 194, 5, 54, 28, 2, 193, 191, 
	3, 2, 2, 2, 193, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 198, 5, 64, 
	33, 2, 196, 199, 5, 22, 12, 2, 197, 199, 5, 20, 11, 2, 198, 196, 3, 2, 
	2, 2, 198, 197, 3, 2, 2, 2, 199, 27, 3, 2, 2, 2, 200, 202, 7, 15, 2, 2, 
	201, 203, 5, 20, 11, 2, 202, 201, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 
	29, 3, 2, 2, 2, 204, 205, 7, 13, 2, 2, 205, 206, 5, 20, 11, 2, 206, 207, 
	7, 44, 2, 2, 207, 208, 5, 14, 8, 2, 208, 212, 7, 45, 2, 2, 209, 211, 5, 
	32, 17, 2, 210, 209, 3, 2, 2, 2, 211, 214, 3, 2, 2, 2, 212, 210, 3, 2, 
	2, 2, 212, 213, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 
	215, 217, 5, 34, 18, 2, 216, 215, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 
	31, 3, 2, 2, 2, 218, 219, 7, 14, 2, 2, 219, 220, 7, 13, 2, 2, 220, 221, 
	5, 20, 11, 2, 221, 222, 7, 44, 2, 2, 222, 223, 5, 14, 8, 2, 223, 224, 7, 
	45, 2, 2, 224, 33, 3, 2, 2, 2, 225, 226, 7, 14, 2, 2, 226, 227, 7, 44, 
	2, 2, 227, 228, 5, 14, 8, 2, 228, 229, 7, 45, 2, 2, 229, 35, 3, 2, 2, 2, 
	230, 239, 5, 46, 24, 2, 231, 239, 5, 40, 21, 2, 232, 239, 5, 42, 22, 2, 
	233, 239, 5, 44, 23, 2, 234, 239, 5, 70, 36, 2, 235, 239, 5, 72, 37, 2, 
	236, 239, 5, 74, 38, 2, 237, 239, 5, 76, 39, 2, 238, 230, 3, 2, 2, 2, 238, 
	231, 3, 2, 2, 2, 238, 232, 3, 2, 2, 2, 238, 233, 3, 2, 2, 2, 238, 234, 
	3, 2, 2, 2, 238, 235, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 238, 237, 3, 2, 
	2, 2, 239, 37, 3, 2, 2, 2, 240, 248, 5, 36, 19, 2, 241, 248, 5, 54, 28, 
	2, 242, 248, 5, 48, 25, 2, 243, 248, 5, 50, 26, 2, 244, 248, 5, 52, 27, 
	2, 245, 248, 5, 68, 35, 2, 246, 248, 5, 20, 11, 2, 247, 240, 3, 2, 2, 2, 
	247, 241, 3, 2, 2, 2, 247, 242, 3, 2, 2, 2, 247, 243, 3, 2, 2, 2, 247, 
	244, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 246, 3, 2, 2, 2, 248, 261, 
	3, 2, 2, 2, 249, 257, 7, 3, 2, 2, 250, 258, 5, 36, 19, 2, 251, 258, 5, 
	54, 28, 2, 252, 258, 5, 48, 25, 2, 253, 258, 5, 50, 26, 2, 254, 258, 5, 
	52, 27, 2, 255, 258, 5, 68, 35, 2, 256, 258, 5, 20, 11, 2, 257, 250, 3, 
	2, 2, 2, 257, 251, 3, 2, 2, 2, 257, 252, 3, 2, 2, 2, 257, 253, 3, 2, 2, 
	2, 257, 254, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 256, 3, 2, 2, 2, 258, 
	260, 3, 2, 2, 2, 259, 249, 3, 2, 2, 2, 260, 263, 3, 2, 2, 2, 261, 259, 
	3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 39, 3, 2, 2, 2, 263, 261, 3, 2, 
	2, 2, 264, 266, 7, 25, 2, 2, 265, 264, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 
	266, 267, 3, 2, 2, 2, 267, 268, 7, 23, 2, 2, 268, 41, 3, 2, 2, 2, 269, 
	271, 7, 25, 2, 2, 270, 269, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 272, 
	3, 2, 2, 2, 272, 273, 7, 52, 2, 2, 273, 43, 3, 2, 2, 2, 274, 275, 7, 49, 
	2, 2, 275, 45, 3, 2, 2, 2, 276, 277, 9, 2, 2, 2, 277, 47, 3, 2, 2, 2, 278, 
	279, 7, 22, 2, 2, 279, 281, 7, 46, 2, 2, 280, 282, 5, 38, 20, 2, 281, 280, 
	3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 284, 7, 47, 
	2, 2, 284, 49, 3, 2, 2, 2, 285, 286, 7, 50, 2, 2, 286, 288, 7, 46, 2, 2, 
	287, 289, 5, 38, 20, 2, 288, 287, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 
	290, 3, 2, 2, 2, 290, 291, 7, 47, 2, 2, 291, 51, 3, 2, 2, 2, 292, 293, 
	7, 51, 2, 2, 293, 295, 7, 46, 2, 2, 294, 296, 5, 38, 20, 2, 295, 294, 3, 
	2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 7, 47, 2, 
	2, 298, 53, 3, 2, 2, 2, 299, 300, 9, 3, 2, 2, 300, 55, 3, 2, 2, 2, 301, 
	302, 9, 4, 2, 2, 302, 57, 3, 2, 2, 2, 303, 304, 9, 5, 2, 2, 304, 59, 3, 
	2, 2, 2, 305, 306, 9, 6, 2, 2, 306, 61, 3, 2, 2, 2, 307, 308, 9, 7, 2, 
	2, 308, 63, 3, 2, 2, 2, 309, 310, 9, 8, 2, 2, 310, 65, 3, 2, 2, 2, 311, 
	312, 7, 34, 2, 2, 312, 67, 3, 2, 2, 2, 313, 314, 5, 54, 28, 2, 314, 318, 
	7, 41, 2, 2, 315, 319, 5, 40, 21, 2, 316, 319, 5, 44, 23, 2, 317, 319, 
	5, 54, 28, 2, 318, 315, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 318, 317, 3, 
	2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 321, 7, 42, 2, 2, 321, 69, 3, 2, 2, 
	2, 322, 323, 7, 4, 2, 2, 323, 71, 3, 2, 2, 2, 324, 325, 7, 5, 2, 2, 325, 
	73, 3, 2, 2, 2, 326, 327, 7, 6, 2, 2, 327, 75, 3, 2, 2, 2, 328, 329, 7, 
	7, 2, 2, 329, 77, 3, 2, 2, 2, 34, 81, 86, 89, 107, 111, 119, 127, 129, 
	137, 141, 147, 157, 159, 168, 178, 180, 189, 193, 198, 202, 212, 216, 238, 
	247, 257, 261, 265, 270, 281, 288, 295, 318,
}
var literalNames = []string{
	"", "','", "'@name'", "'@id'", "'@desc'", "'@sal'", "", "", "'&&'", "'||'", 
	"", "", "", "", "", "", "", "", "", "", "", "", "'+'", "'-'", "'/'", "'*'", 
	"'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'!'", "':='", "'='", "'+='", 
	"'-='", "'*='", "'/='", "'['", "']'", "';'", "'{'", "'}'", "'('", "')'", 
	"'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "CONC", "IF", "ELSE", 
	"RETURN", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END", 
	"SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", 
	"GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", 
	"MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", 
	"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", 
	"DOUBLEDOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
}

var ruleNames = []string{
	"primary", "ruleEntity", "ruleName", "ruleDescription", "salience", "ruleContent", 
	"statements", "statement", "concStatement", "expression", "mathExpression", 
	"expressionAtom", "assignment", "returnStmt", "ifStmt", "elseIfStmt", "elseStmt", 
	"constant", "functionArgs", "integer", "realLiteral", "stringLiteral", 
	"booleanLiteral", "functionCall", "methodCall", "threeLevelCall", "variable", 
	"mathPmOperator", "mathMdOperator", "comparisonOperator", "logicalOperator", 
	"assignOperator", "notOperator", "mapVar", "atName", "atId", "atDesc", 
	"atSal",
}
type gengineParser struct {
	*antlr.BaseParser
}

// NewgengineParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *gengineParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgengineParser(input antlr.TokenStream) *gengineParser {
	this := new(gengineParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "gengine.g4"

	return this
}


// gengineParser tokens.
const (
	gengineParserEOF = antlr.TokenEOF
	gengineParserT__0 = 1
	gengineParserT__1 = 2
	gengineParserT__2 = 3
	gengineParserT__3 = 4
	gengineParserT__4 = 5
	gengineParserNIL = 6
	gengineParserRULE = 7
	gengineParserAND = 8
	gengineParserOR = 9
	gengineParserCONC = 10
	gengineParserIF = 11
	gengineParserELSE = 12
	gengineParserRETURN = 13
	gengineParserTRUE = 14
	gengineParserFALSE = 15
	gengineParserNULL_LITERAL = 16
	gengineParserSALIENCE = 17
	gengineParserBEGIN = 18
	gengineParserEND = 19
	gengineParserSIMPLENAME = 20
	gengineParserINT = 21
	gengineParserPLUS = 22
	gengineParserMINUS = 23
	gengineParserDIV = 24
	gengineParserMUL = 25
	gengineParserEQUALS = 26
	gengineParserGT = 27
	gengineParserLT = 28
	gengineParserGTE = 29
	gengineParserLTE = 30
	gengineParserNOTEQUALS = 31
	gengineParserNOT = 32
	gengineParserASSIGN = 33
	gengineParserSET = 34
	gengineParserPLUSEQUAL = 35
	gengineParserMINUSEQUAL = 36
	gengineParserMULTIEQUAL = 37
	gengineParserDIVEQUAL = 38
	gengineParserLSQARE = 39
	gengineParserRSQARE = 40
	gengineParserSEMICOLON = 41
	gengineParserLR_BRACE = 42
	gengineParserRR_BRACE = 43
	gengineParserLR_BRACKET = 44
	gengineParserRR_BRACKET = 45
	gengineParserDOT = 46
	gengineParserDQUOTA_STRING = 47
	gengineParserDOTTEDNAME = 48
	gengineParserDOUBLEDOTTEDNAME = 49
	gengineParserREAL_LITERAL = 50
	gengineParserSL_COMMENT = 51
	gengineParserWS = 52
)

// gengineParser rules.
const (
	gengineParserRULE_primary = 0
	gengineParserRULE_ruleEntity = 1
	gengineParserRULE_ruleName = 2
	gengineParserRULE_ruleDescription = 3
	gengineParserRULE_salience = 4
	gengineParserRULE_ruleContent = 5
	gengineParserRULE_statements = 6
	gengineParserRULE_statement = 7
	gengineParserRULE_concStatement = 8
	gengineParserRULE_expression = 9
	gengineParserRULE_mathExpression = 10
	gengineParserRULE_expressionAtom = 11
	gengineParserRULE_assignment = 12
	gengineParserRULE_returnStmt = 13
	gengineParserRULE_ifStmt = 14
	gengineParserRULE_elseIfStmt = 15
	gengineParserRULE_elseStmt = 16
	gengineParserRULE_constant = 17
	gengineParserRULE_functionArgs = 18
	gengineParserRULE_integer = 19
	gengineParserRULE_realLiteral = 20
	gengineParserRULE_stringLiteral = 21
	gengineParserRULE_booleanLiteral = 22
	gengineParserRULE_functionCall = 23
	gengineParserRULE_methodCall = 24
	gengineParserRULE_threeLevelCall = 25
	gengineParserRULE_variable = 26
	gengineParserRULE_mathPmOperator = 27
	gengineParserRULE_mathMdOperator = 28
	gengineParserRULE_comparisonOperator = 29
	gengineParserRULE_logicalOperator = 30
	gengineParserRULE_assignOperator = 31
	gengineParserRULE_notOperator = 32
	gengineParserRULE_mapVar = 33
	gengineParserRULE_atName = 34
	gengineParserRULE_atId = 35
	gengineParserRULE_atDesc = 36
	gengineParserRULE_atSal = 37
)

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) AllRuleEntity() []IRuleEntityContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleEntityContext)(nil)).Elem())
	var tst = make([]IRuleEntityContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleEntityContext)
		}
	}

	return tst
}

func (s *PrimaryContext) RuleEntity(i int) IRuleEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleEntityContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleEntityContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gengineParserRULE_primary)
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == gengineParserRULE {
		{
			p.SetState(76)
			p.RuleEntity()
		}


		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IRuleEntityContext is an interface to support dynamic dispatch.
type IRuleEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleEntityContext differentiates from other interfaces.
	IsRuleEntityContext()
}

type RuleEntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleEntityContext() *RuleEntityContext {
	var p = new(RuleEntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleEntity
	return p
}

func (*RuleEntityContext) IsRuleEntityContext() {}

func NewRuleEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleEntityContext {
	var p = new(RuleEntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleEntity

	return p
}

func (s *RuleEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleEntityContext) RULE() antlr.TerminalNode {
	return s.GetToken(gengineParserRULE, 0)
}

func (s *RuleEntityContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *RuleEntityContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(gengineParserBEGIN, 0)
}

func (s *RuleEntityContext) RuleContent() IRuleContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleContentContext)
}

func (s *RuleEntityContext) END() antlr.TerminalNode {
	return s.GetToken(gengineParserEND, 0)
}

func (s *RuleEntityContext) RuleDescription() IRuleDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleDescriptionContext)
}

func (s *RuleEntityContext) Salience() ISalienceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISalienceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISalienceContext)
}

func (s *RuleEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RuleEntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleEntity(s)
	}
}

func (s *RuleEntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleEntity(s)
	}
}

func (s *RuleEntityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleEntity(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) RuleEntity() (localctx IRuleEntityContext) {
	localctx = NewRuleEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gengineParserRULE_ruleEntity)
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
		p.SetState(81)
		p.Match(gengineParserRULE)
	}
	{
		p.SetState(82)
		p.RuleName()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserDQUOTA_STRING {
		{
			p.SetState(83)
			p.RuleDescription()
		}

	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserSALIENCE {
		{
			p.SetState(86)
			p.Salience()
		}

	}
	{
		p.SetState(89)
		p.Match(gengineParserBEGIN)
	}
	{
		p.SetState(90)
		p.RuleContent()
	}
	{
		p.SetState(91)
		p.Match(gengineParserEND)
	}



	return localctx
}


// IRuleNameContext is an interface to support dynamic dispatch.
type IRuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleNameContext differentiates from other interfaces.
	IsRuleNameContext()
}

type RuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleNameContext() *RuleNameContext {
	var p = new(RuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleName
	return p
}

func (*RuleNameContext) IsRuleNameContext() {}

func NewRuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleNameContext {
	var p = new(RuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleName

	return p
}

func (s *RuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleNameContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleName(s)
	}
}

func (s *RuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleName(s)
	}
}

func (s *RuleNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) RuleName() (localctx IRuleNameContext) {
	localctx = NewRuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gengineParserRULE_ruleName)

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
		p.SetState(93)
		p.StringLiteral()
	}



	return localctx
}


// IRuleDescriptionContext is an interface to support dynamic dispatch.
type IRuleDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleDescriptionContext differentiates from other interfaces.
	IsRuleDescriptionContext()
}

type RuleDescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDescriptionContext() *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleDescription
	return p
}

func (*RuleDescriptionContext) IsRuleDescriptionContext() {}

func NewRuleDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleDescription

	return p
}

func (s *RuleDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDescriptionContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RuleDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RuleDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleDescription(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) RuleDescription() (localctx IRuleDescriptionContext) {
	localctx = NewRuleDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gengineParserRULE_ruleDescription)

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
		p.SetState(95)
		p.StringLiteral()
	}



	return localctx
}


// ISalienceContext is an interface to support dynamic dispatch.
type ISalienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalienceContext differentiates from other interfaces.
	IsSalienceContext()
}

type SalienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalienceContext() *SalienceContext {
	var p = new(SalienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_salience
	return p
}

func (*SalienceContext) IsSalienceContext() {}

func NewSalienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalienceContext {
	var p = new(SalienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_salience

	return p
}

func (s *SalienceContext) GetParser() antlr.Parser { return s.parser }

func (s *SalienceContext) SALIENCE() antlr.TerminalNode {
	return s.GetToken(gengineParserSALIENCE, 0)
}

func (s *SalienceContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *SalienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SalienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterSalience(s)
	}
}

func (s *SalienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitSalience(s)
	}
}

func (s *SalienceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitSalience(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) Salience() (localctx ISalienceContext) {
	localctx = NewSalienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gengineParserRULE_salience)

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
		p.SetState(97)
		p.Match(gengineParserSALIENCE)
	}
	{
		p.SetState(98)
		p.Integer()
	}



	return localctx
}


// IRuleContentContext is an interface to support dynamic dispatch.
type IRuleContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleContentContext differentiates from other interfaces.
	IsRuleContentContext()
}

type RuleContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContentContext() *RuleContentContext {
	var p = new(RuleContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleContent
	return p
}

func (*RuleContentContext) IsRuleContentContext() {}

func NewRuleContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContentContext {
	var p = new(RuleContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleContent

	return p
}

func (s *RuleContentContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContentContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *RuleContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RuleContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleContent(s)
	}
}

func (s *RuleContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleContent(s)
	}
}

func (s *RuleContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleContent(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) RuleContent() (localctx IRuleContentContext) {
	localctx = NewRuleContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gengineParserRULE_ruleContent)

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
		p.SetState(100)
		p.Statements()
	}



	return localctx
}


// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) ReturnStmt() IReturnStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gengineParserRULE_statements)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserCONC) | (1 << gengineParserIF) | (1 << gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME || _la == gengineParserDOUBLEDOTTEDNAME {
		{
			p.SetState(102)
			p.Statement()
		}


		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserRETURN {
		{
			p.SetState(108)
			p.ReturnStmt()
		}

	}



	return localctx
}


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *StatementContext) ThreeLevelCall() IThreeLevelCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ConcStatement() IConcStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gengineParserRULE_statement)

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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.IfStmt()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.FunctionCall()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.MethodCall()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.ThreeLevelCall()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.Assignment()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.ConcStatement()
		}

	}


	return localctx
}


// IConcStatementContext is an interface to support dynamic dispatch.
type IConcStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcStatementContext differentiates from other interfaces.
	IsConcStatementContext()
}

type ConcStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcStatementContext() *ConcStatementContext {
	var p = new(ConcStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_concStatement
	return p
}

func (*ConcStatementContext) IsConcStatementContext() {}

func NewConcStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcStatementContext {
	var p = new(ConcStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_concStatement

	return p
}

func (s *ConcStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcStatementContext) CONC() antlr.TerminalNode {
	return s.GetToken(gengineParserCONC, 0)
}

func (s *ConcStatementContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ConcStatementContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ConcStatementContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ConcStatementContext) AllMethodCall() []IMethodCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodCallContext)(nil)).Elem())
	var tst = make([]IMethodCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodCallContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) MethodCall(i int) IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ConcStatementContext) AllThreeLevelCall() []IThreeLevelCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem())
	var tst = make([]IThreeLevelCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThreeLevelCallContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) ThreeLevelCall(i int) IThreeLevelCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *ConcStatementContext) AllAssignment() []IAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignmentContext)(nil)).Elem())
	var tst = make([]IAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignmentContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) Assignment(i int) IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ConcStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConcStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterConcStatement(s)
	}
}

func (s *ConcStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitConcStatement(s)
	}
}

func (s *ConcStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitConcStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) ConcStatement() (localctx IConcStatementContext) {
	localctx = NewConcStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gengineParserRULE_concStatement)
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
		p.SetState(119)
		p.Match(gengineParserCONC)
	}
	{
		p.SetState(120)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 20)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 20))) & ((1 << (gengineParserSIMPLENAME - 20)) | (1 << (gengineParserDOTTEDNAME - 20)) | (1 << (gengineParserDOUBLEDOTTEDNAME - 20)))) != 0) {
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(121)
				p.FunctionCall()
			}


		case 2:
			{
				p.SetState(122)
				p.MethodCall()
			}


		case 3:
			{
				p.SetState(123)
				p.ThreeLevelCall()
			}


		case 4:
			{
				p.SetState(124)
				p.Assignment()
			}

		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(130)
		p.Match(gengineParserRR_BRACE)
	}



	return localctx
}


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) MathExpression() IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *ExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionContext) NotOperator() INotOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotOperatorContext)
}

func (s *ExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *ExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *gengineParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *gengineParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, gengineParserRULE_expression, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(133)
			p.mathExpression(0)
		}


	case 2:
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == gengineParserNOT {
			{
				p.SetState(134)
				p.NotOperator()
			}

		}
		{
			p.SetState(137)
			p.ExpressionAtom()
		}


	case 3:
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == gengineParserNOT {
			{
				p.SetState(138)
				p.NotOperator()
			}

		}
		{
			p.SetState(141)
			p.Match(gengineParserLR_BRACKET)
		}
		{
			p.SetState(142)
			p.expression(0)
		}
		{
			p.SetState(143)
			p.Match(gengineParserRR_BRACKET)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(148)
					p.ComparisonOperator()
				}
				{
					p.SetState(149)
					p.expression(5)
				}


			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(152)
					p.LogicalOperator()
				}
				{
					p.SetState(153)
					p.expression(4)
				}

			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}



	return localctx
}


// IMathExpressionContext is an interface to support dynamic dispatch.
type IMathExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathExpressionContext differentiates from other interfaces.
	IsMathExpressionContext()
}

type MathExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathExpressionContext() *MathExpressionContext {
	var p = new(MathExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathExpression
	return p
}

func (*MathExpressionContext) IsMathExpressionContext() {}

func NewMathExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathExpressionContext {
	var p = new(MathExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathExpression

	return p
}

func (s *MathExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MathExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *MathExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *MathExpressionContext) AllMathExpression() []IMathExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem())
	var tst = make([]IMathExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathExpressionContext)
		}
	}

	return tst
}

func (s *MathExpressionContext) MathExpression(i int) IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *MathExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *MathExpressionContext) MathMdOperator() IMathMdOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathMdOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathMdOperatorContext)
}

func (s *MathExpressionContext) MathPmOperator() IMathPmOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathPmOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathPmOperatorContext)
}

func (s *MathExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MathExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathExpression(s)
	}
}

func (s *MathExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathExpression(s)
	}
}

func (s *MathExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathExpression(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *gengineParser) MathExpression() (localctx IMathExpressionContext) {
	return p.mathExpression(0)
}

func (p *gengineParser) mathExpression(_p int) (localctx IMathExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, gengineParserRULE_mathExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserT__1, gengineParserT__2, gengineParserT__3, gengineParserT__4, gengineParserTRUE, gengineParserFALSE, gengineParserSIMPLENAME, gengineParserINT, gengineParserMINUS, gengineParserDQUOTA_STRING, gengineParserDOTTEDNAME, gengineParserDOUBLEDOTTEDNAME, gengineParserREAL_LITERAL:
		{
			p.SetState(161)
			p.ExpressionAtom()
		}


	case gengineParserLR_BRACKET:
		{
			p.SetState(162)
			p.Match(gengineParserLR_BRACKET)
		}
		{
			p.SetState(163)
			p.mathExpression(0)
		}
		{
			p.SetState(164)
			p.Match(gengineParserRR_BRACKET)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(169)
					p.MathMdOperator()
				}
				{
					p.SetState(170)
					p.mathExpression(5)
				}


			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(173)
					p.MathPmOperator()
				}
				{
					p.SetState(174)
					p.mathExpression(4)
				}

			}

		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}



	return localctx
}


// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionAtomContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ExpressionAtomContext) ThreeLevelCall() IThreeLevelCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *ExpressionAtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionAtomContext) MapVar() IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *ExpressionAtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gengineParserRULE_expressionAtom)

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

	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.FunctionCall()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.MethodCall()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(183)
			p.ThreeLevelCall()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(184)
			p.Constant()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(185)
			p.MapVar()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(186)
			p.Variable()
		}

	}


	return localctx
}


// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AssignOperator() IAssignOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignOperatorContext)
}

func (s *AssignmentContext) MapVar() IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *AssignmentContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) MathExpression() IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gengineParserRULE_assignment)

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
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(189)
			p.MapVar()
		}


	case 2:
		{
			p.SetState(190)
			p.Variable()
		}

	}
	{
		p.SetState(193)
		p.AssignOperator()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(194)
			p.mathExpression(0)
		}


	case 2:
		{
			p.SetState(195)
			p.expression(0)
		}

	}



	return localctx
}


// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_returnStmt
	return p
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(gengineParserRETURN, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gengineParserRULE_returnStmt)
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
		p.SetState(198)
		p.Match(gengineParserRETURN)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__1) | (1 << gengineParserT__2) | (1 << gengineParserT__3) | (1 << gengineParserT__4) | (1 << gengineParserTRUE) | (1 << gengineParserFALSE) | (1 << gengineParserSIMPLENAME) | (1 << gengineParserINT) | (1 << gengineParserMINUS))) != 0) || ((((_la - 32)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 32))) & ((1 << (gengineParserNOT - 32)) | (1 << (gengineParserLR_BRACKET - 32)) | (1 << (gengineParserDQUOTA_STRING - 32)) | (1 << (gengineParserDOTTEDNAME - 32)) | (1 << (gengineParserDOUBLEDOTTEDNAME - 32)) | (1 << (gengineParserREAL_LITERAL - 32)))) != 0) {
		{
			p.SetState(199)
			p.expression(0)
		}

	}



	return localctx
}


// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(gengineParserIF, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *IfStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *IfStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *IfStmtContext) AllElseIfStmt() []IElseIfStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStmtContext)(nil)).Elem())
	var tst = make([]IElseIfStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStmtContext)
		}
	}

	return tst
}

func (s *IfStmtContext) ElseIfStmt(i int) IElseIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStmtContext)
}

func (s *IfStmtContext) ElseStmt() IElseStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gengineParserRULE_ifStmt)
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
		p.SetState(202)
		p.Match(gengineParserIF)
	}
	{
		p.SetState(203)
		p.expression(0)
	}
	{
		p.SetState(204)
		p.Match(gengineParserLR_BRACE)
	}
	{
		p.SetState(205)
		p.Statements()
	}
	{
		p.SetState(206)
		p.Match(gengineParserRR_BRACE)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(207)
				p.ElseIfStmt()
			}


		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserELSE {
		{
			p.SetState(213)
			p.ElseStmt()
		}

	}



	return localctx
}


// IElseIfStmtContext is an interface to support dynamic dispatch.
type IElseIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStmtContext differentiates from other interfaces.
	IsElseIfStmtContext()
}

type ElseIfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStmtContext() *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_elseIfStmt
	return p
}

func (*ElseIfStmtContext) IsElseIfStmtContext() {}

func NewElseIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_elseIfStmt

	return p
}

func (s *ElseIfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(gengineParserELSE, 0)
}

func (s *ElseIfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(gengineParserIF, 0)
}

func (s *ElseIfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ElseIfStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ElseIfStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ElseIfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElseIfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitElseIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) ElseIfStmt() (localctx IElseIfStmtContext) {
	localctx = NewElseIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gengineParserRULE_elseIfStmt)

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
		p.SetState(216)
		p.Match(gengineParserELSE)
	}
	{
		p.SetState(217)
		p.Match(gengineParserIF)
	}
	{
		p.SetState(218)
		p.expression(0)
	}
	{
		p.SetState(219)
		p.Match(gengineParserLR_BRACE)
	}
	{
		p.SetState(220)
		p.Statements()
	}
	{
		p.SetState(221)
		p.Match(gengineParserRR_BRACE)
	}



	return localctx
}


// IElseStmtContext is an interface to support dynamic dispatch.
type IElseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStmtContext differentiates from other interfaces.
	IsElseStmtContext()
}

type ElseStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStmtContext() *ElseStmtContext {
	var p = new(ElseStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_elseStmt
	return p
}

func (*ElseStmtContext) IsElseStmtContext() {}

func NewElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStmtContext {
	var p = new(ElseStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_elseStmt

	return p
}

func (s *ElseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(gengineParserELSE, 0)
}

func (s *ElseStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ElseStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ElseStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) ElseStmt() (localctx IElseStmtContext) {
	localctx = NewElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gengineParserRULE_elseStmt)

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
		p.SetState(223)
		p.Match(gengineParserELSE)
	}
	{
		p.SetState(224)
		p.Match(gengineParserLR_BRACE)
	}
	{
		p.SetState(225)
		p.Statements()
	}
	{
		p.SetState(226)
		p.Match(gengineParserRR_BRACE)
	}



	return localctx
}


// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ConstantContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *ConstantContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) AtName() IAtNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtNameContext)
}

func (s *ConstantContext) AtId() IAtIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtIdContext)
}

func (s *ConstantContext) AtDesc() IAtDescContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtDescContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtDescContext)
}

func (s *ConstantContext) AtSal() IAtSalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtSalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtSalContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gengineParserRULE_constant)

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

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.BooleanLiteral()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Integer()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.RealLiteral()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.StringLiteral()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(232)
			p.AtName()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(233)
			p.AtId()
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(234)
			p.AtDesc()
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(235)
			p.AtSal()
		}

	}


	return localctx
}


// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FunctionArgsContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctionArgsContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionArgsContext) AllMethodCall() []IMethodCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodCallContext)(nil)).Elem())
	var tst = make([]IMethodCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) MethodCall(i int) IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *FunctionArgsContext) AllThreeLevelCall() []IThreeLevelCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem())
	var tst = make([]IThreeLevelCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThreeLevelCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) ThreeLevelCall(i int) IThreeLevelCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *FunctionArgsContext) AllMapVar() []IMapVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapVarContext)(nil)).Elem())
	var tst = make([]IMapVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapVarContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) MapVar(i int) IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *FunctionArgsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gengineParserRULE_functionArgs)
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
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(238)
			p.Constant()
		}


	case 2:
		{
			p.SetState(239)
			p.Variable()
		}


	case 3:
		{
			p.SetState(240)
			p.FunctionCall()
		}


	case 4:
		{
			p.SetState(241)
			p.MethodCall()
		}


	case 5:
		{
			p.SetState(242)
			p.ThreeLevelCall()
		}


	case 6:
		{
			p.SetState(243)
			p.MapVar()
		}


	case 7:
		{
			p.SetState(244)
			p.expression(0)
		}

	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == gengineParserT__0 {
		{
			p.SetState(247)
			p.Match(gengineParserT__0)
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(248)
				p.Constant()
			}


		case 2:
			{
				p.SetState(249)
				p.Variable()
			}


		case 3:
			{
				p.SetState(250)
				p.FunctionCall()
			}


		case 4:
			{
				p.SetState(251)
				p.MethodCall()
			}


		case 5:
			{
				p.SetState(252)
				p.ThreeLevelCall()
			}


		case 6:
			{
				p.SetState(253)
				p.MapVar()
			}


		case 7:
			{
				p.SetState(254)
				p.expression(0)
			}

		}


		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(gengineParserINT, 0)
}

func (s *IntegerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gengineParserRULE_integer)
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
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserMINUS {
		{
			p.SetState(262)
			p.Match(gengineParserMINUS)
		}

	}
	{
		p.SetState(265)
		p.Match(gengineParserINT)
	}



	return localctx
}


// IRealLiteralContext is an interface to support dynamic dispatch.
type IRealLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealLiteralContext differentiates from other interfaces.
	IsRealLiteralContext()
}

type RealLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealLiteralContext() *RealLiteralContext {
	var p = new(RealLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_realLiteral
	return p
}

func (*RealLiteralContext) IsRealLiteralContext() {}

func NewRealLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealLiteralContext {
	var p = new(RealLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_realLiteral

	return p
}

func (s *RealLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RealLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(gengineParserREAL_LITERAL, 0)
}

func (s *RealLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RealLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRealLiteral(s)
	}
}

func (s *RealLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRealLiteral(s)
	}
}

func (s *RealLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRealLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) RealLiteral() (localctx IRealLiteralContext) {
	localctx = NewRealLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gengineParserRULE_realLiteral)
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
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserMINUS {
		{
			p.SetState(267)
			p.Match(gengineParserMINUS)
		}

	}
	{
		p.SetState(270)
		p.Match(gengineParserREAL_LITERAL)
	}



	return localctx
}


// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(gengineParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gengineParserRULE_stringLiteral)

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
		p.SetState(272)
		p.Match(gengineParserDQUOTA_STRING)
	}



	return localctx
}


// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(gengineParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(gengineParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gengineParserRULE_booleanLiteral)
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
		p.SetState(274)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserTRUE || _la == gengineParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(gengineParserSIMPLENAME, 0)
}

func (s *FunctionCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *FunctionCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *FunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gengineParserRULE_functionCall)
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
		p.SetState(276)
		p.Match(gengineParserSIMPLENAME)
	}
	{
		p.SetState(277)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__1) | (1 << gengineParserT__2) | (1 << gengineParserT__3) | (1 << gengineParserT__4) | (1 << gengineParserTRUE) | (1 << gengineParserFALSE) | (1 << gengineParserSIMPLENAME) | (1 << gengineParserINT) | (1 << gengineParserMINUS))) != 0) || ((((_la - 32)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 32))) & ((1 << (gengineParserNOT - 32)) | (1 << (gengineParserLR_BRACKET - 32)) | (1 << (gengineParserDQUOTA_STRING - 32)) | (1 << (gengineParserDOTTEDNAME - 32)) | (1 << (gengineParserDOUBLEDOTTEDNAME - 32)) | (1 << (gengineParserREAL_LITERAL - 32)))) != 0) {
		{
			p.SetState(278)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(281)
		p.Match(gengineParserRR_BRACKET)
	}



	return localctx
}


// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOTTEDNAME, 0)
}

func (s *MethodCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *MethodCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *MethodCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gengineParserRULE_methodCall)
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
		p.SetState(283)
		p.Match(gengineParserDOTTEDNAME)
	}
	{
		p.SetState(284)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__1) | (1 << gengineParserT__2) | (1 << gengineParserT__3) | (1 << gengineParserT__4) | (1 << gengineParserTRUE) | (1 << gengineParserFALSE) | (1 << gengineParserSIMPLENAME) | (1 << gengineParserINT) | (1 << gengineParserMINUS))) != 0) || ((((_la - 32)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 32))) & ((1 << (gengineParserNOT - 32)) | (1 << (gengineParserLR_BRACKET - 32)) | (1 << (gengineParserDQUOTA_STRING - 32)) | (1 << (gengineParserDOTTEDNAME - 32)) | (1 << (gengineParserDOUBLEDOTTEDNAME - 32)) | (1 << (gengineParserREAL_LITERAL - 32)))) != 0) {
		{
			p.SetState(285)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(288)
		p.Match(gengineParserRR_BRACKET)
	}



	return localctx
}


// IThreeLevelCallContext is an interface to support dynamic dispatch.
type IThreeLevelCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThreeLevelCallContext differentiates from other interfaces.
	IsThreeLevelCallContext()
}

type ThreeLevelCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThreeLevelCallContext() *ThreeLevelCallContext {
	var p = new(ThreeLevelCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_threeLevelCall
	return p
}

func (*ThreeLevelCallContext) IsThreeLevelCallContext() {}

func NewThreeLevelCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThreeLevelCallContext {
	var p = new(ThreeLevelCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_threeLevelCall

	return p
}

func (s *ThreeLevelCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ThreeLevelCallContext) DOUBLEDOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOUBLEDOTTEDNAME, 0)
}

func (s *ThreeLevelCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *ThreeLevelCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *ThreeLevelCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *ThreeLevelCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThreeLevelCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ThreeLevelCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterThreeLevelCall(s)
	}
}

func (s *ThreeLevelCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitThreeLevelCall(s)
	}
}

func (s *ThreeLevelCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitThreeLevelCall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) ThreeLevelCall() (localctx IThreeLevelCallContext) {
	localctx = NewThreeLevelCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gengineParserRULE_threeLevelCall)
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
		p.SetState(290)
		p.Match(gengineParserDOUBLEDOTTEDNAME)
	}
	{
		p.SetState(291)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__1) | (1 << gengineParserT__2) | (1 << gengineParserT__3) | (1 << gengineParserT__4) | (1 << gengineParserTRUE) | (1 << gengineParserFALSE) | (1 << gengineParserSIMPLENAME) | (1 << gengineParserINT) | (1 << gengineParserMINUS))) != 0) || ((((_la - 32)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 32))) & ((1 << (gengineParserNOT - 32)) | (1 << (gengineParserLR_BRACKET - 32)) | (1 << (gengineParserDQUOTA_STRING - 32)) | (1 << (gengineParserDOTTEDNAME - 32)) | (1 << (gengineParserDOUBLEDOTTEDNAME - 32)) | (1 << (gengineParserREAL_LITERAL - 32)))) != 0) {
		{
			p.SetState(292)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(295)
		p.Match(gengineParserRR_BRACKET)
	}



	return localctx
}


// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(gengineParserSIMPLENAME, 0)
}

func (s *VariableContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOTTEDNAME, 0)
}

func (s *VariableContext) DOUBLEDOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOUBLEDOTTEDNAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gengineParserRULE_variable)
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
		p.SetState(297)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 20)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 20))) & ((1 << (gengineParserSIMPLENAME - 20)) | (1 << (gengineParserDOTTEDNAME - 20)) | (1 << (gengineParserDOUBLEDOTTEDNAME - 20)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IMathPmOperatorContext is an interface to support dynamic dispatch.
type IMathPmOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathPmOperatorContext differentiates from other interfaces.
	IsMathPmOperatorContext()
}

type MathPmOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathPmOperatorContext() *MathPmOperatorContext {
	var p = new(MathPmOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathPmOperator
	return p
}

func (*MathPmOperatorContext) IsMathPmOperatorContext() {}

func NewMathPmOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathPmOperatorContext {
	var p = new(MathPmOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathPmOperator

	return p
}

func (s *MathPmOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathPmOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(gengineParserPLUS, 0)
}

func (s *MathPmOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *MathPmOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathPmOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MathPmOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathPmOperator(s)
	}
}

func (s *MathPmOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathPmOperator(s)
	}
}

func (s *MathPmOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathPmOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) MathPmOperator() (localctx IMathPmOperatorContext) {
	localctx = NewMathPmOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gengineParserRULE_mathPmOperator)
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
		p.SetState(299)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserPLUS || _la == gengineParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IMathMdOperatorContext is an interface to support dynamic dispatch.
type IMathMdOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathMdOperatorContext differentiates from other interfaces.
	IsMathMdOperatorContext()
}

type MathMdOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathMdOperatorContext() *MathMdOperatorContext {
	var p = new(MathMdOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathMdOperator
	return p
}

func (*MathMdOperatorContext) IsMathMdOperatorContext() {}

func NewMathMdOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathMdOperatorContext {
	var p = new(MathMdOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathMdOperator

	return p
}

func (s *MathMdOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathMdOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(gengineParserMUL, 0)
}

func (s *MathMdOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(gengineParserDIV, 0)
}

func (s *MathMdOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathMdOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MathMdOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathMdOperator(s)
	}
}

func (s *MathMdOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathMdOperator(s)
	}
}

func (s *MathMdOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathMdOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) MathMdOperator() (localctx IMathMdOperatorContext) {
	localctx = NewMathMdOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gengineParserRULE_mathMdOperator)
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
		p.SetState(301)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserDIV || _la == gengineParserMUL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(gengineParserGT, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(gengineParserLT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(gengineParserGTE, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(gengineParserLTE, 0)
}

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gengineParserRULE_comparisonOperator)
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
		p.SetState(303)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserEQUALS) | (1 << gengineParserGT) | (1 << gengineParserLT) | (1 << gengineParserGTE) | (1 << gengineParserLTE) | (1 << gengineParserNOTEQUALS))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(gengineParserAND, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(gengineParserOR, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitLogicalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gengineParserRULE_logicalOperator)
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
		p.SetState(305)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserAND || _la == gengineParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IAssignOperatorContext is an interface to support dynamic dispatch.
type IAssignOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignOperatorContext differentiates from other interfaces.
	IsAssignOperatorContext()
}

type AssignOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignOperatorContext() *AssignOperatorContext {
	var p = new(AssignOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_assignOperator
	return p
}

func (*AssignOperatorContext) IsAssignOperatorContext() {}

func NewAssignOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignOperatorContext {
	var p = new(AssignOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_assignOperator

	return p
}

func (s *AssignOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(gengineParserASSIGN, 0)
}

func (s *AssignOperatorContext) SET() antlr.TerminalNode {
	return s.GetToken(gengineParserSET, 0)
}

func (s *AssignOperatorContext) PLUSEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserPLUSEQUAL, 0)
}

func (s *AssignOperatorContext) MINUSEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUSEQUAL, 0)
}

func (s *AssignOperatorContext) MULTIEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserMULTIEQUAL, 0)
}

func (s *AssignOperatorContext) DIVEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserDIVEQUAL, 0)
}

func (s *AssignOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AssignOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAssignOperator(s)
	}
}

func (s *AssignOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAssignOperator(s)
	}
}

func (s *AssignOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAssignOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) AssignOperator() (localctx IAssignOperatorContext) {
	localctx = NewAssignOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, gengineParserRULE_assignOperator)
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
		p.SetState(307)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (gengineParserASSIGN - 33)) | (1 << (gengineParserSET - 33)) | (1 << (gengineParserPLUSEQUAL - 33)) | (1 << (gengineParserMINUSEQUAL - 33)) | (1 << (gengineParserMULTIEQUAL - 33)) | (1 << (gengineParserDIVEQUAL - 33)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// INotOperatorContext is an interface to support dynamic dispatch.
type INotOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotOperatorContext differentiates from other interfaces.
	IsNotOperatorContext()
}

type NotOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotOperatorContext() *NotOperatorContext {
	var p = new(NotOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_notOperator
	return p
}

func (*NotOperatorContext) IsNotOperatorContext() {}

func NewNotOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotOperatorContext {
	var p = new(NotOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_notOperator

	return p
}

func (s *NotOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *NotOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(gengineParserNOT, 0)
}

func (s *NotOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NotOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterNotOperator(s)
	}
}

func (s *NotOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitNotOperator(s)
	}
}

func (s *NotOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitNotOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) NotOperator() (localctx INotOperatorContext) {
	localctx = NewNotOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, gengineParserRULE_notOperator)

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
		p.SetState(309)
		p.Match(gengineParserNOT)
	}



	return localctx
}


// IMapVarContext is an interface to support dynamic dispatch.
type IMapVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapVarContext differentiates from other interfaces.
	IsMapVarContext()
}

type MapVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapVarContext() *MapVarContext {
	var p = new(MapVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mapVar
	return p
}

func (*MapVarContext) IsMapVarContext() {}

func NewMapVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapVarContext {
	var p = new(MapVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mapVar

	return p
}

func (s *MapVarContext) GetParser() antlr.Parser { return s.parser }

func (s *MapVarContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *MapVarContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *MapVarContext) LSQARE() antlr.TerminalNode {
	return s.GetToken(gengineParserLSQARE, 0)
}

func (s *MapVarContext) RSQARE() antlr.TerminalNode {
	return s.GetToken(gengineParserRSQARE, 0)
}

func (s *MapVarContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *MapVarContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *MapVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MapVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMapVar(s)
	}
}

func (s *MapVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMapVar(s)
	}
}

func (s *MapVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMapVar(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) MapVar() (localctx IMapVarContext) {
	localctx = NewMapVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, gengineParserRULE_mapVar)

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
		p.SetState(311)
		p.Variable()
	}
	{
		p.SetState(312)
		p.Match(gengineParserLSQARE)
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserINT, gengineParserMINUS:
		{
			p.SetState(313)
			p.Integer()
		}


	case gengineParserDQUOTA_STRING:
		{
			p.SetState(314)
			p.StringLiteral()
		}


	case gengineParserSIMPLENAME, gengineParserDOTTEDNAME, gengineParserDOUBLEDOTTEDNAME:
		{
			p.SetState(315)
			p.Variable()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(318)
		p.Match(gengineParserRSQARE)
	}



	return localctx
}


// IAtNameContext is an interface to support dynamic dispatch.
type IAtNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtNameContext differentiates from other interfaces.
	IsAtNameContext()
}

type AtNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtNameContext() *AtNameContext {
	var p = new(AtNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atName
	return p
}

func (*AtNameContext) IsAtNameContext() {}

func NewAtNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtNameContext {
	var p = new(AtNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atName

	return p
}

func (s *AtNameContext) GetParser() antlr.Parser { return s.parser }
func (s *AtNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AtNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtName(s)
	}
}

func (s *AtNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtName(s)
	}
}

func (s *AtNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) AtName() (localctx IAtNameContext) {
	localctx = NewAtNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, gengineParserRULE_atName)

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
		p.SetState(320)
		p.Match(gengineParserT__1)
	}



	return localctx
}


// IAtIdContext is an interface to support dynamic dispatch.
type IAtIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtIdContext differentiates from other interfaces.
	IsAtIdContext()
}

type AtIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtIdContext() *AtIdContext {
	var p = new(AtIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atId
	return p
}

func (*AtIdContext) IsAtIdContext() {}

func NewAtIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtIdContext {
	var p = new(AtIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atId

	return p
}

func (s *AtIdContext) GetParser() antlr.Parser { return s.parser }
func (s *AtIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AtIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtId(s)
	}
}

func (s *AtIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtId(s)
	}
}

func (s *AtIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtId(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) AtId() (localctx IAtIdContext) {
	localctx = NewAtIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, gengineParserRULE_atId)

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
		p.SetState(322)
		p.Match(gengineParserT__2)
	}



	return localctx
}


// IAtDescContext is an interface to support dynamic dispatch.
type IAtDescContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtDescContext differentiates from other interfaces.
	IsAtDescContext()
}

type AtDescContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtDescContext() *AtDescContext {
	var p = new(AtDescContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atDesc
	return p
}

func (*AtDescContext) IsAtDescContext() {}

func NewAtDescContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtDescContext {
	var p = new(AtDescContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atDesc

	return p
}

func (s *AtDescContext) GetParser() antlr.Parser { return s.parser }
func (s *AtDescContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtDescContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AtDescContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtDesc(s)
	}
}

func (s *AtDescContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtDesc(s)
	}
}

func (s *AtDescContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtDesc(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) AtDesc() (localctx IAtDescContext) {
	localctx = NewAtDescContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, gengineParserRULE_atDesc)

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
		p.SetState(324)
		p.Match(gengineParserT__3)
	}



	return localctx
}


// IAtSalContext is an interface to support dynamic dispatch.
type IAtSalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtSalContext differentiates from other interfaces.
	IsAtSalContext()
}

type AtSalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtSalContext() *AtSalContext {
	var p = new(AtSalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atSal
	return p
}

func (*AtSalContext) IsAtSalContext() {}

func NewAtSalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtSalContext {
	var p = new(AtSalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atSal

	return p
}

func (s *AtSalContext) GetParser() antlr.Parser { return s.parser }
func (s *AtSalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtSalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AtSalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtSal(s)
	}
}

func (s *AtSalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtSal(s)
	}
}

func (s *AtSalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtSal(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) AtSal() (localctx IAtSalContext) {
	localctx = NewAtSalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, gengineParserRULE_atSal)

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
		p.SetState(326)
		p.Match(gengineParserT__4)
	}



	return localctx
}


func (p *gengineParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
			var t *ExpressionContext = nil
			if localctx != nil { t = localctx.(*ExpressionContext) }
			return p.Expression_Sempred(t, predIndex)

	case 10:
			var t *MathExpressionContext = nil
			if localctx != nil { t = localctx.(*MathExpressionContext) }
			return p.MathExpression_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gengineParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *gengineParser) MathExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

