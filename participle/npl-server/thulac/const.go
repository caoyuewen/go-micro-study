package thulac

// WordType WordType
type WordType int32

// Word Type
const (
	WtN  WordType = 10 // 名词
	WtNP WordType = 11 // 人名
	WtNS WordType = 12 // 地名
	WtNI WordType = 13 // 机构名
	WtNZ WordType = 14 // 其他专名

	WtV  WordType = 20 // 动词
	WtVM WordType = 21 // 能愿动词
	WtVD WordType = 22 // 趋向动词

	WtM  WordType = 30 // 数词
	WtMQ WordType = 31 // 数量词
	WtQ  WordType = 32 // 量词

	WtT WordType = 40 // 时间词
	WtF WordType = 41 // 方位词
	WtS WordType = 42 // 处所词
	WtA WordType = 43 // 形容词
	WtD WordType = 44 // 副词
	WtH WordType = 45 // 前接成分
	WtK WordType = 46 // 后接成分
	WtI WordType = 47 // 习语
	WtJ WordType = 48 // 简称
	WtR WordType = 49 // 代词

	WtC WordType = 50 // 连词
	WtP WordType = 51 // 介词
	WtU WordType = 52 // 助词
	WtY WordType = 53 // 语气助词
	WtE WordType = 54 // 叹词
	WtO WordType = 55 // 拟声词
	WtG WordType = 56 // 语素

	WtW WordType = 98 // 标点
	WtX WordType = 99 // 其它
)
