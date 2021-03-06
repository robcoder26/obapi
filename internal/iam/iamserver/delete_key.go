package iamserver

import (
	"context"
	"oceanbolt.com/obapi/internal/iam/dao"
	"oceanbolt.com/obapi/rpc/iam"
)

func (s *Server) DeleteKey(ctx context.Context, req *iam.DeleteKeyRequest) (resp *iam.KeyDeletedResponse, err error) {
	db := dao.IamDAO{Ctx: ctx, Ds: s.Ds}

	err = db.DeleteKeyDS(req)
	if err != nil {
		return resp, err
	}

	resp = &iam.KeyDeletedResponse{
		Message: "Key '" + req.ApikeyId + "' successfully deleted",
	}
	return resp, nil
}
