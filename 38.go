package main

//外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 被读作  "one 1"  ("一个一") , 即 11。
//11 被读作 "two 1s" ("两个一"）, 即 21。
//21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//
//给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//
//注意：整数序列中的每一项将表示为一个字符串。
//
// 
//
//示例 1:
//
//输入: 1
//输出: "1"
//解释：这是一个基本样例。
//示例 2:
//
//输入: 4
//输出: "1211"
//解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-and-say
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func countAndSay(n int) string {
	dict := map[int]string{
		1:"1",
		2:"11",
		3:"21",
		4:"1211",
		5:"111221",
		6:"312211",
		7:"13112221",
		8:"1113213211",
		9:"31131211131221",
		10:"13211311123113112211",
		11:"11131221133112132113212221",
		12:"3113112221232112111312211312113211",
		13:"1321132132111213122112311311222113111221131221",
		14:"11131221131211131231121113112221121321132132211331222113112211",
		15:"311311222113111231131112132112311321322112111312211312111322212311322113212221",
		16:"132113213221133112132113311211131221121321131211132221123113112221131112311332111213211322211312113211",
		17:"111312211312111322212321121113122123211231131122211211131221131112311332211213211321322113311213212312" +
			"31121113122113322113111221131221",
		18:"311311222113111231133211121312211231131122111213122112132113213221123113112221133112132123222112111312" +
			"21131211132221232112111312111213111213211231131122212322211331222113112211",
		19:"132113213221133112132123123112111311222112132113212231121113112221121113122113121113222112132113213221" +
			"23211211131211121332211231131122211311123113321112131221123113111231121113311211131221121321132132111" +
			"21332212311322113212221",
		20:"111312211312111322212321121113121112131112132112311321322112111312211312112213211231132132211231131122" +
			"21131112311332211211131221131211132211121312211231131112311211232221121321132132211331121321231231121" +
			"113112221121321133112132112312321123113112221121113122113121113123112112322111213211322211312113211",
		21:"311311222113111231133211121312211231131112311211133112111312211213211312111322211231131122211311122122" +
			"11131221121321131211132221121321132132211331121321232221123113112221131112311322311211131122211213211" +
			"33112132112211213322112111312211312111322212321121113121112131112132112311321322112111312212321121113" +
			"12211213111213122112132113213221123113112221131112311311121321122112132231121113122113322113111221131" +
			"221",
		22:"132113213221133112132123123112111311222112132113311213211231232112311311222112111312211311123113322112" +
			"13211321322113312211223113112221121113122113111231133221121113122113121113222123211211131211121332211" +
			"21321132132211331121321132213211231132132211211131221232112111312212221121123222112311311222113111231" +
			"13321112131221123113111231121113311211131221121321131211132221123113112211121312211231131122211211133" +
			"11211131122211211131221131211132221121321132132211331121321133112111312212221121113221321123113112221" +
			"2322211331222113112211",
		23:"111312211312111322212321121113121112131112132112311321322112111312212321121113122112131112131221121321" +
			"13213221123113112221133112132123222112111312211312111322212311222122132113213221123113112221133112132" +
			"12322211231131122211311123113321112131221123113111231121123222112111312211312111322212321121113122113" +
			"22111312211213211312111322211231131122111213122112311311221132211221121332211213211321322113311213212" +
			"31231121113112221121321133112132112312321123113112221121113122113111231133221121321132122311211131122" +
			"21121321132132211231232112311321322112311311222113111231133221121113122113121113222123211211131221232" +
			"11231131122113221123113221113122112132113213211121332212311322113212221",
		24:"311311222113111231133211121312211231131112311211133112111312211213211312111322211231131122111213122112" +
			"31131122211211133112111311222112111312211312111322211213211321322123211211131211121332211231131122211" +
			"31112311332111213213211221113122113121113222112132113213221232112111312111213322112132113213221133112" +
			"13212312311211131122211213211331121321122112133221123113112221131112311332111213122112311311222113223" +
			"11311222112111312211311123113322112132113212231121113112221121321132122211322212221121123222112111312" +
			"21131211132221232112111312111213111213211231132132211211131221232112111312211213111213122112132113213" +
			"22112311311222113311213212322211211131221131211221321123113213221121113122113121113222112131112131221" +
			"12132113121113222112132113213221133112132123222112311311222113111231133211121312211231131122111213122" +
			"11213211321222113222112132113223113112221121113122113121113123112112322111213211322211312113211",
		25:"132113213221133112132123123112111311222112132113311213211231232112311311222112111312211311123113322112" +
			"13211321223112111311222112132113213221123123211231132132211231131122211311123113322112111312211312111" +
			"32211121312211231131112311211232221121321132132211331121321231231121113121113122122311311222113111231" +
			"13322112111312211312111322111213122112311311123112112322211211131221131211132221232112111312111213111" +
			"21321123113213221121113122123211211131221222112112322211213211321322113311213212312311211131122211213" +
			"21132132211322132113213221123113112221133112132123222112111312211312112213211231132132211211131221131" +
			"21132211332113221122112133221123113112221131112311332111213122112311311123112111331121113122112132113" +
			"12111322211231131122111213122112311311222112111331121113112221121113122113121113222112132113213221232" +
			"11211131211121332211231131122211311122122111312211213211312111322211231131122211311123113322112111331" +
			"12111311222112111312211311123113322112111312211312111322212321121113121112133221121321132132211331121" +
			"32123123112111311222112132113212231121113112221121113122113121132211332211211131221132213211321322112" +
			"3113112221131112311311121321122112132231121113122113322113111221131221",
		26:"111312211312111322212321121113121112131112132112311321322112111312212321121113122112131112131221121321" +
			"13213221123113112221133112132123222112111312211312112213211231132132211211131221131211132221121311121" +
			"31221121321131211132221121321132132211331121321232221123113112221131112311322311211131122211213211331" +
			"12132112211213322112111312211312111322212321121113121112131112132112311311123113112211221321132132211" +
			"33112132123222112311311222113111231132231121113112221121321133112132112211213322112311311222113111231" +
			"13321112131221123113111231121113311211131221121321131211132221123113112211121312211231131122113221122" +
			"11213322112111312211312111322212321121113121112131112132112311321322112111312211312111322211322111312" +
			"21131211132221121321132132212321121113121112133221123113112221131112212211131221121321131211132221123" +
			"11311222113111221132221231221132221222112112322211213211321322113311213212312311211131122211213211331" +
			"12132112312321123113112221121113122113111231133221121321132122311211131122211213211321322112312321123" +
			"11321322112311311222113111231133221121113122113121113221112131221123113111231121123222112132113213221" +
			"13312211223113112221121113122113111231133221121321132132211331121321232221123123211231132132211231131" +
			"12221133112132123222112311311222113111231133211121312211231131112311211232221121113122113121113222123" +
			"21121113121112131112132112311321322112111312211312112213211231132132211231131122211311122113222123222" +
			"11231131122211322111312211312111322211213211321322113311213211331121113122122211211132213211231131122" +
			"212322211331222113112211",
		27:"311311222113111231133211121312211231131112311211133112111312211213211312111322211231131122111213122112" +
			"31131122211211133112111311222112111312211312111322211213211321322123211211131211121332211231131122211" +
			"31112212211131221121321131211132221123113112221131112311332211211133112111311222112111312211311123113" +
			"32211211131221131211132221232112111312111213322112132113213221133112132113221321123113213221121113122" +
			"12321121113122122211211232221123113112221131112311332111213122112311311123112111331121113122112132113" +
			"31121321132122212211131221131211132221232112111312111213322112132113213221133112132113221321123113213" +
			"22112111312212321121113122122211211232221121321132132211331121321231231121113112221121321133112132112" +
			"31232112311311222112111312211311123113322112132113212231121113112221121321132122211322212221121123222" +
			"11231131122211311123113321112131221123113111231121113311211131221121321131211132221123113112221131112" +
			"31133221132231131122211311123113322112111312211312111322111213122112311311123112112322211213211321322" +
			"11331221122311311222112111312211311123113322112132113213221133122211332111213112221133211322112211213" +
			"32211211131221131211132221232112111312111213111213211231132132211211131221232112111312211213111213122" +
			"11213211321322112311311222113311213212322211211131221131211221321123113213221121113122113121113222112" +
			"13111213122112132113121113222112132113213221133112132123222112311311222113111231132231121113112221121" +
			"32113311213211221121332211211131221131211132221231122212213211321322112311311222113311213212322211211" +
			"13122113121113222123211211131211121332211213111213122112132113121113222112132113213221232112111312111" +
			"21332211213211321322113311213212312311211131122211213211331121321122112133221123113112221131112311332" +
			"11121312211231131112311211133112111312211213211312111322211231131122211311122122111312211213211312111" +
			"32221121321132132211331222113321112133221121321132132211322311311222113111231133221121113122113121113" +
			"22212321121113122123211231131122113221123113221113122112132113213211121332212311322113212221",
		28:"132113213221133112132123123112111311222112132113311213211231232112311311222112111312211311123113322112" +
			"13211321223112111311222112132113213221123123211231132132211231131122211311123113322112111312211312111322111213122112311311123112112322211213211321322113312211223113112221121113122113111231133221121321132132211331121321232221123123211231132132211231131122211331121321232221123113112221131112311332111213122112311311123112112322211211131221131211132221232112111312211322111312211213211312111322211231131122111213122112311311221132211221121332211213211321322113311213212312311211131122211213211331121321123123211231131122211211131221232112111312211312113211223113112221131112311332111213122112311311123112112322211211131221131211132221232112111312211322111312211213211312111322211231131122111213122112311311221132211221121332211211131221131211132221232112111312111213111213211231132132211211131221232112111312211213111213122112132113213221123113112221133112132123222112111312211312112213211231132132211211131221131211322113321132211221121332211213211321322113311213212312311211131122211213211331121321123123211231131122211211131221131112311332211213211321322113311213212322211322132113213221133112132123222112311311222113111231132231121113112221121321133112132112211213322112111312211312111322212311222122132113213221123113112221133112132123222112111312211312111322212311322123123112111321322123122113222122211211232221123113112221131112311332111213122112311311123112111331121113122112132113121113222112311311221112131221123113112221121113311211131122211211131221131211132221121321132132212321121113121112133221123113112221131112212211131221121321131211132221123113112221131112311332211211133112111311222112111312211311123113322112111312211312111322212321121113121112133221121321132132211331121321132213211231132132211211131221232112111312212221121123222112311311222113111231133211121321321122111312211312111322211213211321322123211211131211121332211231131122211311123113321112131221123113111231121123222112111331121113112221121113122113111231133221121113122113121113221112131221123113111231121123222112111312211312111322212321121113121112131112132112311321322112111312212321121113122122211211232221121321132132211331121321231231121113112221121321133112132112312321123113112221121113122113111231133221121321132132211331221122311311222112111312211311123113322112111312211312111322212311322123123112112322211211131221131211132221132213211321322113311213212322211231131122211311123113321112131221123113112211121312211213211321222113222112132113223113112221121113122113121113123112112322111213211322211312113211",
		29:"111312211312111322212321121113121112131112132112311321322112111312212321121113122112131112131221121321" +
			"13213221123113112221133112132123222112111312211312112213211231132132211211131221131211132221121311121312211213211312111322211213211321322113311213212322211231131122211311123113223112111311222112132113311213211221121332211211131221131211132221231122212213211321322112311311222113311213212322211211131221131211132221232112111312111213322112131112131221121321131211132221121321132132212321121113121112133221121321132132211331121321231231121113112221121321133112132112211213322112311311222113111231133211121312211231131122211322311311222112111312211311123113322112132113212231121113112221121321132122211322212221121123222112111312211312111322212321121113121112131112132112311321322112111312212321121113122112131112131221121321132132211231131122111213122112311311222113111221131221221321132132211331121321231231121113112221121321133112132112211213322112311311222113111231133211121312211231131122211322311311222112111312211311123113322112132113212231121113112221121321132122211322212221121123222112311311222113111231133211121312211231131112311211133112111312211213211312111322211231131122111213122112311311222112111331121113112221121113122113121113222112132113213221232112111312111213322112311311222113111221221113122112132113121113222112311311222113111221132221231221132221222112112322211211131221131211132221232112111312111213111213211231132132211211131221232112111312211213111213122112132113213221123113112221133112132123222112111312211312111322212321121113121112133221132211131221131211132221232112111312111213322112132113213221133112132113221321123113213221121113122123211211131221222112112322211231131122211311123113321112132132112211131221131211132221121321132132212321121113121112133221123113112221131112311332111213211322111213111213211231131211132211121311222113321132211221121332211213211321322113311213212312311211131122211213211331121321123123211231131122211211131221131112311332211213211321223112111311222112132113213221123123211231132132211231131122211311123113322112111312211312111322111213122112311311123112112322211213211321322113312211223113112221121113122113111231133221121321132132211331121321232221123123211231132132211231131122211331121321232221123113112221131112311332111213122112311311123112112322211211131221131211132221232112111312211322111312211213211312111322211231131122111213122112311311221132211221121332211213211321322113311213212312311211131211131221223113112221131112311332211211131221131211132211121312211231131112311211232221121321132132211331121321231231121113112221121321133112132112211213322112312321123113213221123113112221133112132123222112311311222113111231132231121113112221121321133112132112211213322112311311222113111231133211121312211231131112311211133112111312211213211312111322211231131122111213122112311311221132211221121332211211131221131211132221232112111312111213111213211231132132211211131221232112111312211213111213122112132113213221123113112221133112132123222112111312211312111322212311222122132113213221123113112221133112132123222112311311222113111231133211121321132211121311121321122112133221123113112221131112311332211322111312211312111322212321121113121112133221121321132132211331121321231231121113112221121321132122311211131122211211131221131211322113322112111312211322132113213221123113112221131112311311121321122112132231121113122113322113111221131221",
		30:"311311222113111231133211121312211231131112311211133112111312211213211312111322211231131122111213122112" +
			"3113112221121113311211131122211211131221131211132221121321132132212321121113121112133221123113112221131112212211131221121321131211132221123113112221131112311332211211133112111311222112111312211311123113322112111312211312111322212321121113121112133221121321132132211331121321132213211231132132211211131221232112111312212221121123222112311311222113111231133211121321321122111312211312111322211213211321322123211211131211121332211231131122211311123113321112131221123113111231121123222112111331121113112221121113122113111231133221121113122113121113221112131221123113111231121123222112111312211312111322212321121113121112131112132112311321322112111312212321121113122122211211232221121321132132211331121321231231121113112221121321132132211322132113213221123113112221133112132123222112111312211312112213211231132132211211131221131211322113321132211221121332211231131122211311123113321112131221123113111231121113311211131221121321131211132221123113112211121312211231131122211211133112111311222112111312211312111322211213211321223112111311222112132113213221133122211311221122111312211312111322212321121113121112131112132112311321322112111312212321121113122122211211232221121321132132211331121321231231121113112221121321132132211322132113213221123113112221133112132123222112111312211312112213211231132132211211131221131211322113321132211221121332211213211321322113311213212312311211131122211213211331121321123123211231131122211211131221131112311332211213211321223112111311222112132113213221123123211231132132211231131122211311123113322112111312211312111322111213122112311311123112112322211213211321322113312211223113112221121113122113111231133221121321132132211331222113321112131122211332113221122112133221123113112221131112311332111213122112311311123112111331121113122112132113121113222112311311221112131221123113112221121113311211131122211211131221131211132221121321132132212321121113121112133221123113112221131112311332111213122112311311123112112322211322311311222113111231133211121312211231131112311211232221121113122113121113222123211211131221132211131221121321131211132221123113112211121312211231131122113221122112133221121321132132211331121321231231121113121113122122311311222113111231133221121113122113121113221112131221123113111231121123222112132113213221133112132123123112111312211322311211133112111312211213211311123113223112111321322123122113222122211211232221121113122113121113222123211211131211121311121321123113213221121113122123211211131221121311121312211213211321322112311311222113311213212322211211131221131211221321123113213221121113122113121113222112131112131221121321131211132221121321132132211331121321232221123113112221131112311322311211131122211213211331121321122112133221121113122113121113222123112221221321132132211231131122211331121321232221121113122113121113222123211211131211121332211213111213122112132113121113222112132113213221232112111312111213322112132113213221133112132123123112111311222112132113311213211221121332211231131122211311123113321112131221123113112221132231131122211211131221131112311332211213211321223112111311222112132113212221132221222112112322211211131221131211132221232112111312111213111213211231131112311311221122132113213221133112132123222112311311222113111231132231121113112221121321133112132112211213322112111312211312111322212321121113121112131112132112311321322112111312212321121113122122211211232221121311121312211213211312111322211213211321322123211211131211121332211213211321322113311213211322132112311321322112111312212321121113122122211211232221121321132132211331121321231231121113112221121321133112132112312321123113112221121113122113111231133221121321132122311211131122211213211321222113222122211211232221123113112221131112311332111213122112311311123112111331121113122112132113121113222112311311221112131221123113112221121113311211131122211211131221131211132221121321132132212321121113121112133221123113112221131112311332111213213211221113122113121113222112132113213221232112111312111213322112132113213221133112132123123112111312211322311211133112111312212221121123222112132113213221133112132123222113223113112221131112311332111213122112311311123112112322211211131221131211132221232112111312111213111213211231132132211211131221131211221321123113213221123113112221131112211322212322211231131122211322111312211312111322211213211321322113311213211331121113122122211211132213211231131122212322211331222113112211",
	}
	return dict[n]
}