#include "go-common-stream.h"

#include <gio/gio.h>
#include <glib.h>

const gint goStreamSeekWhence[3] = {
    [G_SEEK_SET] = 0,  // SeekStart
    [G_SEEK_CUR] = 1,  // SeekCurrent
    [G_SEEK_END] = 2,  // SeekEnd
};

gboolean go_stream_seek(guintptr id, goffset offset, GSeekType type, GCancellable *cancellable,
                        GError **error) {
  int whence = goStreamSeekWhence[type];
  gssize res = goStreamSeek(id, offset, whence, error);
  return res >= 0;
};

goffset go_stream_tell(guintptr id) {
  gssize res = goStreamSeek(id, 0, 1, NULL);
  return MAX(res, 0);
};
