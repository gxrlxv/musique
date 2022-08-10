# Music service
Artist service - that allows users to interact with their playlists 
## Music service API

| Endpoint                            | HTTP method | Description                                     |
|-------------------------------------|-------------|-------------------------------------------------|
| /playlists/{playlist_id}/{track_id} | POST        | Add track to a playlist by track id             |
| /playlists/{playlist_id}/{album_id} | POST        | Add tracks from album to a playlist by album id |
| /playlists/{playlist_id}/{track_id} | DELETE      | Delete track from playlist                      |
| /playlists/{playlist_id}/{track_id} | GET         | Get track from playlist by track id             |
| /playlists/{playlist_id}            | GET         | Get all tracks from playlist                    |