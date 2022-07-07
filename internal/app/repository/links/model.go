package links

type Links struct {
	link map[string]string
}

func InitLinksRepository() *Links {
	linksRep := &Links{link: make(map[string]string)}

	return linksRep
}
