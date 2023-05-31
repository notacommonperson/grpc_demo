package rpc

import (
	"context"
	"errors"
	"grpc_demo/idl/pb/task"
)

func TaskCreate(ctx context.Context, req *task.CreateReq) error {
	_, err := taskClient.CreateTask(ctx, req)
	if err != nil {
		return errors.New("task rpc call failed")
	}
	return nil
}

func TaskListGet(ctx context.Context, req *task.GetListReq) ([]*task.Task, error) {
	res, err := taskClient.GetListTask(ctx, req)
	if err != nil {
		return nil, errors.New("task rpc call faild")
	}

	return res.Task, nil
}

func TaskSearch(ctx context.Context, req *task.SearchReq) ([]*task.Task, error) {
	res, err := taskClient.SearchTask(ctx, req)
	if err != nil {
		return nil, errors.New("task rpc call faild")
	}

	return res.Task, nil
}

func TaskUpdate(ctx context.Context, req *task.UpdateReq) (*task.Task, error) {
	res, err := taskClient.UpdataTask(ctx, req)
	if err != nil {
		return nil, errors.New("task rpc call failed")
	}

	return res.Task, nil
}

func TaskDelete(ctx context.Context, req *task.DeleteReq) error {
	_, err := taskClient.DeleteTask(ctx, req)
	if err != nil {
		return errors.New("task rpc call failed")
	}

	return nil
}
