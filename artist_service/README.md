# Artist service
Artist service - service that allows artists to interact with their albums and tracks
## Artist service API

| Endpoint                             | HTTP method | Description             |
|--------------------------------------|-------------|-------------------------|
| /artist/albums                       | POST        | Create album            |   
| /artist/albums/{album_id}            | DELETE      | Delete album by id      |
| /artist/albums/{album_id}            | POST        | Add track to album      |
| /artist/albums/{album_id}/{track_id} | DELETE      | Delete track from album |