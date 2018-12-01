package animaid

//AniMaid is the thing that you can interact with in order to get data from
//whatever your provider might be. All AniMaid implementations should be
//interchangeable, since they all have the same functionallity. However,
//any AniMaid implementation can have additional functions, those however
//can not be used in a general way.
type AniMaid interface {
}
