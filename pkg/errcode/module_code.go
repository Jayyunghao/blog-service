package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")
	ErrorGetTagFail     = NewError(20001006, "获取标签失败")

	ErrorCountArticleFail  = NewError(30010001, "统计文章失败")
	ErrorGetArticleListFail    = NewError(30010002, "获取文章列表失败")
	ErrorCreateArticleFail = NewError(30010003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(30010004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(30010005, "删除文章失败")
	ErrorGetArticleFail     = NewError(30001006, "获取文章失败")

	ErrorUploadFileFail  = NewError(20030001, "上传文件失败")
)
