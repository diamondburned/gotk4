#pragma once

#include <gio/gio.h>
#include <glib.h>

extern gssize goInputStreamRead(guintptr id, void *buf, gsize len, GError **errOut);  // io.Reader
extern gssize goOutputStreamWrite(guintptr id, const void *buf, gsize len,
                                  GError **errOut);                                     // io.Writer
extern gssize goStreamSeek(guintptr id, goffset offset, gint whence, GError **errOut);  // io.Seeker
extern gboolean goStreamClose(guintptr id, GError **errOut);                            // io.Closer

gboolean go_stream_seek(guintptr id, goffset offset, GSeekType type, GCancellable *cancellable,
                        GError **error);

goffset go_stream_tell(guintptr id);
