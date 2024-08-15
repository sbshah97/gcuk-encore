package monitor

import (
	"context"

	"encore.app/api"
	"encore.dev/pubsub"
)

var _ = pubsub.NewSubscription(
	api.Signups, "send-welcome-email",
	pubsub.SubscriptionConfig[*api.Slideshow]{
		Handler: SetSlideshowDetails,
	},
)

func SetSlideshowDetails(ctx context.Context, api *api.Slideshow) error {
	insert(ctx, api.Slideshow.Author)

	return nil
}

// Get retrieves the original URL for the id.
//
//encore:api public method=GET path=/url/:id
func GetAuthor(ctx context.Context, id string) (*api.Slideshow, error) {
	var author_name string
	err := db.QueryRow(ctx, `
        SELECT author_name FROM events
        WHERE id = $1
    `, id).Scan(&author_name)
	return &api.Slideshow{
		Slideshow: api.SlideshowContent{
			Author: author_name,
		},
	}, err
}
