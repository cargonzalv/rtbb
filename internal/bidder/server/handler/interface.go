package handler

import "github.com/adgear/rtb-bidder/internal/bidder"

// Router is proxy type of bidder.Router required by wire resolver.
type Router bidder.Router
