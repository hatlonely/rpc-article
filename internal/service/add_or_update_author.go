package service

//func (s *Service) AddOrUpdateAuthor(ctx context.Context, req *api.Author) (*api.AuthorID, error) {
//	author, err := s.storage.GetAuthorByKey(ctx, req.Key)
//	if err != nil {
//		return nil, err
//	}
//	if author == nil {
//		id, err := s.storage.PutAuthor(ctx, &storage.Author{
//			Key:    req.Key,
//			Name:   req.Name,
//			Avatar: req.Avatar,
//		})
//		if err != nil {
//			return nil, err
//		}
//		return &api.AuthorID{
//			Id: id,
//		}, nil
//	}
//
//	if err := s.storage.UpdateAuthor(ctx, &storage.Author{
//		ID:     author.ID,
//		Key:    req.Key,
//		Name:   req.Name,
//		Avatar: req.Avatar,
//	}); err != nil {
//		return nil, err
//	}
//
//	return &api.AuthorID{
//		Id: author.ID,
//	}, nil
//}
