#include <gio/gfiledescriptorbased.h>
#include <gio/gio.h>
#include <gio/gunixconnection.h>
#include <gio/gunixcredentialsmessage.h>
#include <gio/gunixfdlist.h>
#include <gio/gunixfdmessage.h>
#include <gio/gunixinputstream.h>
#include <gio/gunixmounts.h>
#include <gio/gunixoutputstream.h>
#include <gio/gunixsocketaddress.h>
#include <glib-object.h>
GDBusMessage* _gotk4_gio2_DBusMessageFilterFunction(GDBusConnection*, GDBusMessage*, gboolean, gpointer);
GFile* _gotk4_gio2_VFSFileLookupFunc(GVfs*, char*, gpointer);
GType _gotk4_gio2_DBusProxyTypeFunc(GDBusObjectManagerClient*, gchar*, gchar*, gpointer);
extern void callbackDelete(gpointer);
gboolean _gotk4_gio2_SettingsGetMapping(GVariant*, gpointer*, gpointer);
gint _gotk4_glib2_CompareDataFunc(gconstpointer, gconstpointer, gpointer);
void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
void _gotk4_gio2_DBusSignalCallback(GDBusConnection*, gchar*, gchar*, gchar*, gchar*, GVariant*, gpointer);
void _gotk4_gio2_FileProgressCallback(goffset, goffset, gpointer);