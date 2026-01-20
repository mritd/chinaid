// pinyin_map.go
package metadata

// PinyinMap 汉字到拼音的映射（覆盖所有姓氏和常用名字用字）
var PinyinMap = map[rune]string{
	// === 常见姓氏 ===
	'李': "li", '王': "wang", '张': "zhang", '刘': "liu", '陈': "chen",
	'杨': "yang", '黄': "huang", '赵': "zhao", '周': "zhou", '吴': "wu",
	'徐': "xu", '孙': "sun", '朱': "zhu", '马': "ma", '胡': "hu",
	'郭': "guo", '林': "lin", '何': "he", '高': "gao", '梁': "liang",
	'郑': "zheng", '罗': "luo", '宋': "song", '谢': "xie", '唐': "tang",
	'韦': "wei", '曹': "cao", '许': "xu", '邓': "deng", '萧': "xiao",
	'冯': "feng", '曾': "zeng", '程': "cheng", '蔡': "cai", '彭': "peng",
	'潘': "pan", '袁': "yuan", '于': "yu", '董': "dong", '余': "yu",
	'苏': "su", '叶': "ye", '吕': "lv", '魏': "wei", '蒋': "jiang",
	'田': "tian", '杜': "du", '丁': "ding", '沈': "shen", '姜': "jiang",
	'范': "fan", '江': "jiang", '傅': "fu", '钟': "zhong", '卢': "lu",
	'汪': "wang", '戴': "dai", '崔': "cui", '任': "ren", '陆': "lu",
	'廖': "liao", '姚': "yao", '方': "fang", '金': "jin", '邱': "qiu",
	'夏': "xia", '谭': "tan", '韩': "han", '贾': "jia", '邹': "zou",
	'石': "shi", '熊': "xiong", '孟': "meng", '秦': "qin", '阎': "yan",
	'薛': "xue", '侯': "hou", '雷': "lei", '白': "bai", '龙': "long",
	'段': "duan", '郝': "hao", '孔': "kong", '邵': "shao", '史': "shi",
	'毛': "mao", '常': "chang", '万': "wan", '顾': "gu", '赖': "lai",
	'武': "wu", '康': "kang", '贺': "he", '严': "yan", '尹': "yin",
	'钱': "qian", '施': "shi", '牛': "niu", '洪': "hong", '龚': "gong",

	// === 复姓 ===
	'欧': "ou", '阳': "yang",
	'司': "si", '上': "shang", '官': "guan",
	'诸': "zhu", '葛': "ge",
	'东': "dong",
	'公': "gong",
	'慕': "mu", '容': "rong",
	'甫': "fu",
	'端': "duan", '木': "mu",
	'南': "nan", '宫': "gong",

	// === 常用名字用字 - 男 ===
	'伟': "wei", '强': "qiang", '磊': "lei", '军': "jun", '勇': "yong",
	'杰': "jie", '涛': "tao", '明': "ming", '超': "chao", '华': "hua",
	'刚': "gang", '辉': "hui", '波': "bo", '斌': "bin", '鹏': "peng",
	'飞': "fei", '峰': "feng", '毅': "yi", '威': "wei", '浩': "hao",
	'亮': "liang", '健': "jian", '宁': "ning", '俊': "jun", '凯': "kai",
	'兵': "bing", '锋': "feng", '翔': "xiang", '宇': "yu",
	'轩': "xuan", '豪': "hao", '天': "tian", '佑': "you", '航': "hang",
	'晨': "chen", '曦': "xi", '然': "ran", '睿': "rui", '博': "bo",
	'坤': "kun", '昊': "hao", '铭': "ming", '泽': "ze", '洋': "yang",
	'森': "sen", '翰': "han", '达': "da", '栋': "dong", '政': "zheng",
	'帅': "shuai", '哲': "zhe", '瑞': "rui", '旭': "xu", '彬': "bin",
	'鸿': "hong", '昌': "chang", '松': "song", '楠': "nan", '鑫': "xin",

	// === 常用名字用字 - 女 ===
	'芳': "fang", '娜': "na", '敏': "min", '静': "jing", '丽': "li",
	'艳': "yan", '霞': "xia", '燕': "yan", '玲': "ling", '娟': "juan",
	'萍': "ping", '红': "hong", '梅': "mei", '琴': "qin", '英': "ying",
	'慧': "hui", '莉': "li", '蓉': "rong", '洁': "jie", '颖': "ying",
	'婷': "ting", '雪': "xue", '琳': "lin", '璐': "lu", '倩': "qian",
	'薇': "wei", '妍': "yan", '瑶': "yao", '蕾': "lei", '涵': "han",
	'萱': "xuan", '琪': "qi", '欣': "xin", '怡': "yi", '悦': "yue",
	'诗': "shi", '语': "yu", '嫣': "yan", '若': "ruo", '思': "si",
	'婕': "jie", '茜': "qian", '岚': "lan", '媛': "yuan", '菲': "fei",
	'蓓': "bei", '晶': "jing", '莹': "ying", '蕊': "rui", '露': "lu",
	'萌': "meng", '珊': "shan", '瑾': "jin", '韵': "yun", '雅': "ya",
	'曼': "man", '妮': "ni", '彤': "tong", '晴': "qing", '溪': "xi",

	// === 常用字 ===
	'国': "guo", '建': "jian", '文': "wen", '永': "yong", '海': "hai",
	'子': "zi", '小': "xiao", '大': "da", '中': "zhong", '新': "xin",
	'志': "zhi", '学': "xue", '成': "cheng", '平': "ping", '春': "chun",
	'秀': "xiu", '玉': "yu", '桂': "gui", '淑': "shu", '紫': "zi",
	'梦': "meng", '雨': "yu", '月': "yue", '心': "xin", '美': "mei",
	'爱': "ai", '兰': "lan", '珍': "zhen", '珠': "zhu", '云': "yun",
	'世': "shi", '家': "jia", '德': "de", '安': "an", '富': "fu",
	'贵': "gui", '荣': "rong", '福': "fu", '禄': "lu", '寿': "shou",
	'喜': "xi", '财': "cai", '吉': "ji", '祥': "xiang", '庆': "qing",
	'乐': "le", '嘉': "jia", '佳': "jia", '宝': "bao", '贝': "bei",
	'正': "zheng", '清': "qing", '远': "yuan", '行': "xing", '道': "dao",
}
