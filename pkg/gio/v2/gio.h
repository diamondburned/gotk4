#include <stdlib.h>
#include <glib.h>

struct ActionEntry {
    void* name;
    void* activate;
    void* parameter_type;
    void* state;
    void* change_state;
    void  padding;
};
struct ActionGroupInterface {
          g_iface;
    void* has_action;
    void* list_actions;
    void* get_action_enabled;
    void* get_action_parameter_type;
    void* get_action_state_type;
    void* get_action_state_hint;
    void* get_action_state;
    void* change_action_state;
    void* activate_action;
    void* action_added;
    void* action_removed;
    void* action_enabled_changed;
    void* action_state_changed;
    void* query_action;
};
struct ActionInterface {
          g_iface;
    void* get_name;
    void* get_parameter_type;
    void* get_state_type;
    void* get_state_hint;
    void* get_enabled;
    void* get_state;
    void* change_state;
    void* activate;
};
struct ActionMapInterface {
          g_iface;
    void* lookup_action;
    void* add_action;
    void* remove_action;
};
struct AppInfoIface {
          g_iface;
    void* dup;
    void* equal;
    void* get_id;
    void* get_name;
    void* get_description;
    void* get_executable;
    void* get_icon;
    void* launch;
    void* supports_uris;
    void* supports_files;
    void* launch_uris;
    void* should_show;
    void* set_as_default_for_type;
    void* set_as_default_for_extension;
    void* add_supports_type;
    void* can_remove_supports_type;
    void* remove_supports_type;
    void* can_delete;
    void* do_delete;
    void* get_commandline;
    void* get_display_name;
    void* set_as_last_used_for_type;
    void* get_supported_types;
    void* launch_uris_async;
    void* launch_uris_finish;
};
struct AppLaunchContextClass {
          parent_class;
    void* get_display;
    void* get_startup_notify_id;
    void* launch_failed;
    void* launched;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
};
struct AppLaunchContextPrivate {

};
struct ApplicationClass {
          parent_class;
    void* startup;
    void* activate;
    void* open;
    void* command_line;
    void* local_command_line;
    void* before_emit;
    void* after_emit;
    void* add_platform_data;
    void* quit_mainloop;
    void* run_mainloop;
    void* shutdown;
    void* dbus_register;
    void* dbus_unregister;
    void* handle_local_options;
    void* name_lost;
    void  padding;
};
struct ApplicationCommandLineClass {
          parent_class;
    void* print_literal;
    void* printerr_literal;
    void* get_stdin;
    void  padding;
};
struct ApplicationCommandLinePrivate {

};
struct ApplicationPrivate {

};
struct AsyncInitableIface {
          g_iface;
    void* init_async;
    void* init_finish;
};
struct AsyncResultIface {
          g_iface;
    void* get_user_data;
    void* get_source_object;
    void* is_tagged;
};
struct BufferedInputStreamClass {
          parent_class;
    void* fill;
    void* fill_async;
    void* fill_finish;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct BufferedInputStreamPrivate {

};
struct BufferedOutputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
};
struct BufferedOutputStreamPrivate {

};
struct CancellableClass {
          parent_class;
    void* cancelled;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct CancellablePrivate {

};
struct CharsetConverterClass {
     parent_class;
};
struct ConverterIface {
          g_iface;
    void* convert;
    void* reset;
};
struct ConverterInputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct ConverterInputStreamPrivate {

};
struct ConverterOutputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct ConverterOutputStreamPrivate {

};
struct CredentialsClass {

};
struct DBusAnnotationInfo {
    gint   ref_count;
    void*  key;
    void*  value;
    void** annotations;
};
struct DBusArgInfo {
    gint   ref_count;
    void*  name;
    void*  signature;
    void** annotations;
};
struct DBusErrorEntry {
    gint  error_code;
    void* dbus_error_name;
};
struct DBusInterfaceIface {
          parent_iface;
    void* get_info;
    void* get_object;
    void* set_object;
    void* dup_object;
};
struct DBusInterfaceInfo {
    gint   ref_count;
    void*  name;
    void** methods;
    void** signals;
    void** properties;
    void** annotations;
};
struct DBusInterfaceSkeletonClass {
          parent_class;
    void* get_info;
    void* get_vtable;
    void* get_properties;
    void* flush;
    void  vfunc_padding;
    void* g_authorize_method;
    void  signal_padding;
};
struct DBusInterfaceSkeletonPrivate {

};
struct DBusInterfaceVTable {
    gpointer method_call;
    gpointer get_property;
    gpointer set_property;
    void     padding;
};
struct DBusMethodInfo {
    gint   ref_count;
    void*  name;
    void** in_args;
    void** out_args;
    void** annotations;
};
struct DBusNodeInfo {
    gint   ref_count;
    void*  path;
    void** interfaces;
    void** nodes;
    void** annotations;
};
struct DBusObjectIface {
          parent_iface;
    void* get_object_path;
    void* get_interfaces;
    void* get_interface;
    void* interface_added;
    void* interface_removed;
};
struct DBusObjectManagerClientClass {
          parent_class;
    void* interface_proxy_signal;
    void* interface_proxy_properties_changed;
    void  padding;
};
struct DBusObjectManagerClientPrivate {

};
struct DBusObjectManagerIface {
          parent_iface;
    void* get_object_path;
    void* get_objects;
    void* get_object;
    void* get_interface;
    void* object_added;
    void* object_removed;
    void* interface_added;
    void* interface_removed;
};
struct DBusObjectManagerServerClass {
         parent_class;
    void padding;
};
struct DBusObjectManagerServerPrivate {

};
struct DBusObjectProxyClass {
         parent_class;
    void padding;
};
struct DBusObjectProxyPrivate {

};
struct DBusObjectSkeletonClass {
          parent_class;
    void* authorize_method;
    void  padding;
};
struct DBusObjectSkeletonPrivate {

};
struct DBusPropertyInfo {
    gint   ref_count;
    void*  name;
    void*  signature;
           flags;
    void** annotations;
};
struct DBusProxyClass {
          parent_class;
    void* g_properties_changed;
    void* g_signal;
    void  padding;
};
struct DBusProxyPrivate {

};
struct DBusSignalInfo {
    gint   ref_count;
    void*  name;
    void** args;
    void** annotations;
};
struct DBusSubtreeVTable {
    gpointer enumerate;
    gpointer introspect;
    gpointer dispatch;
    void     padding;
};
struct DataInputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct DataInputStreamPrivate {

};
struct DataOutputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct DataOutputStreamPrivate {

};
struct DatagramBasedInterface {
          g_iface;
    void* receive_messages;
    void* send_messages;
    void* create_source;
    void* condition_check;
    void* condition_wait;
};
struct DesktopAppInfoClass {
     parent_class;
};
struct DesktopAppInfoLookupIface {
          g_iface;
    void* get_default_for_uri_scheme;
};
struct DriveIface {
          g_iface;
    void* changed;
    void* disconnected;
    void* eject_button;
    void* get_name;
    void* get_icon;
    void* has_volumes;
    void* get_volumes;
    void* is_media_removable;
    void* has_media;
    void* is_media_check_automatic;
    void* can_eject;
    void* can_poll_for_media;
    void* eject;
    void* eject_finish;
    void* poll_for_media;
    void* poll_for_media_finish;
    void* get_identifier;
    void* enumerate_identifiers;
    void* get_start_stop_type;
    void* can_start;
    void* can_start_degraded;
    void* start;
    void* start_finish;
    void* can_stop;
    void* stop;
    void* stop_finish;
    void* stop_button;
    void* eject_with_operation;
    void* eject_with_operation_finish;
    void* get_sort_key;
    void* get_symbolic_icon;
    void* is_removable;
};
struct DtlsClientConnectionInterface {
     g_iface;
};
struct DtlsConnectionInterface {
          g_iface;
    void* accept_certificate;
    void* handshake;
    void* handshake_async;
    void* handshake_finish;
    void* shutdown;
    void* shutdown_async;
    void* shutdown_finish;
    void* set_advertised_protocols;
    void* get_negotiated_protocol;
    void* get_binding_data;
};
struct DtlsServerConnectionInterface {
     g_iface;
};
struct EmblemClass {

};
struct EmblemedIconClass {
     parent_class;
};
struct EmblemedIconPrivate {

};
struct FileAttributeInfo {
    void* name;
          type;
          flags;
};
struct FileAttributeInfoList {
    void* infos;
    int   n_infos;
};
struct FileAttributeMatcher {

};
struct FileDescriptorBasedIface {
          g_iface;
    void* get_fd;
};
struct FileEnumeratorClass {
          parent_class;
    void* next_file;
    void* close_fn;
    void* next_files_async;
    void* next_files_finish;
    void* close_async;
    void* close_finish;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
    void* _g_reserved7;
};
struct FileEnumeratorPrivate {

};
struct FileIOStreamClass {
          parent_class;
    void* tell;
    void* can_seek;
    void* seek;
    void* can_truncate;
    void* truncate_fn;
    void* query_info;
    void* query_info_async;
    void* query_info_finish;
    void* get_etag;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct FileIOStreamPrivate {

};
struct FileIconClass {

};
struct FileIface {
             g_iface;
    void*    dup;
    void*    hash;
    void*    equal;
    void*    is_native;
    void*    has_uri_scheme;
    void*    get_uri_scheme;
    void*    get_basename;
    void*    get_path;
    void*    get_uri;
    void*    get_parse_name;
    void*    get_parent;
    void*    prefix_matches;
    void*    get_relative_path;
    void*    resolve_relative_path;
    void*    get_child_for_display_name;
    void*    enumerate_children;
    void*    enumerate_children_async;
    void*    enumerate_children_finish;
    void*    query_info;
    void*    query_info_async;
    void*    query_info_finish;
    void*    query_filesystem_info;
    void*    query_filesystem_info_async;
    void*    query_filesystem_info_finish;
    void*    find_enclosing_mount;
    void*    find_enclosing_mount_async;
    void*    find_enclosing_mount_finish;
    void*    set_display_name;
    void*    set_display_name_async;
    void*    set_display_name_finish;
    void*    query_settable_attributes;
    void*    _query_settable_attributes_async;
    void*    _query_settable_attributes_finish;
    void*    query_writable_namespaces;
    void*    _query_writable_namespaces_async;
    void*    _query_writable_namespaces_finish;
    void*    set_attribute;
    void*    set_attributes_from_info;
    void*    set_attributes_async;
    void*    set_attributes_finish;
    void*    read_fn;
    void*    read_async;
    void*    read_finish;
    void*    append_to;
    void*    append_to_async;
    void*    append_to_finish;
    void*    create;
    void*    create_async;
    void*    create_finish;
    void*    replace;
    void*    replace_async;
    void*    replace_finish;
    void*    delete_file;
    void*    delete_file_async;
    void*    delete_file_finish;
    void*    trash;
    void*    trash_async;
    void*    trash_finish;
    void*    make_directory;
    void*    make_directory_async;
    void*    make_directory_finish;
    void*    make_symbolic_link;
    void*    _make_symbolic_link_async;
    void*    _make_symbolic_link_finish;
    void*    copy;
    void*    copy_async;
    void*    copy_finish;
    void*    move;
    void*    _move_async;
    void*    _move_finish;
    void*    mount_mountable;
    void*    mount_mountable_finish;
    void*    unmount_mountable;
    void*    unmount_mountable_finish;
    void*    eject_mountable;
    void*    eject_mountable_finish;
    void*    mount_enclosing_volume;
    void*    mount_enclosing_volume_finish;
    void*    monitor_dir;
    void*    monitor_file;
    void*    open_readwrite;
    void*    open_readwrite_async;
    void*    open_readwrite_finish;
    void*    create_readwrite;
    void*    create_readwrite_async;
    void*    create_readwrite_finish;
    void*    replace_readwrite;
    void*    replace_readwrite_async;
    void*    replace_readwrite_finish;
    void*    start_mountable;
    void*    start_mountable_finish;
    void*    stop_mountable;
    void*    stop_mountable_finish;
    gboolean supports_thread_contexts;
    void*    unmount_mountable_with_operation;
    void*    unmount_mountable_with_operation_finish;
    void*    eject_mountable_with_operation;
    void*    eject_mountable_with_operation_finish;
    void*    poll_mountable;
    void*    poll_mountable_finish;
    void*    measure_disk_usage;
    void*    measure_disk_usage_async;
    void*    measure_disk_usage_finish;
};
struct FileInfoClass {

};
struct FileInputStreamClass {
          parent_class;
    void* tell;
    void* can_seek;
    void* seek;
    void* query_info;
    void* query_info_async;
    void* query_info_finish;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct FileInputStreamPrivate {

};
struct FileMonitorClass {
          parent_class;
    void* changed;
    void* cancel;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct FileMonitorPrivate {

};
struct FileOutputStreamClass {
          parent_class;
    void* tell;
    void* can_seek;
    void* seek;
    void* can_truncate;
    void* truncate_fn;
    void* query_info;
    void* query_info_async;
    void* query_info_finish;
    void* get_etag;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct FileOutputStreamPrivate {

};
struct FilenameCompleterClass {
          parent_class;
    void* got_completion_data;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
};
struct FilterInputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
};
struct FilterOutputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
};
struct IOExtension {

};
struct IOExtensionPoint {

};
struct IOModuleClass {

};
struct IOModuleScope {

};
struct IOSchedulerJob {

};
struct IOStreamAdapter {

};
struct IOStreamClass {
          parent_class;
    void* get_input_stream;
    void* get_output_stream;
    void* close_fn;
    void* close_async;
    void* close_finish;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
    void* _g_reserved7;
    void* _g_reserved8;
    void* _g_reserved9;
    void* _g_reserved10;
};
struct IOStreamPrivate {

};
struct IconIface {
          g_iface;
    void* hash;
    void* equal;
    void* to_tokens;
    void* from_tokens;
    void* serialize;
};
struct InetAddressClass {
          parent_class;
    void* to_string;
    void* to_bytes;
};
struct InetAddressMaskClass {
     parent_class;
};
struct InetAddressMaskPrivate {

};
struct InetAddressPrivate {

};
struct InetSocketAddressClass {
     parent_class;
};
struct InetSocketAddressPrivate {

};
struct InitableIface {
          g_iface;
    void* init;
};
struct InputMessage {
    void**  address;
    void*   vectors;
    guint   num_vectors;
    gsize   bytes_received;
    gint    flags;
    void*** control_messages;
    void*   num_control_messages;
};
struct InputStreamClass {
          parent_class;
    void* read_fn;
    void* skip;
    void* close_fn;
    void* read_async;
    void* read_finish;
    void* skip_async;
    void* skip_finish;
    void* close_async;
    void* close_finish;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct InputStreamPrivate {

};
struct InputVector {
    gpointer buffer;
    gsize    size;
};
struct ListModelInterface {
          g_iface;
    void* get_item_type;
    void* get_n_items;
    void* get_item;
};
struct ListStoreClass {
     parent_class;
};
struct LoadableIconIface {
          g_iface;
    void* load;
    void* load_async;
    void* load_finish;
};
struct MemoryInputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct MemoryInputStreamPrivate {

};
struct MemoryMonitorInterface {
          g_iface;
    void* low_memory_warning;
};
struct MemoryOutputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct MemoryOutputStreamPrivate {

};
struct MenuAttributeIterClass {
          parent_class;
    void* get_next;
};
struct MenuAttributeIterPrivate {

};
struct MenuLinkIterClass {
          parent_class;
    void* get_next;
};
struct MenuLinkIterPrivate {

};
struct MenuModelClass {
          parent_class;
    void* is_mutable;
    void* get_n_items;
    void* get_item_attributes;
    void* iterate_item_attributes;
    void* get_item_attribute_value;
    void* get_item_links;
    void* iterate_item_links;
    void* get_item_link;
};
struct MenuModelPrivate {

};
struct MountIface {
          g_iface;
    void* changed;
    void* unmounted;
    void* get_root;
    void* get_name;
    void* get_icon;
    void* get_uuid;
    void* get_volume;
    void* get_drive;
    void* can_unmount;
    void* can_eject;
    void* unmount;
    void* unmount_finish;
    void* eject;
    void* eject_finish;
    void* remount;
    void* remount_finish;
    void* guess_content_type;
    void* guess_content_type_finish;
    void* guess_content_type_sync;
    void* pre_unmount;
    void* unmount_with_operation;
    void* unmount_with_operation_finish;
    void* eject_with_operation;
    void* eject_with_operation_finish;
    void* get_default_location;
    void* get_sort_key;
    void* get_symbolic_icon;
};
struct MountOperationClass {
          parent_class;
    void* ask_password;
    void* ask_question;
    void* reply;
    void* aborted;
    void* show_processes;
    void* show_unmount_progress;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
    void* _g_reserved7;
    void* _g_reserved8;
    void* _g_reserved9;
};
struct MountOperationPrivate {

};
struct NativeSocketAddressClass {
     parent_class;
};
struct NativeSocketAddressPrivate {

};
struct NativeVolumeMonitorClass {
          parent_class;
    void* get_mount_for_mount_path;
};
struct NetworkAddressClass {
     parent_class;
};
struct NetworkAddressPrivate {

};
struct NetworkMonitorInterface {
          g_iface;
    void* network_changed;
    void* can_reach;
    void* can_reach_async;
    void* can_reach_finish;
};
struct NetworkServiceClass {
     parent_class;
};
struct NetworkServicePrivate {

};
struct OutputMessage {
    void*  address;
    void*  vectors;
    guint  num_vectors;
    guint  bytes_sent;
    void** control_messages;
    guint  num_control_messages;
};
struct OutputStreamClass {
          parent_class;
    void* write_fn;
    void* splice;
    void* flush;
    void* close_fn;
    void* write_async;
    void* write_finish;
    void* splice_async;
    void* splice_finish;
    void* flush_async;
    void* flush_finish;
    void* close_async;
    void* close_finish;
    void* writev_fn;
    void* writev_async;
    void* writev_finish;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
    void* _g_reserved7;
    void* _g_reserved8;
};
struct OutputStreamPrivate {

};
struct OutputVector {
    gpointer buffer;
    gsize    size;
};
struct PermissionClass {
          parent_class;
    void* acquire;
    void* acquire_async;
    void* acquire_finish;
    void* release;
    void* release_async;
    void* release_finish;
    void  reserved;
};
struct PermissionPrivate {

};
struct PollableInputStreamInterface {
          g_iface;
    void* can_poll;
    void* is_readable;
    void* create_source;
    void* read_nonblocking;
};
struct PollableOutputStreamInterface {
          g_iface;
    void* can_poll;
    void* is_writable;
    void* create_source;
    void* write_nonblocking;
    void* writev_nonblocking;
};
struct ProxyAddressClass {
     parent_class;
};
struct ProxyAddressEnumeratorClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
    void* _g_reserved7;
};
struct ProxyAddressEnumeratorPrivate {

};
struct ProxyAddressPrivate {

};
struct ProxyInterface {
          g_iface;
    void* connect;
    void* connect_async;
    void* connect_finish;
    void* supports_hostname;
};
struct ProxyResolverInterface {
          g_iface;
    void* is_supported;
    void* lookup;
    void* lookup_async;
    void* lookup_finish;
};
struct RemoteActionGroupInterface {
          g_iface;
    void* activate_action_full;
    void* change_action_state_full;
};
struct ResolverClass {
          parent_class;
    void* reload;
    void* lookup_by_name;
    void* lookup_by_name_async;
    void* lookup_by_name_finish;
    void* lookup_by_address;
    void* lookup_by_address_async;
    void* lookup_by_address_finish;
    void* lookup_service;
    void* lookup_service_async;
    void* lookup_service_finish;
    void* lookup_records;
    void* lookup_records_async;
    void* lookup_records_finish;
    void* lookup_by_name_with_flags_async;
    void* lookup_by_name_with_flags_finish;
    void* lookup_by_name_with_flags;
};
struct ResolverPrivate {

};
struct Resource {

};
struct SeekableIface {
          g_iface;
    void* tell;
    void* can_seek;
    void* seek;
    void* can_truncate;
    void* truncate_fn;
};
struct SettingsBackendClass {
          parent_class;
    void* read;
    void* get_writable;
    void* write;
    void* write_tree;
    void* reset;
    void* subscribe;
    void* unsubscribe;
    void* sync;
    void* get_permission;
    void* read_user_value;
    void  padding;
};
struct SettingsBackendPrivate {

};
struct SettingsClass {
          parent_class;
    void* writable_changed;
    void* changed;
    void* writable_change_event;
    void* change_event;
    void  padding;
};
struct SettingsPrivate {

};
struct SettingsSchema {

};
struct SettingsSchemaKey {

};
struct SettingsSchemaSource {

};
struct SimpleActionGroupClass {
         parent_class;
    void padding;
};
struct SimpleActionGroupPrivate {

};
struct SimpleAsyncResultClass {

};
struct SimpleProxyResolverClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct SimpleProxyResolverPrivate {

};
struct SocketAddressClass {
          parent_class;
    void* get_family;
    void* get_native_size;
    void* to_native;
};
struct SocketAddressEnumeratorClass {
          parent_class;
    void* next;
    void* next_async;
    void* next_finish;
};
struct SocketClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
    void* _g_reserved7;
    void* _g_reserved8;
    void* _g_reserved9;
    void* _g_reserved10;
};
struct SocketClientClass {
          parent_class;
    void* event;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
};
struct SocketClientPrivate {

};
struct SocketConnectableIface {
          g_iface;
    void* enumerate;
    void* proxy_enumerate;
    void* to_string;
};
struct SocketConnectionClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
};
struct SocketConnectionPrivate {

};
struct SocketControlMessageClass {
          parent_class;
    void* get_size;
    void* get_level;
    void* get_type;
    void* serialize;
    void* deserialize;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct SocketControlMessagePrivate {

};
struct SocketListenerClass {
          parent_class;
    void* changed;
    void* event;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
};
struct SocketListenerPrivate {

};
struct SocketPrivate {

};
struct SocketServiceClass {
          parent_class;
    void* incoming;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
};
struct SocketServicePrivate {

};
struct SrvTarget {

};
struct StaticResource {
    void*    data;
    gsize    data_len;
    void*    resource;
    void*    next;
    gpointer padding;
};
struct TaskClass {

};
struct TcpConnectionClass {
     parent_class;
};
struct TcpConnectionPrivate {

};
struct TcpWrapperConnectionClass {
     parent_class;
};
struct TcpWrapperConnectionPrivate {

};
struct ThemedIconClass {

};
struct ThreadedSocketServiceClass {
          parent_class;
    void* run;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct ThreadedSocketServicePrivate {

};
struct TlsBackendInterface {
          g_iface;
    void* supports_tls;
    void* get_certificate_type;
    void* get_client_connection_type;
    void* get_server_connection_type;
    void* get_file_database_type;
    void* get_default_database;
    void* supports_dtls;
    void* get_dtls_client_connection_type;
    void* get_dtls_server_connection_type;
};
struct TlsCertificateClass {
          parent_class;
    void* verify;
    void  padding;
};
struct TlsCertificatePrivate {

};
struct TlsClientConnectionInterface {
          g_iface;
    void* copy_session_state;
};
struct TlsConnectionClass {
          parent_class;
    void* accept_certificate;
    void* handshake;
    void* handshake_async;
    void* handshake_finish;
    void* get_binding_data;
    void  padding;
};
struct TlsConnectionPrivate {

};
struct TlsDatabaseClass {
          parent_class;
    void* verify_chain;
    void* verify_chain_async;
    void* verify_chain_finish;
    void* create_certificate_handle;
    void* lookup_certificate_for_handle;
    void* lookup_certificate_for_handle_async;
    void* lookup_certificate_for_handle_finish;
    void* lookup_certificate_issuer;
    void* lookup_certificate_issuer_async;
    void* lookup_certificate_issuer_finish;
    void* lookup_certificates_issued_by;
    void* lookup_certificates_issued_by_async;
    void* lookup_certificates_issued_by_finish;
    void  padding;
};
struct TlsDatabasePrivate {

};
struct TlsFileDatabaseInterface {
         g_iface;
    void padding;
};
struct TlsInteractionClass {
          parent_class;
    void* ask_password;
    void* ask_password_async;
    void* ask_password_finish;
    void* request_certificate;
    void* request_certificate_async;
    void* request_certificate_finish;
    void  padding;
};
struct TlsInteractionPrivate {

};
struct TlsPasswordClass {
          parent_class;
    void* get_value;
    void* set_value;
    void* get_default_warning;
    void  padding;
};
struct TlsPasswordPrivate {

};
struct TlsServerConnectionInterface {
     g_iface;
};
struct UnixConnectionClass {
     parent_class;
};
struct UnixConnectionPrivate {

};
struct UnixCredentialsMessageClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
};
struct UnixCredentialsMessagePrivate {

};
struct UnixFDListClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct UnixFDListPrivate {

};
struct UnixFDMessageClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
};
struct UnixFDMessagePrivate {

};
struct UnixInputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct UnixInputStreamPrivate {

};
struct UnixMountEntry {

};
struct UnixMountMonitorClass {

};
struct UnixMountPoint {

};
struct UnixOutputStreamClass {
          parent_class;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
};
struct UnixOutputStreamPrivate {

};
struct UnixSocketAddressClass {
     parent_class;
};
struct UnixSocketAddressPrivate {

};
struct VfsClass {
          parent_class;
    void* is_active;
    void* get_file_for_path;
    void* get_file_for_uri;
    void* get_supported_uri_schemes;
    void* parse_name;
    void* local_file_add_info;
    void* add_writable_namespaces;
    void* local_file_set_attributes;
    void* local_file_removed;
    void* local_file_moved;
    void* deserialize_icon;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
};
struct VolumeIface {
          g_iface;
    void* changed;
    void* removed;
    void* get_name;
    void* get_icon;
    void* get_uuid;
    void* get_drive;
    void* get_mount;
    void* can_mount;
    void* can_eject;
    void* mount_fn;
    void* mount_finish;
    void* eject;
    void* eject_finish;
    void* get_identifier;
    void* enumerate_identifiers;
    void* should_automount;
    void* get_activation_root;
    void* eject_with_operation;
    void* eject_with_operation_finish;
    void* get_sort_key;
    void* get_symbolic_icon;
};
struct VolumeMonitorClass {
          parent_class;
    void* volume_added;
    void* volume_removed;
    void* volume_changed;
    void* mount_added;
    void* mount_removed;
    void* mount_pre_unmount;
    void* mount_changed;
    void* drive_connected;
    void* drive_disconnected;
    void* drive_changed;
    void* is_supported;
    void* get_connected_drives;
    void* get_volumes;
    void* get_mounts;
    void* get_volume_for_uuid;
    void* get_mount_for_uuid;
    void* adopt_orphan_mount;
    void* drive_eject_button;
    void* drive_stop_button;
    void* _g_reserved1;
    void* _g_reserved2;
    void* _g_reserved3;
    void* _g_reserved4;
    void* _g_reserved5;
    void* _g_reserved6;
};
struct ZlibCompressorClass {
     parent_class;
};
struct ZlibDecompressorClass {
     parent_class;
};