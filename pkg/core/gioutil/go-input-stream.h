#include <gio/gio.h>
#include <glib.h>

#include "go-common-stream.h"

struct _GoInputStreamPrivate {
  guintptr id;
  gboolean isSeekable;
};

typedef struct _GoInputStreamPrivate GoInputStreamPrivate;

struct _GoInputStreamClass {
  GInputStreamClass parent_class;
};

typedef struct _GoInputStreamClass GoInputStreamClass;

struct _GoInputStream {
  GInputStream parent_instance;
  GoInputStreamPrivate *priv;
};

typedef struct _GoInputStream GoInputStream;

static gboolean go_input_stream_seekable_init(GSeekableIface *iface);

G_DEFINE_TYPE_WITH_CODE(GoInputStream, go_input_stream, G_TYPE_INPUT_STREAM,
                        G_ADD_PRIVATE(GoInputStream)
                            G_IMPLEMENT_INTERFACE(G_TYPE_SEEKABLE, go_input_stream_seekable_init));

#define GO_INPUT_STREAM(o) \
  (G_TYPE_CHECK_INSTANCE_CAST((o), go_input_stream_get_type(), GoInputStream))

static void go_input_stream_init(GoInputStream *stream) {
  stream->priv = (GoInputStreamPrivate *)go_input_stream_get_instance_private(stream);
};

static gssize go_input_stream_read(GInputStream *stream, void *buf, gsize len, GCancellable *cancel,
                                   GError **errOut) {
  GoInputStream *gostream = GO_INPUT_STREAM(stream);
  GoInputStreamPrivate *priv = gostream->priv;
  return goInputStreamRead(priv->id, buf, len, errOut);
};

static gboolean go_input_stream_close(GInputStream *stream, GCancellable *cancel, GError **errOut) {
  GoInputStream *gostream = GO_INPUT_STREAM(stream);
  GoInputStreamPrivate *priv = gostream->priv;
  return goStreamClose(priv->id, errOut);
};

static void go_input_stream_class_init(GoInputStreamClass *klass) {
  GInputStreamClass *istream_class = G_INPUT_STREAM_CLASS(klass);
  istream_class->read_fn = go_input_stream_read;
  istream_class->close_fn = go_input_stream_close;
};

GInputStream *go_input_stream_new(guintptr id, gboolean isSeekable) {
  GoInputStream *stream = (GoInputStream *)g_object_new(go_input_stream_get_type(), NULL);

  GoInputStreamPrivate *priv = stream->priv;
  priv->id = id;
  priv->isSeekable = isSeekable;

  return G_INPUT_STREAM(stream);
};

gboolean go_input_stream_can_seek(GSeekable *seekable) {
  GoInputStream *gostream = GO_INPUT_STREAM(seekable);
  GoInputStreamPrivate *priv = gostream->priv;
  return priv->isSeekable;
};

gboolean go_input_stream_seek(GSeekable *seekable, goffset offset, GSeekType type,
                              GCancellable *cancel, GError **errOut) {
  GoInputStream *gostream = GO_INPUT_STREAM(seekable);
  GoInputStreamPrivate *priv = gostream->priv;
  return go_stream_seek(priv->id, offset, type, cancel, errOut);
};

gboolean go_input_stream_can_truncate(GSeekable *seekable) { return FALSE; };

gboolean go_input_stream_truncate(GSeekable *seekable, goffset offset, GCancellable *cancel,
                                  GError **errOut) {
  return FALSE;
};

goffset go_input_stream_tell(GSeekable *seekable) {
  GoInputStream *gostream = GO_INPUT_STREAM(seekable);
  GoInputStreamPrivate *priv = gostream->priv;
  return go_stream_tell(priv->id);
};

gboolean go_input_stream_seekable_init(GSeekableIface *iface) {
  iface->can_seek = go_input_stream_can_seek;
  iface->seek = go_input_stream_seek;
  iface->can_truncate = go_input_stream_can_truncate;
  iface->truncate_fn = go_input_stream_truncate;
  iface->tell = go_input_stream_tell;
  return TRUE;
};
