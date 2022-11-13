package service

import "github.com/three-little-dragons/my-whiteboard-server/internal/app/proto"

func NewPaint(userId int64, nickname string) proto.Paint {
	tab := proto.DrawTab{
		ID: GenerateId(),
		User: proto.User{
			ID:       userId,
			Nickname: nickname,
		},
	}
	paint := proto.Paint{
		Proto:    proto.Proto{Type: proto.TypePaint},
		Version:  "1",
		DrawTab:  tab,
		Elements: map[int64][]proto.Graph{},
	}
	return paint
}
