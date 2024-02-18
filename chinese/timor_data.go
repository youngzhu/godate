package chinese

// 没法共享配置文件
// 先放到变量里，后面有更优雅的方式再调整
var timorData = map[string]string{
	"2024": timorData2024,
}

var timorData2024 = `{
  "code": 0,
  "holidays": [
    {
      "holiday": true,
      "name": "元旦",
      "wage": 3,
      "date": "2024-01-01",
      "rest": 1
    },
    {
      "holiday": false,
      "name": "春节前补班",
      "wage": 1,
      "after": false,
      "target": "春节",
      "date": "2024-02-04",
      "rest": 10
    },
    {
      "holiday": true,
      "name": "初一",
      "wage": 3,
      "date": "2024-02-10",
      "rest": 16
    },
    {
      "holiday": true,
      "name": "初二",
      "wage": 3,
      "date": "2024-02-11",
      "rest": 1
    },
    {
      "holiday": true,
      "name": "初三",
      "wage": 3,
      "date": "2024-02-12"
    },
    {
      "holiday": true,
      "name": "初四",
      "wage": 2,
      "date": "2024-02-13"
    },
    {
      "holiday": true,
      "name": "初五",
      "wage": 2,
      "date": "2024-02-14"
    },
    {
      "holiday": true,
      "name": "初六",
      "wage": 2,
      "date": "2024-02-15"
    },
    {
      "holiday": true,
      "name": "初七",
      "wage": 2,
      "date": "2024-02-16"
    },
    {
      "holiday": true,
      "name": "初八",
      "wage": 2,
      "date": "2024-02-17"
    },
    {
      "holiday": false,
      "name": "春节后补班",
      "wage": 1,
      "after": true,
      "target": "春节",
      "date": "2024-02-18"
    },
    {
      "holiday": true,
      "name": "清明节",
      "wage": 3,
      "date": "2024-04-04",
      "rest": 46
    },
    {
      "holiday": true,
      "name": "清明节",
      "wage": 2,
      "date": "2024-04-05"
    },
    {
      "holiday": true,
      "name": "清明节",
      "wage": 2,
      "date": "2024-04-06"
    },
    {
      "holiday": false,
      "name": "清明节后补班",
      "wage": 1,
      "target": "清明节",
      "after": true,
      "date": "2024-04-07"
    },
    {
      "holiday": false,
      "name": "劳动节前补班",
      "wage": 1,
      "target": "劳动节",
      "after": false,
      "date": "2024-04-28"
    },
    {
      "holiday": true,
      "name": "劳动节",
      "wage": 3,
      "date": "2024-05-01"
    },
    {
      "holiday": true,
      "name": "劳动节",
      "wage": 2,
      "date": "2024-05-02",
      "rest": 1
    },
    {
      "holiday": true,
      "name": "劳动节",
      "wage": 3,
      "date": "2024-05-03"
    },
    {
      "holiday": true,
      "name": "劳动节",
      "wage": 3,
      "date": "2024-05-04"
    },
    {
      "holiday": true,
      "name": "劳动节",
      "wage": 3,
      "date": "2024-05-05"
    },
    {
      "holiday": false,
      "name": "劳动节后补班",
      "after": true,
      "wage": 1,
      "target": "劳动节",
      "date": "2024-05-11"
    },
    {
      "holiday": true,
      "name": "端午节",
      "wage": 2,
      "date": "2024-06-08"
    },
    {
      "holiday": true,
      "name": "端午节",
      "wage": 2,
      "date": "2024-06-09"
    },
    {
      "holiday": true,
      "name": "端午节",
      "wage": 3,
      "date": "2024-06-10"
    },
    {
      "holiday": false,
      "name": "中秋节前补班",
      "after": false,
      "wage": 1,
      "target": "中秋节",
      "date": "2024-09-14",
      "rest": 96
    },
    {
      "holiday": true,
      "name": "中秋节",
      "wage": 2,
      "date": "2024-09-15",
      "rest": 97
    },
    {
      "holiday": true,
      "name": "中秋节",
      "wage": 2,
      "date": "2024-09-16"
    },
    {
      "holiday": true,
      "name": "中秋节",
      "wage": 3,
      "date": "2024-09-17"
    },
    {
      "holiday": false,
      "name": "国庆节前补班",
      "after": false,
      "wage": 1,
      "target": "国庆节",
      "date": "2024-09-29"
    },
    {
      "holiday": true,
      "name": "国庆节",
      "wage": 3,
      "date": "2024-10-01"
    },
    {
      "holiday": true,
      "name": "国庆节",
      "wage": 3,
      "date": "2024-10-02",
      "rest": 1
    },
    {
      "holiday": true,
      "name": "国庆节",
      "wage": 3,
      "date": "2024-10-03"
    },
    {
      "holiday": true,
      "name": "国庆节",
      "wage": 2,
      "date": "2024-10-04"
    },
    {
      "holiday": true,
      "name": "国庆节",
      "wage": 2,
      "date": "2024-10-05"
    },
    {
      "holiday": true,
      "name": "国庆节",
      "wage": 2,
      "date": "2024-10-06",
      "rest": 1
    },
    {
      "holiday": true,
      "name": "国庆节",
      "wage": 2,
      "date": "2024-10-07",
      "rest": 1
    },
    {
      "holiday": false,
      "after": true,
      "wage": 1,
      "name": "国庆节后补班",
      "target": "国庆节",
      "date": "2024-10-12"
    }
  ]
}
`
