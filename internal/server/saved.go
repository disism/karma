package server

import (
	fmt "fmt"
	"github.com/disism/karma/ent"
	"github.com/disism/karma/ent/dirs"
	"github.com/disism/karma/ent/files"
	"github.com/disism/karma/ent/saves"
	"github.com/disism/karma/ent/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	DirAlreadyLinked = "dir already linked"
	DirNotLinked     = "dir not linked"
)

type Saved struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
	Caption    string    `json:"caption"`
	File       *File     `json:"file"`
}

type File struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
	Name string `json:"name"`
	Size string `json:"size"`
}

type AddSavedForm struct {
	Hash    string `form:"hash" json:"hash" binding:"required"`
	Name    string `form:"name" json:"name" binding:"required"`
	Size    string `form:"size" json:"size" binding:"required"`
	Caption string `form:"caption" json:"caption" binding:"max=999"`
}

type CreateSavesQuery struct {
	DirID uint64 `form:"dir_id" json:"dir_id" binding:"numeric"`
}

func FMTSaved(saved *ent.Saves) *Saved {
	r := &Saved{
		ID:         strconv.FormatUint(saved.ID, 10),
		CreateTime: saved.CreateTime,
		UpdateTime: saved.UpdateTime,
		Name:       saved.Name,
		Caption:    saved.Caption,
	}
	if saved.Edges.File != nil {
		r.File = &File{
			ID:   strconv.FormatUint(saved.Edges.File.ID, 10),
			Hash: saved.Edges.File.Hash,
			Name: saved.Edges.File.Name,
			Size: strconv.FormatUint(saved.Edges.File.Size, 10),
		}
	}
	return r
}

func QueryUserSavedByFileID(ctx *gin.Context, client *ent.Client, id uint64) (*ent.Saves, error) {
	query, err := client.Saves.
		Query().
		Where(
			saves.HasFileWith(
				files.IDEQ(id),
			),
			saves.HasOwnerWith(
				users.IDEQ(
					GetUserID(ctx),
				),
			),
		).
		WithDir().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func QueryFileByHash(ctx *gin.Context, client *ent.Client, hash string) (*ent.Files, error) {
	query, err := client.Files.
		Query().
		Where(
			files.HashEQ(hash),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (s *Server) CreateSaves() error {
	defer s.client.Close()

	var q CreateSavesQuery
	if err := s.ctx.ShouldBindQuery(&q); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	var forms []*AddSavedForm
	if err := s.ctx.ShouldBind(&forms); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("create saved tx error: %w", err)
	}
	defer tx.Rollback()

	var exists []*Saved
	var creates []*Saved

	for _, form := range forms {
		f, err := func() (*ent.Files, error) {
			query, err := QueryFileByHash(s.ctx, tx.Client(), form.Hash)
			if err != nil {
				if ent.IsNotFound(err) {
					size, err := strconv.ParseUint(form.Size, 10, 64)
					if err != nil {
						return nil, fmt.Errorf("parse size error: %w", err)
					}
					create, err := tx.Files.Create().
						SetHash(form.Hash).
						SetSize(size).
						SetName(form.Name).
						Save(s.ctx)
					if err != nil {
						return nil, fmt.Errorf("create file error: %w", err)
					}
					return create, nil

				}
				return nil, fmt.Errorf("query file error: %w", err)
			}
			return query, nil
		}()
		if err != nil {
			return err
		}

		saved, err := QueryUserSavedByFileID(s.ctx, tx.Client(), f.ID)
		if err != nil {
			if ent.IsNotFound(err) {
				create, err := tx.Saves.Create().
					SetFile(f).
					SetName(f.Name[:strings.LastIndex(f.Name, ".")]).
					SetCaption(form.Caption).
					SetOwnerID(
						GetUserID(s.ctx),
					).
					Save(s.ctx)
				if err != nil {
					return fmt.Errorf("create saved error: %w", err)
				}
				creates = append(creates, FMTSaved(create))
			} else {
				return fmt.Errorf("query saved error: %w", err)
			}
		} else {
			exists = append(exists, FMTSaved(saved))
		}

		var dir *ent.Dirs
		if q.DirID == 0 {
			dir, err = CreateRootDirIfNot(s.ctx, tx.Client())
			if err != nil {
				return fmt.Errorf("create root dir error: %w", err)
			}
		} else {
			dir, err = tx.Dirs.Query().Where(
				dirs.HasOwnerWith(
					users.IDEQ(
						GetUserID(s.ctx),
					),
				),
				dirs.IDEQ(q.DirID),
			).Only(s.ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					ErrorNoPermission(s.ctx, NoPermission)
				}
				return fmt.Errorf("query dir error: %w", err)
			}
		}

		if err := tx.Saves.Update().
			Where(
				saves.HasOwnerWith(
					users.IDEQ(
						GetUserID(s.ctx),
					),
				),
			).
			AddDir(dir).
			Exec(s.ctx); err != nil {
			return fmt.Errorf("add saved error: %w", err)
		}

	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("add saved commit error: %w", err)
	}

	Success(s.ctx, gin.H{
		"code":    http.StatusCreated,
		"exists":  exists,
		"creates": creates,
	})
	return nil
}

func (s *Server) GetSaves() error {
	defer s.client.Close()

	query, err := s.client.Saves.
		Query().
		Where(
			saves.HasOwnerWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
		).
		WithFile().
		All(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("query saves error: %w", err)
	}

	r := make([]Saved, len(query))
	for i, q := range query {
		r[i] = *FMTSaved(q)
	}

	Success(s.ctx, r)
	return nil
}

type GetSavedParams struct {
	ID uint64 `form:"id" json:"id" binding:"numeric"`
}

func (s *Server) GetSaved() error {
	defer s.client.Close()

	var p GetSavedParams
	if err := s.ctx.ShouldBindUri(&p); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	query, err := s.client.Saves.
		Query().
		Where(
			saves.HasOwnerWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
			saves.IDEQ(p.ID),
		).
		WithFile().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("get saved error: %w", err)
	}

	Success(s.ctx, *FMTSaved(query))
	return nil
}

type EditSavedParams struct {
	ID uint64 `form:"id" json:"id" binding:"numeric"`
}

func (s *Server) EditSaved() error {
	defer s.client.Close()

	var p EditSavedParams
	if err := s.ctx.ShouldBindUri(&p); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	query, err := s.client.Saves.
		Query().
		Where(
			saves.HasOwnerWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
			saves.IDEQ(p.ID),
		).
		WithFile().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("edit saved error: %w", err)
	}
	edit := s.client.Saves.Update().Where(saves.IDEQ(query.ID))
	caption := s.ctx.PostForm("caption")
	switch {
	case strings.TrimSpace(caption) != "":
		edit.SetCaption(s.ctx.PostForm("caption"))
		query.Caption = caption
	}
	if err := edit.Exec(s.ctx); err != nil {
		return fmt.Errorf("edit saved error: %w", err)
	}
	Success(s.ctx, gin.H{
		"code": http.StatusOK,
	})
	return nil
}

type DelSavedParams struct {
	ID uint64 `uri:"id" binding:"numeric,required"`
}

func (s *Server) DelSaved() error {
	defer s.client.Close()

	var p DelSavedParams
	if err := s.ctx.ShouldBindUri(&p); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	query, err := s.client.Saves.
		Query().
		Where(
			saves.HasOwnerWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
			saves.IDEQ(p.ID),
		).
		WithFile().
		WithDir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("del saved error: %w", err)
	}

	// TODO - remove all link dirs.

	if err := s.client.Saves.
		Update().
		Where(
			saves.IDEQ(query.ID),
		).
		RemoveDir(query.Edges.Dir...).
		Exec(s.ctx); err != nil {
		return fmt.Errorf("delete saved error: %w", err)
	}
	if err := s.client.Saves.
		DeleteOne(query).
		Exec(s.ctx); err != nil {
		return fmt.Errorf("delete saved error: %w", err)
	}

	Success(s.ctx, gin.H{"code": http.StatusOK})

	return nil
}

type LinkDirParams struct {
	ID uint64 `form:"id" json:"id" binding:"numeric"`
}

type LinkDirForm struct {
	DirID uint64 `form:"dir_id" json:"dir_id" binding:"numeric"`
}

func (s *Server) LinkDir() error {
	defer s.client.Close()

	var p LinkDirParams
	if err := s.ctx.ShouldBindUri(&p); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}
	var f LinkDirForm
	if err := s.ctx.ShouldBind(&f); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("link dir tx error: %w", err)
	}
	defer tx.Rollback()

	query, err := tx.Saves.Query().Where(
		saves.IDEQ(p.ID),
		saves.HasOwnerWith(
			users.IDEQ(
				GetUserID(s.ctx),
			),
		),
	).
		WithDir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("link dir query saved error: %w", err)
	}

	dirs, err := tx.Dirs.Query().Where(
		dirs.HasOwnerWith(
			users.IDEQ(
				GetUserID(s.ctx),
			),
		),
		dirs.IDEQ(f.DirID),
	).Only(s.ctx)
	if err != nil {
		return fmt.Errorf("link dir query dir error: %w", err)
	}

	if query.Edges.Dir != nil {
		for _, r := range query.Edges.Dir {
			if r.ID == dirs.ID {
				ErrorConflict(s.ctx, DirAlreadyLinked)
				return nil
			}
		}
	}

	if err := tx.Saves.Update().Where(saves.IDEQ(query.ID)).AddDir(dirs).Exec(s.ctx); err != nil {
		return fmt.Errorf("link dir update saved error: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("link dir commit error: %w", err)
	}

	Success(s.ctx, gin.H{
		"code": http.StatusOK,
	})
	return nil
}

type UnlinkDirParams struct {
	ID uint64 `uri:"id" binding:"numeric,required"`
}

func (s *Server) UnlinkDir() error {
	defer s.client.Close()

	var p UnlinkDirParams
	if err := s.ctx.ShouldBindUri(&p); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("unlink dir tx error: %w", err)
	}

	defer tx.Rollback()

	query, err := tx.Saves.
		Query().
		Where(
			saves.IDEQ(p.ID),
			saves.HasOwnerWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
		).
		WithDir().
		Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("unlink dir query saved error: %w", err)
	}

	options := tx.Dirs.Query().Where(dirs.HasOwnerWith(
		users.IDEQ(
			GetUserID(s.ctx),
		),
	))

	dirID, exists := s.ctx.GetPostForm("dir_id")
	if exists {
		parse, err := strconv.ParseUint(dirID, 10, 64)
		if err != nil {
			ErrorBadRequest(s.ctx, err.Error())
			return nil
		}
		options.Where(dirs.IDEQ(parse))
	} else {
		options.Where(dirs.NameEQ(DefaultRootDirName))
	}

	dirs, err := options.Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNoPermission(s.ctx, NoPermission)
			return nil
		}
		return fmt.Errorf("unlink dir query dir error: %w", err)
	}

	if query.Edges.Dir != nil {
		found := func() bool {
			for _, r := range query.Edges.Dir {
				if r.ID == dirs.ID {
					return true
				}
			}
			return false
		}()

		if !found {
			ErrorConflict(s.ctx, DirNotLinked)
			return nil
		}
	}

	if err := tx.Saves.
		Update().
		Where(
			saves.IDEQ(p.ID),
		).
		RemoveDir(dirs).
		Exec(s.ctx); err != nil {
		return fmt.Errorf("unlink dir update saved error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("unlink dir commit error: %w", err)
	}

	Success(s.ctx, gin.H{
		"code": http.StatusOK,
	})

	return nil
}
