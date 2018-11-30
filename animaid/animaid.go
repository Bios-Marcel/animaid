//Package animaid contains the interface, datatypes and convenience methods for the AniMaid library.
package animaid

//Anime contains basic information about an Anime and contains all episodes.
type Anime struct {
	episodes    map[int]Episode
	genres      []string
	name        string
	description string
}

//ReleaseState indicates what level of airing completion an Anime already has.
type ReleaseState int

const (
	//Unreleased means that this Anime is Airing at some point in the future
	//,but doesn't have any aired Episode as of now.
	Unreleased ReleaseState = iota
	//Airing means that this Anime is already released, but not every Episode
	//has aired yet.
	Airing
	//Finished means that this Anime has already been released and all of its
	//Episodes have aired already.
	Finished
)

//GetAmountOfEpisodes returns the total amount of episodes this Anime has.
func (anime *Anime) GetAmountOfEpisodes() int {
	return len(anime.episodes)
}

//GetAmountOfAiredEpisodes returns the amount of already aired episodes.
func (anime *Anime) GetAmountOfAiredEpisodes() int {
	var released int
	for _, episode := range anime.episodes {
		if episode.aired {
			released = released + 1
		}
	}

	return released
}

//GetAmountOfWatchedEpisodes returns the amount of already watched episodes.
//This method doesn't care if there are holes in the watchhistory.
func (anime *Anime) GetAmountOfWatchedEpisodes() int {
	var watched int
	for _, episode := range anime.episodes {
		if episode.watched {
			watched = watched + 1
		}
	}

	return watched
}

//GetReleaseState returns Finished if all episodes have already aired.
//If only some episodes have aired, Airing will be returned instead.
//In case none of the above is true, Unreleased will be returned.
func (anime *Anime) GetReleaseState() ReleaseState {
	airedEpisodes := anime.GetAmountOfAiredEpisodes()

	if airedEpisodes <= 0 {
		return Unreleased
	}

	if anime.GetAmountOfEpisodes() != airedEpisodes {
		return Airing
	}

	return Finished
}

//Episode is a single Episode of an Anime.
type Episode struct {
	aired   bool
	watched bool
}

//HasAired indicates wether this Episode has already been aired.
func (episode *Episode) HasAired() bool {
	return episode.aired
}

//IsWatched indicates wether the user has already watched this episode.
func (episode *Episode) IsWatched() bool {
	return episode.watched
}
