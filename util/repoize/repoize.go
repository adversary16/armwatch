package repoize

type Repository struct {
	get func(...interface{}) (interface{}, error)
	set func(...interface{}) (interface{}, error)
}

func Init() {

}

func Get() {

}

func Set() {

}
