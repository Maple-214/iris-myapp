package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	booksAPI := app.Party("/v3/api")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/homePageContent", commom)
		// POST: http://localhost:8080/books
		// booksAPI.Post("/", create)
	}

	app.Listen(":1111")
}

// Book example.
type SlidesProps struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
}

type CategoryProps struct {
	Id    int    `json:"id"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
}

type RecommendProps struct {
	Image       string `json:"image"`
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
	MallPrice   int    `json:"mallPrice"`
}

type homePageContent struct {
	Slides          []SlidesProps    `json:"slides"`
	Category        []CategoryProps  `json:"category"`
	ShopInfo        string           `json:"shopInfo"`
	Recommend       []RecommendProps `json:"recommend"`
	AdvertesPicture string           `json:"advertesPicture"`
}

func commom(ctx iris.Context) {
	slides := []SlidesProps{
		{Id: 0, Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTBOBdFe_MJJj7v0fJr5tBiGOS64FfmrMkyEQ&usqp=CAU"},
		{Id: 1, Image: "https://image.maigoo.com/upload/images/20221101/1404314044_750x498.jpg"},
		{Id: 2, Image: "https://storage.googleapis.com/pos.shopline.hk/2021/11/onlineshop-vs-marketplace-hero-image.jpg"},
	}

	category := []CategoryProps{
		{Id: 0, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/187640/12/30456/5256/639c2315Ebc95c142/350a8f0c766f5460.png", Title: "京东超市"},
		{Id: 1, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/178015/31/13828/6862/60ec0c04Ee2fd63ac/ccf74d805a059a44.png", Title: "数码电器"},
		{Id: 2, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/54043/33/19389/4660/62b049dbE3b9aef75/2fcd31afd5d702e4.png", Title: "京东新百货"},
		{Id: 3, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/177902/16/13776/5658/60ec0e71E801087f2/a0d5a68bf1461e6d.png", Title: "京东生鲜"},
		{Id: 4, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/196472/7/12807/7127/60ec0ea3Efe11835b/37c65625d94cae75.png", Title: "京东到家"},
		{Id: 5, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/185733/21/13527/6648/60ec0f31E0fea3e0a/d86d463521140bb6.png", Title: "充值缴费"},
		{Id: 6, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/34248/39/16616/4689/62bbbdccE9f11132e/d51caf15f2f412b2.png", Title: "附近好店"},
		{Id: 7, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/37709/6/15279/6118/60ec1046E4b5592c6/a7d6b66354efb141.png", Title: "PLUS会员"},
		{Id: 8, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/191060/24/12861/6818/60ec11f2E67cf5ee6/c264378678b3cd96.png", Title: "京东国际"},
		{Id: 9, Icon: "https://m15.360buyimg.com/mobilecms/jfs/t1/186882/8/12149/5894/60ec1250E9335241a/b22054613aa8ae75.png", Title: "京东拍卖"},
	}

	recommend := []RecommendProps{
		{
			Image:       "https://www.imuraya.co.jp/media_images/2019/03/box-azukiba-maccha.jpg",
			ProductName: "商品名1",
			Price:       20.0,
			MallPrice:   30.0,
		},
		{
			Image:       "https://www.calbee.co.jp/common/utility/binout.php?db=products&f=3152",
			ProductName: "商品名2",
			Price:       30.0,
			MallPrice:   35.0,
		},
		{
			Image:       "https://www.asahi-gf.co.jp/products/images/baby/milk/wakodo/gungun/40195890.jpg",
			ProductName: "商品名3",
			Price:       40.0,
			MallPrice:   48.0,
		},
		{
			Image:       "https://item-shopping.c.yimg.jp/i/l/pc-parts_tkm004",
			ProductName: "商品名4",
			Price:       60.0,
			MallPrice:   78.0,
		},
	}
	home := homePageContent{
		Slides:          slides,
		Category:        category,
		ShopInfo:        "https://www.chuangkit.com/shejixuexi/images/banner2.jpg",
		Recommend:       recommend,
		AdvertesPicture: "https://pic40.photophoto.cn/20160910/0017029583521992_b.jpg",
	}

	ctx.JSON(home)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

// func create(ctx iris.Context) {
// 	var b homePageContent
// 	err := ctx.ReadJSON(&b)
// 	// TIP: use ctx.ReadBody(&b) to bind
// 	// any type of incoming data instead.
// 	if err != nil {
// 		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
// 			Title("Book creation failure").DetailErr(err))
// 		// TIP: use ctx.StopWithError(code, err) when only
// 		// plain text responses are expected on errors.
// 		return
// 	}
// 	ctx.StatusCode(iris.StatusCreated)
// }
