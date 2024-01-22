package server

import (
	"fmt"
	"github.com/disism/karma/ent"
	"github.com/disism/karma/ent/dirs"
	"github.com/disism/karma/ent/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultRootDirName     = "root"
	InvalidDirName         = "Invalid dir name"
	DirExists              = "Dir is exists"
	DisableRootDirCreation = "Root dir creation is not allowed"
	DisableRootDirRenaming = "Root dir renaming is not allowed"
	DisableRootDirDeletion = "Disable root dir deletion"
)

type Dir struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
	Subdirs    []*Subdir `json:"subdirs"`
	Saves      []*Saved  `json:"saves"`
}

type Subdir struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
}

type MKDirForm struct {
	Name string `form:"name" json:"name" binding:"required,min=1,max=225"`
}

type MKDirQuery struct {
	DirID uint64 `form:"dir_id" binding:"numeric"`
}

func CreateRootDirIfNot(ctx *gin.Context, client *ent.Client) (*ent.Dirs, error) {
	query, err := client.Dirs.Query().
		Where(
			dirs.NameEQ(DefaultRootDirName),
			dirs.HasOwnerWith(
				users.IDEQ(
					GetUserID(ctx),
				),
			),
		).
		WithSubdir().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			create, err := client.Dirs.Create().
				SetOwnerID(
					GetUserID(ctx),
				).
				SetName(DefaultRootDirName).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("mk dir create root dir error: %w", err)
			}
			return create, nil
		}
		return nil, fmt.Errorf("mk dir query root dir error: %w", err)
	}
	return query, nil
}

func (s *Server) MKDir() error {
	defer s.client.Close()

	var q MKDirQuery
	if err := s.ctx.ShouldBindQuery(&q); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	var f MKDirForm
	if err := s.ctx.ShouldBind(&f); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("mk dir tx error: %w", err)
	}
	defer tx.Rollback()

	if strings.TrimSpace(f.Name) == "" {
		ErrorBadRequest(s.ctx, InvalidDirName)
		return nil
	}
	query := tx.Dirs.
		Query().
		Where(
			dirs.HasOwnerWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
		)

	r := &ent.Dirs{}
	// If the ID of the query directory is specified.
	if q.DirID != 0 {
		r, err = query.
			Where(
				dirs.IDEQ(q.DirID),
			).
			WithSubdir().
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				ErrorNoPermission(s.ctx, NoPermission)
				return nil
			}
			return fmt.Errorf("mk dir query dir error: %w", err)
		}
	} else {
		// If the ID of the query directory is not specified, the root directory named DefaultRootDirName is queried,
		// and if it is not queried, it is created.
		r, err = query.
			Where(
				dirs.NameEQ(DefaultRootDirName),
			).
			WithSubdir().
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				r, err = tx.Dirs.Create().
					SetOwnerID(
						GetUserID(s.ctx),
					).
					SetName(DefaultRootDirName).
					Save(s.ctx)
				if err != nil {
					return fmt.Errorf("mk dir create root dir error: %w", err)
				}
			}
			return fmt.Errorf("mk dir query root dir error: %w", err)
		}
	}
	if r.Edges.Subdir != nil {
		for _, sub := range r.Edges.Subdir {
			if sub.Name == f.Name {
				ErrorConflict(s.ctx, DirExists)
				return nil
			}
		}
	}

	if strings.EqualFold(DefaultRootDirName, f.Name) {
		ErrorConflict(s.ctx, DisableRootDirCreation)
		return nil
	}

	create, err := tx.Dirs.
		Create().
		SetOwnerID(
			GetUserID(s.ctx),
		).
		SetName(f.Name).
		AddPdir(r).
		Save(s.ctx)
	if err != nil {
		return fmt.Errorf("mk dir create dir error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("mk dir commit error: %w", err)
	}

	Success(s.ctx, &Dir{
		ID:         strconv.FormatUint(create.ID, 10),
		CreateTime: create.CreateTime,
		UpdateTime: create.UpdateTime,
		Name:       create.Name,
		Subdirs:    nil,
	})
	return nil
}

func QueryUserDirByID(ctx *gin.Context, client *ent.Client, id uint64) (*ent.Dirs, error) {
	return client.Dirs.Query().
		Where(
			dirs.IDEQ(id),
			dirs.HasOwnerWith(
				users.IDEQ(
					GetUserID(ctx),
				),
			),
		).
		WithPdir().
		WithSubdir().
		WithOwner().
		WithSaves(
			func(q *ent.SavesQuery) {
				q.WithFile()
			},
		).
		Only(ctx)
}

type ListDirQuery struct {
	DirID uint64 `form:"dir_id" binding:"numeric"`
}

func (s *Server) ListDir() error {
	defer s.client.Close()

	var q ListDirQuery
	if err := s.ctx.ShouldBindQuery(&q); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	options := s.client.Dirs.
		Query().
		Where(
			dirs.HasOwnerWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
		)

	query := &ent.Dirs{}
	var err error
	if q.DirID != 0 {
		query, err = QueryUserDirByID(s.ctx, s.client, q.DirID)
		if err != nil {
			if ent.IsNotFound(err) {
				ErrorNoPermission(s.ctx, NoPermission)
				return nil
			}
			return fmt.Errorf("mk dir query dir error: %w", err)
		}
	} else {
		query, err = options.
			Where(dirs.NameEQ(DefaultRootDirName)).
			WithSaves(
				func(q *ent.SavesQuery) {
					q.WithFile()
				},
			).
			WithSubdir().
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				query, err = s.client.Dirs.Create().SetOwnerID(GetUserID(s.ctx)).SetName(DefaultRootDirName).Save(s.ctx)
				if err != nil {
					return fmt.Errorf("mk dir create root dir error: %w", err)
				}
			}
			return fmt.Errorf("mk dir query root dir error: %w", err)
		}
	}
	r := &Dir{
		ID:         strconv.FormatUint(query.ID, 10),
		CreateTime: query.CreateTime,
		UpdateTime: query.UpdateTime,
		Name:       query.Name,
	}
	if query.Edges.Subdir != nil {
		subdirs := make([]*Subdir, len(query.Edges.Subdir))
		for i, sub := range query.Edges.Subdir {
			subdirs[i] = &Subdir{
				ID:         strconv.FormatUint(sub.ID, 10),
				CreateTime: sub.CreateTime,
				UpdateTime: sub.UpdateTime,
				Name:       sub.Name,
			}
		}
		r.Subdirs = subdirs
	}
	if query.Edges.Saves != nil {
		saves := make([]*Saved, len(query.Edges.Saves))
		for i, sav := range query.Edges.Saves {
			saves[i] = &Saved{
				ID:         strconv.FormatUint(sav.ID, 10),
				CreateTime: sav.CreateTime,
				UpdateTime: sav.UpdateTime,
				Name:       sav.Name,
				Caption:    sav.Caption,
				File: &File{
					ID:   strconv.FormatUint(sav.Edges.File.ID, 10),
					Hash: sav.Edges.File.Hash,
					Name: sav.Edges.File.Name,
					Size: strconv.FormatUint(sav.Edges.File.Size, 10),
				},
			}
		}
		r.Saves = saves
	}
	s.ctx.JSON(http.StatusOK, r)
	return nil
}

func (s *Server) ListDirs() error {
	defer s.client.Close()

	query, err := s.client.Dirs.
		Query().
		Where(
			dirs.HasOwnerWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
		).All(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("list dirs query error: %w", err)
	}

	r := make([]*Dir, len(query))
	for i, d := range query {
		r[i] = &Dir{
			ID:         strconv.FormatUint(d.ID, 10),
			CreateTime: d.CreateTime,
			UpdateTime: d.UpdateTime,
			Name:       d.Name,
		}
	}
	Success(s.ctx, r)
	return nil
}

type RenameDirParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) RenameDir() error {
	defer s.client.Close()

	var p RenameDirParams
	if err := s.ctx.ShouldBindUri(&p); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	name := s.ctx.PostForm("name")

	if strings.EqualFold(name, DefaultRootDirName) {
		ErrorBadRequest(s.ctx, DisableRootDirRenaming)
		return nil
	}
	query, err := QueryUserDirByID(s.ctx, s.client, p.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("rename dir query dir error: %w", err)
	}

	// If the directory name already exists at the same level, it is not allowed to be modified.
	subs, err := query.QueryPdir().WithSubdir().Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("rename dir query dir error: %w", err)
	}
	if subs != nil {
		for _, d := range subs.Edges.Subdir {
			if d.Name == name {
				ErrorBadRequest(s.ctx, DirExists)
				return nil
			}
		}
	}

	if err := s.client.Dirs.Update().Where(dirs.IDEQ(query.ID)).SetName(name).Exec(s.ctx); err != nil {
		return fmt.Errorf("rename dir update dir error: %w", err)
	}
	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type MVDirParams struct {
	ID    uint64 `uri:"id" binding:"required,numeric"`
	NewID uint64 `uri:"new_id" binding:"required,numeric"`
}

func (s *Server) MVDir() error {
	defer s.client.Close()

	var p MVDirParams
	if err := s.ctx.ShouldBindUri(&p); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	if p.ID == p.NewID {
		ErrorBadRequest(s.ctx, "its own subdirectory")
		return nil
	}

	o, err := QueryUserDirByID(s.ctx, s.client, p.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("mv dir query old dir error: %w", err)
	}
	if o.Edges.Pdir != nil {
		for _, par := range o.Edges.Pdir {
			if par.ID == p.NewID {
				ErrorBadRequest(s.ctx, DirExists)
				return nil
			}
		}
	}
	n, err := QueryUserDirByID(s.ctx, s.client, p.NewID)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("mv dir query new dir error: %w", err)
	}
	if n.Edges.Subdir != nil {
		for _, sub := range n.Edges.Subdir {
			if sub.ID == p.ID {
				ErrorBadRequest(s.ctx, DirExists)
				return nil
			}
		}
	}

	if err := s.client.Dirs.Update().Where(dirs.IDEQ(p.ID)).RemovePdir(o.Edges.Pdir...).AddPdir(n).Exec(s.ctx); err != nil {
		return fmt.Errorf("mv dir update dir error: %w", err)
	}
	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}

type RMDirParams struct {
	ID uint64 `uri:"id" binding:"required,numeric"`
}

func (s *Server) RMDir() error {
	defer s.client.Close()

	var p RMDirParams
	if err := s.ctx.ShouldBindUri(&p); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("rm dir tx error: %w", err)
	}
	defer tx.Rollback()

	query, err := QueryUserDirByID(s.ctx, s.client, p.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("rm dir query dir error: %w", err)
	}

	if query.Name == DefaultRootDirName {
		ErrorBadRequest(s.ctx, DisableRootDirDeletion)
		return nil
	}
	// TODO REMOVE ALL LINK FILES:Recursively remove all subdirectories with SAVED links

	// Recursively remove all subdirectories

	remove := make([]*ent.Dirs, 0)
	var re func(query *ent.Dirs) error
	re = func(query *ent.Dirs) error {
		for _, sub := range query.Edges.Subdir {
			remove = append(remove, sub)
			f, err := tx.Dirs.Query().
				Where(dirs.IDEQ(sub.ID)).
				WithSubdir().
				Only(s.ctx)
			if err != nil {
				return fmt.Errorf("rm dir query dir error: %w", err)
			}
			if err := re(f); err != nil {
				return fmt.Errorf("rm dir delete dir error: %w", err)
			}
		}
		return nil
	}
	if err := re(query); err != nil {
		return fmt.Errorf("rm dir delete dir error: %w", err)
	}

	remove = append(remove, query)
	for _, sub := range remove {
		if err := tx.Dirs.DeleteOne(sub).Exec(s.ctx); err != nil {
			return fmt.Errorf("rm dir delete dir error: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("rm dir commit error: %w", err)
	}
	Success(s.ctx, gin.H{"code": http.StatusOK})
	return nil
}
