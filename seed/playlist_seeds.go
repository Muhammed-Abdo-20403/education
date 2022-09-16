package seed

import (
	"context"
	"education/config"
	"fmt"
)

func DO(ctx context.Context) error {
	_, err := config.Client.Playlist.
		Create().
		SetUserID(1).
		SetTitle("videos palylist").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User %w", err)
	}

	_, err = config.Client.Playlist.
		Create().
		SetUserID(2).
		SetTitle("videos palylist").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User %w", err)
	}

	_, err = config.Client.Playlist.
		Create().
		SetUserID(3).
		SetTitle("videos palylist").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User %w", err)
	}

	_, err = config.Client.Playlist.
		Create().
		SetUserID(4).
		SetTitle("videos palylist").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User %w", err)
	}

	_, err = config.Client.Playlist.
		Create().
		SetUserID(5).
		SetTitle("videos palylist").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User %w", err)
	}

	return nil
}
