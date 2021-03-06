// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package declarative

import (
	"time"
)

import (
	"github.com/lxn/walk"
)

type DateEdit struct {
	// Window

	Background       Brush
	ContextMenuItems []MenuItem
	Enabled          Property
	Font             Font
	MaxSize          Size
	MinSize          Size
	Name             string
	OnKeyDown        walk.KeyEventHandler
	OnKeyPress       walk.KeyEventHandler
	OnKeyUp          walk.KeyEventHandler
	OnMouseDown      walk.MouseEventHandler
	OnMouseMove      walk.MouseEventHandler
	OnMouseUp        walk.MouseEventHandler
	OnSizeChanged    walk.EventHandler
	Persistent       bool
	ToolTipText      Property
	Visible          Property

	// Widget

	AlwaysConsumeSpace bool
	Column             int
	ColumnSpan         int
	Row                int
	RowSpan            int
	StretchFactor      int

	// DateEdit

	AssignTo      **walk.DateEdit
	Date          Property
	Format        string
	MaxDate       time.Time
	MinDate       time.Time
	NoneOption    bool
	OnDateChanged walk.EventHandler
}

func (de DateEdit) Create(builder *Builder) error {
	var w *walk.DateEdit
	var err error

	if de.NoneOption {
		w, err = walk.NewDateEditWithNoneOption(builder.Parent())
	} else {
		w, err = walk.NewDateEdit(builder.Parent())
	}
	if err != nil {
		return err
	}

	return builder.InitWidget(de, w, func() error {
		if err := w.SetFormat(de.Format); err != nil {
			return err
		}

		if err := w.SetRange(de.MinDate, de.MaxDate); err != nil {
			return err
		}

		if de.OnDateChanged != nil {
			w.DateChanged().Attach(de.OnDateChanged)
		}

		if de.AssignTo != nil {
			*de.AssignTo = w
		}

		return nil
	})
}
