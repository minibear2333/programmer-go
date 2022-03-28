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

type UpdateInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateInterviewLogic {
	return UpdateInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateInterview  这个接口应该抛弃，不应该全量更新
// TODO 更改为 不能更新created_time、作者的接口，可能需要更改mongo的接口；参数可选传入（需要重新定义req model）
func (l *UpdateInterviewLogic) UpdateInterview(req types.Interview_detail) (resp *types.Interview_detail, err error) {
	if !bson.IsObjectIdHex(req.ID) {
		err = errors.New("面试题识别错误")
		global.LOG.Error("面试题目id识别错误")
		return nil, err
	}
	if !bson.IsObjectIdHex(req.Author.ID) {
		err = errors.New("作者识别错误")
		global.LOG.Error("作者id解析出错")
		return nil, err
	}
	interview := model.Interviews{
		ID: bson.ObjectIdHex(req.ID),
		Author: model.Author{
			ID:   bson.ObjectIdHex(req.Author.ID),
			Name: req.Author.Name,
		},
		Content:     req.Content,
		Bad:         req.Bad,
		ClickNum:    req.ClickNum,
		Comments:    nil,
		Good:        req.Good,
		HardStatus:  req.HardStatus,
		HotNum:      req.HotNum,
		StarNum:     req.StarNum,
		Summary:     req.Summary,
		Tags:        req.Tags,
		Title:       req.Title,
		UpdatedTime: time.Now(),
		CreatedTime: time.Now(),
	}
	err = global.Mongo.InterviewsModel.Update(context.TODO(), &interview)
	if err != nil {
		global.LOG.Error("更新面试题失败:", zap.Error(err))
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
