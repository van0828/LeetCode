package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

//一个网站域名，如"discuss.leetcode.com"，包含了多个子域名。作为顶级域名，常用的有"com"，下一级则有"leetcode.com"，最低的一级为"discuss.leetcode.com"。当我们访问域名"discuss.leetcode.com"时，也同时访问了其父域名"leetcode.com"以及顶级域名 "com"。
//
//给定一个带访问次数和域名的组合，要求分别计算每个域名被访问的次数。其格式为访问次数+空格+地址，例如："9001 discuss.leetcode.com"。
//输入:
//["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
//输出:
//["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
//说明:
//按照假设，会访问"google.mail.com" 900次，"yahoo.com" 50次，"intel.mail.com" 1次，"wiki.org" 5次。
//而对于父域名，会访问"mail.com" 900+1 = 901次，"com" 900 + 50 + 1 = 951次，和 "org" 5 次。

// subdomain-visit-count
func subdomainVisits(cpdomains []string) []string {
	resMap := make(map[string]int, 0)
	for _, str := range cpdomains {
		// 拿到次数
		j := 0
		for ; j < len(str); j++ {
			if str[j] == 32 {
				break
			}
		}
		num, _ := strconv.Atoi(str[:j])

		// 域名拆解
		for k := len(str) - 1; k > j-1; k-- {
			// 空格或者"."后面的字符串都是一个完整的域名
			if str[k] == 46 || str[k] == 32 {
				domain := str[k+1:]
				if _, ok := resMap[domain]; !ok {
					resMap[domain] = num
				} else {
					resMap[domain] += num
				}
			}
		}
	}
	res := make([]string, 0)
	for domain, num := range resMap {
		res = append(res, fmt.Sprintf("%v %v", num, domain))
	}
	return res
}
