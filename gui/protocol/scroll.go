/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

// Scroll indicate scroll behavior in any elements.
// If the element's direction is rtl (right-to-left), then scrollLeft is 0 when the scrollbar is at its rightmost position
// (at the start of the scrolled content), and then increasingly negative as you scroll towards the end of the content.
type Scroll interface {
	// https://developer.mozilla.org/en-US/docs/Web/API/Window/scroll
	Scroll(x, y int, options Scroll_Options)
	// Scrolls to the HTML element with the given id.
	ScrollToID(id string)

	// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
	ScrollWidth() int
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollHeight
	ScrollHeight() int
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
	ScrollTop() int
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
	ScrollLeft() int

	// Returns true if the scrollbars are visible; otherwise, returns false.
	// To hide/diasable/force-show scrollbars use CSS >>
	// 		::-webkit-scrollbar { display: none; }	/* Hide - Chrome/Safari/Webkit */
	// 		scrollbar-width: none;	/* Hide - W3C Candidate Recommendation (Just FireFox) */
	// 		overflow: hidden;		/* Hide and disable scrollbars */
	//		overflow-y: hidden; 	/* Hide and disable vertical scrollbar */
	//		overflow-x: hidden; 	/* Hide and disable horizontal scrollbar */
	// 		overflow: scroll;   	/* Show scrollbars */
	// 		overflow-y: scroll; 	/* Show vertical scrollbar */
	// 		overflow-x: scroll;		/* Show horizontal scrollbar */
	Scrollbars() (x, y bool)

	// EventTarget[Scroll]
}

type Scroll_Options struct {
	Behavior Scroll_Behavior
}

// Specifies the scrolling animate behavior
type Scroll_Behavior uint8

const (
	Scroll_Behavior_Auto Scroll_Behavior = iota
	Scroll_Behavior_Smoothly
	Scroll_Behavior_Instantly // instantly in a single jump
)
