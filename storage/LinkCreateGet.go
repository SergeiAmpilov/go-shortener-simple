package storage

/* external - is the shorted link; internal - native link to the page */

type LinkCreateGet interface {
	CreateLink(internalLink string) (Link, error)
	GetLink(externalLink string) (Link, error)
}
