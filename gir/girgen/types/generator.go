package types

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
)

// ModulePathFunc returns the Go module import path from the given namespace.
// See Generator.ModPath for mroe information.
type ModulePathFunc func(*gir.Namespace) string

// FileGenerator defines a generator instance.
type FileGenerator interface {
	logger.LineLogger
	// CanGenerate checks if a type is going to be generated or not. It is used
	// primarily during type resolving.
	CanGenerate(*Resolved) bool
	// Filters returns the list of matchers that the current generator has.
	Filters() []FilterMatcher
	// ModPath crafts an import path from the given GIR namespace. The import
	// path is assumed to have the same package name as the base file, but major
	// versions are exempted as an edge case.
	ModPath(*gir.Namespace) string
	// Repositories returns the list of known repositories inside the generator.
	Repositories() gir.Repositories
	// Namespace returns the generator's current namespace.
	Namespace() *gir.NamespaceFindResult
	// LinkMode gets the current link mode of the file generator.
	LinkMode() LinkMode
}

type wrappedGenerator struct {
	FileGenerator
	n *gir.NamespaceFindResult
}

func (w wrappedGenerator) Namespace() *gir.NamespaceFindResult { return w.n }

// OverrideNamespace returns a new generator that overrides a generator's
// current namespace.
func OverrideNamespace(gen FileGenerator, nsp *gir.NamespaceFindResult) FileGenerator {
	return wrappedGenerator{gen, nsp}
}

// Find finds the given GIR type from the given generator.
func Find(gen FileGenerator, girType string) *gir.TypeFindResult {
	if girType == "\x00" {
		return nil
	}

	return gen.Repositories().FindType(gen.Namespace(), girType)
}

// AddCallbackHeader is a convenient function around AddCallableHeader that
// takes in a Callback.
func AddCallbackHeader(h *file.Header, source *gir.NamespaceFindResult, callback *gir.Callback) {
	AddCallableHeader(h, source, "", &callback.CallableAttrs)
}

// AddCallableHeader adds an extern C function header from the callable. The
// extern function will have the given name.
func AddCallableHeader(h *file.Header, source *gir.NamespaceFindResult, name string, callable *gir.CallableAttrs) {
	h.AddCallbackHeader(CallableCHeader(source, name, callable))
}

// CallableCHeader renders the C function signature.
func CallableCHeader(source *gir.NamespaceFindResult, name string, callable *gir.CallableAttrs) string {
	var ctail pen.Joints
	if callable.Parameters != nil {
		ctail = pen.NewJoints(", ", len(callable.Parameters.Parameters)+1)

		if callable.Parameters.InstanceParameter != nil {
			ctail.Add(AnyTypeC(callable.Parameters.InstanceParameter.AnyType))
		}
		for _, param := range callable.Parameters.Parameters {
			ctail.Add(AnyTypeC(param.AnyType))
		}
		if callable.Throws {
			ctail.Add("GError**")
		}
	}

	cReturn := "void"
	if callable.ReturnValue != nil {
		cReturn = AnyTypeC(callable.ReturnValue.AnyType)
	}

	if name == "" {
		name = file.CallableExportedName(source, callable)
	}

	return fmt.Sprintf("extern %s %s(%s);", cReturn, name, ctail.Join())
}
