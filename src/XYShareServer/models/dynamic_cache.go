package models

//动态内容
type DynamicContentCache struct {
	NickName string
	Avatar string
	DynamicContent string
	ViewNUM string
	ForwardNum string 
	CommentNum string
	GoodNUM string
	Images string
	Voices string
	PostTime string
}

//动态评论
type DynamicComentCache struct {
	NickName string
	Avatar string
	Contents string
	PostTime string
}

//动态缓存
type DynamicCache struct {
	Dynamic  map[string]string
	Comment  []map[string]string
}