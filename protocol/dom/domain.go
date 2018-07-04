// Code generated by cdpgen. DO NOT EDIT.

// Package dom implements the DOM domain. This domain exposes DOM read/write
// operations. Each DOM Node is represented with its mirror object that has an
// `id`. This `id` can be used to get additional information on the Node,
// resolve it into the JavaScript object wrapper, etc. It is important that
// client receives DOM events only for the nodes that are known to the client.
// Backend keeps track of the nodes that were sent to the client and never
// sends the same node twice. It is client's responsibility to collect
// information about the nodes that were sent to the client.
//
// Note that `iframe` owner elements will return corresponding document
// elements as their child nodes.
package dom

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the DOM domain. This domain exposes DOM
// read/write operations. Each DOM Node is represented with its mirror object
// that has an `id`. This `id` can be used to get additional information on the
// Node, resolve it into the JavaScript object wrapper, etc. It is important
// that client receives DOM events only for the nodes that are known to the
// client. Backend keeps track of the nodes that were sent to the client and
// never sends the same node twice. It is client's responsibility to collect
// information about the nodes that were sent to the client.
//
// Note that `iframe` owner elements will return corresponding document
// elements as their child nodes.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the DOM domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// CollectClassNamesFromSubtree invokes the DOM method. Collects class names
// for the node with given id and all of it's child nodes.
func (d *domainClient) CollectClassNamesFromSubtree(ctx context.Context, args *CollectClassNamesFromSubtreeArgs) (reply *CollectClassNamesFromSubtreeReply, err error) {
	reply = new(CollectClassNamesFromSubtreeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.collectClassNamesFromSubtree", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.collectClassNamesFromSubtree", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "CollectClassNamesFromSubtree", Err: err}
	}
	return
}

// CopyTo invokes the DOM method. Creates a deep copy of the specified node
// and places it into the target container before the given anchor.
func (d *domainClient) CopyTo(ctx context.Context, args *CopyToArgs) (reply *CopyToReply, err error) {
	reply = new(CopyToReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.copyTo", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.copyTo", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "CopyTo", Err: err}
	}
	return
}

// DescribeNode invokes the DOM method. Describes node given its id, does not
// require domain to be enabled. Does not start tracking any objects, can be
// used for automation.
func (d *domainClient) DescribeNode(ctx context.Context, args *DescribeNodeArgs) (reply *DescribeNodeReply, err error) {
	reply = new(DescribeNodeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.describeNode", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.describeNode", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "DescribeNode", Err: err}
	}
	return
}

// Disable invokes the DOM method. Disables DOM agent for the given page.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "DOM.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "Disable", Err: err}
	}
	return
}

// DiscardSearchResults invokes the DOM method. Discards search results from
// the session with the given id. `getSearchResults` should no longer be called
// for that search.
func (d *domainClient) DiscardSearchResults(ctx context.Context, args *DiscardSearchResultsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.discardSearchResults", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.discardSearchResults", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "DiscardSearchResults", Err: err}
	}
	return
}

// Enable invokes the DOM method. Enables DOM agent for the given page.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "DOM.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "Enable", Err: err}
	}
	return
}

// Focus invokes the DOM method. Focuses the given element.
func (d *domainClient) Focus(ctx context.Context, args *FocusArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.focus", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.focus", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "Focus", Err: err}
	}
	return
}

// GetAttributes invokes the DOM method. Returns attributes for the specified
// node.
func (d *domainClient) GetAttributes(ctx context.Context, args *GetAttributesArgs) (reply *GetAttributesReply, err error) {
	reply = new(GetAttributesReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getAttributes", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getAttributes", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetAttributes", Err: err}
	}
	return
}

// GetBoxModel invokes the DOM method. Returns boxes for the given node.
func (d *domainClient) GetBoxModel(ctx context.Context, args *GetBoxModelArgs) (reply *GetBoxModelReply, err error) {
	reply = new(GetBoxModelReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getBoxModel", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getBoxModel", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetBoxModel", Err: err}
	}
	return
}

// GetContentQuads invokes the DOM method. Returns quads that describe node
// position on the page. This method might return multiple quads for inline
// nodes.
func (d *domainClient) GetContentQuads(ctx context.Context, args *GetContentQuadsArgs) (reply *GetContentQuadsReply, err error) {
	reply = new(GetContentQuadsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getContentQuads", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getContentQuads", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetContentQuads", Err: err}
	}
	return
}

// GetDocument invokes the DOM method. Returns the root DOM node (and
// optionally the subtree) to the caller.
func (d *domainClient) GetDocument(ctx context.Context, args *GetDocumentArgs) (reply *GetDocumentReply, err error) {
	reply = new(GetDocumentReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getDocument", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getDocument", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetDocument", Err: err}
	}
	return
}

// GetFlattenedDocument invokes the DOM method. Returns the root DOM node (and
// optionally the subtree) to the caller.
func (d *domainClient) GetFlattenedDocument(ctx context.Context, args *GetFlattenedDocumentArgs) (reply *GetFlattenedDocumentReply, err error) {
	reply = new(GetFlattenedDocumentReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getFlattenedDocument", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getFlattenedDocument", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetFlattenedDocument", Err: err}
	}
	return
}

// GetNodeForLocation invokes the DOM method. Returns node id at given
// location.
func (d *domainClient) GetNodeForLocation(ctx context.Context, args *GetNodeForLocationArgs) (reply *GetNodeForLocationReply, err error) {
	reply = new(GetNodeForLocationReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getNodeForLocation", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getNodeForLocation", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetNodeForLocation", Err: err}
	}
	return
}

// GetOuterHTML invokes the DOM method. Returns node's HTML markup.
func (d *domainClient) GetOuterHTML(ctx context.Context, args *GetOuterHTMLArgs) (reply *GetOuterHTMLReply, err error) {
	reply = new(GetOuterHTMLReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getOuterHTML", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getOuterHTML", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetOuterHTML", Err: err}
	}
	return
}

// GetRelayoutBoundary invokes the DOM method. Returns the id of the nearest
// ancestor that is a relayout boundary.
func (d *domainClient) GetRelayoutBoundary(ctx context.Context, args *GetRelayoutBoundaryArgs) (reply *GetRelayoutBoundaryReply, err error) {
	reply = new(GetRelayoutBoundaryReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getRelayoutBoundary", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getRelayoutBoundary", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetRelayoutBoundary", Err: err}
	}
	return
}

// GetSearchResults invokes the DOM method. Returns search results from given
// `fromIndex` to given `toIndex` from the search with the given identifier.
func (d *domainClient) GetSearchResults(ctx context.Context, args *GetSearchResultsArgs) (reply *GetSearchResultsReply, err error) {
	reply = new(GetSearchResultsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getSearchResults", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getSearchResults", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetSearchResults", Err: err}
	}
	return
}

// MarkUndoableState invokes the DOM method. Marks last undoable state.
func (d *domainClient) MarkUndoableState(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "DOM.markUndoableState", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "MarkUndoableState", Err: err}
	}
	return
}

// MoveTo invokes the DOM method. Moves node into the new container, places it
// before the given anchor.
func (d *domainClient) MoveTo(ctx context.Context, args *MoveToArgs) (reply *MoveToReply, err error) {
	reply = new(MoveToReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.moveTo", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.moveTo", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "MoveTo", Err: err}
	}
	return
}

// PerformSearch invokes the DOM method. Searches for a given string in the
// DOM tree. Use `getSearchResults` to access search results or `cancelSearch`
// to end this search session.
func (d *domainClient) PerformSearch(ctx context.Context, args *PerformSearchArgs) (reply *PerformSearchReply, err error) {
	reply = new(PerformSearchReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.performSearch", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.performSearch", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "PerformSearch", Err: err}
	}
	return
}

// PushNodeByPathToFrontend invokes the DOM method. Requests that the node is
// sent to the caller given its path. // FIXME, use XPath
func (d *domainClient) PushNodeByPathToFrontend(ctx context.Context, args *PushNodeByPathToFrontendArgs) (reply *PushNodeByPathToFrontendReply, err error) {
	reply = new(PushNodeByPathToFrontendReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.pushNodeByPathToFrontend", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.pushNodeByPathToFrontend", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "PushNodeByPathToFrontend", Err: err}
	}
	return
}

// PushNodesByBackendIdsToFrontend invokes the DOM method. Requests that a
// batch of nodes is sent to the caller given their backend node ids.
func (d *domainClient) PushNodesByBackendIdsToFrontend(ctx context.Context, args *PushNodesByBackendIdsToFrontendArgs) (reply *PushNodesByBackendIdsToFrontendReply, err error) {
	reply = new(PushNodesByBackendIdsToFrontendReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.pushNodesByBackendIdsToFrontend", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.pushNodesByBackendIdsToFrontend", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "PushNodesByBackendIdsToFrontend", Err: err}
	}
	return
}

// QuerySelector invokes the DOM method. Executes `querySelector` on a given
// node.
func (d *domainClient) QuerySelector(ctx context.Context, args *QuerySelectorArgs) (reply *QuerySelectorReply, err error) {
	reply = new(QuerySelectorReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.querySelector", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.querySelector", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "QuerySelector", Err: err}
	}
	return
}

// QuerySelectorAll invokes the DOM method. Executes `querySelectorAll` on a
// given node.
func (d *domainClient) QuerySelectorAll(ctx context.Context, args *QuerySelectorAllArgs) (reply *QuerySelectorAllReply, err error) {
	reply = new(QuerySelectorAllReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.querySelectorAll", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.querySelectorAll", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "QuerySelectorAll", Err: err}
	}
	return
}

// Redo invokes the DOM method. Re-does the last undone action.
func (d *domainClient) Redo(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "DOM.redo", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "Redo", Err: err}
	}
	return
}

// RemoveAttribute invokes the DOM method. Removes attribute with given name
// from an element with given id.
func (d *domainClient) RemoveAttribute(ctx context.Context, args *RemoveAttributeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.removeAttribute", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.removeAttribute", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "RemoveAttribute", Err: err}
	}
	return
}

// RemoveNode invokes the DOM method. Removes node with given id.
func (d *domainClient) RemoveNode(ctx context.Context, args *RemoveNodeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.removeNode", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.removeNode", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "RemoveNode", Err: err}
	}
	return
}

// RequestChildNodes invokes the DOM method. Requests that children of the
// node with given id are returned to the caller in form of `setChildNodes`
// events where not only immediate children are retrieved, but all children
// down to the specified depth.
func (d *domainClient) RequestChildNodes(ctx context.Context, args *RequestChildNodesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.requestChildNodes", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.requestChildNodes", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "RequestChildNodes", Err: err}
	}
	return
}

// RequestNode invokes the DOM method. Requests that the node is sent to the
// caller given the JavaScript node object reference. All nodes that form the
// path from the node to the root are also sent to the client as a series of
// `setChildNodes` notifications.
func (d *domainClient) RequestNode(ctx context.Context, args *RequestNodeArgs) (reply *RequestNodeReply, err error) {
	reply = new(RequestNodeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.requestNode", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.requestNode", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "RequestNode", Err: err}
	}
	return
}

// ResolveNode invokes the DOM method. Resolves the JavaScript node object for
// a given NodeId or BackendNodeId.
func (d *domainClient) ResolveNode(ctx context.Context, args *ResolveNodeArgs) (reply *ResolveNodeReply, err error) {
	reply = new(ResolveNodeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.resolveNode", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.resolveNode", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "ResolveNode", Err: err}
	}
	return
}

// SetAttributeValue invokes the DOM method. Sets attribute for an element
// with given id.
func (d *domainClient) SetAttributeValue(ctx context.Context, args *SetAttributeValueArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.setAttributeValue", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.setAttributeValue", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "SetAttributeValue", Err: err}
	}
	return
}

// SetAttributesAsText invokes the DOM method. Sets attributes on element with
// given id. This method is useful when user edits some existing attribute
// value and types in several attribute name/value pairs.
func (d *domainClient) SetAttributesAsText(ctx context.Context, args *SetAttributesAsTextArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.setAttributesAsText", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.setAttributesAsText", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "SetAttributesAsText", Err: err}
	}
	return
}

// SetFileInputFiles invokes the DOM method. Sets files for the given file
// input element.
func (d *domainClient) SetFileInputFiles(ctx context.Context, args *SetFileInputFilesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.setFileInputFiles", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.setFileInputFiles", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "SetFileInputFiles", Err: err}
	}
	return
}

// SetInspectedNode invokes the DOM method. Enables console to refer to the
// node with given id via $x (see Command Line API for more details $x
// functions).
func (d *domainClient) SetInspectedNode(ctx context.Context, args *SetInspectedNodeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.setInspectedNode", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.setInspectedNode", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "SetInspectedNode", Err: err}
	}
	return
}

// SetNodeName invokes the DOM method. Sets node name for a node with given
// id.
func (d *domainClient) SetNodeName(ctx context.Context, args *SetNodeNameArgs) (reply *SetNodeNameReply, err error) {
	reply = new(SetNodeNameReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.setNodeName", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.setNodeName", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "SetNodeName", Err: err}
	}
	return
}

// SetNodeValue invokes the DOM method. Sets node value for a node with given
// id.
func (d *domainClient) SetNodeValue(ctx context.Context, args *SetNodeValueArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.setNodeValue", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.setNodeValue", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "SetNodeValue", Err: err}
	}
	return
}

// SetOuterHTML invokes the DOM method. Sets node HTML markup, returns new
// node id.
func (d *domainClient) SetOuterHTML(ctx context.Context, args *SetOuterHTMLArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.setOuterHTML", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.setOuterHTML", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "SetOuterHTML", Err: err}
	}
	return
}

// Undo invokes the DOM method. Undoes the last performed action.
func (d *domainClient) Undo(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "DOM.undo", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "Undo", Err: err}
	}
	return
}

// GetFrameOwner invokes the DOM method. Returns iframe node that owns iframe
// with the given domain.
func (d *domainClient) GetFrameOwner(ctx context.Context, args *GetFrameOwnerArgs) (reply *GetFrameOwnerReply, err error) {
	reply = new(GetFrameOwnerReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "DOM.getFrameOwner", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DOM.getFrameOwner", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOM", Op: "GetFrameOwner", Err: err}
	}
	return
}

func (d *domainClient) AttributeModified(ctx context.Context) (AttributeModifiedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.attributeModified", d.conn)
	if err != nil {
		return nil, err
	}
	return &attributeModifiedClient{Stream: s}, nil
}

type attributeModifiedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *attributeModifiedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *attributeModifiedClient) Recv() (*AttributeModifiedReply, error) {
	event := new(AttributeModifiedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "AttributeModified Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) AttributeRemoved(ctx context.Context) (AttributeRemovedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.attributeRemoved", d.conn)
	if err != nil {
		return nil, err
	}
	return &attributeRemovedClient{Stream: s}, nil
}

type attributeRemovedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *attributeRemovedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *attributeRemovedClient) Recv() (*AttributeRemovedReply, error) {
	event := new(AttributeRemovedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "AttributeRemoved Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) CharacterDataModified(ctx context.Context) (CharacterDataModifiedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.characterDataModified", d.conn)
	if err != nil {
		return nil, err
	}
	return &characterDataModifiedClient{Stream: s}, nil
}

type characterDataModifiedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *characterDataModifiedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *characterDataModifiedClient) Recv() (*CharacterDataModifiedReply, error) {
	event := new(CharacterDataModifiedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "CharacterDataModified Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ChildNodeCountUpdated(ctx context.Context) (ChildNodeCountUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.childNodeCountUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &childNodeCountUpdatedClient{Stream: s}, nil
}

type childNodeCountUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *childNodeCountUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *childNodeCountUpdatedClient) Recv() (*ChildNodeCountUpdatedReply, error) {
	event := new(ChildNodeCountUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "ChildNodeCountUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ChildNodeInserted(ctx context.Context) (ChildNodeInsertedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.childNodeInserted", d.conn)
	if err != nil {
		return nil, err
	}
	return &childNodeInsertedClient{Stream: s}, nil
}

type childNodeInsertedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *childNodeInsertedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *childNodeInsertedClient) Recv() (*ChildNodeInsertedReply, error) {
	event := new(ChildNodeInsertedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "ChildNodeInserted Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ChildNodeRemoved(ctx context.Context) (ChildNodeRemovedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.childNodeRemoved", d.conn)
	if err != nil {
		return nil, err
	}
	return &childNodeRemovedClient{Stream: s}, nil
}

type childNodeRemovedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *childNodeRemovedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *childNodeRemovedClient) Recv() (*ChildNodeRemovedReply, error) {
	event := new(ChildNodeRemovedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "ChildNodeRemoved Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) DistributedNodesUpdated(ctx context.Context) (DistributedNodesUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.distributedNodesUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &distributedNodesUpdatedClient{Stream: s}, nil
}

type distributedNodesUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *distributedNodesUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *distributedNodesUpdatedClient) Recv() (*DistributedNodesUpdatedReply, error) {
	event := new(DistributedNodesUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "DistributedNodesUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) DocumentUpdated(ctx context.Context) (DocumentUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.documentUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &documentUpdatedClient{Stream: s}, nil
}

type documentUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *documentUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *documentUpdatedClient) Recv() (*DocumentUpdatedReply, error) {
	event := new(DocumentUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "DocumentUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) InlineStyleInvalidated(ctx context.Context) (InlineStyleInvalidatedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.inlineStyleInvalidated", d.conn)
	if err != nil {
		return nil, err
	}
	return &inlineStyleInvalidatedClient{Stream: s}, nil
}

type inlineStyleInvalidatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *inlineStyleInvalidatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *inlineStyleInvalidatedClient) Recv() (*InlineStyleInvalidatedReply, error) {
	event := new(InlineStyleInvalidatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "InlineStyleInvalidated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) PseudoElementAdded(ctx context.Context) (PseudoElementAddedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.pseudoElementAdded", d.conn)
	if err != nil {
		return nil, err
	}
	return &pseudoElementAddedClient{Stream: s}, nil
}

type pseudoElementAddedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *pseudoElementAddedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *pseudoElementAddedClient) Recv() (*PseudoElementAddedReply, error) {
	event := new(PseudoElementAddedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "PseudoElementAdded Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) PseudoElementRemoved(ctx context.Context) (PseudoElementRemovedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.pseudoElementRemoved", d.conn)
	if err != nil {
		return nil, err
	}
	return &pseudoElementRemovedClient{Stream: s}, nil
}

type pseudoElementRemovedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *pseudoElementRemovedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *pseudoElementRemovedClient) Recv() (*PseudoElementRemovedReply, error) {
	event := new(PseudoElementRemovedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "PseudoElementRemoved Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) SetChildNodes(ctx context.Context) (SetChildNodesClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.setChildNodes", d.conn)
	if err != nil {
		return nil, err
	}
	return &setChildNodesClient{Stream: s}, nil
}

type setChildNodesClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *setChildNodesClient) GetStream() rpcc.Stream { return c.Stream }

func (c *setChildNodesClient) Recv() (*SetChildNodesReply, error) {
	event := new(SetChildNodesReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "SetChildNodes Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ShadowRootPopped(ctx context.Context) (ShadowRootPoppedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.shadowRootPopped", d.conn)
	if err != nil {
		return nil, err
	}
	return &shadowRootPoppedClient{Stream: s}, nil
}

type shadowRootPoppedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *shadowRootPoppedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *shadowRootPoppedClient) Recv() (*ShadowRootPoppedReply, error) {
	event := new(ShadowRootPoppedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "ShadowRootPopped Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ShadowRootPushed(ctx context.Context) (ShadowRootPushedClient, error) {
	s, err := rpcc.NewStream(ctx, "DOM.shadowRootPushed", d.conn)
	if err != nil {
		return nil, err
	}
	return &shadowRootPushedClient{Stream: s}, nil
}

type shadowRootPushedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *shadowRootPushedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *shadowRootPushedClient) Recv() (*ShadowRootPushedReply, error) {
	event := new(ShadowRootPushedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "DOM", Op: "ShadowRootPushed Recv", Err: err}
	}
	return event, nil
}
