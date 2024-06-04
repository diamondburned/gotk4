#pragma once

#include <gio/gio.h>
#include <glib.h>

#include "go-common-stream.h"

struct _GoOutputStreamPrivate {
  guintptr id;
  gboolean isSeekable;
};

typedef struct _GoOutputStreamPrivate GoOutputStreamPrivate;

struct _GoOutputStreamClass {
  GOutputStreamClass parent_class;
};

typedef struct _GoOutputStreamClass GoOutputStreamClass;

struct _GoOutputStream {
  GOutputStream parent_instance;
  GoOutputStreamPrivate *priv;
};

typedef struct _GoOutputStream GoOutputStream;

static gboolean go_output_stream_seekable_init(GSeekableIface *iface);

G_DEFINE_TYPE_WITH_CODE(GoOutputStream, go_output_stream, G_TYPE_OUTPUT_STREAM,
                        G_ADD_PRIVATE(GoOutputStream)
                            G_IMPLEMENT_INTERFACE(G_TYPE_SEEKABLE, go_output_stream_seekable_init));

#define GO_OUTPUT_STREAM(o) \
  (G_TYPE_CHECK_INSTANCE_CAST((o), go_output_stream_get_type(), GoOutputStream))

static void go_output_stream_init(GoOutputStream *stream) {
  stream->priv = (GoOutputStreamPrivate *)go_output_stream_get_instance_private(stream);
};

static gssize go_output_stream_write(GOutputStream *stream, const void *buf, gsize len,
                                     GCancellable *cancel, GError **errOut) {
  GoOutputStream *gostream = GO_OUTPUT_STREAM(stream);
  GoOutputStreamPrivate *priv = gostream->priv;
  return goOutputStreamWrite(priv->id, buf, len, errOut);
};

static gboolean go_output_stream_close(GOutputStream *stream, GCancellable *cancel,
                                       GError **errOut) {
  GoOutputStream *gostream = GO_OUTPUT_STREAM(stream);
  GoOutputStreamPrivate *priv = gostream->priv;
  return goStreamClose(priv->id, errOut);
};

static void go_output_stream_class_init(GoOutputStreamClass *klass) {
  GOutputStreamClass *ostream_class = G_OUTPUT_STREAM_CLASS(klass);
  ostream_class->write_fn = go_output_stream_write;
  ostream_class->close_fn = go_output_stream_close;
};

GOutputStream *go_output_stream_new(guintptr id, gboolean isSeekable) {
  GoOutputStream *stream = (GoOutputStream *)g_object_new(go_output_stream_get_type(), NULL);

  GoOutputStreamPrivate *priv = stream->priv;
  priv->id = id;
  priv->isSeekable = isSeekable;

  return G_OUTPUT_STREAM(stream);
};

gboolean go_output_stream_can_seek(GSeekable *seekable) {
  GoOutputStream *gostream = GO_OUTPUT_STREAM(seekable);
  GoOutputStreamPrivate *priv = gostream->priv;
  return priv->isSeekable;
};

gboolean go_output_stream_seek(GSeekable *seekable, goffset offset, GSeekType type,
                               GCancellable *cancel, GError **errOut) {
  GoOutputStream *gostream = GO_OUTPUT_STREAM(seekable);
  GoOutputStreamPrivate *priv = gostream->priv;
  return go_stream_seek(priv->id, offset, type, cancel, errOut);
};

gboolean go_output_stream_can_truncate(GSeekable *seekable) { return FALSE; };

gboolean go_output_stream_truncate(GSeekable *seekable, goffset offset, GCancellable *cancel,
                                   GError **errOut) {
  return FALSE;
};

goffset go_output_stream_tell(GSeekable *seekable) {
  GoOutputStream *gostream = GO_OUTPUT_STREAM(seekable);
  GoOutputStreamPrivate *priv = gostream->priv;
  return go_stream_tell(priv->id);
};

gboolean go_output_stream_seekable_init(GSeekableIface *iface) {
  iface->can_seek = go_output_stream_can_seek;
  iface->seek = go_output_stream_seek;
  iface->can_truncate = go_output_stream_can_truncate;
  iface->truncate_fn = go_output_stream_truncate;
  iface->tell = go_output_stream_tell;
  return TRUE;
};
