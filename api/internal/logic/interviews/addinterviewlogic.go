package interviews

import (
	"context"
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/model"
	"go.uber.org/zap"
	"time"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddInterviewLogic {
	return AddInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddInterviewLogic) AddInterview(req types.ReqInterviewAdd) (resp *types.Interview_detail, err error) {
	if !bson.IsObjectIdHex(req.Author.ID) {
		err = errors.New("作者识别错误")
		global.LOG.Error("作者id解析出错")
		return nil, err
	}
	interview := model.Interviews{
		Author: model.Author{
			ID:   bson.ObjectIdHex(req.Author.ID),
			Name: req.Author.Name,
		},
		Content:     req.Content,
		Bad:         0,
		ClickNum:    0,
		Comments:    nil,
		CreatedTime: time.Now(),
		Good:        0,
		HardStatus:  req.HardStatus,
		HotNum:      0,
		StarNum:     0,
		Summary:     req.Summary,
		Tags:        req.Tags,
		Title:       req.Title,
		UpdatedTime: time.Now(),
	}
	err = global.Mongo.InterviewsModel.Insert(context.TODO(), &interview)
	if err != nil {
		global.LOG.Error("创建面试题失败", zap.Error(err))
		return nil, err
	}
	return &types.Interview_detail{
		Interview: types.Interview{
			ID: interview.ID.Hex(),
			Author: types.Author{
				ID:   interview.Author.ID.Hex(),
				Name: interview.Author.Name,
			},
			ClickNum:    interview.ClickNum,
			Good:        interview.Good,
			HardStatus:  interview.HardStatus,
			HotNum:      interview.HotNum,
			Summary:     interview.Summary,
			Tags:        interview.Tags,
			Title:       interview.Title,
			UpdatedTime: interview.UpdatedTime.Unix(),
			Status:      false,
		},
		StarNum:     interview.StarNum,
		Bad:         interview.Bad,
		Content:     interview.Content,
		CreatedTime: interview.CreatedTime.Unix(),
	}, nil
}
