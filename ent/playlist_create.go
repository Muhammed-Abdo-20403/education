// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"education/ent/playlist"
	"education/ent/upload"
	"education/ent/user"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaylistCreate is the builder for creating a Playlist entity.
type PlaylistCreate struct {
	config
	mutation *PlaylistMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (pc *PlaylistCreate) SetUserID(i int) *PlaylistCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetTitle sets the "title" field.
func (pc *PlaylistCreate) SetTitle(s string) *PlaylistCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PlaylistCreate) SetUser(u *User) *PlaylistCreate {
	return pc.SetUserID(u.ID)
}

// AddUploadIDs adds the "uploads" edge to the Upload entity by IDs.
func (pc *PlaylistCreate) AddUploadIDs(ids ...int) *PlaylistCreate {
	pc.mutation.AddUploadIDs(ids...)
	return pc
}

// AddUploads adds the "uploads" edges to the Upload entity.
func (pc *PlaylistCreate) AddUploads(u ...*Upload) *PlaylistCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddUploadIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (pc *PlaylistCreate) Mutation() *PlaylistMutation {
	return pc.mutation
}

// Save creates the Playlist in the database.
func (pc *PlaylistCreate) Save(ctx context.Context) (*Playlist, error) {
	var (
		err  error
		node *Playlist
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Playlist)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlaylistMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlaylistCreate) SaveX(ctx context.Context) *Playlist {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlaylistCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlaylistCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlaylistCreate) check() error {
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Playlist.user_id"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Playlist.title"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Playlist.user"`)}
	}
	return nil
}

func (pc *PlaylistCreate) sqlSave(ctx context.Context) (*Playlist, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PlaylistCreate) createSpec() (*Playlist, *sqlgraph.CreateSpec) {
	var (
		_node = &Playlist{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: playlist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlist.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldTitle,
		})
		_node.Title = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist.UserTable,
			Columns: []string{playlist.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.UploadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playlist.UploadsTable,
			Columns: []string{playlist.UploadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: upload.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlaylistCreateBulk is the builder for creating many Playlist entities in bulk.
type PlaylistCreateBulk struct {
	config
	builders []*PlaylistCreate
}

// Save creates the Playlist entities in the database.
func (pcb *PlaylistCreateBulk) Save(ctx context.Context) ([]*Playlist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Playlist, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaylistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlaylistCreateBulk) SaveX(ctx context.Context) []*Playlist {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlaylistCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlaylistCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}