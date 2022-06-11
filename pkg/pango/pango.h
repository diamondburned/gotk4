#include <stdlib.h>
#include <glib.h>

struct Analysis {
    gpointer shape_engine;
    gpointer lang_engine;
    void*    font;
    guint8   level;
    guint8   gravity;
    guint8   flags;
    guint8   script;
    void*    language;
    void*    extra_attrs;
};
struct AttrClass {
          type;
    void* copy;
    void* destroy;
    void* equal;
};
struct AttrColor {
     attr;
     color;
};
struct AttrFloat {
           attr;
    double value;
};
struct AttrFontDesc {
          attr;
    void* desc;
};
struct AttrFontFeatures {
          attr;
    void* features;
};
struct AttrInt {
        attr;
    int value;
};
struct AttrIterator {

};
struct AttrLanguage {
          attr;
    void* value;
};
struct AttrList {

};
struct AttrShape {
             attr;
             ink_rect;
             logical_rect;
    gpointer data;
    gpointer copy_func;
             destroy_func;
};
struct AttrSize {
          attr;
    int   size;
    guint absolute  : 1;
};
struct AttrString {
          attr;
    void* value;
};
struct Attribute {
    void* klass;
    guint start_index;
    guint end_index;
};
struct Color {
    guint16 red;
    guint16 green;
    guint16 blue;
};
struct ContextClass {

};
struct FontClass {
          parent_class;
    void* describe;
    void* get_coverage;
    void* get_glyph_extents;
    void* get_metrics;
    void* get_font_map;
    void* describe_absolute;
    void* get_features;
    void* create_hb_font;
};
struct FontDescription {

};
struct FontFaceClass {
          parent_class;
    void* get_face_name;
    void* describe;
    void* list_sizes;
    void* is_synthesized;
    void* get_family;
    void* _pango_reserved3;
    void* _pango_reserved4;
};
struct FontFamilyClass {
          parent_class;
    void* list_faces;
    void* get_name;
    void* is_monospace;
    void* is_variable;
    void* get_face;
    void* _pango_reserved2;
};
struct FontMapClass {
          parent_class;
    void* load_font;
    void* list_families;
    void* load_fontset;
    void* shape_engine_type;
    void* get_serial;
    void* changed;
    void* get_family;
    void* get_face;
};
struct FontMetrics {
    guint ref_count;
    int   ascent;
    int   descent;
    int   height;
    int   approximate_char_width;
    int   approximate_digit_width;
    int   underline_position;
    int   underline_thickness;
    int   strikethrough_position;
    int   strikethrough_thickness;
};
struct FontsetClass {
          parent_class;
    void* get_font;
    void* get_metrics;
    void* get_language;
    void* foreach;
    void* _pango_reserved1;
    void* _pango_reserved2;
    void* _pango_reserved3;
    void* _pango_reserved4;
};
struct FontsetSimpleClass {

};
struct GlyphGeometry {
    gint32 width;
    gint32 x_offset;
    gint32 y_offset;
};
struct GlyphInfo {
    guint32 glyph;
            geometry;
            attr;
};
struct GlyphItem {
    void* item;
    void* glyphs;
};
struct GlyphItemIter {
    void* glyph_item;
    void* text;
    int   start_glyph;
    int   start_index;
    int   start_char;
    int   end_glyph;
    int   end_index;
    int   end_char;
};
struct GlyphString {
    gint  num_glyphs;
    void* glyphs;
    void* log_clusters;
    gint  space;
};
struct GlyphVisAttr {
    guint is_cluster_start  : 1;
};
struct Item {
    gint offset;
    gint length;
    gint num_chars;
         analysis;
};
struct Language {

};
struct LayoutClass {

};
struct LayoutIter {

};
struct LayoutLine {
    void* layout;
    gint  start_index;
    gint  length;
    void* runs;
    guint is_paragraph_start  : 1;
    guint resolved_dir        : 3;
};
struct LogAttr {
    guint is_line_break                : 1;
    guint is_mandatory_break           : 1;
    guint is_char_break                : 1;
    guint is_white                     : 1;
    guint is_cursor_position           : 1;
    guint is_word_start                : 1;
    guint is_word_end                  : 1;
    guint is_sentence_boundary         : 1;
    guint is_sentence_start            : 1;
    guint is_sentence_end              : 1;
    guint backspace_deletes_character  : 1;
    guint is_expandable_space          : 1;
    guint is_word_boundary             : 1;
};
struct Matrix {
    double xx;
    double xy;
    double yx;
    double yy;
    double x0;
    double y0;
};
struct Rectangle {
    int x;
    int y;
    int width;
    int height;
};
struct RendererClass {
          parent_class;
    void* draw_glyphs;
    void* draw_rectangle;
    void* draw_error_underline;
    void* draw_shape;
    void* draw_trapezoid;
    void* draw_glyph;
    void* part_changed;
    void* begin;
    void* end;
    void* prepare_run;
    void* draw_glyph_item;
    void* _pango_reserved2;
    void* _pango_reserved3;
    void* _pango_reserved4;
};
struct RendererPrivate {

};
struct ScriptIter {

};
struct TabArray {

};