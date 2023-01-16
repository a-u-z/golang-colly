package main

import (
	"log"

	"github.com/gocolly/colly"
)

// 基本要素，以下三個都要以同一個 *colly.Collector 為基礎
// 1. NewCollector
// 2. OnHTML/OnXML
// 3. OnVisit

func Co(url string) {
	// 實例預設的 collector，有以下參數可以設定。 ref: https://www.readfog.com/a/1644908070754684928
	c := colly.NewCollector(
	// colly.AllowedDomains(url), //只能造訪這些網站，白名單作用
	// colly.AllowURLRevisit(),              // 允許對同一 URL 進行多次下載
	// colly.Async(true), // 設置爲異步請求，設定此參數要注意併發，看是否因為還沒收到 response 程式就執行完畢
	// colly.Debugger(&debug.LogDebugger{}), // 開啓debug
	// colly.MaxDepth(2),                    // 爬取頁面深度,最多爲兩層
	// colly.MaxBodySize(1024*1024),         // 響應正文最大字節數
	// colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"),
	// colly.IgnoreRobotsTxt(), //忽略目標機器中的`robots.txt`聲明
	)

	// OnHTML 的第一個參數使用 goquerySelector
	// 如要抓取 id => #main-container
	// 如要抓取 tag 標籤 => title, div, span, a
	// 如要抓取擁有 class 的 tag => div[class], a[href]
	// 其他 goquerySelector 寫法詳情請參考：golang goquery selector(选择器) 示例大全(https://www.flysnow.org/2018/01/20/golang-goquery-examples-selector)

	// 收到 html 後，可以依據第一個參數，執行 callback func
	c.OnHTML("div[class=r-ent] > div[class=title] > a", func(e *colly.HTMLElement) {
		log.Printf("here is e:%+v", e.Text)
	})

	c.Visit(url)
}

// ref: https://ithelp.ithome.com.tw/articles/10247126
