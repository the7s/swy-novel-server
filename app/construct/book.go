package construct

import "github.com/the7s/swy-novel-server/app/model"

var CategoryConst = map[int]model.Category{
	0:  {Id: 0, Tag: "all", Name: "全部"},
	1:  {Id: 1, Tag: "all/chanId21", Name: "玄幻"},
	2:  {Id: 2, Tag: "all/chanId1", Name: "奇幻"},
	3:  {Id: 3, Tag: "all/chanId2", Name: "武侠"},
	4:  {Id: 4, Tag: "all/chanId22", Name: "仙侠"},
	5:  {Id: 5, Tag: "all/chanId4", Name: "都市"},
	6:  {Id: 6, Tag: "all/chanId15", Name: "现实"},
	7:  {Id: 7, Tag: "all/chanId6", Name: "军事"},
	8:  {Id: 8, Tag: "all/chanId5", Name: "历史"},
	9:  {Id: 9, Tag: "all/chanId7", Name: "游戏"},
	10: {Id: 10, Tag: "all/chanId8", Name: "体育"},
	11: {Id: 11, Tag: "all/chanId9", Name: "科幻"},
	12: {Id: 12, Tag: "all/chanId20109", Name: "诸天无限"},
	13: {Id: 13, Tag: "all/chanId10", Name: "悬疑"},
	14: {Id: 14, Tag: "all/chanId12", Name: "轻小说"},
	15: {Id: 15, Tag: "all/chanId20076", Name: "短片"},
}
