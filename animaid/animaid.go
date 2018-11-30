package animaid

type Anime struct {
	episodes    map[int]Episode
	genres      []string
	name        string
	description string
}

type ReleaseState int

const (
	UNRELEASED ReleaseState = iota
	AIRING
	FINISHED
	CANCELED
)

func (anime *Anime) GetAmountOfEpisodes() int {
	return length(anime.episodes)
}

func (anime *Anime) GetAmountOfReleasedEpisodes() int {
	//TODO
	return 0
}

func (anime *Anime) GetAmountOfWatchedEpisodes() int {
	//TODO
	return 0
}

func (anime *Anime) GetReleaseState() ReleaseState {
	//TODO
	return UNRELEASED
}

type Episdode struct {
	released bool
	watched  bool
}
