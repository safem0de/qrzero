package entity

type FileExistRequest struct {
    Path string `json:"path" binding:"required"`
}

type FileExistResponse struct {
    Exists bool `json:"exists"`
}
