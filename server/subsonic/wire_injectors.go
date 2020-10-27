//+build wireinject

package subsonic

import (
	"github.com/deluan/navidrome/core"
	"github.com/google/wire"
)

var allProviders = wire.NewSet(
	NewSystemController,
	NewBrowsingController,
	NewAlbumListController,
	NewMediaAnnotationController,
	NewPlaylistsController,
	NewSearchingController,
	NewUsersController,
	NewMediaRetrievalController,
	NewStreamController,
	NewBookmarksController,
	core.NewNowPlayingRepository,
	wire.FieldsOf(new(*Router), "Artwork", "Streamer", "Archiver", "DataStore", "ExternalInfo"),
)

func initSystemController(router *Router) *SystemController {
	panic(wire.Build(allProviders))
}

func initBrowsingController(router *Router) *BrowsingController {
	panic(wire.Build(allProviders))
}

func initAlbumListController(router *Router) *AlbumListController {
	panic(wire.Build(allProviders))
}

func initMediaAnnotationController(router *Router) *MediaAnnotationController {
	panic(wire.Build(allProviders))
}

func initPlaylistsController(router *Router) *PlaylistsController {
	panic(wire.Build(allProviders))
}

func initSearchingController(router *Router) *SearchingController {
	panic(wire.Build(allProviders))
}

func initUsersController(router *Router) *UsersController {
	panic(wire.Build(allProviders))
}

func initMediaRetrievalController(router *Router) *MediaRetrievalController {
	panic(wire.Build(allProviders))
}

func initStreamController(router *Router) *StreamController {
	panic(wire.Build(allProviders))
}

func initBookmarksController(router *Router) *BookmarksController {
	panic(wire.Build(allProviders))
}
