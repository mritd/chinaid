package metadata

// Province 省份信息
type Province struct {
	Name   string // 省份全称："北京市"
	Short  string // 简称："北京"
	Code   string // 省级代码："11"
	Cities []City // 下属城市/区县
}

// City 城市/区县信息
type City struct {
	Name      string   // 名称："朝阳区"
	AreaCodes []string // 6位地区码
}

// Provinces 全国省份数据
var Provinces = []Province{
	{
		Name: "北京市", Short: "北京", Code: "11",
		Cities: []City{
			{Name: "东城区", AreaCodes: []string{"110101"}},
			{Name: "西城区", AreaCodes: []string{"110102"}},
			{Name: "朝阳区", AreaCodes: []string{"110105"}},
			{Name: "丰台区", AreaCodes: []string{"110106"}},
			{Name: "石景山区", AreaCodes: []string{"110107"}},
			{Name: "海淀区", AreaCodes: []string{"110108"}},
			{Name: "顺义区", AreaCodes: []string{"110113"}},
			{Name: "通州区", AreaCodes: []string{"110112"}},
			{Name: "大兴区", AreaCodes: []string{"110115"}},
			{Name: "昌平区", AreaCodes: []string{"110114"}},
		},
	},
	{
		Name: "天津市", Short: "天津", Code: "12",
		Cities: []City{
			{Name: "和平区", AreaCodes: []string{"120101"}},
			{Name: "河东区", AreaCodes: []string{"120102"}},
			{Name: "河西区", AreaCodes: []string{"120103"}},
			{Name: "南开区", AreaCodes: []string{"120104"}},
			{Name: "河北区", AreaCodes: []string{"120105"}},
			{Name: "滨海新区", AreaCodes: []string{"120116"}},
			{Name: "东丽区", AreaCodes: []string{"120110"}},
			{Name: "西青区", AreaCodes: []string{"120111"}},
		},
	},
	{
		Name: "河北省", Short: "河北", Code: "13",
		Cities: []City{
			{Name: "石家庄市", AreaCodes: []string{"130100", "130102"}},
			{Name: "唐山市", AreaCodes: []string{"130200", "130202"}},
			{Name: "秦皇岛市", AreaCodes: []string{"130300"}},
			{Name: "邯郸市", AreaCodes: []string{"130400"}},
			{Name: "保定市", AreaCodes: []string{"130600"}},
			{Name: "张家口市", AreaCodes: []string{"130700"}},
			{Name: "廊坊市", AreaCodes: []string{"131000"}},
		},
	},
	{
		Name: "山西省", Short: "山西", Code: "14",
		Cities: []City{
			{Name: "太原市", AreaCodes: []string{"140100", "140105"}},
			{Name: "大同市", AreaCodes: []string{"140200"}},
			{Name: "阳泉市", AreaCodes: []string{"140300"}},
			{Name: "长治市", AreaCodes: []string{"140400"}},
			{Name: "晋城市", AreaCodes: []string{"140500"}},
			{Name: "运城市", AreaCodes: []string{"140800"}},
		},
	},
	{
		Name: "内蒙古自治区", Short: "内蒙古", Code: "15",
		Cities: []City{
			{Name: "呼和浩特市", AreaCodes: []string{"150100", "150102"}},
			{Name: "包头市", AreaCodes: []string{"150200"}},
			{Name: "乌海市", AreaCodes: []string{"150300"}},
			{Name: "赤峰市", AreaCodes: []string{"150400"}},
			{Name: "鄂尔多斯市", AreaCodes: []string{"150600"}},
		},
	},
	{
		Name: "辽宁省", Short: "辽宁", Code: "21",
		Cities: []City{
			{Name: "沈阳市", AreaCodes: []string{"210100", "210102"}},
			{Name: "大连市", AreaCodes: []string{"210200", "210202"}},
			{Name: "鞍山市", AreaCodes: []string{"210300"}},
			{Name: "抚顺市", AreaCodes: []string{"210400"}},
			{Name: "本溪市", AreaCodes: []string{"210500"}},
			{Name: "丹东市", AreaCodes: []string{"210600"}},
		},
	},
	{
		Name: "吉林省", Short: "吉林", Code: "22",
		Cities: []City{
			{Name: "长春市", AreaCodes: []string{"220100", "220102"}},
			{Name: "吉林市", AreaCodes: []string{"220200"}},
			{Name: "四平市", AreaCodes: []string{"220300"}},
			{Name: "辽源市", AreaCodes: []string{"220400"}},
			{Name: "通化市", AreaCodes: []string{"220500"}},
		},
	},
	{
		Name: "黑龙江省", Short: "黑龙江", Code: "23",
		Cities: []City{
			{Name: "哈尔滨市", AreaCodes: []string{"230100", "230102"}},
			{Name: "齐齐哈尔市", AreaCodes: []string{"230200"}},
			{Name: "牡丹江市", AreaCodes: []string{"231000"}},
			{Name: "佳木斯市", AreaCodes: []string{"230800"}},
			{Name: "大庆市", AreaCodes: []string{"230600"}},
		},
	},
	{
		Name: "上海市", Short: "上海", Code: "31",
		Cities: []City{
			{Name: "黄浦区", AreaCodes: []string{"310101"}},
			{Name: "徐汇区", AreaCodes: []string{"310104"}},
			{Name: "长宁区", AreaCodes: []string{"310105"}},
			{Name: "静安区", AreaCodes: []string{"310106"}},
			{Name: "普陀区", AreaCodes: []string{"310107"}},
			{Name: "虹口区", AreaCodes: []string{"310109"}},
			{Name: "杨浦区", AreaCodes: []string{"310110"}},
			{Name: "浦东新区", AreaCodes: []string{"310115"}},
			{Name: "闵行区", AreaCodes: []string{"310112"}},
			{Name: "宝山区", AreaCodes: []string{"310113"}},
		},
	},
	{
		Name: "江苏省", Short: "江苏", Code: "32",
		Cities: []City{
			{Name: "南京市", AreaCodes: []string{"320100", "320102"}},
			{Name: "无锡市", AreaCodes: []string{"320200"}},
			{Name: "徐州市", AreaCodes: []string{"320300"}},
			{Name: "常州市", AreaCodes: []string{"320400"}},
			{Name: "苏州市", AreaCodes: []string{"320500", "320505"}},
			{Name: "南通市", AreaCodes: []string{"320600"}},
			{Name: "扬州市", AreaCodes: []string{"321000"}},
		},
	},
	{
		Name: "浙江省", Short: "浙江", Code: "33",
		Cities: []City{
			{Name: "杭州市", AreaCodes: []string{"330100", "330102"}},
			{Name: "宁波市", AreaCodes: []string{"330200"}},
			{Name: "温州市", AreaCodes: []string{"330300"}},
			{Name: "嘉兴市", AreaCodes: []string{"330400"}},
			{Name: "湖州市", AreaCodes: []string{"330500"}},
			{Name: "绍兴市", AreaCodes: []string{"330600"}},
			{Name: "金华市", AreaCodes: []string{"330700"}},
		},
	},
	{
		Name: "安徽省", Short: "安徽", Code: "34",
		Cities: []City{
			{Name: "合肥市", AreaCodes: []string{"340100", "340102"}},
			{Name: "芜湖市", AreaCodes: []string{"340200"}},
			{Name: "蚌埠市", AreaCodes: []string{"340300"}},
			{Name: "淮南市", AreaCodes: []string{"340400"}},
			{Name: "马鞍山市", AreaCodes: []string{"340500"}},
			{Name: "安庆市", AreaCodes: []string{"340800"}},
		},
	},
	{
		Name: "福建省", Short: "福建", Code: "35",
		Cities: []City{
			{Name: "福州市", AreaCodes: []string{"350100", "350102"}},
			{Name: "厦门市", AreaCodes: []string{"350200", "350203"}},
			{Name: "莆田市", AreaCodes: []string{"350300"}},
			{Name: "泉州市", AreaCodes: []string{"350500"}},
			{Name: "漳州市", AreaCodes: []string{"350600"}},
		},
	},
	{
		Name: "江西省", Short: "江西", Code: "36",
		Cities: []City{
			{Name: "南昌市", AreaCodes: []string{"360100", "360102"}},
			{Name: "景德镇市", AreaCodes: []string{"360200"}},
			{Name: "萍乡市", AreaCodes: []string{"360300"}},
			{Name: "九江市", AreaCodes: []string{"360400"}},
			{Name: "赣州市", AreaCodes: []string{"360700"}},
		},
	},
	{
		Name: "山东省", Short: "山东", Code: "37",
		Cities: []City{
			{Name: "济南市", AreaCodes: []string{"370100", "370102"}},
			{Name: "青岛市", AreaCodes: []string{"370200", "370202"}},
			{Name: "淄博市", AreaCodes: []string{"370300"}},
			{Name: "枣庄市", AreaCodes: []string{"370400"}},
			{Name: "东营市", AreaCodes: []string{"370500"}},
			{Name: "烟台市", AreaCodes: []string{"370600"}},
			{Name: "潍坊市", AreaCodes: []string{"370700"}},
			{Name: "威海市", AreaCodes: []string{"371000"}},
		},
	},
	{
		Name: "河南省", Short: "河南", Code: "41",
		Cities: []City{
			{Name: "郑州市", AreaCodes: []string{"410100", "410102"}},
			{Name: "开封市", AreaCodes: []string{"410200"}},
			{Name: "洛阳市", AreaCodes: []string{"410300"}},
			{Name: "平顶山市", AreaCodes: []string{"410400"}},
			{Name: "安阳市", AreaCodes: []string{"410500"}},
			{Name: "新乡市", AreaCodes: []string{"410700"}},
		},
	},
	{
		Name: "湖北省", Short: "湖北", Code: "42",
		Cities: []City{
			{Name: "武汉市", AreaCodes: []string{"420100", "420102"}},
			{Name: "黄石市", AreaCodes: []string{"420200"}},
			{Name: "十堰市", AreaCodes: []string{"420300"}},
			{Name: "宜昌市", AreaCodes: []string{"420500"}},
			{Name: "襄阳市", AreaCodes: []string{"420600"}},
			{Name: "荆州市", AreaCodes: []string{"421000"}},
		},
	},
	{
		Name: "湖南省", Short: "湖南", Code: "43",
		Cities: []City{
			{Name: "长沙市", AreaCodes: []string{"430100", "430102"}},
			{Name: "株洲市", AreaCodes: []string{"430200"}},
			{Name: "湘潭市", AreaCodes: []string{"430300"}},
			{Name: "衡阳市", AreaCodes: []string{"430400"}},
			{Name: "邵阳市", AreaCodes: []string{"430500"}},
			{Name: "岳阳市", AreaCodes: []string{"430600"}},
		},
	},
	{
		Name: "广东省", Short: "广东", Code: "44",
		Cities: []City{
			{Name: "广州市", AreaCodes: []string{"440100", "440103", "440104", "440105"}},
			{Name: "深圳市", AreaCodes: []string{"440300", "440303", "440304", "440305"}},
			{Name: "珠海市", AreaCodes: []string{"440400"}},
			{Name: "汕头市", AreaCodes: []string{"440500"}},
			{Name: "佛山市", AreaCodes: []string{"440600"}},
			{Name: "东莞市", AreaCodes: []string{"441900"}},
			{Name: "中山市", AreaCodes: []string{"442000"}},
			{Name: "惠州市", AreaCodes: []string{"441300"}},
		},
	},
	{
		Name: "广西壮族自治区", Short: "广西", Code: "45",
		Cities: []City{
			{Name: "南宁市", AreaCodes: []string{"450100", "450102"}},
			{Name: "柳州市", AreaCodes: []string{"450200"}},
			{Name: "桂林市", AreaCodes: []string{"450300"}},
			{Name: "梧州市", AreaCodes: []string{"450400"}},
			{Name: "北海市", AreaCodes: []string{"450500"}},
		},
	},
	{
		Name: "海南省", Short: "海南", Code: "46",
		Cities: []City{
			{Name: "海口市", AreaCodes: []string{"460100", "460105"}},
			{Name: "三亚市", AreaCodes: []string{"460200"}},
			{Name: "三沙市", AreaCodes: []string{"460300"}},
			{Name: "儋州市", AreaCodes: []string{"460400"}},
		},
	},
	{
		Name: "重庆市", Short: "重庆", Code: "50",
		Cities: []City{
			{Name: "渝中区", AreaCodes: []string{"500101"}},
			{Name: "江北区", AreaCodes: []string{"500105"}},
			{Name: "沙坪坝区", AreaCodes: []string{"500106"}},
			{Name: "九龙坡区", AreaCodes: []string{"500107"}},
			{Name: "南岸区", AreaCodes: []string{"500108"}},
			{Name: "渝北区", AreaCodes: []string{"500112"}},
			{Name: "巴南区", AreaCodes: []string{"500113"}},
		},
	},
	{
		Name: "四川省", Short: "四川", Code: "51",
		Cities: []City{
			{Name: "成都市", AreaCodes: []string{"510100", "510104", "510105"}},
			{Name: "自贡市", AreaCodes: []string{"510300"}},
			{Name: "攀枝花市", AreaCodes: []string{"510400"}},
			{Name: "泸州市", AreaCodes: []string{"510500"}},
			{Name: "德阳市", AreaCodes: []string{"510600"}},
			{Name: "绵阳市", AreaCodes: []string{"510700"}},
		},
	},
	{
		Name: "贵州省", Short: "贵州", Code: "52",
		Cities: []City{
			{Name: "贵阳市", AreaCodes: []string{"520100", "520102"}},
			{Name: "六盘水市", AreaCodes: []string{"520200"}},
			{Name: "遵义市", AreaCodes: []string{"520300"}},
			{Name: "安顺市", AreaCodes: []string{"520400"}},
		},
	},
	{
		Name: "云南省", Short: "云南", Code: "53",
		Cities: []City{
			{Name: "昆明市", AreaCodes: []string{"530100", "530102"}},
			{Name: "曲靖市", AreaCodes: []string{"530300"}},
			{Name: "玉溪市", AreaCodes: []string{"530400"}},
			{Name: "保山市", AreaCodes: []string{"530500"}},
			{Name: "昭通市", AreaCodes: []string{"530600"}},
			{Name: "大理市", AreaCodes: []string{"532901"}},
		},
	},
	{
		Name: "西藏自治区", Short: "西藏", Code: "54",
		Cities: []City{
			{Name: "拉萨市", AreaCodes: []string{"540100", "540102"}},
			{Name: "日喀则市", AreaCodes: []string{"540200"}},
			{Name: "昌都市", AreaCodes: []string{"540300"}},
			{Name: "林芝市", AreaCodes: []string{"540400"}},
		},
	},
	{
		Name: "陕西省", Short: "陕西", Code: "61",
		Cities: []City{
			{Name: "西安市", AreaCodes: []string{"610100", "610102", "610103"}},
			{Name: "铜川市", AreaCodes: []string{"610200"}},
			{Name: "宝鸡市", AreaCodes: []string{"610300"}},
			{Name: "咸阳市", AreaCodes: []string{"610400"}},
			{Name: "渭南市", AreaCodes: []string{"610500"}},
			{Name: "延安市", AreaCodes: []string{"610600"}},
		},
	},
	{
		Name: "甘肃省", Short: "甘肃", Code: "62",
		Cities: []City{
			{Name: "兰州市", AreaCodes: []string{"620100", "620102"}},
			{Name: "嘉峪关市", AreaCodes: []string{"620200"}},
			{Name: "金昌市", AreaCodes: []string{"620300"}},
			{Name: "白银市", AreaCodes: []string{"620400"}},
			{Name: "天水市", AreaCodes: []string{"620500"}},
		},
	},
	{
		Name: "青海省", Short: "青海", Code: "63",
		Cities: []City{
			{Name: "西宁市", AreaCodes: []string{"630100", "630102"}},
			{Name: "海东市", AreaCodes: []string{"630200"}},
			{Name: "海北州", AreaCodes: []string{"632200"}},
			{Name: "黄南州", AreaCodes: []string{"632300"}},
		},
	},
	{
		Name: "宁夏回族自治区", Short: "宁夏", Code: "64",
		Cities: []City{
			{Name: "银川市", AreaCodes: []string{"640100", "640104"}},
			{Name: "石嘴山市", AreaCodes: []string{"640200"}},
			{Name: "吴忠市", AreaCodes: []string{"640300"}},
			{Name: "固原市", AreaCodes: []string{"640400"}},
		},
	},
	{
		Name: "新疆维吾尔自治区", Short: "新疆", Code: "65",
		Cities: []City{
			{Name: "乌鲁木齐市", AreaCodes: []string{"650100", "650102"}},
			{Name: "克拉玛依市", AreaCodes: []string{"650200"}},
			{Name: "吐鲁番市", AreaCodes: []string{"650400"}},
			{Name: "哈密市", AreaCodes: []string{"650500"}},
			{Name: "喀什地区", AreaCodes: []string{"653100"}},
		},
	},
	{
		Name: "台湾省", Short: "台湾", Code: "71",
		Cities: []City{
			{Name: "台北市", AreaCodes: []string{"710100"}},
			{Name: "高雄市", AreaCodes: []string{"710200"}},
			{Name: "台中市", AreaCodes: []string{"710300"}},
			{Name: "台南市", AreaCodes: []string{"710400"}},
		},
	},
	{
		Name: "香港特别行政区", Short: "香港", Code: "81",
		Cities: []City{
			{Name: "香港岛", AreaCodes: []string{"810100"}},
			{Name: "九龙", AreaCodes: []string{"810200"}},
			{Name: "新界", AreaCodes: []string{"810300"}},
		},
	},
	{
		Name: "澳门特别行政区", Short: "澳门", Code: "82",
		Cities: []City{
			{Name: "澳门半岛", AreaCodes: []string{"820100"}},
			{Name: "氹仔岛", AreaCodes: []string{"820200"}},
		},
	},
}

// ProvinceMap 省份名称到 Province 的映射
var ProvinceMap map[string]*Province

// AreaCodeMap 6位地区码到省市信息的映射
var AreaCodeMap map[string]struct {
	Province string
	City     string
}

func init() {
	ProvinceMap = make(map[string]*Province)
	AreaCodeMap = make(map[string]struct {
		Province string
		City     string
	})

	for i := range Provinces {
		p := &Provinces[i]
		ProvinceMap[p.Name] = p
		ProvinceMap[p.Short] = p

		for _, city := range p.Cities {
			for _, code := range city.AreaCodes {
				AreaCodeMap[code] = struct {
					Province string
					City     string
				}{Province: p.Short, City: city.Name}
			}
		}
	}
}
